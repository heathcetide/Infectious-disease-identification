// Copyright (c) 2026 heathcetide. All rights reserved.

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"gorm.io/gorm"
)

var clinicalDiseaseNames = []string{
	"流感", "新型冠状病毒感染", "感染性腹泻", "结核病", "斑疹伤寒",
	"登革热（蚊媒传染病）", "疟疾（蚊媒传染病）", "森林脑炎（蜱媒传染病）",
	"发热伴血小板减少综合征（蜱媒传染病）", "炭疽", "鼠疫", "流行性出血热",
}

var hospitalCatalog = []models.Hospital{
	{HospitalName: "青海省交通医院西宁院区", HospitalPhoneNumber: "0971-6268120", Address: "青海省西宁市城东区互助中路12号"},
	{HospitalName: "格尔木市人民医院", HospitalPhoneNumber: "0979-8412120", Address: "青海省格尔木市昆仑路88号"},
	{HospitalName: "青藏铁路疾病预防控制站（格尔木）", HospitalPhoneNumber: "0979-8458000", Address: "青海省格尔木市青藏路"},
	{HospitalName: "那曲市人民医院", HospitalPhoneNumber: "0896-3822120", Address: "西藏自治区那曲市色尼区浙江中路"},
	{HospitalName: "西藏自治区人民医院", HospitalPhoneNumber: "0891-6332462", Address: "西藏自治区拉萨市城关区林廓北路"},
	{HospitalName: "拉萨市人民医院", HospitalPhoneNumber: "0891-6322202", Address: "西藏自治区拉萨市城关区北京中路"},
	{HospitalName: "海西州人民医院", HospitalPhoneNumber: "0977-8223120", Address: "青海省海西蒙古族藏族自治州德令哈市"},
	{HospitalName: "安多县人民医院", HospitalPhoneNumber: "0896-3612120", Address: "西藏自治区那曲市安多县"},
}

func seedClinicalDemo(db *gorm.DB) error {
	if err := replaceHospitals(db); err != nil {
		return err
	}
	if err := clearPropagandaMaterials(db); err != nil {
		return err
	}
	if err := ensureWorkEnvironments(db); err != nil {
		return err
	}
	if err := seedDiagnosesAndLabs(db); err != nil {
		return err
	}
	if err := seedSatisfactionDemo(db); err != nil {
		return err
	}
	// Normalize leftover disease labels used by dashboard charts.
	_ = db.Model(&models.DailyHealthStatus{}).
		Where("disease_type_name IN ?", []string{"高原反应相关不适", "上呼吸道感染"}).
		Update("disease_type_name", "流感").Error
	_ = db.Model(&models.DailyHealthStatus{}).
		Where("disease_type_name = ?", "流行性热出血").
		Update("disease_type_name", "流行性出血热").Error
	_ = db.Model(&models.DiseaseScoring{}).
		Where("disease_type_name IN ?", []string{"高原反应相关不适", "上呼吸道感染"}).
		Update("disease_type_name", "流感").Error
	_ = db.Model(&models.DiseaseScoring{}).
		Where("disease_type_name = ?", "流行性热出血").
		Update("disease_type_name", "流行性出血热").Error
	return nil
}

func replaceHospitals(db *gorm.DB) error {
	_ = db.Where("1 = 1").Delete(&models.Hospital{}).Error
	for _, h := range hospitalCatalog {
		row := h
		if err := db.Create(&row).Error; err != nil {
			return err
		}
	}
	fmt.Printf("hospitals replaced: %d\n", len(hospitalCatalog))
	return nil
}

func clearPropagandaMaterials(db *gorm.DB) error {
	res := db.Where("1 = 1").Delete(&models.DiseaseLearningMaterial{})
	if res.Error != nil {
		return res.Error
	}
	_ = db.Where("1 = 1").Delete(&models.MaterialLearningRecord{}).Error
	fmt.Printf("propaganda materials cleared: %d\n", res.RowsAffected)
	if err := seedRealPropagandaMaterials(db); err != nil {
		return err
	}
	return seedPropagandaMediaToOSS(db)
}

