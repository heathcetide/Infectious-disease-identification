// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/pkg/common/response"
	"gorm.io/gorm"
)

// CheckInDetailVO aggregates a check-in with scores for CDC review.
type CheckInDetailVO struct {
	Status           models.DailyHealthStatus  `json:"status"`
	User             *models.User              `json:"user,omitempty"`
	Profile          *models.EmployeeProfile   `json:"profile,omitempty"`
	DiseaseScoring   []models.DiseaseScoring   `json:"diseaseScoringList"`
	AIDiseaseScoring []models.AIDiseaseScoring `json:"aiDiseaseScoringList,omitempty"`
	Hybrid           *HybridMeta               `json:"hybrid,omitempty"`
}

// GetCheckInDetail loads check-in by status ID with user and scores.
func GetCheckInDetail(db *gorm.DB, statusID uint) (*CheckInDetailVO, error) {
	var status models.DailyHealthStatus
	if err := db.First(&status, statusID).Error; err != nil {
		return nil, err
	}
	vo := &CheckInDetailVO{Status: status}
	if u, err := (models.User{}).FindByID(db, status.UserID); err == nil {
		vo.User = u
	}
	var profile models.EmployeeProfile
	if err := db.Where("user_id = ?", status.UserID).First(&profile).Error; err == nil {
		vo.Profile = &profile
	}
	day := status.CheckInDate
	_ = db.Where("status_id = ? AND symptom_time = ?", statusID, day).Find(&vo.DiseaseScoring).Error
	_ = db.Where("status_id = ? AND symptom_time = ?", statusID, day).Find(&vo.AIDiseaseScoring).Error
	disagree := status.ScoringDisagreement != nil && *status.ScoringDisagreement
	confirmed := status.ScoringConfirmedAt != nil
	if disagree || len(vo.AIDiseaseScoring) > 0 {
		vo.Hybrid = &HybridMeta{
			IsAI:              true,
			Disagreement:      disagree,
			NeedsHumanConfirm: disagree && !confirmed, // 仅管理端：需疾控复核
			Confirmed:         confirmed,
			Message:           statusHybridMessage(disagree, confirmed),
		}
	}
	return vo, nil
}

// CheckInLocationVO is location-focused check-in info (flat for admin drawer).
type CheckInLocationVO = CheckInLocationFlatVO

// GetCheckInLocation returns flat personal + location fields for the admin drawer.
func GetCheckInLocation(db *gorm.DB, statusID uint) (*CheckInLocationVO, error) {
	var status models.DailyHealthStatus
	if err := db.First(&status, statusID).Error; err != nil {
		return nil, err
	}
	vo := &CheckInLocationVO{
		Longitude:    status.Longitude,
		Latitude:     status.Latitude,
		Province:     status.Province,
		City:         status.City,
		County:       status.County,
		LocationName: status.LocationName,
		CheckInDate:  dateParts(status.CheckInDate),
		Address:      joinLocation(status.Province, status.City, status.County, status.LocationName),
	}
	if u, err := (models.User{}).FindByID(db, status.UserID); err == nil {
		vo.Name = u.Name
		vo.UserType = roleDisplayName(u.Role)
	}
	var profile models.EmployeeProfile
	if err := db.Where("user_id = ?", status.UserID).First(&profile).Error; err == nil {
		vo.Gender = profile.Gender
		vo.Age = profile.Age
		vo.Ethnicity = profile.Ethnicity
		vo.Department = profile.Department
	}
	return vo, nil
}

func dateParts(t time.Time) []int {
	if t.IsZero() {
		return nil
	}
	return []int{t.Year(), int(t.Month()), t.Day()}
}

func joinLocation(parts ...string) string {
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return strings.Join(out, "")
}

