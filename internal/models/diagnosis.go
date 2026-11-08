// Copyright (c) 2026 heathcetide. All rights reserved.

package models

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/constants"
)

// DiagnosisResults is a confirmed diagnosis submission for a user.
type DiagnosisResults struct {
	DiagnosisResultsID uint `json:"diagnosisResultsId" gorm:"primaryKey;autoIncrement"`
	UserID             uint `json:"userId" gorm:"index;not null"`

	SubmissionTime      *time.Time `json:"submissionTime,omitempty" gorm:"type:date"`
	SubmissionUserID    uint       `json:"submissionUserId,omitempty"`
	SubmissionUserName  string     `json:"submissionUserName,omitempty" gorm:"size:100"`
	SubmissionUserType  string     `json:"submissionUserType,omitempty" gorm:"size:50"`
	DiseaseType         string     `json:"diseaseType,omitempty" gorm:"size:100"`
	OtherDiseaseName    string     `json:"otherDiseaseName,omitempty" gorm:"size:100"`
	PlagueSubtype       string     `json:"plagueSubtype,omitempty" gorm:"size:50"`
	AnthraxSubtype      string     `json:"anthraxSubtype,omitempty" gorm:"size:50"`
	DiscoveryMethod     string     `json:"discoveryMethod,omitempty" gorm:"size:100"`
	OtherDiscoveryMethodName string `json:"otherDiscoveryMethodName,omitempty" gorm:"size:100"`
	DiseaseOutcome      string     `json:"diseaseOutcome,omitempty" gorm:"size:50"`
	HospitalName        string     `json:"hospitalName,omitempty" gorm:"size:200"`
	AdmissionDate       *time.Time `json:"admissionDate,omitempty" gorm:"type:date"`
	DischargeDate       *time.Time `json:"dischargeDate,omitempty" gorm:"type:date"`
	DeathDate           *time.Time `json:"deathDate,omitempty" gorm:"type:date"`
	AdmissionSymptomsAndSigns string `json:"admissionSymptomsAndSigns,omitempty" gorm:"type:text"`
	RegistrationClassification string `json:"registrationClassification,omitempty" gorm:"size:50"`
	OtherRegistrationDetails   string `json:"otherRegistrationDetails,omitempty" gorm:"size:200"`
}

func (DiagnosisResults) TableName() string { return constants.TableDiagnosisResults }

// DiagnosisReport is a file attachment for a diagnosis result.
type DiagnosisReport struct {
	ReportID           uint `json:"reportId" gorm:"primaryKey;autoIncrement"`
	DiagnosisResultsID uint `json:"diagnosisResultsId" gorm:"index;not null"`

	ReportType string    `json:"reportType,omitempty" gorm:"size:50"`
	FileName   string    `json:"fileName,omitempty" gorm:"size:255"`
	FilePath   string    `json:"filePath,omitempty" gorm:"size:500"`
	FileType   string    `json:"fileType,omitempty" gorm:"size:50"`
	UploadDate time.Time `json:"uploadDate"`

	// P3 report understanding: OCR → structured metrics.
	OCRStatus   string `json:"ocrStatus,omitempty" gorm:"size:20"`  // pending|done|failed|skipped
	OCRText     string `json:"ocrText,omitempty" gorm:"type:text"`
	MetricsJSON string `json:"metricsJson,omitempty" gorm:"type:text"`
	OCRError    string `json:"ocrError,omitempty" gorm:"size:500"`
}

func (DiagnosisReport) TableName() string { return constants.TableDiagnosisReports }

