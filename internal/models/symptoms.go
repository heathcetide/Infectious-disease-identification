// Copyright (c) 2026 heathcetide. All rights reserved.

package models

import (
	"time"

	"github.com/heathcetide/infectious-disease/internal/constants"
)

// GeneralSymptoms records whole-body symptoms for a daily health status.
type GeneralSymptoms struct {
	GeneralSymptomsID uint `json:"generalSymptomsId" gorm:"primaryKey;autoIncrement"`
	StatusID          uint `json:"statusId" gorm:"index;not null"`

	HasFever                               bool    `json:"hasFever" gorm:"default:false"`
	HighestTemperature                     float64 `json:"highestTemperature,omitempty"`
	FeverDuration                          int     `json:"feverDuration,omitempty"`
	HasChills                              bool    `json:"hasChills" gorm:"default:false"`
	HasSweating                            bool    `json:"hasSweating" gorm:"default:false"`
	HasFatigue                             bool    `json:"hasFatigue" gorm:"default:false"`
	HasHeadache                            bool    `json:"hasHeadache" gorm:"default:false"`
	HasMusclePain                          bool    `json:"hasMusclePain" gorm:"default:false"`
	HasJointPain                           bool    `json:"hasJointPain" gorm:"default:false"`
	HasLymphNodeSwelling                   bool    `json:"hasLymphNodeSwelling" gorm:"default:false"`
	HasLymphNodeSwellingGroin              bool    `json:"hasLymphNodeSwellingGroin" gorm:"default:false"`
	HasLymphNodeSwellingArmpit             bool    `json:"hasLymphNodeSwellingArmpit" gorm:"default:false"`
	HasLymphNodeSwellingSubclavian         bool    `json:"hasLymphNodeSwellingSubclavian" gorm:"default:false"`
	HasLymphNodeSwellingNeck               bool    `json:"hasLymphNodeSwellingNeck" gorm:"default:false"`
	HasLymphNodeSwellingElbow              bool    `json:"hasLymphNodeSwellingElbow" gorm:"default:false"`
	HasLymphNodeSwellingPopliteal          bool    `json:"hasLymphNodeSwellingPopliteal" gorm:"default:false"`
	HasCyanosis                            bool    `json:"hasCyanosis" gorm:"default:false"`
	HasSubcutaneousAndMucosalBleedingSpots bool    `json:"hasSubcutaneousAndMucosalBleedingSpots" gorm:"default:false"`
	HasPainfulRedRash                      bool    `json:"hasPainfulRedRash" gorm:"default:false"`
	HasBloodBlisters                       bool    `json:"hasBloodBlisters" gorm:"default:false"`
	HasSkinUlcer                           bool    `json:"hasSkinUlcer" gorm:"default:false"`
	HasCongestiveOrPetechialRash           bool    `json:"hasCongestiveOrPetechialRash" gorm:"default:false"`
	HasDehydration                         bool    `json:"hasDehydration" gorm:"default:false"`
	Dehydration                            string  `json:"dehydration,omitempty" gorm:"size:20"`
	HasItchyRash                           bool    `json:"hasItchyRash" gorm:"default:false"`
	ItchyRashOnFingers                     bool    `json:"itchyRashOnFingers" gorm:"default:false"`
	ItchyRashOnBackOfHands                 bool    `json:"itchyRashOnBackOfHands" gorm:"default:false"`
	ItchyRashOnUpperLimbs                  bool    `json:"itchyRashOnUpperLimbs" gorm:"default:false"`
	ItchyRashOnLowerLimbs                  bool    `json:"itchyRashOnLowerLimbs" gorm:"default:false"`
	ItchyRashOnFeet                        bool    `json:"itchyRashOnFeet" gorm:"default:false"`
	ItchyRashOnFace                        bool    `json:"itchyRashOnFace" gorm:"default:false"`
	ItchyRashOnOther                       bool    `json:"itchyRashOnOther" gorm:"default:false"`
	HasEdema                               bool    `json:"hasEdema" gorm:"default:false"`
	HasNightSweats                         bool    `json:"hasNightSweats" gorm:"default:false"`
	HasWeightLoss                          bool    `json:"hasWeightLoss" gorm:"default:false"`
	HasExhaustion                          bool    `json:"hasExhaustion" gorm:"default:false"`
}

func (GeneralSymptoms) TableName() string { return constants.TableGeneralSymptoms }