// PageCheckInsJoined paginates check-ins with user/profile columns for the admin table.
func PageCheckInsJoined(db *gorm.DB, q PageQuery) (*response.Page, error) {
	if db == nil {
		return response.NewPage([]CheckInListItem{}, 0, 1, 10), nil
	}
	page, size := q.Normalize()
	query := db.Table("daily_health_statuses AS d").
		Joins("LEFT JOIN users u ON u.id = d.user_id").
		Joins("LEFT JOIN employee_profiles p ON p.user_id = d.user_id")

	if q.IsHealth == "true" || q.IsHealth == "1" {
		query = query.Where("d.is_health = ?", true)
	} else if q.IsHealth == "false" || q.IsHealth == "0" {
		query = query.Where("d.is_health = ?", false)
	}
	if q.Name != "" {
		query = query.Where("u.name LIKE ?", "%"+q.Name+"%")
	}
	if q.PhoneNumber != "" {
		like := "%" + q.PhoneNumber + "%"
		query = query.Where("u.phone_number LIKE ? OR u.phone LIKE ?", like, like)
	}
	if q.Department != "" {
		query = query.Where("p.department LIKE ?", "%"+q.Department+"%")
	}
	if q.SpecificOccupation != "" {
		query = query.Where("p.specific_occupation LIKE ?", "%"+q.SpecificOccupation+"%")
	}
	if q.CheckInDateBegin != "" {
		if t, err := ParseDate(q.CheckInDateBegin); err == nil {
			query = query.Where("d.check_in_date >= ?", t)
		}
	}
	if q.CheckInDateEnd != "" {
		if t, err := ParseDate(q.CheckInDateEnd); err == nil {
			query = query.Where("d.check_in_date < ?", t.AddDate(0, 0, 1))
		}
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	type row struct {
		StatusID           uint      `gorm:"column:status_id"`
		UserID             uint      `gorm:"column:user_id"`
		CheckInDate        time.Time `gorm:"column:check_in_date"`
		IsHealth           *bool     `gorm:"column:is_health"`
		DiseaseTypeName    string    `gorm:"column:disease_type_name"`
		Longitude          float64   `gorm:"column:longitude"`
		Latitude           float64   `gorm:"column:latitude"`
		Province           string    `gorm:"column:province"`
		City               string    `gorm:"column:city"`
		County             string    `gorm:"column:county"`
		LocationName       string    `gorm:"column:location_name"`
		Name               string    `gorm:"column:name"`
		PhoneNumber        string    `gorm:"column:phone_number"`
		Phone              string    `gorm:"column:phone"`
		Role               string    `gorm:"column:role"`
		Gender             string    `gorm:"column:gender"`
		Age                int       `gorm:"column:age"`
		Department         string    `gorm:"column:department"`
		SpecificOccupation string    `gorm:"column:specific_occupation"`
	}
	var rows []row
	if err := query.
		Select(`d.status_id, d.user_id, d.check_in_date, d.is_health, d.disease_type_name,
			d.longitude, d.latitude, d.province, d.city, d.county, d.location_name,
			u.name, u.phone_number, u.phone, u.role,
			p.gender, p.age, p.department, p.specific_occupation`).
		Order("d.status_id DESC").
		Offset((page - 1) * size).Limit(size).
		Scan(&rows).Error; err != nil {
		return nil, err
	}

	list := make([]CheckInListItem, 0, len(rows))
	for _, r := range rows {
		phone := r.PhoneNumber
		if phone == "" {
			phone = r.Phone
		}
		item := CheckInListItem{
			StatusID:           r.StatusID,
			UserID:             r.UserID,
			UserType:           roleDisplayName(r.Role),
			Name:               r.Name,
			PhoneNumber:        phone,
			Gender:             r.Gender,
			Age:                r.Age,
			Department:         r.Department,
			SpecificOccupation: r.SpecificOccupation,
			IsHealth:           r.IsHealth,
			LocationName:       displayLocation(r.Province, r.City, r.County, r.LocationName),
			Province:           r.Province,
			City:               r.City,
			County:             r.County,
			CheckInDate:        dateParts(r.CheckInDate),
			DiseaseTypeName:    r.DiseaseTypeName,
			Longitude:          r.Longitude,
			Latitude:           r.Latitude,
		}
		list = append(list, item)
	}
	return response.NewPage(list, total, page, size), nil
}

func displayLocation(province, city, county, name string) string {
	if full := joinLocation(province, city, county, name); full != "" {
		return full
	}
	return name
}

// ExportCheckInsCSV exports all daily health rows.
func ExportCheckInsCSV(db *gorm.DB) ([]byte, error) {
	var list []models.DailyHealthStatus
	if err := db.Order("status_id ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	_ = w.Write([]string{"statusId", "userId", "checkInDate", "isHealth", "province", "city", "county", "locationName", "longitude", "latitude"})
	for _, s := range list {
		health := ""
		if s.IsHealth != nil {
			health = strconv.FormatBool(*s.IsHealth)
		}
		_ = w.Write([]string{
			strconv.FormatUint(uint64(s.StatusID), 10),
			strconv.FormatUint(uint64(s.UserID), 10),
			s.CheckInDate.Format("2006-01-02"),
			health, s.Province, s.City, s.County, s.LocationName,
			strconv.FormatFloat(s.Longitude, 'f', -1, 64),
			strconv.FormatFloat(s.Latitude, 'f', -1, 64),
		})
	}
	w.Flush()
	return buf.Bytes(), w.Error()
}

// CDCStatisticsData matches Java StatisticsDataVO core fields.
type CDCStatisticsData struct {
	UserNumber        int64 `json:"userNumber"`
	CheckInNumber     int64 `json:"checkInNumber"`
	UserHealthNumber  int64 `json:"userHealthNumber"`
	UserDiseaseNumber int64 `json:"userDiseaseNumber"`
	CheckInToday      int64 `json:"checkInToday"`
}

// QueryCDCStatisticsData loads dashboard statistics for a day (empty date = today).
func QueryCDCStatisticsData(db *gorm.DB, date string) (*CDCStatisticsData, error) {
	dayStart, dayEnd := todayDateRange()
	if strings.TrimSpace(date) != "" {
		if t, err := time.ParseInLocation("2006-01-02", strings.TrimSpace(date), time.Local); err == nil {
			dayStart = dateOnly(t)
			dayEnd = dayStart.AddDate(0, 0, 1)
		}
	}
	var out CDCStatisticsData
	if err := db.Model(&models.User{}).Where("status = ?", models.UserStatusActive).Count(&out.UserNumber).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&models.DailyHealthStatus{}).Count(&out.CheckInNumber).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&models.DailyHealthStatus{}).
		Where("check_in_date >= ? AND check_in_date < ?", dayStart, dayEnd).
		Count(&out.CheckInToday).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&models.DailyHealthStatus{}).
		Where("check_in_date >= ? AND check_in_date < ? AND is_health = ?", dayStart, dayEnd, true).
		Count(&out.UserHealthNumber).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&models.DailyHealthStatus{}).
		Where("check_in_date >= ? AND check_in_date < ? AND is_health = ?", dayStart, dayEnd, false).
		Count(&out.UserDiseaseNumber).Error; err != nil {
		return nil, err
	}
	return &out, nil
}

