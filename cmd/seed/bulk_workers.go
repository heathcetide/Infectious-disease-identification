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

const targetRailwayWorkers = 12000

// Nationwide sparse stations (leave / transfer / affiliated projects).
var nationwideStations = []railStation{
	{Name: "昆明南站附近驻点", Province: "云南省", City: "昆明市", County: "呈贡区", Longitude: 102.8615, Latitude: 24.8526, Altitude: 1890},
	{Name: "大理站工务点", Province: "云南省", City: "大理白族自治州", County: "大理市", Longitude: 100.2480, Latitude: 25.6860, Altitude: 1976},
	{Name: "丽江站驻地", Province: "云南省", City: "丽江市", County: "古城区", Longitude: 100.2490, Latitude: 26.8550, Altitude: 2393},
	{Name: "景洪站驻点", Province: "云南省", City: "西双版纳傣族自治州", County: "景洪市", Longitude: 100.7980, Latitude: 22.0020, Altitude: 550},
	{Name: "楚雄工务驻点", Province: "云南省", City: "楚雄彝族自治州", County: "楚雄市", Longitude: 101.5460, Latitude: 25.0320, Altitude: 1770},
	{Name: "曲靖西站驻点", Province: "云南省", City: "曲靖市", County: "麒麟区", Longitude: 103.7960, Latitude: 25.4960, Altitude: 1860},
	{Name: "成都东站换乘点", Province: "四川省", City: "成都市", County: "成华区", Longitude: 104.1410, Latitude: 30.6300, Altitude: 500},
	{Name: "西昌站工区", Province: "四川省", City: "凉山彝族自治州", County: "西昌市", Longitude: 102.2640, Latitude: 27.8950, Altitude: 1500},
	{Name: "攀枝花站驻点", Province: "四川省", City: "攀枝花市", County: "东区", Longitude: 101.7190, Latitude: 26.5820, Altitude: 1200},
	{Name: "广元站工区", Province: "四川省", City: "广元市", County: "利州区", Longitude: 105.8440, Latitude: 32.4350, Altitude: 490},
	{Name: "重庆西站驻点", Province: "重庆市", City: "重庆市", County: "沙坪坝区", Longitude: 106.4350, Latitude: 29.5040, Altitude: 350},
	{Name: "万州北站驻点", Province: "重庆市", City: "重庆市", County: "万州区", Longitude: 108.4080, Latitude: 30.8070, Altitude: 260},
	{Name: "贵阳北站驻点", Province: "贵州省", City: "贵阳市", County: "观山湖区", Longitude: 106.6730, Latitude: 26.6470, Altitude: 1100},
	{Name: "遵义站驻点", Province: "贵州省", City: "遵义市", County: "汇川区", Longitude: 106.9270, Latitude: 27.7370, Altitude: 900},
	{Name: "南宁东站驻点", Province: "广西壮族自治区", City: "南宁市", County: "青秀区", Longitude: 108.4140, Latitude: 22.8400, Altitude: 80},
	{Name: "柳州站工区", Province: "广西壮族自治区", City: "柳州市", County: "城中区", Longitude: 109.4220, Latitude: 24.3260, Altitude: 90},
	{Name: "广州南站中转点", Province: "广东省", City: "广州市", County: "番禺区", Longitude: 113.2690, Latitude: 22.9890, Altitude: 20},
	{Name: "深圳北站驻点", Province: "广东省", City: "深圳市", County: "龙华区", Longitude: 114.0290, Latitude: 22.6100, Altitude: 50},
	{Name: "韶关站驻点", Province: "广东省", City: "韶关市", County: "浈江区", Longitude: 113.5970, Latitude: 24.8110, Altitude: 70},
	{Name: "海口东站驻点", Province: "海南省", City: "海口市", County: "美兰区", Longitude: 110.3600, Latitude: 20.0300, Altitude: 10},
	{Name: "三亚站驻点", Province: "海南省", City: "三亚市", County: "吉阳区", Longitude: 109.5080, Latitude: 18.2970, Altitude: 10},
	{Name: "长沙南站驻点", Province: "湖南省", City: "长沙市", County: "雨花区", Longitude: 113.0660, Latitude: 28.1360, Altitude: 60},
	{Name: "怀化南站驻点", Province: "湖南省", City: "怀化市", County: "鹤城区", Longitude: 109.9980, Latitude: 27.5490, Altitude: 260},
	{Name: "武汉站驻点", Province: "湖北省", City: "武汉市", County: "洪山区", Longitude: 114.4210, Latitude: 30.6070, Altitude: 30},
	{Name: "宜昌东站驻点", Province: "湖北省", City: "宜昌市", County: "伍家岗区", Longitude: 111.3590, Latitude: 30.6550, Altitude: 80},
	{Name: "郑州东站中转点", Province: "河南省", City: "郑州市", County: "金水区", Longitude: 113.7760, Latitude: 34.7600, Altitude: 110},
	{Name: "洛阳龙门站驻点", Province: "河南省", City: "洛阳市", County: "洛龙区", Longitude: 112.4510, Latitude: 34.5930, Altitude: 150},
	{Name: "西安北站中转点", Province: "陕西省", City: "西安市", County: "未央区", Longitude: 108.9390, Latitude: 34.3760, Altitude: 400},
	{Name: "宝鸡南站驻点", Province: "陕西省", City: "宝鸡市", County: "渭滨区", Longitude: 107.1700, Latitude: 34.3370, Altitude: 620},
	{Name: "延安站驻点", Province: "陕西省", City: "延安市", County: "宝塔区", Longitude: 109.4900, Latitude: 36.5850, Altitude: 960},
	{Name: "兰州西站驻点", Province: "甘肃省", City: "兰州市", County: "安宁区", Longitude: 103.7500, Latitude: 36.0680, Altitude: 1520},
	{Name: "天水南站驻点", Province: "甘肃省", City: "天水市", County: "麦积区", Longitude: 105.8930, Latitude: 34.5630, Altitude: 1140},
	{Name: "嘉峪关站工区", Province: "甘肃省", City: "嘉峪关市", County: "嘉峪关市", Longitude: 98.2770, Latitude: 39.7860, Altitude: 1520},
	{Name: "敦煌站驻点", Province: "甘肃省", City: "酒泉市", County: "敦煌市", Longitude: 94.6620, Latitude: 40.1420, Altitude: 1140},
	{Name: "银川站驻点", Province: "宁夏回族自治区", City: "银川市", County: "金凤区", Longitude: 106.1680, Latitude: 38.4870, Altitude: 1110},
	{Name: "中卫站驻点", Province: "宁夏回族自治区", City: "中卫市", County: "沙坡头区", Longitude: 105.1970, Latitude: 37.5000, Altitude: 1220},
	{Name: "呼和浩特东站驻点", Province: "内蒙古自治区", City: "呼和浩特市", County: "新城区", Longitude: 111.7670, Latitude: 40.8490, Altitude: 1050},
	{Name: "包头站驻点", Province: "内蒙古自治区", City: "包头市", County: "昆都仑区", Longitude: 109.8220, Latitude: 40.6570, Altitude: 1060},
	{Name: "乌鲁木齐站休整点", Province: "新疆维吾尔自治区", City: "乌鲁木齐市", County: "新市区", Longitude: 87.5550, Latitude: 43.8400, Altitude: 800},
	{Name: "吐鲁番北站驻点", Province: "新疆维吾尔自治区", City: "吐鲁番市", County: "高昌区", Longitude: 89.1900, Latitude: 43.0700, Altitude: 30},
	{Name: "哈密站工区", Province: "新疆维吾尔自治区", City: "哈密市", County: "伊州区", Longitude: 93.5130, Latitude: 42.8330, Altitude: 760},
	{Name: "喀什站驻点", Province: "新疆维吾尔自治区", City: "喀什地区", County: "喀什市", Longitude: 75.9890, Latitude: 39.4680, Altitude: 1290},
	{Name: "库尔勒站驻点", Province: "新疆维吾尔自治区", City: "巴音郭楞蒙古自治州", County: "库尔勒市", Longitude: 86.1460, Latitude: 41.7630, Altitude: 930},
	{Name: "哈尔滨西站驻点", Province: "黑龙江省", City: "哈尔滨市", County: "南岗区", Longitude: 126.5780, Latitude: 45.7060, Altitude: 140},
	{Name: "齐齐哈尔站驻点", Province: "黑龙江省", City: "齐齐哈尔市", County: "铁锋区", Longitude: 123.9180, Latitude: 47.3550, Altitude: 150},
	{Name: "长春西站驻点", Province: "吉林省", City: "长春市", County: "绿园区", Longitude: 125.1970, Latitude: 43.8810, Altitude: 220},
	{Name: "吉林站驻点", Province: "吉林省", City: "吉林市", County: "昌邑区", Longitude: 126.5700, Latitude: 43.8700, Altitude: 190},
	{Name: "沈阳北站驻点", Province: "辽宁省", City: "沈阳市", County: "沈河区", Longitude: 123.3840, Latitude: 41.8150, Altitude: 50},
	{Name: "大连北站驻点", Province: "辽宁省", City: "大连市", County: "甘井子区", Longitude: 121.6030, Latitude: 39.0070, Altitude: 30},
	{Name: "北京西站中转点", Province: "北京市", City: "北京市", County: "丰台区", Longitude: 116.3210, Latitude: 39.8950, Altitude: 50},
	{Name: "北京丰台站驻点", Province: "北京市", City: "北京市", County: "丰台区", Longitude: 116.3040, Latitude: 39.8500, Altitude: 50},
	{Name: "天津西站驻点", Province: "天津市", City: "天津市", County: "红桥区", Longitude: 117.1560, Latitude: 39.1570, Altitude: 10},
	{Name: "石家庄站驻点", Province: "河北省", City: "石家庄市", County: "桥西区", Longitude: 114.4780, Latitude: 38.0090, Altitude: 80},
	{Name: "张家口站驻点", Province: "河北省", City: "张家口市", County: "桥东区", Longitude: 114.8860, Latitude: 40.7690, Altitude: 720},
	{Name: "济南西站驻点", Province: "山东省", City: "济南市", County: "槐荫区", Longitude: 116.8900, Latitude: 36.6690, Altitude: 50},
	{Name: "青岛北站驻点", Province: "山东省", City: "青岛市", County: "李沧区", Longitude: 120.3750, Latitude: 36.1700, Altitude: 20},
	{Name: "南京南站中转点", Province: "江苏省", City: "南京市", County: "雨花台区", Longitude: 118.7990, Latitude: 31.9690, Altitude: 20},
	{Name: "徐州东站驻点", Province: "江苏省", City: "徐州市", County: "贾汪区", Longitude: 117.3060, Latitude: 34.2760, Altitude: 40},
	{Name: "上海虹桥站中转点", Province: "上海市", City: "上海市", County: "闵行区", Longitude: 121.3200, Latitude: 31.1940, Altitude: 10},
	{Name: "杭州东站驻点", Province: "浙江省", City: "杭州市", County: "上城区", Longitude: 120.2130, Latitude: 30.2910, Altitude: 20},
	{Name: "宁波站驻点", Province: "浙江省", City: "宁波市", County: "海曙区", Longitude: 121.5460, Latitude: 29.8640, Altitude: 10},
	{Name: "合肥南站驻点", Province: "安徽省", City: "合肥市", County: "包河区", Longitude: 117.3170, Latitude: 31.7980, Altitude: 30},
	{Name: "黄山北站驻点", Province: "安徽省", City: "黄山市", County: "屯溪区", Longitude: 118.2860, Latitude: 29.7290, Altitude: 160},
	{Name: "南昌西站驻点", Province: "江西省", City: "南昌市", County: "红谷滩区", Longitude: 115.7940, Latitude: 28.6270, Altitude: 30},
	{Name: "赣州西站驻点", Province: "江西省", City: "赣州市", County: "南康区", Longitude: 114.7650, Latitude: 25.7130, Altitude: 130},
	{Name: "福州站驻点", Province: "福建省", City: "福州市", County: "晋安区", Longitude: 119.3200, Latitude: 26.0850, Altitude: 20},
	{Name: "厦门北站驻点", Province: "福建省", City: "厦门市", County: "集美区", Longitude: 118.0710, Latitude: 24.6400, Altitude: 20},
	{Name: "太原南站驻点", Province: "山西省", City: "太原市", County: "小店区", Longitude: 112.6060, Latitude: 37.7910, Altitude: 800},
	{Name: "大同南站驻点", Province: "山西省", City: "大同市", County: "平城区", Longitude: 113.3660, Latitude: 40.0180, Altitude: 1050},
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

	// Light history only for newly created workers; dense season fills the screen window.
	checkIns := make([]models.DailyHealthStatus, 0, 4000)
	flush := func() error {
		if len(checkIns) == 0 {
			return nil
		}
		const batch = 500
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

	for wi, nw := range newWorkers {
		u := nw.user
		home := nw.station
		wantDays := 10 + (wi % 6)
		added := 0
		for dayOffset := 1; dayOffset < 40 && added < wantDays; dayOffset++ {
			day := today.AddDate(0, 0, -dayOffset)
			rate := 0.88
			if day.Weekday() == time.Saturday || day.Weekday() == time.Sunday {
				rate = 0.72
			}
			if rng.Float64() > rate {
				continue
			}
			st := pickDenseCheckInStation(rng, home)
			lng := st.Longitude + (rng.Float64()-0.5)*0.06
			lat := st.Latitude + (rng.Float64()-0.5)*0.05
			h := true
			disease := ""
			unhealthyP := 0.08
			if st.Altitude > 4500 {
				unhealthyP = 0.14
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
			if len(checkIns) >= 2500 {
				if err := flush(); err != nil {
					return err
				}
				fmt.Printf("  flushed new-worker check-ins mid-run (%d/%d)\n", wi+1, len(newWorkers))
			}
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
