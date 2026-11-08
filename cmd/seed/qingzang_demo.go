// Copyright (c) 2026 heathcetide. All rights reserved.

package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Qinghai–Tibet Railway corridor stations (approx. WGS84).
type railStation struct {
	Name      string
	Province  string
	City      string
	County    string
	Longitude float64
	Latitude  float64
	Altitude  int // meters, for work environment
}

var qingzangStations = []railStation{
	{Name: "西宁站工区", Province: "青海省", City: "西宁市", County: "城东区", Longitude: 101.7782, Latitude: 36.6232, Altitude: 2295},
	{Name: "西宁曹家堡机场接驳点", Province: "青海省", City: "西宁市", County: "城东区", Longitude: 102.0430, Latitude: 36.5280, Altitude: 2180},
	{Name: "湟源站工务点", Province: "青海省", City: "西宁市", County: "湟源县", Longitude: 101.2570, Latitude: 36.6820, Altitude: 2630},
	{Name: "海晏站驻地", Province: "青海省", City: "海北藏族自治州", County: "海晏县", Longitude: 100.9940, Latitude: 36.8960, Altitude: 3010},
	{Name: "青海湖东岸养护点", Province: "青海省", City: "海北藏族自治州", County: "海晏县", Longitude: 100.7400, Latitude: 36.7800, Altitude: 3200},
	{Name: "刚察站驻点", Province: "青海省", City: "海北藏族自治州", County: "刚察县", Longitude: 100.1460, Latitude: 37.3260, Altitude: 3300},
	{Name: "天峻站工区", Province: "青海省", City: "海西蒙古族藏族自治州", County: "天峻县", Longitude: 99.0230, Latitude: 37.3000, Altitude: 3400},
	{Name: "乌兰站工区", Province: "青海省", City: "海西蒙古族藏族自治州", County: "乌兰县", Longitude: 98.4798, Latitude: 36.9304, Altitude: 2960},
	{Name: "茶卡盐湖转运点", Province: "青海省", City: "海西蒙古族藏族自治州", County: "乌兰县", Longitude: 99.0800, Latitude: 36.7400, Altitude: 3050},
	{Name: "德令哈站驻地", Province: "青海省", City: "海西蒙古族藏族自治州", County: "德令哈市", Longitude: 97.3701, Latitude: 37.3743, Altitude: 2980},
	{Name: "德令哈西物资场", Province: "青海省", City: "海西蒙古族藏族自治州", County: "德令哈市", Longitude: 97.2200, Latitude: 37.3100, Altitude: 2950},
	{Name: "大柴旦驻点", Province: "青海省", City: "海西蒙古族藏族自治州", County: "海西蒙古族藏族自治州直辖", Longitude: 95.3600, Latitude: 37.8500, Altitude: 3170},
	{Name: "格尔木站工区", Province: "青海省", City: "海西蒙古族藏族自治州", County: "格尔木市", Longitude: 94.9033, Latitude: 36.4024, Altitude: 2808},
	{Name: "格尔木南山口营地", Province: "青海省", City: "海西蒙古族藏族自治州", County: "格尔木市", Longitude: 94.7800, Latitude: 36.1200, Altitude: 3100},
	{Name: "纳赤台养护点", Province: "青海省", City: "海西蒙古族藏族自治州", County: "格尔木市", Longitude: 94.4500, Latitude: 35.8700, Altitude: 3600},
	{Name: "昆仑山站附近营地", Province: "青海省", City: "海西蒙古族藏族自治州", County: "格尔木市", Longitude: 94.0550, Latitude: 35.6730, Altitude: 4767},
	{Name: "西大滩施工点", Province: "青海省", City: "海西蒙古族藏族自治州", County: "格尔木市", Longitude: 94.2000, Latitude: 35.7400, Altitude: 4100},
	{Name: "不冻泉观测点", Province: "青海省", City: "海西蒙古族藏族自治州", County: "格尔木市", Longitude: 93.7800, Latitude: 35.4200, Altitude: 4600},
	{Name: "可可西里站驻地", Province: "青海省", City: "海西蒙古族藏族自治州", County: "格尔木市", Longitude: 93.3400, Latitude: 35.1430, Altitude: 4500},
	{Name: "楚玛尔河桥工点", Province: "青海省", City: "海西蒙古族藏族自治州", County: "格尔木市", Longitude: 93.2000, Latitude: 35.1800, Altitude: 4550},
	{Name: "五道梁养护点", Province: "青海省", City: "海西蒙古族藏族自治州", County: "格尔木市", Longitude: 93.0800, Latitude: 35.2100, Altitude: 4612},
	{Name: "风火山隧道口营地", Province: "青海省", City: "玉树藏族自治州", County: "治多县", Longitude: 92.9000, Latitude: 34.6800, Altitude: 4900},
	{Name: "沱沱河站工区", Province: "青海省", City: "玉树藏族自治州", County: "治多县", Longitude: 92.4430, Latitude: 34.2170, Altitude: 4547},
	{Name: "沱沱河沿生活区", Province: "青海省", City: "玉树藏族自治州", County: "治多县", Longitude: 92.5000, Latitude: 34.2500, Altitude: 4530},
	{Name: "通天河桥头驻点", Province: "青海省", City: "玉树藏族自治州", County: "治多县", Longitude: 92.3000, Latitude: 33.9500, Altitude: 4600},
	{Name: "雁石坪驻地", Province: "青海省", City: "玉树藏族自治州", County: "杂多县", Longitude: 92.0400, Latitude: 33.5900, Altitude: 4720},
	{Name: "布强格养护点", Province: "青海省", City: "玉树藏族自治州", County: "杂多县", Longitude: 91.9500, Latitude: 33.3000, Altitude: 4800},
	{Name: "唐古拉站工区", Province: "西藏自治区", City: "那曲市", County: "安多县", Longitude: 91.8700, Latitude: 33.0060, Altitude: 5068},
	{Name: "唐古拉山口观测点", Province: "西藏自治区", City: "那曲市", County: "安多县", Longitude: 91.9200, Latitude: 32.8800, Altitude: 5231},
	{Name: "安多站驻地", Province: "西藏自治区", City: "那曲市", County: "安多县", Longitude: 91.6820, Latitude: 32.2650, Altitude: 4702},
	{Name: "安多南材料场", Province: "西藏自治区", City: "那曲市", County: "安多县", Longitude: 91.7200, Latitude: 32.1800, Altitude: 4680},
	{Name: "那曲站工区", Province: "西藏自治区", City: "那曲市", County: "色尼区", Longitude: 92.0512, Latitude: 31.4762, Altitude: 4507},
	{Name: "那曲西生活区", Province: "西藏自治区", City: "那曲市", County: "色尼区", Longitude: 91.9800, Latitude: 31.4500, Altitude: 4490},
	{Name: "古露站驻点", Province: "西藏自治区", City: "那曲市", County: "色尼区", Longitude: 91.6500, Latitude: 30.9500, Altitude: 4600},
	{Name: "当雄站驻地", Province: "西藏自治区", City: "拉萨市", County: "当雄县", Longitude: 91.1010, Latitude: 30.4780, Altitude: 4290},
	{Name: "当雄温泉转运点", Province: "西藏自治区", City: "拉萨市", County: "当雄县", Longitude: 91.0500, Latitude: 30.5200, Altitude: 4200},
	{Name: "羊八井养护点", Province: "西藏自治区", City: "拉萨市", County: "堆龙德庆区", Longitude: 90.5480, Latitude: 30.0920, Altitude: 4300},
	{Name: "羊八井地热站驻点", Province: "西藏自治区", City: "拉萨市", County: "堆龙德庆区", Longitude: 90.5000, Latitude: 30.0600, Altitude: 4280},
	{Name: "堆龙德庆工务点", Province: "西藏自治区", City: "拉萨市", County: "堆龙德庆区", Longitude: 90.9600, Latitude: 29.7200, Altitude: 3700},
	{Name: "拉萨站工区", Province: "西藏自治区", City: "拉萨市", County: "堆龙德庆区", Longitude: 91.0640, Latitude: 29.6550, Altitude: 3648},
	{Name: "拉萨西货场驻点", Province: "西藏自治区", City: "拉萨市", County: "堆龙德庆区", Longitude: 90.9800, Latitude: 29.6700, Altitude: 3650},
	{Name: "拉萨东郊医疗接驳点", Province: "西藏自治区", City: "拉萨市", County: "城关区", Longitude: 91.1800, Latitude: 29.6500, Altitude: 3650},
}