// DiagnosisPersonalInfo holds vitals and lab values for a diagnosis.
type DiagnosisPersonalInfo struct {
	DiagnosisPersonalInfoID uint `json:"diagnosisPersonalInfoId" gorm:"primaryKey;autoIncrement"`
	DiagnosisResultsID      uint `json:"diagnosisResultsId" gorm:"index;not null"`

	Height             float64 `json:"height"`
	Weight             float64 `json:"weight"`
	BloodPressure      string  `json:"bloodPressure,omitempty" gorm:"size:20"`
	BloodSugarFasting  float64 `json:"bloodSugarFasting,omitempty"`
	BloodSugarRandom   float64 `json:"bloodSugarRandom,omitempty"`
	CholesterolTotal   float64 `json:"cholesterolTotal,omitempty"`
	Triglycerides      float64 `json:"triglycerides,omitempty"`
	CholesterolLipid   float64 `json:"cholesterolLipid,omitempty"`
	HDL                float64 `json:"hdl,omitempty"`
	LDL                float64 `json:"ldl,omitempty"`
	BloodOxygen        float64 `json:"bloodOxygen,omitempty"`
}

func (DiagnosisPersonalInfo) TableName() string { return constants.TableDiagnosisPersonalInfo }

// DiagnosisComplications records complication flags for a diagnosis.
type DiagnosisComplications struct {
	DiagnosisComplicationID uint `json:"diagnosisComplicationId" gorm:"primaryKey;autoIncrement"`
	DiagnosisResultsID      uint `json:"diagnosisResultsId" gorm:"index;not null"`

	HasComplications             bool   `json:"hasComplications" gorm:"default:false"`
	HasViralPneumonia            bool   `json:"hasViralPneumonia" gorm:"default:false"`
	HasBacterialPneumonia        bool   `json:"hasBacterialPneumonia" gorm:"default:false"`
	HasFungalPneumonia           bool   `json:"hasFungalPneumonia" gorm:"default:false"`
	HasRespiratoryFailure        bool   `json:"hasRespiratoryFailure" gorm:"default:false"`
	HasArds                      bool   `json:"hasArds" gorm:"default:false"`
	HasDic                       bool   `json:"hasDic" gorm:"default:false"`
	HasToxicShock                bool   `json:"hasToxicShock" gorm:"default:false"`
	HasMeningitis                bool   `json:"hasMeningitis" gorm:"default:false"`
	HasEncephalitis              bool   `json:"hasEncephalitis" gorm:"default:false"`
	HasBacteremia                bool   `json:"hasBacteremia" gorm:"default:false"`
	HasMyocarditis               bool   `json:"hasMyocarditis" gorm:"default:false"`
	HasAcuteLungInjury           bool   `json:"hasAcuteLungInjury" gorm:"default:false"`
	HasAcuteKidneyInjury         bool   `json:"hasAcuteKidneyInjury" gorm:"default:false"`
	HasEpilepsy                  bool   `json:"hasEpilepsy" gorm:"default:false"`
	HasToxicHepatitis            bool   `json:"hasToxicHepatitis" gorm:"default:false"`
	HasFluidOverload             bool   `json:"hasFluidOverload" gorm:"default:false"`
	HasSuperinfection            bool   `json:"hasSuperinfection" gorm:"default:false"`
	HasHypoxemia                 bool   `json:"hasHypoxemia" gorm:"default:false"`
	HasSepticShock               bool   `json:"hasSepticShock" gorm:"default:false"`
	HasMetabolicAcidosis         bool   `json:"hasMetabolicAcidosis" gorm:"default:false"`
	HasElectrolyteImbalance      bool   `json:"hasElectrolyteImbalance" gorm:"default:false"`
	HasHemolysis                 bool   `json:"hasHemolysis" gorm:"default:false"`
	HasGastrointestinalBleeding  bool   `json:"hasGastrointestinalBleeding" gorm:"default:false"`
	HasCoagulationDisorder       bool   `json:"hasCoagulationDisorder" gorm:"default:false"`
	HasMultipleOrganFailure      bool   `json:"hasMultipleOrganFailure" gorm:"default:false"`
	HasBronchitis                bool   `json:"hasBronchitis" gorm:"default:false"`
	HasBronchopneumonia          bool   `json:"hasBronchopneumonia" gorm:"default:false"`
	HasAnemia                    bool   `json:"hasAnemia" gorm:"default:false"`
	HasHepatosplenomegaly        bool   `json:"hasHepatosplenomegaly" gorm:"default:false"`
	HasBrainEdema                bool   `json:"hasBrainEdema" gorm:"default:false"`
	HasHus                       bool   `json:"hasHus" gorm:"default:false"`
	HasHematemesis               bool   `json:"hasHematemesis" gorm:"default:false"`
	HasMelena                    bool   `json:"hasMelena" gorm:"default:false"`
	HasHemoptysis                bool   `json:"hasHemoptysis" gorm:"default:false"`
	HasAbdominalHemorrhage       bool   `json:"hasAbdominalHemorrhage" gorm:"default:false"`
	HasNasalBleeding             bool   `json:"hasNasalBleeding" gorm:"default:false"`
	HasVaginalBleeding           bool   `json:"hasVaginalBleeding" gorm:"default:false"`
	HasHypertensiveEncephalopathy bool  `json:"hasHypertensiveEncephalopathy" gorm:"default:false"`
	HasIntracranialHemorrhage    bool   `json:"hasIntracranialHemorrhage" gorm:"default:false"`
	HasPulmonaryEdema            bool   `json:"hasPulmonaryEdema" gorm:"default:false"`
	HasSecondaryInfection        bool   `json:"hasSecondaryInfection" gorm:"default:false"`
	HasSpontaneousRenalRupture   bool   `json:"hasSpontaneousRenalRupture" gorm:"default:false"`
	HasMyocardialDamage          bool   `json:"hasMyocardialDamage" gorm:"default:false"`
	HasLiverDamage               bool   `json:"hasLiverDamage" gorm:"default:false"`
	OtherComplications           bool   `json:"otherComplications" gorm:"default:false"`
	OtherComplicationsName       string `json:"otherComplicationsName,omitempty" gorm:"size:500"`
}