// RespiratorySymptoms records respiratory symptoms for a daily health status.
type RespiratorySymptoms struct {
	RespiratorySymptomsID uint `json:"respiratorySymptomsId" gorm:"primaryKey;autoIncrement"`
	StatusID              uint `json:"statusId" gorm:"index;not null"`

	HasCough               bool `json:"hasCough" gorm:"default:false"`
	HasSputum              bool `json:"hasSputum" gorm:"default:false"`
	HasBloodySputum        bool `json:"hasBloodySputum" gorm:"default:false"`
	HasHemoptysis          bool `json:"hasHemoptysis" gorm:"default:false"`
	HasSoreThroat          bool `json:"hasSoreThroat" gorm:"default:false"`
	HasDryThroat           bool `json:"hasDryThroat" gorm:"default:false"`
	HasNasalCongestion     bool `json:"hasNasalCongestion" gorm:"default:false"`
	HasRunnyNose           bool `json:"hasRunnyNose" gorm:"default:false"`
	HasChestTightness      bool `json:"hasChestTightness" gorm:"default:false"`
	HasShortnessOfBreath   bool `json:"hasShortnessOfBreath" gorm:"default:false"`
	HasBreathingDifficulty bool `json:"hasBreathingDifficulty" gorm:"default:false"`
	HasHoarseness          bool `json:"hasHoarseness" gorm:"default:false"`
	HasBronchitis          bool `json:"hasBronchitis" gorm:"default:false"`
}

func (RespiratorySymptoms) TableName() string { return constants.TableRespiratorySymptoms }

// DigestiveSymptoms records digestive symptoms for a daily health status.
type DigestiveSymptoms struct {
	DigestiveSymptomsID uint `json:"digestiveSymptomsId" gorm:"primaryKey;autoIncrement"`
	StatusID            uint `json:"statusId" gorm:"index;not null"`

	ArrheaFrequencyGEThreeTimesPerDay bool   `json:"arrheaFrequencyGEThreeTimesPerDay" gorm:"default:false"`
	DiarrheaTimesPerDay               int    `json:"diarrheaTimesPerDay,omitempty"`
	HasStoolType1                     bool   `json:"hasStoolType1" gorm:"default:false"`
	StoolType1Detail                  string `json:"stoolType1Detail,omitempty" gorm:"size:100"`
	HasStoolType2                     bool   `json:"hasStoolType2" gorm:"default:false"`
	StoolType2Detail                  string `json:"stoolType2Detail,omitempty" gorm:"size:100"`
	DiarrheaMode                      string `json:"diarrheaMode,omitempty" gorm:"size:50"`
	StoolAmount                       string `json:"stoolAmount,omitempty" gorm:"size:20"`
	StoolOdor                         string `json:"stoolOdor,omitempty" gorm:"size:50"`
	HasVomiting                       bool   `json:"hasVomiting" gorm:"default:false"`
	VomitingMode                      string `json:"vomitingMode,omitempty" gorm:"size:50"`
	HasNausea                         bool   `json:"hasNausea" gorm:"default:false"`
	HasAppetiteLoss                   bool   `json:"hasAppetiteLoss" gorm:"default:false"`
	HasAbdominalDistension            bool   `json:"hasAbdominalDistension" gorm:"default:false"`
	HasAbdominalPain                  bool   `json:"hasAbdominalPain" gorm:"default:false"`
	HasBorborygmus                    bool   `json:"hasBorborygmus" gorm:"default:false"`
	HasUpperAbdominalDiscomfort       bool   `json:"hasUpperAbdominalDiscomfort" gorm:"default:false"`
	HasConstipation                   bool   `json:"hasConstipation" gorm:"default:false"`
	HasOliguriaOrAnuria               bool   `json:"hasOliguriaOrAnuria" gorm:"default:false"`
}

func (DigestiveSymptoms) TableName() string { return constants.TableDigestiveSymptoms }

// CirculatorySymptoms records circulatory symptoms for a daily health status.
type CirculatorySymptoms struct {
	CirculatorySymptomsID uint `json:"circulatorySymptomsId" gorm:"primaryKey;autoIncrement"`
	StatusID              uint `json:"statusId" gorm:"index;not null"`

	HasArrhythmia       bool `json:"hasArrhythmia" gorm:"default:false"`
	HasChestPain        bool `json:"hasChestPain" gorm:"default:false"`
	HasRapidPulse       bool `json:"hasRapidPulse" gorm:"default:false"`
	HasPalpitation      bool `json:"hasPalpitation" gorm:"default:false"`
	HasLowBloodPressure bool `json:"hasLowBloodPressure" gorm:"default:false"`
}

