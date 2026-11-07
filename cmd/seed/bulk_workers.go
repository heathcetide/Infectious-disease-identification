// Copyright (c) 2026 heathcetide. All rights reserved.

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const targetRailwayWorkers = 8000

// Nationwide sparse stations (leave / transfer / affiliated projects).
var nationwideStations = []railStation{
	{Name: "昆明南站附近驻点", Province: "云南省", City: "昆明市", County: "呈贡区", Longitude: 102.8615, Latitude: 24.8526, Altitude: 1890},
	{Name: "大理站工务点", Province: "云南省", City: "大理白族自治州", County: "大理市", Longitude: 100.2480, Latitude: 25.6860, Altitude: 1976},
	{Name: "丽江站驻地", Province: "云南省", City: "丽江市", County: "古城区", Longitude: 100.2490, Latitude: 26.8550, Altitude: 2393},
	{Name: "景洪站驻点", Province: "云南省", City: "西双版纳傣族自治州", County: "景洪市", Longitude: 100.7980, Latitude: 22.0020, Altitude: 550},
	{Name: "成都东站换乘点", Province: "四川省", City: "成都市", County: "成华区", Longitude: 104.1410, Latitude: 30.6300, Altitude: 500},
	{Name: "西昌站工区", Province: "四川省", City: "凉山彝族自治州", County: "西昌市", Longitude: 102.2640, Latitude: 27.8950, Altitude: 1500},
	{Name: "重庆西站驻点", Province: "重庆市", City: "重庆市", County: "沙坪坝区", Longitude: 106.4350, Latitude: 29.5040, Altitude: 350},
	{Name: "贵阳北站驻点", Province: "贵州省", City: "贵阳市", County: "观山湖区", Longitude: 106.6730, Latitude: 26.6470, Altitude: 1100},
	{Name: "南宁东站驻点", Province: "广西壮族自治区", City: "南宁市", County: "青秀区", Longitude: 108.4140, Latitude: 22.8400, Altitude: 80},
	{Name: "广州南站中转点", Province: "广东省", City: "广州市", County: "番禺区", Longitude: 113.2690, Latitude: 22.9890, Altitude: 20},
	{Name: "深圳北站驻点", Province: "广东省", City: "深圳市", County: "龙华区", Longitude: 114.0290, Latitude: 22.6100, Altitude: 50},
	{Name: "海口东站驻点", Province: "海南省", City: "海口市", County: "美兰区", Longitude: 110.3600, Latitude: 20.0300, Altitude: 10},
	{Name: "长沙南站驻点", Province: "湖南省", City: "长沙市", County: "雨花区", Longitude: 113.0660, Latitude: 28.1360, Altitude: 60},
	{Name: "武汉站驻点", Province: "湖北省", City: "武汉市", County: "洪山区", Longitude: 114.4210, Latitude: 30.6070, Altitude: 30},
	{Name: "郑州东站中转点", Province: "河南省", City: "郑州市", County: "金水区", Longitude: 113.7760, Latitude: 34.7600, Altitude: 110},
	{Name: "西安北站中转点", Province: "陕西省", City: "西安市", County: "未央区", Longitude: 108.9390, Latitude: 34.3760, Altitude: 400},
	{Name: "兰州西站驻点", Province: "甘肃省", City: "兰州市", County: "安宁区", Longitude: 103.7500, Latitude: 36.0680, Altitude: 1520},
	{Name: "银川站驻点", Province: "宁夏回族自治区", City: "银川市", County: "金凤区", Longitude: 106.1680, Latitude: 38.4870, Altitude: 1110},
	{Name: "呼和浩特东站驻点", Province: "内蒙古自治区", City: "呼和浩特市", County: "新城区", Longitude: 111.7670, Latitude: 40.8490, Altitude: 1050},
	{Name: "乌鲁木齐站休整点", Province: "新疆维吾尔自治区", City: "乌鲁木齐市", County: "新市区", Longitude: 87.5550, Latitude: 43.8400, Altitude: 800},
	{Name: "喀什站驻点", Province: "新疆维吾尔自治区", City: "喀什地区", County: "喀什市", Longitude: 75.9890, Latitude: 39.4680, Altitude: 1290},
	{Name: "哈尔滨西站驻点", Province: "黑龙江省", City: "哈尔滨市", County: "南岗区", Longitude: 126.5780, Latitude: 45.7060, Altitude: 140},
	{Name: "长春西站驻点", Province: "吉林省", City: "长春市", County: "绿园区", Longitude: 125.1970, Latitude: 43.8810, Altitude: 220},
	{Name: "沈阳北站驻点", Province: "辽宁省", City: "沈阳市", County: "沈河区", Longitude: 123.3840, Latitude: 41.8150, Altitude: 50},
	{Name: "北京西站中转点", Province: "北京市", City: "北京市", County: "丰台区", Longitude: 116.3210, Latitude: 39.8950, Altitude: 50},
	{Name: "天津西站驻点", Province: "天津市", City: "天津市", County: "红桥区", Longitude: 117.1560, Latitude: 39.1570, Altitude: 10},
	{Name: "石家庄站驻点", Province: "河北省", City: "石家庄市", County: "桥西区", Longitude: 114.4780, Latitude: 38.0090, Altitude: 80},
	{Name: "济南西站驻点", Province: "山东省", City: "济南市", County: "槐荫区", Longitude: 116.8900, Latitude: 36.6690, Altitude: 50},
	{Name: "南京南站中转点", Province: "江苏省", City: "南京市", County: "雨花台区", Longitude: 118.7990, Latitude: 31.9690, Altitude: 20},
	{Name: "上海虹桥站中转点", Province: "上海市", City: "上海市", County: "闵行区", Longitude: 121.3200, Latitude: 31.1940, Altitude: 10},
	{Name: "杭州东站驻点", Province: "浙江省", City: "杭州市", County: "上城区", Longitude: 120.2130, Latitude: 30.2910, Altitude: 20},
	{Name: "合肥南站驻点", Province: "安徽省", City: "合肥市", County: "包河区", Longitude: 117.3170, Latitude: 31.7980, Altitude: 30},
	{Name: "南昌西站驻点", Province: "江西省", City: "南昌市", County: "红谷滩区", Longitude: 115.7940, Latitude: 28.6270, Altitude: 30},
	{Name: "福州站驻点", Province: "福建省", City: "福州市", County: "晋安区", Longitude: 119.3200, Latitude: 26.0850, Altitude: 20},
	{Name: "太原南站驻点", Province: "山西省", City: "太原市", County: "小店区", Longitude: 112.6060, Latitude: 37.7910, Altitude: 800},
}

