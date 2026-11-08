<script>
import { diagnosisSubmitService,getReportService } from '../../services/clock'
import { tip } from '../../utils/tip'
import { baseURL } from '../../utils/http'
import ReportUpload from '../../components/report-upload/report-upload.vue'
export default {
  components: { ReportUpload },
  data() {
    return {
		submitting: false,
		formData: {
		  stoolReport: [],   // 粪便检查报告文件数组
		  bloodReport: [],   // 血液检查报告文件数组
		  imageReport: [],
		  ultrasoundReport: [],
		  urineReport: [],
		  medicalRecord: [],
		},
		AllCheckInfo : {
		  diagnosisResultsEmpDTO: {
			diagnosisDate:'',
		    admissionDate: '',
		    admissionSymptomsAndSigns: '',
		    anthraxSubtype: null,
		    deathDate: null,
		    dischargeDate: '',
		    discoveryMethod: '',
		    diseaseOutcome: null,
		    diseaseType: null,
		    hospitalName: '',
		    otherDiscoveryMethodName: null,
		    otherDiseaseName: null,
		    otherRegistrationDetails: null,
		    plagueSubtype: null,
		    registrationClassification: null
		  },
		  diagnosiscirculatorysymptoms: {
		    hasArrhythmia: false,
		    hasChestPain: false,
		    hasPalpitation: false,
		    hasRapidPulse: false
		  },
		  diagnosiscomplications: {
		    hasAbdominalHemorrhage: false,
		    hasAcuteKidneyInjury: false,
		    hasAcuteLungInjury: false,
		    hasAnemia: false,
		    hasArds: false,
		    hasBacteremia: false,
		    hasBacterialPneumonia: false,
		    hasBrainEdema: false,
		    hasBronchitis: false,
		    hasBronchopneumonia: false,
		    hasCoagulationDisorder: false,
		    hasComplications: false,
		    hasDic: false,
		    hasElectrolyteImbalance: false,
		    hasEncephalitis: false,
		    hasEpilepsy: false,
		    hasFluidOverload: false,
		    hasFungalPneumonia: false,
		    hasGastrointestinalBleeding: false,
		    hasHematemesis: false,
		    hasHemolysis: false,
		    hasHemoptysis: false,
		    hasHepatosplenomegaly: false,
		    hasHus: false,
		    hasHypertensiveEncephalopathy: false,
		    hasHypoxemia: false,
		    hasIntracranialHemorrhage: false,
		    hasLiverDamage: false,
		    hasMelena: false,
		    hasMeningitis: false,
		    hasMetabolicAcidosis: false,
		    hasMultipleOrganFailure: false,
		    hasMyocardialDamage: false,
		    hasMyocarditis: false,
		    hasNasalBleeding: false,
		    hasPulmonaryEdema: false,
		    hasRespiratoryFailure: false,
		    hasSecondaryInfection: false,
		    hasSepticShock: false,
		    hasSpontaneousRenalRupture: false,
		    hasSuperinfection: false,
		    hasToxicHepatitis: false,
		    hasToxicShock: false,
		    hasVaginalBleeding: false,
		    hasViralPneumonia: false,
		    otherComplications: false,
		    otherComplicationsName: null
		  },
		  diagnosisdigestivesymptoms: {
		    diarrheaFrequencyGEThreeTimesPerDay: null,
		    diarrheaMode: null,
		    diarrheaTimesPerDay: '',
		    hasAbdominalDistension: false,
		    hasAbdominalPain: false,
		    hasAppetiteLoss: false,
		    hasBorborygmus: false,
		    hasConstipation: false,
		    hasNausea: false,
		    hasOliguriaOrAnuria: false,
		    hasStoolType1: false,
		    hasStoolType2: false,
		    hasUpperAbdominalDiscomfort: false,
		    hasVomiting: null,
		    stoolAmount: null,
		    stoolOdor: null,
		    stoolType1Detail: null,
		    stoolType2Detail: null,
		    vomitingMode: null
		  },
		  diagnosisexaminations: {
		    hasBloodTest: false,
		    hasImaging: false,
		    hasStoolTest: false,
		    hasUltrasound: false,
		    hasUrineTest: false
		  },
		  diagnosisgeneralsymptoms: {
		    bleedingCavity: null,
		    bleedingVolume: null,
		    dehydrationSeverity: null,
		    feverDuration: null,
		    hasBleeding: false,
		    hasBloodBlisters: false,
		    hasBloodPressureDrop: false,
		    hasBrudzinskiSign: false,
		    hasChills: false,
		    hasCongestiveOrPetechialRash: false,
		    hasCyanosis: false,
		    hasDehydration: false,
		    hasEdema: false,
		    hasExhaustion: false,
		    hasFatigue: false,
		    hasFever: false,
		    hasHeadache: false,
		    hasHighIntracranialPressure: false,
		    hasItchyRash: false,
		    hasJointPain: false,
		    hasKernigSign: false,
		    hasLymphNodeSwelling: false,
		    hasLymphNodeSwellingArmpit: false,
		    hasLymphNodeSwellingElbow: false,
		    hasLymphNodeSwellingGroin: false,
		    hasLymphNodeSwellingNeck: false,
		    hasLymphNodeSwellingPopliteal: false,
		    hasLymphNodeSwellingSubclavian: false,
		    hasMusclePain: false,
		    hasNightSweats: false,
		    hasOtherOrganDamage: false,
		    hasPainfulRedRash: false,
		    hasPlasmaLeakage: false,
		    hasSevereAnemia: false,
		    hasShockSyndrome: false,
		    hasSkinUlcer: false,
		    hasSubcutaneousAndMucosalBleedingCavity: false,
		    hasSubcutaneousAndMucosalBleedingSpots: false,
		    hasSweating: false,
		    hasTurbidCerebrospinalFluid: false,
		    hasWeightLoss: false,
		    hasWhiteBloodCells: false,
		    highestTemperature: '',
		    itchyRashOnBackOfHands: false,
		    itchyRashOnFace: false,
		    itchyRashOnFeet: false,
		    itchyRashOnFingers: false,
		    itchyRashOnLowerLimbs: false,
		    itchyRashOnOther: false,
		    itchyRashOnUpperLimbs: false,
		    lowestBloodPressure: ''
		  },
		  diagnosislocalsymptoms: {
		    hasConjunctivalCongestion: false,
		    hasConjunctivitis: false,
		    hasEyelidEdema: false,
		    hasFacialRednessOrPallor: false,
		    hasGumBleeding: false,
		    hasHepatomegaly: false,
		    hasKidneyPain: false,
		    hasLipCyanosis: false,
		    hasLowBackPain: false,
		    hasMucosalBleedingInOralOrNasal: false,
		    hasOrbitalPain: false,
		    hasRednessNeckChest: false,
		    hasSkinBleedingPoints: false,
		    hasSkinPallorOrBruising: false,
		    hasSmellTasteLoss: false,
		    hasSplenomegaly: false
		  },
		  diagnosisneurologicalsymptoms: {
		    hasComa: false,
		    hasConsciousnessDisorder: false,
		    hasConvulsion: false,
		    hasDelirium: false,
		    hasDizziness: false,
		    hasDrowsiness: false,
		    hasHearingLoss: false,
		    hasMuscleParalysisLimbs: false,
		    hasMuscleParalysisNeckShoulder: false,
		    hasRestlessness: false,
		    hasSpeechDisorder: false,
		    hasStiffNeck: false,
		    hasSwallowingDifficulty: false,
		    hasTinnitus: false
		  },
		  diagnosisothersymptoms: {
		    hasCalfMusclePain: false,
		    hasForcedPosture: false,
		    hasOtherSymptoms: false,
		    hasPeriodicAttack: false,
		    hasRapidProgress: false,
		    hasSuddenOnset: false,
		    otherSymptomsName: null
		  },
		  diagnosispersonalinfo: {
		    bloodOxygen: '',
		    bloodPressure: '',
		    bloodSugarFasting: '',
		    bloodSugarRandom: '',
		    cholesterolLipid: '',
		    cholesterolTotal: '',
		    hdl: '',
		    height: '',
		    ldl: '',
		    triglycerides: '',
		    weight: '',
		  },
		  diagnosisrespiratorysymptoms: {
		    hasBloodySputum: false,
		    hasBreathingDifficulty: false,
		    hasBronchitis: false,
		    hasChestTightness: false,
		    hasCough: false,
		    hasDryThroat: false,
		    hasHemoptysis: false,
		    hasHoarseness: false,
		    hasLungRales: false,
		    hasNasalCongestion: false,
		    hasPharyngealHyperemia: false,
		    hasRunnyNose: false,
		    hasShortnessOfBreath: false,
		    hasSoreThroat: false,
		    hasSputum: false
		  },
		  fileIdList: [],
		  files:[]
		},
	  
	  temporaryChoice1: false,
	  temporaryChoice2: false,
	  otherChoices: false,
	  SeafoodChoices: false,
	  isViewMode: false,
      sonpages: 1,
	  pages : 1,
    };
  },
  onLoad(options) {
    // `options` 是页面加载时获取的传参对象
    const { date } = options;
	console.log('dedao',date,'那么isViewMode=',this.isViewMode)
    if (date) {
      this.isViewMode = true;
      // 调用获取服务方法以查询数据
      this.getDiagnosisData(date);
    }
  },
  methods: {
	async getDiagnosisData(date) {
		getReportService(date)
				.then(res => {
					const data = (res && res.data) || {}
					this.AllCheckInfo = data
					const files = Array.isArray(data.files) ? data.files : []
					files.forEach(file => {
					    const field = file.reportId
					    if (!field) return
					    const fileUrl = `${baseURL}/api/v1/infect-files/diagnosis/${field}`
					    const newFile = {
					        name: `${file.reportType || 'report'}_${field}.pdf`,
					        extname: "pdf",
					        url: fileUrl,
					        path: fileUrl,
					    }
					    if (file.reportType === "粪便常规检查报告") {
					        this.formData.stoolReport.push(newFile)
					    } else if (file.reportType === "血液检查报告") {
					        this.formData.bloodReport.push(newFile)
					    } else if (file.reportType === "影像学检查报告") {
					        this.formData.imageReport.push(newFile)
					    } else if (file.reportType === "B超报告") {
					        this.formData.ultrasoundReport.push(newFile)
					    } else if (file.reportType === "尿常规检查报告") {
					        this.formData.urineReport.push(newFile)
					    } else if (file.reportType === "病案首页") {
					        this.formData.medicalRecord.push(newFile)
					    }
					})
				})
				.catch(error => this.submitError(error))
	},
	handleSubmit() {
	    if (this.pages == 1 && this.sonpages < 2) {
	        this.sonpages++;
	    } else if (this.pages < 11) {
	        this.pages++;
	    } else {
	        this.handleSubmits();
	    }
	},
	onDateChange1(e) {
	  this.AllCheckInfo.diagnosisResultsEmpDTO.diagnosisDate = e.detail.value;
	},
	uploadFile(filePath, reportType, index) {
	  return new Promise((resolve, reject) => {
	    uni.uploadFile({
	      url: '/api/v1/diagnosis/file',
	      filePath: filePath,
	      name: 'file',
	      formData: {
	        reportType: reportType,
	        index: String(index),
	      },
	      success: (res) => {
			try {
			  const body = JSON.parse(res.data);
			  const fileId = body.data?.report?.reportId ?? body.data;
			  if (fileId != null) this.AllCheckInfo.fileIdList.push(fileId);
			  resolve(res);
			} catch (err) {
			  reject(err);
			}
	      },
	      fail: (err) => {
	        reject(err);
	      }
	    });
	  });
	},
	
	handleSubmits() {
	  if (this.isViewMode || this.submitting) return;
	  if (this.AllCheckInfo.diagnosisResultsEmpDTO.diseaseType == null || this.AllCheckInfo.diagnosisResultsEmpDTO.discoveryMethod == null
		|| this.AllCheckInfo.diagnosisResultsEmpDTO.diseaseOutcome == null || this.AllCheckInfo.diagnosisResultsEmpDTO.diagnosisDate == '') {
	    uni.showToast({
	      title: '请填写基本信息',
	      icon: 'none',
	      duration: 2000,
	    });
	    return;
	  }
	  wx.showModal({
	    content: '提交后不可修改，是否确认提交？',
	    cancelText: '取消',
	    confirmText: '确认提交',
	    success: async (res) => {
	      if (res.confirm) {
	        this.submitting = true;
	        tip.loading('正在提交…');
	        const uploadPromises = [];
	        const enqueue = (list, type) => {
	          list.forEach((file, index) => {
	            if (file.path && !String(file.path).startsWith('http')) {
	              uploadPromises.push(this.uploadFile(file.path, type, index));
	            }
	          });
	        };
	        this.AllCheckInfo.fileIdList = [];
	        enqueue(this.formData.stoolReport, '粪便常规检查报告');
	        enqueue(this.formData.bloodReport, '血液检查报告');
	        enqueue(this.formData.imageReport, '影像学检查报告');
	        enqueue(this.formData.ultrasoundReport, 'B超报告');
	        enqueue(this.formData.urineReport, '尿常规检查报告');
	        enqueue(this.formData.medicalRecord, '病案首页');
	        try {
	          if (uploadPromises.length) {
	            tip.loading('正在上传报告…');
	          }
	          await Promise.all(uploadPromises);
	          tip.loading('正在提交诊断…');
	          const result = await diagnosisSubmitService(this.AllCheckInfo);
	          this.submitSuccess(result);
	        } catch (error) {
	          console.error('上传/提交出现错误: ', error);
	          this.submitError(error);
	        } finally {
	          this.submitting = false;
	          tip.hide();
	        }
	      }
	    }
	  });
	},
	submitError(error) {
	  console.log('报错了',error)
	  this.submitting = false;
	  tip.hide();
	  tip.fail('提交失败');
	},
	submitSuccess(response) {
	  console.log("注意注意",response)
	  const ok = response?.code === 200 || response?.code === 1 || response?.data;
	  if (ok) {
	    tip.success('信息提交成功');
	  } else {
	    uni.showModal({
	      title: '信息提交失败',
	      content: response.msg || '信息提交失败',
	      showCancel: false,
	      confirmText: '确定',
	    });
	  }
	},
	handleReport(propName) {
	  console.log(`Changing ${propName}:`, this.reportForm[propName]);
	  this.reportForm[propName] = !this.reportForm[propName];
	},
	isImage(filePath) {
	  const extension = filePath.split('.').pop().toLowerCase();
	  return ['jpg', 'png'].includes(extension);
	},
	handlereturn(){
		if(this.pages >1){
			this.pages--;
		}else if( this.pages == 1 && this.sonpages > 1){
			this.sonpages--;
		}
	},
	handleRegistrationClassificationChange(value){
		this.AllCheckInfo.diagnosisResultsEmpDTO.registrationClassification = value;
		console.log('已经选择',this.AllCheckInfo.diagnosisResultsEmpDTO.registrationClassification)
	},
	onDateChange(e) {
	  this.AllCheckInfo.diagnosisResultsEmpDTO.admissionDate = e.detail.value;
	},
	handleDiseaseTypeChange(value){
		this.AllCheckInfo.diagnosisResultsEmpDTO.diseaseType = value;
		console.log('已经选择',this.AllCheckInfo.diagnosisResultsEmpDTO.diseaseType)
	},
	handleDiseaseOutcomeChange(value){
		this.AllCheckInfo.diagnosisResultsEmpDTO.diseaseOutcome = value;
		console.log('已经选择',this.AllCheckInfo.diagnosisResultsEmpDTO.diseaseOutcome)
	},
	handleDiscoveryMethodChange(value){
		this.AllCheckInfo.diagnosisResultsEmpDTO.discoveryMethod = value;
		console.log('已经选择',this.AllCheckInfo.diagnosisResultsEmpDTO.discoveryMethod)
	},
	showRadio(maleValue){
		this.AllCheckInfo.diagnosisgeneralsymptoms.dehydrationSeverity = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisgeneralsymptoms.dehydrationSeverity)
	},
	handleCheckboxChange(propName) {
	  console.log(`Changing ${propName}:`, this.AllCheckInfo.diagnosisgeneralsymptoms[propName]);
	  this.AllCheckInfo.diagnosisgeneralsymptoms[propName] = !this.AllCheckInfo.diagnosisgeneralsymptoms[propName];
	},
	handleRespiratorysymptomsBox(propName){
	  this.AllCheckInfo.diagnosisrespiratorysymptoms[propName] = !this.AllCheckInfo.diagnosisrespiratorysymptoms[propName];
	  console.log(`Changing :`, this.AllCheckInfo.diagnosisrespiratorysymptoms[propName]);
	},
	handleDigestFrequency(propName){
		console.log(`Changing ${propName}:`, this.AllCheckInfo.diagnosisdigestivesymptoms[propName]);
		this.AllCheckInfo.diagnosisdigestivesymptoms[propName] = !this.AllCheckInfo.diagnosisdigestivesymptoms[propName];
	},
	showDigestRadio(maleValue){
		this.AllCheckInfo.diagnosisdigestivesymptoms.diarrheaMode = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisdigestivesymptoms.diarrheaMode)
	},
	showDigestRadioBystool1(maleValue){
		this.AllCheckInfo.diagnosisdigestivesymptoms.stoolType1Detail = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisdigestivesymptoms.stoolType1Detail)
	},
	showDigestRadioBystool2(maleValue){
		this.AllCheckInfo.diagnosisdigestivesymptoms.stoolType2Detail = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisdigestivesymptoms.stoolType2Detail)
	},
	showDigestVomitingMode(maleValue){
		this.AllCheckInfo.diagnosisdigestivesymptoms.vomitingMode = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisdigestivesymptoms.vomitingMode)
	},
	handleCirculatorysymptoms(propName){
		console.log(`Changing ${propName}:`, this.AllCheckInfo.diagnosiscirculatorysymptoms[propName]);
		this.AllCheckInfo.diagnosiscirculatorysymptoms[propName] = !this.AllCheckInfo.diagnosiscirculatorysymptoms[propName];
	},
	handleNeurologicalSymptoms(propName){
		console.log(`Changing ${propName}:`, this.AllCheckInfo.diagnosisneurologicalsymptoms[propName]);
		this.AllCheckInfo.diagnosisneurologicalsymptoms[propName] = !this.AllCheckInfo.diagnosisneurologicalsymptoms[propName];
	},
	handleLocalSymptoms(propName){
		console.log(`Changing ${propName}:`, this.AllCheckInfo.diagnosislocalsymptoms[propName]);
		this.AllCheckInfo.diagnosislocalsymptoms[propName] = ! this.AllCheckInfo.diagnosislocalsymptoms[propName];
	},
	handleOtherSymptoms(propName){
		console.log(`Changing ${propName}:`, this.AllCheckInfo.diagnosisothersymptoms[propName]);
		this.AllCheckInfo.diagnosisothersymptoms[propName] = !this.AllCheckInfo.diagnosisothersymptoms[propName];
	},
	showSleepQuality(maleValue){
		this.AllCheckInfo.diagnosisothersymptoms.sleepQuality = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisothersymptoms.sleepQuality)
	},
	showNutritionStatus(maleValue){
		this.AllCheckInfo.diagnosisothersymptoms.nutritionStatus = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisothersymptoms.nutritionStatus)
	},
	showWorkLifeStress(maleValue){
		this.AllCheckInfo.diagnosisothersymptoms.workLifeStress = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisothersymptoms.workLifeStress)
	},
	shwoContactWithFeverPatient(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.contactWithFeverPatient = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.contactWithFeverPatient)
	},
	showContactWithDiarrheaPatient(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.contactWithDiarrheaPatient = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.contactWithDiarrheaPatient)
	},
	showContactWithRashPatient(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.contactWithRashPatient = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.contactWithRashPatient)
	},
	showContactWithTuberculosisPatient(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.contactWithTuberculosisPatient = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.contactWithTuberculosisPatient)
	},
	handleFactorSymptoms(propName){
		console.log(`Changing ${propName}:`, this.AllCheckInfo.diagnosisriskFactorsAndExposure[propName]);
		this.AllCheckInfo.diagnosisriskFactorsAndExposure[propName] = !this.AllCheckInfo.diagnosisriskFactorsAndExposure[propName];
	},
	showGroupOutbreak(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.groupOutbreak = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.groupOutbreak)
	},
	showOutdoorStayOrWorkWithinMonth(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.outdoorStayOrWorkWithinMonth = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.outdoorStayOrWorkWithinMonth)
	},
	showWildActivityHistory(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.wildActivityHistory = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.wildActivityHistory)
	},
	showPlagueArea(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.plagueArea = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.plagueArea)
	},
	showAnthraxArea(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.anthraxArea = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.anthraxArea)
	},
	showMalariaArea(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.malariaArea = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.malariaArea)
	},
	showTickBite(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.tickBite = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.tickBite)
	},
	showFleaBite(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.fleaBite = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.fleaBite)
	},
	showContactWithRatIn2Months(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.contactWithRatIn2Months = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.contactWithRatIn2Months)
	},
	showRatContactMethod(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.ratContactMethod = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.ratContactMethod)
	},
	showEatenFoodContaminatedByRatFeces(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.eatenFoodContaminatedByRatFeces = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.eatenFoodContaminatedByRatFeces)
	},
	showDrunkWaterFromDitchesOrPonds(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.drunkWaterFromDitchesOrPonds = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.drunkWaterFromDitchesOrPonds)
	},
	showRestedNearRatHoles(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.restedNearRatHoles = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.restedNearRatHoles)
	},
	showRatOrRatDroppingsAtWorkplace(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.ratOrRatDroppingsAtWorkplace = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.ratOrRatDroppingsAtWorkplace)
	},
	showContactWithPatientExcreta(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.contactWithPatientExcreta = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.contactWithPatientExcreta)
	},
	showOtherWildActivityName(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.otherWildActivityName = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.otherWildActivityName)
	},
	showContactWithAnimals(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.contactWithAnimals = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.contactWithAnimals)
	},
	showContactWithAnimalProducts(maleValue){
		this.AllCheckInfo.diagnosisriskFactorsAndExposure.contactWithAnimalProducts = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisriskFactorsAndExposure.contactWithAnimalProducts)
	},
	showStoolOdor(maleValue){
		this.AllCheckInfo.diagnosisdigestivesymptoms.stoolOdor = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisdigestivesymptoms.stoolOdor)
	},
	showStoolAmount(maleValue){
		this.AllCheckInfo.diagnosisdigestivesymptoms.stoolAmount = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisdigestivesymptoms.stoolAmount)
	},
	hanleGeneralSymprion(propName){
		console.log(`Changing ${propName}:`, this.AllCheckInfo.diagnosisgeneralsymptoms[propName]);
		this.AllCheckInfo.diagnosisgeneralsymptoms[propName] = !this.AllCheckInfo.diagnosisgeneralsymptoms[propName];
	},
	handleDiagnosisexaminations(propName){
		if (this.isViewMode) return;
		this.AllCheckInfo.diagnosisexaminations[propName] = !this.AllCheckInfo.diagnosisexaminations[propName];
	},
	showAnthraxSubtype(value){
		this.AllCheckInfo.diagnosisResultsEmpDTO.anthraxSubtype = value;
		console.log(`Changing ${value}:`, this.AllCheckInfo.diagnosisResultsEmpDTO.anthraxSubtype)
	},
	handleComplications(propName){
		console.log(`Changing ${propName}:`, this.AllCheckInfo.diagnosiscomplications[propName]);
		this.AllCheckInfo.diagnosiscomplications[propName] = !this.AllCheckInfo.diagnosiscomplications[propName];
	},
	setReportFiles(reportType, files) {
	  this.formData[reportType] = files || [];
	},
  },
};
</script>

