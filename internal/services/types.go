// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
)

// CheckInDTO matches frontend clockFormParams (health → isHealth on persist).
type CheckInDTO struct {
	City         string  `json:"city"`
	County       string  `json:"county"`
	Health       bool    `json:"health"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	LocationName string  `json:"locationName"`
	Province     string  `json:"province"`
}

// WorkEnvironmentDTO matches frontend clockEnvironment camelCase JSON.
type WorkEnvironmentDTO struct {
	AvgActivityFaceToFace       int     `json:"avgActivityFaceToFace"`
	AvgActivityPhysicalContact  int     `json:"avgActivityPhysicalContact"`
	AvgActivityRoomPeople       int     `json:"avgActivityRoomPeople"`
	AvgBreakfastFaceToFace      int     `json:"avgBreakfastFaceToFace"`
	AvgBreakfastPhysicalContact int     `json:"avgBreakfastPhysicalContact"`
	AvgBreakfastRoomPeople      int     `json:"avgBreakfastRoomPeople"`
	AvgDinnerFaceToFace         int     `json:"avgDinnerFaceToFace"`
	AvgDinnerPhysicalContact    int     `json:"avgDinnerPhysicalContact"`
	AvgDinnerRoomPeople         int     `json:"avgDinnerRoomPeople"`
	AvgLunchFaceToFace          int     `json:"avgLunchFaceToFace"`
	AvgLunchPhysicalContact     int     `json:"avgLunchPhysicalContact"`
	AvgLunchRoomPeople          int     `json:"avgLunchRoomPeople"`
	AvgMonthlyPressure          float64 `json:"avgMonthlyPressure"`
	AvgWorkFaceToFace           int     `json:"avgWorkFaceToFace"`
	AvgWorkPhysicalContact      int     `json:"avgWorkPhysicalContact"`
	AvgWorkRoomPeople           int     `json:"avgWorkRoomPeople"`
	CollectiveActivitiesCount   int     `json:"collectiveActivitiesCount"`
	DormitoryArea               float64 `json:"dormitoryArea"`
	DormitoryBuildingNumber     int     `json:"dormitoryBuildingNumber"`
	DormitoryOccupants          int     `json:"dormitoryOccupants"`
	DormitoryRoomNumber         int     `json:"dormitoryRoomNumber"`
	MaxTemperature              float64 `json:"maxTemperature"`
	MinTemperature              float64 `json:"minTemperature"`
	SubmissionTime              string  `json:"submissionTime"`
	VentilationCondition        string  `json:"ventilationCondition"`
	WindowOpenArea              float64 `json:"windowOpenArea"`
	WorkStationAltitude         int     `json:"workStationAltitude"`
	WorkStationName             string  `json:"workStationName"`
}

// AllSymptoms matches frontend AllCheckInfo payload keys.
type AllSymptoms struct {
	Generalsymptoms       models.GeneralSymptoms       `json:"generalsymptoms"`
	Respiratorysymptoms   models.RespiratorySymptoms   `json:"respiratorysymptoms"`
	Digestivesymptoms     DigestiveSymptomsDTO         `json:"digestivesymptoms"`
	Circulatorysymptoms   models.CirculatorySymptoms   `json:"circulatorysymptoms"`
	NeurologicalSymptoms  models.NeurologicalSymptoms  `json:"neurologicalSymptoms"`
	LocalSymptoms         models.LocalSymptoms         `json:"localSymptoms"`
	OtherSymptoms         models.OtherSymptoms         `json:"otherSymptoms"`
	RiskFactorsAndExposure RiskFactorsDTO              `json:"riskFactorsAndExposure"`
}

// DigestiveSymptomsDTO maps frontend arrheaFrequencyGEThreeTimesPerDay to the model field.
type DigestiveSymptomsDTO struct {
	models.DigestiveSymptoms
	ArrheaFrequencyGEThreeTimesPerDay bool `json:"arrheaFrequencyGEThreeTimesPerDay"`
}

// ToModel copies DTO fields into models.DigestiveSymptoms.
func (d DigestiveSymptomsDTO) ToModel() models.DigestiveSymptoms {
	m := d.DigestiveSymptoms
	if d.ArrheaFrequencyGEThreeTimesPerDay {
		m.ArrheaFrequencyGEThreeTimesPerDay = true
	}
	return m
}

// RiskFactorsDTO maps a few frontend keys that omit the "is" / "has" prefix.
type RiskFactorsDTO struct {
	models.RiskFactorsAndExposure
	GroupDining        bool `json:"groupDining"`
	DiningOut          bool `json:"diningOut"`
	DiningUnclear      bool `json:"diningUnclear"`
	HillyOrMountainous bool `json:"hillyOrMountainous"`
	Plain              bool `json:"plain"`
	OtherResidenceType bool `json:"otherResidenceType"`
}

// ToModel merges aliased booleans into models.RiskFactorsAndExposure.
func (r RiskFactorsDTO) ToModel() models.RiskFactorsAndExposure {
	m := r.RiskFactorsAndExposure
	if r.GroupDining {
		m.IsGroupDining = true
	}
	if r.DiningOut {
		m.IsDiningOut = true
	}
	if r.DiningUnclear {
		m.IsDiningUnclear = true
	}
	if r.HillyOrMountainous {
		m.IsHillyOrMountainous = true
	}
	if r.Plain {
		m.IsPlain = true
	}
	if r.OtherResidenceType {
		m.IsOtherResidenceType = true
	}
	return m
}

// DailyCheckInVO is returned for daily check-in queries.
type DailyCheckInVO struct {
	HasSignedToday      bool                     `json:"hasSignedToday"`
	Dailyhealthstatus   *models.DailyHealthStatus `json:"dailyhealthstatus,omitempty"`
	DiseasescoringList  []models.DiseaseScoring  `json:"diseasescoringList,omitempty"`
}

// MonthlyHealthDay is one calendar day in the month summary.
type MonthlyHealthDay struct {
	Date     string `json:"date"`
	Health   *bool  `json:"health,omitempty"`
	Diagnose bool   `json:"diagnose,omitempty"`
	Examine  bool   `json:"examine,omitempty"`
}

// DiagnosisResultsEmpDTO matches frontend diagnosisResultsEmpDTO.
type DiagnosisResultsEmpDTO struct {
	AdmissionDate              string `json:"admissionDate"`
	AdmissionSymptomsAndSigns  string `json:"admissionSymptomsAndSigns"`
	AnthraxSubtype             string `json:"anthraxSubtype"`
	DeathDate                  string `json:"deathDate,omitempty"`
	DischargeDate              string `json:"dischargeDate"`
	DiscoveryMethod            string `json:"discoveryMethod"`
	DiseaseOutcome             string `json:"diseaseOutcome"`
	DiseaseType                string `json:"diseaseType"`
	HospitalName               string `json:"hospitalName"`
	OtherDiscoveryMethodName   string `json:"otherDiscoveryMethodName,omitempty"`
	OtherDiseaseName           string `json:"otherDiseaseName,omitempty"`
	OtherRegistrationDetails   string `json:"otherRegistrationDetails,omitempty"`
	PlagueSubtype              string `json:"plagueSubtype,omitempty"`
	RegistrationClassification string `json:"registrationClassification"`
}

// DiagnosisDigestiveDTO maps frontend diarrheaFrequencyGEThreeTimesPerDay.
type DiagnosisDigestiveDTO struct {
	models.DiagnosisDigestiveSymptoms
	DiarrheaFrequencyGEThreeTimesPerDay bool `json:"diarrheaFrequencyGEThreeTimesPerDay"`
}

func (d DiagnosisDigestiveDTO) ToModel() models.DiagnosisDigestiveSymptoms {
	m := d.DiagnosisDigestiveSymptoms
	if d.DiarrheaFrequencyGEThreeTimesPerDay {
		m.IsDiarrheaFrequencyGEThreeTimesPerDay = true
	}
	return m
}

// AllDiagnosisDTO is the full diagnosis submission payload from the frontend.
type AllDiagnosisDTO struct {
	DiagnosisResultsEmpDTO       DiagnosisResultsEmpDTO                `json:"diagnosisResultsEmpDTO"`
	Diagnosiscirculatorysymptoms models.DiagnosisCirculatorySymptoms   `json:"diagnosiscirculatorysymptoms"`
	Diagnosiscomplications       models.DiagnosisComplications         `json:"diagnosiscomplications"`
	Diagnosisdigestivesymptoms   DiagnosisDigestiveDTO                 `json:"diagnosisdigestivesymptoms"`
	Diagnosisexaminations        models.DiagnosisExaminations          `json:"diagnosisexaminations"`
	Diagnosisgeneralsymptoms     models.DiagnosisGeneralSymptoms       `json:"diagnosisgeneralsymptoms"`
	Diagnosislocalsymptoms       models.DiagnosisLocalSymptoms         `json:"diagnosislocalsymptoms"`
	Diagnosisneurologicalsymptoms models.DiagnosisNeurologicalSymptoms `json:"diagnosisneurologicalsymptoms"`
	Diagnosisothersymptoms       models.DiagnosisOtherSymptoms         `json:"diagnosisothersymptoms"`
	Diagnosispersonalinfo        models.DiagnosisPersonalInfo          `json:"diagnosispersonalinfo"`
	Diagnosisrespiratorysymptoms models.DiagnosisRespiratorySymptoms   `json:"diagnosisrespiratorysymptoms"`
	FileIDList                   []uint                                `json:"fileIdList"`
}

// LabTestReportDTO matches frontend lab report submission.
type LabTestReportDTO struct {
	PathogenicTestResults       string `json:"pathogenicTestResults"`
	IsVirusAntigenTestDone      bool   `json:"virusAntigenTestDone"`
	IsVirusNucleicAcidTestDone  bool   `json:"virusNucleicAcidTestDone"`
	IsVirusCultureIsolationDone bool   `json:"virusCultureIsolationDone"`
	IsSerologicalTestDone       bool   `json:"serologicalTestDone"`
	LabTestFileIDs              []uint `json:"labTestFileIds,omitempty"`
}

// LearningMaterialsPageDTO matches pageSelect query body.
type LearningMaterialsPageDTO struct {
	DiseaseTypeName string `json:"diseaseTypeName" form:"diseaseTypeName"`
	PageNo          int    `json:"pageNo" form:"pageNo"`
	PageSize        int    `json:"pageSize" form:"pageSize"`
}

// SubmitSatisfactionDTO matches frontend submitParams.
type SubmitSatisfactionDTO struct {
	PhoneNumber       string `json:"phoneNumber"`
	RatingID          uint   `json:"ratingID"`
	RatingTime        string `json:"ratingTime"`
	SatisfactionLevel string `json:"satisfactionLevel"`
	SurveyID          uint   `json:"surveyID"`
	UserID            uint   `json:"userID"`
}

// FeedbackDTO is user feedback submission.
type FeedbackDTO struct {
	Name          string `json:"name"`
	PhoneNumber   string `json:"phoneNumber"`
	FeedbackTitle string `json:"feedbackTitle"`
	FeedbackText  string `json:"feedbackText"`
}

// ChangePasswordDTO matches password change request.
type ChangePasswordDTO struct {
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

// PageQuery is a generic admin list query.
type PageQuery struct {
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
	Keyword  string `json:"keyword"`
}

// Normalize applies defaults for pagination.
func (p *PageQuery) Normalize() (page, size int) {
	page = p.PageNo
	if page < 1 {
		page = 1
	}
	size = p.PageSize
	if size < 1 {
		size = 10
	}
	if size > 100 {
		size = 100
	}
	return page, size
}

// SymptomWeightUpdateDTO updates a single weight row.
type SymptomWeightUpdateDTO struct {
	SymptomWeightingID uint    `json:"symptomWeightingId"`
	WeightScore        float64 `json:"weightScore"`
}

// DashboardCounts aggregates simple admin dashboard metrics.
type DashboardCounts struct {
	UserCount           int64 `json:"userCount"`
	CheckInTodayCount   int64 `json:"checkInTodayCount"`
	DiagnosisTodayCount int64 `json:"diagnosisTodayCount"`
	LabReportTodayCount int64 `json:"labReportTodayCount"`
	FeedbackCount       int64 `json:"feedbackCount"`
}

// NeedRatingVO tells the client whether to show the satisfaction survey.
type NeedRatingVO struct {
	NeedRating bool `json:"needRating"`
}

// ParseDate parses yyyy-MM-dd in local timezone.
func ParseDate(s string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", s, time.Local)
}
