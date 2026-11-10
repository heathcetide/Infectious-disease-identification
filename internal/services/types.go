// Copyright (c) 2026 heathcetide. All rights reserved.

package services

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/models"
	"github.com/heathcetide/infectious-disease/pkg/utils"
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
// Numeric fields use utils.FlexInt/FlexFloat because uni-app text inputs send strings.
type WorkEnvironmentDTO struct {
	AvgActivityFaceToFace       utils.FlexInt   `json:"avgActivityFaceToFace"`
	AvgActivityPhysicalContact  utils.FlexInt   `json:"avgActivityPhysicalContact"`
	AvgActivityRoomPeople       utils.FlexInt   `json:"avgActivityRoomPeople"`
	AvgBreakfastFaceToFace      utils.FlexInt   `json:"avgBreakfastFaceToFace"`
	AvgBreakfastPhysicalContact utils.FlexInt   `json:"avgBreakfastPhysicalContact"`
	AvgBreakfastRoomPeople      utils.FlexInt   `json:"avgBreakfastRoomPeople"`
	AvgDinnerFaceToFace         utils.FlexInt   `json:"avgDinnerFaceToFace"`
	AvgDinnerPhysicalContact    utils.FlexInt   `json:"avgDinnerPhysicalContact"`
	AvgDinnerRoomPeople         utils.FlexInt   `json:"avgDinnerRoomPeople"`
	AvgLunchFaceToFace          utils.FlexInt   `json:"avgLunchFaceToFace"`
	AvgLunchPhysicalContact     utils.FlexInt   `json:"avgLunchPhysicalContact"`
	AvgLunchRoomPeople          utils.FlexInt   `json:"avgLunchRoomPeople"`
	AvgMonthlyPressure          utils.FlexFloat `json:"avgMonthlyPressure"`
	AvgWorkFaceToFace           utils.FlexInt   `json:"avgWorkFaceToFace"`
	AvgWorkPhysicalContact      utils.FlexInt   `json:"avgWorkPhysicalContact"`
	AvgWorkRoomPeople           utils.FlexInt   `json:"avgWorkRoomPeople"`
	CollectiveActivitiesCount   utils.FlexInt   `json:"collectiveActivitiesCount"`
	DormitoryArea               utils.FlexFloat `json:"dormitoryArea"`
	DormitoryBuildingNumber     utils.FlexInt   `json:"dormitoryBuildingNumber"`
	DormitoryOccupants          utils.FlexInt   `json:"dormitoryOccupants"`
	DormitoryRoomNumber         utils.FlexInt   `json:"dormitoryRoomNumber"`
	MaxTemperature              utils.FlexFloat `json:"maxTemperature"`
	MinTemperature              utils.FlexFloat `json:"minTemperature"`
	SubmissionTime              string          `json:"submissionTime"`
	VentilationCondition        string          `json:"ventilationCondition"`
	WindowOpenArea              utils.FlexFloat `json:"windowOpenArea"`
	WorkStationAltitude         utils.FlexInt   `json:"workStationAltitude"`
	WorkStationName             string          `json:"workStationName"`
}

// AllSymptoms matches frontend AllCheckInfo payload keys.
type AllSymptoms struct {
	Generalsymptoms        GeneralSymptomsDTO          `json:"generalsymptoms"`
	Respiratorysymptoms    models.RespiratorySymptoms  `json:"respiratorysymptoms"`
	Digestivesymptoms      DigestiveSymptomsDTO        `json:"digestivesymptoms"`
	Circulatorysymptoms    models.CirculatorySymptoms  `json:"circulatorysymptoms"`
	NeurologicalSymptoms   models.NeurologicalSymptoms `json:"neurologicalSymptoms"`
	LocalSymptoms          models.LocalSymptoms        `json:"localSymptoms"`
	OtherSymptoms          models.OtherSymptoms        `json:"otherSymptoms"`
	RiskFactorsAndExposure RiskFactorsDTO              `json:"riskFactorsAndExposure"`
	// SymptomPhotoIDs: skin/rash photos classified by YOLO/VLM (P4 auxiliary).
	SymptomPhotoIDs []uint `json:"symptomPhotoIds,omitempty"`
}