// DiseaseNameCount is a disease label with count/score sum.
type DiseaseNameCount struct {
	DiseaseTypeName            string  `json:"diseaseTypeName"`
	DiagnosisDiseaseTypeName   string  `json:"diagnosisDiseaseTypeName"`
	Count                      int64   `json:"count"`
	TotalScore                 float64 `json:"totalScore"`
}

func normalizeScreenDiseaseName(name string) string {
	name = strings.TrimSpace(name)
	switch name {
	case "", "健康":
		return "健康"
	case "高原反应相关不适", "上呼吸道感染":
		return "流感"
	case "流行性出血热":
		return "流行性热出血"
	default:
		return name
	}
}

// CDCDiseaseToday aggregates check-in disease labels for a day (incl. healthy).
// date is optional yyyy-MM-dd; empty means today.
func CDCDiseaseToday(db *gorm.DB, date string) ([]DiseaseNameCount, error) {
	todayStart, todayEnd := todayDateRange()
	if strings.TrimSpace(date) != "" {
		if t, err := time.ParseInLocation("2006-01-02", strings.TrimSpace(date), time.Local); err == nil {
			todayStart = dateOnly(t)
			todayEnd = todayStart.AddDate(0, 0, 1)
		}
	}
	var statuses []models.DailyHealthStatus
	if err := db.Where("check_in_date >= ? AND check_in_date < ?", todayStart, todayEnd).
		Find(&statuses).Error; err != nil {
		return nil, err
	}
	counts := map[string]int64{}
	for _, s := range statuses {
		label := "健康"
		if s.IsHealth != nil && !*s.IsHealth {
			label = normalizeScreenDiseaseName(s.DiseaseTypeName)
			if label == "健康" {
				label = "流感"
			}
		}
		counts[label]++
	}
	// Enrich with scoring totals when present.
	scoreBy := map[string]float64{}
	type scoreRow struct {
		Name  string
		Score float64
	}
	var srows []scoreRow
	_ = db.Model(&models.DiseaseScoring{}).
		Select("disease_type_name as name, sum(score) as score").
		Where("symptom_time >= ? AND symptom_time < ?", todayStart, todayEnd).
		Group("disease_type_name").
		Scan(&srows).Error
	for _, r := range srows {
		scoreBy[normalizeScreenDiseaseName(r.Name)] += r.Score
	}

	out := make([]DiseaseNameCount, 0, len(counts))
	for name, cnt := range counts {
		out = append(out, DiseaseNameCount{
			DiseaseTypeName:          name,
			DiagnosisDiseaseTypeName: name,
			Count:                    cnt,
			TotalScore:               scoreBy[name],
		})
	}
	return out, nil
}

// UserStation is a map point for today's check-ins.
type UserStation struct {
	UserID          uint    `json:"userId,string"`
	StatusID        uint    `json:"statusId"`
	Name            string  `json:"name"`
	PhoneNumber     string  `json:"phoneNumber"`
	Longitude       float64 `json:"longitude"`
	Latitude        float64 `json:"latitude"`
	Province        string  `json:"province,omitempty"`
	City            string  `json:"city,omitempty"`
	County          string  `json:"county,omitempty"`
	LocationName    string  `json:"locationName"`
	DiseaseTypeName string  `json:"diseaseTypeName,omitempty"`
	MaxScore        float64 `json:"maxScore"`
	IsHighRisk      bool    `json:"isHighRisk"`
}

// RegionAlert summarizes clustered case density for map early-warning overlays.
type RegionAlert struct {
	RegionKey     string  `json:"regionKey"`
	RegionName    string  `json:"regionName"`
	Province      string  `json:"province,omitempty"`
	City          string  `json:"city,omitempty"`
	County        string  `json:"county,omitempty"`
	Longitude     float64 `json:"longitude"`
	Latitude      float64 `json:"latitude"`
	CaseCount     int     `json:"caseCount"`
	HighRiskCount int     `json:"highRiskCount"`
	TopDisease    string  `json:"topDisease,omitempty"`
	Level         string  `json:"level"` // watch | warning | critical
	Message       string  `json:"message"`
}

// UserStationsMapDTO is the map payload: raw points + regional alerts.
type UserStationsMapDTO struct {
	Stations     []UserStation `json:"stations"`
	RegionAlerts []RegionAlert `json:"regionAlerts"`
}

const regionHighRiskScore = 80.0