func (CirculatorySymptoms) TableName() string { return constants.TableCirculatorySymptoms }

// NeurologicalSymptoms records neurological symptoms for a daily health status.
type NeurologicalSymptoms struct {
	NeurologicalSymptomsID uint `json:"neurologicalSymptomsId" gorm:"primaryKey;autoIncrement"`
	StatusID               uint `json:"statusId" gorm:"index;not null"`

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

func (NeurologicalSymptoms) TableName() string { return constants.TableNeurologicalSymptoms }

// LocalSymptoms records localized symptoms for a daily health status.
type LocalSymptoms struct {
	LocalSymptomsID uint `json:"localSymptomsId" gorm:"primaryKey;autoIncrement"`
	StatusID        uint `json:"statusId" gorm:"index;not null"`

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
	HasMucosalBleedingInOralOrNasal bool `json:"hasMucosalBleedingInOralOrNasal" gorm:"default:false"`
	HasSkinBleedingPoints           bool `json:"hasSkinBleedingPoints" gorm:"default:false"`
}

func (LocalSymptoms) TableName() string { return constants.TableLocalSymptoms }

// OtherSymptoms records miscellaneous symptoms for a daily health status.
type OtherSymptoms struct {
	OtherSymptomsID uint `json:"otherSymptomsId" gorm:"primaryKey;autoIncrement"`
	StatusID        uint `json:"statusId" gorm:"index;not null"`

	HasSuddenOnset    bool   `json:"hasSuddenOnset" gorm:"default:false"`
	HasRapidProgress  bool   `json:"hasRapidProgress" gorm:"default:false"`
	HasPeriodicAttack bool   `json:"hasPeriodicAttack" gorm:"default:false"`
	HasForcedPosture  bool   `json:"hasForcedPosture" gorm:"default:false"`
	HasCalfMusclePain bool   `json:"hasCalfMusclePain" gorm:"default:false"`
	SleepQuality      string `json:"sleepQuality,omitempty" gorm:"size:20"`
	NutritionStatus   string `json:"nutritionStatus,omitempty" gorm:"size:20"`
	WorkLifeStress    string `json:"workLifeStress,omitempty" gorm:"size:20"`
	SymptomSeverity   string `json:"symptomSeverity,omitempty" gorm:"size:20"`
	HasOtherSymptoms  bool   `json:"hasOtherSymptoms" gorm:"default:false"`
	OtherSymptomsName string `json:"otherSymptomsName,omitempty" gorm:"size:500"`
}

func (OtherSymptoms) TableName() string { return constants.TableOtherSymptoms }

// RiskFactorsAndExposure records exposure and risk factors for a daily health status.
type RiskFactorsAndExposure struct {
	RiskFactorsAndExposureID uint `json:"riskFactorsAndExposureId" gorm:"primaryKey;autoIncrement"`
	StatusID                 uint `json:"statusId" gorm:"index;not null"`

	ContactWithFeverPatient         string     `json:"contactWithFeverPatient,omitempty" gorm:"size:20"`
	ContactWithDiarrheaPatient      string     `json:"contactWithDiarrheaPatient,omitempty" gorm:"size:20"`
	ContactWithRashPatient          string     `json:"contactWithRashPatient,omitempty" gorm:"size:20"`
	ContactWithTuberculosisPatient  string     `json:"contactWithTuberculosisPatient,omitempty" gorm:"size:20"`
	IsGroupDining                   bool       `json:"isGroupDining" gorm:"default:false"`
	IsDiningOut                     bool       `json:"isDiningOut" gorm:"default:false"`
	IsDiningUnclear                 bool       `json:"isDiningUnclear" gorm:"default:false"`
	HasDrinkingRawWater             bool       `json:"hasDrinkingRawWater" gorm:"default:false"`
	HasEatingRawFood                bool       `json:"hasEatingRawFood" gorm:"default:false"`
	HasEatingColdCookedFood         bool       `json:"hasEatingColdCookedFood" gorm:"default:false"`
	HasEatingSeafood                bool       `json:"hasEatingSeafood" gorm:"default:false"`
	FoodNames                       string     `json:"foodNames,omitempty" gorm:"size:500"`
	DiningPlace                     string     `json:"diningPlace,omitempty" gorm:"size:100"`
	GroupOutbreak                   string     `json:"groupOutbreak,omitempty" gorm:"size:20"`
	OutdoorStayOrWorkWithinMonth    string     `json:"outdoorStayOrWorkWithinMonth,omitempty" gorm:"size:20"`
	WildActivityHistory             string     `json:"wildActivityHistory,omitempty" gorm:"size:100"`
	OtherWildActivityName           string     `json:"otherWildActivityName,omitempty" gorm:"size:100"`
	PlagueArea                      string     `json:"plagueArea,omitempty" gorm:"size:20"`
	SpecificPlagueLocation          string     `json:"specificPlagueLocation,omitempty" gorm:"size:200"`
	AnthraxArea                     string     `json:"anthraxArea,omitempty" gorm:"size:20"`
	SpecificAnthraxLocation         string     `json:"specificAnthraxLocation,omitempty" gorm:"size:200"`
	MalariaArea                     string     `json:"malariaArea,omitempty" gorm:"size:20"`
	SpecificMalariaLocation         string     `json:"specificMalariaLocation,omitempty" gorm:"size:200"`
	ContactWithAnimalProducts       string     `json:"contactWithAnimalProducts,omitempty" gorm:"size:20"`
	ContactDate                     *time.Time `json:"contactDate,omitempty" gorm:"type:date"`
	ContactWithAnimals              string     `json:"contactWithAnimals,omitempty" gorm:"size:20"`
	ContactPlace                    string     `json:"contactPlace,omitempty" gorm:"size:200"`
	HasContactWithRat               bool       `json:"hasContactWithRat" gorm:"default:false"`
	HasContactWithMarmot            bool       `json:"hasContactWithMarmot" gorm:"default:false"`
	HasContactWithLice              bool       `json:"hasContactWithLice" gorm:"default:false"`
	HasContactWithFlea              bool       `json:"hasContactWithFlea" gorm:"default:false"`
	HasContactWithCat               bool       `json:"hasContactWithCat" gorm:"default:false"`
	HasContactWithBird              bool       `json:"hasContactWithBird" gorm:"default:false"`
	HasContactWithPig               bool       `json:"hasContactWithPig" gorm:"default:false"`
	HasContactWithHorse             bool       `json:"hasContactWithHorse" gorm:"default:false"`
	HasContactWithMule              bool       `json:"hasContactWithMule" gorm:"default:false"`
	HasContactWithDonkey            bool       `json:"hasContactWithDonkey" gorm:"default:false"`
	HasContactWithCow               bool       `json:"hasContactWithCow" gorm:"default:false"`
	HasContactWithSheep             bool       `json:"hasContactWithSheep" gorm:"default:false"`
	HasContactWithDog               bool       `json:"hasContactWithDog" gorm:"default:false"`
	HasContactWithRabbit            bool       `json:"hasContactWithRabbit" gorm:"default:false"`
	HasContactWithWolf              bool       `json:"hasContactWithWolf" gorm:"default:false"`
	HasContactWithTiger             bool       `json:"hasContactWithTiger" gorm:"default:false"`
	HasContactWithOtherAnimal       bool       `json:"hasContactWithOtherAnimal" gorm:"default:false"`
	OtherAnimalName                 string     `json:"otherAnimalName,omitempty" gorm:"size:100"`
	IsHillyOrMountainous            bool       `json:"isHillyOrMountainous" gorm:"default:false"`
	IsPlain                         bool       `json:"isPlain" gorm:"default:false"`
	IsOtherResidenceType            bool       `json:"isOtherResidenceType" gorm:"default:false"`
	HasMosquitoBite                 bool       `json:"hasMosquitoBite" gorm:"default:false"`
	TickBite                        string     `json:"tickBite,omitempty" gorm:"size:20"`
	FleaBite                        string     `json:"fleaBite,omitempty" gorm:"size:20"`
	ContactWithRatIn2Months         string     `json:"contactWithRatIn2Months,omitempty" gorm:"size:20"`
	RatContactMethod                string     `json:"ratContactMethod,omitempty" gorm:"size:100"`
	EatenFoodContaminatedByRatFeces string     `json:"eatenFoodContaminatedByRatFeces,omitempty" gorm:"size:20"`
	DrunkWaterFromDitchesOrPonds    string     `json:"drunkWaterFromDitchesOrPonds,omitempty" gorm:"size:20"`
	RestedNearRatHoles              string     `json:"restedNearRatHoles,omitempty" gorm:"size:20"`
	RatOrRatDroppingsAtWorkplace    string     `json:"ratOrRatDroppingsAtWorkplace,omitempty" gorm:"size:20"`
	ContactWithPatientExcreta       string     `json:"contactWithPatientExcreta,omitempty" gorm:"size:20"`
}

func (RiskFactorsAndExposure) TableName() string { return constants.TableRiskFactorsAndExposure }
