<script>
import { diagnosisSubmitService,getReportService } from '../../services/clock'
export default {
  data() {
    return {
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
					console.log(res.data)
					console.log(res.data.diagnosiscirculatorysymptoms)
					this.AllCheckInfo = res.data
					this.AllCheckInfo.files.forEach(file => {
					    const field = file.reportId; // 假设fileId是你想用的字段
						console.log('id为',field);
					    const newFile = {
					        name: `${file.reportType}_report_${field}.pdf`, // 文件名可自定义
					        extname: "pdf", // 假设都是pdf格式，若有不同可根据实际情况修改
					        url: `https://ds.sccdc.cn/common/file/getDiagnosisReport/${field}`
					    };
									
					    if (file.reportType === "粪便常规检查报告") {
					        this.formData.stoolReport.push(newFile);
					    } else if (file.reportType === "血液检查报告") {
					        this.formData.bloodReport.push(newFile);
					    } else if (file.reportType === "影像学检查报告") {
					        this.formData.imageReport.push(newFile);
					    } else if (file.reportType === "B超报告") {
					        this.formData.ultrasoundReport.push(newFile);
					    } else if (file.reportType === "尿常规检查报告") {
					        this.formData.urineReport.push(newFile);
					    } else if (file.reportType === "病案首页") {
					        this.formData.medicalRecord.push(newFile);
					    }
					});
				})
				.catch(error => this.submitError(error));
				console.log(this.AllCheckInfo.files);
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
	chooseFile(reportType) {
	  wx.chooseMessageFile({
	    count: 3, // 最多可以选择3个文件
	    type: 'file',
	    extension: ['pdf', 'jpg', 'png'], // 支持的文件类型
	    success: (res) => {
	      const files = res.tempFiles.map(file => ({
	        path: file.path,
	        name: file.name
	      }));
	      this.formData[reportType] = [...this.formData[reportType], ...files];
	    }
	  });
	},
	
	uploadFile(filePath, reportType, index) {
	  return new Promise((resolve, reject) => {
	    wx.uploadFile({
	      url: 'https://ds.sccdc.cn/railwayemployee/diagnosis/file',  // 替换为你的接口地址
	      filePath: filePath,
	      name: 'file',  // 根据你的接口定义
	      formData: {
	        reportType: reportType,  // 上传的报告类型
	        index: index,             // 当前文件的索引
	      },
	      header: {
	        'Content-Type': 'application/x-www-form-urlencoded'
	      },
	      success: (res) => {
	        console.log(`文件上传成功: ${res}`);
			const fileId = parseInt(JSON.parse(res.data).data);
			console.log('添加了id',this.AllCheckInfo.fileIdList,"666",fileId)
	        this.AllCheckInfo.fileIdList.push(fileId);
	        resolve(res);
	      },
	      fail: (err) => {
	        console.error(`文件上传失败: ${err}`);
	        reject(err);
	      }
	    });
	  });
	},
	
	handleSubmits() {
	  if (this.isViewMode) return; 
	  if (this.AllCheckInfo.diagnosisResultsEmpDTO.diseaseType == null || this.AllCheckInfo.diagnosisResultsEmpDTO.discoveryMethod == null
		|| this.AllCheckInfo.diagnosisResultsEmpDTO.diseaseOutcome == null || this.AllCheckInfo.diagnosisResultsEmpDTO.diagnosisDate == '') {
	    uni.showToast({
	      title: '请填写基本信息',
	      icon: 'none',
	      duration: 2000,
	    });
	    return; // Stop the submission process
	  }
	  wx.showModal({
	    content: '提交后不可修改，是否确认提交？',
	    cancelText: '取消',
	    confirmText: '确认提交',
	    success: async (res) => {
	      if (res.confirm) {
	        // 上传所有报告的文件
	        const uploadPromises = [];
	        // 上传粪便检查报告
	        this.formData.stoolReport.forEach((file, index) => {
	          uploadPromises.push(this.uploadFile(file.path, '粪便常规检查报告', index));
	        });
	        // 上传血液检查报告
	        this.formData.bloodReport.forEach((file, index) => {
	          uploadPromises.push(this.uploadFile(file.path, '血液检查报告', index));
	        });
			this.formData.imageReport.forEach((file,index) => {
			  uploadPromises.push(this.uploadFile(file.path, '影像学检查报告', index));
			});
			this.formData.ultrasoundReport.forEach((file,index) => {
			  uploadPromises.push(this.uploadFile(file.path, 'B超报告', index));
			});
			this.formData.urineReport.forEach((file,index) => {
			  uploadPromises.push(this.uploadFile(file.path, '尿常规检查报告', index));
			});
			this.formData.medicalRecord.forEach((file,index) => {
			  uploadPromises.push(this.uploadFile(file.path, '病案首页', index));
			});
	        try {
	          // 等待所有文件上传完成
	          await Promise.all(uploadPromises);
	          console.log('所有文件上传完成');
	          // 提交表单
	          diagnosisSubmitService(this.AllCheckInfo)
	            .then((res) => this.submitSuccess(res))
	            .catch((error) => this.submitError(error));
	
	          wx.showToast({
	            title: '提交成功',
	            icon: 'success',
	            duration: 2000
	          });
	        } catch (error) {
	          console.error('上传过程出现错误: ', error);
	          wx.showToast({
	            title: '提交失败',
	            icon: 'none',
	            duration: 2000
	          });
	        }
	      } else if (res.cancel) {
	        console.log('用户点击取消');
	      }
	    }
	  });
	},
	submitError(error) {
	  console.log('报错了',error)
	  uni.showToast({
	    title: '提交失败',
	    icon: 'none',
	    duration: 1500
	  });
	},
	submitSuccess(response) {
	  console.log("注意注意",response)
	  if (response.code == 1) {
	    uni.showToast({
	      title: '信息提交成功',
	      icon: 'success',
	      duration: 1800
	    });
	  } else {
	    uni.showModal({
	      title: '信息提交失败',
	      content: response.msg,
	      showCancel: false,
	      confirmText: '确定',
	      success: function (res) {}
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
		this.AllCheckInfo.diagnosisotherSymptoms.sleepQuality = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisotherSymptoms.sleepQuality)
	},
	showNutritionStatus(maleValue){
		this.AllCheckInfo.diagnosisotherSymptoms.nutritionStatus = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisotherSymptoms.nutritionStatus)
	},
	showWorkLifeStress(maleValue){
		this.AllCheckInfo.diagnosisotherSymptoms.workLifeStress = maleValue;
		console.log('已经选择',this.AllCheckInfo.diagnosisotherSymptoms.workLifeStress)
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
		console.log(`Changing ${propName}:`, this.AllCheckInfo.diagnosisexaminations[propName]);
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
	handleFileChange(event, reportType) {
	    console.log(event, '66666');
	
		const selectFile =  [{
			path: event.tempFilePaths[0],
			name: event.tempFiles[0].name,
		}]
		console.log('666666666655555',selectFile)
	    this.formData[reportType] = [...this.formData[reportType], ...selectFile];
	
	    console.log(this.formData);
	},
  },
};
</script>
<template>
	<view class="container">
	  <!-- part1 -->
		<view class="content3" v-show="pages == 1">
		  <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
	      <view class="header">
	        <text class="title">确诊</text>疾病（必填）
	      </view>
		  <view class="part1" v-show="sonpages == 1">
	      <view class="section">
			  <view class="sections1">
				  <view class="section-title3">
				  	<text class="label">确诊疾病</text>
				  </view>
				  <view class="row3">
					<radio-group>
						<view style="margin: 5px 0">
							<radio :disabled="isViewMode" color="#285AC8" @tap="handleDiseaseTypeChange('新型冠状病毒感染')" :checked="AllCheckInfo.diagnosisresults.diseaseType == '新型冠状病毒感染'">新型冠状病毒感染</radio>
						</view>
						<view style="margin: 5px 0">
							<radio :disabled="isViewMode" color="#285AC8" @tap="handleDiseaseTypeChange('流感')" :checked="AllCheckInfo.diagnosisresults.diseaseType == '流感'" >流感</radio>
						</view>
						<view style="margin: 5px 0">
							<radio :disabled="isViewMode" color="#285AC8" @tap="handleDiseaseTypeChange('鼠疫')"  :checked="AllCheckInfo.diagnosisresults.diseaseType == '鼠疫'" >鼠疫</radio>
						</view>
						<view style="margin: 5px 0">
							<radio :disabled="isViewMode" color="#285AC8" @tap="handleDiseaseTypeChange('感染性腹泻')" :checked="AllCheckInfo.diagnosisresults.diseaseType == '感染性腹泻'" >感染性腹泻</radio>
						</view>
						<view style="margin: 5px 0">
							<radio :disabled="isViewMode" color="#285AC8" @tap="handleDiseaseTypeChange('炭疽')" :checked="AllCheckInfo.diagnosisresults.diseaseType == '炭疽'">炭疽</radio>
						</view>
						<view style="margin: 5px 0; margin-left: 25px;" v-show="AllCheckInfo.diagnosisResultsEmpDTO.diseaseType == '炭疽'">
							<radio-group>
							<view class="option-button" style="margin-top: 10px;">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showAnthraxSubtype('皮肤炭疽')" :checked="AllCheckInfo.diagnosisresults.anthraxSubtype == '皮肤炭疽'">皮肤炭疽</radio>
							<radio :disabled="isViewMode" color="#285AC8" @tap="showAnthraxSubtype('肠炭疽')" :checked="AllCheckInfo.diagnosisresults.anthraxSubtype == '肠炭疽'"  style="margin-left: 20px;">肠炭疽</radio>
							</view>
							<view class="option-button" style="margin-top: 10px;">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showAnthraxSubtype('肺炭疽')"  :checked="AllCheckInfo.diagnosisresults.anthraxSubtype == '肺炭疽'">肺炭疽</radio>
							<radio :disabled="isViewMode" color="#285AC8" @tap="showAnthraxSubtype('脑膜炎型炭疽')" style="margin-left: 20px;" :checked="AllCheckInfo.diagnosisresults.anthraxSubtype == '脑膜炎型炭疽'" >脑膜炎型炭疽</radio>
							</view>
							<view class="option-button" style="margin-top: 10px;">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showAnthraxSubtype('败血症型炭疽')" :checked="AllCheckInfo.diagnosisresults.anthraxSubtype == '败血症型炭疽'" >败血症型炭疽</radio>
							</view>
							</radio-group>
						</view>
						<view style="margin: 5px 0">
							<radio :disabled="isViewMode" color="#285AC8" @tap="handleDiseaseTypeChange('登革热（蚊媒传染病）')" :checked="AllCheckInfo.diagnosisresults.diseaseType == '登革热（蚊媒传染病）'" >登革热（蚊媒传染病）</radio>
						</view>
						<view style="margin: 5px 0">
							<radio :disabled="isViewMode" color="#285AC8" @tap="handleDiseaseTypeChange('疟疾（蚊媒传染病）')" :checked="AllCheckInfo.diagnosisresults.diseaseType == '疟疾（蚊媒传染病）'" >疟疾（蚊媒传染病）</radio>
						</view>
						<view style="margin: 5px 0"> 
							<radio :disabled="isViewMode" color="#285AC8" @tap="handleDiseaseTypeChange('森林脑炎（蜱媒传染病）')" :checked="AllCheckInfo.diagnosisresults.diseaseType == '森林脑炎（蜱媒传染病）'" >森林脑炎（蜱媒传染病）</radio>
						</view>
						<view style="margin: 5px 0">
							<radio :disabled="isViewMode" color="#285AC8" @tap="handleDiseaseTypeChange('发热伴血小板减少综合征（蜱媒传染病）')" :checked="AllCheckInfo.diagnosisresults.diseaseType == '发热伴血小板减少综合征（蜱媒传染病）'" >发热伴血小板减少综合征（蜱媒传染病）</radio>
						</view>
						<view style="margin: 5px 0">
							<radio :disabled="isViewMode" color="#285AC8" @tap="handleDiseaseTypeChange('斑疹伤寒')" :checked="AllCheckInfo.diagnosisresults.diseaseType == '斑疹伤寒'">斑疹伤寒</radio>
						</view>
						<view style="margin: 5px 0">
							<radio :disabled="isViewMode" color="#285AC8" @tap="handleDiseaseTypeChange('流行性出血热')" :checked="AllCheckInfo.diagnosisresults.diseaseType == '流行性出血热'" >流行性出血热</radio>
						</view>
						<view style="margin: 5px 0">
							<radio :disabled="isViewMode" color="#285AC8" @tap="handleDiseaseTypeChange('其他')" :checked="AllCheckInfo.diagnosisresults.diseaseType == '其他'" >其他</radio>
							<input :disabled="isViewMode" type="text" placeholder="请输入其他疾病名称" v-model="AllCheckInfo.diagnosisResultsEmpDTO.otherDiseaseName" class="text-input" style="float: right;background-color: #FAFAFA;" v-show="AllCheckInfo.diagnosisResultsEmpDTO.diseaseType == '其他'"/>
						</view>
					</radio-group>
				</view>
			</view>
	      </view>
		  
		  <view class="section">
			<view class="sections1">
		  		<view class="section-title3">
		  			<text class="label">病例发现途径</text>
		  		</view>
			<view class="row3">
				<radio-group>
					<view style="margin: 10px 0">
						<radio :disabled="isViewMode" color="#285AC8" @tap="handleDiscoveryMethodChange('APP填报')" :checked="AllCheckInfo.diagnosisresults.discoveryMethod == 'APP填报'" >APP填报</radio>
						<radio :disabled="isViewMode" color="#285AC8" @tap="handleDiscoveryMethodChange('主动就诊')" :checked="AllCheckInfo.diagnosisresults.discoveryMethod == '主动就诊'"  style="float: right; margin-left: 8vh;">主动就诊</radio>
					</view>
					<view style="margin: 10px 0">
						<radio :disabled="isViewMode" color="#285AC8" @tap="handleDiscoveryMethodChange('常规监测发现')" :checked="AllCheckInfo.diagnosisresults.discoveryMethod == '常规监测发现'" >常规监测发现</radio>
					</view>
					<view style="margin: 10px 0">
						<radio :disabled="isViewMode" color="#285AC8" @tap="handleDiscoveryMethodChange('其他')" :checked="AllCheckInfo.diagnosisresults.discoveryMethod == '其他'" >其他</radio>
					</view>
				</radio-group>
			</view>
		  	</view>
			<view class="row3" v-show="!isViewMode">
			  <view class="form-item3">
			    <view class="section-title3" style="margin-left: -25px;">
			      <text class="label">提交时间</text>
			    </view>
			    <picker mode="date" value="2024-01-01" start="2024-01-01" end="2099-12-31" @change="onDateChange1" class="time-picker">
			      <view class="picker">
			        {{ AllCheckInfo.diagnosisResultsEmpDTO.diagnosisDate || '请选择日期' }}
			      </view>
			    </picker>
			  </view>
			</view>
		  </view>
		  <view class="section">
			<view class="sections1">
				<view class="section-title3">
				  	<text class="label">疾病转归</text>
				</view>
				  <view class="row3">
				  	<radio-group>
				    <view style="margin: 10px 0">
				        <radio :disabled="isViewMode" color="#285AC8" @tap="handleDiseaseOutcomeChange('治愈')" :checked="AllCheckInfo.diagnosisresults.diseaseOutcome == '治愈'" >治愈</radio>
				  		<radio :disabled="isViewMode" color="#285AC8" @tap="handleDiseaseOutcomeChange('好转')" :checked="AllCheckInfo.diagnosisresults.diseaseOutcome == '好转'"  style="float: right; margin-left: 8vh;">好转</radio>
				    </view>
				    <view style="margin: 10px 0">
				      <radio :disabled="isViewMode" color="#285AC8" @tap="handleDiseaseOutcomeChange('未愈')" :checked="AllCheckInfo.diagnosisresults.diseaseOutcome == '未愈'" >未愈</radio>
				    </view>
				    </radio-group>
				  </view>
				</view>
			</view>
			<button class="single-submit-button" @tap="handleSubmit()">下一步</button>
			</view>
		  <view class="part2" v-show="sonpages == 2">
		  <view class="section">
		  			  <view class="sections1">
		  				  <view class="section-title3">
		  				  	<text class="label">确诊医院名称</text>
		  				  </view>
		  				  <view class="row3">
							<input type="text" placeholder="请输入确诊医院名称" v-model="AllCheckInfo.diagnosisResultsEmpDTO.hospitalName" style="height: 40px;background-color: #FAFAFA;padding-left: 9px;width: 70vh;"/>
		  				  </view>
		  			</view>
		  </view>
<!-- 		  <view class="section">
		  			<view class="sections1">
		  		<view class="section-title3">
		  			<text class="label">就诊/入院日期</text>
		  		</view>
		  			<view class="row3">
		  				<picker mode="date" value="2024-01-01" start="2024-01-01" end="2099-12-31" @change="onDateChange" class="time-picker">
		  				  <view style="float: right; margin-top: 6px; margin-left: 16px;">
		  				    {{ AllCheckInfo.diagnosisResultsEmpDTO.admissionDate || '请选择日期' }}
		  				  </view>
		  				</picker>
		  			</view>
		  	</view>
		  </view> -->
		  <view class="section">
		  			  <view class="sections1">
		  				  <view class="section-title3">
		  				  	<text class="label">就诊/入院时症状和体征</text>
		  				  </view>
		  				  <view class="row3">
		  					<input :disabled="isViewMode" type="text" placeholder="请输入就诊/入院时症状和体征" v-model="AllCheckInfo.diagnosisResultsEmpDTO.admissionSymptomsAndSigns" style="height: 40px;background-color: #FAFAFA;padding-left: 9px;width: 70vh;"/>
		  				  </view>
		  			</view>
		  </view>
		  <view class="section">
		  			<view class="sections1">
		  				<view class="section-title3">
		  				  	<text class="label">登记分类</text>
		  				</view>
		  				  <view class="row3">
		  				  	<radio-group>
		  				    <view style="margin: 10px 0">
		  				        <radio :disabled="isViewMode" color="#285AC8" @tap="handleRegistrationClassificationChange('新患者')" :checked="AllCheckInfo.diagnosisresults.registrationClassification == '新患者'" >新患者</radio>
		  				  		<radio :disabled="isViewMode" color="#285AC8" @tap="handleRegistrationClassificationChange('复发')" :checked="AllCheckInfo.diagnosisresults.registrationClassification == '复发'" style="float: right; margin-left: 8vh;">复发</radio>
		  				    </view>
		  				    <view style="margin: 10px 0">
		  				      <radio :disabled="isViewMode" color="#285AC8" @tap="handleRegistrationClassificationChange('返回')" :checked="AllCheckInfo.diagnosisresults.registrationClassification == '返回'" >返回</radio>
		  				  	  <radio :disabled="isViewMode" color="#285AC8" @tap="handleRegistrationClassificationChange('治疗失败')"  :checked="AllCheckInfo.diagnosisresults.registrationClassification == '治疗失败'"  style="float: right; margin-left: 8vh;">治疗失败</radio>
		  				    </view>
							<view style="margin: 10px 0">
							  <radio :disabled="isViewMode" color="#285AC8" @tap="handleRegistrationClassificationChange('其他')" :checked="AllCheckInfo.diagnosisresults.registrationClassification == '其他'" >其他（请详述）</radio>
							</view>
							<input :disabled="isViewMode" type="text" placeholder="请详述" v-model="AllCheckInfo.diagnosisResultsEmpDTO.otherRegistrationDetails" style="height: 40px;background-color: #FAFAFA;padding-left: 9px;width: 40vh;"/>
		  				    </radio-group>
		  				  </view>
		  				</view>
		  			</view>
		  			<view class="button-container">
		  			<button class="return-button" @tap="handlereturn">上一步</button>
		  			<button class="submit-button" @tap="handleSubmit">下一步</button>
		  			</view>
		  			</view>
		</view>
		<view class="content3"  v-show="pages == 2">
			<image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
			<view class="header">
			  <text class="title">基本</text>数据
			</view>
				<view class="section">
				<view class="sections1">
					<view class="section-title3">
						<text class="label">身高、体重、血压</text>
					</view>
						<view class="row3">
							<view class="form-item3">
								<text style="float: left;">身高</text>
								<input :disabled="isViewMode" type="text" v-model="AllCheckInfo.diagnosispersonalinfo.height" placeholder="请输入身高(cm)" class="text-input" style="float: right;width: 120px; text-align: right;"/>
							</view>
						</view>
						<view class="row3">
							<view class="form-item3">
								<text style="float: left;">体重</text>
								<input type="text" placeholder="请输入体重(kg)" v-model="AllCheckInfo.diagnosispersonalinfo.weight" class="text-input" style="float: right;width: 120px; text-align: right;"/>
							</view>
						</view>
						<view class="row3">
							<view class="form-item3">
								<text style="float: left;">血压</text>
								<input :disabled="isViewMode" type="text" placeholder="请输入血压(mmHg)" v-model="AllCheckInfo.diagnosispersonalinfo.bloodPressure" class="text-input" style="float: right;width: 120px; text-align: right;"/>
							</view>
						</view>
					<view class="section-title3">
						<text class="label">血糖</text>
					</view>
						<view class="row3">
							<view class="form-item3">
								<text style="float: left;">空腹血糖</text>
								<input :disabled="isViewMode" type="text" placeholder="请输入空腹血糖(mmol/L)" v-model="AllCheckInfo.diagnosispersonalinfo.bloodSugarFasting" class="text-input" style="text-align: right; float: right;width: 160px;"/>
							</view>
						</view>
						<view class="row3">
							<view class="form-item3">
								<text style="float: left;">随机血糖</text>
								<input :disabled="isViewMode" type="text" placeholder="请输入空腹血糖(mmol/L)" v-model="AllCheckInfo.diagnosispersonalinfo.bloodSugarRandom" class="text-input" style="text-align: right; float: right;width: 160px;"/>
							</view>
						</view>
					<view class="section-title3">
						<text class="label">血脂</text>
					</view>
						<view class="row3">
							<view class="form-item3">
								<text style="float: left;">总胆固醇</text>
								<input :disabled="isViewMode" type="text" placeholder="请输入总胆固醇(mmol/L)" v-model="AllCheckInfo.diagnosispersonalinfo.cholesterolTotal" class="text-input" style="text-align: right; float: right;width: 160px;"/>
							</view>
						</view>
						<view class="row3">
							<view class="form-item3">
								<text style="float: left;">甘油三酯</text>
								<input :disabled="isViewMode" type="text" placeholder="请输入甘油三酯(mmol/L)" v-model="AllCheckInfo.diagnosispersonalinfo.triglycerides" class="text-input" style="text-align: right; float: right;width: 160px;"/>
							</view>
						</view>
						<view class="row3">
							<view class="form-item3">
								<text style="float: left;">胆固醇酯</text>
								<input :disabled="isViewMode" type="text" placeholder="请输入胆固醇酯(mmol/L)" v-model="AllCheckInfo.diagnosispersonalinfo.cholesterolLipid" class="text-input" style="text-align: right; float: right;width: 160px;"/>
							</view>
						</view>
						<view class="row3">
							<view class="form-item3">
								<text style="float: left;">高密度脂蛋白</text>
								<input :disabled="isViewMode" type="text" placeholder="请输入高密度脂蛋白(mmol/L)" v-model="AllCheckInfo.diagnosispersonalinfo.hdl" class="text-input" style="text-align: right; float: right;width: 160px;"/>
							</view>
						</view>
						<view class="row3">
							<view class="form-item3">
								<text style="float: left;">低密度脂蛋白</text>
								<input :disabled="isViewMode" type="text" placeholder="请输入低密度脂蛋白(mmol/L)" v-model="AllCheckInfo.diagnosispersonalinfo.ldl" class="text-input" style="text-align: right; float: right;width: 160px;"/>
							</view>
						</view>
						<view class="section-title3">
							<text class="label">血氧饱和度</text>
						</view>
						<view class="row3">
							<view class="form-item3">
								<text style="float: left;">血氧饱和度</text>
								<input :disabled="isViewMode" type="text" placeholder="请输入血氧饱和度(%)" v-model="AllCheckInfo.diagnosispersonalinfo.bloodOxygen" class="text-input" style="text-align: right; float: right;width: 160px;"/>
							</view>
						</view>
				</view>
				<view class="button-container">
				<button class="return-button" @tap="handlereturn">上一步</button>
				<button class="submit-button" @tap="handleSubmit">下一步</button>
				</view>
		</view>
	  </view>
	  
	  
	  <view class="container" v-show="pages == 3">
	    <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
	  	<view class="content3">
	        <view class="header">
	          <text class="title">全身</text>症状
	        </view>
	  	  <!-- 第一个 -->
	        <view class="section">
	  		<view class="sections1">
	  			<view class="section-title3">
	  				<text class="label">发热情况</text>
	  			</view>
	  			<view class="row3">
	  				<text class="label">是否发热</text>
	  				<radio-group class="options">
	  					<label class="radio">
	  						<radio :disabled="isViewMode" @tap="hanleGeneralSymprion('hasFever')" color="#285AC8" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasFever" >是</radio>
	  					</label>
	  					<label class="radio">
	  						<radio :disabled="isViewMode" @tap="hanleGeneralSymprion('hasFever')" color="#285AC8" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasFever" >否</radio>
	  					</label>
	  				</radio-group>
	  			</view>
	  		</view>
	          <view class="row3">
	            <text class="label">最高温度</text>
	  		  <input :disabled="isViewMode" type="text" v-model="AllCheckInfo.diagnosisgeneralsymptoms.highestTemperature" placeholder="请输入温度" style="width: 10vh;"/>
	            <text class="label">持续时间</text>
	            <input :disabled="isViewMode" type="text" v-model="AllCheckInfo.diagnosisgeneralsymptoms.feverDuration" placeholder="请输入小时数" style="width: 10vh;"/>
	          </view>
	        </view>
	    
	  		<!-- 第二个 -->
	        <view class="section">
	  		  <view class="sections1">
	  			  <view class="section-title3">
	  			  	<text class="label">请选择症状</text>
	  			  </view>
	          <view class="row3">
	            <checkbox-group class="options">
	  			<view class="order-options">
	              <label class="checkbox">
	                <checkbox :disabled="isViewMode" color="#285AC8" value="chills"  :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasChills"  @tap="handleCheckboxChange('hasChills')">寒战</checkbox>
	              </label>
	              <label class="checkbox">
	                <checkbox :disabled="isViewMode" color="#285AC8" value="sweating" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasSweating"  @tap="handleCheckboxChange('hasSweating')">大汗</checkbox>
	              </label>
	              <label class="checkbox">
	                <checkbox :disabled="isViewMode" color="#285AC8" value="fatigue"  :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasFatigue"  @tap="handleCheckboxChange('hasFatigue')">乏力</checkbox>
	              </label>
	              <label class="checkbox">
	                <checkbox  :disabled="isViewMode" color="#285AC8" value="thirst" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasHeadache"  @tap="handleCheckboxChange('hasHeadache')">头痛</checkbox>
	              </label>
	  			</view>
	  			<view class="order-options" style="margin-right: 60px;">
	  			<label>
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('hasMusclePain')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasMusclePain" >肌肉酸痛</checkbox>
	  			</label>
	  			<label style="margin-left: 20px;">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('hasJointPain')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasJointPain" >关节酸痛</checkbox>
	  			</label>
	  			</view>
	            </checkbox-group>
	          </view>
	  		</view>
	        </view>
	    
	  		<!-- 第三个 -->
	  		<view class="section">
	  					<view class="sections1">
	  						<view class="section-title3">
	  							<text class="label">皮肤、疹、疮、淋巴结情况</text>
	  						</view>
	  					</view>
	  					<view class="section" style="margin-bottom: -6px;">
	  					  <checkbox-group>
	  					    <label>
	  					      <checkbox :disabled="isViewMode" color="#285AC8" value="淋巴结肿大" @tap="handleCheckboxChange('hasLymphNodeSwelling')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasLymphNodeSwelling" />
	  					      淋巴结肿大
	  					    </label>
	  						<view class="order-options" style="margin-left: 30px;" v-if="AllCheckInfo.diagnosisgeneralsymptoms.hasLymphNodeSwelling == true">
	  					    <label>
	  					      <checkbox :disabled="isViewMode" color="#285AC8" value="腹股沟" @tap="handleCheckboxChange('hasLymphNodeSwellingGroin')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasLymphNodeSwellingGroin" />
	  					      腹股沟
	  					    </label>
	  					    <label>
	  					      <checkbox :disabled="isViewMode" color="#285AC8" value="腋下" @tap="handleCheckboxChange('hasLymphNodeSwellingArmpit')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasLymphNodeSwellingArmpit" />
	  					      腋下
	  					    </label>
	  					    <label>
	  					      <checkbox :disabled="isViewMode" color="#285AC8" value="锁骨下" @tap="handleCheckboxChange('hasLymphNodeSwellingSubclavian')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasLymphNodeSwellingSubclavian" />
	  					      锁骨下
	  					    </label>
	  						</view>
	  						<view class="order-options" style="margin-left: 30px;" v-if="AllCheckInfo.diagnosisgeneralsymptoms.hasLymphNodeSwelling == true">
	  					    <label>
	  					      <checkbox :disabled="isViewMode" color="#285AC8" value="颈部" @tap="handleCheckboxChange('hasLymphNodeSwellingNeck')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasLymphNodeSwellingNeck" />
	  					      颈部
	  					    </label>
	  					    <label>
	  					      <checkbox :disabled="isViewMode" color="#285AC8" value="肘部" @tap="handleCheckboxChange('hasLymphNodeSwellingElbow')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasLymphNodeSwellingElbow" />
	  					      肘部
	  					    </label>
	  					    <label>
	  					      <checkbox :disabled="isViewMode" color="#285AC8" value="腘窝" @tap="handleCheckboxChange('hasLymphNodeSwellingPopliteal')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasLymphNodeSwellingPopliteal" />
	  					      腘窝
	  					    </label>
	  						</view>
	  					  </checkbox-group>
	  					</view>
	  					<view class="checkboxs3" >
	  					      <checkbox-group>
	  							<view style="display: flex; flex-direction: column;margin-left: 12px;">
	  					        <label class="choice-item">
	  					          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('hasLipCyanosis')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasLipCyanosis" />
	  					          口唇、颜面、四肢及全身皮肤发绀
	  					        </label>
	  					        <label class="choice-item">
	  					          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('hasSubcutaneousAndMucosalBleedingSpots')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasSubcutaneousAndMucosalBleedingSpots" />
	  					          皮下及粘膜出血或出血点
	  					        </label>
	  					        <label class="choice-item">
	  					          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('hasPainfulRedRash')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasPainfulRedRash" />
	  					          皮肤可见剧痛的红色丘疹
	  					        </label>
	  					        <label class="choice-item">
	  					          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('hasBloodBlisters')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasBloodBlisters" />
	  					          皮肤可见周边呈灰黑色、基底坚硬的血性水泡
	  					        </label>
	  					        <label class="choice-item">
	  					          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('hasSkinUlcer')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasSkinUlcer" />
	  					          皮肤出现呈灰黑色创面的溃疡
	  					        </label>
	  					        <label class="choice-item">
	  					          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('hasCongestiveOrPetechialRash')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasCongestiveOrPetechialRash" />
	  					          皮疹：充血性或点状出血疹
	  					        </label>
	  					        <label class="choice-item">
	  					          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('hasPainfulRedRash')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasPainfulRedRash" />
	  					          皮疹：起初压之褪色，后期压之不退，多见于四肢、躯干、面部一般没有
	  					        </label>
	  					        <label class="choice-item">
	  					          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('hasItchyRash')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.hasItchyRash" />
	  					          瘙痒性丘疹/水疱
	  					        </label>
	  							</view>
	  					      </checkbox-group>
	  					    </view>
	  					    <view class="inner-checkboxs3" v-if="AllCheckInfo.diagnosisgeneralsymptoms.hasItchyRash == true">
	  					      <checkbox-group>
	  							<view class="order-options">
	  					        <label>
	  					          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('itchyRashOnFingers')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.itchyRashOnFingers" />
	  					          手指
	  					        </label>
	  					        <label>
	  					          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('itchyRashOnBackOfHands')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.itchyRashOnBackOfHands" />
	  					          手背
	  					        </label>
	  					        <label>
	  					          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('itchyRashOnUpperLimbs')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.itchyRashOnUpperLimbs" />
	  					          上肢
	  					        </label>
	  							</view>
	  							<view class="order-options">
	  					        <label>
	  					          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('itchyRashOnLowerLimbs')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.itchyRashOnLowerLimbs" />
	  					          下肢
	  					        </label>
	  					        <label>
	  					          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('itchyRashOnFeet')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.itchyRashOnFeet" />
	  					          足背
	  					        </label>
	  					        <label>
	  					          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('itchyRashOnFace')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.itchyRashOnFace" />
	  					          面部
	  					        </label>
	  							</view>
	  							<view class="order-options">
	  					        <label>
	  					          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('itchyRashOnOther')" :checked="AllCheckInfo.diagnosisgeneralsymptoms.itchyRashOnOther" />
	  					          其他
	  					        </label>
	  							</view>
	  					      </checkbox-group>
	  					    </view>
	  		</view>
	        <view class="section">
	  		  <view class="sections1">
	  			  <view class="section-title3">
	  			  	<text class="label">失水情况</text>
	  			  </view>
	  			  <view class="row3">
	  			  	<text class="label">失水</text>
	  			  	<radio-group class="options">
	  			  		<label class="radio">
	  			  			<radio :disabled="isViewMode" color="#285AC8" :checked="AllCheckInfo.diagnosisgeneralsymptoms.dehydrationSeverity === '轻度'"  @tap="showRadio('轻度')">轻度</radio>
	  			  		</label>
	  			  		<label class="radio">
	  			  			<radio :disabled="isViewMode" color="#285AC8" :checked="AllCheckInfo.diagnosisgeneralsymptoms.dehydrationSeverity === '中度'"  @tap="showRadio('中度')">中度</radio>
	  			  		</label>
	  					<label class="radio">
	  						<radio :disabled="isViewMode" color="#285AC8" :checked="AllCheckInfo.diagnosisgeneralsymptoms.dehydrationSeverity === '重度'"  @tap="showRadio('重度')">重度</radio>
	  					</label>
	  			  	</radio-group>
	  			  </view>
	  		  </view>
	  	  </view>
	        <view class="button-container">
	        <button class="return-button" @tap="handlereturn">上一步</button>
	        <button class="submit-button" @tap="handleSubmit">下一步</button>
	        </view>
	  	  </view>
	  	<uni-custom-tabbar></uni-custom-tabbar>
	  </view>
	  <view class="container" v-show="pages == 4">
	    <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
	  	<view class="content3">
	        <view class="header">
	          <text class="title">呼吸</text>系统症状
	        </view>
	        <view class="section">
	  		  <view class="sections1">
	  			  <view class="section-title3">
	  			  	<text class="label">请选择症状</text>
	  			  </view>
	          <view class="row3">
	            <checkbox-group class="options">
	  			<view class="order-options" style="margin-bottom: -5px;">
	              <label class="checkbox">
	                <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasCough')" :checked="AllCheckInfo.diagnosisrespiratorysymptoms.hasCough"  >咳嗽</checkbox>
	              </label>
	              <label class="checkbox">
	                <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasSputum')" :checked="AllCheckInfo.diagnosisrespiratorysymptoms.hasSputum" >咳痰</checkbox>
	              </label>
	              <label class="checkbox">
	                <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasBloodySputum')" :checked="AllCheckInfo.diagnosisrespiratorysymptoms.hasBloodySputum" >血痰</checkbox>
	              </label>
	              <label class="checkbox">
	                <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasHemoptysis')" :checked="AllCheckInfo.diagnosisrespiratorysymptoms.hasHemoptysis" >咳血</checkbox>
	              </label>
	  			</view>
	  			<view class="order-options" style="margin-bottom: -5px;">
	  			<label class="checkbox">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasSoreThroat')" :checked="AllCheckInfo.diagnosisrespiratorysymptoms.hasSoreThroat" >咽痛</checkbox>
	  			</label>
	  			<label class="checkbox">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasDryThroat')" :checked="AllCheckInfo.diagnosisrespiratorysymptoms.hasDryThroat" >咽干</checkbox>
	  			</label>
	  			<label class="checkbox">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasNasalCongestion')" :checked="AllCheckInfo.diagnosisrespiratorysymptoms.hasNasalCongestion" >鼻塞</checkbox>
	  			</label>
	  			<label class="checkbox">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasRunnyNose')" :checked="AllCheckInfo.diagnosisrespiratorysymptoms.hasRunnyNose" >流涕</checkbox>
	  			</label>
	  			</view>
	  			<view class="order-options" style="margin-bottom: -5px;">
	  			<label style="margin-left: 5px;">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasChestTightness')" :checked="AllCheckInfo.diagnosisrespiratorysymptoms.hasChestTightness" >胸闷</checkbox>
	  			</label>
	  			<label style="margin-left: 18px;" >
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasShortnessOfBreath')" :checked="AllCheckInfo.diagnosisrespiratorysymptoms.hasShortnessOfBreath" >气促</checkbox>
	  			</label>
	  			</view>
	  			<view class="order-options" style="margin-bottom: -5px; ">
	  			<label style="margin-left: 5px;">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasBreathingDifficulty')" :checked="AllCheckInfo.diagnosisrespiratorysymptoms.hasBreathingDifficulty" >呼吸困难</checkbox>
	  			</label>
	  			<label style="margin-left: 18px;">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasHoarseness')" :checked="AllCheckInfo.diagnosisrespiratorysymptoms.hasHoarseness" >声音嘶哑</checkbox>
	  			</label>
	  			</view>
	  			<view class="order-options" style="margin-bottom: -5px; margin-right: 160px;">
	  			<label  style="margin-left: 5px;">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasBronchitis')" :checked="AllCheckInfo.diagnosisrespiratorysymptoms.hasBronchitis" >支气管炎/支气管肺炎</checkbox>
	  			</label>
	  			</view>
	            </checkbox-group>
	          </view>
	  		</view>
	        </view>
	  	  <view class="button-container">
	  	  <button class="return-button" @tap="handlereturn">上一步</button>
	  	  <button class="submit-button" @tap="handleSubmit">下一步</button>
	  	  </view>
	  	  </view>
	  	<uni-custom-tabbar></uni-custom-tabbar>
	  </view>
	  <view class="container" v-show="pages == 5">
	    <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
	  	<view class="content3">
	      <view class="header">
	        <text class="title">消化</text>系统症状
	      </view>
	  	
	  	<!-- 第一个 -->
	  	<view class="section">
	  				<view class="sections1">
	  					<view class="section-title3">
	  						<text class="label">腹泻情况</text>
	  					</view>
	  					<view class="row3">
	  						<text class="label">腹泻次数</text>
	  						<radio-group class="options">
	  							<label class="radio">
	  								<radio :disabled="isViewMode" color="#285AC8" @tap="handleDigestFrequency('diarrheaFrequencyGEThreeTimesPerDay')" :checked="AllCheckInfo.diagnosisdigestivesymptoms.diarrheaFrequencyGEThreeTimesPerDay">少于3次/天</radio>
	  							</label>
	  							<label class="radio">
	  								<radio :disabled="isViewMode" color="#285AC8" @tap="handleDigestFrequency('diarrheaFrequencyGEThreeTimesPerDay')" :checked="!AllCheckInfo.diagnosisdigestivesymptoms.diarrheaFrequencyGEThreeTimesPerDay" >大于等于3次/天</radio>
	  							</label>
	  						</radio-group>
	  					</view>
	  					<input type="number" placeholder="请输入具体次数" v-model="AllCheckInfo.diagnosisdigestivesymptoms.diarrheaTimesPerDay" v-show="AllCheckInfo.diagnosisdigestivesymptoms.diarrheaFrequencyGEThreeTimesPerDay"/>
	  					<view class="row3">
	  						<text class="label">腹泻方式</text>
	  						<radio-group class="options">
	  							<view class="option-button">
	  							<label class="radio">
	  								<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadio('里急后重')"  :checked="AllCheckInfo.diagnosisdigestivesymptoms.diarrheaMode == '里急后重'" >里急后重</radio>
	  							</label>
	  							<label class="radio"  style="margin-top: 10px;">
	  								<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadio('失禁')"  :checked="AllCheckInfo.diagnosisdigestivesymptoms.diarrheaMode == '失禁'" >失禁</radio>
	  							</label>
	  							</view>
	  							<view class="option-button">
	  							<label class="radio">
	  								<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadio('畅通')" :checked="AllCheckInfo.diagnosisdigestivesymptoms.diarrheaMode == '通畅'" >畅通</radio>
	  							</label>
	  							<label class="radio" style="margin-top: 10px;">
	  								<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadio('绞痛')" :checked="AllCheckInfo.diagnosisdigestivesymptoms.diarrheaMode == '绞痛'" >绞痛</radio>
	  							</label>
	  							</view>
	  						</radio-group>
	  					</view>
	  				</view>
	  	</view>
	  	<!-- 第二个 -->
	  	<view class="section">
	  				<view class="sections1">
	  					<view class="section-title3">
	  						<text class="label">粪便情况</text>
	  					</view>
	  					<view class="row3">
	  						<text class="label">粪便量</text>
	  						<radio-group class="options">
	  							<label class="radio">
	  								<radio :disabled="isViewMode" color="#285AC8" @tap="showStoolAmount('多')"  :checked="AllCheckInfo.diagnosisdigestivesymptoms.stoolAmount == '多'" >多</radio>
	  							</label>
	  							<label class="radio">
	  								<radio :disabled="isViewMode" color="#285AC8" @tap="showStoolAmount('少')"  :checked="AllCheckInfo.diagnosisdigestivesymptoms.stoolAmount == '少'" >少</radio>
	  							</label>
	  						</radio-group>
	  					</view>
	  					<view class="row3">
	  						<text class="label">粪便气味</text>
	  						<radio-group class="options">
	  							<label class="radio">
	  								<radio :disabled="isViewMode" color="#285AC8" @tap="showStoolOdor('恶臭')" :checked="AllCheckInfo.diagnosisdigestivesymptoms.stoolOdor == '恶臭'" >恶臭</radio>
	  							</label>
	  							<label class="radio">
	  								<radio :disabled="isViewMode" @tap="showStoolOdor('无恶臭')" color="#285AC8" :checked="AllCheckInfo.diagnosisdigestivesymptoms.stoolOdor == '无恶臭'" >无恶臭</radio>
	  							</label>
	  						</radio-group>
	  					</view>
	  					<view class="checkboxs3" >
	  					      <checkbox-group>
	  							<view style="display: flex; flex-direction: column;">
	  					        <label>
	  					          <checkbox :checked="temporaryChoice1 == true" value="粪便性状1" color="#285AC8" @tap="temporaryChoice1 = !temporaryChoice1" :disabled="temporaryChoice2 == true"/>
	  					          粪便性状1
	  					        </label>
	  							<radio-group class="options" v-show="temporaryChoice1" style="margin-top: 6px;" >
	  								<label class="radio">
	  									<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadioBystool1('水样')"  :checked="AllCheckInfo.diagnosisdigestivesymptoms.stoolType1Detail == '水样'" >水样</radio>
	  								</label>
	  								<label class="radio">
	  									<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadioBystool1('米泔样')" :checked="AllCheckInfo.diagnosisdigestivesymptoms.stoolType1Detail == '米泔样'" >米泔样</radio>
	  								</label>
	  								<label class="radio">
	  									<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadioBystool1('洗肉水样')" :checked="AllCheckInfo.diagnosisdigestivesymptoms.stoolType1Detail == '洗肉水样'" >洗肉水样</radio>
	  								</label>
	  							</radio-group>
	  					        <label style="margin-top: 10px;">
	  					          <checkbox :checked="temporaryChoice2 == true" value="粪便性状2" color="#285AC8" @tap="temporaryChoice2 = !temporaryChoice2" :disabled="temporaryChoice1 == true" />
	  					          粪便性状2
	  					        </label>
	  							<radio-group class="options" v-show="temporaryChoice2" style="margin-top: 6px;">
	  								<label class="radio">
	  									<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadioBystool2('黑便')" :checked="AllCheckInfo.diagnosisdigestivesymptoms.stoolType2Detail == '黑便'" >黑便</radio>
	  								</label>
	  								<label class="radio">
	  									<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadioBystool2('脓血')" :checked="AllCheckInfo.diagnosisdigestivesymptoms.stoolType2Detail == '脓血'" >脓血</radio>
	  								</label>
	  								<label class="radio">
	  									<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadioBystool2('大块粘膜')" :checked="AllCheckInfo.diagnosisdigestivesymptoms.stoolType2Detail == '大块粘膜'" >大块粘膜</radio>
	  								</label>
	  							</radio-group>
	  							</view>
	  					      </checkbox-group>
	  					    </view>
	  				</view>
	  	</view>
	  	
	  	
	  	<!-- 第二个 -->
	  	<view class="section">
	  				<view class="sections1">
	  					<view class="section-title3">
	  						<text class="label">呕吐情况</text>
	  					</view>
	  					<view class="row3">
	  						<text class="label">是否呕吐</text>
	  						<radio-group class="options">
	  							<label class="radio">
	  								<radio :disabled="isViewMode" color="#285AC8" @tap="handleDigestFrequency('hasVomiting')" :checked="AllCheckInfo.diagnosisdigestivesymptoms.hasVomiting" >是</radio>
	  							</label>
	  							<label class="radio">
	  								<radio :disabled="isViewMode"  color="#285AC8" @tap="handleDigestFrequency('hasVomiting')" :checked="!AllCheckInfo.diagnosisdigestivesymptoms.hasVomiting">否</radio>
	  							</label>
	  						</radio-group>
	  					</view>
	  					<view class="row3">
	  						<text class="label">呕吐方式</text>
	  						<radio-group class="options">
	  							<view class="option-button" style="margin-top: 20px;">
	  							<label class="radio">
	  								<radio :disabled="isViewMode"  color="#285AC8" @tap="showDigestVomitingMode('喷射状')" :checked="AllCheckInfo.diagnosisdigestivesymptoms.vomitingMode == '喷射状'" >喷射状</radio>
	  							</label>
	  							<label class="radio" style="margin-top: 10px;">
	  								<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestVomitingMode('先泻后吐')" :checked="AllCheckInfo.diagnosisdigestivesymptoms.vomitingMode == '先泻后吐'" >先泻后吐</radio>
	  							</label>
	  							</view>
	  							<view class="option-button" style="margin-top: 20px;">
	  							<label class="radio">
	  								<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestVomitingMode('先吐后泻')" :checked="AllCheckInfo.diagnosisdigestivesymptoms.vomitingMode == '先吐后泻'" >先吐后泻</radio>
	  							</label>
	  							</view>
	  						</radio-group>
	  					</view>
	  				</view>
	  	</view>
	  		<!-- 第三个 -->
	  		<view class="section">
	  					<view class="sections1">
	  						<view class="section-title3">
	  							<text class="label">请选择其他消化系统症状</text>
	  						</view>
	  					</view>
	  					<view class="checkboxs3" style="margin-top: 10px;" >
	  					      <checkbox-group>
	  							<view class="order-options" style="margin-left: 60px;">
	  					        <label>
	  					          <checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisdigestivesymptoms.hasNausea"  value="恶心" color="#285AC8" @tap="handleDigestFrequency('hasNausea')" />
	  					          恶心
	  					        </label>
	  					        <label>
	  					          <checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisdigestivesymptoms.hasAbdominalDistension"  value="腹胀" color="#285AC8" @tap="handleDigestFrequency('hasAbdominalDistension')"/>
	  					          腹胀
	  					        </label>
	  					        <label>
	  					          <checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisdigestivesymptoms.hasUpperAbdominalDiscomfort"  value="上腹部不适" color="#285AC8" @tap="handleDigestFrequency('hasUpperAbdominalDiscomfort')" />
	  					          上腹部不适
	  					        </label>
	  							</view>
	  							<view class="order-options" style="margin-left: 60px;">
	  					        <label>
	  					          <checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisdigestivesymptoms.hasAbdominalPain"  value="腹痛" color="#285AC8" @tap="handleDigestFrequency('hasAbdominalPain')" />
	  					          腹痛
	  					        </label>
	  					        <label>
	  					          <checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisdigestivesymptoms.hasBorborygmus" value="腹鸣" color="#285AC8" @tap="handleDigestFrequency('hasBorborygmus')"/>
	  					          腹鸣
	  					        </label>
	  					        <label>
	  					          <checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisdigestivesymptoms.hasConstipation"  value="便秘" color="#285AC8"  @tap="handleDigestFrequency('hasConstipation')"/>
	  					          便秘
	  					        </label>
	  							</view>
	  							<view class="order-options" style="margin-left: 60px;">
	  					        <label>
	  					          <checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisdigestivesymptoms.hasOliguriaOrAnuria" value="少尿或无尿" color="#285AC8" @tap="handleDigestFrequency('hasOliguriaOrAnuria')"/>
	  					          少尿或无尿
	  					        </label>
	  					        <label>
	  					          <checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisdigestivesymptoms.hasAppetiteLoss" value="食欲减退" color="#285AC8" @tap="handleDigestFrequency('hasAppetiteLoss')"/>
	  					          食欲减退
	  					        </label>
	  							</view>
	  					      </checkbox-group>
	  					    </view>
	  		</view>
	        <view class="button-container">
	        <button class="return-button" @tap="handlereturn">上一步</button>
	        <button class="submit-button" @tap="handleSubmit">下一步</button>
	        </view>
	  	  </view>
	  	<uni-custom-tabbar></uni-custom-tabbar>
	  </view>
	  <view class="container" v-show="pages == 6">
	    <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
	  	<view class="content3">
	        <view class="header">
	          <text class="title">循环</text>系统症状
	        </view>
	        <view class="section">
	  		  <view class="sections1">
	  			  <view class="section-title3">
	  			  	<text class="label">请选择症状</text>
	  			  </view>
	          <view class="row3">
	            <checkbox-group class="options">
	  			<view class="order-options" >
	              <label style="margin-left: 5px;">
	                <checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosiscirculatorysymptoms.hasArrhythmia" color="#285AC8" @tap="handleCirculatorysymptoms('hasArrhythmia')">心律不齐</checkbox>
	              </label>
	              <label style="margin-left: 15px;">
	                <checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosiscirculatorysymptoms.hasChestPain" color="#285AC8" @tap="handleCirculatorysymptoms('hasChestPain')">胸痛</checkbox>
	              </label>
	              <label style="margin-left: 15px;">
	                <checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosiscirculatorysymptoms.hasRapidPulse" color="#285AC8" @tap="handleCirculatorysymptoms('hasRapidPulse')">脉搏细速</checkbox>
	              </label>
	  			</view>
	  			<view class="order-options" style="margin-right: 130px;">
	              <label style="margin-left: 5px;">
	                <checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosiscirculatorysymptoms.hasPalpitation" color="#285AC8" @tap="handleCirculatorysymptoms('hasPalpitation')">心悸</checkbox>
	              </label>
	  			<label style="margin-left: 15px;">
	  				<checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosiscirculatorysymptoms.hasLowBloodPressure" color="#285AC8" @tap="handleCirculatorysymptoms('hasLowBloodPressure')">低血压</checkbox>
	  			</label>
	  			</view>
	            </checkbox-group>
	          </view>
	  		</view>
	        </view>
	  	  <view class="button-container">
	  	  <button class="return-button" @tap="handlereturn">上一步</button>
	        <button class="submit-button" @tap="handleSubmit">下一步</button>
	  	  </view>
	  	  </view>
	  	<uni-custom-tabbar></uni-custom-tabbar>
	  </view>
	  <view class="container" v-show="pages == 7">
	    <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
	  	<view class="content3">
	        <view class="header">
	          <text class="title">神经</text>系统症状
	        </view>
	        <view class="section">
	  		  <view class="sections1">
	  			  <view class="section-title3">
	  			  	<text class="label">请选择症状</text>
	  			  </view>
	          <view class="row3">
	            <checkbox-group class="options">
	  			<view class="order-options">
	              <label class="checkbox">
	                <checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisneurologicalsymptoms.hasDelirium" color="#285AC8" @tap="handleNeurologicalSymptoms('hasDelirium')">儋妄</checkbox>
	              </label>
	              <label class="checkbox">
	                <checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisneurologicalsymptoms.hasConvulsion" color="#285AC8" @tap="handleNeurologicalSymptoms('hasConvulsion')">抽搐</checkbox>
	              </label>
	              <label class="checkbox">
	                <checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisneurologicalsymptoms.hasDrowsiness" color="#285AC8" @tap="handleNeurologicalSymptoms('hasDrowsiness')">昏睡</checkbox>
	              </label>
	              <label class="checkbox">
	                <checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisneurologicalsymptoms.hasComa" color="#285AC8" @tap="handleNeurologicalSymptoms('hasComa')">昏迷</checkbox>
	              </label>
	  			</view>
	  			<view class="order-options" style="margin-right: 115px;">
	  			<label style="margin-left: 5px;">
	  				<checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisneurologicalsymptoms.hasStiffNeck" color="#285AC8" @tap="handleNeurologicalSymptoms('hasStiffNeck')">颈部强直</checkbox>
	  			</label>
	  			<label style="margin-left: 15px;">
	  				<checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisneurologicalsymptoms.hasRestlessness" color="#285AC8" @tap="handleNeurologicalSymptoms('hasRestlessness')">烦躁不安</checkbox>
	  			</label>
	  			</view>
	  			<view class="order-options" style="margin-right: 146px;">
	  			<label style="margin-left: 5px;">
	  				<checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisneurologicalsymptoms.hasMuscleParalysisNeckShoulder" color="#285AC8" @tap="handleNeurologicalSymptoms('hasMuscleParalysisNeckShoulder')">颈肩部肌肉迟缓性瘫痪</checkbox>
	  			</label>
	  			</view>
	  			<view class="order-options" style="margin-right: 156px;">
	  			<label style="margin-left: 5px;">
	  				<checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisneurologicalsymptoms.hasMuscleParalysisLimbs" color="#285AC8" @tap="handleNeurologicalSymptoms('hasMuscleParalysisLimbs')">肢体肌肉迟缓性瘫痪</checkbox>
	  			</label>
	  			</view>
	  			<view class="order-options"  style="margin-right: 115px;">
	  			<label style="margin-left: 5px;">
	  				<checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisneurologicalsymptoms.hasSwallowingDifficulty" color="#285AC8" @tap="handleNeurologicalSymptoms('hasSwallowingDifficulty')">吞咽困难</checkbox>
	  			</label>
	  			<label style="margin-left: 15px;">
	  				<checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisneurologicalsymptoms.hasSpeechDisorder" color="#285AC8" @tap="handleNeurologicalSymptoms('hasSpeechDisorder')">语言障碍</checkbox>
	  			</label>
	  			</view>
	  			<view class="order-options" style="margin-right: 15px;">
	  			<label style="margin-left: 5px;" >
	  				<checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisneurologicalsymptoms.hasConsciousnessDisorder" color="#285AC8" @tap="handleNeurologicalSymptoms('hasConsciousnessDisorder')">意识障碍或惊厥</checkbox>
	  			</label>
	  			<label class="checkbox">
	  				<checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisneurologicalsymptoms.hasDizziness" color="#285AC8" @tap="handleNeurologicalSymptoms('hasDizziness')">头晕</checkbox>
	  			</label>
	  			<label class="checkbox">
	  				<checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisneurologicalsymptoms.hasTinnitus" color="#285AC8" @tap="handleNeurologicalSymptoms('hasTinnitus')">耳鸣</checkbox>
	  			</label>
	  			</view>
	  			<view class="order-options"  style="margin-right: 225px;">
	  			<label style="margin-left: 5px;">
	  				<checkbox :disabled="isViewMode" :checked="AllCheckInfo.diagnosisneurologicalsymptoms.hasHearingLoss" color="#285AC8" @tap="handleNeurologicalSymptoms('hasHearingLoss')">听力减退</checkbox>
	  			</label>
	  			</view>
	            </checkbox-group>
	          </view>
	  		</view>
	        </view>
	  	  <view class="button-container">
	  	  <button class="return-button" @tap="handlereturn">上一步</button>
	        <button class="submit-button" @tap="handleSubmit">下一步</button>
	  	  </view>
	  	  </view>
	  	<uni-custom-tabbar></uni-custom-tabbar>
	  </view>
	  <view class="container" v-show="pages == 8">
	    <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
	  	<view class="content3">
	        <view class="header">
	          <text class="title">局部</text>症状
	        </view>
	        <view class="section">
	  		  <view class="sections1">
	  			  <view class="section-title3">
	  			  	<text class="label">请选择症状</text>
	  			  </view>
	          <view class="row3">
	            <checkbox-group class="options" >
	  			<view class="order-options" style="margin-right: 50px;">
	              <label style="margin-left: 5px;">
	                <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasConjunctivitis')" :checked="AllCheckInfo.diagnosislocalsymptoms.hasConjunctivitis">结膜炎</checkbox>
	              </label>
	              <label style="margin-left: 15px;">
	                <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasConjunctivalCongestion')" :checked="AllCheckInfo.diagnosislocalsymptoms.hasConjunctivalCongestion" >眼结膜充血、肿胀疼痛</checkbox>
	              </label>
	  			</view>
	  			<view class="order-options" style="margin-right: 25px;">
	              <label style="margin-left: 5px;">
	                <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasEyelidEdema')" :checked="AllCheckInfo.diagnosislocalsymptoms.hasEyelidEdema">上下眼睑水肿</checkbox>
	              </label>
	              <label style="margin-left: 15px;">
	                <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasSmellTasteLoss')" :checked="AllCheckInfo.diagnosislocalsymptoms.hasSmellTasteLoss">嗅觉味觉减退或丧失</checkbox>
	              </label>
	  			</view>
	  			<view class="order-options" style="margin-right: 10px;">
	  			<label style="margin-left: 5px;">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasFacialRednessOrPallor')" :checked="AllCheckInfo.diagnosislocalsymptoms.hasFacialRednessOrPallor" >颜面潮红或苍白</checkbox>
	  			</label>
	  			<label style="margin-left: 15px;">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasRednessNeckChest')" :checked="AllCheckInfo.diagnosislocalsymptoms.hasRednessNeckChest" >潮红（颈部、胸部）</checkbox>
	  			</label>
	  			</view>
	  			<view class="order-options" style="margin-right: 65px;">
	  			<label style="margin-left: 5px;">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasLipCyanosis')" :checked="AllCheckInfo.diagnosislocalsymptoms.hasLipCyanosis" >口唇发绀</checkbox>
	  			</label>
	  			<label style="margin-left: 15px;">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasSkinPallorOrBruising')" :checked="AllCheckInfo.diagnosislocalsymptoms.hasSkinPallorOrBruising" >皮肤苍白或带青紫</checkbox>
	  			</label>
	  			</view>
	  			<view class="order-options" style="margin-right: 120px;">
	  			<label style="margin-left: 5px;">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasGumBleeding')" :checked="AllCheckInfo.diagnosislocalsymptoms.hasGumBleeding" >牙龈出血</checkbox>
	  			</label>
	  			<label style="margin-left: 15px;">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasKidneyPain')" :checked="AllCheckInfo.diagnosislocalsymptoms.hasKidneyPain" >肾区疼痛</checkbox>
	  			</label>
	  			</view>
	  			<view class="order-options" style="margin-right: 160px;">
	  			<label style="margin-left: 5px;">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasLowBackPain')" :checked="AllCheckInfo.diagnosislocalsymptoms.hasLowBackPain" >腰痛</checkbox>
	  			</label>
	  			<label style="margin-left: 15px;">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasOrbitalPain')" :checked="AllCheckInfo.diagnosislocalsymptoms.hasOrbitalPain" >眼眶痛</checkbox>
	  			</label>
	  			</view>
	  			<view class="order-options" style="margin-right: 105px;">
	  			<label style="margin-left: 5px;">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasMucosalBleedingInOralOrNasal')" :checked="AllCheckInfo.diagnosislocalsymptoms.hasMucosalBleedingInOralOrNasal">口腔、鼻腔等处黏膜有出血点</checkbox>
	  			</label>
	  			</view>
	  			<view class="order-options" style="margin-right: 35px;" >
	  			<label style="margin-left: 5px;">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasSkinBleedingPoints')" :checked="AllCheckInfo.diagnosislocalsymptoms.hasSkinBleedingPoints" >腋下/上臂/胸部或其他部位有皮肤出血点</checkbox>
	  			</label>
	  			</view>
	            </checkbox-group>
	          </view>
	  		</view>
	        </view>
	  	  <view class="button-container">
	  	  <button class="return-button" @tap="handlereturn">上一步</button>
	        <button class="submit-button" @tap="handleSubmit">下一步</button>
	  	  </view>
	  	  </view>
	  	<uni-custom-tabbar></uni-custom-tabbar>
	  </view>
	  <view class="container" v-show="pages == 9">
	    <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
	  	<view class="content3">
	        <view class="header">
	          <text class="title">其他</text>
	        </view>
	        <view class="section">
	  		  <view class="sections1">
	  			  <view class="section-title3">
	  			  	<text class="label">发病情况</text>
	  			  </view>
	          <view class="row3">
	            <checkbox-group class="options">
	  			<view class="order-options" style="margin-right: 70px;">
	              <label style="margin-left: 5px;">
	                <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleOtherSymptoms('hasSuddenOnset')" :checked="AllCheckInfo.diagnosisothersymptoms.hasSuddenOnset" >突然发病</checkbox>
	              </label>
	              <label style="margin-left: 15px;">
	                <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleOtherSymptoms('hasRapidProgress')" :checked="AllCheckInfo.diagnosislocalsymptoms.hasRapidProgress" >病情进展迅速</checkbox>
	              </label>
	  			</view>
	  			<view class="order-options" style="margin-right: 80px;">
	              <label style="margin-left: 5px;">
	                <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleOtherSymptoms('hasPeriodicAttack')" :checked="AllCheckInfo.diagnosislocalsymptoms.hasPeriodicAttack" >周期性发作</checkbox>
	              </label>
	              <label style="margin-left: 15px;">
	                <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleOtherSymptoms('hasForcedPosture')" :checked="AllCheckInfo.diagnosislocalsymptoms.hasForcedPosture" >强迫体位</checkbox>
	              </label>
	  			</view>
	  			<view class="order-options" style="margin-right: 200px;">
	  			<label style="margin-left: 5px;">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleOtherSymptoms('hasCalfMusclePain')" :checked="AllCheckInfo.diagnosislocalsymptoms.hasCalfMusclePain" >腓肠肌疼痛</checkbox>
	  			</label>
	  			</view>
	            </checkbox-group>
	          </view>
	  		</view>
	        </view>
	  	  <view class="button-container">
	  	  <button class="return-button" @tap="handlereturn">上一步</button>
	        <button class="submit-button" @tap="handleSubmit">下一步</button>
	  	  </view>
	  	  </view>
	  	<uni-custom-tabbar></uni-custom-tabbar>
	  </view>
	  
	  <view class="container" v-show="pages == 10">
	    <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
	  	<view class="content3">
	        <view class="header">
	          <text class="title">并发</text>症
	        </view>
	        <view class="section">
			  <view class="sections1">
			  	<view class="section-title3">
			  		<text class="label">并发症情况</text>
			  	</view>
			  	<view class="row3">
			  		<text class="label">有无并发症</text>
			  		<radio-group class="options">
			  			<label class="radio">
			  				<radio :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasComplications')" :checked="AllCheckInfo.diagnosisotherSymptoms.sleepQuality == '有'" >有</radio>
			  			</label>
			  			<label class="radio">
			  				<radio :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasComplications')" :checked="AllCheckInfo.diagnosisotherSymptoms.sleepQuality == '无'" >无</radio>
			  			</label>
			  		</radio-group>
			  	</view>
			  </view>
	  		  <view class="sections1">
	  			  <view class="section-title3">
	  			  	<text class="label">选择并发症</text>
	  			  </view>
	          <view class="row3">
	            <checkbox-group class="options">
	  			<view class="order-options" style="margin-right: 70px;">
	              <label style="margin-left: 5px;">
	                <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasViralPneumonia')" :checked="AllCheckInfo.diagnosisothersymptoms.hasSuddenOnset" >病毒性肺炎</checkbox>
	              </label>
	              <label style="margin-left: 15px;">
	                <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasBacterialPneumonia')" :checked="AllCheckInfo.diagnosisothersymptoms.hasRapidProgress" >细菌性肺炎</checkbox>
	              </label>
	  			</view>
				<view class="order-options" style="margin-right: 70px;">
				  <label style="margin-left: 5px;">
				    <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasFungalPneumonia')" :checked="AllCheckInfo.diagnosisothersymptoms.hasSuddenOnset" >真菌性肺炎</checkbox>
				  </label>
				  <label style="margin-left: 15px;">
				    <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasRespiratoryFailure')" :checked="AllCheckInfo.diagnosisothersymptoms.hasRapidProgress" >呼吸衰竭</checkbox>
				  </label>
				</view>
				<view class="order-options" style="margin-right: 70px;">
				  <label style="margin-left: 5px;">
				    <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasArds')" :checked="AllCheckInfo.diagnosisothersymptoms.hasSuddenOnset" >急性呼吸窘迫综合征(ARDS)</checkbox>
				  </label>
				</view>
				<view class="order-options" style="margin-right: 70px;">
				  <label style="margin-left: 5px;">
				    <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasDic')" :checked="AllCheckInfo.diagnosisothersymptoms.hasSuddenOnset" >弥散性血管内凝血（DIC）</checkbox>
				  </label>
				</view>
				<view class="order-options" style="margin-right: 70px;">
				  <label style="margin-left: 5px;">
				    <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasToxicShock')" :checked="AllCheckInfo.diagnosisothersymptoms.hasSuddenOnset" >感染中毒性休克</checkbox>
				  </label>
				  <label style="margin-left: 15px;">
				    <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasMeningitis')" :checked="AllCheckInfo.diagnosisothersymptoms.hasRapidProgress" >脑膜炎</checkbox>
				  </label>
				</view>
				<view class="order-options" style="margin-right: 70px;">
				  <label style="margin-left: 5px;">
				    <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasEncephalitis')" :checked="AllCheckInfo.diagnosisothersymptoms.hasSuddenOnset">脑炎</checkbox>
				  </label>
				  <label style="margin-left: 15px;">
				    <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasBacteremia')" :checked="AllCheckInfo.diagnosisothersymptoms.hasRapidProgress" >菌血症</checkbox>
				  </label>
				  <label style="margin-left: 15px;">
				    <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasMyocarditis')":checked="AllCheckInfo.diagnosisothersymptoms.hasRapidProgress" >心肌炎</checkbox>
				  </label>
				</view>
				<view class="order-options" style="margin-right: 70px;">
				  <label style="margin-left: 5px;">
				    <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasAcuteLungInjury')" :checked="AllCheckInfo.diagnosisothersymptoms.hasSuddenOnset">急性肺损伤/ARDS</checkbox>
				  </label>
				  <label style="margin-left: 15px;">
				    <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasAcuteKidneyInjury')" :checked="AllCheckInfo.diagnosisothersymptoms.hasRapidProgress" >急性肾损伤</checkbox>
				  </label>
				</view>
				<view class="order-options" style="margin-right: 70px;">
				  <label style="margin-left: 5px;">
				    <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasEpilepsy')" :checked="AllCheckInfo.diagnosisothersymptoms.hasSuddenOnset">癫痫</checkbox>
				  </label>
				  <label style="margin-left: 15px;">
				    <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasToxicHepatitis')" :checked="AllCheckInfo.diagnosisothersymptoms.hasRapidProgress" >中毒性肝炎</checkbox>
				  </label>
				</view>
				<view class="order-options" style="margin-right: 70px;">
				  <label style="margin-left: 5px;">
				    <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasFluidOverload')" :checked="AllCheckInfo.diagnosisothersymptoms.hasSuddenOnset" >输液过量</checkbox>
				  </label>
				  <label style="margin-left: 15px;">
				    <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasSuperinfection')" :checked="AllCheckInfo.diagnosisothersymptoms.hasRapidProgress" >二重感染</checkbox>
				  </label>
				</view>
				<view class="order-options" style="margin-right: 70px;">
				  <label style="margin-left: 5px;">
				    <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasHypoxemia')" :checked="AllCheckInfo.diagnosisothersymptoms.hasSuddenOnset" >低血氧症</checkbox>
				  </label>
				  <label style="margin-left: 15px;">
				    <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasSepticShock')" :checked="AllCheckInfo.diagnosisothersymptoms.hasRapidProgress" >脓毒症休克</checkbox>
				  </label>
				</view>
	  			<view class="order-options" style="margin-right: 80px;">
	              <label style="margin-left: 5px;">
	                <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasMetabolicAcidosis')" :checked="AllCheckInfo.diagnosisothersymptoms.hasPeriodicAttack" >代谢性酸中毒</checkbox>
	              </label>
	              <label style="margin-left: 15px;">
	                <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasElectrolyteImbalance')" :checked="AllCheckInfo.diagnosisothersymptoms.hasForcedPosture" >电解质及酸碱失衡</checkbox>
	              </label>
	  			</view>
	  			<view class="order-options" >
	  			<label style="margin-left: 5px;">
	  				<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasHemolysis')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >急性血管内溶血</checkbox>
	  			</label>
				<label style="margin-left: 15px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasGastrointestinalBleeding')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >消化道出血</checkbox>
				</label>
	  			</view>
				<view class="order-options">
				<label style="margin-left: 5px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasCoagulationDisorder')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >出凝血功能障碍</checkbox>
				</label>
				</view>
				<view class="order-options" >
				<label style="margin-left: 5px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasMultipleOrganFailure')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >多器官功能衰竭</checkbox>
				</label>
				<label style="margin-left: 15px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasBronchitis')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >支气管炎</checkbox>
				</label>
				</view>
				<view class="order-options">
				<label style="margin-left: 5px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasBronchopneumonia')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >支气管肺炎</checkbox>
				</label>
				<label style="margin-left: 15px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasAnemia')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >贫血</checkbox>
				</label>
				</view>
				<view class="order-options" >
				<label style="margin-left: 5px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasHepatosplenomegaly')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >肝脾肿大</checkbox>
				</label>
				<label style="margin-left: 15px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasBrainEdema')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >脑水肿</checkbox>
				</label>
				</view>
				
				<view class="order-options" >
				<label style="margin-left: 5px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasHus')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >溶血性尿毒综合征</checkbox>
				</label>
				<label style="margin-left: 15px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasHematemesis')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >呕血</checkbox>
				</label>
				</view>
				<view class="order-options" >
				<label style="margin-left: 5px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasMelena')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >便血</checkbox>
				</label>
				<label style="margin-left: 15px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasHemoptysis')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >咯血</checkbox>
				</label>
				<label style="margin-left: 15px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasAbdominalHemorrhage')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >腹腔出血</checkbox>
				</label>
				</view>
				<view class="order-options" >
				<label style="margin-left: 5px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasNasalBleeding')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >鼻出血</checkbox>
				</label>
				<label style="margin-left: 15px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasVaginalBleeding')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >阴道出血</checkbox>
				</label>
				</view>
				
				<view class="order-options" >
				<label style="margin-left: 5px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasHypertensiveEncephalopathy')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >高血压脑病</checkbox>
				</label>
				<label style="margin-left: 15px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasIntracranialHemorrhage')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >颅内出血</checkbox>
				</label>
				</view>
				<view class="order-options" >
				<label style="margin-left: 5px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasPulmonaryEdema')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >肺水肿</checkbox>
				</label>
				<label style="margin-left: 15px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasSecondaryInfection')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >继发性感染</checkbox>
				</label>
				</view>
				<view class="order-options" >
				<label style="margin-left: 5px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasSpontaneousRenalRupture')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >自发性肾破裂</checkbox>
				</label>
				<label style="margin-left: 15px;">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleComplications('hasMyocardialDamage')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >心肌损害</checkbox>
				</label>
				</view>
				<view class="order-options" >
				<label style="margin-left: 5px;">
					<checkbox v-show="!isViewMode" color="#285AC8" @tap="handleComplications('hasLiverDamage')" :checked="AllCheckInfo.diagnosisothersymptoms.hasCalfMusclePain" >肝损害</checkbox>
				</label>
				</view>
				<view class="order-options" >
				<checkbox color="#285AC8" style="margin-bottom: 10px;" @tap="handleComplications('otherComplications')">其他</checkbox>
					<input :disabled="isViewMode" type="number" placeholder="请填写其他症状" style="margin-left: 30px;" v-model="AllCheckInfo.diagnosiscomplications.otherComplicationsName"/>
				</view>
				</checkbox-group>
	          </view>
	  		</view>
	        </view>
	  	  <view class="button-container">
	  	  <button class="return-button" @tap="handlereturn">上一步</button>
	        <button class="submit-button" @tap="handleSubmit">下一步</button>
	  	  </view>
	  	  </view>
	  	<uni-custom-tabbar></uni-custom-tabbar>
	  </view>
	  <view class="container" v-show="pages == 11">
	    <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
	    <view class="content3">
			<view class="header">
			  <text class="title">检测</text>报告
			</view>
	      <!-- 病原学及血清学检查 -->
	      <view class="section">
	        <view class="sections1">
	          <view class="section-title3">
	            <text class="label">病原学及血清学检查（可多选）</text>
	          </view>
	          <view class="checkboxs3">
	            <checkbox-group>
	              <view class="order-options">
	                <label class="checkbox-item">
	                  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleDiagnosisexaminations('hasBloodTest')" :checked="AllCheckInfo.diagnosisexaminations.hasBloodTest"  />
	                  血常规
	                </label>
	                <label class="checkbox-item">
	                  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleDiagnosisexaminations('hasImaging')"  :checked="AllCheckInfo.diagnosisexaminations.hasImaging"  />
	                  影像学检查
	                </label>
	              </view>
	              <view class="order-options">
	                <label class="checkbox-item">
	                  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleDiagnosisexaminations('hasUrineTest')"  :checked="AllCheckInfo.diagnosisexaminations.hasUrineTest" />
	                  尿常规
	                </label>
	                <label class="checkbox-item">
	                  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleDiagnosisexaminations('hasStoolTest')" :checked="AllCheckInfo.diagnosisexaminations.hasStoolTest"  />
	                  粪便常规
	                </label>
	              </view>
				  <view class="order-options">
				    <label class="checkbox-item">
				      <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleDiagnosisexaminations('hasUltrasound')" :checked="AllCheckInfo.diagnosisexaminations.hasUltrasound" />
				      B超
				    </label>
				  </view>
	            </checkbox-group>
	          </view>
	  
	          <view class="row3" style="margin-top: 45px; margin-bottom: 70px;">
	            <view class="form-item3">
	              <view class="section-title3"  style="margin-left: -25px;">
	                <text class="label">胸部X线或 CT或核磁共振等影像学检查报告</text>
	              </view>
				  <view class="example-body">
				      <uni-file-picker 
					      fileMediatype="image"
					      :show-type="'grid'"
					      :disabled="isViewMode"
					      :value="formData.imageReport"
					      :files-list="formData.imageReport"
				          :limit="3" 
				          :file-type="['image', 'pdf']" 
				          @select="handleFileChange($event, 'imageReport')"
				          title="最多选择3个文件 (PDF, JPG, PNG)">
				      </uni-file-picker>
				  </view>
<!-- 	              <button @tap="chooseFile('imageReport')" class="center-button" style="float: right;">文件上传</button>
	              <view v-if="formData.imageReport.length > 0" class="file-preview">
	                <view v-for="(file, index) in formData.imageReport" :key="index" class="file-item">
	                  <image v-if="isImage(file.path)" :src="file.path" mode="aspectFit" style="width: 100px; height: 100px;" />
	                  <view v-else>{{ file.name }}</view>
	                </view>
	              </view> -->
	            </view>
	          </view>
	  
	          <view class="row3" style="margin-top: 85px; margin-bottom: 80px;">
	            <view class="form-item3">
	              <view class="section-title3"  style="margin-left: -25px;">
	                <text class="label">B超报告</text>
	              </view>
				  <view class="example-body">
				      <uni-file-picker 
					      fileMediatype="image"
					      :show-type="'grid'"
					      :disabled="isViewMode"
						  :value="formData.ultrasoundReport"
						  :files-list="formData.ultrasoundReport"
				          :limit="3" 
				          :file-type="['image', 'pdf']" 
				          @select="handleFileChange($event, 'ultrasoundReport')"
				          title="最多选择3个文件 (PDF, JPG, PNG)">
				      </uni-file-picker>
				  </view>
<!-- 	              <button @tap="chooseFile('ultrasoundReport')" class="center-button" style="float: right;">文件上传</button>
	              <view v-if="formData.ultrasoundReport.length > 0" class="file-preview">
	                <view v-for="(file, index) in formData.ultrasoundReport" :key="index" class="file-item">
	                  <image v-if="isImage(file.path)" :src="file.path" mode="aspectFit" style="width: 100px; height: 100px;" />
	                  <view v-else>{{ file.name }}</view>
	                </view>
	              </view> -->
	            </view>
	          </view>
	          <view class="row3" style="margin-top: 85px; margin-bottom: 80px;">
	            <view class="form-item3">
	              <view class="section-title3"  style="margin-left: -25px;">
	                <text class="label">血液检查报告（血常规、C反应蛋白、血沉等）</text>
	              </view>
				  <view class="example-body">
				      <uni-file-picker 
					      fileMediatype="image"
						  :show-type="'grid'"
					      :disabled="isViewMode"
				          :limit="3" 
						  :value="formData.bloodReport"
						  :files-list="formData.bloodReport"
				          :file-type="['image', 'pdf']" 
				          @select="handleFileChange($event, 'bloodReport')"
				          title="最多选择3个文件 (PDF, JPG, PNG)">
				      </uni-file-picker>
				  </view>
<!-- 	              <button @tap="chooseFile('bloodReport')" class="center-button" style="float: right;">文件上传</button>
	              <view v-if="formData.bloodReport.length > 0" class="file-preview">
	                <view v-for="(file, index) in formData.bloodReport" :key="index" class="file-item">
	                  <image v-if="isImage(file.path)" :src="file.path" mode="aspectFit" style="width: 100px; height: 100px;" />
	                  <view v-else>{{ file.name }}</view>
	                </view>
	              </view> -->
	            </view>
	          </view>
			  <view class="row3" style="margin-top: 95px; margin-bottom: 80px;">
			    <view class="form-item3">
			      <view class="section-title3"  style="margin-left: -25px;">
			        <text class="label">粪便常规检查报告</text>
			      </view>
				  <view class="example-body">
				      <uni-file-picker 
					      fileMediatype="image"
					      :show-type="'grid'"
					      :disabled="isViewMode"
					      :value="formData.stoolReport"
					      :files-list="formData.stoolReport"
				          :limit="3" 
				          :file-type="['image', 'pdf']" 
				          @select="handleFileChange($event, 'stoolReport')"
				          title="最多选择3个文件 (PDF, JPG, PNG)">
				      </uni-file-picker>
				  </view>
<!-- 			      <button @tap="chooseFile('stoolReport')" class="center-button" style="float: right;">文件上传</button>
			      <view v-if="formData.stoolReport.length > 0" class="file-preview">
			        <view v-for="(file, index) in formData.stoolReport" :key="index" class="file-item">
			          <image v-if="isImage(file.path)" :src="file.path" mode="aspectFit" style="width: 100px; height: 100px;" />
			          <view v-else>{{ file.name }}</view>
			        </view>
			      </view> -->
			    </view>
			  </view>
			  
			  <view class="row3" style="margin-top: 85px; margin-bottom: 85px;">
			    <view class="form-item3">
			      <view class="section-title3"  style="margin-left: -25px;">
			        <text class="label">尿常规检查报告</text>
			      </view>
				  <view class="example-body">
				      <uni-file-picker 
					      fileMediatype="image"
					      :show-type="'grid'"
					      :disabled="isViewMode"
					      :value="formData.urineReport"
					      :files-list="formData.urineReport"
				          :limit="3" 
				          :file-type="['image', 'pdf']" 
				          @select="handleFileChange($event, 'urineReport')"
				          title="最多选择3个文件 (PDF, JPG, PNG)">
				      </uni-file-picker>
				  </view>
<!-- 			      <button @tap="chooseFile('urineReport')" class="center-button" style="float: right;">文件上传</button>
			      <view v-if="formData.urineReport.length > 0" class="file-preview">
			        <view v-for="(file, index) in formData.urineReport" :key="index" class="file-item">
			          <image v-if="isImage(file.path)" :src="file.path" mode="aspectFit" style="width: 100px; height: 100px;" />
			          <view v-else>{{ file.name }}</view>
			        </view>
			      </view> -->
			    </view>
			  </view>
			  
			  <view class="row3" style="margin-top: 85px; margin-bottom: 80px;">
			    <view class="form-item3">
			      <view class="section-title3"  style="margin-left: -25px;">
			        <text class="label">病案首页</text>
			      </view>
				  <view class="example-body">
				      <uni-file-picker 
					      fileMediatype="image"
					      :show-type="'grid'"
					      :disabled="isViewMode"
					      :value="formData.medicalRecord"
					      :files-list="formData.medicalRecord"
				          :limit="3" 
				          :file-type="['image', 'pdf']" 
				          @select="handleFileChange($event, 'medicalRecord')"
				          title="最多选择3个文件 (PDF, JPG, PNG)">
				      </uni-file-picker>
				  </view>
<!-- 			      <button @tap="chooseFile('medicalRecord')" class="center-button" style="float: right;">文件上传</button>
			      <view v-if="formData.medicalRecord.length > 0" class="file-preview">
			        <view v-for="(file, index) in formData.medicalRecord" :key="index" class="file-item">
			          <image v-if="isImage(file.path)" :src="file.path" mode="aspectFit" style="width: 100px; height: 100px;" />
			          <view v-else>{{ file.name }}</view>
			        </view>
			      </view> -->
			    </view>
			  </view>
	  
	  
	        </view>
	  	  <view class="button-container" style="margin-top: 14vh;">
	  	  <button class="return-button" @tap="handlereturn">上一步</button>
	        <button class="submit-button" @tap="handleSubmit">提交</button>
	  	  </view>
	      </view>
	    </view>
	  </view>
		<uni-custom-tabbar class="tabbar"></uni-custom-tabbar>
	</view>
</template>
<style>
.container {
	width: 100%;
	height: 100vh;
	position: relative; /* 设置为相对定位，使内部元素能够相对于容器定位 */
}
	
.background-image {
	position: absolute; /* 设置为绝对定位，使其占据整个容器 */
	top: 0;
	right: 0;
	bottom: 0;
	left: 0;
	width: 100%;
	height: 100%;
	z-index: -1; /* 将图片置于内容下方 */
}
.form-item3{
	width: 100%;
	margin: 10px 8px;
	width: 343px;
	height: 50px;
	opacity: 1;
	border-radius: 5px;
	background: rgba(255, 255, 255, 1);
	padding: 16px;
	margin-bottom: 8px;
	display: block;
}
.content3{
	position: relative;
	padding-top: 11vh;
	border-top-left-radius: 40px;
	border-top-right-radius: 40px;
	padding-bottom: 13vh;
}
.header {
  position: relative;
  text-align: left;
  margin-bottom: 20px;
  font-size: 26px;
  font-weight: bold;
  padding: 0 10px;
}

.title {
  position: relative;
}
  .title::after {
    content: "";
    position: absolute;
    bottom: -4px;
    left: 0;
    width: 100%;
    height: 6px;
    background-color: #285AC8;
  }
.section {
  margin-top: 15px;
  margin-bottom: 10px;
  background-color: #ffffff;
  padding: 12px;
  border-radius: 8px;
}
.section-title3{
	margin-left: -3px;
	padding-left: 4px;
	border-left: 8px solid #285AC8;
}
.row3 {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 6px;
  margin-bottom: 18px;
}
.checkboxs3{
	display: block;
}
.inner-checkboxs3{
	margin-left: 4vh;
}
.label {
  font-size: 16px;
  color: #333333;
}

.options {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between; /* 确保选项在一行内均匀分布 */
}

.checkbox {
  display: flex;
  align-items: center;
  margin: 2px; /* 添加四周的边距 */
}
@media (min-width: 200px) {
  .options .checkbox {
    flex-basis: calc(25% - 10px); /* 每个选项占25%，减去左右各5px的边距 */
  }
}
.radio, .checkbox {
  display: flex;
  align-items: center;
  margin-right: 10px;
}

.picker {
  padding: 5px 10px;
  border-radius: 4px;
  background-color: #ffffff;
}
.button-container {
  display: flex;
  justify-content: space-between; /* 均匀分配剩余空间 */
  align-items: center;
}

.submit-button {
  width: 50%;
  height: 44px;
  background-color: #285AC8;
  color: #ffffff;
  text-align: center;
  border-radius: 8px;
  font-size: 16px;
  margin: 0 24px;
}
.return-button{
	width: 30%;
	height: 44px;
	background-color: #FAFAFA;
	border: 1px solid #285AC8;
}
.single-submit-button{
	height: 44px;
	background-color: #285AC8;
	color: #ffffff;
	text-align: center;
	border-radius: 8px;
	font-size: 16px;
	margin: 0 24px;
}
.tabbar{
	z-index: 999;
}
.order-options{
	display: block;
	margin: 15px 30px;
	display: flex; 
	justify-content: space-between;
}
.choice-item{
	margin: 5px 0;
}
.center-button {
  width: 90px;
  height: 90px;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 14px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
.time-picker{
	border: none;
	margin-top: -18px;
	margin-left: 180px;
}
.form-item4{
	width: 100%;
	margin: 3px 2px;
	opacity: 1;
	border-radius: 5px;
	background: rgba(255, 255, 255, 1);
	padding: 8px;
	margin-bottom: 8px;
}
.example-body {
	padding: 10px;
	padding-top: 0;
}

.custom-image-box {
	/* #ifndef APP-NVUE */
	display: flex;
	/* #endif */
	flex-direction: row;
	justify-content: space-between;
	align-items: center;
}

.text {
	font-size: 14px;
	color: #333;
}
input{
	opacity: 0.7;
}
</style>