func regionDisplayName(province, city, county, location string) (key, name string) {
	county = strings.TrimSpace(county)
	city = strings.TrimSpace(city)
	province = strings.TrimSpace(province)
	location = strings.TrimSpace(location)
	// Aggregate primarily at city/prefecture level so sparse county points still raise alerts.
	switch {
	case city != "" && province != "":
		return province + "|" + city, province + " · " + city
	case city != "":
		return city, city
	case county != "" && province != "":
		return province + "|" + county, province + " · " + county
	case province != "":
		return province, province
	case location != "":
		return "loc|" + location, location
	default:
		return "unknown", "未知区域"
	}
}

func buildRegionAlerts(stations []UserStation) []RegionAlert {
	type agg struct {
		Province, City, County string
		LngSum, LatSum         float64
		Count, HighRisk        int
		DiseaseHits            map[string]int
	}
	by := map[string]*agg{}
	for _, s := range stations {
		if s.Longitude == 0 && s.Latitude == 0 {
			continue
		}
		key, _ := regionDisplayName(s.Province, s.City, s.County, s.LocationName)
		a := by[key]
		if a == nil {
			a = &agg{
				Province: s.Province, City: s.City, County: s.County,
				DiseaseHits: map[string]int{},
			}
			by[key] = a
		}
		a.Count++
		a.LngSum += s.Longitude
		a.LatSum += s.Latitude
		if s.IsHighRisk {
			a.HighRisk++
		}
		disease := strings.TrimSpace(s.DiseaseTypeName)
		if disease != "" {
			a.DiseaseHits[disease]++
		}
	}
	out := make([]RegionAlert, 0, len(by))
	for key, a := range by {
		if a.Count < 1 {
			continue
		}
		_, name := regionDisplayName(a.Province, a.City, a.County, "")
		if name == "未知区域" && key != "unknown" {
			name = key
		}
		topDisease := ""
		topHits := 0
		for d, n := range a.DiseaseHits {
			if n > topHits {
				topHits = n
				topDisease = d
			}
		}
		level := ""
		switch {
		case a.HighRisk >= 5 || a.Count >= 20:
			level = "critical"
		case a.HighRisk >= 2 || a.Count >= 10:
			level = "warning"
		case a.HighRisk >= 1 || a.Count >= 5:
			level = "watch"
		default:
			continue // below alert threshold
		}
		msg := fmt.Sprintf("%s当日打卡 %d 例", name, a.Count)
		if a.HighRisk > 0 {
			msg += fmt.Sprintf("，高风险 %d 例", a.HighRisk)
		}
		if topDisease != "" {
			msg += fmt.Sprintf("，主要病种「%s」", topDisease)
		}
		switch level {
		case "critical":
			msg += "，建议启动属地重点监测与现场核查"
		case "warning":
			msg += "，建议加强属地巡查与复核"
		default:
			msg += "，请持续关注"
		}
		out = append(out, RegionAlert{
			RegionKey: key, RegionName: name,
			Province: a.Province, City: a.City, County: a.County,
			Longitude: a.LngSum / float64(a.Count), Latitude: a.LatSum / float64(a.Count),
			CaseCount: a.Count, HighRiskCount: a.HighRisk,
			TopDisease: topDisease, Level: level, Message: msg,
		})
	}
	sort.Slice(out, func(i, j int) bool {
		rank := map[string]int{"critical": 3, "warning": 2, "watch": 1}
		if rank[out[i].Level] != rank[out[j].Level] {
			return rank[out[i].Level] > rank[out[j].Level]
		}
		if out[i].HighRiskCount != out[j].HighRiskCount {
			return out[i].HighRiskCount > out[j].HighRiskCount
		}
		return out[i].CaseCount > out[j].CaseCount
	})
	return out
}