var bulkOccs = []string{
	"否", "否", "否", "否", "否", "否",
	"医务人员", "病原微生物检测人员", "野生动物接触相关人员", "家禽、家畜养殖人员",
}

// seedBulkWorkers expands railway workforce to targetRailwayWorkers with nationwide sparse check-ins.
func seedBulkWorkers(db *gorm.DB, passwordHash string) error {
	rng := rand.New(rand.NewSource(20261107 + 99))
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	var railwayN int64
	if err := db.Model(&models.User{}).Where("role = ?", models.UserRoleRailway).Count(&railwayN).Error; err != nil {
		return err
	}
	need := targetRailwayWorkers - int(railwayN)
	if need < 0 {
		need = 0
	}
	fmt.Printf("bulk workers: have=%d need=%d target=%d\n", railwayN, need, targetRailwayWorkers)

	type worker struct {
		user    models.User
		station railStation
	}
	newWorkers := make([]worker, 0, need)

	// Create missing users: phone 13980xxxxxx continuing from 1..N
	for i := 1; len(newWorkers) < need; i++ {
		phone := fmt.Sprintf("13980%06d", i)
		var exists models.User
		err := db.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)}).
			Where("phone_number = ?", phone).First(&exists).Error
		if err == nil {
			continue
		}
		if err != gorm.ErrRecordNotFound {
			return err
		}

		st := pickHomeStation(rng, i)
		surname := workerSurnames[rng.Intn(len(workerSurnames))]
		given := workerGiven[rng.Intn(len(workerGiven))]
		name := surname + given
		gender := "男"
		if i%8 == 0 {
			gender = "女"
		}
		age := 22 + rng.Intn(30)
		eth := ethnicities[rng.Intn(len(ethnicities))]
		dept := departments[rng.Intn(len(departments))]
		edu := educationLevels[rng.Intn(len(educationLevels))]
		occ := bulkOccs[rng.Intn(len(bulkOccs))]

		u := models.User{
			Username:    phone,
			Email:       fmt.Sprintf("qt-worker-%04d@railway.demo", i),
			Name:        name,
			PhoneNumber: phone,
			Phone:       phone,
			Password:    passwordHash,
			Role:        models.UserRoleRailway,
			Status:      models.UserStatusActive,
		}
		if err := db.Create(&u).Error; err != nil {
			return fmt.Errorf("create worker %s: %w", phone, err)
		}

		birth := time.Date(now.Year()-age, time.Month(1+rng.Intn(12)), 1+rng.Intn(28), 0, 0, 0, 0, time.Local)
		plateauStart := time.Date(2016+rng.Intn(8), time.Month(1+rng.Intn(12)), 1+rng.Intn(28), 0, 0, 0, 0, time.Local)
		idNumber := fmt.Sprintf("63%02d%02d%04d%02d%02d%04d", 1+rng.Intn(28), 1+rng.Intn(12), birth.Year(), int(birth.Month()), birth.Day(), 1000+i%9000)
		if len(idNumber) > 18 {
			idNumber = idNumber[:18]
		}
		profile := models.EmployeeProfile{
			UserID:                   u.ID,
			Gender:                   gender,
			Age:                      age,
			BirthDate:                &birth,
			Ethnicity:                eth,
			EducationLevel:           edu,
			Department:               dept,
			WorkOnPlateauStartDate:   &plateauStart,
			IDNumber:                 idNumber,
			SpecificOccupation:       occ,
			EmergencyContactName:     surname + "家属",
			EmergencyContactPhone:    fmt.Sprintf("138%08d", 30000000+i),
			EmergencyContactRelation: "家属",
			Height:                   158 + rng.Intn(28),
			Weight:                   52 + rng.Intn(35),
			SmokingStatus:            []string{"从不吸", "以前吸", "现在吸"}[rng.Intn(3)],
			DrinkingStatus:           []string{"不喝酒", "偶尔喝酒（小于3次/周）", "经常喝酒（大于等于3次/周）"}[rng.Intn(3)],
		}
		if err := db.Create(&profile).Error; err != nil {
			return err
		}
		newWorkers = append(newWorkers, worker{user: u, station: st})
		if len(newWorkers)%100 == 0 {
			fmt.Printf("  created users %d/%d\n", len(newWorkers), need)
		}
	}

	// Also backfill check-ins for existing bulk phones that may lack history.
	var allWorkers []models.User
	if err := db.Where("role = ? AND phone_number LIKE ?", models.UserRoleRailway, "13980%").
		Find(&allWorkers).Error; err != nil {
		return err
	}

	checkIns := make([]models.DailyHealthStatus, 0, 8000)
	flush := func() error {
		if len(checkIns) == 0 {
			return nil
		}
		const batch = 300
		for i := 0; i < len(checkIns); i += batch {
			end := i + batch
			if end > len(checkIns) {
				end = len(checkIns)
			}
			if err := db.Create(checkIns[i:end]).Error; err != nil {
				return err
			}
		}
		checkIns = checkIns[:0]
		return nil
	}

	for wi, u := range allWorkers {
		home := pickHomeStation(rng, wi+1)
		// Infer from phone index when possible.
		var phoneIdx int
		_, _ = fmt.Sscanf(u.PhoneNumber, "13980%d", &phoneIdx)
		if phoneIdx > 0 {
			home = pickHomeStation(rng, phoneIdx)
		}

		var existingDays int64
		_ = db.Model(&models.DailyHealthStatus{}).Where("user_id = ?", u.ID).Count(&existingDays)
		// Dense Nov–Dec fill is handled by seedDenseSeason; keep bulk history light.
		wantDays := 12 + (wi % 8)
		if int(existingDays) >= wantDays {
			continue
		}

		added := 0
		for dayOffset := 0; dayOffset < 50 && int(existingDays)+added < wantDays; dayOffset++ {
			day := today.AddDate(0, 0, -dayOffset)
			var cnt int64
			_ = db.Model(&models.DailyHealthStatus{}).
				Where("user_id = ? AND check_in_date = ?", u.ID, day).Count(&cnt)
			if cnt > 0 {
				continue
			}
			rate := 0.86
			if day.Weekday() == time.Saturday || day.Weekday() == time.Sunday {
				rate = 0.68
			}
			if rng.Float64() > rate {
				continue
			}

			st := home
			// ~18% of check-ins happen off home base (nationwide sparse).
			if rng.Float64() < 0.18 {
				st = nationwideStations[rng.Intn(len(nationwideStations))]
			} else if rng.Float64() < 0.10 {
				st = qingzangStations[rng.Intn(len(qingzangStations))]
			}

			lng := st.Longitude + (rng.Float64()-0.5)*0.06
			lat := st.Latitude + (rng.Float64()-0.5)*0.05
			h := true
			disease := ""
			unhealthyP := 0.07
			if st.Altitude > 4500 {
				unhealthyP = 0.12
			}
			if dayOffset == 0 {
				unhealthyP += 0.03
			}
			if rng.Float64() < unhealthyP {
				h = false
				disease = symptomDiseases[rng.Intn(len(symptomDiseases))]
			}
			checkTime := time.Date(day.Year(), day.Month(), day.Day(), 6+rng.Intn(5), rng.Intn(60), rng.Intn(60), 0, time.Local)
			hh := h
			checkIns = append(checkIns, models.DailyHealthStatus{
				UserID:          u.ID,
				CheckInDate:     day,
				CheckInTime:     checkTime,
				IsHealth:        &hh,
				DiseaseTypeName: disease,
				Longitude:       round6(lng),
				Latitude:        round6(lat),
				Province:        st.Province,
				City:            st.City,
				County:          st.County,
				LocationName:    st.Name,
			})
			added++
			if len(checkIns) >= 2000 {
				if err := flush(); err != nil {
					return err
				}
				fmt.Printf("  flushed check-ins mid-run (user idx %d)\n", wi)
			}
		}
		if err := ensureSparseOutliers(db, rng, u.ID, today, 2+wi%3); err != nil {
			return err
		}
	}
	if err := flush(); err != nil {
		return err
	}

	var userN, checkN, todayN int64
	_ = db.Model(&models.User{}).Where("role = ?", models.UserRoleRailway).Count(&userN)
	_ = db.Model(&models.DailyHealthStatus{}).Count(&checkN)
	_ = db.Model(&models.DailyHealthStatus{}).
		Where("check_in_date >= ? AND check_in_date < ?", today, today.AddDate(0, 0, 1)).
		Count(&todayN)
	fmt.Printf("bulk done: railway=%d check-ins=%d today=%d\n", userN, checkN, todayN)
	return nil
}

