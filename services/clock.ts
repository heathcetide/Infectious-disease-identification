import { http } from '@/utils/http'

type learningMaterials = {
	diseaseTypeName: string,
	pageNo: number,
	pageSize: number
}

export const getPageLearningMaterials = (data: learningMaterials) =>{
	return http({
		method: 'POST',
		url: '/railwayemployee/query/pageSelectDiseaseLearningMaterials',
		data,
	})
}

type clockFormParams = {
	city: string,
	county: string,
	health: boolean,
	latitude: number,
	longitude: number,
	locationName: string,
	province: string,
}

export const clockSubmitService = (data: clockFormParams) =>{
	return http({
		method: 'POST',
		url: '/railwayemployee/checkin/checkin',
		data,
	})
}

//提交-修改当月工作环境信息
type clockEnvironment ={
	avgActivityFaceToFace: number,
	avgActivityPhysicalContact: number,
	avgActivityRoomPeople:number,
	avgBreakfastFaceToFace: number,
	avgBreakfastPhysicalContact: number,
	avgBreakfastRoomPeople: number,
	avgDinnerFaceToFace: number,
	avgDinnerPhysicalContact: number,
	avgDinnerRoomPeople: number,
	avgLunchFaceToFace: number,
	avgLunchPhysicalContact: number,
	avgLunchRoomPeople: number,
	avgMonthlyPressure: number,
	avgWorkFaceToFace: number,
	avgWorkPhysicalContact: number,
	avgWorkRoomPeople: number,
	collectiveActivitiesCount: number,
	dormitoryArea: number,
	dormitoryBuildingNumber: number,
	dormitoryOccupants: number,
	dormitoryRoomNumber: number,
	maxTemperature: number,
	minTemperature: number,
	submissionTime: string,
	ventilationCondition: string,
	windowOpenArea: number,
	workStationAltitude: number,
	workStationName: string,
}

export const clockEnvironmentService = (data: clockEnvironment) => {
	return http({
		method: 'POST',
		url: '/railwayemployee/checkin/month',
		data,
	})
}

type clockInfo ={
	date:string,
}
//查看当日的打卡情况
export const clockSituationService = (data: clockInfo) =>{
	return http({
		method:'GET',
		url: '/railwayemployee/checkin/select/daily',
		data,
	})
}