func (DiagnosisComplications) TableName() string { return constants.TableDiagnosisComplications }

// DiagnosisExaminations records which examinations were performed for a diagnosis.
type DiagnosisExaminations struct {
	ExaminationID      uint `json:"examinationId" gorm:"primaryKey;autoIncrement"`
	DiagnosisResultsID uint `json:"diagnosisResultsId" gorm:"index;not null"`

	HasBloodTest   bool `json:"hasBloodTest" gorm:"default:false"`
	HasUrineTest   bool `json:"hasUrineTest" gorm:"default:false"`
	HasStoolTest   bool `json:"hasStoolTest" gorm:"default:false"`
	HasImaging     bool `json:"hasImaging" gorm:"default:false"`
	HasUltrasound  bool `json:"hasUltrasound" gorm:"default:false"`
}

func (DiagnosisExaminations) TableName() string { return constants.TableDiagnosisExaminations }

// DiagnosisGeneralSymptoms records general symptoms for a diagnosis case.
type DiagnosisGeneralSymptoms struct {
	DiagnosisGeneralSymptomID uint `json:"diagnosisGeneralSymptomId" gorm:"primaryKey;autoIncrement"`
	DiagnosisResultsID        uint `json:"diagnosisResultsId" gorm:"index;not null"`

	HasFever                             bool    `json:"hasFever" gorm:"default:false"`
	HighestTemperature                   float64 `json:"highestTemperature,omitempty"`
	FeverDuration                        int     `json:"feverDuration,omitempty"`
	HasChills                            bool    `json:"hasChills" gorm:"default:false"`
	HasSweating                          bool    `json:"hasSweating" gorm:"default:false"`
	HasFatigue                           bool    `json:"hasFatigue" gorm:"default:false"`
	HasHeadache                          bool    `json:"hasHeadache" gorm:"default:false"`
	HasMusclePain                        bool    `json:"hasMusclePain" gorm:"default:false"`
	HasJointPain                         bool    `json:"hasJointPain" gorm:"default:false"`
	HasLymphNodeSwelling                 bool    `json:"hasLymphNodeSwelling" gorm:"default:false"`
	HasLymphNodeSwellingGroin            bool    `json:"hasLymphNodeSwellingGroin" gorm:"default:false"`
	HasLymphNodeSwellingArmpit           bool    `json:"hasLymphNodeSwellingArmpit" gorm:"default:false"`
	HasLymphNodeSwellingSubclavian       bool    `json:"hasLymphNodeSwellingSubclavian" gorm:"default:false"`
	HasLymphNodeSwellingNeck             bool    `json:"hasLymphNodeSwellingNeck" gorm:"default:false"`
	HasLymphNodeSwellingElbow            bool    `json:"hasLymphNodeSwellingElbow" gorm:"default:false"`
	HasLymphNodeSwellingPopliteal        bool    `json:"hasLymphNodeSwellingPopliteal" gorm:"default:false"`
	HasCyanosis                          bool    `json:"hasCyanosis" gorm:"default:false"`
	HasSubcutaneousAndMucosalBleedingSpots bool  `json:"hasSubcutaneousAndMucosalBleedingSpots" gorm:"default:false"`
	HasSevereAnemia                      bool    `json:"hasSevereAnemia" gorm:"default:false"`
	HasPainfulRedRash                    bool    `json:"hasPainfulRedRash" gorm:"default:false"`
	HasBloodBlisters                     bool    `json:"hasBloodBlisters" gorm:"default:false"`
	HasSkinUlcer                         bool    `json:"hasSkinUlcer" gorm:"default:false"`
	HasSubcutaneousAndMucosalBleedingCavity bool `json:"hasSubcutaneousAndMucosalBleedingCavity" gorm:"default:false"`
	HasCongestiveOrPetechialRash         bool    `json:"hasCongestiveOrPetechialRash" gorm:"default:false"`
	HasPressureInsensitiveRash           bool    `json:"hasPressureInsensitiveRash" gorm:"default:false"`
	HasDehydration                       bool    `json:"hasDehydration" gorm:"default:false"`
	DehydrationSeverity                  string  `json:"dehydrationSeverity,omitempty" gorm:"size:20"`
	HasItchyRash                         bool    `json:"hasItchyRash" gorm:"default:false"`
	ItchyRashOnFingers                   bool    `json:"itchyRashOnFingers" gorm:"default:false"`
	ItchyRashOnBackOfHands               bool    `json:"itchyRashOnBackOfHands" gorm:"default:false"`
	ItchyRashOnUpperLimbs                bool    `json:"itchyRashOnUpperLimbs" gorm:"default:false"`
	ItchyRashOnLowerLimbs                bool    `json:"itchyRashOnLowerLimbs" gorm:"default:false"`
	ItchyRashOnFeet                       bool    `json:"itchyRashOnFeet" gorm:"default:false"`
	ItchyRashOnFace                        bool    `json:"itchyRashOnFace" gorm:"default:false"`
	ItchyRashOnOther                       bool    `json:"itchyRashOnOther" gorm:"default:false"`
	HasEdema                               bool    `json:"hasEdema" gorm:"default:false"`
	HasNightSweats                         bool    `json:"hasNightSweats" gorm:"default:false"`
	HasWeightLoss                          bool    `json:"hasWeightLoss" gorm:"default:false"`
	HasExhaustion                          bool    `json:"hasExhaustion" gorm:"default:false"`
	HasWhiteBloodCells                     bool    `json:"hasWhiteBloodCells" gorm:"default:false"`
	HasBloodPressureDrop                   bool    `json:"hasBloodPressureDrop" gorm:"default:false"`
	LowestBloodPressure                    int     `json:"lowestBloodPressure,omitempty"`
	HasShockSyndrome                       bool    `json:"hasShockSyndrome" gorm:"default:false"`
	HasHighIntracranialPressure            bool    `json:"hasHighIntracranialPressure" gorm:"default:false"`
	HasTurbidCerebrospinalFluid            bool    `json:"hasTurbidCerebrospinalFluid" gorm:"default:false"`
	HasBleeding                            bool    `json:"hasBleeding" gorm:"default:false"`
	BleedingVolume                         int     `json:"bleedingVolume,omitempty"`
	BleedingCavity                         string  `json:"bleedingCavity,omitempty" gorm:"size:100"`
	HasPlasmaLeakage                       bool    `json:"hasPlasmaLeakage" gorm:"default:false"`
	HasOtherOrganDamage                    bool    `json:"hasOtherOrganDamage" gorm:"default:false"`
	HasKernigSign                          bool    `json:"hasKernigSign" gorm:"default:false"`
	HasBrudzinskiSign                      bool    `json:"hasBrudzinskiSign" gorm:"default:false"`
}