// GeneralSymptomsDTO accepts stringy numeric inputs from uni-app text fields.
type GeneralSymptomsDTO struct {
	models.GeneralSymptoms
	HighestTemperature utils.FlexFloat `json:"highestTemperature"`
	FeverDuration      utils.FlexInt   `json:"feverDuration"`
}

// ToModel copies DTO fields into models.GeneralSymptoms.
func (d GeneralSymptomsDTO) ToModel() models.GeneralSymptoms {
	m := d.GeneralSymptoms
	m.HighestTemperature = d.HighestTemperature.Float64()
	m.FeverDuration = d.FeverDuration.Int()
	return m
}

// DigestiveSymptomsDTO maps frontend arrheaFrequencyGEThreeTimesPerDay to the model field.
type DigestiveSymptomsDTO struct {
	models.DigestiveSymptoms
	ArrheaFrequencyGEThreeTimesPerDay bool          `json:"arrheaFrequencyGEThreeTimesPerDay"`
	DiarrheaTimesPerDay               utils.FlexInt `json:"diarrheaTimesPerDay"`
}

// ToModel copies DTO fields into models.DigestiveSymptoms.
func (d DigestiveSymptomsDTO) ToModel() models.DigestiveSymptoms {
	m := d.DigestiveSymptoms
	if d.ArrheaFrequencyGEThreeTimesPerDay {
		m.ArrheaFrequencyGEThreeTimesPerDay = true
	}
	m.DiarrheaTimesPerDay = d.DiarrheaTimesPerDay.Int()
	return m
}

// RiskFactorsDTO maps a few frontend keys that omit the "is" / "has" prefix.
type RiskFactorsDTO struct {
	models.RiskFactorsAndExposure
	ContactDate        utils.FlexDate `json:"contactDate"`
	GroupDining        bool           `json:"groupDining"`
	DiningOut          bool           `json:"diningOut"`
	DiningUnclear      bool           `json:"diningUnclear"`
	HillyOrMountainous bool           `json:"hillyOrMountainous"`
	Plain              bool           `json:"plain"`
	OtherResidenceType bool           `json:"otherResidenceType"`
}

// ToModel merges aliased booleans into models.RiskFactorsAndExposure.
func (r RiskFactorsDTO) ToModel() models.RiskFactorsAndExposure {
	m := r.RiskFactorsAndExposure
	m.ContactDate = r.ContactDate.Ptr()
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
	HasSignedToday     bool                      `json:"hasSignedToday"`
	HasMonthSigned     bool                      `json:"hasMonthSigned"` // monthly work-environment filled
	Dailyhealthstatus  *models.DailyHealthStatus `json:"dailyhealthstatus,omitempty"`
	DiseasescoringList []ScoredDiseaseItem       `json:"diseasescoringList,omitempty"`
	Hybrid             *HybridMeta               `json:"hybrid,omitempty"`
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
	DiarrheaFrequencyGEThreeTimesPerDay bool          `json:"diarrheaFrequencyGEThreeTimesPerDay"`
	DiarrheaTimesPerDay                 utils.FlexInt `json:"diarrheaTimesPerDay"`
}

func (d DiagnosisDigestiveDTO) ToModel() models.DiagnosisDigestiveSymptoms {
	m := d.DiagnosisDigestiveSymptoms
	if d.DiarrheaFrequencyGEThreeTimesPerDay {
		m.IsDiarrheaFrequencyGEThreeTimesPerDay = true
	}
	m.DiarrheaTimesPerDay = d.DiarrheaTimesPerDay.Int()
	return m
}

func DiagnosisDigestiveFromModel(m models.DiagnosisDigestiveSymptoms) DiagnosisDigestiveDTO {
	return DiagnosisDigestiveDTO{
		DiagnosisDigestiveSymptoms:          m,
		DiarrheaFrequencyGEThreeTimesPerDay: m.IsDiarrheaFrequencyGEThreeTimesPerDay,
		DiarrheaTimesPerDay:                 utils.FlexInt(m.DiarrheaTimesPerDay),
	}
}