// Official China CDC / NHC public education articles (real URLs).
func seedRealPropagandaMaterials(db *gorm.DB) error {
	type mat struct {
		disease, title, typ, link string
		seq                       int
		pub                       time.Time
	}
	items := []mat{
		{"流感", "秋冬流感季：流感疫苗怎么打、重点人群如何防护（中国疾控中心）", "网页链接", "https://www.chinacdc.cn/jkkp/crb/bcr/202611/t20261104_1842705.html", 1, time.Date(2026, 11, 4, 0, 0, 0, 0, time.Local)},
		{"流感", "中国流感疫苗预防接种技术指南（2026—2027）（中国疾控中心）", "网页链接", "https://www.chinacdc.cn/jkyj/mygh02/jswj2026/202609/t20260902_1839739.html", 2, time.Date(2026, 9, 2, 0, 0, 0, 0, time.Local)},
		{"流感", "这些人得流感后更容易发展为重症：如何预防与应对（中国疾控中心）", "网页链接", "https://www.chinacdc.cn/jkkp/crb/bcr/202511/t20251103_313231.html", 3, time.Date(2025, 11, 3, 0, 0, 0, 0, time.Local)},
		{"流感", "除了打疫苗，这几件事也是预防流感的关键（中国疾控中心）", "网页链接", "https://www.chinacdc.cn/jkkp/crb/bcr/202512/t20251222_314212.html", 4, time.Date(2025, 12, 22, 0, 0, 0, 0, time.Local)},
		{"感染性腹泻", "细菌性腹泻高发季：病从口入怎么防（中国疾控中心）", "网页链接", "https://www.chinacdc.cn/jkkp/crb/bcr/202507/t20250703_308170.html", 1, time.Date(2025, 7, 3, 0, 0, 0, 0, time.Local)},
		{"感染性腹泻", "志贺氏菌的预防与控制（中国疾控中心）", "网页链接", "https://www.chinacdc.cn/jkkp/crb/ycr/202608/t20260827_1839558.html", 2, time.Date(2026, 8, 27, 0, 0, 0, 0, time.Local)},
		{"鼠疫", "预防鼠疫，你需要知道的这些事（中国疾控中心）", "网页链接", "https://www.chinacdc.cn/jkkp/ggws/syyr/202408/t20240825_295563.html", 1, time.Date(2023, 9, 27, 0, 0, 0, 0, time.Local)},
		{"炭疽", "正确认识炭疽，科学预防炭疽（中国疾控中心）", "网页链接", "https://www.chinacdc.cn/jkkp/crb/ycr/202408/t20240822_294959.html", 1, time.Date(2024, 8, 16, 0, 0, 0, 0, time.Local)},
		{"炭疽", "这种古老的传染病从未离开：认识肺炭疽（中国疾控中心）", "网页链接", "https://www.chinacdc.cn/jkkp/crb/ycr/202610/t20261019_1842361.html", 2, time.Date(2026, 10, 19, 0, 0, 0, 0, time.Local)},
		{"结核病", "肺结核核心知识：传播途径、就诊时机与规范治疗（中国疾控中心）", "网页链接", "https://www.chinacdc.cn/jkkp/crb/ycr/202409/t20240925_300847.html", 1, time.Date(2024, 9, 25, 0, 0, 0, 0, time.Local)},
		{"登革热（蚊媒传染病）", "登革热是什么、有哪些症状、如何防蚊（中国疾控中心）", "网页链接", "https://www.chinacdc.cn/jkkp/crb/ycr/202507/t20250716_308513.html", 1, time.Date(2025, 7, 16, 0, 0, 0, 0, time.Local)},
		{"登革热（蚊媒传染病）", "夏季健康提示：登革热等蚊媒传染病防护（中国疾控中心）", "网页链接", "https://www.chinacdc.cn/jkts/202606/t20260610_1836703.html", 2, time.Date(2026, 6, 10, 0, 0, 0, 0, time.Local)},
		{"疟疾（蚊媒传染病）", "严防境外疟疾输入，守护跨境出行健康（中国疾控中心）", "网页链接", "https://www.chinacdc.cn/jkkp/crb/ycr/202608/t20260826_1839384.html", 1, time.Date(2026, 8, 26, 0, 0, 0, 0, time.Local)},
		{"疟疾（蚊媒传染病）", "2026年全国疟疾宣传日：主动告知旅居史（中国疾控中心）", "网页链接", "https://www.chinacdc.cn/jkkp/crb/ycr/202604/t20260429_1835401.html", 2, time.Date(2026, 4, 29, 0, 0, 0, 0, time.Local)},
		{"发热伴血小板减少综合征（蜱媒传染病）", "发热伴血小板减少综合征：蜱叮咬传播与个人防护（中国疾控中心）", "网页链接", "https://www.chinacdc.cn/jkkp/crb/ycr/202603/t20260325_315906.html", 1, time.Date(2026, 3, 25, 0, 0, 0, 0, time.Local)},
		{"发热伴血小板减少综合征（蜱媒传染病）", "秋季健康提示：发热伴血小板减少综合征防护要点（中国疾控中心）", "网页链接", "https://www.chinacdc.cn/jkts/202609/t20260907_1839864.html", 2, time.Date(2026, 9, 7, 0, 0, 0, 0, time.Local)},
		{"森林脑炎（蜱媒传染病）", "野外作业防蜱叮咬与蜱媒传染病防护提示（中国疾控中心）", "网页链接", "https://www.chinacdc.cn/jkts/202606/t20260610_1836703.html", 1, time.Date(2026, 6, 10, 0, 0, 0, 0, time.Local)},
		{"新型冠状病毒感染", "新型冠状病毒感染不同风险人群防护指南（国家卫生健康委）", "网页链接", "http://www.nhc.gov.cn/jkj/c100063/202001/d7dca5a82e5f4253af8a19c3921ecd29.shtml", 1, time.Date(2020, 1, 31, 0, 0, 0, 0, time.Local)},
		{"新型冠状病毒感染", "公众科学戴口罩指引（国家卫生健康委）", "网页链接", "https://www.nhc.gov.cn/jkj/c100063/202003/3d91f77851b045519466ee954cb11dce.shtml", 2, time.Date(2020, 3, 18, 0, 0, 0, 0, time.Local)},
		{"新型冠状病毒感染", "公共场所新型冠状病毒感染卫生防护指南（国家卫生健康委）", "网页链接", "https://www.nhc.gov.cn/jkj/c100063/202001/dbe42bfc662a4631b053d42a102edec2.shtml", 3, time.Date(2020, 1, 31, 0, 0, 0, 0, time.Local)},
		{"斑疹伤寒", "夏秋蚊虫活跃期：常见虫媒传染病健康提示（中国疾控中心）", "网页链接", "https://www.chinacdc.cn/jkts/202606/t20260610_1836703.html", 1, time.Date(2026, 6, 10, 0, 0, 0, 0, time.Local)},
		{"流行性热出血", "流行性出血热早期症状与预防要点（中国疾控中心）", "网页链接", "https://www.chinacdc.cn/jkkp/crb/ycr/index_2.html", 1, time.Date(2025, 2, 20, 0, 0, 0, 0, time.Local)},
	}

	typeIDByName := map[string]int{}
	var dts []models.DiseaseType
	_ = db.Find(&dts).Error
	for _, d := range dts {
		typeIDByName[d.DiseaseTypeName] = int(d.DiseaseTypeID)
	}

	off := false
	created := 0
	for _, it := range items {
		pub := it.pub
		m := models.DiseaseLearningMaterial{
			DiseaseTypeID:   typeIDByName[it.disease],
			DiseaseTypeName: it.disease,
			Title:           it.title,
			MaterialType:    it.typ,
			Link:            it.link,
			SequenceNumber:  it.seq,
			PublishDate:     &pub,
			IsDeleted:       &off,
		}
		if err := db.Create(&m).Error; err != nil {
			return err
		}
		created++
	}
	fmt.Printf("real propaganda materials seeded: %d\n", created)
	return nil
}

