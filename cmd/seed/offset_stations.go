// Copyright (c) 2026 heathcetide. All rights reserved.

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Mid/small-town anchors — avoid provincial capitals / mega hubs.
var obscureAnchors = []railStation{
	{Name: "共和县工区", Province: "青海省", City: "海南藏族自治州", County: "共和县", Longitude: 100.62, Latitude: 36.28, Altitude: 3200},
	{Name: "兴海县养护点", Province: "青海省", City: "海南藏族自治州", County: "兴海县", Longitude: 99.99, Latitude: 35.58, Altitude: 3600},
	{Name: "都兰县工务点", Province: "青海省", City: "海西蒙古族藏族自治州", County: "都兰县", Longitude: 98.09, Latitude: 36.30, Altitude: 3200},
	{Name: "曲麻莱县营地", Province: "青海省", City: "玉树藏族自治州", County: "曲麻莱县", Longitude: 95.80, Latitude: 34.13, Altitude: 4200},
	{Name: "称多县工点", Province: "青海省", City: "玉树藏族自治州", County: "称多县", Longitude: 97.11, Latitude: 33.37, Altitude: 3800},
	{Name: "杂多县桥隧点", Province: "青海省", City: "玉树藏族自治州", County: "杂多县", Longitude: 95.30, Latitude: 32.89, Altitude: 4100},
	{Name: "班玛县驻点", Province: "青海省", City: "果洛藏族自治州", County: "班玛县", Longitude: 100.74, Latitude: 32.93, Altitude: 3600},
	{Name: "门源县工务点", Province: "青海省", City: "海北藏族自治州", County: "门源县", Longitude: 101.62, Latitude: 37.37, Altitude: 2800},
	{Name: "聂荣县驻地", Province: "西藏自治区", City: "那曲市", County: "聂荣县", Longitude: 92.30, Latitude: 32.11, Altitude: 4600},
	{Name: "比如县工区", Province: "西藏自治区", City: "那曲市", County: "比如县", Longitude: 93.68, Latitude: 31.48, Altitude: 4000},
	{Name: "嘉黎县营地", Province: "西藏自治区", City: "那曲市", County: "嘉黎县", Longitude: 93.23, Latitude: 30.64, Altitude: 4500},
	{Name: "尼木县驻点", Province: "西藏自治区", City: "拉萨市", County: "尼木县", Longitude: 90.16, Latitude: 29.43, Altitude: 3800},
	{Name: "墨竹工卡县营地", Province: "西藏自治区", City: "拉萨市", County: "墨竹工卡县", Longitude: 91.73, Latitude: 29.83, Altitude: 3800},
	{Name: "南木林县驻点", Province: "西藏自治区", City: "日喀则市", County: "南木林县", Longitude: 89.10, Latitude: 29.68, Altitude: 4000},
	{Name: "江孜县工区", Province: "西藏自治区", City: "日喀则市", County: "江孜县", Longitude: 89.60, Latitude: 28.92, Altitude: 4000},
	{Name: "工布江达县驻点", Province: "西藏自治区", City: "林芝市", County: "工布江达县", Longitude: 93.25, Latitude: 29.89, Altitude: 3400},
	{Name: "波密县养护点", Province: "西藏自治区", City: "林芝市", County: "波密县", Longitude: 95.77, Latitude: 29.86, Altitude: 2800},
	{Name: "左贡县驻点", Province: "西藏自治区", City: "昌都市", County: "左贡县", Longitude: 97.84, Latitude: 29.67, Altitude: 3800},
	{Name: "洛隆县工区", Province: "西藏自治区", City: "昌都市", County: "洛隆县", Longitude: 95.83, Latitude: 30.74, Altitude: 3700},
	{Name: "若尔盖县工点", Province: "四川省", City: "阿坝藏族羌族自治州", County: "若尔盖县", Longitude: 102.96, Latitude: 33.58, Altitude: 3500},
	{Name: "红原县驻地", Province: "四川省", City: "阿坝藏族羌族自治州", County: "红原县", Longitude: 102.55, Latitude: 32.79, Altitude: 3500},
	{Name: "色达县营地", Province: "四川省", City: "甘孜藏族自治州", County: "色达县", Longitude: 100.33, Latitude: 32.27, Altitude: 3900},
	{Name: "石渠县驻点", Province: "四川省", City: "甘孜藏族自治州", County: "石渠县", Longitude: 98.10, Latitude: 32.98, Altitude: 4200},
	{Name: "理塘县材料场", Province: "四川省", City: "甘孜藏族自治州", County: "理塘县", Longitude: 100.27, Latitude: 30.00, Altitude: 4000},
	{Name: "木里县营地", Province: "四川省", City: "凉山彝族自治州", County: "木里县", Longitude: 101.28, Latitude: 27.93, Altitude: 2500},
	{Name: "德钦县工务点", Province: "云南省", City: "迪庆藏族自治州", County: "德钦县", Longitude: 98.92, Latitude: 28.49, Altitude: 3400},
	{Name: "维西县养护点", Province: "云南省", City: "迪庆藏族自治州", County: "维西县", Longitude: 99.29, Latitude: 27.18, Altitude: 2300},
	{Name: "宁蒗县驻点", Province: "云南省", City: "丽江市", County: "宁蒗县", Longitude: 100.85, Latitude: 27.28, Altitude: 2700},
	{Name: "兰坪县工点", Province: "云南省", City: "怒江傈僳族自治州", County: "兰坪县", Longitude: 99.42, Latitude: 26.45, Altitude: 2300},
	{Name: "肃南县工区", Province: "甘肃省", City: "张掖市", County: "肃南县", Longitude: 99.61, Latitude: 38.84, Altitude: 2300},
	{Name: "阿克塞县营地", Province: "甘肃省", City: "酒泉市", County: "阿克塞县", Longitude: 94.34, Latitude: 39.63, Altitude: 2700},
	{Name: "玛曲县工区", Province: "甘肃省", City: "甘南藏族自治州", County: "玛曲县", Longitude: 102.07, Latitude: 33.97, Altitude: 3500},
	{Name: "碌曲县养护点", Province: "甘肃省", City: "甘南藏族自治州", County: "碌曲县", Longitude: 102.49, Latitude: 34.59, Altitude: 3200},
	{Name: "天祝县材料场", Province: "甘肃省", City: "武威市", County: "天祝县", Longitude: 103.14, Latitude: 36.98, Altitude: 2500},
	{Name: "若羌县营地", Province: "新疆维吾尔自治区", City: "巴音郭楞蒙古自治州", County: "若羌县", Longitude: 88.17, Latitude: 39.02, Altitude: 900},
	{Name: "且末县材料场", Province: "新疆维吾尔自治区", City: "巴音郭楞蒙古自治州", County: "且末县", Longitude: 85.53, Latitude: 38.14, Altitude: 1300},
	{Name: "尼勒克县营地", Province: "新疆维吾尔自治区", City: "伊犁哈萨克自治州", County: "尼勒克县", Longitude: 82.51, Latitude: 43.80, Altitude: 1200},
	{Name: "乌恰县工务点", Province: "新疆维吾尔自治区", City: "克孜勒苏柯尔克孜自治州", County: "乌恰县", Longitude: 75.26, Latitude: 39.72, Altitude: 2200},
	{Name: "莎车县驻点", Province: "新疆维吾尔自治区", City: "喀什地区", County: "莎车县", Longitude: 77.24, Latitude: 38.41, Altitude: 1300},
	{Name: "额济纳旗营地", Province: "内蒙古自治区", City: "阿拉善盟", County: "额济纳旗", Longitude: 101.07, Latitude: 41.96, Altitude: 1000},
	{Name: "四子王旗材料场", Province: "内蒙古自治区", City: "乌兰察布市", County: "四子王旗", Longitude: 111.70, Latitude: 41.53, Altitude: 1400},
	{Name: "武隆区工区", Province: "重庆市", City: "重庆市", County: "武隆区", Longitude: 107.76, Latitude: 29.33, Altitude: 300},
	{Name: "石柱县驻点", Province: "重庆市", City: "重庆市", County: "石柱县", Longitude: 108.11, Latitude: 30.00, Altitude: 500},
	{Name: "威宁县营地", Province: "贵州省", City: "毕节市", County: "威宁县", Longitude: 104.28, Latitude: 26.86, Altitude: 2200},
	{Name: "册亨县养护点", Province: "贵州省", City: "黔西南布依族苗族自治州", County: "册亨县", Longitude: 105.81, Latitude: 24.98, Altitude: 800},
	{Name: "荔波县材料场", Province: "贵州省", City: "黔南布依族苗族自治州", County: "荔波县", Longitude: 107.88, Latitude: 25.41, Altitude: 500},
	{Name: "靖西市工区", Province: "广西壮族自治区", City: "百色市", County: "靖西市", Longitude: 106.41, Latitude: 23.13, Altitude: 700},
	{Name: "天峨县营地", Province: "广西壮族自治区", City: "河池市", County: "天峨县", Longitude: 107.17, Latitude: 25.00, Altitude: 400},
	{Name: "通道县养护点", Province: "湖南省", City: "怀化市", County: "通道县", Longitude: 109.78, Latitude: 26.16, Altitude: 400},
	{Name: "花垣县工点", Province: "湖南省", City: "湘西土家族苗族自治州", County: "花垣县", Longitude: 109.48, Latitude: 28.57, Altitude: 400},
	{Name: "竹溪县驻点", Province: "湖北省", City: "十堰市", County: "竹溪县", Longitude: 109.72, Latitude: 32.32, Altitude: 500},
	{Name: "神农架林区工点", Province: "湖北省", City: "神农架林区", County: "松柏镇", Longitude: 110.68, Latitude: 31.74, Altitude: 1200},
	{Name: "商南县材料场", Province: "陕西省", City: "商洛市", County: "商南县", Longitude: 110.88, Latitude: 33.53, Altitude: 600},
	{Name: "留坝县驻点", Province: "陕西省", City: "汉中市", County: "留坝县", Longitude: 106.92, Latitude: 33.62, Altitude: 1100},
	{Name: "岚皋县工区", Province: "陕西省", City: "安康市", County: "岚皋县", Longitude: 108.90, Latitude: 32.31, Altitude: 500},
	{Name: "吴起县营地", Province: "陕西省", City: "延安市", County: "吴起县", Longitude: 108.18, Latitude: 36.93, Altitude: 1300},
	{Name: "靖边县养护点", Province: "陕西省", City: "榆林市", County: "靖边县", Longitude: 108.80, Latitude: 37.60, Altitude: 1400},
	{Name: "府谷县驻点", Province: "陕西省", City: "榆林市", County: "府谷县", Longitude: 111.07, Latitude: 39.03, Altitude: 900},
}