var workerSurnames = []string{
	"王", "李", "张", "刘", "陈", "杨", "赵", "黄", "周", "吴",
	"徐", "孙", "马", "朱", "胡", "郭", "何", "高", "林", "罗",
	"郑", "梁", "谢", "宋", "唐", "许", "韩", "冯", "邓", "曹",
	"彭", "曾", "肖", "田", "董", "袁", "潘", "于", "蒋", "蔡",
	"余", "杜", "叶", "程", "魏", "苏", "吕", "丁", "任", "卢",
	"姚", "沈", "钟", "姜", "崔", "谭", "陆", "汪", "范", "金",
	"石", "廖", "贾", "夏", "韦", "傅", "方", "白", "邹", "孟",
}

var workerGiven = []string{
	"伟", "强", "磊", "军", "勇", "杰", "涛", "明", "超", "辉",
	"华", "鹏", "飞", "刚", "平", "保国", "建军", "志强", "海波", "晓东",
	"文斌", "国庆", "永强", "立新", "春生", "宝山", "德明", "正清", "成龙", "学军",
	"扎西", "尼玛", "多吉", "达瓦", "格桑", "次仁", "洛桑", "顿珠", "旺堆", "巴桑",
	"才让", "索南", "旦增", "拉巴", "平措", "央金", "卓玛", "德吉", "白玛", "曲珍",
}