// DiagnosisGeneralDTO accepts stringy numeric inputs from uni-app.
type DiagnosisGeneralDTO struct {
	models.DiagnosisGeneralSymptoms
	HighestTemperature  utils.FlexFloat `json:"highestTemperature"`
	FeverDuration       utils.FlexInt   `json:"feverDuration"`
	LowestBloodPressure utils.FlexInt   `json:"lowestBloodPressure"`
	BleedingVolume      utils.FlexInt   `json:"bleedingVolume"`
}

func (d DiagnosisGeneralDTO) ToModel() models.DiagnosisGeneralSymptoms {
	m := d.DiagnosisGeneralSymptoms
	m.HighestTemperature = d.HighestTemperature.Float64()
	m.FeverDuration = d.FeverDuration.Int()
	m.LowestBloodPressure = d.LowestBloodPressure.Int()
	m.BleedingVolume = d.BleedingVolume.Int()
	return m
}

func DiagnosisGeneralFromModel(m models.DiagnosisGeneralSymptoms) DiagnosisGeneralDTO {
	return DiagnosisGeneralDTO{
		DiagnosisGeneralSymptoms: m,
		HighestTemperature:       utils.FlexFloat(m.HighestTemperature),
		FeverDuration:            utils.FlexInt(m.FeverDuration),
		LowestBloodPressure:      utils.FlexInt(m.LowestBloodPressure),
		BleedingVolume:           utils.FlexInt(m.BleedingVolume),
	}
}

// DiagnosisPersonalDTO accepts stringy vitals from uni-app text inputs.
type DiagnosisPersonalDTO struct {
	models.DiagnosisPersonalInfo
	Height            utils.FlexFloat `json:"height"`
	Weight            utils.FlexFloat `json:"weight"`
	BloodSugarFasting utils.FlexFloat `json:"bloodSugarFasting"`
	BloodSugarRandom  utils.FlexFloat `json:"bloodSugarRandom"`
	CholesterolTotal  utils.FlexFloat `json:"cholesterolTotal"`
	Triglycerides     utils.FlexFloat `json:"triglycerides"`
	CholesterolLipid  utils.FlexFloat `json:"cholesterolLipid"`
	HDL               utils.FlexFloat `json:"hdl"`
	LDL               utils.FlexFloat `json:"ldl"`
	BloodOxygen       utils.FlexFloat `json:"bloodOxygen"`
}

func (d DiagnosisPersonalDTO) ToModel() models.DiagnosisPersonalInfo {
	m := d.DiagnosisPersonalInfo
	m.Height = d.Height.Float64()
	m.Weight = d.Weight.Float64()
	m.BloodSugarFasting = d.BloodSugarFasting.Float64()
	m.BloodSugarRandom = d.BloodSugarRandom.Float64()
	m.CholesterolTotal = d.CholesterolTotal.Float64()
	m.Triglycerides = d.Triglycerides.Float64()
	m.CholesterolLipid = d.CholesterolLipid.Float64()
	m.HDL = d.HDL.Float64()
	m.LDL = d.LDL.Float64()
	m.BloodOxygen = d.BloodOxygen.Float64()
	return m
}

func DiagnosisPersonalFromModel(m models.DiagnosisPersonalInfo) DiagnosisPersonalDTO {
	return DiagnosisPersonalDTO{
		DiagnosisPersonalInfo: m,
		Height:                utils.FlexFloat(m.Height),
		Weight:                utils.FlexFloat(m.Weight),
		BloodSugarFasting:     utils.FlexFloat(m.BloodSugarFasting),
		BloodSugarRandom:      utils.FlexFloat(m.BloodSugarRandom),
		CholesterolTotal:      utils.FlexFloat(m.CholesterolTotal),
		Triglycerides:         utils.FlexFloat(m.Triglycerides),
		CholesterolLipid:      utils.FlexFloat(m.CholesterolLipid),
		HDL:                   utils.FlexFloat(m.HDL),
		LDL:                   utils.FlexFloat(m.LDL),
		BloodOxygen:           utils.FlexFloat(m.BloodOxygen),
	}
}

