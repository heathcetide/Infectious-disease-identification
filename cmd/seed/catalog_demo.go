// Copyright (c) 2026 heathcetide. All rights reserved.

package main

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/heathcetide/infectious-disease/internal/models"
)

func silentDB(db *gorm.DB) *gorm.DB {
	return db.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)})
}

// Realistic subtypes used in Chinese infectious-disease reporting / clinical forms.
var diseaseSubtypeCatalog = map[string][]string{
	"鼠疫":         {"腺鼠疫", "肺鼠疫", "败血症型鼠疫"},
	"炭疽":         {"皮肤炭疽", "肠炭疽", "肺炭疽"},
	"疟疾（蚊媒传染病）":  {"间日疟", "恶性疟", "三日疟", "卵形疟"},
	"登革热（蚊媒传染病）": {"登革热", "登革出血热"},
	"结核病":        {"肺结核", "肺外结核"},
	"流感":         {"甲型流感", "乙型流感"},
	"感染性腹泻":      {"细菌性腹泻", "病毒性腹泻", "寄生虫性腹泻"},
	"斑疹伤寒":       {"流行性斑疹伤寒", "地方性斑疹伤寒"},
	"流行性出血热":     {"轻型", "中型", "重型", "危重型"},
	"新型冠状病毒感染":   {"轻型", "普通型", "重型", "危重型"},
}

func seedDiseaseSubtypes(db *gorm.DB) error {
	var types []models.DiseaseType
	if err := db.Find(&types).Error; err != nil {
		return err
	}
	for _, dt := range types {
		subs, ok := diseaseSubtypeCatalog[dt.DiseaseTypeName]
		has := ok && len(subs) > 0
		_ = db.Model(&dt).Update("has_subtype", has).Error
		if !has {
			continue
		}
		for _, name := range subs {
			var cnt int64
			_ = db.Model(&models.DiseaseSubtype{}).
				Where("disease_type_id = ? AND subtype_name = ?", dt.DiseaseTypeID, name).
				Count(&cnt)
			if cnt > 0 {
				continue
			}
			if err := db.Create(&models.DiseaseSubtype{
				DiseaseTypeID: dt.DiseaseTypeID,
				SubtypeName:   name,
			}).Error; err != nil {
				return fmt.Errorf("subtype %s/%s: %w", dt.DiseaseTypeName, name, err)
			}
		}
	}
	return nil
}