func (DiagnosisGeneralSymptoms) TableName() string { return constants.TableDiagnosisGeneralSymptoms }

// DiagnosisRespiratorySymptoms records respiratory symptoms for a diagnosis case.
type DiagnosisRespiratorySymptoms struct {
	DiagnosisRespiratorySymptomID uint `json:"diagnosisRespiratorySymptomId" gorm:"primaryKey;autoIncrement"`
	DiagnosisResultsID            uint `json:"diagnosisResultsId" gorm:"index;not null"`

	HasCough               bool `json:"hasCough" gorm:"default:false"`
	HasSputum              bool `json:"hasSputum" gorm:"default:false"`
	HasBloodySputum        bool `json:"hasBloodySputum" gorm:"default:false"`
	HasHemoptysis          bool `json:"hasHemoptysis" gorm:"default:false"`
	HasSoreThroat          bool `json:"hasSoreThroat" gorm:"default:false"`
	HasDryThroat           bool `json:"hasDryThroat" gorm:"default:false"`
	HasNasalCongestion     bool `json:"hasNasalCongestion" gorm:"default:false"`
	HasRunnyNose           bool `json:"hasRunnyNose" gorm:"default:false"`
	HasPharyngealHyperemia bool `json:"hasPharyngealHyperemia" gorm:"default:false"`
	HasChestTightness      bool `json:"hasChestTightness" gorm:"default:false"`
	HasShortnessOfBreath   bool `json:"hasShortnessOfBreath" gorm:"default:false"`
	HasBreathingDifficulty bool `json:"hasBreathingDifficulty" gorm:"default:false"`
	HasHoarseness          bool `json:"hasHoarseness" gorm:"default:false"`
	HasBronchitis          bool `json:"hasBronchitis" gorm:"default:false"`
	HasLungRales           bool `json:"hasLungRales" gorm:"default:false"`
}