// CDCUserStations returns check-in locations + regional alerts for a day.
func CDCUserStations(db *gorm.DB, date string) (*UserStationsMapDTO, error) {
	start, end := todayDateRange()
	if date != "" {
		if t, err := ParseDate(date); err == nil {
			start = dateOnly(t)
			end = start.AddDate(0, 0, 1)
		}
	}
	var statuses []models.DailyHealthStatus
	if err := db.Where("check_in_date >= ? AND check_in_date < ?", start, end).
		Find(&statuses).Error; err != nil {
		return nil, err
	}
	statusIDs := make([]uint, 0, len(statuses))
	for _, s := range statuses {
		statusIDs = append(statusIDs, s.StatusID)
	}
	maxScoreBy := map[uint]float64{}
	diseaseBy := map[uint]string{}
	if len(statusIDs) > 0 {
		type scoreRow struct {
			StatusID        uint
			DiseaseTypeName string
			Score           float64
		}
		var rows []scoreRow
		_ = db.Model(&models.DiseaseScoring{}).
			Select("status_id", "disease_type_name", "score").
			Where("status_id IN ?", statusIDs).
			Find(&rows).Error
		for _, r := range rows {
			if r.Score >= maxScoreBy[r.StatusID] {
				maxScoreBy[r.StatusID] = r.Score
				if name := normalizeScreenDiseaseName(r.DiseaseTypeName); name != "" {
					diseaseBy[r.StatusID] = name
				}
			}
		}
		var aiRows []scoreRow
		_ = db.Model(&models.AIDiseaseScoring{}).
			Select("status_id", "disease_type_name", "score").
			Where("status_id IN ?", statusIDs).
			Find(&aiRows).Error
		for _, r := range aiRows {
			if r.Score >= maxScoreBy[r.StatusID] {
				maxScoreBy[r.StatusID] = r.Score
				if name := normalizeScreenDiseaseName(r.DiseaseTypeName); name != "" {
					diseaseBy[r.StatusID] = name
				}
			}
		}
	}

	userCache := map[uint]*models.User{}
	out := make([]UserStation, 0, len(statuses))
	for _, s := range statuses {
		if s.Longitude == 0 && s.Latitude == 0 {
			continue
		}
		st := UserStation{
			UserID: s.UserID, StatusID: s.StatusID,
			Longitude: s.Longitude, Latitude: s.Latitude,
			Province: s.Province, City: s.City, County: s.County,
			LocationName: s.LocationName,
			DiseaseTypeName: normalizeScreenDiseaseName(s.DiseaseTypeName),
			MaxScore:        maxScoreBy[s.StatusID],
		}
		if st.DiseaseTypeName == "" {
			st.DiseaseTypeName = diseaseBy[s.StatusID]
		}
		st.IsHighRisk = st.MaxScore >= regionHighRiskScore
		u := userCache[s.UserID]
		if u == nil {
			if found, err := (models.User{}).FindByID(db, s.UserID); err == nil {
				userCache[s.UserID] = found
				u = found
			}
		}
		if u != nil {
			st.Name = u.Name
			st.PhoneNumber = userPhoneField(u)
		}
		out = append(out, st)
	}
	return &UserStationsMapDTO{
		Stations:     out,
		RegionAlerts: buildRegionAlerts(out),
	}, nil
}

func userPhoneField(u *models.User) string {
	if u.PhoneNumber != "" {
		return u.PhoneNumber
	}
	return u.Phone
}

// CheckInSearchItem is a search result row.
type CheckInSearchItem struct {
	StatusID                   uint      `json:"statusId"`
	UserID                     uint      `json:"userId"`
	Name                       string    `json:"name"`
	PhoneNumber                string    `json:"phoneNumber"`
	Gender                     string    `json:"gender"`
	Department                 string    `json:"department"`
	CheckInDate                time.Time `json:"checkInDate"`
	IsHealth                   *bool     `json:"isHealth"`
	DiagnosisDiseaseTypeName   string    `json:"diagnosisDiseaseTypeName"`
	DiseaseTypeName            string    `json:"diseaseTypeName"`
}

// CDCCheckInSearch finds recent check-ins matching q (phone or name).
func CDCCheckInSearch(db *gorm.DB, q string, limit int) ([]CheckInSearchItem, error) {
	q = strings.TrimSpace(q)
	if limit < 1 {
		limit = 20
	}
	query := db.Model(&models.DailyHealthStatus{}).Order("status_id DESC").Limit(limit)
	if q != "" {
		like := "%" + q + "%"
		var userIDs []uint
		_ = db.Model(&models.User{}).
			Where("name LIKE ? OR phone_number LIKE ? OR phone LIKE ?", like, like, like).
			Pluck("id", &userIDs)
		if len(userIDs) > 0 {
			query = query.Where("user_id IN ?", userIDs)
		} else {
			return []CheckInSearchItem{}, nil
		}
	}
	var statuses []models.DailyHealthStatus
	if err := query.Find(&statuses).Error; err != nil {
		return nil, err
	}
	out := make([]CheckInSearchItem, 0, len(statuses))
	for _, s := range statuses {
		label := normalizeScreenDiseaseName(s.DiseaseTypeName)
		item := CheckInSearchItem{
			StatusID: s.StatusID, UserID: s.UserID, CheckInDate: s.CheckInDate, IsHealth: s.IsHealth,
			DiseaseTypeName: label, DiagnosisDiseaseTypeName: label,
		}
		if u, err := (models.User{}).FindByID(db, s.UserID); err == nil {
			item.Name = u.Name
			item.PhoneNumber = userPhoneField(u)
		}
		var profile models.EmployeeProfile
		if err := db.Where("user_id = ?", s.UserID).First(&profile).Error; err == nil {
			item.Gender = profile.Gender
			item.Department = profile.Department
		}
		out = append(out, item)
	}
	return out, nil
}

// CheckInDailySeriesRequest is the POST body for daily series.
type CheckInDailySeriesRequest struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

// CheckInDailyNumber is count per day.
type CheckInDailyNumber struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}

// CDCCheckInDailySeries counts check-ins per day in range.
func CDCCheckInDailySeries(db *gorm.DB, req CheckInDailySeriesRequest) ([]CheckInDailyNumber, error) {
	start, err := ParseDate(req.StartDate)
	if err != nil {
		return nil, ErrInvalidDate
	}
	end, err := ParseDate(req.EndDate)
	if err != nil {
		return nil, ErrInvalidDate
	}
	end = end.AddDate(0, 0, 1)
	var list []CheckInDailyNumber
	for d := start; d.Before(end); d = d.AddDate(0, 0, 1) {
		dayEnd := d.AddDate(0, 0, 1)
		var cnt int64
		if err := db.Model(&models.DailyHealthStatus{}).
			Where("check_in_date >= ? AND check_in_date < ?", d, dayEnd).
			Count(&cnt).Error; err != nil {
			return nil, err
		}
		list = append(list, CheckInDailyNumber{Date: d.Format("2006-01-02"), Count: cnt})
	}
	return list, nil
}