func ensureWorkEnvironments(db *gorm.DB) error {
	rng := rand.New(rand.NewSource(20261107 + 77))
	var workers []models.User
	if err := db.Where("role = ?", models.UserRoleRailway).Limit(280).Find(&workers).Error; err != nil {
		return err
	}
	yearMonth := time.Now().Format("2006-01")
	created := 0
	for i, u := range workers {
		var cnt int64
		_ = db.Model(&models.WorkEnvironmentInfo{}).
			Where("user_id = ? AND year_month = ?", u.ID, yearMonth).Count(&cnt)
		if cnt > 0 {
			continue
		}
		st := qingzangStations[i%len(qingzangStations)]
		if i%6 == 0 {
			st = nationwideStations[i%len(nationwideStations)]
		}
		env := models.WorkEnvironmentInfo{
			UserID:                  u.ID,
			YearMonth:               yearMonth,
			SubmissionTime:          time.Now().Add(-time.Duration(rng.Intn(20)) * 24 * time.Hour),
			WorkStationAltitude:     st.Altitude + rng.Intn(40) - 10,
			WorkStationName:         st.Name,
			DormitoryBuildingNumber: 1 + rng.Intn(6),
			DormitoryRoomNumber:     101 + rng.Intn(40),
			DormitoryOccupants:      4 + rng.Intn(4),
			DormitoryArea:           28 + float64(rng.Intn(18)),
			WindowOpenArea:          1.2 + rng.Float64()*2,
			VentilationCondition:    []string{"良好", "一般", "较差"}[rng.Intn(3)],
			AvgWorkRoomPeople:       6 + rng.Intn(10),
			AvgMonthlyPressure:      480 + float64(rng.Intn(120)),
			MaxTemperature:          6 + float64(rng.Intn(16)),
			MinTemperature:          -16 + float64(rng.Intn(14)),
		}
		if err := db.Create(&env).Error; err != nil {
			return err
		}
		created++
	}
	fmt.Printf("work environments created=%d\n", created)
	return nil
}