func (DiagnosisRespiratorySymptoms) TableName() string { return constants.TableDiagnosisRespiratorySymptoms }

// DiagnosisDigestiveSymptoms records digestive symptoms for a diagnosis case.
type DiagnosisDigestiveSymptoms struct {
	DiagnosisDigestiveSymptomID uint `json:"diagnosisDigestiveSymptomId" gorm:"primaryKey;autoIncrement"`
	DiagnosisResultsID          uint `json:"diagnosisResultsId" gorm:"index;not null"`

	IsDiarrheaFrequencyGEThreeTimesPerDay bool   `json:"isDiarrheaFrequencyGEThreeTimesPerDay" gorm:"default:false"`
	DiarrheaTimesPerDay                   int    `json:"diarrheaTimesPerDay,omitempty"`
	HasStoolType1                         bool   `json:"hasStoolType1" gorm:"default:false"`
	StoolType1Detail                      string `json:"stoolType1Detail,omitempty" gorm:"size:100"`
	HasStoolType2                         bool   `json:"hasStoolType2" gorm:"default:false"`
	StoolType2Detail                      string `json:"stoolType2Detail,omitempty" gorm:"size:100"`
	DiarrheaMode                          string `json:"diarrheaMode,omitempty" gorm:"size:50"`
	StoolAmount                           string `json:"stoolAmount,omitempty" gorm:"size:20"`
	StoolOdor                             string `json:"stoolOdor,omitempty" gorm:"size:50"`
	HasVomiting                           bool   `json:"hasVomiting" gorm:"default:false"`
	VomitingMode                          string `json:"vomitingMode,omitempty" gorm:"size:50"`
	HasNausea                             bool   `json:"hasNausea" gorm:"default:false"`
	HasAppetiteLoss                       bool   `json:"hasAppetiteLoss" gorm:"default:false"`
	HasAbdominalDistension                bool   `json:"hasAbdominalDistension" gorm:"default:false"`
	HasAbdominalPain                      bool   `json:"hasAbdominalPain" gorm:"default:false"`
	HasBorborygmus                        bool   `json:"hasBorborygmus" gorm:"default:false"`
	HasUpperAbdominalDiscomfort           bool   `json:"hasUpperAbdominalDiscomfort" gorm:"default:false"`
	HasConstipation                       bool   `json:"hasConstipation" gorm:"default:false"`
	HasOliguriaOrAnuria                   bool   `json:"hasOliguriaOrAnuria" gorm:"default:false"`
}