func pickHomeStation(rng *rand.Rand, i int) railStation {
	// ~72% Qingzang corridor home bases, rest nationwide.
	if i%7 == 0 || i%11 == 0 {
		return nationwideStations[rng.Intn(len(nationwideStations))]
	}
	if i%5 == 0 {
		return qingzangStations[5+rng.Intn(len(qingzangStations)-5)]
	}
	return qingzangStations[(i-1)%len(qingzangStations)]
}

func ensureSparseOutliers(db *gorm.DB, rng *rand.Rand, userID uint, today time.Time, n int) error {
	var outlierCnt int64
	_ = db.Model(&models.DailyHealthStatus{}).
		Where("user_id = ? AND province NOT IN ?", userID,
			[]string{"青海省", "西藏自治区"}).
		Count(&outlierCnt)
	if int(outlierCnt) >= n {
		return nil
	}
	need := n - int(outlierCnt)
	for i := 0; i < need*4 && need > 0; i++ {
		day := today.AddDate(0, 0, -(1 + rng.Intn(50)))
		var cnt int64
		_ = db.Model(&models.DailyHealthStatus{}).
			Where("user_id = ? AND check_in_date = ?", userID, day).Count(&cnt)
		if cnt > 0 {
			continue
		}
		st := nationwideStations[rng.Intn(len(nationwideStations))]
		h := true
		if rng.Float64() < 0.06 {
			h = false
		}
		disease := ""
		if !h {
			disease = symptomDiseases[rng.Intn(len(symptomDiseases))]
		}
		checkTime := time.Date(day.Year(), day.Month(), day.Day(), 7+rng.Intn(4), rng.Intn(60), 0, 0, time.Local)
		hh := h
		rec := models.DailyHealthStatus{
			UserID:          userID,
			CheckInDate:     day,
			CheckInTime:     checkTime,
			IsHealth:        &hh,
			DiseaseTypeName: disease,
			Longitude:       round6(st.Longitude + (rng.Float64()-0.5)*0.05),
			Latitude:        round6(st.Latitude + (rng.Float64()-0.5)*0.04),
			Province:        st.Province,
			City:            st.City,
			County:          st.County,
			LocationName:    st.Name,
		}
		if err := db.Create(&rec).Error; err != nil {
			return err
		}
		need--
	}
	return nil
}
