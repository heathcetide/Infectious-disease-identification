<script>
import { clockRecordService, clockSituationService } from '../../services/clock'
import store from '../../store/index.ts';
/**
 * 获取任意时间
 */
function getDate(date, AddDayCount = 0) {
  if (!date) {
    date = new Date();
  }
  if (typeof date !== 'object') {
    date = date.replace(/-/g, '/');
  }
  const dd = new Date(date);

  dd.setDate(dd.getDate() + AddDayCount); // 获取AddDayCount天后的日期

  const y = dd.getFullYear();
  const m = dd.getMonth() + 1 < 10 ? '0' + (dd.getMonth() + 1) : dd.getMonth() + 1; // 获取当前月份的日期，不足10补0
  const d = dd.getDate() < 10 ? '0' + dd.getDate() : dd.getDate(); // 获取当前几号，不足10补0
  return {
    fullDate: y + '-' + m + '-' + d,
    year: y,
    month: m,
    date: d,
    day: dd.getDay()
  };
}

export default {
  data() {
    return {
	  dateForm:{
	  	date: '',
		diagnose: '',
		examine: '',
		health: '',
	  },
	  dateForms:{
		date: '',
	  },
	  reportForm: {
	    isSerologicalTestDone: false,
	    isVirusAntigenTestDone: false,
	    isVirusNucleicAcidTestDone: false,
	    isVirusCultureIsolationDone: false,
	    labTestFiles: []  ,// 新增的 labTestFileIds 数组
	  },
	  todayInfo: '',
	  yearMonth: '',
	  checkFormData : {
		stoolReport: [],   // 粪便检查报告文件数组
		bloodReport: []    ,// 血液检查报告文件数组
		imageReport: [],
		ultrasoundReport: [],
		urineReport: [],
		medicalRecord: [],
	  },
	  formData: {
	    date: '',
	    stoolReport : [
	    ],   // 粪便检查报告文件数组
	    vomitReport : [
	      // 可以添加更多文件
	    ],   // 呕吐物检查报告文件数组
	    bloodReport: [
	      // 可以添加更多文件
	    ]    // 血液检查报告文件数组
	  },
      showCalendar: false,
      info: {
        lunar: true,
        range: true,
        insert: false,
        selected: [] // 日期打卡记录
      },
	  tempRecords:[],
      records: {
      },
	  diseases: [
	    { name: '新型冠状病毒', rate: '90%', link: '/pages/healthProtection/healthProtection' },
	    { name: '流感', rate: '80%', link: '/pages/healthProtection/healthProtection' },
	    { name: '鼠疫', rate: '75%', link: '/pages/healthProtection/healthProtection' },
	    { name: '感染性腹泻', rate: '60%', link: '/pages/healthProtection/healthProtection' },
	    { name: '炭疽', rate: '50%', link: '/pages/healthProtection/healthProtection' },
	  ],
	  calendarPages: true,
	  currentDateNow: '',
	  examineReport1: false,
	  examineReport2: false,
	  sonpages: 1,
	  pages : 1,
	  formData: {
	  		  date: '',
	  		  stoolReport : [
	  		    // 可以添加更多文件
	  		  ],   // 粪便检查报告文件数组
	  		  vomitReport : [
	  		    // 可以添加更多文件
	  		  ],   // 呕吐物检查报告文件数组
	  		  bloodReport: [
	  		    // 可以添加更多文件
	  		  ]    // 血液检查报告文件数组
	  		},
	  		AllCheckInfo : {
	  		  diagnosisResultsEmpDTO: {
	  		    admissionDate: '',
	  		    admissionSymptomsAndSigns: '',
	  		    anthraxSubtype: '',
	  		    deathDate: null,
	  		    dischargeDate: '',
	  		    discoveryMethod: '',
	  		    diseaseOutcome: '',
	  		    diseaseType: '',
	  		    hospitalName: '',
	  		    otherDiscoveryMethodName: null,
	  		    otherDiseaseName: null,
	  		    otherRegistrationDetails: null,
	  		    plagueSubtype: null,
	  		    registrationClassification: ''
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
	  		    diarrheaFrequencyGEThreeTimesPerDay: false,
	  		    diarrheaMode: null,
	  		    diarrheaTimesPerDay: 0,
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
	  		    stoolAmount: '',
	  		    stoolOdor: '',
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
	  		    bleedingCavity: '',
	  		    bleedingVolume: '',
	  		    dehydrationSeverity: '',
	  		    feverDuration: '',
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
	  		  fileIdList: []
	  		},
	  temporaryChoice1: false,
	  temporaryChoice2: false,
	  otherChoices: false,
	  SeafoodChoices: false,
    };
  },
  onReady() {
    this.$nextTick(() => {
      this.showCalendar = true;
    });

    // 模拟请求异步同步数据
    setTimeout(() => {
      this.info.date = getDate(new Date(), -30).fullDate;
      this.info.startDate = getDate(new Date(), -60).fullDate;
      this.info.endDate = getDate(new Date(), 30).fullDate;
      this.info.selected = Object.keys(this.records).map((date) => {
        return {
          date: date,
          info: this.records[date]
        };
      });
    }, 300);
  },
  mounted() {
	  this.initMonthSwitch()
	  clockRecordService(this.yearMonth)
			.then(res => {
				console.log(res.data)
				this.tempRecords = res.data; // 保存所有的记录数据
		        res.data.forEach(item => {
		          this.updateRecords(item);
		        });
			})
			.catch(error => this.submitError(error));
	  const diseasescoring = store.state.diseasescoring;
	  const formattedDiseases = diseasescoring.map((item) => ({
	    diseaseTypeName: item.diseaseTypeName,
	    score: item.score,
	    link: '/pages/healthProtection/healthProtection',
	  }));
	  // 对 formattedDiseases 进行排序
	  formattedDiseases.sort((a, b) => b.score - a.score);
	  // 将排序后的数据赋值给 diseases
	  this.diseases = formattedDiseases;
  },
  methods: {
    preventDelete() {
        // 防止删除操作
        console.log("删除操作被禁止");
        return false; // 返回 false 以防止删除
    },
	formatDate(date) {
	    const year = date.getFullYear();
	    const month = String(date.getMonth() + 1).padStart(2, '0'); // 月份从0开始，所以加1
	    const day = String(date.getDate()).padStart(2, '0');
	    return `${year}-${month}-${day}`;
	},
	checkSuccess(response) {
	  console.log("666666666", response);
	  // 接收到的数据进行排序
	  const sortedDiseases = response.data.diseasescoringList.sort((a, b) => b.score - a.score);
	  this.diseases = sortedDiseases;
	},
	checkError(response){
		console.log(response)
	},  
	updateRecords(data) {
	    const { date, health, diagnose, examine } = data;
	    const key = date;
	    if (!this.records[key]) {
	      this.records[key] = [];
	    }
	    this.records[key][0] = health ? 1 : 0;
	    this.records[key][1] = diagnose ? 1 : 0;
	    this.records[key][2] = examine ? 1 : 0;
	
	    console.log('Updated records:', this.records);
        this.$nextTick(() => {
          this.showCalendar = true;
        });
        
        // 模拟请求异步同步数据
        setTimeout(() => {
          this.info.date = getDate(new Date(), -30).fullDate;
          this.info.startDate = getDate(new Date(), -60).fullDate;
          this.info.endDate = getDate(new Date(), 30).fullDate;
          this.info.selected = Object.keys(this.records).map((date) => {
            return {
              date: date,
              info: this.records[date]
            };
          });
        }, 300);
	},
	initMonthSwitch() {
	    // 获取当前日期
	    const currentDate = new Date();
	    const currentYear = currentDate.getFullYear();
	    const currentMonth = String(currentDate.getMonth() + 1).padStart(2, '0'); // 确保月份为两位数
	    const currentDay = String(currentDate.getDate()).padStart(2, '0'); // 确保日期为两位数
	
	    // 存储当前的年月信息
	    this.currentYear = currentYear;
	    this.currentMonth = currentMonth;
	    this.currentDay = currentDay;
	    this.yearMonth = `${currentYear}-${currentMonth}`; // 拼接成 YYYY-MM 格式
	    this.todayInfo = `${currentYear}-${currentMonth}-${currentDay}`; // 拼接成 YYYY-MM-DD 格式
		
	    console.log(this.todayInfo);
		console.log('nownownownownow',this.yearMonth)
	},
	isImage(filePath) {
	  const extension = filePath.split('.').pop().toLowerCase();
	  return ['jpg', 'png'].includes(extension);
	},
	onDateChange(e) {
	  this.formData.date = e.detail.value;
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
	      url: 'https://ds.sccdc.cn/railwayemployee/labtest/file',  // 替换为你的接口地址
	      filePath: filePath,
	      name: 'file',  // 根据你的接口定义
	      formData: {
	        reportType: reportType,  // 上传的报告类型
	        index: index             // 当前文件的索引
	      },
	      header: {
	        'Content-Type': 'application/x-www-form-urlencoded'
	      },
	      success: (res) => {
	        console.log(`文件上传成功: ${res}`);
	        const fileId = parseInt(res.data); // 假设返回的是一个数字
	        this.reportForm.labTestFiles.push(fileId);
	        resolve(res);
	      },
	      fail: (err) => {
	        console.error(`文件上传失败: ${err}`);
	        reject(err);
	      }
	    });
	  });
	},
    change(e) {
	  // 检查所选日期是否在未来
      const currentDate = getDate(new Date()).fullDate;
      const selectedDate = e.fulldate;
      console.log('change 返回:', e,"66666666",currentDate,selectedDate);
      this.currentDateNow = e.fulldate;

      if (selectedDate > currentDate) {
        this.showUnmarkedPopup(selectedDate);
      }else{
		this.calendarPages = !this.calendarPages;
	  }
	  this.todayInfo = selectedDate;
	  
	  // 查找与选中日期匹配的记录
	  const selectedRecord = this.tempRecords.find(record => record.date === this.todayInfo);
	  
	  if (selectedRecord) {
	    // 如果找到匹配的记录，赋值 diagnose, examine, health 属性
	    this.dateForm.diagnose = selectedRecord.diagnose;
	    this.dateForm.examine = selectedRecord.examine;
	    this.dateForm.health = selectedRecord.health;
	  } else {
	    // 如果没有找到匹配的记录，设置默认值
	    this.dateForm.diagnose = false;
	    this.dateForm.examine = false;
	    this.dateForm.health = false;
	  }
	  
	  console.log('测试',this.dateForm)
	  this.dateForms.date = this.todayInfo;
	  clockSituationService(this.dateForms)
	  	.then(res => this.checkSuccess(res))
	  	.catch(error => this.checkError(error));
    },
	showUnmarkedPopup(date) {
	  uni.showToast({
	    title: `日期 ${date} 尚未打卡`,
	    icon: 'none',
	    duration: 2000
	  });
	},
    monthSwitch(e) {
      console.log('monthSwitch 返回:', e);
	  this.yearMonth = e.year + '-'+ e.month
	  clockRecordService(this.yearMonth)
	  			.then(res => {
	  				console.log(res.data)
	  		        res.data.forEach(item => {
	  		          this.updateRecords(item);
	  		        });
	  			})
	  			.catch(error => this.submitError(error));
    },
    navigateToDay(day) {
      const selectedDate = day.fullDate;
      wx.navigateTo({
        url: `/pages/targetPage/targetPage?date=${selectedDate}`
      });
    },
	onDateChange(e) {
	  this.formData.date = e.detail.value;
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
	        index: index             // 当前文件的索引
	      },
	      header: {
	        'Content-Type': 'application/x-www-form-urlencoded'
	      },
	      success: (res) => {
	        console.log(`文件上传成功: ${res}`);
	        const fileId = parseInt(res.data); // 假设返回的是一个数字
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
	          uploadPromises.push(this.uploadFile(file.path, '粪便检查报告', index));
	        });
	
	        // 上传呕吐物检查报告
	        this.formData.vomitReport.forEach((file, index) => {
	          uploadPromises.push(this.uploadFile(file.path, '呕吐物检查报告', index));
	        });
	
	        // 上传血液检查报告
	        this.formData.bloodReport.forEach((file, index) => {
	          uploadPromises.push(this.uploadFile(file.path, '血液检查报告', index));
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
	checkExamineData() {
	    console.log('日期准备', this.todayInfo);
	    // 检查 examine 的状态
	    if (this.dateForm.diagnose === false) {
	        // 如果 examine 为 false，显示提示信息并阻止跳转
	        uni.showToast({
	            title: '用户当日未提交检测数据',
	            icon: 'none',
	            duration: 2000
	        });
	    } else {
	        // 如果 examine 为 true，跳转到 diagnosis 页面并携带日期参数
	        const dateParam = this.todayInfo;
	        uni.navigateTo({
	            url: `/pages/diagnosis/diagnosis?date=${dateParam}`,
	        });
	    }
	},
	checkHealthData(){
		console.log('日期准备', this.todayInfo);
		if (this.dateForm.health === true) {
		    uni.showToast({
		        title: '用户打卡健康，未提交当日症状',
		        icon: 'none',
		        duration: 2000
		    });
		} else {
		    // 如果 examine 为 true，跳转到 diagnosis 页面并携带日期参数
		    const dateParam = this.todayInfo;
		    uni.navigateTo({
		        url: `/pages/todaySymptoms/todaySymptoms?date=${dateParam}`,
		    });
		}
	},
	checkDiagnosisData(){
		console.log('日期准备', this.todayInfo);
		if (this.dateForm.examine === false) {
		    uni.showToast({
		        title: '用户当日未提交检测数据',
		        icon: 'none',
		        duration: 2000
		    });
		} else {
		    // 如果 examine 为 true，跳转到 diagnosis 页面并携带日期参数
		    const dateParam = this.todayInfo;
		    uni.navigateTo({
		        url: `/pages/experimentalInput/experimentalInput?date=${dateParam}`,
		    });
		}
	},
	handleFileChange(event, reportType) {
		// Handle file selection
		console.log(reportType, event);
		// Update formData based on selected files
		this.formData[reportType] = event.files;
	},
  }
};
</script>