type AllCheckInfo = {
  circulatorysymptoms: {
    circulatorySymptomsId: number,
    hasArrhythmia: boolean,
    hasChestPain: boolean,
    hasLowBloodPressure: boolean,
    hasPalpitation: boolean,
    hasRapidPulse: boolean,
    statusId: number
  },
  digestivesymptoms: {
    arrheaFrequencyGEThreeTimesPerDay: boolean,
    diarrheaMode: string,
    diarrheaTimesPerDay: number,
    digestiveSymptomsId: number,
    hasAbdominalDistension: boolean,
    hasAbdominalPain: boolean,
    hasAppetiteLoss: boolean,
    hasBorborygmus: boolean,
    hasConstipation: boolean,
    hasNausea: boolean,
    hasOliguriaOrAnuria: boolean,
    hasStoolType1: boolean,
    hasStoolType2: boolean,
    hasUpperAbdominalDiscomfort: boolean,
    hasVomiting: boolean,
    statusId: number,
    stoolAmount: string,
    stoolOdor: string,
    stoolType1Detail: string,
    stoolType2Detail: string,
    vomitingMode: string
  },
  generalsymptoms: {
    dehydration: string,
    feverDuration: number,
    generalSymptomsId: number,
    hasBloodBlisters: boolean,
    hasChills: boolean,
    hasCongestiveOrPetechialRash: boolean,
    hasCyanosis: boolean,
    hasDehydration: boolean,
    hasEdema: boolean,
    hasExhaustion: boolean,
    hasFatigue: boolean,
    hasFever: boolean,
    hasHeadache: boolean,
    hasItchyRash: boolean,
    hasJointPain: boolean,
    hasLymphNodeSwelling: boolean,
    hasLymphNodeSwellingArmpit: boolean,
    hasLymphNodeSwellingElbow: boolean,
    hasLymphNodeSwellingGroin: boolean,
    hasLymphNodeSwellingNeck: boolean,
    hasLymphNodeSwellingPopliteal: boolean,
    hasLymphNodeSwellingSubclavian: boolean,
    hasMusclePain: boolean,
	hasNightSweats: boolean,
    hasPainfulRedRash: boolean,
    hasSkinUlcer: boolean,
    hasSubcutaneousAndMucosalBleedingSpots: boolean,
    hasSweating: boolean,
	hasWeightLoss: boolean,
    highestTemperature: number,
    itchyRashOnBackOfHands: boolean,
    itchyRashOnFace: boolean,
    itchyRashOnFeet: boolean,
    itchyRashOnFingers: boolean,
    itchyRashOnLowerLimbs: boolean,
    itchyRashOnOther:boolean,
    itchyRashOnUpperLimbs:boolean,
    statusId: number
  },
  localSymptoms: {
    hasConjunctivalCongestion: boolean,
    hasConjunctivitis: boolean,
    hasEyelidEdema: boolean,
    hasFacialRednessOrPallor: boolean,
    hasGumBleeding: boolean,
    hasKidneyPain: boolean,
    hasLipCyanosis: boolean,
    hasLowBackPain: boolean,
    hasMucosalBleedingInOralOrNasal: boolean,
    hasOrbitalPain: boolean,
    hasRednessNeckChest: boolean,
    hasSkinBleedingPoints: boolean,
    hasSkinPallorOrBruising: boolean,
    hasSmellTasteLoss: boolean,
    localSymptomsId: number,
    statusId: number
  },
  neurologicalSymptoms: {
    hasComa: boolean,
    hasConsciousnessDisorder: boolean,
    hasConvulsion: boolean,
    hasDelirium: boolean,
    hasDizziness: boolean,
    hasDrowsiness: boolean,
    hasHearingLoss: boolean,
    hasMuscleParalysisLimbs: boolean,
    hasMuscleParalysisNeckShoulder: boolean,
    hasRestlessness: boolean,
    hasSpeechDisorder: boolean,
    hasStiffNeck: boolean,
    hasSwallowingDifficulty: boolean,
    hasTinnitus: boolean,
    neurologicalSymptomsId: number,
    statusId: number
  },
  otherSymptoms: {
    hasCalfMusclePain: boolean,
    hasForcedPosture: boolean,
    hasOtherSymptoms: boolean,
    hasPeriodicAttack: boolean,
    hasRapidProgress: boolean,
    hasSuddenOnset: boolean,
    nutritionStatus: string,
    otherSymptomsId: number,
    otherSymptomsName: string,
    sleepQuality: string,
    statusId: number,
    symptomSeverity: string,
    workLifeStress: string
  },
  respiratorysymptoms: {
    hasBloodySputum: boolean,
    hasBreathingDifficulty: boolean,
    hasBronchitis: boolean,
    hasChestTightness: boolean,
    hasCough: boolean,
    hasDryThroat: boolean,
    hasHemoptysis: boolean,
    hasHoarseness: boolean,
    hasNasalCongestion: boolean,
    hasRunnyNose: boolean,
    hasShortnessOfBreath: boolean,
    hasSoreThroat: boolean,
    hasSputum: boolean,
    respiratorySymptomsId: number,
    statusId: number
  },
  riskFactorsAndExposure: {
    anthraxArea: string,
    contactDate: string,
    contactPlace: string,
    contactWithAnimalProducts: string,
    contactWithAnimals: string,
    contactWithDiarrheaPatient: string,
    contactWithFeverPatient: string,
    contactWithPatientExcreta: string,
    contactWithRashPatient: string,
    contactWithRatIn2Months: string,
    contactWithTuberculosisPatient: string,
    diningOut: boolean,
    diningPlace: string,
    diningUnclear: boolean,
    drunkWaterFromDitchesOrPonds: string,
    eatenFoodContaminatedByRatFeces: string,
    fleaBite: string,
    foodNames: string,
    groupDining: boolean,
    groupOutbreak: string,
    hasContactWithBird: boolean,
    hasContactWithCat: boolean,
    hasContactWithCow: boolean,
    hasContactWithDog: boolean,
    hasContactWithDonkey: boolean,
    hasContactWithFlea: boolean,
    hasContactWithHorse: boolean,
    hasContactWithLice: boolean,
    hasContactWithMarmot: boolean,
    hasContactWithMule: boolean,
    hasContactWithOtherAnimal: boolean,
    hasContactWithPig: boolean,
    hasContactWithRabbit: boolean,
    hasContactWithRat: boolean,
    hasContactWithSheep: boolean,
    hasContactWithTiger: boolean,
    hasContactWithWolf: boolean,
    hasDrinkingRawWater: boolean,
    hasEatingColdCookedFood: boolean,
    hasEatingRawFood: boolean,
    hasEatingSeafood: boolean,
    hasMosquitoBite: boolean,
    hillyOrMountainous: boolean,
    malariaArea: string,
    otherAnimalName: string,
    otherResidenceType: boolean,
    otherWildActivityName: string,
    outdoorStayOrWorkWithinMonth: string,
    plagueArea: string,
    plain: boolean,
    ratContactMethod: string,
    ratOrRatDroppingsAtWorkplace: string,
    restedNearRatHoles: string,
    riskFactorsAndExposureId: number,
    specificAnthraxLocation: string,
    specificMalariaLocation: string,
    specificPlagueLocation: string,
    statusId: number,
    tickBite: string,
    wildActivityHistory: string
  }
}
export const submitAllCheckService = (data: AllCheckInfo) =>{
	return http({
		method:'POST',
		url: '/railwayemployee/checkin/all',
		data,
	})
}


