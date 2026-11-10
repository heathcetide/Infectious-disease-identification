// Copyright (c) 2026 heathcetide. All rights reserved.

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"

	"github.com/heathcetide/infectious-disease/internal/models"
)

var clinicalDiseaseNames = []string{
	"流感", "新型冠状病毒感染", "感染性腹泻", "结核病", "斑疹伤寒",
	"登革热（蚊媒传染病）", "疟疾（蚊媒传染病）", "森林脑炎（蜱媒传染病）",
	"发热伴血小板减少综合征（蜱媒传染病）", "炭疽", "鼠疫", "流行性出血热",
}

var hospitalCatalog = []models.Hospital{
	// 西宁 / 海东
	{HospitalName: "青海省交通医院西宁院区", HospitalPhoneNumber: "0971-6268120", Address: "青海省西宁市城东区互助中路12号", Latitude: 36.6175, Longitude: 101.8234},
	{HospitalName: "青海大学附属医院", HospitalPhoneNumber: "0971-6231557", Address: "青海省西宁市城西区同仁路29号", Latitude: 36.6170, Longitude: 101.7658},
	{HospitalName: "青海省人民医院", HospitalPhoneNumber: "0971-8066186", Address: "青海省西宁市城东区共和路2号", Latitude: 36.6212, Longitude: 101.7986},
	{HospitalName: "西宁市第一人民医院", HospitalPhoneNumber: "0971-8239173", Address: "青海省西宁市城中区南大街", Latitude: 36.6178, Longitude: 101.7783},
	{HospitalName: "青海红十字医院", HospitalPhoneNumber: "0971-8247738", Address: "青海省西宁市城中区南大街55号", Latitude: 36.6165, Longitude: 101.7801},
	{HospitalName: "海东市人民医院", HospitalPhoneNumber: "0972-8612120", Address: "青海省海东市乐都区滨河中路", Latitude: 36.4821, Longitude: 102.4015},
	// 格尔木 / 海西沿线
	{HospitalName: "格尔木市人民医院", HospitalPhoneNumber: "0979-8412120", Address: "青海省格尔木市昆仑路88号", Latitude: 36.4020, Longitude: 94.9030},
	{HospitalName: "格尔木市第三人民医院", HospitalPhoneNumber: "0979-8422880", Address: "青海省格尔木市黄河路", Latitude: 36.4068, Longitude: 94.9152},
	{HospitalName: "青藏铁路疾病预防控制站（格尔木）", HospitalPhoneNumber: "0979-8458000", Address: "青海省格尔木市青藏路", Latitude: 36.4155, Longitude: 94.9088},
	{HospitalName: "格尔木市中蒙医医院", HospitalPhoneNumber: "0979-8413668", Address: "青海省格尔木市柴达木中路", Latitude: 36.3985, Longitude: 94.9011},
	{HospitalName: "海西州人民医院", HospitalPhoneNumber: "0977-8223120", Address: "青海省海西蒙古族藏族自治州德令哈市", Latitude: 37.3697, Longitude: 97.3701},
	{HospitalName: "德令哈市蒙藏医院", HospitalPhoneNumber: "0977-8225600", Address: "青海省德令哈市黄河路", Latitude: 37.3742, Longitude: 97.3608},
	{HospitalName: "乌兰县人民医院", HospitalPhoneNumber: "0977-8242120", Address: "青海省海西州乌兰县希里沟镇", Latitude: 36.9305, Longitude: 98.4796},
	{HospitalName: "都兰县人民医院", HospitalPhoneNumber: "0977-8452120", Address: "青海省海西州都兰县察汗乌苏镇", Latitude: 36.2984, Longitude: 98.0892},
	// 那曲 / 藏北
	{HospitalName: "那曲市人民医院", HospitalPhoneNumber: "0896-3822120", Address: "西藏自治区那曲市色尼区浙江中路", Latitude: 31.4762, Longitude: 92.0512},
	{HospitalName: "那曲市藏医院", HospitalPhoneNumber: "0896-3823456", Address: "西藏自治区那曲市色尼区山东路", Latitude: 31.4801, Longitude: 92.0488},
	{HospitalName: "安多县人民医院", HospitalPhoneNumber: "0896-3612120", Address: "西藏自治区那曲市安多县", Latitude: 32.2642, Longitude: 91.6815},
	{HospitalName: "聂荣县人民医院", HospitalPhoneNumber: "0896-3662120", Address: "西藏自治区那曲市聂荣县", Latitude: 32.1078, Longitude: 92.3035},
	{HospitalName: "比如县人民医院", HospitalPhoneNumber: "0896-3642120", Address: "西藏自治区那曲市比如县", Latitude: 31.4796, Longitude: 93.6802},
	{HospitalName: "班戈县人民医院", HospitalPhoneNumber: "0896-3632120", Address: "西藏自治区那曲市班戈县", Latitude: 31.3920, Longitude: 90.0186},
	// 拉萨 / 日喀则 / 山南
	{HospitalName: "西藏自治区人民医院", HospitalPhoneNumber: "0891-6332462", Address: "西藏自治区拉萨市城关区林廓北路", Latitude: 29.6573, Longitude: 91.1325},
	{HospitalName: "拉萨市人民医院", HospitalPhoneNumber: "0891-6322202", Address: "西藏自治区拉萨市城关区北京中路", Latitude: 29.6548, Longitude: 91.1142},
	{HospitalName: "西藏自治区藏医院", HospitalPhoneNumber: "0891-6323115", Address: "西藏自治区拉萨市城关区娘热路", Latitude: 29.6685, Longitude: 91.1120},
	{HospitalName: "西藏军区总医院", HospitalPhoneNumber: "0891-6255120", Address: "西藏自治区拉萨市城关区金珠西路", Latitude: 29.6502, Longitude: 91.0955},
	{HospitalName: "日喀则市人民医院", HospitalPhoneNumber: "0892-8822120", Address: "西藏自治区日喀则市桑珠孜区青岛路", Latitude: 29.2672, Longitude: 88.8808},
	{HospitalName: "山南市人民医院", HospitalPhoneNumber: "0893-7822120", Address: "西藏自治区山南市乃东区浙江大道", Latitude: 29.2370, Longitude: 91.7715},
	{HospitalName: "林芝市人民医院", HospitalPhoneNumber: "0894-5822120", Address: "西藏自治区林芝市巴宜区广东路", Latitude: 29.6544, Longitude: 94.3615},
	{HospitalName: "昌都市人民医院", HospitalPhoneNumber: "0895-4822120", Address: "西藏自治区昌都市卡若区四川中路", Latitude: 31.1405, Longitude: 97.1785},
	// 甘肃 / 兰青线衔接
	{HospitalName: "兰州大学第一医院", HospitalPhoneNumber: "0931-8356240", Address: "甘肃省兰州市城关区东岗西路1号", Latitude: 36.0512, Longitude: 103.8556},
	{HospitalName: "甘肃省人民医院", HospitalPhoneNumber: "0931-8281635", Address: "甘肃省兰州市城关区东岗西路204号", Latitude: 36.0548, Longitude: 103.8662},
	{HospitalName: "中国铁路兰州局集团有限公司兰州铁路卫生监督所", HospitalPhoneNumber: "0931-4922120", Address: "甘肃省兰州市城关区火车站东路", Latitude: 36.0345, Longitude: 103.8502},
	{HospitalName: "天祝藏族自治县人民医院", HospitalPhoneNumber: "0935-3122120", Address: "甘肃省武威市天祝县华藏寺镇", Latitude: 36.9718, Longitude: 103.1420},
	{HospitalName: "武威市凉州医院", HospitalPhoneNumber: "0935-2212120", Address: "甘肃省武威市凉州区东大街", Latitude: 37.9283, Longitude: 102.6380},
	{HospitalName: "嘉峪关市第一人民医院", HospitalPhoneNumber: "0937-6222120", Address: "甘肃省嘉峪关市雄关东路", Latitude: 39.7865, Longitude: 98.2892},
	{HospitalName: "敦煌市医院", HospitalPhoneNumber: "0937-8822120", Address: "甘肃省酒泉市敦煌市阳关中路", Latitude: 40.1415, Longitude: 94.6628},
	// 四川 / 成昆及进藏通道
	{HospitalName: "四川大学华西医院", HospitalPhoneNumber: "028-85422114", Address: "四川省成都市武侯区国学巷37号", Latitude: 30.6425, Longitude: 104.0658},
	{HospitalName: "成都市第三人民医院", HospitalPhoneNumber: "028-61318682", Address: "四川省成都市青羊区青龙街82号", Latitude: 30.6702, Longitude: 104.0555},
	{HospitalName: "中国铁路成都局集团有限公司疾病预防控制所", HospitalPhoneNumber: "028-86432200", Address: "四川省成都市金牛区一环路北二段", Latitude: 30.6915, Longitude: 104.0622},
	{HospitalName: "雅安市人民医院", HospitalPhoneNumber: "0835-2222120", Address: "四川省雅安市雨城区正和路", Latitude: 30.0105, Longitude: 103.0412},
	{HospitalName: "康定市人民医院", HospitalPhoneNumber: "0836-2832120", Address: "四川省甘孜州康定市炉城南路", Latitude: 30.0495, Longitude: 101.9648},
	// 陕西 / 枢纽城市
	{HospitalName: "西安交通大学第一附属医院", HospitalPhoneNumber: "029-85323112", Address: "陕西省西安市雁塔西路277号", Latitude: 34.2185, Longitude: 108.9422},
	{HospitalName: "陕西省人民医院", HospitalPhoneNumber: "029-85251331", Address: "陕西省西安市碑林区友谊西路256号", Latitude: 34.2412, Longitude: 108.9285},
	{HospitalName: "中国铁路西安局集团有限公司疾病预防控制所", HospitalPhoneNumber: "029-82322200", Address: "陕西省西安市新城区太华南路", Latitude: 34.2788, Longitude: 108.9735},
	// 北京 / 铁路系统示范点（开发机定位常见）
	{HospitalName: "中国铁路北京局集团有限公司疾病预防控制所", HospitalPhoneNumber: "010-51821200", Address: "北京市东城区毛家湾胡同甲13号", Latitude: 39.9042, Longitude: 116.4074},
	{HospitalName: "北京佑安医院", HospitalPhoneNumber: "010-83997146", Address: "北京市丰台区右安门外西头条8号", Latitude: 39.8665, Longitude: 116.3612},
	{HospitalName: "首都医科大学附属北京地坛医院", HospitalPhoneNumber: "010-84322000", Address: "北京市朝阳区京顺东街8号", Latitude: 40.0128, Longitude: 116.4185},
}