func seedSatisfactionDemo(db *gorm.DB) error {
	rng := rand.New(rand.NewSource(20261107 + 17))

	// Fix empty placeholder row if present.
	_ = db.Model(&models.SatisfactionSurvey{}).
		Where("survey_title = ? OR survey_title IS NULL", "").
		Updates(map[string]any{
			"survey_title":   "2026年职工健康筛查系统使用满意度（一季度）",
			"survey_content": "本问卷面向一线职工，了解您对健康筛查小程序的整体使用感受。请结合日常打卡、宣教阅读与异常提醒等场景，评价操作是否顺畅、信息是否清楚、是否对您的健康管理有帮助。",
			"is_open":        true,
		}).Error

	surveys := []models.SatisfactionSurvey{
		{
			SurveyTitle:   "2026年职工健康筛查系统使用满意度（一季度）",
			SurveyContent: "本问卷面向一线职工，了解您对健康筛查小程序的整体使用感受。请结合日常打卡、宣教阅读与异常提醒等场景，评价操作是否顺畅、信息是否清楚、是否对您的健康管理有帮助。",
			IsOpen:        true,
			CreateTime:    time.Date(2026, 3, 15, 9, 0, 0, 0, time.Local),
		},
		{
			SurveyTitle:   "高原工区宣教材料阅读体验调查",
			SurveyContent: "请评价高原工区传染病与高原相关不适宣教材料：内容是否易懂、是否贴合现场实际、阅读后是否清楚该如何防护与报告。您的反馈将用于优化宣教重点与表述方式。",
			IsOpen:        true,
			CreateTime:    time.Date(2026, 6, 20, 10, 30, 0, 0, time.Local),
		},
		{
			SurveyTitle:   "日打卡流程与定位准确性反馈",
			SurveyContent: "请从「入口是否好找、定位是否准确、提交是否稳定、失败提示是否清楚」等方面评价每日健康打卡流程，便于我们优先修复影响出勤打卡的问题。",
			IsOpen:        false,
			CreateTime:    time.Date(2026, 8, 8, 14, 0, 0, 0, time.Local),
		},
		{
			SurveyTitle:   "疾控预警短信/站内提醒满意度",
			SurveyContent: "当系统对重点关注人员或异常症状发出提醒时，请评价通知是否及时、渠道是否合适、文案是否可理解、是否清楚下一步该做什么。",
			IsOpen:        true,
			CreateTime:    time.Date(2026, 9, 12, 11, 0, 0, 0, time.Local),
		},
		{
			SurveyTitle:   "医护诊断协同与报告回传体验",
			SurveyContent: "请从医护协同视角评价：诊断信息录入是否方便、检验/报告查看是否及时、结果回传链路是否完整，以及跨岗位协作中最影响效率的环节。",
			IsOpen:        false,
			CreateTime:    time.Date(2026, 10, 5, 16, 20, 0, 0, time.Local),
		},
		{
			SurveyTitle:   "2026年秋季传染病防控专项满意度",
			SurveyContent: "针对秋季呼吸道传染病防控，请评价宣教覆盖、筛查提醒、工区管理措施是否到位，以及您是否感到被充分告知风险与防护要点。",
			IsOpen:        true,
			CreateTime:    time.Date(2026, 10, 28, 9, 40, 0, 0, time.Local),
		},
		{
			SurveyTitle:   "高原工区宿舍通风与健康管理满意度",
			SurveyContent: "请结合宿舍居住体验，评价通风习惯引导、消杀安排、健康管理提醒是否实用，以及是否有仍需加强的卫生或空间管理问题。",
			IsOpen:        true,
			CreateTime:    time.Date(2026, 7, 18, 10, 0, 0, 0, time.Local),
		},
		{
			SurveyTitle:   "移动端症状打卡易用性调查",
			SurveyContent: "请评价小程序症状打卡的易用性：页面是否简洁、选项是否好理解、定位与提交耗时是否可接受，以及最希望改进的一步是什么。",
			IsOpen:        false,
			CreateTime:    time.Date(2026, 5, 22, 15, 30, 0, 0, time.Local),
		},
		{
			SurveyTitle:   "疾控重点人员随访服务评价",
			SurveyContent: "若您曾被标记为重点关注人员，请评价随访是否及时、沟通是否尊重、健康建议是否可执行，以及对后续随访频次与方式的期望。",
			IsOpen:        true,
			CreateTime:    time.Date(2026, 9, 30, 9, 10, 0, 0, time.Local),
		},
		{
			SurveyTitle:   "铁路医院协同转诊体验调查",
			SurveyContent: "请评价从现场发现异常到铁路医院协同处置的全流程：转诊指引是否清楚、对接是否顺畅、检查与结果回传是否及时，以及您认为最需要优化的环节。",
			IsOpen:        true,
			CreateTime:    time.Date(2026, 8, 26, 11, 20, 0, 0, time.Local),
		},
	}

	for _, s := range surveys {
		var existing models.SatisfactionSurvey
		err := silentDB(db).Where("survey_title = ?", s.SurveyTitle).First(&existing).Error
		if err == nil {
			// Refresh copy/content for already-seeded rows so admin list stays current.
			_ = db.Model(&existing).Updates(map[string]any{
				"survey_content": s.SurveyContent,
				"is_open":        s.IsOpen,
			}).Error
			continue
		}
		if err != gorm.ErrRecordNotFound {
			return err
		}
		row := s
		if err := db.Create(&row).Error; err != nil {
			return err
		}
	}

	var allSurveys []models.SatisfactionSurvey
	if err := db.Find(&allSurveys).Error; err != nil {
		return err
	}
	var workers []models.User
	if err := db.Where("role = ?", models.UserRoleRailway).Find(&workers).Error; err != nil {
		return err
	}
	if len(workers) == 0 {
		return nil
	}

	levels := []string{"非常满意", "满意", "满意", "满意", "一般", "一般", "不满意", "非常不满意"}
	for _, survey := range allSurveys {
		var ratingCnt int64
		_ = db.Model(&models.SatisfactionRating{}).Where("survey_id = ?", survey.SurveyID).Count(&ratingCnt)
		if ratingCnt >= 12 {
			continue
		}
		n := 18 + rng.Intn(16)
		if n > len(workers) {
			n = len(workers)
		}
		perm := rng.Perm(len(workers))[:n]
		for _, idx := range perm {
			u := workers[idx]
			var exists int64
			_ = db.Model(&models.SatisfactionRating{}).
				Where("survey_id = ? AND user_id = ?", survey.SurveyID, u.ID).
				Count(&exists)
			if exists > 0 {
				continue
			}
			level := levels[rng.Intn(len(levels))]
			// Open surveys skew more positive.
			if survey.IsOpen && rng.Float64() < 0.55 {
				level = []string{"非常满意", "满意", "满意"}[rng.Intn(3)]
			}
			dayOffset := rng.Intn(40)
			ratingTime := survey.CreateTime.Add(time.Duration(dayOffset)*24*time.Hour +
				time.Duration(rng.Intn(10))*time.Hour)
			if ratingTime.After(time.Now()) {
				ratingTime = time.Now().Add(-time.Duration(rng.Intn(72)) * time.Hour)
			}
			phone := u.PhoneNumber
			if phone == "" {
				phone = u.Phone
			}
			rec := models.SatisfactionRating{
				UserID:            u.ID,
				SurveyID:          survey.SurveyID,
				PhoneNumber:       phone,
				SatisfactionLevel: level,
				RatingTime:        ratingTime,
			}
			if err := db.Create(&rec).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

// Off-corridor sparse check-in points (leave / business travel / transfer).
var scatteredStations = []railStation{
	{Name: "昆明南站附近驻点", Province: "云南省", City: "昆明市", County: "呈贡区", Longitude: 102.8615, Latitude: 24.8526, Altitude: 1890},
	{Name: "大理站工务点", Province: "云南省", City: "大理白族自治州", County: "大理市", Longitude: 100.2480, Latitude: 25.6860, Altitude: 1976},
	{Name: "丽江站驻地", Province: "云南省", City: "丽江市", County: "古城区", Longitude: 100.2490, Latitude: 26.8550, Altitude: 2393},
	{Name: "成都东站换乘点", Province: "四川省", City: "成都市", County: "成华区", Longitude: 104.1410, Latitude: 30.6300, Altitude: 500},
	{Name: "兰州西站驻点", Province: "甘肃省", City: "兰州市", County: "安宁区", Longitude: 103.7500, Latitude: 36.0680, Altitude: 1520},
	{Name: "乌鲁木齐站休整点", Province: "新疆维吾尔自治区", City: "乌鲁木齐市", County: "新市区", Longitude: 87.5550, Latitude: 43.8400, Altitude: 800},
	{Name: "西安北站中转点", Province: "陕西省", City: "西安市", County: "未央区", Longitude: 108.9390, Latitude: 34.3760, Altitude: 400},
	{Name: "重庆西站驻点", Province: "重庆市", City: "重庆市", County: "沙坪坝区", Longitude: 106.4350, Latitude: 29.5040, Altitude: 350},
}

func seedScatteredCheckIns(db *gorm.DB) error {
	rng := rand.New(rand.NewSource(20261107 + 31))
	var workers []models.User
	if err := db.Where("role = ? AND phone_number LIKE ?", models.UserRoleRailway, "13980%").
		Order("id ASC").Find(&workers).Error; err != nil {
		return err
	}
	if len(workers) == 0 {
		return nil
	}
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	// Pick a handful of workers for sparse off-corridor check-ins.
	pickN := 10
	if pickN > len(workers) {
		pickN = len(workers)
	}
	perm := rng.Perm(len(workers))[:pickN]
	created := 0
	for _, wi := range perm {
		u := workers[wi]
		// 3–6 historical off-corridor days + maybe today for 2 workers.
		days := 3 + rng.Intn(4)
		for d := 0; d < days; d++ {
			dayOffset := 1 + rng.Intn(40)
			if d == 0 && rng.Float64() < 0.35 {
				dayOffset = 0 // occasional today outlier
			}
			day := today.AddDate(0, 0, -dayOffset)
			st := scatteredStations[rng.Intn(len(scatteredStations))]
			var cnt int64
			_ = db.Model(&models.DailyHealthStatus{}).
				Where("user_id = ? AND check_in_date = ?", u.ID, day).
				Count(&cnt)
			if cnt > 0 {
				continue
			}
			h := true
			if rng.Float64() < 0.08 {
				h = false
			}
			checkTime := time.Date(day.Year(), day.Month(), day.Day(), 7+rng.Intn(5), rng.Intn(60), 0, 0, time.Local)
			disease := ""
			if !h {
				disease = symptomDiseases[rng.Intn(len(symptomDiseases))]
			}
			rec := models.DailyHealthStatus{
				UserID:          u.ID,
				CheckInDate:     day,
				CheckInTime:     checkTime,
				IsHealth:        &h,
				DiseaseTypeName: disease,
				Longitude:       round6(st.Longitude + (rng.Float64()-0.5)*0.04),
				Latitude:        round6(st.Latitude + (rng.Float64()-0.5)*0.03),
				Province:        st.Province,
				City:            st.City,
				County:          st.County,
				LocationName:    st.Name,
			}
			if err := db.Create(&rec).Error; err != nil {
				return err
			}
			created++
		}
	}
	fmt.Printf("scattered check-ins touched=%d\n", created)
	return nil
}

func fillStaffProfiles(db *gorm.DB) error {
	type staffSpec struct {
		phone, gender, ethnicity, edu, dept, occ string
		age                                      int
		plateau                                  time.Time
	}
	specs := []staffSpec{
		{"13800000000", "男", "汉族", "大专/本科", "工程技术部", "否", 36, time.Date(2019, 3, 1, 0, 0, 0, 0, time.Local)},
		{"13800000001", "男", "汉族", "大专/本科", "系统管理部", "否", 42, time.Date(2016, 5, 1, 0, 0, 0, 0, time.Local)},
		{"13800000002", "女", "汉族", "大专/本科", "疾病预防控制科", "医务人员", 38, time.Date(2017, 8, 12, 0, 0, 0, 0, time.Local)},
		{"13800000003", "女", "回族", "大专/本科", "格尔木铁路医院内科", "医务人员", 34, time.Date(2018, 4, 20, 0, 0, 0, 0, time.Local)},
	}
	for _, sp := range specs {
		var u models.User
		if err := db.Where("phone_number = ?", sp.phone).First(&u).Error; err != nil {
			continue
		}
		var p models.EmployeeProfile
		err := silentDB(db).Where("user_id = ?", u.ID).First(&p).Error
		plateau := sp.plateau
		birth := time.Date(time.Now().Year()-sp.age, 6, 15, 0, 0, 0, 0, time.Local)
		fields := map[string]any{
			"gender":                     sp.gender,
			"age":                        sp.age,
			"birth_date":                 birth,
			"ethnicity":                  sp.ethnicity,
			"education_level":            sp.edu,
			"department":                 sp.dept,
			"specific_occupation":        sp.occ,
			"work_on_plateau_start_date": plateau,
			"id_number":                  fmt.Sprintf("632801%04d0615%04d", 1980+sp.age%20, u.ID%10000),
			"emergency_contact_name":     "紧急联系人",
			"emergency_contact_phone":    "13820001111",
			"emergency_contact_relation": "家属",
			"height":                     165 + sp.age%15,
			"weight":                     58 + sp.age%20,
			"smoking_status":             "从不吸",
			"drinking_status":            "偶尔喝酒（小于3次/周）",
		}
		if err == gorm.ErrRecordNotFound {
			p = models.EmployeeProfile{UserID: u.ID}
			if err = db.Create(&p).Error; err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
		if err = db.Model(&p).Updates(fields).Error; err != nil {
			return err
		}
	}

	// Ensure every railway profile has non-empty required list columns.
	var profiles []models.EmployeeProfile
	if err := db.Find(&profiles).Error; err != nil {
		return err
	}
	occs := []string{"否", "否", "否", "否", "医务人员", "病原微生物检测人员", "野生动物接触相关人员", "家禽、家畜养殖人员"}
	rng := rand.New(rand.NewSource(99))
	for i := range profiles {
		p := &profiles[i]
		upd := map[string]any{}
		if p.Gender == "" {
			upd["gender"] = "男"
		}
		if p.Age == 0 {
			upd["age"] = 28 + rng.Intn(25)
		}
		if p.Ethnicity == "" {
			upd["ethnicity"] = ethnicities[rng.Intn(len(ethnicities))]
		}
		if p.EducationLevel == "" {
			upd["education_level"] = educationLevels[rng.Intn(len(educationLevels))]
		}
		if p.Department == "" {
			upd["department"] = departments[rng.Intn(len(departments))]
		}
		if p.SpecificOccupation == "" {
			upd["specific_occupation"] = occs[rng.Intn(len(occs))]
		}
		if p.WorkOnPlateauStartDate == nil {
			t := time.Date(2017+rng.Intn(7), time.Month(1+rng.Intn(12)), 1+rng.Intn(28), 0, 0, 0, 0, time.Local)
			upd["work_on_plateau_start_date"] = t
		}
		if len(upd) > 0 {
			_ = db.Model(p).Updates(upd).Error
		}
	}
	return nil
}