export const environmentInfoService = (data: string) =>{
	return http({
		method: 'POST',
		url: '/railwayemployee/checkin/environment/info',
		data,
	})
}


type DiagnosisSubmit = {
  diagnosisResultsEmpDTO: {
    admissionDate: string;
    admissionSymptomsAndSigns: string;
    anthraxSubtype: string;
    deathDate: string;
    dischargeDate: string;
    discoveryMethod: string;
    diseaseOutcome: string;
    diseaseType: string;
    hospitalName: string;
    otherDiscoveryMethodName: string;
    otherDiseaseName: string;
    otherRegistrationDetails: string;
    plagueSubtype: string;
    registrationClassification: string;
  };
  diagnosiscirculatorysymptoms: {
    hasArrhythmia: boolean;
    hasChestPain: boolean;
    hasPalpitation: boolean;
    hasRapidPulse: boolean;
  };
  diagnosiscomplications: {
    hasAbdominalHemorrhage: boolean;
    hasAcuteKidneyInjury: boolean;
    hasAcuteLungInjury: boolean;
    hasAnemia: boolean;
    hasArds: boolean;
    hasBacteremia: boolean;
    hasBacterialPneumonia: boolean;
    hasBrainEdema: boolean;
    hasBronchitis: boolean;
    hasBronchopneumonia: boolean;
    hasCoagulationDisorder: boolean;
    hasComplications: boolean;
    hasDic: boolean;
    hasElectrolyteImbalance: boolean;
    hasEncephalitis: boolean;
    hasEpilepsy: boolean;
    hasFluidOverload: boolean;
    hasFungalPneumonia: boolean;
    hasGastrointestinalBleeding: boolean;
    hasHematemesis: boolean;
    hasHemolysis: boolean;
    hasHemoptysis: boolean;
    hasHepatosplenomegaly: boolean;
    hasHus: boolean;
    hasHypertensiveEncephalopathy: boolean;
    hasHypoxemia: boolean;
    hasIntracranialHemorrhage: boolean;
    hasLiverDamage: boolean;
    hasMelena: boolean;
    hasMeningitis: boolean;
    hasMetabolicAcidosis: boolean;
    hasMultipleOrganFailure: boolean;
    hasMyocardialDamage: boolean;
    hasMyocarditis: boolean;
    hasNasalBleeding: boolean;
    hasPulmonaryEdema: boolean;
    hasRespiratoryFailure: boolean;
    hasSecondaryInfection: boolean;
    hasSepticShock: boolean;
    hasSpontaneousRenalRupture: boolean;
    hasSuperinfection: boolean;
    hasToxicHepatitis: boolean;
    hasToxicShock: boolean;
    hasVaginalBleeding: boolean;
    hasViralPneumonia: boolean;
    otherComplications: boolean;
    otherComplicationsName: string;
  };
  diagnosisdigestivesymptoms: {
    diarrheaFrequencyGEThreeTimesPerDay: boolean;
    diarrheaMode: string;
    diarrheaTimesPerDay: number;
    hasAbdominalDistension: boolean;
    hasAbdominalPain: boolean;
    hasAppetiteLoss: boolean;
    hasBorborygmus: boolean;
    hasConstipation: boolean;
    hasNausea: boolean;
    hasOliguriaOrAnuria: boolean;
    hasStoolType1: boolean;
    hasStoolType2: boolean;
    hasUpperAbdominalDiscomfort: boolean;
    hasVomiting: boolean;
    stoolAmount: string;
    stoolOdor: string;
    stoolType1Detail: string;
    stoolType2Detail: string;
    vomitingMode: string;
  };
  diagnosisexaminations: {
    hasBloodTest: boolean;
    hasImaging: boolean;
    hasStoolTest: boolean;
    hasUltrasound: boolean;
    hasUrineTest: boolean;
  };
  diagnosisgeneralsymptoms: {
    bleedingCavity: string;
    bleedingVolume: number;
    dehydrationSeverity: string;
    feverDuration: number;
    hasBleeding: boolean;
    hasBloodBlisters: boolean;
    hasBloodPressureDrop: boolean;
    hasBrudzinskiSign: boolean;
    hasChills: boolean;
    hasCongestiveOrPetechialRash: boolean;
    hasCyanosis: boolean;
    hasDehydration: boolean;
    hasEdema: boolean;
    hasExhaustion: boolean;
    hasFatigue: boolean;
    hasFever: boolean;
    hasHeadache: boolean;
    hasHighIntracranialPressure: boolean;
    hasItchyRash: boolean;
    hasJointPain: boolean;
    hasKernigSign: boolean;
    hasLymphNodeSwelling: boolean;
    hasLymphNodeSwellingArmpit: boolean;
    hasLymphNodeSwellingElbow: boolean;
    hasLymphNodeSwellingGroin: boolean;
    hasLymphNodeSwellingNeck: boolean;
    hasLymphNodeSwellingPopliteal: boolean;
    hasLymphNodeSwellingSubclavian: boolean;
    hasMusclePain: boolean;
    hasNightSweats: boolean;
    hasOtherOrganDamage: boolean;
    hasPainfulRedRash: boolean;
    hasPlasmaLeakage: boolean;
    hasSevereAnemia: boolean;
    hasShockSyndrome: boolean;
    hasSkinUlcer: boolean;
    hasSubcutaneousAndMucosalBleedingCavity: boolean;
    hasSubcutaneousAndMucosalBleedingSpots: boolean;
    hasSweating: boolean;
    hasTurbidCerebrospinalFluid: boolean;
    hasWeightLoss: boolean;
    hasWhiteBloodCells: boolean;
    highestTemperature: number;
    itchyRashOnBackOfHands: boolean;
    itchyRashOnFace: boolean;
    itchyRashOnFeet: boolean;
    itchyRashOnFingers: boolean;
    itchyRashOnLowerLimbs: boolean;
    itchyRashOnOther: boolean;
    itchyRashOnUpperLimbs: boolean;
    lowestBloodPressure: number;
  };
  diagnosislocalsymptoms: {
    hasConjunctivalCongestion: boolean;
    hasConjunctivitis: boolean;
    hasEyelidEdema: boolean;
    hasFacialRednessOrPallor: boolean;
    hasGumBleeding: boolean;
    hasHepatomegaly: boolean;
    hasKidneyPain: boolean;
    hasLipCyanosis: boolean;
    hasLowBackPain: boolean;
    hasMucosalBleedingInOralOrNasal: boolean;
    hasOrbitalPain: boolean;
    hasRednessNeckChest: boolean;
    hasSkinBleedingPoints: boolean;
    hasSkinPallorOrBruising: boolean;
    hasSmellTasteLoss: boolean;
    hasSplenomegaly: boolean;
  };
  diagnosisneurologicalsymptoms: {
    hasComa: boolean;
    hasConsciousnessDisorder: boolean;
    hasConvulsion: boolean;
    hasDelirium: boolean;
    hasDizziness: boolean;
    hasDrowsiness: boolean;
    hasHearingLoss: boolean;
    hasMuscleParalysisLimbs: boolean;
    hasMuscleParalysisNeckShoulder: boolean;
    hasRestlessness: boolean;
    hasSpeechDisorder: boolean;
    hasStiffNeck: boolean;
    hasSwallowingDifficulty: boolean;
    hasTinnitus: boolean;
  };
  diagnosisothersymptoms: {
    hasCalfMusclePain: boolean;
    hasForcedPosture: boolean;
    hasOtherSymptoms: boolean;
    hasPeriodicAttack: boolean;
    hasRapidProgress: boolean;
    hasSuddenOnset: boolean;
    otherSymptomsName: string;
  };
  diagnosispersonalinfo: {
    bloodOxygen: number;
    bloodPressure: string;
    bloodSugarFasting: number;
    bloodSugarRandom: number;
    cholesterolLipid: number;
    cholesterolTotal: number;
    hdl: number;
    height: number;
    ldl: number;
    triglycerides: number;
    weight: number;
  };
  diagnosisrespiratorysymptoms: {
    hasBloodySputum: boolean;
    hasBreathingDifficulty: boolean;
    hasBronchitis: boolean;
    hasChestTightness: boolean;
    hasCough: boolean;
    hasDryThroat: boolean;
    hasHemoptysis: boolean;
    hasHoarseness: boolean;
    hasLungRales: boolean;
    hasNasalCongestion: boolean;
    hasPharyngealHyperemia: boolean;
    hasRunnyNose: boolean;
    hasShortnessOfBreath: boolean;
    hasSoreThroat: boolean;
    hasSputum: boolean;
  };
  fileIdList: number[];
}

export const diagnosisSubmitService = (data: DiagnosisSubmit) =>{
	return http({
		method: 'POST',
		url: '/railwayemployee/diagnosis/result',
		data,
	})
}

export const clockRecordService = (data: String) =>{
	return http({
		method: 'POST',
		url: '/railwayemployee/checkin/select/month',
		data
	})
}

export const getReportService = (data: String) =>{
	return http({
		method: 'POST',
		url: '/railwayemployee/diagnosis/report',
		data
	})
}

export const getLocalReportService = (data: String) =>{
	return http({
		method: 'POST',
		url: '/railwayemployee/labtest/get/report',
		data
	})
}

export const getDailyReportService = (data: String) =>{
	return http({
		method: 'POST',
		url: '/railwayemployee/checkin/get/daily',
		data
	})
}