// ImportantUserInfo highlights users with unhealthy check-in today.
type ImportantUserInfo struct {
	UserID                     uint   `json:"userId"`
	Name                       string `json:"name"`
	PhoneNumber                string `json:"phoneNumber"`
	Gender                     string `json:"gender"`
	Department                 string `json:"department"`
	IsHealth                   *bool  `json:"isHealth"`
	DiagnosisDiseaseTypeName   string `json:"diagnosisDiseaseTypeName"`
	Reason                     string `json:"reason"`
}

// CDCImportantUsers lists users flagged today.
func CDCImportantUsers(db *gorm.DB) ([]ImportantUserInfo, error) {
	todayStart, todayEnd := todayDateRange()
	var statuses []models.DailyHealthStatus
	if err := db.Where("check_in_date >= ? AND check_in_date < ? AND is_health = ?",
		todayStart, todayEnd, false).Find(&statuses).Error; err != nil {
		return nil, err
	}
	out := make([]ImportantUserInfo, 0, len(statuses))
	for _, s := range statuses {
		h := false
		info := ImportantUserInfo{
			UserID:                   s.UserID,
			IsHealth:                 &h,
			DiagnosisDiseaseTypeName: normalizeScreenDiseaseName(s.DiseaseTypeName),
			Reason:                   "今日症状异常打卡",
		}
		if info.DiagnosisDiseaseTypeName == "健康" || info.DiagnosisDiseaseTypeName == "" {
			info.DiagnosisDiseaseTypeName = "流感"
		}
		if u, err := (models.User{}).FindByID(db, s.UserID); err == nil {
			info.Name = u.Name
			info.PhoneNumber = userPhoneField(u)
		}
		var profile models.EmployeeProfile
		if err := db.Where("user_id = ?", s.UserID).First(&profile).Error; err == nil {
			info.Gender = profile.Gender
			info.Department = profile.Department
		}
		out = append(out, info)
	}
	return out, nil
}

// DiseaseStatisticsQuery filters disease time-series for the disease-report chart.
type DiseaseStatisticsQuery struct {
	DiseaseTypeName string
	DiseaseList     []string
	StartDate       string
	EndDate         string
	SourceType      string // 自动诊断 | 医院诊断
	Granularity     string // day | hour
}

// DiseaseTimeSeriesPoint matches admin disease-report.vue series points.
type DiseaseTimeSeriesPoint struct {
	Key   []string `json:"key"`
	Value int64    `json:"value"`
}