func (DiagnosisDigestiveSymptoms) TableName() string { return constants.TableDiagnosisDigestiveSymptoms }

// DiagnosisCirculatorySymptoms records circulatory symptoms for a diagnosis case.
type DiagnosisCirculatorySymptoms struct {
	DiagnosisCirculatorySymptomID uint `json:"diagnosisCirculatorySymptomId" gorm:"primaryKey;autoIncrement"`
	DiagnosisResultsID            uint `json:"diagnosisResultsId" gorm:"index;not null"`

	HasArrhythmia  bool `json:"hasArrhythmia" gorm:"default:false"`
	HasChestPain   bool `json:"hasChestPain" gorm:"default:false"`
	HasRapidPulse  bool `json:"hasRapidPulse" gorm:"default:false"`
	HasPalpitation bool `json:"hasPalpitation" gorm:"default:false"`
}

func (DiagnosisCirculatorySymptoms) TableName() string { return constants.TableDiagnosisCirculatorySymptoms }

// DiagnosisNeurologicalSymptoms records neurological symptoms for a diagnosis case.
type DiagnosisNeurologicalSymptoms struct {
	DiagnosisNeurologicalSymptomID uint `json:"diagnosisNeurologicalSymptomId" gorm:"primaryKey;autoIncrement"`
	DiagnosisResultsID             uint `json:"diagnosisResultsId" gorm:"index;not null"`

	HasDelirium                    bool `json:"hasDelirium" gorm:"default:false"`
	HasConvulsion                  bool `json:"hasConvulsion" gorm:"default:false"`
	HasDrowsiness                  bool `json:"hasDrowsiness" gorm:"default:false"`
	HasComa                        bool `json:"hasComa" gorm:"default:false"`
	HasStiffNeck                   bool `json:"hasStiffNeck" gorm:"default:false"`
	HasRestlessness                bool `json:"hasRestlessness" gorm:"default:false"`
	HasMuscleParalysisNeckShoulder bool `json:"hasMuscleParalysisNeckShoulder" gorm:"default:false"`
	HasMuscleParalysisLimbs        bool `json:"hasMuscleParalysisLimbs" gorm:"default:false"`
	HasSwallowingDifficulty        bool `json:"hasSwallowingDifficulty" gorm:"default:false"`
	HasSpeechDisorder              bool `json:"hasSpeechDisorder" gorm:"default:false"`
	HasConsciousnessDisorder       bool `json:"hasConsciousnessDisorder" gorm:"default:false"`
	HasDizziness                   bool `json:"hasDizziness" gorm:"default:false"`
	HasTinnitus                    bool `json:"hasTinnitus" gorm:"default:false"`
	HasHearingLoss                 bool `json:"hasHearingLoss" gorm:"default:false"`
}