var departments = []string{
	"工程技术部", "测量队", "物资设备部", "安全部", "试验室",
	"综合管理部", "项目管理层", "征拆协调部", "合约部", "财务部",
}

var ethnicities = []string{"汉族", "藏族", "回族", "汉族", "汉族", "藏族", "汉族", "蒙古族"}
var educationLevels = []string{"初中", "中专/高中", "大专/本科", "大专/本科", "小学及以下", "中专/高中"}

var symptomDiseases = []string{
	"流感", "流感", "流感", "新型冠状病毒感染", "新型冠状病毒感染",
	"感染性腹泻", "感染性腹泻", "结核病", "斑疹伤寒", "流行性出血热",
	"登革热（蚊媒传染病）", "疟疾（蚊媒传染病）", "森林脑炎（蜱媒传染病）",
	"发热伴血小板减少综合征（蜱媒传染病）", "鼠疫", "炭疽",
	"高原反应相关不适", "上呼吸道感染",
}

// seedQingzangDemo creates realistic railway workers + check-ins along the Qinghai–Tibet Railway.
func seedQingzangDemo(db *gorm.DB, passwordHash string) error {
	// Always relocate leftover placeholder check-ins onto the corridor.
	if err := relocatePlaceholderCheckIns(db); err != nil {
		return err
	}
	if err := ensureDemoRailwayOnCorridor(db); err != nil {
		return err
	}

	var existing int64
	if err := db.Model(&models.User{}).
		Where("phone_number LIKE ?", "13980%").
		Count(&existing).Error; err != nil {
		return err
	}
	if existing >= 40 {
		fmt.Println("qingzang demo data already present, skip bulk seed")
		return nil
	}

	rng := rand.New(rand.NewSource(20261107))
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	type worker struct {
		user     models.User
		station  railStation
		dept     string
		gender   string
		age      int
		ethnicity string
	}

	workers := make([]worker, 0, 48)
	for i := 1; i <= 48; i++ {
		phone := fmt.Sprintf("13980%06d", i)
		var exists models.User
		err := db.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)}).
			Where("phone_number = ?", phone).First(&exists).Error
		if err == nil {
			st := qingzangStations[(i-1)%len(qingzangStations)]
			workers = append(workers, worker{user: exists, station: st})
			continue
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		st := qingzangStations[(i-1)%len(qingzangStations)]
		// Bias more workers toward mid/high plateau segments.
		if i%7 == 0 {
			st = qingzangStations[10+rng.Intn(6)] // Tanggula–Lhasa
		} else if i%5 == 0 {
			st = qingzangStations[5+rng.Intn(5)] // Kunlun–Tuotuohe
		}

		surname := workerSurnames[rng.Intn(len(workerSurnames))]
		given := workerGiven[rng.Intn(len(workerGiven))]
		name := surname + given
		gender := "男"
		if i%9 == 0 {
			gender = "女"
		}
		age := 22 + rng.Intn(28)
		eth := ethnicities[rng.Intn(len(ethnicities))]
		if eth == "藏族" && rng.Float64() < 0.55 {
			// Prefer Tibetan-style given names already in list; keep surname mix.
		}
		dept := departments[rng.Intn(len(departments))]
		edu := educationLevels[rng.Intn(len(educationLevels))]

		u := models.User{
			Username:    phone,
			Email:       fmt.Sprintf("qt-worker-%02d@railway.demo", i),
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
		plateauStart := time.Date(2018+rng.Intn(6), time.Month(1+rng.Intn(12)), 1+rng.Intn(28), 0, 0, 0, 0, time.Local)
		idTail := fmt.Sprintf("%02d%02d%02d%04d", birth.Year()%100, int(birth.Month()), birth.Day(), 1000+i)
		idNumber := fmt.Sprintf("63%02d%02d%s", 1+rng.Intn(28), 1+rng.Intn(12), idTail)
		if len(idNumber) > 18 {
			idNumber = idNumber[:18]
		}

		profile := models.EmployeeProfile{
			UserID:                 u.ID,
			Gender:                 gender,
			Age:                    age,
			BirthDate:              &birth,
			Ethnicity:              eth,
			EducationLevel:         edu,
			Department:             dept,
			WorkOnPlateauStartDate: &plateauStart,
			IDNumber:               idNumber,
			SpecificOccupation:     "否",
			EmergencyContactName:   surname + "家属",
			EmergencyContactPhone:  fmt.Sprintf("138%08d", 20000000+i),
			EmergencyContactRelation: "家属",
			Height:                 160 + rng.Intn(25),
			Weight:                 55 + rng.Intn(30),
			SmokingStatus:          []string{"从不吸", "以前吸", "现在吸"}[rng.Intn(3)],
			DrinkingStatus:         []string{"不喝酒", "偶尔喝酒（小于3次/周）", "经常喝酒（大于等于3次/周）"}[rng.Intn(3)],
		}
		if gender == "女" {
			profile.IsPregnant = false
		}
		if rng.Float64() < 0.18 {
			profile.HasMedicalHistory = true
			profile.HasHypertension = rng.Float64() < 0.4
			profile.HasTuberculosis = rng.Float64() < 0.1
		}
		if err := db.Create(&profile).Error; err != nil {
			return fmt.Errorf("profile %s: %w", phone, err)
		}

		workers = append(workers, worker{
			user: u, station: st, dept: dept, gender: gender, age: age, ethnicity: eth,
		})
	}

	// Hospitals along the corridor (replace empty demo if only the placeholder exists).
	var hospCount int64
	_ = db.Model(&models.Hospital{}).Count(&hospCount)
	if hospCount <= 1 {
		_ = db.Where("hospital_name = ?", "示例铁路医院").Delete(&models.Hospital{})
		hospitals := []models.Hospital{
			{HospitalName: "青海省交通医院西宁院区", HospitalPhoneNumber: "0971-6268120", Address: "青海省西宁市城东区互助中路"},
			{HospitalName: "格尔木市人民医院", HospitalPhoneNumber: "0979-8412120", Address: "青海省格尔木市昆仑路"},
			{HospitalName: "那曲市人民医院", HospitalPhoneNumber: "0896-3822120", Address: "西藏自治区那曲市色尼区"},
			{HospitalName: "西藏自治区人民医院", HospitalPhoneNumber: "0891-6332462", Address: "西藏自治区拉萨市城关区林廓北路"},
			{HospitalName: "青藏铁路疾病预防控制站（格尔木）", HospitalPhoneNumber: "0979-8458000", Address: "青海省格尔木市青藏路"},
		}
		if err := db.Create(&hospitals).Error; err != nil {
			return err
		}
	}

	yearMonth := today.Format("2006-01")
	checkIns := make([]models.DailyHealthStatus, 0, len(workers)*40)
	scorings := make([]models.DiseaseScoring, 0, 200)

	for wi, w := range workers {
		// Monthly work environment for most workers.
		var envCnt int64
		_ = db.Model(&models.WorkEnvironmentInfo{}).
			Where("user_id = ? AND year_month = ?", w.user.ID, yearMonth).
			Count(&envCnt)
		if envCnt == 0 && wi%6 != 0 { // leave a few unsigned for realism
			alt := w.station.Altitude + rng.Intn(80) - 20
			env := models.WorkEnvironmentInfo{
				UserID:                     w.user.ID,
				YearMonth:                  yearMonth,
				SubmissionTime:             today.Add(-time.Duration(rng.Intn(10)) * 24 * time.Hour),
				WorkStationAltitude:        alt,
				WorkStationName:            w.station.Name,
				DormitoryBuildingNumber:    1 + rng.Intn(8),
				DormitoryRoomNumber:        100 + rng.Intn(40),
				DormitoryOccupants:         4 + rng.Intn(5),
				DormitoryArea:              28 + float64(rng.Intn(20)),
				WindowOpenArea:             1.2 + rng.Float64()*2,
				VentilationCondition:       []string{"良好", "一般", "较差"}[rng.Intn(3)],
				AvgWorkRoomPeople:          6 + rng.Intn(12),
				AvgWorkFaceToFace:          2 + rng.Intn(6),
				AvgWorkPhysicalContact:     rng.Intn(4),
				AvgBreakfastRoomPeople:     8 + rng.Intn(20),
				AvgBreakfastFaceToFace:     2 + rng.Intn(5),
				AvgBreakfastPhysicalContact: rng.Intn(3),
				AvgLunchRoomPeople:         10 + rng.Intn(25),
				AvgLunchFaceToFace:         2 + rng.Intn(6),
				AvgLunchPhysicalContact:    rng.Intn(3),
				AvgDinnerRoomPeople:        8 + rng.Intn(20),
				AvgDinnerFaceToFace:        2 + rng.Intn(5),
				AvgDinnerPhysicalContact:   rng.Intn(3),
				CollectiveActivitiesCount:  rng.Intn(6),
				AvgActivityRoomPeople:      10 + rng.Intn(30),
				AvgActivityFaceToFace:      rng.Intn(8),
				AvgActivityPhysicalContact: rng.Intn(4),
				AvgMonthlyPressure:         480 + float64(rng.Intn(120)), // hPa-ish at altitude
				MaxTemperature:             8 + float64(rng.Intn(18)),
				MinTemperature:             -18 + float64(rng.Intn(16)),
			}
			if err := db.Create(&env).Error; err != nil {
				return err
			}
		}

		// ~45 days of check-ins with realistic attendance.
		for dayOffset := 0; dayOffset < 45; dayOffset++ {
			day := today.AddDate(0, 0, -dayOffset)
			// Weekends slightly lower check-in rate; high plateau segments slightly higher.
			rate := 0.88
			if day.Weekday() == time.Saturday || day.Weekday() == time.Sunday {
				rate = 0.72
			}
			if w.station.Altitude > 4500 {
				rate += 0.05
			}
			if rng.Float64() > rate {
				continue
			}

			// Skip if already exists (re-run safe).
			var cnt int64
			_ = db.Model(&models.DailyHealthStatus{}).
				Where("user_id = ? AND check_in_date = ?", w.user.ID, day).
				Count(&cnt)
			if cnt > 0 {
				continue
			}

			hour := 6 + rng.Intn(4)
			minute := rng.Intn(60)
			checkTime := time.Date(day.Year(), day.Month(), day.Day(), hour, minute, rng.Intn(60), 0, time.Local)

			// Small GPS jitter around assigned station.
			lng := w.station.Longitude + (rng.Float64()-0.5)*0.08
			lat := w.station.Latitude + (rng.Float64()-0.5)*0.06
			// Occasionally check in at a neighboring station (duty travel).
			locName := w.station.Name
			prov, city, county := w.station.Province, w.station.City, w.station.County
			if rng.Float64() < 0.12 {
				nb := qingzangStations[rng.Intn(len(qingzangStations))]
				lng, lat = nb.Longitude+(rng.Float64()-0.5)*0.05, nb.Latitude+(rng.Float64()-0.5)*0.04
				locName, prov, city, county = nb.Name, nb.Province, nb.City, nb.County
			}

			healthy := true
			disease := ""
			// Plateau + winter-ish days: slightly more symptoms; overall ~9% unhealthy.
			unhealthyP := 0.08
			if w.station.Altitude > 4500 {
				unhealthyP = 0.14
			}
			if dayOffset == 0 {
				unhealthyP += 0.04 // ensure today has some flags for dashboard
			}
			if rng.Float64() < unhealthyP {
				healthy = false
				disease = symptomDiseases[rng.Intn(len(symptomDiseases))]
				if disease == "高原反应相关不适" && w.station.Altitude < 4000 {
					disease = "上呼吸道感染"
				}
			}
			h := healthy
			rec := models.DailyHealthStatus{
				UserID:          w.user.ID,
				CheckInDate:     day,
				CheckInTime:     checkTime,
				IsHealth:        &h,
				DiseaseTypeName: disease,
				Longitude:       round6(lng),
				Latitude:        round6(lat),
				Province:        prov,
				City:            city,
				County:          county,
				LocationName:    locName,
			}
			checkIns = append(checkIns, rec)
		}
	}

	// Batch insert check-ins.
	const batch = 200
	for i := 0; i < len(checkIns); i += batch {
		end := i + batch
		if end > len(checkIns) {
			end = len(checkIns)
		}
		if err := db.Create(checkIns[i:end]).Error; err != nil {
			return fmt.Errorf("check-ins batch: %w", err)
		}
	}

	// Attach disease scoring for unhealthy check-ins (esp. today) for dashboard charts.
	var unhealthy []models.DailyHealthStatus
	if err := db.Where("is_health = ? AND check_in_date >= ?", false, today.AddDate(0, 0, -14)).
		Find(&unhealthy).Error; err != nil {
		return err
	}
	for _, s := range unhealthy {
		name := s.DiseaseTypeName
		if name == "" {
			name = "流感"
		}
		if name == "高原反应相关不适" || name == "上呼吸道感染" {
			name = "流感"
		}
		day := s.CheckInDate
		score := 35 + rng.Float64()*45
		scorings = append(scorings, models.DiseaseScoring{
			StatusID:        s.StatusID,
			SymptomTime:     &day,
			DiseaseTypeName: name,
			Score:           math.Round(score*10) / 10,
		})
		// Secondary lower score for a related label sometimes.
		if rng.Float64() < 0.35 {
			alt := "感染性腹泻"
			if name == "感染性腹泻" {
				alt = "流感"
			}
			scorings = append(scorings, models.DiseaseScoring{
				StatusID:        s.StatusID,
				SymptomTime:     &day,
				DiseaseTypeName: alt,
				Score:           math.Round((15+rng.Float64()*20)*10) / 10,
			})
		}
	}
	if len(scorings) > 0 {
		if err := db.Create(&scorings).Error; err != nil {
			return fmt.Errorf("disease scoring: %w", err)
		}
	}

	var userN, checkN, todayN int64
	_ = db.Model(&models.User{}).Where("role = ?", models.UserRoleRailway).Count(&userN)
	_ = db.Model(&models.DailyHealthStatus{}).Count(&checkN)
	_ = db.Model(&models.DailyHealthStatus{}).
		Where("check_in_date >= ? AND check_in_date < ?", today, today.AddDate(0, 0, 1)).
		Count(&todayN)
	fmt.Printf("qingzang demo: railway users=%d check-ins=%d today=%d\n", userN, checkN, todayN)
	return nil
}