func seedClinicalDemo(db *gorm.DB) error {
	if err := ensureHospitalsUpsert(db); err != nil {
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

// ensureHospitalsUpsert inserts missing hospitals by name and refreshes coords/phone for known rows.
func ensureHospitalsUpsert(db *gorm.DB) error {
	added, updated := 0, 0
	for _, h := range hospitalCatalog {
		var existing models.Hospital
		err := db.Where("hospital_name = ?", h.HospitalName).First(&existing).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			row := h
			if err = db.Create(&row).Error; err != nil {
				return err
			}
			added++
			continue
		}
		if err != nil {
			return err
		}
		patch := map[string]interface{}{
			"hospital_phone_number": h.HospitalPhoneNumber,
			"address":               h.Address,
			"latitude":              h.Latitude,
			"longitude":             h.Longitude,
		}
		if err := db.Model(&existing).Updates(patch).Error; err != nil {
			return err
		}
		updated++
	}
	fmt.Printf("hospitals upsert: added=%d updated=%d catalog=%d\n", added, updated, len(hospitalCatalog))
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
			UserID:                     u.ID,
			SubmissionTime:             &sub,
			SubmissionUserID:           medical.ID,
			SubmissionUserName:         medical.Name,
			SubmissionUserType:         "专职医护人员",
			DiseaseType:                disease,
			DiscoveryMethod:            methods[rng.Intn(len(methods))],
			DiseaseOutcome:             outcome,
			HospitalName:               hosp,
			AdmissionDate:              &admit,
			DischargeDate:              &dis,
			AdmissionSymptomsAndSigns:  "发热、乏力、咳嗽",
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
				IsVirusAntigenTestDone:      &antigen,
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