var siteSuffixes = []string{
	"东侧养护点", "西侧工区", "南沟材料场", "北坡巡检点", "桥头驻点",
	"隧道口营地", "生活区", "临时便道点", "边坡监测点", "砂石料场",
	"配电房外工点", "涵洞养护点", "取水点驻地", "防风障巡检点", "应急物资点",
}

// offsetStations is built once: ≥300 unique mid/small-town offset points.
var offsetStations []railStation

func init() {
	rng := rand.New(rand.NewSource(20261108 + 300))
	offsetStations = make([]railStation, 0, 360)
	seen := map[string]struct{}{}
	add := func(st railStation) {
		key := fmt.Sprintf("%.4f,%.4f,%s", st.Longitude, st.Latitude, st.Name)
		if _, ok := seen[key]; ok {
			return
		}
		seen[key] = struct{}{}
		offsetStations = append(offsetStations, st)
	}
	for _, a := range obscureAnchors {
		add(a)
	}
	for len(offsetStations) < 320 {
		base := obscureAnchors[rng.Intn(len(obscureAnchors))]
		ang := rng.Float64() * 2 * math.Pi
		dist := 0.15 + rng.Float64()*0.65
		st := base
		st.Longitude = round6(base.Longitude + math.Cos(ang)*dist)
		st.Latitude = round6(base.Latitude + math.Sin(ang)*dist*0.85)
		st.Altitude = base.Altitude + rng.Intn(400) - 120
		if st.Altitude < 50 {
			st.Altitude = 50
		}
		label := base.County
		if label == "" {
			label = base.Name
		}
		st.Name = label + "·" + siteSuffixes[rng.Intn(len(siteSuffixes))] + fmt.Sprintf("%02d", len(offsetStations)%97+1)
		add(st)
	}
}