<template>
<view class="allPages">
  <view class="container1" v-show="calendarPages">
    <image src="../../static/other-background.png" mode="widthFix" class="background-image1"></image>
    <view class="content1" v-if="showCalendar">
      <view>
        <uni-calendar 
          class="uni-calendar--hook" 
          :selected="info.selected" 
          :showMonth="false" 
          @change="change" 
          @monthSwitch="monthSwitch" 
          @tapDay="navigateToDay" 
        />
      </view>
    </view>
    <uni-custom-tabbar class="tabbar"></uni-custom-tabbar>
  </view>
  

  <view class="container" v-show="!calendarPages">
    <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
    <image src="../../static/virus.png" mode="aspectFit" class="virus-image1"></image>
    <view class="clock-info1">
      <view class="clock-text1">
        <text>您已经在{{currentDateNow}}打卡</text>
      </view>
      <text>
        <text class="clock-texts1">根据您提交的信息</text>
        <text class="clock-texts1">您罹患或感染以下疾病的</text>
		<text class="clock-texts1">风险为：</text>
      </text>
    </view>
    <view class="content">
      <view class="list-container1">
        <view v-if="diseases.length > 0">
          <view class="list-item1" v-for="(item, index) in diseases" :key="index">
            <view class="disease-info1">
              <view class="index-box1">
                <text class="index1">{{ index + 1 }}</text>
              </view>
              <view class="text-box">
                <text class="disease-name1">{{ item.diseaseTypeName }}</text>
                <navigator class="health-protection1" :url="item.link">
                  <text>健康防护</text>
                  <text>></text>
                </navigator>
              </view>
            </view>
            <view class="rate-box1">
              <text class="rate1">{{ item.score }}%</text>
              <text class="rate-label1">患病率</text>
            </view>
          </view>
        </view>
        <view v-else class="no-data-container">
          <text class="no-data">暂无疾病风险记录</text>
        </view>
      </view>
	  <div class="button-container1" :style="{'margin-top': diseases.length <= 0 ? '320px' : ''}">
	    <button class="button1" @tap="checkHealthData()">查看当日症状</button>
	    <button class="button1" @tap="checkDiagnosisData()">查看诊断数据</button>
	    <button class="button1" @tap="checkExamineData()">查看检测数据</button>
	  </div>
    </view>
    <uni-custom-tabbar></uni-custom-tabbar>
  </view>