// AllDiagnosisDTO is the full diagnosis submission payload from the frontend.
type AllDiagnosisDTO struct {
	DiagnosisResultsEmpDTO        DiagnosisResultsEmpDTO               `json:"diagnosisResultsEmpDTO"`
	Diagnosiscirculatorysymptoms  models.DiagnosisCirculatorySymptoms  `json:"diagnosiscirculatorysymptoms"`
	Diagnosiscomplications        models.DiagnosisComplications        `json:"diagnosiscomplications"`
	Diagnosisdigestivesymptoms    DiagnosisDigestiveDTO                `json:"diagnosisdigestivesymptoms"`
	Diagnosisexaminations         models.DiagnosisExaminations         `json:"diagnosisexaminations"`
	Diagnosisgeneralsymptoms      DiagnosisGeneralDTO                  `json:"diagnosisgeneralsymptoms"`
	Diagnosislocalsymptoms        models.DiagnosisLocalSymptoms        `json:"diagnosislocalsymptoms"`
	Diagnosisneurologicalsymptoms models.DiagnosisNeurologicalSymptoms `json:"diagnosisneurologicalsymptoms"`
	Diagnosisothersymptoms        models.DiagnosisOtherSymptoms        `json:"diagnosisothersymptoms"`
	Diagnosispersonalinfo         DiagnosisPersonalDTO                 `json:"diagnosispersonalinfo"`
	Diagnosisrespiratorysymptoms  models.DiagnosisRespiratorySymptoms  `json:"diagnosisrespiratorysymptoms"`
	FileIDList                    []uint                               `json:"fileIdList"`
	// Files is populated on read so the mini-app can restore uploaded reports.
	Files []models.DiagnosisReport `json:"files,omitempty"`
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
	PageNo             int    `json:"pageNo"`
	PageSize           int    `json:"pageSize"`
	Keyword            string `json:"keyword"`
	UserType           string `json:"userType"`
	Name               string `json:"name"`
	PhoneNumber        string `json:"phoneNumber"`
	Department         string `json:"department"`
	SpecificOccupation string `json:"specificOccupation"`
	Text               string `json:"text"` // disease keyword alias used by admin UI
	IsHealth           string `json:"isHealth"`
	CheckInDateBegin   string `json:"checkInDateBegin"`
	CheckInDateEnd     string `json:"checkInDateEnd"`
	WorkStationName    string `json:"workStationName"`
	BeginYearMonth     string `json:"beginYearMonth"`
	EndYearMonth       string `json:"endYearMonth"`
	BeginDate          string `json:"beginDate"`
	EndDate            string `json:"endDate"`
	FeedbackTitle      string `json:"feedbackTitle"`
}

// CheckInListItem is the CDC check-in table row with joined user profile.
type CheckInListItem struct {
	StatusID           uint    `json:"statusId"`
	UserID             uint    `json:"userId"`
	UserType           string  `json:"userType"`
	Name               string  `json:"name"`
	PhoneNumber        string  `json:"phoneNumber"`
	Gender             string  `json:"gender"`
	Age                int     `json:"age"`
	Department         string  `json:"department"`
	SpecificOccupation string  `json:"specificOccupation"`
	IsHealth           *bool   `json:"isHealth"`
	LocationName       string  `json:"locationName"`
	Province           string  `json:"province"`
	City               string  `json:"city"`
	County             string  `json:"county"`
	CheckInDate        []int   `json:"checkInDate"`
	DiseaseTypeName    string  `json:"diseaseTypeName"`
	Longitude          float64 `json:"longitude"`
	Latitude           float64 `json:"latitude"`
}

// CheckInLocationFlatVO is the location drawer payload expected by admin UI.
type CheckInLocationFlatVO struct {
	UserType     string  `json:"userType"`
	Name         string  `json:"name"`
	Age          int     `json:"age"`
	Gender       string  `json:"gender"`
	Ethnicity    string  `json:"ethnicity"`
	Department   string  `json:"department"`
	CheckInDate  []int   `json:"checkInDate"`
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
	Province     string  `json:"province"`
	City         string  `json:"city"`
	County       string  `json:"county"`
	LocationName string  `json:"locationName"`
	Address      string  `json:"address"`
}

