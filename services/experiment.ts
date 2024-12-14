import { http } from '@/utils/http'

type reportParams ={
	pathogenicTestResults: string,
	serologicalTestDone: boolean,
	virusAntigenTestDone: boolean,
	virusCultureIsolationDone: boolean,
	virusNucleicAcidTestDone: boolean
}

export const reportSubmitService = (data: reportParams) =>{
	return http({
		method: 'POST',
		url: '/railwayemployee/labtest/report',
		data,
	})
}

type AllCheckInfo = {
  diagnosisResultsEmpDTO: {
    admissionDate: string;
    admissionSymptomsAndSigns: string;
    anthraxSubtype: string;
    deathDate?: string;
    dischargeDate: string;
    discoveryMethod: string;
    diseaseOutcome: string;
    diseaseType: string;
    hospitalName: string;
    otherDiscoveryMethodName?: string;
    otherDiseaseName?: string;
    otherRegistrationDetails?: string;
    plagueSubtype?: string;
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
    otherComplicationsName?: string;
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
    otherSymptomsName?: string;
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
  fileIdList: any[];
};

export const diagnosisService = (data : AllCheckInfo) =>{
	return http({
		method: 'POST',
		url: '/railwayemployee/diagnosis/result',
		data,
	})
}