</view>
</template>
<style lang="scss">
.container1 {
  width: 100%;
  height: 100%;
  position: relative; /* 设置为相对定位，使内部元素能够相对于容器定位 */
}
  
.background-image1 {
  position: absolute; /* 设置为绝对定位，使其占据整个容器 */
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1; /* 将图片置于内容下方 */
}

.tabbar {
  z-index: 999;
}

.content1 {
  position: relative;
  padding-top: 14vh;
  border-top-left-radius: 40px;
  border-top-right-radius: 40px;
  padding-bottom: 13vh;
  margin: 0 16px;
  border-radius: 5px;
}

.uni-calendar--hook .uni-calendar-item__text {
  border: 3px solid transparent; /* 默认无边框颜色 */
}

.uni-calendar--hook .uni-calendar-item__text[data-state='normal'] {
  border-color: #4caf50; /* 绿色边框 */
}

.uni-calendar--hook .uni-calendar-item__text[data-state='missed'] {
  border-color: #ffeb3b; /* 黄色边框 */
}

.uni-calendar--hook .uni-calendar-item__text[data-state='problem'] {
  border-image: linear-gradient(to right, #ffeb3b 50%, #f44336 50%);
  border-image-slice: 1;
  border-radius: 50%; /* 圆形边框 */
}

.uni-calendar--hook .uni-calendar-item__text[data-state='severe'] {
  border-image: linear-gradient(to right, #ffeb3b 33%, #f44336 33%, #9c27b0 34%);
  border-image-slice: 1;
  border-radius: 50%; /* 圆形边框 */
}

.uni-calendar--hook .uni-calendar-item__text.today {
  background-color: #285AC8; /* 当日背景颜色 */
  color: #ffffff;
}

.container {
  width: 100%;
  height: 100%;
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
.button-container1 {
  display: flex; /* 使用Flex布局使按钮水平排列 */
  justify-content: space-around; /* 平均分配按钮间的空间 */
  // margin-top: 320px; /* 为按钮容器添加上边距 */
  position: sticky;
  bottom: 0px;
  max-height: 60px;
  background-color: #FFFFFF;
}
.button1 {
  padding: 5px 10px; /* 内边距设置 */
  font-size: 13px; /* 文字大小 */
  border: none; /* 去除边框 */
  cursor: pointer; /* 鼠标悬停时指针变为手型 */
  background-color: #FFFFFF; /* 按钮背景颜色 */
  color: #285AC8; /* 文字颜色 */
  text-align: center; /* 文字居中 */
  text-decoration: none; /* 移除下划线 */
  transition: 0.3s; /* 过渡效果 */
  border-radius: 5px; /* 边框圆角 */
  border: 2px solid #285AC8;
  margin-left: 2px;
  margin-right: 2px;
}
.button1:hover {
  background-color: #285AC8; /* 鼠标悬停时的背景色变化 */
  color: #FFFFFF;
}

.clock-info1 {
  position: fixed;
  bottom: 70vh; /* 调整位置，使其离下面的内容更近 */
  font-size: 20px;
  color: #646464;
  left: 5vh;
}

.clock-text1 {
  display: flex; /* 使用flex布局，使文字和图标在同一行 */
  align-items: center; /* 垂直居中对齐 */
  margin-bottom: 3vh; /* 减小底部间距 */
}

.clock-texts1 {
  font-size: 18px;
  display: block;
  font-weight: 800;
}

.virus-image1 {
  position: absolute;
  width: 170px;
  right: -20px;
  top: 12vh;
}

.content {
  position: fixed;
  bottom: 110px;
  left: 0;
  right: 0;
  width: 100%;
  height: 50vh;
  overflow-y: auto; /* 如果内容超出容器，则显示滚动条 */
  background-color: #ffffff;
  z-index: 20;
  border-top-left-radius: 40px;
  border-top-right-radius: 40px;
}

.list-container1 {
  padding: 16px;
}

.list-item1 {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #ffffff;
  padding: 12px;
  margin-bottom: 10px;
  border-radius: 8px;
}

.disease-info1 {
  display: flex;
  justify-content: space-between;
}

.disease-name1 {
  font-size: 16px;
  color: #333333;
  margin-bottom: 4px;
}

.health-protection1 {
  display: flex;
  align-items: center;
  color: #007aff;
  font-size: 14px;
  text-decoration: none;
}

.rate-box1 {
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #285AC8;
  color: #ffffff;
  padding: 8px;
  border-radius: 8px;
  min-width: 39px;
  min-height: 47px;
}

.rate1 {
  font-size: 16px;
  font-weight: bold;
}

.rate-label1 {
  font-size: 12px;
}

.index-box1 {
  margin-right: 12px;
}

.index1 {
  font-size: 16px;
  color: #E5E5E5;
  font-weight: bold;
}
// .content {
//   position: fixed;
//   bottom: 60px; /* Adjusted from 110px to create more space for the tabbar */
//   left: 0;
//   right: 0;
//   width: 100%;
//   height: calc(50vh - 60px); /* Ensure the content does not overlap the tabbar */
//   overflow-y: auto; /* Keeps scrolling for overflowed content */
//   background-color: #ffffff;
//   z-index: 20;
//   border-top-left-radius: 40px;
//   border-top-right-radius: 40px;
//   box-shadow: 0 -2px 4px rgba(0, 0, 0, 0.1); /* Added subtle shadow for visual separation */
// }

.footer {
  height: 10vh;
  display: flex;
  justify-content: space-around;
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  width: 100%;
  background-color: #ffffff;
  padding: 10px 0;
  box-shadow: 0 -2px 4px rgba(0, 0, 0, 0.1);
  z-index: 25; /* Ensure the footer stays on top of other content */
}

.footer-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  font-size: 14px;
  color: #000000;
}

.footer-item image {
  width: 24px;
  height: 24px;
  margin-bottom: 5px;
}

// .button-container1 {
//   margin-top: 20px;
//   display: flex;
//   flex-direction: column;
//   align-items: center;
// }

.button1 .button2 {
  width: 90%;
  padding: 12px;
  margin: 8px 0;
  background-color: #285AC8;
  color: #ffffff;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  text-align: center;
}

/* Adjusted dropdown styles to ensure it does not cover the footer */
.dropdown-content {
  position: fixed;
  bottom: 60px; /* Ensures dropdown does not overlap with the footer */
  left: 0;
  width: 100%;
  background: #ffffff;
  border-top-left-radius: 8px;
  border-top-right-radius: 8px;
}

.container2 {
	width: 100%;
	height: 100%;
	position: relative; /* 设置为相对定位，使内部元素能够相对于容器定位 */
}
.content2{
	position: relative;
	padding-top: 14vh;
	border-top-left-radius: 40px;
	border-top-right-radius: 40px;
	padding-bottom: 13vh;
}
.section2 {
  margin-bottom: 20px;
  padding: 12px;
  border-radius: 8px;
  color: #333333;
}
.section-title2{
	margin-left: -3px;
	padding-left: 4px;
	border-left: 8px solid #285AC8;
}
.row2 {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 24px;
  margin-bottom: 36px;
}

.checkboxs2{
	margin: 7px 30px;
}
.order-options2{
	display: block; 
	margin: 15px 0; 
	display: flex; 
	justify-content: space-between; 
}
.label2 {
  font-size: 16px;
  color: #333333;
  display: block;
}
.time-picker2{
	margin-top: -18px;
	margin-left: 180px;
}
.center-button2 {
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
  margin-bottom: 12px;
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
  margin: 5px; /* 添加四周的边距 */
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
  border: 1px solid #cccccc;
  border-radius: 4px;
  background-color: #ffffff;
}
.button-container {
  display: flex;
  justify-content: space-between; /* 均匀分配剩余空间 */
  align-items: center;
}
.no-data-container {
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.no-data {
  font-size: 22px;
  color: #285AC8;
  font-weight: bold;
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
</style>
          <!-- 粪便检查报告 -->
<!--          <view class="row2">
            <view class="form-item">
              <view class="section-title">
                <text class="label2">粪便检查报告：</text>
              </view>
              <button @tap="chooseFile('stoolReport')" class="center-button2">文件上传</button>
              <view v-if="formData.stoolReport.length > 0" class="file-preview">
                <view v-for="(file, index) in formData.stoolReport" :key="index" class="file-item">
                  <image v-if="isImage(file.path)" :src="file.path" mode="aspectFit" style="width: 100px; height: 100px;" />
                  <view v-else>{{ file.name }}</view>
                </view>
              </view>
            </view>
          </view> -->
          <!-- 呕吐物检查报告 -->
<!--          <view class="row2">
            <view class="form-item">
              <view class="section-title">
                <text class="label2">呕吐物检查报告：</text>
              </view>
              <button @tap="chooseFile('vomitReport')" class="center-button2">文件上传</button>
              <view v-if="formData.vomitReport.length > 0" class="file-preview">
                <view v-for="(file, index) in formData.vomitReport" :key="index" class="file-item">
                  <image v-if="isImage(file.path)" :src="file.path" mode="aspectFit" style="width: 100px; height: 100px;" />
                  <view v-else>{{ file.name }}</view>
                </view>
              </view>
            </view>
          </view> -->
          <!-- 血液检查报告 -->
<!--          <view class="row2">
            <view class="form-item">
              <view class="section-title">
                <text class="label2">血液检查报告：</text>
              </view>
              <button @tap="chooseFile('bloodReport')" class="center-button2">文件上传</button>
              <view v-if="formData.bloodReport.length > 0" class="file-preview">
                <view v-for="(file, index) in formData.bloodReport" :key="index" class="file-item">
                  <image v-if="isImage(file.path)" :src="file.path" mode="aspectFit" style="width: 100px; height: 100px;" />
                  <view v-else>{{ file.name }}</view>
                </view>
              </view>
            </view>
          </view> -->