func (DiagnosisNeurologicalSymptoms) TableName() string { return constants.TableDiagnosisNeurologicalSymptoms }

// DiagnosisLocalSymptoms records localized symptoms for a diagnosis case.
type DiagnosisLocalSymptoms struct {
	DiagnosisLocalSymptomID uint `json:"diagnosisLocalSymptomId" gorm:"primaryKey;autoIncrement"`
	DiagnosisResultsID      uint `json:"diagnosisResultsId" gorm:"index;not null"`

	HasConjunctivitis               bool `json:"hasConjunctivitis" gorm:"default:false"`
	HasConjunctivalCongestion       bool `json:"hasConjunctivalCongestion" gorm:"default:false"`
	HasEyelidEdema                  bool `json:"hasEyelidEdema" gorm:"default:false"`
	HasSmellTasteLoss               bool `json:"hasSmellTasteLoss" gorm:"default:false"`
	HasFacialRednessOrPallor        bool `json:"hasFacialRednessOrPallor" gorm:"default:false"`
	HasRednessNeckChest             bool `json:"hasRednessNeckChest" gorm:"default:false"`
	HasLipCyanosis                  bool `json:"hasLipCyanosis" gorm:"default:false"`
	HasSkinPallorOrBruising         bool `json:"hasSkinPallorOrBruising" gorm:"default:false"`
	HasGumBleeding                  bool `json:"hasGumBleeding" gorm:"default:false"`
	HasKidneyPain                   bool `json:"hasKidneyPain" gorm:"default:false"`
	HasLowBackPain                  bool `json:"hasLowBackPain" gorm:"default:false"`
	HasOrbitalPain                  bool `json:"hasOrbitalPain" gorm:"default:false"`
	HasHepatomegaly                 bool `json:"hasHepatomegaly" gorm:"default:false"`
	HasSplenomegaly                 bool `json:"hasSplenomegaly" gorm:"default:false"`
	HasMucosalBleedingInOralOrNasal bool `json:"hasMucosalBleedingInOralOrNasal" gorm:"default:false"`
	HasSkinBleedingPoints           bool `json:"hasSkinBleedingPoints" gorm:"default:false"`
}

func (DiagnosisLocalSymptoms) TableName() string { return constants.TableDiagnosisLocalSymptoms }

// DiagnosisOtherSymptoms records other symptoms for a diagnosis case.
type DiagnosisOtherSymptoms struct {
	DiagnosisOtherSymptomID uint `json:"diagnosisOtherSymptomId" gorm:"primaryKey;autoIncrement"`
	DiagnosisResultsID      uint `json:"diagnosisResultsId" gorm:"index;not null"`

	HasSuddenOnset    bool   `json:"hasSuddenOnset" gorm:"default:false"`
	HasRapidProgress  bool   `json:"hasRapidProgress" gorm:"default:false"`
	HasPeriodicAttack bool   `json:"hasPeriodicAttack" gorm:"default:false"`
	HasForcedPosture  bool   `json:"hasForcedPosture" gorm:"default:false"`
	HasCalfMusclePain bool   `json:"hasCalfMusclePain" gorm:"default:false"`
	HasOtherSymptoms  bool   `json:"hasOtherSymptoms" gorm:"default:false"`
	OtherSymptomsName string `json:"otherSymptomsName,omitempty" gorm:"size:500"`
}

func (DiagnosisOtherSymptoms) TableName() string { return constants.TableDiagnosisOtherSymptoms }