func seedDiagnosesAndLabs(db *gorm.DB) error {
	rng := rand.New(rand.NewSource(20261107 + 55))
	var existingDiag int64
	_ = db.Model(&models.DiagnosisResults{}).Count(&existingDiag)
	if existingDiag >= 80 {
		fmt.Println("diagnosis/lab already seeded, skip")
		return nil
	}

	var workers []models.User
	if err := db.Where("role = ?", models.UserRoleRailway).Order("id ASC").Limit(220).Find(&workers).Error; err != nil {
		return err
	}
	var medical models.User
	_ = silentDB(db).Where("role = ?", models.UserRoleMedical).First(&medical).Error

	outcomes := []string{"治愈", "治愈", "未愈", "好转", "死亡"}
	methods := []string{"主动监测", "被动监测", "健康打卡预警", "门诊就诊", "住院排查"}
	createdD, createdL := 0, 0

	for i, u := range workers {
		if i%2 != 0 && rng.Float64() > 0.45 {
			continue
		}
		disease := clinicalDiseaseNames[rng.Intn(len(clinicalDiseaseNames))]
		hosp := hospitalCatalog[rng.Intn(len(hospitalCatalog))].HospitalName
		day := time.Now().AddDate(0, 0, -(1 + rng.Intn(50)))
		admit := day
		dis := day.AddDate(0, 0, 3+rng.Intn(12))
		sub := day
		outcome := outcomes[rng.Intn(len(outcomes))]
		if outcome == "死亡" && rng.Float64() > 0.08 {
			outcome = "未愈"
		}
		rec := models.DiagnosisResults{
			UserID:             u.ID,
			SubmissionTime:     &sub,
			SubmissionUserID:   medical.ID,
			SubmissionUserName: medical.Name,
			SubmissionUserType: "专职医护人员",
			DiseaseType:        disease,
			DiscoveryMethod:    methods[rng.Intn(len(methods))],
			DiseaseOutcome:     outcome,
			HospitalName:       hosp,
			AdmissionDate:      &admit,
			DischargeDate:      &dis,
			AdmissionSymptomsAndSigns: "发热、乏力、咳嗽",
			RegistrationClassification: "临床诊断病例",
		}
		if disease == "鼠疫" {
			rec.PlagueSubtype = []string{"腺鼠疫", "肺鼠疫"}[rng.Intn(2)]
		}
		if disease == "炭疽" {
			rec.AnthraxSubtype = []string{"皮肤炭疽", "肠炭疽"}[rng.Intn(2)]
		}
		if err := db.Create(&rec).Error; err != nil {
			return err
		}
		createdD++

		// Matching disease scoring for accuracy charts.
		_ = db.Create(&models.DiseaseScoring{
			StatusID:        0,
			SymptomTime:     &day,
			DiseaseTypeName: disease,
			Score:           55 + rng.Float64()*40,
		}).Error

		if i%3 == 0 || rng.Float64() < 0.4 {
			antigen := rng.Float64() < 0.7
			na := rng.Float64() < 0.55
			culture := rng.Float64() < 0.2
			sero := rng.Float64() < 0.35
			upload := day.Add(24 * time.Hour)
			lab := models.LabTestReport{
				UserID:                      u.ID,
				UploadDate:                  &upload,
				IsVirusAntigenTestDone:     &antigen,
				IsVirusNucleicAcidTestDone:  &na,
				IsVirusCultureIsolationDone: &culture,
				IsSerologicalTestDone:       &sero,
				PathogenicTestResults:       fmt.Sprintf("%s 相关病原学检测：%s", disease, []string{"阳性", "阴性", "待复核"}[rng.Intn(3)]),
			}
			if err := db.Create(&lab).Error; err != nil {
				return err
			}
			createdL++
		}
	}
	fmt.Printf("diagnosis=%d lab=%d\n", createdD, createdL)
	return nil
}