// AdminUserListItem is the admin user table row (Java-compatible field names).
// UserID is a string so snowflake IDs survive JSON→JS without precision loss.
type AdminUserListItem struct {
	UserID                 string `json:"userId"`
	UserType               string `json:"userType"`
	Name                   string `json:"name"`
	PhoneNumber            string `json:"phoneNumber"`
	Gender                 string `json:"gender"`
	Age                    int    `json:"age"`
	Ethnicity              string `json:"ethnicity"`
	EducationLevel         string `json:"educationLevel"`
	WorkOnPlateauStartDate []int  `json:"workOnPlateauStartDate"`
	Department             string `json:"department"`
	SpecificOccupation     string `json:"specificOccupation"`
	IsActived              bool   `json:"isActived"`
}

// DiseaseSubtypeItem is nested under disease type page rows.
type DiseaseSubtypeItem struct {
	SubtypeID   uint   `json:"subtypeId"`
	SubtypeName string `json:"subtypeName"`
}

// DiseaseTypePageItem is admin disease basic-data table row.
type DiseaseTypePageItem struct {
	DiseaseTypeID   uint                 `json:"diseaseTypeId"`
	DiseaseTypeName string               `json:"diseaseTypeName"`
	HasSubtype      bool                 `json:"hasSubtype"`
	SubDiseaseList  []DiseaseSubtypeItem `json:"subDiseaseList"`
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

// SatisfactionSurveyBrief is a lightweight open survey for the mini-app picker.
type SatisfactionSurveyBrief struct {
	SurveyID      uint   `json:"surveyID"`
	SurveyTitle   string `json:"surveyTitle"`
	SurveyContent string `json:"surveyContent"`
	IsOpen        bool   `json:"isOpen"`
}

// NeedRatingVO tells the client whether to show the satisfaction survey.
type NeedRatingVO struct {
	NeedRating bool                      `json:"needRating"`
	Surveys    []SatisfactionSurveyBrief `json:"surveys"`
	// legacy single-survey fields (first open survey) for old clients
	SurveyID      uint   `json:"surveyID,omitempty"`
	SurveyTitle   string `json:"surveyTitle,omitempty"`
	SurveyContent string `json:"surveyContent,omitempty"`
}

// ScoredDiseaseItem is one disease row returned to the mini-app (AI or expert).
type ScoredDiseaseItem struct {
	ScoringID       uint       `json:"scoringId,omitempty"`
	StatusID        uint       `json:"statusId,omitempty"`
	SymptomTime     *time.Time `json:"symptomTime,omitempty"`
	DiseaseTypeID   int        `json:"diseaseTypeId,omitempty"`
	DiseaseTypeName string     `json:"diseaseTypeName"`
	Score           float64    `json:"score"`
	Reason          string     `json:"reason,omitempty"`
	Evidence        []string   `json:"evidence,omitempty"`
	ExpertScore     float64    `json:"expertScore,omitempty"`
}

// HybridMeta describes expert vs AI consistency for P1 hybrid reasoning.
type HybridMeta struct {
	IsAI              bool     `json:"isAI"`
	Disagreement      bool     `json:"disagreement"`
	NeedsHumanConfirm bool     `json:"needsHumanConfirm"`
	Confirmed         bool     `json:"confirmed"`
	LLMDegraded       bool     `json:"llmDegraded"`
	ExpertTop         []string `json:"expertTop,omitempty"`
	AITop             []string `json:"aiTop,omitempty"`
	Message           string   `json:"message,omitempty"`
}

// HybridScoringResult is the symptoms-submit payload (backward-compatible list + hybrid).
type HybridScoringResult struct {
	DiseasescoringList []ScoredDiseaseItem `json:"diseasescoringList"`
	Hybrid             HybridMeta          `json:"hybrid"`
	HighRisk           bool                `json:"highRisk,omitempty"`
	RiskAlerts         []models.RiskAlert  `json:"riskAlerts,omitempty"`
	NearestHospital    string              `json:"nearestHospital,omitempty"`
}

// ParseDate parses yyyy-MM-dd in local timezone.
func ParseDate(s string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", s, time.Local)
}
