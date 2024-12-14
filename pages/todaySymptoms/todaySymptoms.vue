<!-- 全身症状页面 -->
<script>
import store from '../../store/index.ts'
import { clockSubmitService, submitAllCheckService, getDailyReportService } from '../../services/clock'
export default {
  data() {
    return {
	  AllCheckInfos: {
		  
	  },
	  AllCheckInfo : {
		circulatorysymptoms: {
		    circulatorySymptomsId: 0,
		    hasArrhythmia: false,
		    hasChestPain: false,
		    hasLowBloodPressure: false,
		    hasPalpitation: false,
		    hasRapidPulse: false,
		    statusId: 0
		},
		digestivesymptoms: {
		    arrheaFrequencyGEThreeTimesPerDay: false,
		    diarrheaMode: null,
		    diarrheaTimesPerDay: '',
		    digestiveSymptomsId: 0,
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
		    hasVomiting: false,
		    statusId: 0,
		    stoolAmount: null,
		    stoolOdor: null,
		    stoolType1Detail: null,
		    stoolType2Detail: null,
		    vomitingMode: '先泻后吐',
		},
		generalsymptoms: {
		    dehydration: null,
		    feverDuration: '',
		    generalSymptomsId: 0,
		    hasBloodBlisters: false,
		    hasChills: false,
		    hasCongestiveOrPetechialRash: false,
		    hasCyanosis: false,
		    hasDehydration: false,
		    hasEdema: false,
		    hasExhaustion: false,
		    hasFatigue: false,
		    hasFever: false,
		    hasHeadache: false,
		    hasItchyRash: false,
		    hasJointPain: false,
		    hasLymphNodeSwelling: false,
		    hasLymphNodeSwellingArmpit: false,
		    hasLymphNodeSwellingElbow: false,
		    hasLymphNodeSwellingGroin: false,
		    hasLymphNodeSwellingNeck: false,
		    hasLymphNodeSwellingPopliteal: false,
		    hasLymphNodeSwellingSubclavian: false,
		    hasMusclePain: false,
		    hasNightSweats: false,
		    hasPainfulRedRash: false,
		    hasSkinUlcer: false,
		    hasSubcutaneousAndMucosalBleedingSpots: false,
		    hasSweating: false,
		    hasWeightLoss: false,
		    highestTemperature: '',
		    itchyRashOnBackOfHands: false,
		    itchyRashOnFace: false,
		    itchyRashOnFeet: false,
		    itchyRashOnFingers: false,
		    itchyRashOnLowerLimbs: false,
		    itchyRashOnOther: false,
		    itchyRashOnUpperLimbs: false,
		    statusId: 0
		},
		localSymptoms: {
		    hasConjunctivalCongestion: false,
		    hasConjunctivitis: false,
		    hasEyelidEdema: false,
		    hasFacialRednessOrPallor: false,
		    hasGumBleeding: false,
		    hasKidneyPain: false,
		    hasLipCyanosis: false,
		    hasLowBackPain: false,
		    hasMucosalBleedingInOralOrNasal: false,
		    hasOrbitalPain: false,
		    hasRednessNeckChest: false,
		    hasSkinBleedingPoints: false,
		    hasSkinPallorOrBruising: false,
		    hasSmellTasteLoss: false,
		    localSymptomsId: 0,
		    statusId: 0
		},
		neurologicalSymptoms: {
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
		    hasTinnitus: false,
		    neurologicalSymptomsId: 0,
		    statusId: 0
		},
		otherSymptoms: {
		    hasCalfMusclePain: false,
		    hasForcedPosture: false,
		    hasOtherSymptoms: false,
		    hasPeriodicAttack: false,
		    hasRapidProgress: false,
		    hasSuddenOnset: false,
		    nutritionStatus: '好',
		    otherSymptomsId: 0,
		    otherSymptomsName: '',
		    sleepQuality: '好',
		    statusId: 0,
		    symptomSeverity: '轻',
		    workLifeStress: '小',
		},
		respiratorysymptoms: {
			hasRunnyNose: false,
		    hasBloodySputum: false,
		    hasBreathingDifficulty: false,
		    hasBronchitis: false,
		    hasChestTightness: false,
		    hasCough: false,
		    hasDryThroat: false,
		    hasHemoptysis: false,
		    hasHoarseness: false,
		    hasNasalCongestion: false,
		    hasShortnessOfBreath: false,
		    hasSoreThroat: false,
		    hasSputum: false,
		    respiratorySymptomsId: 0,
		    statusId: 0
		},
		riskFactorsAndExposure: {
		    anthraxArea: '否',
		    contactDate: '',
		    contactPlace: '',
		    contactWithAnimalProducts: '否',
		    contactWithAnimals: '否',
		    contactWithDiarrheaPatient: '否',
		    contactWithFeverPatient: '否',
		    contactWithPatientExcreta: '否',
		    contactWithRashPatient: '否',
		    contactWithRatIn2Months: '否',
		    contactWithTuberculosisPatient: '否',
		    diningOut: false,
		    diningPlace: '',
		    diningUnclear: false,
		    drunkWaterFromDitchesOrPonds: '否',
		    eatenFoodContaminatedByRatFeces: '否',
		    fleaBite: '否',
		    foodNames: '',
		    groupDining: false,
		    groupOutbreak: '否',
		    hasContactWithBird: false,
		    hasContactWithCat: false,
		    hasContactWithCow: false,
		    hasContactWithDog: false,
		    hasContactWithDonkey: false,
		    hasContactWithFlea: false,
		    hasContactWithHorse: false,
		    hasContactWithLice: false,
		    hasContactWithMarmot: false,
		    hasContactWithMule: false,
		    hasContactWithOtherAnimal: false,
		    hasContactWithPig: false,
		    hasContactWithRabbit: false,
		    hasContactWithRat: false,
		    hasContactWithSheep: false,
		    hasContactWithTiger: false,
		    hasContactWithWolf: false,
		    hasDrinkingRawWater: false,
		    hasEatingColdCookedFood: false,
		    hasEatingRawFood: false,
		    hasEatingSeafood: false,
		    hasMosquitoBite: false,
		    hillyOrMountainous: false,
		    malariaArea: '否',
		    otherAnimalName: '',
		    otherResidenceType: false,
		    otherWildActivityName: '',
		    outdoorStayOrWorkWithinMonth: '否',
		    plagueArea: '否',
		    plain: false,
		    ratContactMethod: null,
		    ratOrRatDroppingsAtWorkplace: '否',
		    restedNearRatHoles: '否',
		    riskFactorsAndExposureId: 0,
		    specificAnthraxLocation: '',
		    specificMalariaLocation: '',
		    specificPlagueLocation: '',
		    statusId: 0,
		    tickBite: '否',
		    wildActivityHistory: '其他'
		}
	  },
	  formData: {
		option: '' // 用于存储选中的选项
	  },
	  temporaryChoice1: false,
	  temporaryChoice2: false,
	  otherChoices: false,
	  SeafoodChoices: false,
	  isViewMode: false, 
	  pages: 1,
	  sonPage: 1,
    };
  },
  onLoad(options) {
    // `options` 是页面加载时获取的传参对象
    const { date } = options;
	console.log('dedao',date,'那么isViewMode',this.isViewMode)
    if ( date) {
      this.isViewMode = true;
      // 调用获取服务方法以查询数据
      this.getDiagnosisData(date);
    }
  },
  methods: {
    async getDiagnosisData(date) {
      try {
        const response = await getDailyReportService(date);
        if (response && response.data) {
          this.AllCheckInfo = response.data;
        }
      } catch (error) {
        console.error('获取诊断数据失败:', error);
      }
    },
    toggleSymptom(symptom) {
      if (this.symptoms.includes(symptom)) {
        this.symptoms = this.symptoms.filter(s => s !== symptom);
      } else {
        this.symptoms.push(symptom);
      }
    },
    handleSubmit() {
      // 处理提交逻辑
      console.log('提交数据', {
        fever: this.fever,
        duration: this.selectedDuration,
        symptoms: this.symptoms,
      });
	  if(this.pages <8){
		this.pages++;
	  }else{
		  if(this.sonPage<=5){
			  this.sonPage++;
		  }
	  }
	  
	  // uni.navigateTo({
	  // 	url: '/pages/respiratorySymptoms/respiratorySymptoms'
	  // })
    },
	showRadio(maleValue){
		this.AllCheckInfo.generalsymptoms.dehydration = maleValue;
		console.log('已经选择',this.AllCheckInfo.generalsymptoms.dehydration)
	},
	handleCheckboxChange(propName) {
	  console.log(`Changing ${propName}:`, this.AllCheckInfo.generalsymptoms[propName]);
	  this.AllCheckInfo.generalsymptoms[propName] = !this.AllCheckInfo.generalsymptoms[propName];
	},
	handleRespiratorysymptomsBox(propName){
	  console.log(`Changing ${propName}:`, this.AllCheckInfo.respiratorysymptoms[propName]);
	  this.AllCheckInfo.respiratorysymptoms[propName] = !this.AllCheckInfo.respiratorysymptoms[propName];
	},
	handlereturn(){
		if(this.pages == 8 && this.sonPage >1){
			this.sonPage--;
		}else{
			this.pages--;
		}
	},
	submitSuccess1(response) {
		console.log(response)
	},
	handleFinallySubmit(){
		uni.showModal({
            title: '提示',
            content: '您确定要提交吗？',
            confirmText: '确定',
            cancelText: '取消',
            success: (res) => {
                if (res.confirm) {
                    const that = this; // 保存当前的 `this` 上下文
                    const clockForm = store.state.clockFormParams;
                    // 使用顺序执行的方式确保顺序执行
                    clockSubmitService(clockForm)
                        .then((res) => {
                            this.submitSuccess1(res);
            
                            // 在 `clockSubmitService` 完成后再执行 `submitAllCheckService`
                            return submitAllCheckService(that.AllCheckInfo);
                        })
                        .then((res) => {
                            that.submitSuccess(res);
                        })
                        .catch((error) => {
                            that.submitError(error);
                        });
                }
            }
        });
	},
	submitError(error) {
		uni.showToast({ 
			title: '信息提交失败',
			icon: 'none', // 可选，默认为none，还有success, loading, warning, error等
			duration: 1500, // 显示持续时间，单位为毫秒，默认为1500
		});
	},
	submitSuccess(response) {
		console.log(response);
		console.log(response.data);
		store.dispatch('setDiseasescoring', response.data);
	    if (response.code == 1) {
	        uni.showToast({
	          title: '信息提交成功',
	          icon: 'success',
			  duration:1800,
	        });
			setTimeout(() => {
				// 进行页面跳转
				uni.navigateTo({
					url: '/pages/clockFail/clockFail'
				});
        }, 50);
	    } else {
	        uni.showModal({
	            title: '信息提交失败',
	            content: '信息提交失败。',
	            showCancel: false,
	            confirmText: '确定',
	            success: function (res) {
	            }
	        });
	    }
	},
	handleDigestFrequency(propName){
		console.log(`Changing ${propName}:`, this.AllCheckInfo.digestivesymptoms[propName]);
		this.AllCheckInfo.digestivesymptoms[propName] = !this.AllCheckInfo.digestivesymptoms[propName];
	},
	showDigestRadio(maleValue){
		this.AllCheckInfo.digestivesymptoms.diarrheaMode = maleValue;
		console.log('已经选择',this.AllCheckInfo.digestivesymptoms.diarrheaMode)
	},
	showDigestRadioBystool1(maleValue){
		this.AllCheckInfo.digestivesymptoms.stoolType1Detail = maleValue;
		console.log('已经选择',this.AllCheckInfo.digestivesymptoms.stoolType1Detail)
	},
	showDigestRadioBystool2(maleValue){
		this.AllCheckInfo.digestivesymptoms.stoolType2Detail = maleValue;
		console.log('已经选择',this.AllCheckInfo.digestivesymptoms.stoolType2Detail)
	},
	showDigestVomitingMode(maleValue){
		this.AllCheckInfo.digestivesymptoms.vomitingMode = maleValue;
		console.log('已经选择',this.AllCheckInfo.digestivesymptoms.vomitingMode)
	},
	handleCirculatorysymptoms(propName){
		console.log(`Changing ${propName}:`, this.AllCheckInfo.circulatorysymptoms[propName]);
		this.AllCheckInfo.circulatorysymptoms[propName] = !this.AllCheckInfo.circulatorysymptoms[propName];
	},
	handleNeurologicalSymptoms(propName){
		console.log(`Changing ${propName}:`, this.AllCheckInfo.neurologicalSymptoms[propName]);
		this.AllCheckInfo.neurologicalSymptoms[propName] = !this.AllCheckInfo.neurologicalSymptoms[propName];
	},
	handleLocalSymptoms(propName){
		console.log(`Changing ${propName}:`, this.AllCheckInfo.localSymptoms[propName]);
		this.AllCheckInfo.localSymptoms[propName] = !this.AllCheckInfo.localSymptoms[propName];
	},
	handleOtherSymptoms(propName){
		console.log(`Changing ${propName}:`, this.AllCheckInfo.otherSymptoms[propName]);
		this.AllCheckInfo.otherSymptoms[propName] = !this.AllCheckInfo.otherSymptoms[propName];
	},
	showSleepQuality(maleValue){
		this.AllCheckInfo.otherSymptoms.sleepQuality = maleValue;
		console.log('已经选择',this.AllCheckInfo.otherSymptoms.sleepQuality)
	},
	showNutritionStatus(maleValue){
		this.AllCheckInfo.otherSymptoms.nutritionStatus = maleValue;
		console.log('已经选择',this.AllCheckInfo.otherSymptoms.nutritionStatus)
	},
	showWorkLifeStress(maleValue){
		this.AllCheckInfo.otherSymptoms.workLifeStress = maleValue;
		console.log('已经选择',this.AllCheckInfo.otherSymptoms.workLifeStress)
	},
	showSymptomSeverity(maleValue){
		this.AllCheckInfo.otherSymptoms.symptomSeverity = maleValue;
		console.log('已经选择',this.AllCheckInfo.otherSymptoms.symptomSeverity)
	},
	shwoContactWithFeverPatient(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.contactWithFeverPatient = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.contactWithFeverPatient)
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
		this.AllCheckInfo.riskFactorsAndExposure.contactWithTuberculosisPatient = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.contactWithTuberculosisPatient)
	},
	handleFactorSymptoms(propName){
		console.log(`Changing ${propName}:`, this.AllCheckInfo.riskFactorsAndExposure[propName]);
		this.AllCheckInfo.riskFactorsAndExposure[propName] = !this.AllCheckInfo.riskFactorsAndExposure[propName];
	},
	showGroupOutbreak(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.groupOutbreak = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.groupOutbreak)
	},
	showOutdoorStayOrWorkWithinMonth(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.outdoorStayOrWorkWithinMonth = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.outdoorStayOrWorkWithinMonth)
	},
	showWildActivityHistory(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.wildActivityHistory = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.wildActivityHistory)
	},
	showPlagueArea(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.plagueArea = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.plagueArea)
	},
	showAnthraxArea(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.anthraxArea = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.anthraxArea)
	},
	showMalariaArea(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.malariaArea = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.malariaArea)
	},
	showTickBite(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.tickBite = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.tickBite)
	},
	showFleaBite(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.fleaBite = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.fleaBite)
	},
	showContactWithRatIn2Months(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.contactWithRatIn2Months = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.contactWithRatIn2Months)
		if (this.AllCheckInfo.riskFactorsAndExposure.contactWithRatIn2Months == '否'){
			this.AllCheckInfo.riskFactorsAndExposure.ratContactMethod = null;
		}
	},
	showRatContactMethod(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.ratContactMethod = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.ratContactMethod)
	},
	showEatenFoodContaminatedByRatFeces(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.eatenFoodContaminatedByRatFeces = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.eatenFoodContaminatedByRatFeces)
	},
	showDrunkWaterFromDitchesOrPonds(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.drunkWaterFromDitchesOrPonds = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.drunkWaterFromDitchesOrPonds)
	},
	showRestedNearRatHoles(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.restedNearRatHoles = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.restedNearRatHoles)
	},
	showRatOrRatDroppingsAtWorkplace(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.ratOrRatDroppingsAtWorkplace = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.ratOrRatDroppingsAtWorkplace)
	},
	showContactWithPatientExcreta(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.contactWithPatientExcreta = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.contactWithPatientExcreta)
	},
	showOtherWildActivityName(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.otherWildActivityName = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.otherWildActivityName)
	},
	showContactWithAnimals(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.contactWithAnimals = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.contactWithAnimals)
		if (this.AllCheckInfo.riskFactorsAndExposure.contactWithAnimals === '否'){
			this.AllCheckInfo.riskFactorsAndExposure.hasContactWithRat = false;
			this.AllCheckInfo.riskFactorsAndExposure.hasContactWithLice = false;
			this.AllCheckInfo.riskFactorsAndExposure.hasContactWithFlea = false;
			this.AllCheckInfo.riskFactorsAndExposure.hasContactWithCat = false;
			this.AllCheckInfo.riskFactorsAndExposure.hasContactWithMarmot = false;
			this.AllCheckInfo.riskFactorsAndExposure.hasContactWithBird = false;
			this.AllCheckInfo.riskFactorsAndExposure.hasContactWithPig = false;
			this.AllCheckInfo.riskFactorsAndExposure.hasContactWithHorse = false;
			this.AllCheckInfo.riskFactorsAndExposure.hasContactWithMule = false;
			this.AllCheckInfo.riskFactorsAndExposure.hasContactWithDonkey = false;
			this.AllCheckInfo.riskFactorsAndExposure.hasContactWithCow = false;
			this.AllCheckInfo.riskFactorsAndExposure.hasContactWithSheep = false;
			this.AllCheckInfo.riskFactorsAndExposure.hasContactWithDog = false;
			this.AllCheckInfo.riskFactorsAndExposure.hasContactWithRabbit = false;
			this.AllCheckInfo.riskFactorsAndExposure.hasContactWithWolf = false;
			this.AllCheckInfo.riskFactorsAndExposure.hasContactWithTiger = false;
		}
	},
	showContactWithAnimalProducts(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.contactWithAnimalProducts = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.contactWithAnimalProducts)
	},
	showStoolOdor(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.stoolOdor = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.stoolOdor)
	},
	showStoolAmount(maleValue){
		this.AllCheckInfo.riskFactorsAndExposure.stoolAmount = maleValue;
		console.log('已经选择',this.AllCheckInfo.riskFactorsAndExposure.stoolAmount)
	},
	onDateChange(e) {
	  this.AllCheckInfo.riskFactorsAndExposure.contactDate = e.detail.value;
	},
  },
};
</script>
<template>
	<view>
	<view class="container" v-show="pages == 1">
	  <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
		<view class="content">
	      <view class="header">
	        <text class="title">全身</text>症状
	      </view>
		  <!-- 第一个 -->
	      <view class="section">
			<view class="sections1">
				<view class="section-title">
					<text class="label">发热情况</text>
				</view>
				<view class="row">
					<text class="label">是否发热</text>
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" @tap="handleCheckboxChange('hasFever')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasFever" color="#285AC8" >是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" @tap="handleCheckboxChange('hasFever')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasFever" color="#285AC8"  >否</radio>
						</label>
					</radio-group>
				</view>
			</view>
	        <view class="row">
	          <text class="label">最高温度</text>
			  <input :disabled="isViewMode" type="text" v-model="AllCheckInfo.generalsymptoms.highestTemperature" placeholder="请输入温度" style="width: 10vh;"/>
	          <text class="label">持续时间</text>
	          <input :disabled="isViewMode" type="text" v-model="AllCheckInfo.generalsymptoms.feverDuration" placeholder="请输入小时数" style="width: 10vh;"/>
	        </view>
	      </view>
	  
			<!-- 第二个 -->
	      <view class="section">
			  <view class="sections1">
				  <view class="section-title">
				  	<text class="label">请选择症状</text>
				  </view>
	        <view class="row">
	          <checkbox-group class="options">
				<view class="order-options">
	            <label class="checkbox1">
	              <checkbox :disabled="isViewMode" color="#285AC8" value="chills" @tap="handleCheckboxChange('hasChills')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasChills">寒战</checkbox>
	            </label>
	            <label class="checkbox1">
	              <checkbox :disabled="isViewMode" color="#285AC8" value="sweating"  @tap="handleCheckboxChange('hasSweating')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasSweating">大汗</checkbox>
	            </label>
	            <label class="checkbox1">
	              <checkbox :disabled="isViewMode" color="#285AC8" value="fatigue" @tap="handleCheckboxChange('hasFatigue')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasFatigue">乏力</checkbox>
	            </label>
	            <label class="checkbox1">
	              <checkbox :disabled="isViewMode" color="#285AC8" value="thirst" @tap="handleCheckboxChange('hasHeadache')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasHeadache">头痛</checkbox>
	            </label>
				</view>
				<view class="order-options" style="margin-right: 70px;">
				<label class="checkbox">
					<checkbox :disabled="isViewMode" color="#285AC8" value="masule" @tap="handleCheckboxChange('hasMusclePain')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasMusclePain">肌肉酸痛</checkbox>
				</label>
				<label class="checkbox">
					<checkbox :disabled="isViewMode" color="#285AC8" value="masule" @tap="handleCheckboxChange('hasJointPain')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasJointPain">关节酸痛</checkbox>
				</label>
				</view>
	          </checkbox-group>
	        </view>
			</view>
	      </view>
	  
			<!-- 第三个 -->
			<view class="section">
						<view class="sections1">
							<view class="section-title">
								<text class="label">皮肤、疹、疮、淋巴结情况</text>
							</view>
						</view>
						<view class="section" style="margin-bottom: -6px;">
						  <checkbox-group>
						    <label>
						      <checkbox :disabled="isViewMode" color="#285AC8" value="淋巴结肿大" @tap="handleCheckboxChange('hasLymphNodeSwelling')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasLymphNodeSwelling" />
						      淋巴结肿大
						    </label>
							<view class="order-options" style="margin-left: 30px;" v-if="AllCheckInfo.generalsymptoms.hasLymphNodeSwelling == true">
						    <label>
						      <checkbox :disabled="isViewMode" color="#285AC8" value="腹股沟" @tap="handleCheckboxChange('hasLymphNodeSwellingGroin')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasLymphNodeSwellingGroin" />
						      腹股沟
						    </label>
						    <label>
						      <checkbox :disabled="isViewMode" color="#285AC8" value="腋下" @tap="handleCheckboxChange('hasLymphNodeSwellingArmpit')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasLymphNodeSwellingArmpit"/>
						      腋下
						    </label>
						    <label>
						      <checkbox :disabled="isViewMode" color="#285AC8" value="锁骨下" @tap="handleCheckboxChange('hasLymphNodeSwellingSubclavian')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasLymphNodeSwellingSubclavian"/>
						      锁骨下
						    </label>
							</view>
							<view class="order-options" style="margin-left: 30px;" v-if="AllCheckInfo.generalsymptoms.hasLymphNodeSwelling == true">
						    <label>
						      <checkbox :disabled="isViewMode" color="#285AC8" value="颈部" @tap="handleCheckboxChange('hasLymphNodeSwellingNeck')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasLymphNodeSwellingNeck"/>
						      颈部
						    </label>
						    <label>
						      <checkbox :disabled="isViewMode" color="#285AC8" value="肘部" @tap="handleCheckboxChange('hasLymphNodeSwellingElbow')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasLymphNodeSwellingElbow"/>
						      肘部
						    </label>
						    <label>
						      <checkbox :disabled="isViewMode" color="#285AC8" value="腘窝" @tap="handleCheckboxChange('hasLymphNodeSwellingPopliteal')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasLymphNodeSwellingPopliteal"/>
						      腘窝
						    </label>
							</view>
						  </checkbox-group>
						</view>
						<view class="checkboxs" >
						      <checkbox-group>
								<view style="display: flex; flex-direction: column;margin-left: 12px;">
						        <label class="choice-item">
						          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('hasLipCyanosis')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasLipCyanosis"/>
						          口唇、颜面、四肢及全身皮肤发绀
						        </label>
						        <label class="choice-item">
						          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('hasSubcutaneousAndMucosalBleedingSpots')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasSubcutaneousAndMucosalBleedingSpots"/>
						          皮下及粘膜出血或出血点
						        </label>
						        <label class="choice-item">
						          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('hasPainfulRedRash')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasPainfulRedRash"/>
						          皮肤可见剧痛的红色丘疹
						        </label>
						        <label class="choice-item">
						          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('hasBloodBlisters')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasBloodBlisters"/>
						          皮肤可见周边呈灰黑色、基底坚硬的血性水泡
						        </label>
						        <label class="choice-item">
						          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('hasSkinUlcer')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasSkinUlcer"/>
						          皮肤出现呈灰黑色创面的溃疡
						        </label>
						        <label class="choice-item">
						          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('hasCongestiveOrPetechialRash')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasCongestiveOrPetechialRash"/>
						          皮疹：充血性或点状出血疹
						        </label>
						        <label class="choice-item">
						          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('hasPainfulRedRash')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasPainfulRedRash"/>
						          皮疹：起初压之褪色，后期压之不退，多见于四肢、躯干、面部一般没有
						        </label>
						        <label class="choice-item">
						          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('hasItchyRash')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.hasItchyRash"/>
						          瘙痒性丘疹/水疱
						        </label>
								</view>
						      </checkbox-group>
						    </view>
						    <view class="inner-checkboxs" v-if="AllCheckInfo.generalsymptoms.hasItchyRash == true">
						      <checkbox-group>
								<view class="order-options">
						        <label>
						          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('itchyRashOnFingers')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.itchyRashOnFingers"/>
						          手指
						        </label>
						        <label>
						          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('itchyRashOnBackOfHands')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.itchyRashOnBackOfHands"/>
						          手背
						        </label>
						        <label>
						          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('itchyRashOnUpperLimbs')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.itchyRashOnUpperLimbs"/>
						          上肢
						        </label>
								</view>
								<view class="order-options">
						        <label>
						          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('itchyRashOnLowerLimbs')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.itchyRashOnLowerLimbs"/>
						          下肢
						        </label>
						        <label>
						          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('itchyRashOnFeet')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.itchyRashOnFeet"/>
						          足背
						        </label>
						        <label>
						          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('itchyRashOnFace')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.itchyRashOnFace"/>
						          面部
						        </label>
								</view>
								<view class="order-options">
						        <label>
						          <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCheckboxChange('itchyRashOnOther')" :checked="!isViewMode ? false : AllCheckInfo.generalsymptoms.itchyRashOnOther"/>
						          其他
						        </label>
								</view>
						      </checkbox-group>
						    </view>
			</view>
	      <view class="section">
			  <view class="sections1">
				  <view class="section-title">
				  	<text class="label">失水情况</text>
				  </view>
				  <view class="row">
				  	<text class="label">失水</text>
				  	<radio-group class="options">
				  		<label class="radio">
				  			<radio :disabled="isViewMode" color="#285AC8" value="轻" v-model="AllCheckInfo.generalsymptoms.dehydration" :checked="!isViewMode ? false :AllCheckInfo.generalsymptoms.dehydration === '轻度'" @tap="showRadio('轻度')">轻度</radio>
				  		</label>
				  		<label class="radio">
				  			<radio :disabled="isViewMode" color="#285AC8" value="中" v-model="AllCheckInfo.generalsymptoms.dehydration" :checked="!isViewMode ? false :AllCheckInfo.generalsymptoms.dehydration === '中度'" @tap="showRadio('中度')">中度</radio>
				  		</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" value="重" v-model="AllCheckInfo.generalsymptoms.dehydration" :checked="!isViewMode ? false :AllCheckInfo.generalsymptoms.dehydration === '重度'" @tap="showRadio('重度')">重度</radio>
						</label>
				  	</radio-group>
				  </view>
			  </view>
		  </view>
	      <button class="single-submit-button" @tap="handleSubmit">下一步</button>
		  </view>
		<uni-custom-tabbar></uni-custom-tabbar>
	</view>
	<view class="container" v-show="pages == 2">
	  <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
		<view class="content">
	      <view class="header">
	        <text class="title">呼吸</text>系统症状
	      </view>
	      <view class="section">
			  <view class="sections1">
				  <view class="section-title">
				  	<text class="label">请选择症状</text>
				  </view>
	        <view class="row">
	          <checkbox-group class="options">
				<view class="order-options" style="margin-bottom: -5px;">
	            <label class="checkbox1">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasCough')" :checked="!isViewMode ? false : AllCheckInfo.respiratorysymptoms.hasCough">咳嗽</checkbox>
	            </label>
	            <label class="checkbox1">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasSputum')" :checked="!isViewMode ? false : AllCheckInfo.respiratorysymptoms.hasSputum" >咳痰</checkbox>
	            </label>
	            <label class="checkbox1">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasBloodySputum')" :checked="!isViewMode ? false : AllCheckInfo.respiratorysymptoms.hasBloodySputum">血痰</checkbox>
	            </label>
	            <label class="checkbox1">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasHemoptysis')" :checked="!isViewMode ? false : AllCheckInfo.respiratorysymptoms.hasHemoptysis">咳血</checkbox>
	            </label>
				</view>
				<view class="order-options" style="margin-bottom: -5px;">
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasSoreThroat')" :checked="!isViewMode ? false : AllCheckInfo.respiratorysymptoms.hasSoreThroat">咽痛</checkbox>
				</label>
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8"  @tap="handleRespiratorysymptomsBox('hasDryThroat')" :checked="!isViewMode ? false : AllCheckInfo.respiratorysymptoms.hasDryThroat">咽干</checkbox>
				</label>
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasNasalCongestion')" :checked="!isViewMode ? false : AllCheckInfo.respiratorysymptoms.hasNasalCongestion">鼻塞</checkbox>
				</label>
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasRunnyNose')" :checked="!isViewMode ? false : AllCheckInfo.respiratorysymptoms.hasRunnyNose">流涕</checkbox>
				</label>
				</view>
				<view class="order-options" style="margin-bottom: -5px; margin-right: 142px;">
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasChestTightness')" :checked="!isViewMode ? false : AllCheckInfo.respiratorysymptoms.hasChestTightness">胸闷</checkbox>
				</label>
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasShortnessOfBreath')" :checked="!isViewMode ? false : AllCheckInfo.respiratorysymptoms.hasShortnessOfBreath">气促</checkbox>
				</label>
				</view>
				<view class="order-options" style="margin-bottom: -5px; margin-right: 85px; ">
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasBreathingDifficulty')" :checked="!isViewMode ? false : AllCheckInfo.respiratorysymptoms.hasBreathingDifficulty">呼吸困难</checkbox>
				</label>
				<label class="checkbox1" >
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasHoarseness')" :checked="!isViewMode ? false : AllCheckInfo.respiratorysymptoms.hasHoarseness">声音嘶哑</checkbox>
				</label>
				</view>
				<view class="order-options" style="margin-bottom: -5px; margin-right: 110px;">
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleRespiratorysymptomsBox('hasBronchitis')" :checked="!isViewMode ? false : AllCheckInfo.respiratorysymptoms.hasBronchitis">支气管炎/支气管肺炎</checkbox>
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
	<view class="container" v-show="pages == 3">
	  <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
		<view class="content">
	    <view class="header">
	      <text class="title">消化</text>系统症状
	    </view>
		
		<!-- 第一个 -->
		<view class="section">
					<view class="sections1">
						<view class="section-title">
							<text class="label">腹泻情况</text>
						</view>
						<view class="row">
							<text class="label">腹泻次数</text>
							<radio-group class="options">
								<label class="radio">
									<radio :disabled="isViewMode" color="#285AC8" @tap="handleDigestFrequency('arrheaFrequencyGEThreeTimesPerDay')" :checked="!isViewMode ? false : !AllCheckInfo.digestivesymptoms.arrheaFrequencyGEThreeTimesPerDay" >少于3次/天</radio>
								</label>
								<label class="radio">
									<radio :disabled="isViewMode" color="#285AC8" @tap="handleDigestFrequency('arrheaFrequencyGEThreeTimesPerDay')" :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.arrheaFrequencyGEThreeTimesPerDay">大于等于3次/天</radio>
								</label>
							</radio-group>
						</view>
						<input type="number" placeholder="请输入具体次数" v-model="AllCheckInfo.digestivesymptoms.diarrheaTimesPerDay" v-show="!AllCheckInfo.digestivesymptoms.arrheaFrequencyGEThreeTimesPerDay"/>
						<view class="row">
							<text class="label">腹泻方式</text>
							<radio-group class="options">
								<view class="option-button">
								<label class="radio">
									<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadio('里急后重')"  :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.diarrheaMode == '里急后重'">里急后重</radio>
								</label>
								<label class="radio" style="margin-top: 10px;">
									<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadio('失禁')"  :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.diarrheaMode == '失禁'">失禁</radio>
								</label>
								</view>
								<view class="option-button">
								<label class="radio">
									<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadio('畅通')"  :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.diarrheaMode == '畅通'">畅通</radio>
								</label>
								<label class="radio" style="margin-top: 10px;">
									<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadio('绞痛')"  :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.diarrheaMode == '绞痛'">绞痛</radio>
								</label>
								</view>
							</radio-group>
						</view>
					</view>
		</view>
		<!-- 第二个 -->
		<view class="section">
					<view class="sections1">
						<view class="section-title">
							<text class="label">粪便情况</text>
						</view>
						<view class="row">
							<text class="label">粪便量</text>
							<radio-group class="options">
								<label class="radio">
									<radio :disabled="isViewMode" color="#285AC8" @tap="showStoolAmount('多')" :checked="!isViewMode ? false: AllCheckInfo.riskFactorsAndExposure.stoolAmount == '多'">多</radio>
								</label>
								<label class="radio">
									<radio :disabled="isViewMode" color="#285AC8" @tap="showStoolAmount('少')" :checked="!isViewMode ? false: AllCheckInfo.riskFactorsAndExposure.stoolAmount == '少'">>少</radio>
								</label>
							</radio-group>
						</view>
						<view class="row">
							<text class="label">粪便气味</text>
							<radio-group class="options">
								<label class="radio">
									<radio :disabled="isViewMode" color="#285AC8" @tap="showStoolOdor('恶臭')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.stoolOdor  == '恶臭'">恶臭</radio>
								</label>
								<label class="radio">
									<radio :disabled="isViewMode" color="#285AC8" @tap="showStoolOdor('无恶臭')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.stoolOdor  == '无恶臭'">无恶臭</radio>
								</label>
							</radio-group>
						</view>
						<view class="checkboxs" >
						      <checkbox-group>
								<view style="display: flex; flex-direction: column;">
						        <label @tap="temporaryChoice1 = !temporaryChoice1">
	  					          <checkbox :checked="temporaryChoice1 == true" value="粪便性状1" color="#285AC8" :disabled="temporaryChoice2 == true"/>
						          粪便性状1
						        </label>
								<radio-group class="options" v-show="temporaryChoice1">
									<label class="radio">
										<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadioBystool1('水样')" :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.stoolType1Detail == '水样'">水样</radio>
									</label>
									<label class="radio">
										<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadioBystool1('米泔样')" :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.stoolType1Detail == '米泔样'">米泔样</radio>
									</label>
									<label class="radio">
										<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadioBystool1('洗肉水样')" :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.stoolType1Detail == '洗肉水样'">洗肉水样</radio>
									</label>
								</radio-group>
						        <label style="margin-top: 10px;" @tap="temporaryChoice2 = !temporaryChoice2">
								  <checkbox :checked="temporaryChoice2 == true" value="粪便性状2" color="#285AC8" :disabled="temporaryChoice1 == true" />
						          粪便性状2
						        </label>
								<radio-group class="options" v-show="temporaryChoice2">
									<label class="radio">
										<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadioBystool2('黑便')" :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.stoolType2Detail == '黑便'">黑便</radio>
									</label>
									<label class="radio">
										<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadioBystool2('脓血')" :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.stoolType2Detail == '脓血'">脓血</radio>
									</label>
									<label class="radio">
										<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestRadioBystool2('大块粘膜')" :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.stoolType2Detail == '大块粘膜'">大块粘膜</radio>
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
						<view class="section-title">
							<text class="label">呕吐情况</text>
						</view>
						<view class="row">
							<text class="label">是否呕吐</text>
							<radio-group class="options">
								<label class="radio">
									<radio :disabled="isViewMode" color="#285AC8" @tap="handleDigestFrequency('hasVomiting')" :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.hasVomiting">是</radio>
								</label>
								<label class="radio">
									<radio :disabled="isViewMode" color="#285AC8" @tap="handleDigestFrequency('hasVomiting')" :checked="!isViewMode ? false : !AllCheckInfo.digestivesymptoms.hasVomiting">否</radio>
								</label>
							</radio-group>
						</view>
						<view class="row">
							<text class="label">呕吐方式</text>
							<radio-group class="options">
								<view class="option-button" style="margin-top: 20px;">
								<label class="radio">
									<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestVomitingMode('喷射状')" :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.vomitingMode == '喷射状'">喷射状</radio>
								</label>
								<label class="radio" style="margin-top: 10px;">
									<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestVomitingMode('先泻后吐')" :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.vomitingMode == '先泻后吐'">先泄后吐</radio>
								</label>
								</view>
								<view class="option-button" style="margin-top: 20px;">
								<label class="radio">
									<radio :disabled="isViewMode" color="#285AC8" @tap="showDigestVomitingMode('先吐后泻')" :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.vomitingMode == '先吐后泻'">先吐后泄</radio>
								</label>
								</view>
							</radio-group>
						</view>
					</view>
		</view>
			<!-- 第三个 -->
			<view class="section">
						<view class="sections1">
							<view class="section-title">
								<text class="label">请选择其他消化系统症状</text>
							</view>
						</view>
						<view class="checkboxs" style="margin-top: 10px;" >
						      <checkbox-group>
								<view class="order-options" style="margin-left: 60px;">
						        <label>
						          <checkbox :disabled="isViewMode" color="#285AC8" value="恶心" @tap="handleDigestFrequency('hasNausea')" :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.hasNausea"/>
						          恶心
						        </label>
						        <label>
						          <checkbox :disabled="isViewMode" color="#285AC8" value="腹胀" @tap="handleDigestFrequency('hasAbdominalDistension')" :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.hasAbdominalDistension"/>
						          腹胀
						        </label>
						        <label>
						          <checkbox :disabled="isViewMode" color="#285AC8" value="上腹部不适" @tap="handleDigestFrequency('hasUpperAbdominalDiscomfort')" :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.hasUpperAbdominalDiscomfort" />
						          上腹部不适
						        </label>
								</view>
								<view class="order-options" style="margin-left: 60px;">
						        <label>
						          <checkbox :disabled="isViewMode" color="#285AC8" value="腹痛" @tap="handleDigestFrequency('hasAbdominalPain')" :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.hasAbdominalPain"/>
						          腹痛
						        </label>
						        <label>
						          <checkbox :disabled="isViewMode" color="#285AC8" value="腹鸣" @tap="handleDigestFrequency('hasBorborygmus')" :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.hasBorborygmus"/>
						          腹鸣
						        </label>
						        <label>
						          <checkbox :disabled="isViewMode" color="#285AC8" value="便秘"  @tap="handleDigestFrequency('hasConstipation')" :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.hasConstipation"/>
						          便秘
						        </label>
								</view>
								<view class="order-options" style="margin-left: 60px;">
						        <label>
						          <checkbox :disabled="isViewMode" color="#285AC8" value="少尿或无尿" @tap="handleDigestFrequency('hasOliguriaOrAnuria')" :checked="!isViewMode ? false: AllCheckInfo.digestivesymptoms.hasOliguriaOrAnuria"/>
						          少尿或无尿
						        </label>
						        <label>
						          <checkbox :disabled="isViewMode" color="#285AC8" value="食欲减退" @tap="handleDigestFrequency('hasAppetiteLoss')" :checked="!isViewMode ? false : AllCheckInfo.digestivesymptoms.hasAppetiteLoss"/>
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
	<view class="container" v-show="pages == 4">
	  <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
		<view class="content">
	      <view class="header">
	        <text class="title">循环</text>系统症状
	      </view>
	      <view class="section">
			  <view class="sections1">
				  <view class="section-title">
				  	<text class="label">请选择症状</text>
				  </view>
	        <view class="row">
	          <checkbox-group class="options">
				<view class="order-options" >
	            <label class="checkbox">
	              <checkbox  :disabled="isViewMode" color="#285AC8" @tap="handleCirculatorysymptoms('hasArrhythmia')" :checked="!isViewMode ? false : AllCheckInfo.circulatorysymptoms.hasArrhythmia">心律不齐</checkbox>
	            </label>
	            <label class="checkbox">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCirculatorysymptoms('hasChestPain')" :checked="!isViewMode ? false : AllCheckInfo.circulatorysymptoms.hasChestPain">胸痛</checkbox>
	            </label>
	            <label class="checkbox">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCirculatorysymptoms('hasRapidPulse')" :checked="!isViewMode ? false : AllCheckInfo.circulatorysymptoms.hasRapidPulse">脉搏细速</checkbox>
	            </label>
				</view>
				<view class="order-options" style="margin-right: 130px;">
	            <label class="checkbox">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCirculatorysymptoms('hasPalpitation')" :checked="!isViewMode ? false : AllCheckInfo.circulatorysymptoms.hasPalpitation">心悸</checkbox>
	            </label>
				<label class="checkbox">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleCirculatorysymptoms('hasLowBloodPressure')" :checked="!isViewMode ? false : AllCheckInfo.circulatorysymptoms.hasLowBloodPressure">低血压</checkbox>
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
		<view class="content">
	      <view class="header">
	        <text class="title">神经</text>系统症状
	      </view>
	      <view class="section">
			  <view class="sections1">
				  <view class="section-title">
				  	<text class="label">请选择症状</text>
				  </view>
	        <view class="row">
	          <checkbox-group class="options">
				<view class="order-options">
	            <label class="checkbox1">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleNeurologicalSymptoms('hasDelirium')" :checked="!isViewMode ? false : AllCheckInfo.neurologicalSymptoms.hasDelirium">儋妄</checkbox>
	            </label>
	            <label class="checkbox1">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleNeurologicalSymptoms('hasConvulsion')" :checked="!isViewMode ? false : AllCheckInfo.neurologicalSymptoms.hasConvulsion">抽搐</checkbox>
	            </label>
	            <label class="checkbox1">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleNeurologicalSymptoms('hasDrowsiness')" :checked="!isViewMode ? false : AllCheckInfo.neurologicalSymptoms.hasDrowsiness">昏睡</checkbox>
	            </label>
	            <label class="checkbox1">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleNeurologicalSymptoms('hasComa')" :checked="!isViewMode ? false : AllCheckInfo.neurologicalSymptoms.hasComa">昏迷</checkbox>
	            </label>
				</view>
				<view class="order-options" style="margin-right: 91px;">
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleNeurologicalSymptoms('hasStiffNeck')" :checked="!isViewMode ? false : AllCheckInfo.neurologicalSymptoms.hasStiffNeck">颈部强直</checkbox>
				</label>
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleNeurologicalSymptoms('hasRestlessness')" :checked="!isViewMode ? false : AllCheckInfo.neurologicalSymptoms.hasRestlessness">烦躁不安</checkbox>
				</label>
				</view>
				<view class="order-options" style="margin-right: 106px;">
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleNeurologicalSymptoms('hasMuscleParalysisNeckShoulder')" :checked="!isViewMode ? false : AllCheckInfo.neurologicalSymptoms.hasMuscleParalysisNeckShoulder">颈肩部肌肉迟缓性瘫痪</checkbox>
				</label>
				</view>
				<view class="order-options" style="margin-right: 120px;">
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleNeurologicalSymptoms('hasMuscleParalysisLimbs')" :checked="!isViewMode ? false : AllCheckInfo.neurologicalSymptoms.hasMuscleParalysisLimbs">肢体肌肉迟缓性瘫痪</checkbox>
				</label>
				</view>
				<view class="order-options"  style="margin-right: 91px;">
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleNeurologicalSymptoms('hasSwallowingDifficulty')" :checked="!isViewMode ? false : AllCheckInfo.neurologicalSymptoms.hasSwallowingDifficulty">吞咽困难</checkbox>
				</label>
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleNeurologicalSymptoms('hasSpeechDisorder')" :checked="!isViewMode ? false : AllCheckInfo.neurologicalSymptoms.hasSpeechDisorder">语言障碍</checkbox>
				</label>
				</view>
				<view class="order-options" style="margin-right: 10px;">
				<label class="checkbox1" >
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleNeurologicalSymptoms('hasConsciousnessDisorder')" :checked="!isViewMode ? false : AllCheckInfo.neurologicalSymptoms.hasConsciousnessDisorder">意识障碍或惊厥</checkbox>
				</label>
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleNeurologicalSymptoms('hasDizziness')" :checked="!isViewMode ? false : AllCheckInfo.neurologicalSymptoms.hasDizziness">头晕</checkbox>
				</label>
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleNeurologicalSymptoms('hasTinnitus')" :checked="!isViewMode ? false : AllCheckInfo.neurologicalSymptoms.hasTinnitus">耳鸣</checkbox>
				</label>
				</view>
				<view class="order-options"  style="margin-right: 190px;">
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleNeurologicalSymptoms('hasHearingLoss')" :checked="!isViewMode ? false : AllCheckInfo.neurologicalSymptoms.hasHearingLoss">听力减退</checkbox>
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
	<view class="container" v-show="pages == 6">
	  <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
		<view class="content">
	      <view class="header">
	        <text class="title">局部</text>症状
	      </view>
	      <view class="section">
			  <view class="sections1">
				  <view class="section-title">
				  	<text class="label">请选择症状</text>
				  </view>
	        <view class="row">
	          <checkbox-group class="options" >
				<view class="order-options" style="margin-right: 50px;">
	            <label class="checkbox1">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasConjunctivitis')" :checked="!isViewMode ? false :AllCheckInfo.localSymptoms.hasConjunctivitis">结膜炎</checkbox>
	            </label>
	            <label class="checkbox1">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasConjunctivalCongestion')" :checked="!isViewMode ? false : AllCheckInfo.localSymptoms.hasConjunctivalCongestion">眼结膜充血、肿胀疼痛</checkbox>
	            </label>
				</view>
				<view class="order-options" style="margin-right: 25px;">
	            <label class="checkbox1">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasEyelidEdema')" :checked="!isViewMode ? false : AllCheckInfo.localSymptoms.hasEyelidEdema">上下眼睑水肿</checkbox>
	            </label>
	            <label class="checkbox1">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasSmellTasteLoss')" :checked="!isViewMode ? false : AllCheckInfo.localSymptoms.hasSmellTasteLoss">嗅觉味觉减退或丧失</checkbox>
	            </label>
				</view>
				<view class="order-options" style="margin-right: 10px;">
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasFacialRednessOrPallor')" :checked="!isViewMode ? false : AllCheckInfo.localSymptoms.hasFacialRednessOrPallor">颜面潮红或苍白</checkbox>
				</label>
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasRednessNeckChest')" :checked="!isViewMode ? false : AllCheckInfo.localSymptoms.hasRednessNeckChest">潮红（颈部、胸部）</checkbox>
				</label>
				</view>
				<view class="order-options" style="margin-right: 65px;">
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasLipCyanosis')" :checked="!isViewMode ? false : AllCheckInfo.localSymptoms.hasLipCyanosis">口唇发绀</checkbox>
				</label>
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasSkinPallorOrBruising')" :checked="!isViewMode ? false : AllCheckInfo.localSymptoms.hasSkinPallorOrBruising">皮肤苍白或带青紫</checkbox>
				</label>
				</view>
				<view class="order-options" style="margin-right: 120px;">
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasGumBleeding')" :checked="!isViewMode ? false : AllCheckInfo.localSymptoms.hasGumBleeding">牙龈出血</checkbox>
				</label>
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasKidneyPain')" :checked="!isViewMode ? false : AllCheckInfo.localSymptoms.hasKidneyPain">肾区疼痛</checkbox>
				</label>
				</view>
				<view class="order-options" style="margin-right: 160px;">
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasLowBackPain')" :checked="!isViewMode ? false : AllCheckInfo.localSymptoms.hasLowBackPain">腰痛</checkbox>
				</label>
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasOrbitalPain')" :checked="!isViewMode ? false : AllCheckInfo.localSymptoms.hasOrbitalPain">眼眶痛</checkbox>
				</label>
				</view>
				<view class="order-options" style="margin-right: 105px;">
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasMucosalBleedingInOralOrNasal')" :checked="!isViewMode ? false : AllCheckInfo.localSymptoms.hasMucosalBleedingInOralOrNasal">口腔、鼻腔等处黏膜有出血点</checkbox>
				</label>
				</view>
				<view class="order-options" style="margin-right: 35px;" >
				<label class="checkbox1">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleLocalSymptoms('hasSkinBleedingPoints')" :checked="!isViewMode ? false : AllCheckInfo.localSymptoms.hasSkinBleedingPoints">腋下/上臂/胸部或其他部位有皮肤出血点</checkbox>
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
		<view class="content">
	      <view class="header">
	        <text class="title">其他</text>
	      </view>
	      <view class="section">
			  <view class="sections1">
				  <view class="section-title">
				  	<text class="label">发病情况</text>
				  </view>
	        <view class="row">
	          <checkbox-group class="options">
				<view class="order-options" style="margin-right: 70px;">
	            <label class="checkbox">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleOtherSymptoms('hasSuddenOnset')" :checked="!isViewMode ? false : AllCheckInfo.otherSymptoms.hasSuddenOnset">突然发病</checkbox>
	            </label>
	            <label class="checkbox">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleOtherSymptoms('hasRapidProgress')" :checked="!isViewMode ? false : AllCheckInfo.otherSymptoms.hasRapidProgress">病情进展迅速</checkbox>
	            </label>
				</view>
				<view class="order-options" style="margin-right: 80px;">
	            <label class="checkbox">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleOtherSymptoms('hasPeriodicAttack')" :checked="!isViewMode ? false : AllCheckInfo.otherSymptoms.hasPeriodicAttack">周期性发作</checkbox>
	            </label>
	            <label class="checkbox">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleOtherSymptoms('hasForcedPosture')" :checked="!isViewMode ? false : AllCheckInfo.otherSymptoms.hasForcedPosture">强迫体位</checkbox>
	            </label>
				</view>
				<view class="order-options" style="margin-right: 200px;">
				<label class="checkbox">
					<checkbox :disabled="isViewMode" color="#285AC8" @tap="handleOtherSymptoms('hasCalfMusclePain')" :checked="!isViewMode ? false : AllCheckInfo.otherSymptoms.hasCalfMusclePain">腓肠肌疼痛</checkbox>
				</label>
				</view>
	          </checkbox-group>
	        </view>
			</view>
			
			<view class="sections1">
				<view class="section-title">
					<text class="label">生活情况</text>
				</view>
				<view class="row">
					<text class="label">睡眠状况</text>
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showSleepQuality('好')" :checked="!isViewMode ? false : AllCheckInfo.otherSymptoms.sleepQuality == '好'">好</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showSleepQuality('一般')" :checked="!isViewMode ? false : AllCheckInfo.otherSymptoms.sleepQuality == '一般'">一般</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showSleepQuality('差')" :checked="!isViewMode ? false : AllCheckInfo.otherSymptoms.sleepQuality == '差'">差</radio>
						</label>
					</radio-group>
				</view>
				<view class="row">
					<text class="label">营养状况</text>
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showNutritionStatus('好')" :checked="!isViewMode ? false : AllCheckInfo.otherSymptoms.nutritionStatus == '好'">好</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showNutritionStatus('一般')" :checked="!isViewMode ? false : AllCheckInfo.otherSymptoms.nutritionStatus == '一般'">一般</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showNutritionStatus('差')" :checked="!isViewMode ? false : AllCheckInfo.otherSymptoms.nutritionStatus == '差'">差</radio>
						</label>
					</radio-group>
				</view>
				<view class="row">
					<text class="label">工作和生活压力</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showWorkLifeStress('大')" :checked="!isViewMode ? false : AllCheckInfo.otherSymptoms.workLifeStress == '大'">大</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showWorkLifeStress('一般')" :checked="!isViewMode ? false : AllCheckInfo.otherSymptoms.workLifeStress == '一般'">一般</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showWorkLifeStress('小')" :checked="!isViewMode ? false : AllCheckInfo.otherSymptoms.workLifeStress == '小'">小</radio>
						</label>
					</radio-group>
				</view>
				<view class="row">
					<text class="label" style="width: 100px;">自我感觉症状的严重程度</text>
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showSymptomSeverity('轻')" :checked="!isViewMode ? false : AllCheckInfo.otherSymptoms.symptomSeverity == '轻'">轻</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showSymptomSeverity('中')" :checked="!isViewMode ? false : AllCheckInfo.otherSymptoms.symptomSeverity == '中'">中</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showSymptomSeverity('重')" :checked="!isViewMode ? false : AllCheckInfo.otherSymptoms.symptomSeverity == '重'">重</radio>
						</label>
					</radio-group>
				</view>
				<checkbox :disabled="isViewMode" color="#285AC8" style="margin-bottom: 10px;" @tap="otherChoices = !otherChoices" :checked="!isViewMode ? false : AllCheckInfo.otherSymptoms.otherSymptomsName != ''">其他</checkbox>
				<input type="text" placeholder="请填写其他症状" v-show="otherChoices" v-model="AllCheckInfo.otherSymptoms.otherSymptomsName"/>
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
		<view class="content">
	      <view class="header">
	        <text class="title">危险</text>因素与暴露史
	      </view>
	      <view class="section">
			  <!-- part1 -->
			<view class="sections1" v-show="sonPage==1">
				<view class="section-title">
					<text class="label">与其他患者接触史</text>
				</view>
				<view class="row">
					<text class="label">是否接触过发热或有呼吸道症状的患者</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="shwoContactWithFeverPatient('是')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.contactWithFeverPatient == '是'">是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="shwoContactWithFeverPatient('否')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.contactWithFeverPatient == '否'">否</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="shwoContactWithFeverPatient('不清楚')" :checked="!isViewMode ? false :AllCheckInfo.riskFactorsAndExposure.contactWithFeverPatient == '不清楚'">不清楚</radio>
						</label>
					</radio-group>
				</view>
				<view class="row">
					<text class="label">是否接触过有腹泻症状的患者</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showContactWithDiarrheaPatient('是')"  :checked="!isViewMode ? false :AllCheckInfo.riskFactorsAndExposure.contactWithDiarrheaPatient== '是'">是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showContactWithDiarrheaPatient('否')" :checked="!isViewMode ? false :AllCheckInfo.riskFactorsAndExposure.contactWithDiarrheaPatient== '否'">否</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showContactWithDiarrheaPatient('不清楚')" :checked="!isViewMode ? false :AllCheckInfo.riskFactorsAndExposure.contactWithDiarrheaPatient== '不清楚'">不清楚</radio>
						</label>
					</radio-group>
				</view>
				<view class="row">
					<text class="label">是否接触过有皮疹的患者</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showContactWithRashPatient('是')" :checked="!isViewMode ? false :AllCheckInfo.riskFactorsAndExposure.contactWithRashPatient== '是'">是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showContactWithRashPatient('否')" :checked="!isViewMode ? false :AllCheckInfo.riskFactorsAndExposure.contactWithRashPatient== '否'">否</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showContactWithRashPatient('不清楚')" :checked="!isViewMode ? false :AllCheckInfo.riskFactorsAndExposure.contactWithRashPatient== '不清楚'">不清楚</radio>
						</label>
					</radio-group>
				</view>
				<view class="row">
					<text class="label">是否接触过诊断为结核病的患者</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showContactWithTuberculosisPatient('是')" :checked="!isViewMode ? false :AllCheckInfo.riskFactorsAndExposure.contactWithTuberculosisPatient== '是'">是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showContactWithTuberculosisPatient('否')" :checked="!isViewMode ? false :AllCheckInfo.riskFactorsAndExposure.contactWithTuberculosisPatient== '否'">否</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showContactWithTuberculosisPatient('不清楚')" :checked="!isViewMode ? false :AllCheckInfo.riskFactorsAndExposure.contactWithTuberculosisPatient== '不清楚'">不清楚</radio>
						</label>
					</radio-group>
				</view>
			</view>
			
			<!-- part2 -->
			<view class="sections1" v-show="sonPage==2">
				<view class="section-title">
					<text class="label">就餐情况（五天内）</text>
				</view>
				<view class="row">
				<checkbox-group class="options">
				<view class="order-options" style="margin-right: 60px;">
				<label class="checkbox">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('groupDining')" :checked="!isViewMode ? false: AllCheckInfo.riskFactorsAndExposure.groupDining">集体堂食就餐</checkbox>
	            </label>
	            <label class="checkbox">
	              <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('diningOut')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.diningOut">外出就餐</checkbox>
	            </label>
				</view>
				<view class="order-options" style="margin-right: 213px;">
				<label class="checkbox">
				  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('diningUnclear')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.diningUnclear">不清楚</checkbox>
				</label>
				</view>
				</checkbox-group>
				</view>
				
				<view class="section-title">
					<text class="label">食物情况（五天内）</text>
				</view>
				<view class="row">
				<checkbox-group class="options">
					<label class="checkbox">
				  <checkbox :disabled="isViewMode" color="#285AC8"  @tap="handleFactorSymptoms('hasDrinkingRawWater')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasDrinkingRawWater">饮生水</checkbox>
				</label>
				<label class="checkbox">
				  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasEatingRawFood')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasEatingRawFood">吃生冷食品</checkbox>
				</label>
				<label class="checkbox1">
				  <checkbox :disabled="isViewMode" color="#285AC8"  @tap="handleFactorSymptoms('hasEatingColdCookedFood')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasEatingColdCookedFood">熟食冷吃</checkbox>
				</label>
				</checkbox-group>
				</view>
				<checkbox-group class="options" >
					<label class="checkbox1" style="margin-right: 225px;">
				       <checkbox :disabled="isViewMode" color="#285AC8" @tap="SeafoodChoices = !SeafoodChoices">海水产品</checkbox>
					</label>
				</checkbox-group>
				<view v-show="SeafoodChoices" class="row" style="margin-left: 50px;">
					食物名称<input type="text" placeholder="请输入食物名称" v-model="AllCheckInfo.riskFactorsAndExposure.foodNames"  :disabled="isViewMode"/>
				</view>
				
				<view class="section-title" style="margin: 10px 0;">
					<text class="label">就餐地点（五天内）</text>
				</view>
				<view class="row" style="margin-top: 10px; margin-left: 10px;">
				就餐地点<input :disabled="isViewMode" type="text" placeholder="请输入就餐地点" v-model="AllCheckInfo.riskFactorsAndExposure.diningPlace"/>
				</view>
				
				<view class="row">
					<text class="label">同一家庭、办公室、车间等集体单位是否有聚集性发病</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showGroupOutbreak('是')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.groupOutbreak == '是'">是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showGroupOutbreak('否')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.groupOutbreak == '否'">否</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showGroupOutbreak('不清楚')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.groupOutbreak == '不清楚'">不清楚</radio>
						</label>
					</radio-group>
				</view>
				<view class="row">
					<text class="label">发病前1月内是否在野外/户外住宿或工作或活动</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showOutdoorStayOrWorkWithinMonth('是')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.outdoorStayOrWorkWithinMonth == '是'">是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showOutdoorStayOrWorkWithinMonth('否')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.outdoorStayOrWorkWithinMonth == '否'">否</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showOutdoorStayOrWorkWithinMonth('不清楚')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.outdoorStayOrWorkWithinMonth == '不清楚'">不清楚</radio>
						</label>
					</radio-group>
				</view>
				<view class="row">
					<text class="label">野外/户外活动史</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<view class="order-options">
						<label class="radio" style="margin-bottom: 5px;">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showWildActivityHistory('种地')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.wildActivityHistory == '种地'">种地</radio>
						</label>
						<label class="radio" style="margin-bottom: 5px;">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showWildActivityHistory('割草')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.wildActivityHistory == '割草'">割草</radio>
						</label>
						<label class="radio" style="margin-bottom: 5px;">
							<radio :disabled="isViewMode"color="#285AC8" @tap="showWildActivityHistory('打猎')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.wildActivityHistory == '打猎'">打猎</radio>
						</label>
						</view>
						<view class="order-options">
						<label class="radio" style="margin-bottom: 5px;">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showWildActivityHistory('采茶')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.wildActivityHistory == '采茶'">采茶</radio>
						</label>
						<label class="radio" style="margin-bottom: 5px;">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showWildActivityHistory('放牧')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.wildActivityHistory == '放牧'">放牧</radio>
						</label>
						<label class="radio" style="margin-bottom: 5px;">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showWildActivityHistory('采伐')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.wildActivityHistory == '采伐'">采伐</radio>
						</label>
						</view>
						<view class="order-options" style="margin-right: 83px;">
						<label class="radio" style="margin-bottom: 5px;">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showWildActivityHistory('旅游')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.wildActivityHistory == '旅游'">旅游</radio>
						</label>
						<label class="radio" style="margin-bottom: 5px;">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showWildActivityHistory('其他')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.wildActivityHistory == '其他'">其他</radio>
						</label>
						</view>
					</radio-group>
				</view>
			</view>
			
			<!-- part3 -->
			<view class="sections1" v-show="sonPage==3">
				<view class="section-title">
					<text class="label">疾病流行区</text>
				</view>
				<view class="row">
					<text class="label">是否到过鼠疫流行地区</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showPlagueArea('是')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.plagueArea == '是'">是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showPlagueArea('否')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.plagueArea == '否'">否</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showPlagueArea('不清楚')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.plagueArea== '不清楚'">不清楚</radio>
						</label>
					</radio-group>
				</view>
				<veiw class="row" style="margin-left: 20px;" v-show="AllCheckInfo.riskFactorsAndExposure.plagueArea == '是'">
					具体地点<input :disabled="isViewMode" type="text" placeholder="请输入具体地点" style="margin-right: 30px;" v-model="AllCheckInfo.riskFactorsAndExposure.specificPlagueLocation"/>
				</veiw>
				<view class="row">
					<text class="label">是否到过炭疽流行地区</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showAnthraxArea('是')"  :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.anthraxArea== '是'">是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showAnthraxArea('否')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.anthraxArea== '否'">否</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showAnthraxArea('不清楚')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.anthraxArea== '不清楚'">不清楚</radio>
						</label>
					</radio-group>
				</view>
				<veiw class="row" style="margin-left: 20px;" v-show="AllCheckInfo.riskFactorsAndExposure.anthraxArea == '是'">
					具体地点<input :disabled="isViewMode" type="text" placeholder="请输入具体地点" style="margin-right: 30px;" v-model="AllCheckInfo.riskFactorsAndExposure.specificAnthraxLocation"/>
				</veiw>
				<view class="row">
					<text class="label">是否到过疟疾流行地区</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showMalariaArea('是')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.malariaArea== '是'">是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showMalariaArea('否')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.malariaArea== '否'">否</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showMalariaArea('不清楚')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.malariaArea== '不清楚'">不清楚</radio>
						</label>
					</radio-group>
				</view>
				<veiw class="row" style="margin-left: 20px;" v-show="AllCheckInfo.riskFactorsAndExposure.malariaArea == '是'">
					具体地点<input :disabled="isViewMode" type="text" placeholder="请输入具体地点" style="margin-right: 30px;" v-model="AllCheckInfo.riskFactorsAndExposure.specificMalariaLocation"/>
				</veiw>
			</view>
			
			
			<!-- part4 -->
			<view class="sections1" v-show="sonPage==4">
				<view class="section-title">
					<text class="label">动物接触史</text>
				</view>
				<view class="row">
					<text class="label">接触过动物制品</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showContactWithAnimalProducts('是')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.contactWithAnimalProducts== '是'">是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showContactWithAnimalProducts('否')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.contactWithAnimalProducts== '否'">否</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showContactWithAnimalProducts('不清楚')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.contactWithAnimalProducts== '不清楚'">不清楚</radio>
						</label>
					</radio-group>
				</view>
				<view class="row" style="margin-left: 20px;" v-show="AllCheckInfo.riskFactorsAndExposure.contactWithAnimalProducts == '是'">
					触摸时间				
					<picker :disabled="isViewMode" mode="date" value="2024-01-01" start="2024-01-01" end="2099-12-31" @change="onDateChange" class="time-picker">
					  <view style="float: right;">
					    {{ AllCheckInfo.riskFactorsAndExposure.contactDate || '请选择日期' }}
					  </view>
					</picker>
				</view>
				<view class="row">
					<text class="label">接触过动物</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showContactWithAnimals('是')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.contactWithAnimals== '是'">是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showContactWithAnimals('否')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.contactWithAnimals== '否'">否</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showContactWithAnimals('不清楚')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.contactWithAnimals== '不清楚'">不清楚</radio>
						</label>
					</radio-group>
				</view>
				<view class="row" style="margin-left: 20px;" v-show="AllCheckInfo.riskFactorsAndExposure.contactWithAnimals == '是'">
					触摸地点<input :disabled="isViewMode" type="text" placeholder="请输入触摸地点" style="margin-right: 30px;" v-model="AllCheckInfo.riskFactorsAndExposure.contactPlace"/>
				</view>
				<view class="row" v-show="AllCheckInfo.riskFactorsAndExposure.contactWithAnimals == '是'">
					<checkbox-group class="options">
					<view class="option-button" style="margin: 5px 0;">
					<label class="checkbox">
					  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasContactWithRat')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasEatingColdCookedFood">鼠</checkbox>
					</label>
					<label class="checkbox">
					  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasContactWithLice')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasEatingColdCookedFood">虱</checkbox>
					</label>
					<label class="checkbox">
					  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasContactWithFlea')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasEatingColdCookedFood">蚤</checkbox>
					</label>
					<label class="checkbox">
					  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasContactWithCat')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasEatingColdCookedFood">猫</checkbox>
					</label>
					</view>
					<view class="option-button" style="margin: 5px 0;">
					<label class="checkbox">
					  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasContactWithMarmot')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasContactWithMarmot">旱獭</checkbox>
					</label>
					<label class="checkbox">
					  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasContactWithBird')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasContactWithBird">鸟</checkbox>
					</label>
					<label class="checkbox">
					  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasContactWithPig')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasContactWithPig">猪</checkbox>
					</label>
					<label class="checkbox">
					  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasContactWithHorse')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasContactWithHorse">马</checkbox>
					</label>
					</view>
					<view class="option-button" style="margin: 5px 0;">
					<label class="checkbox">
					  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasContactWithMule')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasContactWithMule">骡</checkbox>
					</label>
					<label class="checkbox">
					  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasContactWithDonkey')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasContactWithDonkey">驴</checkbox>
					</label>
					<label class="checkbox">
					  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasContactWithCow')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasContactWithCow">牛</checkbox>
					</label>
					<label class="checkbox">
					  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasContactWithSheep')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasContactWithSheep">羊</checkbox>
					</label>
					</view>
					<view class="option-button" style="margin: 5px 0;">
					<label class="checkbox">
					  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasContactWithDog')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasContactWithDog">狗</checkbox>
					</label>
					<label class="checkbox">
					  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasContactWithRabbit')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasContactWithRabbit">兔</checkbox>
					</label>
					<label class="checkbox">
					  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasContactWithWolf')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasContactWithWolf">狼</checkbox>
					</label>
					<label class="checkbox">
					  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasContactWithTiger')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasContactWithTiger">虎</checkbox>
					</label>
					</view>
					<view class="option-button" style="margin-right: 265px;">
					<label class="checkbox">
					  <checkbox :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasContactWithOtherAnimal')" :checked="!isViewMode ? false: AllCheckInfo.riskFactorsAndExposure.hasContactWithOtherAnimal">其他</checkbox>
					</label>
					</view>
					</checkbox-group>
				</view>
			</view>
			
			<!-- part5 -->
			<view class="sections1" v-show="sonPage==5">
				<view class="section-title">
					<text class="label">发病前1个月居住地类型</text>
				</view>
				<view class="row">
				<checkbox-group class="options">
					<label class="checkbox">
			      <checkbox :disabled="isViewMode" color="#285AC8" @tap="showOtherWildActivityName('丘陵或山区')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.otherWildActivityName == '丘陵或山区'">丘陵或山区</checkbox>
			    </label>
			    <label class="checkbox">
			      <checkbox :disabled="isViewMode" color="#285AC8" @tap="showOtherWildActivityName('平原')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.otherWildActivityName == '平原'">平原</checkbox>
			    </label>
				<label class="checkbox">
				  <checkbox :disabled="isViewMode" color="#285AC8" @tap="showOtherWildActivityName('其他')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.otherWildActivityName == '其他'">其他</checkbox>
				</label>
				</checkbox-group>
				</view>
				
				<view class="section-title">
					<text class="label">14天内是否明确被蚊虫叮咬过</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasMosquitoBite')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.hasMosquitoBite">是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="handleFactorSymptoms('hasMosquitoBite')" :checked="!isViewMode ? false : !AllCheckInfo.riskFactorsAndExposure.hasMosquitoBite">否</radio>
						</label>
					</radio-group>
				</view>
				
				<view class="section-title">
					<text class="label">发病前2周内是否被蜱叮咬过</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showTickBite('是')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.tickBite == '是'">是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showTickBite('否')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.tickBite == '否'">否</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showTickBite('不清楚')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.tickBite == '不清楚'">不清楚</radio>
						</label>
					</radio-group>
				</view>
				
				<view class="section-title">
					<text class="label">发病是否被跳蚤叮咬过</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showFleaBite('是')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.fleaBite == '是'">是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showFleaBite('否')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.fleaBite == '否'">否</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showFleaBite('不清楚')" :checked="!isViewMode ? false: AllCheckInfo.riskFactorsAndExposure.fleaBite == '不清楚'">不清楚</radio>
						</label>
					</radio-group>
				</view>
				
				<view class="section-title">
					<text class="label">发病前2个月内是否接触鼠类</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showContactWithRatIn2Months('是')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.contactWithRatIn2Months  == '是'">是</radio>
						</label>
						<label class="radio">
							<radio color="#285AC8" @tap="showContactWithRatIn2Months('否')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.contactWithRatIn2Months  == '否'">否</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" @tap="showContactWithRatIn2Months('不清楚')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.contactWithRatIn2Months  == '不清楚'">不清楚</radio>
						</label>
					</radio-group>
				</view>
				<view class="row" v-show="AllCheckInfo.riskFactorsAndExposure.contactWithRatIn2Months == '是'">
					<text class="label">接触方式</text>
				</view>
				<view class="row" v-show="AllCheckInfo.riskFactorsAndExposure.contactWithRatIn2Months == '是'">
					<radio-group class="options">
						<view class="order-options" style="margin-right: 60px;">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showRatContactMethod('挖鼠洞')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.ratContactMethod == '挖鼠洞'">挖鼠洞</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showRatContactMethod('拿鼠')" :checked="!isViewMode ? false: AllCheckInfo.riskFactorsAndExposure.ratContactMethod == '拿鼠'">拿鼠</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showRatContactMethod('鼠咬')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.ratContactMethod == '鼠咬'">鼠咬</radio>
						</label>
						</view>
						<view class="order-options" style="margin-right: 30px;">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showRatContactMethod('鼠尿')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.ratContactMethod == '鼠尿'">鼠尿</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showRatContactMethod('鼠粪')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.ratContactMethod == '鼠粪'">鼠粪</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showRatContactMethod('鼠血污染手')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.ratContactMethod == '鼠血污染手'">鼠血污染手</radio>
						</label>
						</view>
						<view class="order-options" style="margin-right: 105px;">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showRatContactMethod('手被鼠夹打伤')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.ratContactMethod == '手被鼠夹打伤'">手被鼠夹打伤</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showRatContactMethod('其他')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.ratContactMethod == '其他'">其他</radio>
						</label>
						</view>
					</radio-group>
				</view>
				<view class="section-title">
					<text class="label">发病前1月内是否吃过被鼠排泄物污染的食物</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showEatenFoodContaminatedByRatFeces('是')"  :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.eatenFoodContaminatedByRatFeces == '是'">是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showEatenFoodContaminatedByRatFeces('否')"  :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.eatenFoodContaminatedByRatFeces == '否'">否</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showEatenFoodContaminatedByRatFeces('不清楚')"  :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.eatenFoodContaminatedByRatFeces == '不清楚'">不清楚</radio>
						</label>
					</radio-group>
				</view>
				
				<view class="section-title">
					<text class="label">发病前1月内是否在野外喝过沟（塘） 水</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showDrunkWaterFromDitchesOrPonds('是')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.drunkWaterFromDitchesOrPonds  == '是'">是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showDrunkWaterFromDitchesOrPonds('否')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.drunkWaterFromDitchesOrPonds  == '否'">否</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showDrunkWaterFromDitchesOrPonds('不清楚')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.drunkWaterFromDitchesOrPonds  == '不清楚'">不清楚</radio>
						</label>
					</radio-group>
				</view>
				
				<view class="section-title">
					<text class="label">发病前9天内是否在鼠洞附近坐卧</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showRestedNearRatHoles('是')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.restedNearRatHoles  == '是'">是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showRestedNearRatHoles('否')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.restedNearRatHoles  == '否'">否</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showRestedNearRatHoles('不清楚')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.restedNearRatHoles  == '不清楚'">不清楚</radio>
						</label>
					</radio-group>
				</view>
				
				<view class="section-title">
					<text class="label">工作场所有无鼠或鼠排泄物</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showRatOrRatDroppingsAtWorkplace('是')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.ratOrRatDroppingsAtWorkplace  == '是'">是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showRatOrRatDroppingsAtWorkplace('否')" :checked="!isViewMode ? false: AllCheckInfo.riskFactorsAndExposure.ratOrRatDroppingsAtWorkplace  == '否'">否</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showRatOrRatDroppingsAtWorkplace('不清楚')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.ratOrRatDroppingsAtWorkplace  == '不清楚'">不清楚</radio>
						</label>
					</radio-group>
				</view>
				
				<view class="section-title">
					<text class="label">发病前 1 月内是否接触过病人排泄物（血、尿、便等）</text>
				</view>
				<view class="row">
					<radio-group class="options">
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showContactWithPatientExcreta('是')" :checked="!isViewMode ? false: AllCheckInfo.riskFactorsAndExposure.contactWithPatientExcreta  == '是'">是</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showContactWithPatientExcreta('否')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.contactWithPatientExcreta  == '否'">否</radio>
						</label>
						<label class="radio">
							<radio :disabled="isViewMode" color="#285AC8" @tap="showContactWithPatientExcreta('不清楚')" :checked="!isViewMode ? false : AllCheckInfo.riskFactorsAndExposure.contactWithPatientExcreta  == '不清楚'">不清楚</radio>
						</label>
					</radio-group>
				</view>
			</view>
	      </view>
		  <view class="button-container" v-show="sonPage!=5">
		  <button class="return-button" @tap="handlereturn">上一步</button>
	      <button class="submit-button" @tap="handleSubmit">下一步</button>
		  </view>
		  <view class="button-container" v-show="sonPage==5">
		  <button class="return-button" @tap="handlereturn">上一步</button>
		  <button class="submit-button" @tap="handleFinallySubmit" v-show="!isViewMode">提交</button>
		  </view>
		  </view>
		<uni-custom-tabbar></uni-custom-tabbar>
	</view>
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
.content{
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
  margin-bottom: 20px;
  background-color: #ffffff;
  padding: 12px;
  border-radius: 8px;
}
.section-title{
	margin-left: -3px;
	padding-left: 4px;
	border-left: 8px solid #285AC8;
}
.row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 6px;
  margin-bottom: 12px;
}
.checkboxs{

}
.inner-checkboxs{
	margin-left: 5vh;
}
.label {
  font-size: 16px;
  color: #333333;
}

.options {
  display: flex;
  justify-content: space-around;
  flex-wrap: wrap;
}
.checkbox1{
  display: flex;
  align-items: center;
  margin-left: 7.5px;
  margin-right: 7.5px;
}
.radio, .checkbox {
  display: flex;
  align-items: center;
  margin-left: 15px;
  margin-right: 15px;
}

.picker {
  padding: 5px 10px;
  border: 1px solid #cccccc;
  border-radius: 4px;
  background-color: #ffffff;
}

.button-container {
  display: flex;
  justify-content: space-between; /* 均匀分配剩余空间 */
  align-items: center;
  margin-top: 5vh; /* 可根据需要调整 */
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
	border: 2px solid #285AC8;
	color: #285AC8;
	line-height: 40px;
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
.order-options{
	display: block; 
	margin: 15px 0; 
	display: flex; 
	justify-content: space-between; 
}
.choice-item{
	margin: 5px 0;
}
input{
	opacity: 0.7;
}
</style>