func round6(v float64) float64 {
	return math.Round(v*1e6) / 1e6
}

func relocatePlaceholderCheckIns(db *gorm.DB) error {
	st := qingzangStations[4] // 格尔木站工区
	res := db.Model(&models.DailyHealthStatus{}).
		Where("location_name IN ? OR (longitude BETWEEN ? AND ? AND latitude BETWEEN ? AND ?)",
			[]string{"测试点", "示例地点", "演示地点"},
			103.5, 105.0, 30.0, 31.2).
		Updates(map[string]any{
			"location_name": st.Name,
			"longitude":     st.Longitude,
			"latitude":      st.Latitude,
			"province":      st.Province,
			"city":          st.City,
			"county":        st.County,
		})
	return res.Error
}

// ensureDemoRailwayOnCorridor gives 演示职工 a Qingzang station profile + today's check-in if missing.
func ensureDemoRailwayOnCorridor(db *gorm.DB) error {
	var u models.User
	if err := db.Where("phone_number = ?", "13800000000").First(&u).Error; err != nil {
		return nil
	}
	st := qingzangStations[4]
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	var profile models.EmployeeProfile
	err := db.Where("user_id = ?", u.ID).First(&profile).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		birth := time.Date(1988, 5, 12, 0, 0, 0, 0, time.Local)
		plateau := time.Date(2019, 3, 1, 0, 0, 0, 0, time.Local)
		profile = models.EmployeeProfile{
			UserID:                   u.ID,
			Gender:                   "男",
			Age:                      36,
			BirthDate:                &birth,
			Ethnicity:                "汉族",
			EducationLevel:           "大专/本科",
			Department:               "工程技术部",
			WorkOnPlateauStartDate:   &plateau,
			IDNumber:                 "632801198805121234",
			SpecificOccupation:       "否",
			EmergencyContactName:     "王家属",
			EmergencyContactPhone:    "13820000000",
			EmergencyContactRelation: "家属",
			Height:                   172,
			Weight:                   68,
			SmokingStatus:            "从不吸",
			DrinkingStatus:           "偶尔喝酒（小于3次/周）",
		}
		if err := db.Create(&profile).Error; err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else if profile.Department == "" {
		_ = db.Model(&profile).Updates(map[string]any{
			"gender":          "男",
			"age":             36,
			"department":      "工程技术部",
			"ethnicity":       "汉族",
			"education_level": "大专/本科",
			"id_number":       "632801198805121234",
		}).Error
	}

	var cnt int64
	_ = db.Model(&models.DailyHealthStatus{}).
		Where("user_id = ? AND check_in_date >= ? AND check_in_date < ?", u.ID, today, today.AddDate(0, 0, 1)).
		Count(&cnt)
	if cnt == 0 {
		h := true
		checkTime := time.Date(today.Year(), today.Month(), today.Day(), 7, 35, 0, 0, time.Local)
		rec := models.DailyHealthStatus{
			UserID:       u.ID,
			CheckInDate:  today,
			CheckInTime:  checkTime,
			IsHealth:     &h,
			Longitude:    st.Longitude,
			Latitude:     st.Latitude,
			Province:     st.Province,
			City:         st.City,
			County:       st.County,
			LocationName: st.Name,
		}
		if err := db.Create(&rec).Error; err != nil {
			return err
		}
	} else {
		_ = db.Model(&models.DailyHealthStatus{}).
			Where("user_id = ? AND check_in_date >= ? AND check_in_date < ?", u.ID, today, today.AddDate(0, 0, 1)).
			Updates(map[string]any{
				"location_name": st.Name,
				"longitude":     st.Longitude,
				"latitude":      st.Latitude,
				"province":      st.Province,
				"city":          st.City,
				"county":        st.County,
			}).Error
	}
	return nil
}