// CDCDiseaseStatistics returns map[diseaseName][]points for the report chart.
func CDCDiseaseStatistics(db *gorm.DB, q DiseaseStatisticsQuery) (map[string][]DiseaseTimeSeriesPoint, error) {
	out := map[string][]DiseaseTimeSeriesPoint{}
	if db == nil {
		return out, nil
	}
	start, end, err := parseDateRange(q.StartDate, q.EndDate)
	if err != nil {
		return nil, ErrInvalidDate
	}
	names := q.DiseaseList
	if len(names) == 0 && q.DiseaseTypeName != "" {
		names = []string{q.DiseaseTypeName}
	}
	for _, n := range names {
		out[n] = []DiseaseTimeSeriesPoint{}
	}

	sameDay := start.Format("2006-01-02") == end.Format("2006-01-02")
	useHour := q.Granularity == "hour" || (q.Granularity == "" && sameDay)

	type aggRow struct {
		Name string
		Bucket string
		Cnt  int64
	}
	var rows []aggRow
	src := strings.TrimSpace(q.SourceType)
	if src == "医院诊断" {
		if useHour {
			err = db.Raw(`
				SELECT disease_type AS name,
					printf('%02d', CAST(strftime('%H', submission_time) AS INTEGER) / 6 * 6) AS bucket,
					COUNT(*) AS cnt
				FROM diagnosis_results
				WHERE submission_time >= ? AND submission_time <= ?
				GROUP BY disease_type, bucket
			`, start, end).Scan(&rows).Error
		} else {
			err = db.Raw(`
				SELECT disease_type AS name,
					strftime('%Y-%m-%d', submission_time) AS bucket,
					COUNT(*) AS cnt
				FROM diagnosis_results
				WHERE submission_time >= ? AND submission_time <= ?
				GROUP BY disease_type, bucket
			`, start, end).Scan(&rows).Error
		}
	} else {
		// 自动诊断：优先 disease_scoring；若无则回退到打卡预测病种
		if useHour {
			// symptom_time is date-only; join check-in time when available.
			err = db.Raw(`
				SELECT ds.disease_type_name AS name,
					printf('%02d', CAST(strftime('%H', COALESCE(dhs.check_in_time, ds.symptom_time)) AS INTEGER) / 6 * 6) AS bucket,
					COUNT(*) AS cnt
				FROM disease_scoring ds
				LEFT JOIN daily_health_statuses dhs ON dhs.status_id = ds.status_id
				WHERE ds.symptom_time >= ? AND ds.symptom_time <= ?
				GROUP BY ds.disease_type_name, bucket
			`, start, end).Scan(&rows).Error
		} else {
			err = db.Raw(`
				SELECT disease_type_name AS name,
					strftime('%Y-%m-%d', symptom_time) AS bucket,
					COUNT(*) AS cnt
				FROM disease_scoring
				WHERE symptom_time >= ? AND symptom_time <= ?
				GROUP BY disease_type_name, bucket
			`, start, end).Scan(&rows).Error
		}
		if err == nil && len(rows) == 0 {
			if useHour {
				err = db.Raw(`
					SELECT disease_type_name AS name,
						printf('%02d', CAST(strftime('%H', check_in_time) AS INTEGER) / 6 * 6) AS bucket,
						COUNT(*) AS cnt
					FROM daily_health_statuses
					WHERE check_in_date >= ? AND check_in_date <= ? AND is_health = 0 AND disease_type_name <> ''
					GROUP BY disease_type_name, bucket
				`, start, end).Scan(&rows).Error
			} else {
				err = db.Raw(`
					SELECT disease_type_name AS name,
						strftime('%Y-%m-%d', check_in_date) AS bucket,
						COUNT(*) AS cnt
					FROM daily_health_statuses
					WHERE check_in_date >= ? AND check_in_date <= ? AND is_health = 0 AND disease_type_name <> ''
					GROUP BY disease_type_name, bucket
				`, start, end).Scan(&rows).Error
			}
		}
	}
	if err != nil {
		return nil, err
	}

	hourLabel := map[string]string{
		"00": "00-06时", "06": "06-12时", "12": "12-18时", "18": "18-24时",
	}
	for _, r := range rows {
		if len(names) > 0 {
			found := false
			for _, n := range names {
				if n == r.Name {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}
		var key []string
		if useHour {
			lab := hourLabel[r.Bucket]
			if lab == "" {
				lab = r.Bucket
			}
			key = []string{lab}
		} else {
			parts := strings.Split(r.Bucket, "-")
			key = parts
		}
		out[r.Name] = append(out[r.Name], DiseaseTimeSeriesPoint{Key: key, Value: r.Cnt})
	}
	return out, nil
}

// WorkEnvironmentListItem joins employee name for CDC work-env table.
type WorkEnvironmentListItem struct {
	models.WorkEnvironmentInfo
	Name string `json:"name"`
}

// PageWorkEnvironmentsJoined paginates work environments with user name.
func PageWorkEnvironmentsJoined(db *gorm.DB, q PageQuery) (*response.Page, error) {
	if db == nil {
		return response.NewPage([]WorkEnvironmentListItem{}, 0, 1, 10), nil
	}
	page, size := q.Normalize()
	query := db.Table("work_environment_info AS w").
		Joins("LEFT JOIN users u ON u.id = w.user_id")
	if q.WorkStationName != "" {
		query = query.Where("w.work_station_name LIKE ?", "%"+q.WorkStationName+"%")
	}
	if q.Keyword != "" {
		query = query.Where("w.work_station_name LIKE ? OR u.name LIKE ?", "%"+q.Keyword+"%", "%"+q.Keyword+"%")
	}
	if q.BeginYearMonth != "" {
		query = query.Where("w.year_month >= ?", q.BeginYearMonth)
	}
	if q.EndYearMonth != "" {
		query = query.Where("w.year_month <= ?", q.EndYearMonth)
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	type row struct {
		models.WorkEnvironmentInfo
		Name string `gorm:"column:name"`
	}
	var rows []row
	if err := query.Select("w.*, u.name AS name").
		Order("w.environment_id DESC").
		Offset((page - 1) * size).Limit(size).
		Scan(&rows).Error; err != nil {
		return nil, err
	}
	list := make([]WorkEnvironmentListItem, 0, len(rows))
	for _, r := range rows {
		list = append(list, WorkEnvironmentListItem{
			WorkEnvironmentInfo: r.WorkEnvironmentInfo,
			Name:                r.Name,
		})
	}
	return response.NewPage(list, total, page, size), nil
}

// GetWorkEnvironmentByID loads work environment row.
func GetWorkEnvironmentByID(db *gorm.DB, id uint) (*models.WorkEnvironmentInfo, error) {
	var info models.WorkEnvironmentInfo
	if err := db.First(&info, id).Error; err != nil {
		return nil, err
	}
	return &info, nil
}

// ExportWorkEnvironmentsCSV exports work environment rows.
func ExportWorkEnvironmentsCSV(db *gorm.DB) ([]byte, error) {
	var list []models.WorkEnvironmentInfo
	if err := db.Order("environment_id ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	_ = w.Write([]string{"environmentId", "userId", "yearMonth", "workStationName"})
	for _, e := range list {
		_ = w.Write([]string{
			strconv.FormatUint(uint64(e.EnvironmentID), 10),
			strconv.FormatUint(uint64(e.UserID), 10),
			e.YearMonth, e.WorkStationName,
		})
	}
	w.Flush()
	return buf.Bytes(), w.Error()
}

// LearningMaterialCreateDTO for POST/PUT materials (multipart or JSON).
// Legacy form field "filePath" is accepted as an alias for Link (CDN URL).
type LearningMaterialCreateDTO struct {
	DiseaseTypeName string `json:"diseaseTypeName" form:"diseaseTypeName"`
	Title           string `json:"title" form:"title"`
	MaterialType    string `json:"materialType" form:"materialType"`
	Link            string `json:"link" form:"link"`
	FilePath        string `json:"filePath" form:"filePath"` // deprecated alias → Link
	SequenceNumber  int    `json:"sequenceNumber" form:"sequenceNumber"`
}

func resolveMaterialLink(dto *LearningMaterialCreateDTO) string {
	link := strings.TrimSpace(dto.Link)
	if link == "" {
		link = strings.TrimSpace(dto.FilePath)
	}
	return link
}

// CreateLearningMaterial inserts a material (aligned with Java saveMaterial).
func CreateLearningMaterial(db *gorm.DB, dto LearningMaterialCreateDTO) (*models.DiseaseLearningMaterial, error) {
	deleted := false
	now := time.Now()
	link := resolveMaterialLink(&dto)

	var seq int
	_ = db.Model(&models.DiseaseLearningMaterial{}).Select("COALESCE(MAX(sequence_number),0)").Scan(&seq)
	if dto.SequenceNumber > 0 {
		seq = dto.SequenceNumber
	} else {
		seq++
	}

	diseaseID := 0
	if name := strings.TrimSpace(dto.DiseaseTypeName); name != "" {
		var dt models.DiseaseType
		if err := db.Where("disease_type_name = ?", name).First(&dt).Error; err == nil {
			diseaseID = int(dt.DiseaseTypeID)
		}
	}

	m := models.DiseaseLearningMaterial{
		DiseaseTypeID:   diseaseID,
		DiseaseTypeName: dto.DiseaseTypeName,
		Title:           dto.Title,
		MaterialType:    dto.MaterialType,
		Link:            link,
		SequenceNumber:  seq,
		PublishDate:     &now,
		IsDeleted:       &deleted,
	}
	if err := db.Create(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

// GetLearningMaterialByID loads one material.
func GetLearningMaterialByID(db *gorm.DB, id uint) (*models.DiseaseLearningMaterial, error) {
	var m models.DiseaseLearningMaterial
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

// UpdateLearningMaterial updates fields.
func UpdateLearningMaterial(db *gorm.DB, id uint, dto LearningMaterialCreateDTO) (*models.DiseaseLearningMaterial, error) {
	m, err := GetLearningMaterialByID(db, id)
	if err != nil {
		return nil, err
	}
	updates := map[string]any{}
	if dto.DiseaseTypeName != "" {
		updates["disease_type_name"] = dto.DiseaseTypeName
		var dt models.DiseaseType
		if err := db.Where("disease_type_name = ?", dto.DiseaseTypeName).First(&dt).Error; err == nil {
			updates["disease_type_id"] = int(dt.DiseaseTypeID)
		}
	}
	if dto.Title != "" {
		updates["title"] = dto.Title
	}
	if dto.MaterialType != "" {
		updates["material_type"] = dto.MaterialType
	}
	if link := resolveMaterialLink(&dto); link != "" {
		updates["link"] = link
	}
	if dto.SequenceNumber > 0 {
		updates["sequence_number"] = dto.SequenceNumber
	}
	if len(updates) > 0 {
		if err := db.Model(m).Updates(updates).Error; err != nil {
			return nil, err
		}
	}
	return GetLearningMaterialByID(db, id)
}

// SetLearningMaterialStatus sets IsDeleted (enable/disable).
func SetLearningMaterialStatus(db *gorm.DB, id uint, disabled bool) error {
	m, err := GetLearningMaterialByID(db, id)
	if err != nil {
		return err
	}
	m.IsDeleted = &disabled
	if disabled {
		now := time.Now()
		m.DeletedDate = &now
	}
	return db.Model(m).Updates(map[string]any{
		"is_deleted":   disabled,
		"deleted_date": m.DeletedDate,
	}).Error
}

// MaterialSequenceDirection body.
type MaterialSequenceDirection struct {
	Direction string `json:"direction"` // up|down
}

// MoveLearningMaterialSequence swaps sequence with neighbor in same disease group.
func MoveLearningMaterialSequence(db *gorm.DB, id uint, direction string) error {
	cur, err := GetLearningMaterialByID(db, id)
	if err != nil {
		return err
	}
	var neighbor models.DiseaseLearningMaterial
	q := db.Where("disease_type_name = ?", cur.DiseaseTypeName)
	if strings.EqualFold(direction, "up") {
		q = q.Where("sequence_number < ?", cur.SequenceNumber).Order("sequence_number DESC")
	} else {
		q = q.Where("sequence_number > ?", cur.SequenceNumber).Order("sequence_number ASC")
	}
	if err := q.First(&neighbor).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return db.Transaction(func(tx *gorm.DB) error {
		a, b := cur.SequenceNumber, neighbor.SequenceNumber
		if err := tx.Model(cur).Update("sequence_number", b).Error; err != nil {
			return err
		}
		return tx.Model(&neighbor).Update("sequence_number", a).Error
	})
}
