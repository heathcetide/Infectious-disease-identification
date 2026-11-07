// Copyright (c) 2026 heathcetide. All rights reserved.

package models

// ScoringDiseaseNames are the twelve fixed disease names used in health scoring.
// DailyhealthstatusServiceImpl.java uses variant labels (e.g. mosquito/tick suffixes);
// this slice matches the canonical frontend scoring list.
var ScoringDiseaseNames = []string{
	"感染性腹泻",
	"流行性出血热",
	"炭疽",
	"鼠疫",
	"霍乱",
	"细菌性痢疾",
	"伤寒和副伤寒",
	"疟疾",
	"登革热",
	"布鲁氏菌病",
	"新型冠状病毒感染",
	"流感",
}

// JavaScoringDiseaseNames mirrors the disease labels in DailyhealthstatusServiceImpl.java.
var JavaScoringDiseaseNames = []string{
	"感染性腹泻",
	"流行性出血热",
	"炭疽",
	"发热伴血小板减少综合征（蜱媒传染病）",
	"登革热（蚊媒传染病）",
	"疟疾（蚊媒传染病）",
	"森林脑炎（蜱媒传染病）",
	"新型冠状病毒感染",
	"鼠疫",
	"斑疹伤寒",
	"结核病",
	"流感",
}