<template>
  <view class="page">
  <view class="page-step" v-show="pages == 1 && sonpages == 1">
    <image class="bg" src="/static/other-background.png" mode="aspectFill" />
    <scroll-view scroll-y class="scroll" :show-scrollbar="false">
      <view class="page-header"><text class="title-accent">确诊</text>疾病（必填）</view>
      <view class="card">
        <view class="card-title">确诊疾病</view>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.diseaseType == '新型冠状病毒感染' }" @tap="!isViewMode && handleDiseaseTypeChange('新型冠状病毒感染')">新型冠状病毒感染</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.diseaseType == '流感' }" @tap="!isViewMode && handleDiseaseTypeChange('流感')">流感</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.diseaseType == '鼠疫' }" @tap="!isViewMode && handleDiseaseTypeChange('鼠疫')">鼠疫</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.diseaseType == '感染性腹泻' }" @tap="!isViewMode && handleDiseaseTypeChange('感染性腹泻')">感染性腹泻</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.diseaseType == '炭疽' }" @tap="!isViewMode && handleDiseaseTypeChange('炭疽')">炭疽</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.diseaseType == '登革热（蚊媒传染病）' }" @tap="!isViewMode && handleDiseaseTypeChange('登革热（蚊媒传染病）')">登革热（蚊媒传染病）</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.diseaseType == '疟疾（蚊媒传染病）' }" @tap="!isViewMode && handleDiseaseTypeChange('疟疾（蚊媒传染病）')">疟疾（蚊媒传染病）</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.diseaseType == '森林脑炎（蜱媒传染病）' }" @tap="!isViewMode && handleDiseaseTypeChange('森林脑炎（蜱媒传染病）')">森林脑炎（蜱媒传染病）</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.diseaseType == '发热伴血小板减少综合征（蜱媒传染病）' }" @tap="!isViewMode && handleDiseaseTypeChange('发热伴血小板减少综合征（蜱媒传染病）')">发热伴血小板减少综合征（蜱媒传染病）</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.diseaseType == '斑疹伤寒' }" @tap="!isViewMode && handleDiseaseTypeChange('斑疹伤寒')">斑疹伤寒</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.diseaseType == '流行性出血热' }" @tap="!isViewMode && handleDiseaseTypeChange('流行性出血热')">流行性出血热</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.diseaseType == '其他' }" @tap="!isViewMode && handleDiseaseTypeChange('其他')">其他</view>
        </view>
        <view class="chip-grid nested" v-show="AllCheckInfo.diagnosisResultsEmpDTO.diseaseType == '炭疽'">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.anthraxSubtype == '皮肤炭疽' }" @tap="!isViewMode && showAnthraxSubtype('皮肤炭疽')">皮肤炭疽</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.anthraxSubtype == '肠炭疽' }" @tap="!isViewMode && showAnthraxSubtype('肠炭疽')">肠炭疽</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.anthraxSubtype == '肺炭疽' }" @tap="!isViewMode && showAnthraxSubtype('肺炭疽')">肺炭疽</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.anthraxSubtype == '脑膜炎型炭疽' }" @tap="!isViewMode && showAnthraxSubtype('脑膜炎型炭疽')">脑膜炎型炭疽</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.anthraxSubtype == '败血症型炭疽' }" @tap="!isViewMode && showAnthraxSubtype('败血症型炭疽')">败血症型炭疽</view>
        </view>
        <input v-show="AllCheckInfo.diagnosisResultsEmpDTO.diseaseType == '其他'" :disabled="isViewMode" type="text" placeholder="请输入其他疾病名称" v-model="AllCheckInfo.diagnosisResultsEmpDTO.otherDiseaseName" class="field-input" />
      </view>
      <view class="card">
        <view class="card-title">病例发现途径</view>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.discoveryMethod == 'APP填报' }" @tap="!isViewMode && handleDiscoveryMethodChange('APP填报')">APP填报</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.discoveryMethod == '主动就诊' }" @tap="!isViewMode && handleDiscoveryMethodChange('主动就诊')">主动就诊</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.discoveryMethod == '常规监测发现' }" @tap="!isViewMode && handleDiscoveryMethodChange('常规监测发现')">常规监测发现</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.discoveryMethod == '其他' }" @tap="!isViewMode && handleDiscoveryMethodChange('其他')">其他</view>
        </view>
      </view>
      <view class="card">
        <view class="card-title">提交时间</view>
        <picker mode="date" value="2024-01-01" start="2024-01-01" end="2099-12-31" @change="onDateChange1" v-show="!isViewMode">
          <view class="picker">{{ AllCheckInfo.diagnosisResultsEmpDTO.diagnosisDate || '请选择日期' }}</view>
        </picker>
      </view>
      <view class="card">
        <view class="card-title">疾病转归</view>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.diseaseOutcome == '治愈' }" @tap="!isViewMode && handleDiseaseOutcomeChange('治愈')">治愈</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.diseaseOutcome == '好转' }" @tap="!isViewMode && handleDiseaseOutcomeChange('好转')">好转</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.diseaseOutcome == '未愈' }" @tap="!isViewMode && handleDiseaseOutcomeChange('未愈')">未愈</view>
        </view>
      </view>
        <view class="button-container button-container--single">
          <button class="submit-button" @tap="handleSubmit">下一步</button>
        </view>
      <view class="bottom-space" />
    </scroll-view>
  </view>
  <view class="page-step" v-show="pages == 1 && sonpages == 2">
    <image class="bg" src="/static/other-background.png" mode="aspectFill" />
    <scroll-view scroll-y class="scroll" :show-scrollbar="false">
      <view class="page-header"><text class="title-accent">确诊</text>信息</view>
      <view class="card">
        <view class="card-title">确诊医院名称</view>
        <view class="field">
          <input :disabled="isViewMode" type="text" placeholder="请输入确诊医院名称" v-model="AllCheckInfo.diagnosisResultsEmpDTO.hospitalName" class="field-input" />
        </view>
      </view>
      <view class="card">
        <view class="card-title">就诊/入院时症状和体征</view>
        <view class="field">
          <input :disabled="isViewMode" type="text" placeholder="请输入就诊/入院时症状和体征" v-model="AllCheckInfo.diagnosisResultsEmpDTO.admissionSymptomsAndSigns" class="field-input" />
        </view>
      </view>
      <view class="card">
        <view class="card-title">登记分类</view>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.registrationClassification == '新患者' }" @tap="!isViewMode && handleRegistrationClassificationChange('新患者')">新患者</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.registrationClassification == '复发' }" @tap="!isViewMode && handleRegistrationClassificationChange('复发')">复发</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.registrationClassification == '返回' }" @tap="!isViewMode && handleRegistrationClassificationChange('返回')">返回</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.registrationClassification == '治疗失败' }" @tap="!isViewMode && handleRegistrationClassificationChange('治疗失败')">治疗失败</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisResultsEmpDTO.registrationClassification == '其他' }" @tap="!isViewMode && handleRegistrationClassificationChange('其他')">其他</view>
        </view>
        <input :disabled="isViewMode" type="text" placeholder="请详述" v-model="AllCheckInfo.diagnosisResultsEmpDTO.otherRegistrationDetails" class="field-input mt" />
      </view>
        <view class="button-container">
          <button class="return-button" @tap="handlereturn">上一步</button>
          <button class="submit-button" @tap="handleSubmit">下一步</button>
        </view>
      <view class="bottom-space" />
    </scroll-view>
  </view>
  <view class="page-step" v-show="pages == 2">
    <image class="bg" src="/static/other-background.png" mode="aspectFill" />
    <scroll-view scroll-y class="scroll" :show-scrollbar="false">
      <view class="page-header"><text class="title-accent">基本</text>数据</view>
      <view class="card">
        <view class="card-title">身高、体重、血压</view>
        <view class="field">
          <text class="field-label">身高 (cm)</text>
          <input :disabled="isViewMode" class="field-input" type="text" v-model="AllCheckInfo.diagnosispersonalinfo.height" placeholder="请输入身高" />
        </view>
        <view class="field">
          <text class="field-label">体重 (kg)</text>
          <input class="field-input" type="text" v-model="AllCheckInfo.diagnosispersonalinfo.weight" placeholder="请输入体重" />
        </view>
        <view class="field">
          <text class="field-label">血压 (mmHg)</text>
          <input :disabled="isViewMode" class="field-input" type="text" v-model="AllCheckInfo.diagnosispersonalinfo.bloodPressure" placeholder="请输入血压" />
        </view>

      </view>
      <view class="card">
        <view class="card-title">血糖</view>
        <view class="field">
          <text class="field-label">空腹血糖 (mmol/L)</text>
          <input :disabled="isViewMode" class="field-input" type="text" v-model="AllCheckInfo.diagnosispersonalinfo.bloodSugarFasting" placeholder="请输入空腹血糖" />
        </view>
        <view class="field">
          <text class="field-label">随机血糖 (mmol/L)</text>
          <input :disabled="isViewMode" class="field-input" type="text" v-model="AllCheckInfo.diagnosispersonalinfo.bloodSugarRandom" placeholder="请输入随机血糖" />
        </view>

      </view>
      <view class="card">
        <view class="card-title">血脂</view>
        <view class="field">
          <text class="field-label">总胆固醇</text>
          <input :disabled="isViewMode" class="field-input" type="text" v-model="AllCheckInfo.diagnosispersonalinfo.cholesterolTotal" placeholder="请输入总胆固醇" />
        </view>
        <view class="field">
          <text class="field-label">甘油三酯</text>
          <input :disabled="isViewMode" class="field-input" type="text" v-model="AllCheckInfo.diagnosispersonalinfo.triglycerides" placeholder="请输入甘油三酯" />
        </view>
        <view class="field">
          <text class="field-label">胆固醇酯</text>
          <input :disabled="isViewMode" class="field-input" type="text" v-model="AllCheckInfo.diagnosispersonalinfo.cholesterolLipid" placeholder="请输入胆固醇酯" />
        </view>
        <view class="field">
          <text class="field-label">高密度脂蛋白</text>
          <input :disabled="isViewMode" class="field-input" type="text" v-model="AllCheckInfo.diagnosispersonalinfo.hdl" placeholder="请输入高密度脂蛋白" />
        </view>
        <view class="field">
          <text class="field-label">低密度脂蛋白</text>
          <input :disabled="isViewMode" class="field-input" type="text" v-model="AllCheckInfo.diagnosispersonalinfo.ldl" placeholder="请输入低密度脂蛋白" />
        </view>

      </view>
      <view class="card">
        <view class="card-title">血氧饱和度</view>
        <view class="field">
          <text class="field-label">血氧饱和度 (%)</text>
          <input :disabled="isViewMode" class="field-input" type="text" v-model="AllCheckInfo.diagnosispersonalinfo.bloodOxygen" placeholder="请输入血氧饱和度" />
        </view>
      </view>
        <view class="button-container">
          <button class="return-button" @tap="handlereturn">上一步</button>
          <button class="submit-button" @tap="handleSubmit">下一步</button>
        </view>
      <view class="bottom-space" />
    </scroll-view>
  </view>
  <view class="page-step" v-show="pages == 3">
    <image class="bg" src="/static/other-background.png" mode="aspectFill" />
    <scroll-view scroll-y class="scroll" :show-scrollbar="false">
      <view class="page-header"><text class="title-accent">全身</text>症状</view>
      <view class="card">
        <view class="card-title">发热情况</view>
        <text class="field-label">是否发热</text>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasFever }" @tap="!isViewMode && hanleGeneralSymprion('hasFever')">是</view>
          <view class="chip" :class="{ on: !AllCheckInfo.diagnosisgeneralsymptoms.hasFever }" @tap="!isViewMode && hanleGeneralSymprion('hasFever')">否</view>
        </view>
        <view class="field-row-inline">
          <view class="field flex1">
            <text class="field-label">最高温度</text>
            <input :disabled="isViewMode" class="field-input" type="text" v-model="AllCheckInfo.diagnosisgeneralsymptoms.highestTemperature" placeholder="请输入温度" />
          </view>
          <view class="field flex1">
            <text class="field-label">持续时间（小时）</text>
            <input :disabled="isViewMode" class="field-input" type="text" v-model="AllCheckInfo.diagnosisgeneralsymptoms.feverDuration" placeholder="请输入小时数" />
          </view>
        </view>
      </view>
      <view class="card">
        <view class="card-title">常见症状</view>
        <text class="card-sub">可多选</text>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasChills }" @tap="!isViewMode && handleCheckboxChange('hasChills')">寒战</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasSweating }" @tap="!isViewMode && handleCheckboxChange('hasSweating')">大汗</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasFatigue }" @tap="!isViewMode && handleCheckboxChange('hasFatigue')">乏力</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasHeadache }" @tap="!isViewMode && handleCheckboxChange('hasHeadache')">头痛</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasMusclePain }" @tap="!isViewMode && handleCheckboxChange('hasMusclePain')">肌肉酸痛</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasJointPain }" @tap="!isViewMode && handleCheckboxChange('hasJointPain')">关节酸痛</view>
        </view>
      </view>
      <view class="card">
        <view class="card-title">皮肤、疹、疮、淋巴结</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasLymphNodeSwelling }" @tap="!isViewMode && handleCheckboxChange('hasLymphNodeSwelling')">淋巴结肿大</view>
        <view class="chip-grid nested" v-if="AllCheckInfo.diagnosisgeneralsymptoms.hasLymphNodeSwelling == true">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasLymphNodeSwellingGroin }" @tap="!isViewMode && handleCheckboxChange('hasLymphNodeSwellingGroin')">腹股沟</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasLymphNodeSwellingArmpit }" @tap="!isViewMode && handleCheckboxChange('hasLymphNodeSwellingArmpit')">腋下</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasLymphNodeSwellingSubclavian }" @tap="!isViewMode && handleCheckboxChange('hasLymphNodeSwellingSubclavian')">锁骨下</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasLymphNodeSwellingNeck }" @tap="!isViewMode && handleCheckboxChange('hasLymphNodeSwellingNeck')">颈部</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasLymphNodeSwellingElbow }" @tap="!isViewMode && handleCheckboxChange('hasLymphNodeSwellingElbow')">肘部</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasLymphNodeSwellingPopliteal }" @tap="!isViewMode && handleCheckboxChange('hasLymphNodeSwellingPopliteal')">腘窝</view>
        </view>
        <view class="chip-grid">
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasLipCyanosis }" @tap="!isViewMode && handleCheckboxChange('hasLipCyanosis')">口唇、颜面、四肢及全身皮肤发绀</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasSubcutaneousAndMucosalBleedingSpots }" @tap="!isViewMode && handleCheckboxChange('hasSubcutaneousAndMucosalBleedingSpots')">皮下及粘膜出血或出血点</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasPainfulRedRash }" @tap="!isViewMode && handleCheckboxChange('hasPainfulRedRash')">皮肤可见剧痛的红色丘疹</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasPainfulRedRash }" @tap="!isViewMode && handleCheckboxChange('hasPainfulRedRash')">皮疹：起初压之褪色，后期压之不退，多见于四肢、躯干、面部一般没有</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasBloodBlisters }" @tap="!isViewMode && handleCheckboxChange('hasBloodBlisters')">皮肤可见周边呈灰黑色、基底坚硬的血性水泡</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasSkinUlcer }" @tap="!isViewMode && handleCheckboxChange('hasSkinUlcer')">皮肤出现呈灰黑色创面的溃疡</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasCongestiveOrPetechialRash }" @tap="!isViewMode && handleCheckboxChange('hasCongestiveOrPetechialRash')">皮疹：充血性或点状出血疹</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.hasItchyRash }" @tap="!isViewMode && handleCheckboxChange('hasItchyRash')">瘙痒性丘疹/水疱</view>
        </view>        <view class="chip-grid nested" v-if="AllCheckInfo.diagnosisgeneralsymptoms.hasItchyRash == true">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.itchyRashOnFingers }" @tap="!isViewMode && handleCheckboxChange('itchyRashOnFingers')">手指</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.itchyRashOnBackOfHands }" @tap="!isViewMode && handleCheckboxChange('itchyRashOnBackOfHands')">手背</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.itchyRashOnUpperLimbs }" @tap="!isViewMode && handleCheckboxChange('itchyRashOnUpperLimbs')">上肢</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.itchyRashOnLowerLimbs }" @tap="!isViewMode && handleCheckboxChange('itchyRashOnLowerLimbs')">下肢</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.itchyRashOnFeet }" @tap="!isViewMode && handleCheckboxChange('itchyRashOnFeet')">足背</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.itchyRashOnFace }" @tap="!isViewMode && handleCheckboxChange('itchyRashOnFace')">面部</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.itchyRashOnOther }" @tap="!isViewMode && handleCheckboxChange('itchyRashOnOther')">其他</view>
        </view>
      </view>
      <view class="card">
        <view class="card-title">失水情况</view>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.dehydrationSeverity === '轻度' }" @tap="!isViewMode && showRadio('轻度')">轻度</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.dehydrationSeverity === '中度' }" @tap="!isViewMode && showRadio('中度')">中度</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisgeneralsymptoms.dehydrationSeverity === '重度' }" @tap="!isViewMode && showRadio('重度')">重度</view>
        </view>
      </view>
        <view class="button-container">
          <button class="return-button" @tap="handlereturn">上一步</button>
          <button class="submit-button" @tap="handleSubmit">下一步</button>
        </view>
      <view class="bottom-space" />
    </scroll-view>
  </view>
  <view class="page-step" v-show="pages == 4">
    <image class="bg" src="/static/other-background.png" mode="aspectFill" />
    <scroll-view scroll-y class="scroll" :show-scrollbar="false">
      <view class="page-header"><text class="title-accent">呼吸</text>系统症状</view>
      <view class="card">
        <view class="card-title">请选择症状</view>
        <text class="card-sub">可多选</text>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisrespiratorysymptoms.hasCough }" @tap="!isViewMode && handleRespiratorysymptomsBox('hasCough')">咳嗽</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisrespiratorysymptoms.hasSputum }" @tap="!isViewMode && handleRespiratorysymptomsBox('hasSputum')">咳痰</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisrespiratorysymptoms.hasBloodySputum }" @tap="!isViewMode && handleRespiratorysymptomsBox('hasBloodySputum')">血痰</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisrespiratorysymptoms.hasHemoptysis }" @tap="!isViewMode && handleRespiratorysymptomsBox('hasHemoptysis')">咳血</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisrespiratorysymptoms.hasSoreThroat }" @tap="!isViewMode && handleRespiratorysymptomsBox('hasSoreThroat')">咽痛</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisrespiratorysymptoms.hasDryThroat }" @tap="!isViewMode && handleRespiratorysymptomsBox('hasDryThroat')">咽干</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisrespiratorysymptoms.hasNasalCongestion }" @tap="!isViewMode && handleRespiratorysymptomsBox('hasNasalCongestion')">鼻塞</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisrespiratorysymptoms.hasRunnyNose }" @tap="!isViewMode && handleRespiratorysymptomsBox('hasRunnyNose')">流涕</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisrespiratorysymptoms.hasChestTightness }" @tap="!isViewMode && handleRespiratorysymptomsBox('hasChestTightness')">胸闷</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisrespiratorysymptoms.hasShortnessOfBreath }" @tap="!isViewMode && handleRespiratorysymptomsBox('hasShortnessOfBreath')">气促</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisrespiratorysymptoms.hasBreathingDifficulty }" @tap="!isViewMode && handleRespiratorysymptomsBox('hasBreathingDifficulty')">呼吸困难</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisrespiratorysymptoms.hasHoarseness }" @tap="!isViewMode && handleRespiratorysymptomsBox('hasHoarseness')">声音嘶哑</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosisrespiratorysymptoms.hasBronchitis }" @tap="!isViewMode && handleRespiratorysymptomsBox('hasBronchitis')">支气管炎/支气管肺炎</view>
        </view>
      </view>
        <view class="button-container">
          <button class="return-button" @tap="handlereturn">上一步</button>
          <button class="submit-button" @tap="handleSubmit">下一步</button>
        </view>
      <view class="bottom-space" />
    </scroll-view>
  </view>
  <view class="page-step" v-show="pages == 5">
    <image class="bg" src="/static/other-background.png" mode="aspectFill" />
    <scroll-view scroll-y class="scroll" :show-scrollbar="false">
      <view class="page-header"><text class="title-accent">消化</text>系统症状</view>
      <view class="card">
        <view class="card-title">腹泻情况</view>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.diarrheaFrequencyGEThreeTimesPerDay }" @tap="!isViewMode && handleDigestFrequency('diarrheaFrequencyGEThreeTimesPerDay')">少于3次/天</view>
          <view class="chip" :class="{ on: !AllCheckInfo.diagnosisdigestivesymptoms.diarrheaFrequencyGEThreeTimesPerDay }" @tap="!isViewMode && handleDigestFrequency('diarrheaFrequencyGEThreeTimesPerDay')">大于等于3次/天</view>
        </view>
        <input class="field-input mt" type="number" placeholder="请输入具体次数" v-model="AllCheckInfo.diagnosisdigestivesymptoms.diarrheaTimesPerDay" v-show="AllCheckInfo.diagnosisdigestivesymptoms.diarrheaFrequencyGEThreeTimesPerDay"/>
        <text class="field-label mt">腹泻方式</text>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.diarrheaMode == '里急后重' }" @tap="!isViewMode && showDigestRadio('里急后重')">里急后重</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.diarrheaMode == '失禁' }" @tap="!isViewMode && showDigestRadio('失禁')">失禁</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.diarrheaMode == '通畅' }" @tap="!isViewMode && showDigestRadio('通畅')">通畅</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.diarrheaMode == '绞痛' }" @tap="!isViewMode && showDigestRadio('绞痛')">绞痛</view>
        </view>
      </view>
      <view class="card">
        <view class="card-title">粪便情况</view>
        <text class="field-label">粪便量</text>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.stoolAmount == '多' }" @tap="!isViewMode && showStoolAmount('多')">多</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.stoolAmount == '少' }" @tap="!isViewMode && showStoolAmount('少')">少</view>
        </view>
        <text class="field-label mt">粪便气味</text>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.stoolOdor == '恶臭' }" @tap="!isViewMode && showStoolOdor('恶臭')">恶臭</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.stoolOdor == '无恶臭' }" @tap="!isViewMode && showStoolOdor('无恶臭')">无恶臭</view>
        </view>
        <view class="chip-grid">
          <view class="chip chip-wide" :class="{ on: temporaryChoice1 }" @tap="!isViewMode && !temporaryChoice2 && (temporaryChoice1 = !temporaryChoice1)">粪便性状1</view>
          <view class="chip chip-wide" :class="{ on: temporaryChoice2 }" @tap="!isViewMode && !temporaryChoice1 && (temporaryChoice2 = !temporaryChoice2)">粪便性状2</view>
        </view>
        <view class="chip-grid nested" v-show="temporaryChoice1">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.stoolType1Detail == '水样' }" @tap="!isViewMode && showDigestRadioBystool1('水样')">水样</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.stoolType1Detail == '米泔样' }" @tap="!isViewMode && showDigestRadioBystool1('米泔样')">米泔样</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.stoolType1Detail == '洗肉水样' }" @tap="!isViewMode && showDigestRadioBystool1('洗肉水样')">洗肉水样</view>
        </view>
        <view class="chip-grid nested" v-show="temporaryChoice2">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.stoolType2Detail == '黑便' }" @tap="!isViewMode && showDigestRadioBystool2('黑便')">黑便</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.stoolType2Detail == '脓血' }" @tap="!isViewMode && showDigestRadioBystool2('脓血')">脓血</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.stoolType2Detail == '大块粘膜' }" @tap="!isViewMode && showDigestRadioBystool2('大块粘膜')">大块粘膜</view>
        </view>
      </view>
      <view class="card">
        <view class="card-title">呕吐情况</view>
        <text class="field-label">是否呕吐</text>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.hasVomiting }" @tap="!isViewMode && handleDigestFrequency('hasVomiting')">是</view>
          <view class="chip" :class="{ on: !AllCheckInfo.diagnosisdigestivesymptoms.hasVomiting }" @tap="!isViewMode && handleDigestFrequency('hasVomiting')">否</view>
        </view>
        <text class="field-label mt">呕吐方式</text>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.vomitingMode == '喷射状' }" @tap="!isViewMode && showDigestVomitingMode('喷射状')">喷射状</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.vomitingMode == '先泻后吐' }" @tap="!isViewMode && showDigestVomitingMode('先泻后吐')">先泻后吐</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.vomitingMode == '先吐后泻' }" @tap="!isViewMode && showDigestVomitingMode('先吐后泻')">先吐后泻</view>
        </view>
      </view>
      <view class="card">
        <view class="card-title">其他消化系统症状</view>
        <text class="card-sub">可多选</text>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.hasNausea }" @tap="!isViewMode && handleDigestFrequency('hasNausea')">恶心</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.hasAbdominalDistension }" @tap="!isViewMode && handleDigestFrequency('hasAbdominalDistension')">腹胀</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.hasUpperAbdominalDiscomfort }" @tap="!isViewMode && handleDigestFrequency('hasUpperAbdominalDiscomfort')">上腹部不适</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.hasAbdominalPain }" @tap="!isViewMode && handleDigestFrequency('hasAbdominalPain')">腹痛</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.hasBorborygmus }" @tap="!isViewMode && handleDigestFrequency('hasBorborygmus')">腹鸣</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.hasConstipation }" @tap="!isViewMode && handleDigestFrequency('hasConstipation')">便秘</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.hasOliguriaOrAnuria }" @tap="!isViewMode && handleDigestFrequency('hasOliguriaOrAnuria')">少尿或无尿</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisdigestivesymptoms.hasAppetiteLoss }" @tap="!isViewMode && handleDigestFrequency('hasAppetiteLoss')">食欲减退</view>
        </view>
      </view>
        <view class="button-container">
          <button class="return-button" @tap="handlereturn">上一步</button>
          <button class="submit-button" @tap="handleSubmit">下一步</button>
        </view>
      <view class="bottom-space" />
    </scroll-view>
  </view>
  <view class="page-step" v-show="pages == 6">
    <image class="bg" src="/static/other-background.png" mode="aspectFill" />
    <scroll-view scroll-y class="scroll" :show-scrollbar="false">
      <view class="page-header"><text class="title-accent">循环</text>系统症状</view>
      <view class="card">
        <view class="card-title">请选择症状</view>
        <text class="card-sub">可多选</text>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscirculatorysymptoms.hasArrhythmia }" @tap="!isViewMode && handleCirculatorysymptoms('hasArrhythmia')">心律不齐</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscirculatorysymptoms.hasChestPain }" @tap="!isViewMode && handleCirculatorysymptoms('hasChestPain')">胸痛</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscirculatorysymptoms.hasRapidPulse }" @tap="!isViewMode && handleCirculatorysymptoms('hasRapidPulse')">脉搏细速</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscirculatorysymptoms.hasPalpitation }" @tap="!isViewMode && handleCirculatorysymptoms('hasPalpitation')">心悸</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscirculatorysymptoms.hasLowBloodPressure }" @tap="!isViewMode && handleCirculatorysymptoms('hasLowBloodPressure')">低血压</view>
        </view>
      </view>
        <view class="button-container">
          <button class="return-button" @tap="handlereturn">上一步</button>
          <button class="submit-button" @tap="handleSubmit">下一步</button>
        </view>
      <view class="bottom-space" />
    </scroll-view>
  </view>
  <view class="page-step" v-show="pages == 7">
    <image class="bg" src="/static/other-background.png" mode="aspectFill" />
    <scroll-view scroll-y class="scroll" :show-scrollbar="false">
      <view class="page-header"><text class="title-accent">神经</text>系统症状</view>
      <view class="card">
        <view class="card-title">请选择症状</view>
        <text class="card-sub">可多选</text>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisneurologicalsymptoms.hasDelirium }" @tap="!isViewMode && handleNeurologicalSymptoms('hasDelirium')">儋妄</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisneurologicalsymptoms.hasConvulsion }" @tap="!isViewMode && handleNeurologicalSymptoms('hasConvulsion')">抽搐</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisneurologicalsymptoms.hasDrowsiness }" @tap="!isViewMode && handleNeurologicalSymptoms('hasDrowsiness')">昏睡</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisneurologicalsymptoms.hasComa }" @tap="!isViewMode && handleNeurologicalSymptoms('hasComa')">昏迷</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisneurologicalsymptoms.hasStiffNeck }" @tap="!isViewMode && handleNeurologicalSymptoms('hasStiffNeck')">颈部强直</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisneurologicalsymptoms.hasRestlessness }" @tap="!isViewMode && handleNeurologicalSymptoms('hasRestlessness')">烦躁不安</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosisneurologicalsymptoms.hasMuscleParalysisNeckShoulder }" @tap="!isViewMode && handleNeurologicalSymptoms('hasMuscleParalysisNeckShoulder')">颈肩部肌肉迟缓性瘫痪</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosisneurologicalsymptoms.hasMuscleParalysisLimbs }" @tap="!isViewMode && handleNeurologicalSymptoms('hasMuscleParalysisLimbs')">肢体肌肉迟缓性瘫痪</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisneurologicalsymptoms.hasSwallowingDifficulty }" @tap="!isViewMode && handleNeurologicalSymptoms('hasSwallowingDifficulty')">吞咽困难</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisneurologicalsymptoms.hasSpeechDisorder }" @tap="!isViewMode && handleNeurologicalSymptoms('hasSpeechDisorder')">语言障碍</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisneurologicalsymptoms.hasConsciousnessDisorder }" @tap="!isViewMode && handleNeurologicalSymptoms('hasConsciousnessDisorder')">意识障碍或惊厥</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisneurologicalsymptoms.hasDizziness }" @tap="!isViewMode && handleNeurologicalSymptoms('hasDizziness')">头晕</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisneurologicalsymptoms.hasTinnitus }" @tap="!isViewMode && handleNeurologicalSymptoms('hasTinnitus')">耳鸣</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisneurologicalsymptoms.hasHearingLoss }" @tap="!isViewMode && handleNeurologicalSymptoms('hasHearingLoss')">听力减退</view>
        </view>
      </view>
        <view class="button-container">
          <button class="return-button" @tap="handlereturn">上一步</button>
          <button class="submit-button" @tap="handleSubmit">下一步</button>
        </view>
      <view class="bottom-space" />
    </scroll-view>
  </view>
  <view class="page-step" v-show="pages == 8">
    <image class="bg" src="/static/other-background.png" mode="aspectFill" />
    <scroll-view scroll-y class="scroll" :show-scrollbar="false">
      <view class="page-header"><text class="title-accent">局部</text>症状</view>
      <view class="card">
        <view class="card-title">请选择症状</view>
        <text class="card-sub">可多选</text>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosislocalsymptoms.hasConjunctivitis }" @tap="!isViewMode && handleLocalSymptoms('hasConjunctivitis')">结膜炎</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosislocalsymptoms.hasConjunctivalCongestion }" @tap="!isViewMode && handleLocalSymptoms('hasConjunctivalCongestion')">眼结膜充血、肿胀疼痛</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosislocalsymptoms.hasEyelidEdema }" @tap="!isViewMode && handleLocalSymptoms('hasEyelidEdema')">上下眼睑水肿</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosislocalsymptoms.hasSmellTasteLoss }" @tap="!isViewMode && handleLocalSymptoms('hasSmellTasteLoss')">嗅觉味觉减退或丧失</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosislocalsymptoms.hasFacialRednessOrPallor }" @tap="!isViewMode && handleLocalSymptoms('hasFacialRednessOrPallor')">颜面潮红或苍白</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosislocalsymptoms.hasRednessNeckChest }" @tap="!isViewMode && handleLocalSymptoms('hasRednessNeckChest')">潮红（颈部、胸部）</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosislocalsymptoms.hasLipCyanosis }" @tap="!isViewMode && handleLocalSymptoms('hasLipCyanosis')">口唇发绀</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosislocalsymptoms.hasSkinPallorOrBruising }" @tap="!isViewMode && handleLocalSymptoms('hasSkinPallorOrBruising')">皮肤苍白或带青紫</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosislocalsymptoms.hasGumBleeding }" @tap="!isViewMode && handleLocalSymptoms('hasGumBleeding')">牙龈出血</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosislocalsymptoms.hasKidneyPain }" @tap="!isViewMode && handleLocalSymptoms('hasKidneyPain')">肾区疼痛</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosislocalsymptoms.hasLowBackPain }" @tap="!isViewMode && handleLocalSymptoms('hasLowBackPain')">腰痛</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosislocalsymptoms.hasOrbitalPain }" @tap="!isViewMode && handleLocalSymptoms('hasOrbitalPain')">眼眶痛</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosislocalsymptoms.hasMucosalBleedingInOralOrNasal }" @tap="!isViewMode && handleLocalSymptoms('hasMucosalBleedingInOralOrNasal')">口腔、鼻腔等处黏膜有出血点</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosislocalsymptoms.hasSkinBleedingPoints }" @tap="!isViewMode && handleLocalSymptoms('hasSkinBleedingPoints')">腋下/上臂/胸部或其他部位有皮肤出血点</view>
        </view>
      </view>
        <view class="button-container">
          <button class="return-button" @tap="handlereturn">上一步</button>
          <button class="submit-button" @tap="handleSubmit">下一步</button>
        </view>
      <view class="bottom-space" />
    </scroll-view>
  </view>
  <view class="page-step" v-show="pages == 9">
    <image class="bg" src="/static/other-background.png" mode="aspectFill" />
    <scroll-view scroll-y class="scroll" :show-scrollbar="false">
      <view class="page-header"><text class="title-accent">其他</text>症状</view>
      <view class="card">
        <view class="card-title">发病情况</view>
        <text class="card-sub">可多选</text>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisothersymptoms.hasSuddenOnset }" @tap="!isViewMode && handleOtherSymptoms('hasSuddenOnset')">突然发病</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisothersymptoms.hasRapidProgress }" @tap="!isViewMode && handleOtherSymptoms('hasRapidProgress')">病情进展迅速</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisothersymptoms.hasPeriodicAttack }" @tap="!isViewMode && handleOtherSymptoms('hasPeriodicAttack')">周期性发作</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisothersymptoms.hasForcedPosture }" @tap="!isViewMode && handleOtherSymptoms('hasForcedPosture')">强迫体位</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain }" @tap="!isViewMode && handleOtherSymptoms('hasCalfMusclePain')">腓肠肌疼痛</view>
        </view>
      </view>
        <view class="button-container">
          <button class="return-button" @tap="handlereturn">上一步</button>
          <button class="submit-button" @tap="handleSubmit">下一步</button>
        </view>
      <view class="bottom-space" />
    </scroll-view>
  </view>
  <view class="page-step" v-show="pages == 10">
    <image class="bg" src="/static/other-background.png" mode="aspectFill" />
    <scroll-view scroll-y class="scroll" :show-scrollbar="false">
      <view class="page-header"><text class="title-accent">并发</text>症</view>
      <view class="card">
        <view class="card-title">并发症情况</view>
        <text class="field-label">有无并发症</text>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasComplications }" @tap="!isViewMode && handleComplications('hasComplications')">有</view>
          <view class="chip" :class="{ on: !AllCheckInfo.diagnosiscomplications.hasComplications }" @tap="!isViewMode && handleComplications('hasComplications')">无</view>
        </view>
      </view>
      <view class="card">
        <view class="card-title">选择并发症</view>
        <text class="card-sub">可多选</text>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasViralPneumonia }" @tap="!isViewMode && handleComplications('hasViralPneumonia')">病毒性肺炎</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasBacterialPneumonia }" @tap="!isViewMode && handleComplications('hasBacterialPneumonia')">细菌性肺炎</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasFungalPneumonia }" @tap="!isViewMode && handleComplications('hasFungalPneumonia')">真菌性肺炎</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasRespiratoryFailure }" @tap="!isViewMode && handleComplications('hasRespiratoryFailure')">呼吸衰竭</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosiscomplications.hasArds }" @tap="!isViewMode && handleComplications('hasArds')">急性呼吸窘迫综合征(ARDS)</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosiscomplications.hasDic }" @tap="!isViewMode && handleComplications('hasDic')">弥散性血管内凝血（DIC）</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasToxicShock }" @tap="!isViewMode && handleComplications('hasToxicShock')">感染中毒性休克</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasMeningitis }" @tap="!isViewMode && handleComplications('hasMeningitis')">脑膜炎</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasEncephalitis }" @tap="!isViewMode && handleComplications('hasEncephalitis')">脑炎</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasBacteremia }" @tap="!isViewMode && handleComplications('hasBacteremia')">菌血症</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasMyocarditis }" @tap="!isViewMode && handleComplications('hasMyocarditis')">心肌炎</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosiscomplications.hasAcuteLungInjury }" @tap="!isViewMode && handleComplications('hasAcuteLungInjury')">急性肺损伤/ARDS</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasAcuteKidneyInjury }" @tap="!isViewMode && handleComplications('hasAcuteKidneyInjury')">急性肾损伤</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasEpilepsy }" @tap="!isViewMode && handleComplications('hasEpilepsy')">癫痫</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasToxicHepatitis }" @tap="!isViewMode && handleComplications('hasToxicHepatitis')">中毒性肝炎</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasFluidOverload }" @tap="!isViewMode && handleComplications('hasFluidOverload')">输液过量</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasSuperinfection }" @tap="!isViewMode && handleComplications('hasSuperinfection')">二重感染</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasHypoxemia }" @tap="!isViewMode && handleComplications('hasHypoxemia')">低血氧症</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasSepticShock }" @tap="!isViewMode && handleComplications('hasSepticShock')">脓毒症休克</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasMetabolicAcidosis }" @tap="!isViewMode && handleComplications('hasMetabolicAcidosis')">代谢性酸中毒</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosiscomplications.hasElectrolyteImbalance }" @tap="!isViewMode && handleComplications('hasElectrolyteImbalance')">电解质及酸碱失衡</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasHemolysis }" @tap="!isViewMode && handleComplications('hasHemolysis')">急性血管内溶血</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasGastrointestinalBleeding }" @tap="!isViewMode && handleComplications('hasGastrointestinalBleeding')">消化道出血</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosiscomplications.hasCoagulationDisorder }" @tap="!isViewMode && handleComplications('hasCoagulationDisorder')">出凝血功能障碍</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosiscomplications.hasMultipleOrganFailure }" @tap="!isViewMode && handleComplications('hasMultipleOrganFailure')">多器官功能衰竭</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasBronchitis }" @tap="!isViewMode && handleComplications('hasBronchitis')">支气管炎</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasBronchopneumonia }" @tap="!isViewMode && handleComplications('hasBronchopneumonia')">支气管肺炎</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasAnemia }" @tap="!isViewMode && handleComplications('hasAnemia')">贫血</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasHepatosplenomegaly }" @tap="!isViewMode && handleComplications('hasHepatosplenomegaly')">肝脾肿大</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasBrainEdema }" @tap="!isViewMode && handleComplications('hasBrainEdema')">脑水肿</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosiscomplications.hasHus }" @tap="!isViewMode && handleComplications('hasHus')">溶血性尿毒综合征</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasHematemesis }" @tap="!isViewMode && handleComplications('hasHematemesis')">呕血</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasMelena }" @tap="!isViewMode && handleComplications('hasMelena')">便血</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasHemoptysis }" @tap="!isViewMode && handleComplications('hasHemoptysis')">咯血</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasAbdominalHemorrhage }" @tap="!isViewMode && handleComplications('hasAbdominalHemorrhage')">腹腔出血</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasNasalBleeding }" @tap="!isViewMode && handleComplications('hasNasalBleeding')">鼻出血</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasVaginalBleeding }" @tap="!isViewMode && handleComplications('hasVaginalBleeding')">阴道出血</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasHypertensiveEncephalopathy }" @tap="!isViewMode && handleComplications('hasHypertensiveEncephalopathy')">高血压脑病</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasIntracranialHemorrhage }" @tap="!isViewMode && handleComplications('hasIntracranialHemorrhage')">颅内出血</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasPulmonaryEdema }" @tap="!isViewMode && handleComplications('hasPulmonaryEdema')">肺水肿</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasSecondaryInfection }" @tap="!isViewMode && handleComplications('hasSecondaryInfection')">继发性感染</view>
          <view class="chip chip-wide" :class="{ on: AllCheckInfo.diagnosiscomplications.hasSpontaneousRenalRupture }" @tap="!isViewMode && handleComplications('hasSpontaneousRenalRupture')">自发性肾破裂</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasMyocardialDamage }" @tap="!isViewMode && handleComplications('hasMyocardialDamage')">心肌损害</view>
          <view v-show="!isViewMode" class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.hasLiverDamage }" @tap="!isViewMode && handleComplications('hasLiverDamage')">肝损害</view>
        </view>
      </view>
      <view class="card">
        <view class="card-title">其他并发症</view>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosiscomplications.otherComplications }" @tap="!isViewMode && handleComplications('otherComplications')">其他</view>
        </view>
        <input :disabled="isViewMode" type="text" placeholder="请填写其他症状" class="field-input mt" v-model="AllCheckInfo.diagnosiscomplications.otherComplicationsName"/>
      </view>
        <view class="button-container">
          <button class="return-button" @tap="handlereturn">上一步</button>
          <button class="submit-button" @tap="handleSubmit">下一步</button>
        </view>
      <view class="bottom-space" />
    </scroll-view>
  </view>
  <view class="page-step" v-show="pages == 11">
    <image class="bg" src="/static/other-background.png" mode="aspectFill" />
    <scroll-view scroll-y class="scroll" :show-scrollbar="false">
      <view class="page-header"><text class="title-accent">检测</text>报告</view>
      <view class="card">
        <view class="card-title">已做检查（可多选）</view>
        <view class="chip-grid">
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisexaminations.hasBloodTest }" @tap="handleDiagnosisexaminations('hasBloodTest')">血常规</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisexaminations.hasImaging }" @tap="handleDiagnosisexaminations('hasImaging')">影像学检查</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisexaminations.hasUrineTest }" @tap="handleDiagnosisexaminations('hasUrineTest')">尿常规</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisexaminations.hasStoolTest }" @tap="handleDiagnosisexaminations('hasStoolTest')">粪便常规</view>
          <view class="chip" :class="{ on: AllCheckInfo.diagnosisexaminations.hasUltrasound }" @tap="handleDiagnosisexaminations('hasUltrasound')">B超</view>
        </view>
      </view>
      <view class="card">
        <view class="card-title">影像学检查报告</view>
        <text class="card-sub">胸部 X 线 / CT / 核磁共振等</text>
        <report-upload :files="formData.imageReport" :disabled="isViewMode" @update:files="setReportFiles('imageReport', $event)" />
      </view>
      <view class="card">
        <view class="card-title">B超报告</view>
        <report-upload :files="formData.ultrasoundReport" :disabled="isViewMode" @update:files="setReportFiles('ultrasoundReport', $event)" />
      </view>
      <view class="card">
        <view class="card-title">血液检查报告</view>
        <text class="card-sub">血常规、C反应蛋白、血沉等</text>
        <report-upload :files="formData.bloodReport" :disabled="isViewMode" @update:files="setReportFiles('bloodReport', $event)" />
      </view>
      <view class="card">
        <view class="card-title">粪便常规检查报告</view>
        <report-upload :files="formData.stoolReport" :disabled="isViewMode" @update:files="setReportFiles('stoolReport', $event)" />
      </view>
      <view class="card">
        <view class="card-title">尿常规检查报告</view>
        <report-upload :files="formData.urineReport" :disabled="isViewMode" @update:files="setReportFiles('urineReport', $event)" />
      </view>
      <view class="card">
        <view class="card-title">病案首页</view>
        <report-upload :files="formData.medicalRecord" :disabled="isViewMode" @update:files="setReportFiles('medicalRecord', $event)" />
      </view>
      <view class="button-container">
        <button class="return-button" @tap="handlereturn">上一步</button>
        <button class="submit-button" @tap="handleSubmit">提交</button>
      </view>
      <view class="bottom-space" />
    </scroll-view>
  </view>
    <uni-custom-tabbar class="tabbar" />
    <view class="submit-loading-mask" v-if="submitting">
      <view class="submit-loading-card">
        <view class="submit-loading-spinner"></view>
        <text class="submit-loading-title">正在提交</text>
        <text class="submit-loading-desc">报告上传与诊断信息提交中，请稍候…</text>
      </view>
    </view>
  </view>