func pickOffsetStation(rng *rand.Rand) railStation {
	if len(offsetStations) == 0 {
		return obscureAnchors[0]
	}
	st := offsetStations[rng.Intn(len(offsetStations))]
	st.Longitude = round6(st.Longitude + (rng.Float64()-0.5)*0.04)
	st.Latitude = round6(st.Latitude + (rng.Float64()-0.5)*0.03)
	return st
}

// pickDenseCheckInStation prefers obscure offset points so the heatmap is not capital-heavy.
func pickDenseCheckInStation(rng *rand.Rand, home railStation) railStation {
	roll := rng.Float64()
	switch {
	case roll < 0.70:
		return pickOffsetStation(rng)
	case roll < 0.82:
		st := home
		st.Longitude = round6(home.Longitude + (rng.Float64()-0.5)*0.35)
		st.Latitude = round6(home.Latitude + (rng.Float64()-0.5)*0.28)
		st.Name = home.Name + "·偏移工点"
		return st
	case roll < 0.90:
		return qingzangStations[rng.Intn(len(qingzangStations))]
	case roll < 0.96:
		return nationwideStations[rng.Intn(len(nationwideStations))]
	default:
		return pickOffsetStation(rng)
	}
}

// seedRelocateCheckIns rewrites dense-season coordinates onto the offset catalog.
// Fast path: assign each status_id to a station bucket, then one UPDATE per bucket (SQLite-friendly).
func seedRelocateCheckIns(db *gorm.DB) error {
	rng := rand.New(rand.NewSource(20261108 + 777))
	silent := db.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)})
	start := time.Date(2026, 11, 8, 0, 0, 0, 0, time.Local)
	end := time.Date(2027, 8, 1, 0, 0, 0, 0, time.Local)

	type idRow struct{ StatusID uint }
	var ids []idRow
	if err := silent.Model(&models.DailyHealthStatus{}).
		Select("status_id").
		Where("check_in_date >= ? AND check_in_date < ?", start, end).
		Order("status_id ASC").
		Find(&ids).Error; err != nil {
		return err
	}
	fmt.Printf("relocate scan: candidates=%d catalog=%d\n", len(ids), len(offsetStations))
	if len(ids) == 0 || len(offsetStations) == 0 {
		return nil
	}

	// ~75% relocate: bucket status IDs by station index for bulk UPDATE.
	buckets := make([][]uint, len(offsetStations))
	skipped := 0
	for _, row := range ids {
		if rng.Float64() < 0.25 {
			skipped++
			continue
		}
		si := rng.Intn(len(offsetStations))
		buckets[si] = append(buckets[si], row.StatusID)
	}

	updated := 0
	const chunk = 800
	for si, list := range buckets {
		if len(list) == 0 {
			continue
		}
		st := offsetStations[si]
		// tiny per-bucket jitter so points aren't pixel-identical
		lng := round6(st.Longitude + (rng.Float64()-0.5)*0.05)
		lat := round6(st.Latitude + (rng.Float64()-0.5)*0.04)
		for i := 0; i < len(list); i += chunk {
			j := i + chunk
			if j > len(list) {
				j = len(list)
			}
			part := list[i:j]
			res := silent.Model(&models.DailyHealthStatus{}).
				Where("status_id IN ?", part).
				Updates(map[string]any{
					"longitude":     lng,
					"latitude":      lat,
					"province":      st.Province,
					"city":          st.City,
					"county":        st.County,
					"location_name": st.Name,
				})
			if res.Error != nil {
				return res.Error
			}
			updated += int(res.RowsAffected)
		}
		if si%40 == 0 {
			fmt.Printf("relocate bucket %d/%d updated=%d skipped=%d\n", si+1, len(buckets), updated, skipped)
		}
	}
	fmt.Printf("relocate done: offset_catalog=%d updated=%d skipped=%d\n", len(offsetStations), updated, skipped)
	return nil
}