</template>
<style>
.submit-loading-mask {
  position: fixed;
  left: 0;
  top: 0;
  right: 0;
  bottom: 0;
  z-index: 9999;
  background: rgba(15, 23, 42, 0.45);
  display: flex;
  align-items: center;
  justify-content: center;
}
.submit-loading-card {
  width: 72%;
  max-width: 280px;
  padding: 28px 20px 24px;
  border-radius: 14px;
  background: #ffffff;
  display: flex;
  flex-direction: column;
  align-items: center;
  box-shadow: 0 12px 40px rgba(15, 23, 42, 0.18);
}
.submit-loading-spinner {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  border: 3px solid #dbe4f5;
  border-top-color: #285AC8;
  animation: submit-spin 0.8s linear infinite;
  margin-bottom: 14px;
}
.submit-loading-title {
  font-size: 16px;
  font-weight: 600;
  color: #0f172a;
  margin-bottom: 6px;
}
.submit-loading-desc {
  font-size: 13px;
  color: #64748b;
  text-align: center;
  line-height: 1.45;
}
@keyframes submit-spin {
  to { transform: rotate(360deg); }
}
.page {
  position: relative;
  width: 100%;
  min-height: 100vh;
  background: #eef3fb;
}
.page-step {
  position: relative;
  width: 100%;
  min-height: 100vh;
}
.bg {
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 36vh;
  z-index: 0;
}
.scroll {
  position: relative;
  z-index: 1;
  height: 100vh;
  box-sizing: border-box;
  padding: 11vh 16px 0;
}
.page-header {
  font-size: 22px;
  font-weight: 700;
  color: #1f2a44;
  margin-bottom: 14px;
  padding-left: 4px;
}
.title-accent {
  position: relative;
}
.title-accent::after {
  content: "";
  position: absolute;
  bottom: -4px;
  left: 0;
  width: 100%;
  height: 5px;
  background-color: #285ac8;
}
.card {
  background: #fff;
  border-radius: 14px;
  padding: 16px 14px;
  margin-bottom: 12px;
  box-shadow: 0 4px 16px rgba(40, 90, 200, 0.06);
}
.card-title {
  font-size: 15px;
  font-weight: 600;
  color: #1f2a44;
  padding-left: 8px;
  border-left: 3px solid #285ac8;
  line-height: 1.2;
  margin-bottom: 10px;
}
.card-sub {
  display: block;
  font-size: 12px;
  color: #94a3b8;
  margin-bottom: 10px;
  padding-left: 11px;
}
.chip-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}
.chip-grid.nested {
  margin-top: 8px;
  padding-left: 4px;
}
.chip {
  padding: 8px 12px;
  border-radius: 999px;
  background: #f3f6fb;
  color: #4a5568;
  font-size: 13px;
  border: 1px solid transparent;
  line-height: 1.35;
}
.chip-wide {
  width: 100%;
  border-radius: 10px;
  text-align: left;
}
.chip.on {
  background: rgba(40, 90, 200, 0.1);
  color: #285ac8;
  border-color: rgba(40, 90, 200, 0.35);
  font-weight: 600;
}
.field {
  margin-bottom: 10px;
}
.field-label {
  display: block;
  font-size: 13px;
  color: #4a5568;
  margin-bottom: 6px;
}
.field-input {
  width: 100%;
  height: 42px;
  padding: 0 12px;
  box-sizing: border-box;
  border-radius: 10px;
  background: #f5f7fb;
  font-size: 15px;
  color: #1a202c;
}
.picker {
  height: 42px;
  line-height: 42px;
  padding: 0 12px;
  border-radius: 10px;
  background: #f5f7fb;
  font-size: 15px;
  color: #1a202c;
}
.field-row-inline {
  display: flex;
  gap: 10px;
  margin-top: 10px;
}
.flex1 {
  flex: 1;
  min-width: 0;
}
.mt {
  margin-top: 10px;
}
.button-container {
  display: flex;
  gap: 12px;
  margin-top: 8px;
  padding: 0 2px;
}
.button-container--single .submit-button {
  flex: 1;
  margin: 0;
}
.return-button,
.submit-button {
  flex: 1;
  height: 44px;
  line-height: 44px;
  border-radius: 12px;
  font-size: 16px;
  margin: 0;
}
.return-button {
  background: #fafafa;
  color: #285ac8;
  border: 1px solid #285ac8;
}
.return-button::after {
  border: none;
}
.submit-button {
  background: #285ac8;
  color: #fff;
  border: none;
}
.submit-button::after {
  border: none;
}
.bottom-space {
  height: 14vh;
}
.tabbar {
  z-index: 999;
}
</style>
