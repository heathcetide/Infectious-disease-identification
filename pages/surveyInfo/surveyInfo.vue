<script>
import { submitUserInfoService, getUserInfoService } from '../../services/login';
export default {
  data() {
    return {
	  staticForm: {
		medicalPersonnelTypeRange: ['医生','护士','疾控现场工作人员','试验室检测人员','其他'],
		specificOccupationRange: ['否','医务人员','病原微生物检测人员','野生动物接触相关人员','家禽、家畜养殖人员'],
		departmentRange: ['安全部','财务部','测量队','工程技术部','合约部','试验室','物资设备部','项目管理层','征拆协调部','综合管理部'],
		educationLevelRange: ['小学及以下','初中','中专/高中','大专/本科','硕士及以上'],
		ethnicityRange: ['汉族','藏族','彝族','其他少数民族'],
		drinkingStatusRange: ['不喝酒','偶尔喝酒（小于3次/周）','经常喝酒（大于等于3次/周）','经常喝大量酒（白酒≥3次/周，每次≥2两）'],
		smokingStatusRange: ['现在吸','以前吸','从不吸'],
	  },
	  submitForm: {
		name: '',
		gender: '男',
		pregnant: 'false',
		pregnancyWeeks: '',
		idNumber:'',
		age: '',
		ethnicity:'',
		educationLevel: '',
		department:'',
		specificOccupation: '',
		workOnPlateauStartDate:'',
		medicalPersonnelType: '',
		phoneNumber:'',
		otherPhoneNumber:'',
		
		hasMedicalHistory:'',
		hasHypertension: false,
		hasDiabetes: false,
		hasHyperlipidemia: false,
		hasHyperuricemia: false,
		hasCoronaryHeartDisease: false,
		hasStroke: false,
		hasOtherCardiovascularDiseases: false,
		hasAsthma:false,
		hasCOPD: false,
		hasPepticUlcer:false,
		hasMalignantTumor:false,
		hasLungCancer:false,
		hasOtherCancer: false,
		otherCancerName: '',
		hasSevereMentalDisorders: false,
		hasTuberculosis: false,
		hasHepatitis: false,
		hasOccupationalDisease: false,
		hasChronicKidneyDisease: false,
		hasChronicLiverDisease: false,
		hasImmunodeficiency: false,
		hasTyphus: false,
		isPostpartumInSixWeeks: false,
		hasDustExposure: false,
		hasOtherDiseases: false,
		otherDiseasesName: '',
		smokingStatus:'',
		drinkingStatus:'',
		height: '',
		weight: '',
		isVaccinatedForCOVID: false,
		isVaccinatedForFlu: false,
		isVaccinatedForPlague: false,
		isVaccinatedForBCG: false,
		isVaccinatedForHepatitis: false,
		emergencyContactName: '',
		emergencyContactPhoneNumber: '',
		emergencyContactRelation: '',
	  },
	  pages: 1,
    };
  },
  mounted(){
	  getUserInfoService()
			.then(res => this.getInfoSuccess(res))
			.catch(error => this.getInfoError(error));
  },
  methods: {
	getInfoSuccess(response) {
		console.log(response)
	    this.submitForm = response.data;
	},
	getInfoError(error) {
	    console.error('获取信息失败', error);
		uni.showToast({ 
			title: '获取信息失败，请联系管理员',
			icon: 'none', // 可选，默认为none，还有success, loading, warning, error等
			duration: 2000, // 显示持续时间，单位为毫秒，默认为1500
		});
	},
	handleDrinkingStatusChange(e){
		this.submitForm.drinkingStatus = this.staticForm.drinkingStatusRange[e.detail.value];
	},
	handleSmokingStatusChange(e){
		this.submitForm.smokingStatus = this.staticForm.smokingStatusRange[e.detail.value];
	},
	handleMedicalPersonnelTypeChange(e){
		this.submitForm.medicalPersonnelType = this.staticForm.medicalPersonnelTypeRange[e.detail.value];
	},
	handleSpecificOccupationChange(e){
		this.submitForm.specificOccupation = this.staticForm.specificOccupationRange[e.detail.value];
	},
	handleDepartmentChange(e){
		this.submitForm.department = this.staticForm.departmentRange[e.detail.value];
	},
	handleEducationLevelChange(e){
		this.submitForm.educationLevel = this.staticForm.educationLevelRange[e.detail.value];
	},
	handleethnicityChange(e){
		this.submitForm.ethnicity = this.staticForm.ethnicityRange[e.detail.value];
	},
	handleHasMedicalHistory(hasMedicalHistory){
		this.submitForm.hasMedicalHistory = hasMedicalHistory
	},
	submitError(error) {
		uni.showToast({ 
			title: '信息失败',
			icon: 'none', // 可选，默认为none，还有success, loading, warning, error等
			duration: 8000, // 显示持续时间，单位为毫秒，默认为1500
		});
	},
	submitSuccess(response) {
		console.log('提交信息成功',response)
	    if (response.code == 1) {
	        uni.showToast({
	          title: '信息修改成功',
	          icon: 'success',
			  duration:8000,
			  success: function(res){
				  uni.redirectTo({
				  	url:'/pages/clock/clock'
				  })
			  }
	        });
	    } else {
	        uni.showModal({
	            title: '信息修改失败',
	            content: '信息修改失败。',
	            showCancel: false,
	            confirmText: '确定',
	            success: function (res) {
	            }
	        });
	    }
	},
    handleSubmit() {
	  if( this.pages >= 4){
		uni.showModal({
            title: '提示',
            content: '您确定要提交吗？',
            confirmText: '确定',
            cancelText: '取消',
            success: (res) => {
				if (res.confirm) {
                    const that = this; // 保存当前的 this 上下文
                    submitUserInfoService(that.submitForm)
                        .then((res) => that.submitSuccess(res))
                        .catch((error) => that.submitError(error));
                }
            }
        });
	  }else{
		  this.pages++;
	  }
    },
	handlereturn(){
		this.pages--;
	},
	showRadio(maleValue){
		this.submitForm.gender = maleValue
		console.log('gender为',this.gender)
	},
	showPregnant(pregnantValue){
		this.submitForm.pregnant = pregnantValue
		console.log('pregnant为',this.pregnant)
	},
	onDateChange(e) {
	  this.submitForm.workOnPlateauStartDate = e.detail.value;
	},
    handleCheckboxChange(propName) {
      this.submitForm[propName] = !this.submitForm[propName];
	  console.log(`Changing ${propName}:`, this.submitForm[propName]);
    },
  },
};
</script>
<template>
	<view class="container">
	  <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
		<view class="content">
		<!-- 第一个 -->
		<view class="section" v-show="pages == 1">
				<view class="sections1">
					<view class="section-title">
						<text class="label">基本信息</text>
					</view>
						<view class="row">
							<view class="form-item">
								<text style="float: left;">姓名</text>
								<input type="text" v-model="submitForm.name" placeholder="请输入姓名" class="text-input" style="float: right;"/>
							</view>
						</view>
						<view class="row" style="margin: 7px 24px;">
						    <text class="label">性别</text>
						    <radio-group class="options">
						        <label class="radio">
						            <radio color="#285AC8" v-model="submitForm.gender" value="男" @tap="showRadio('男')" :checked="submitForm.gender == '男'"/>男
						        </label>
						        <label class="radio">
						            <radio color="#285AC8" v-model="submitForm.gender" value="女" @tap="showRadio('女')" :checked="submitForm.gender == '女'"/>女
						        </label>
						    </radio-group>
						</view>
						<view class="row" style="margin: 7px 24px;" v-if="submitForm.gender === '女'">
							<text class="label">是否为孕妇</text>
							<radio-group class="options">
								<label class="radio">
									<radio color="#285AC8" v-model="submitForm.pregnant" value="true" @tap="showPregnant('true')" :checked="submitForm.isPregnant == true">是</radio>
								</label>
								<label class="radio">
									<radio color="#285AC8" v-model="submitForm.pregnant" value="false" @tap="showPregnant('false')" :checked="submitForm.isPregnant == false">否</radio>
								</label>
							</radio-group>
						</view>
						<view class="row" v-show="submitForm.isPregnant === true">
							<view class="form-item" style="margin-left: 24px ;">
								<text style="float: left;">孕周</text>
								<input type="text" placeholder="请输入孕周" v-model="submitForm.pregnancyWeeks" class="text-input" style="float: right;"/>
							</view>
						</view>
						<view class="row">
							<view class="form-item">
								<text style="float: left;">身份证号</text>
								<input type="text" placeholder="请输入身份证号" v-model="submitForm.idNumber" class="text-input" style="float: right;"/>
							</view>
						</view>
						<view class="row">
							<view class="form-item">
								<text style="float: left;">年龄</text>
								<input type="text" placeholder="请输入年龄" v-model="submitForm.age" class="text-input" style="float: right;"/>
							</view>
						</view>
						<view class="row">
							<view class="form-item">
								<text style="float: left;">民族</text>
								<picker mode="selector" :range="staticForm.ethnicityRange" @change="handleethnicityChange">
									<view style="float: right;">{{ submitForm.ethnicity || '请选择' }}</view>
								</picker>							
							</view>
						</view>
						<view class="row">
							<view class="form-item">
								<text style="float: left;">受教育程度</text>
								<picker mode="selector" :range="staticForm.educationLevelRange" @change="handleEducationLevelChange">
									<view style="float: right;">{{ submitForm.educationLevel || '请选择' }}</view>
								</picker>	
							</view>
						</view>
					</view>
					<button class="single-submit-button" @tap="handleSubmit">下一步</button>
		</view>
		<!-- 第二个 -->
		<view class="section" v-show="pages == 2">
					<view class="sections1">
						<view class="section-title">
							<text class="label">工作信息</text>
						</view>
						<view class="row">
							<view class="form-item">
								<text style="float: left;">来高原工作时间</text>
								<picker mode="date" value="2024-01-01" start="2024-01-01" end="2099-12-31" @change="onDateChange" class="time-picker">
								  <view style="float: right;">
								    {{ submitForm.workOnPlateauStartDate || '请选择日期' }}
								  </view>
								</picker>
							</view>
						</view>
						<view class="row">
							<view class="form-item">
								<text style="float: left;">部门工种</text>
								<picker mode="selector" :range="staticForm.departmentRange" @change="handleDepartmentChange">
									<view style="float: right;">{{ submitForm.department || '请选择' }}</view>
								</picker>	
							</view>
						</view>
						<view class="row">
							<view class="form-item">
								<text style="float: left;">是否为以下特定职业人群</text>
								<picker mode="selector" :range="staticForm.specificOccupationRange" @change="handleSpecificOccupationChange">
									<view style="float: right;">{{ submitForm.specificOccupation || '请选择' }}</view>
								</picker>	
							</view>
						</view>
						<view class="row">
							<view class="form-item">
								<text style="float: left; font-size: 12px;">如为医务人员，请选择具体工作性质</text>
								<picker mode="selector" :range="staticForm.medicalPersonnelTypeRange" @change="handleMedicalPersonnelTypeChange">
									<view style="float: right;">{{ submitForm.medicalPersonnelType || '请选择' }}</view>
								</picker>
							</view>
						</view>
						<view>
							<text class="label" style="margin: 24px;">联系电话</text>
<!-- 							<view class="form-item">
								<text style="float: left;">手机</text>
								<input type="text" placeholder="请输入手机电话" v-model="submitForm.phoneNumber" class="text-input" style="float: right;"/>
							</view> -->
							<view class="form-item">
							    <text style="float: left;">其他联系电话</text>
							    <input type="text" placeholder="请输入其他联系电话" v-model="submitForm.otherPhoneNumber" class="text-input" style="float: right; line-height: 30px; height: 30px;"/>
							</view>
							
						</view>
						<view>
							<text class="label" style="margin: 24px;">紧急联系人</text>
							<view class="form-item">
								<text style="float: left;">姓名</text>
								<input type="text" placeholder="请输入紧急联系人姓名" v-model="submitForm.emergencyContactName" class="text-input" style="float: right;"/>
							</view>
							<view class="form-item">
								<text style="float: left;">联系电话</text>
								<input type="text" placeholder="请输入联系电话" v-model="submitForm.emergencyContactPhoneNumber" class="text-input" style="float: right;"/>
							</view>
							<view class="form-item">
								<text style="float: left;">关系</text>
								<input type="text" placeholder="请输入与该联系人关系" v-model="submitForm.emergencyContactRelation" class="text-input" style="float: right;"/>
							</view>
						</view>
					</view>
					<view class="button-container">
					<button class="return-button" @tap="handlereturn">上一步</button>
					<button class="submit-button" @tap="handleSubmit">下一步</button>
					</view>
		</view>
		<!-- 第三个 -->
		<view class="section" v-show="pages == 3">
					<view class="sections1">
						<view class="section-title">
							<text class="label">既往病史基本情况</text>
						</view>
						<view class="row">
							<text class="label" style="margin: 7px 15px;">既往病史和基本情况</text>
							<radio-group class="options">
								<label class="radio">
									<radio color="#285AC8" v-model="hasMedicalHistory" @tap="handleHasMedicalHistory('true')" :checked="submitForm.hasMedicalHistory == true">有</radio>
								</label>
								<label class="radio">
									<radio color="#285AC8" v-model="hasMedicalHistory" @tap="handleHasMedicalHistory('false')" :checked="submitForm.hasMedicalHistory == false">无</radio>
								</label>
							</radio-group>
						</view>
						<view class="checkboxs" v-show="submitForm.hasMedicalHistory == true">
						      <checkbox-group>
								<view class="order-options">
						        <label class="checkbox-item">
						            <checkbox color="#285AC8" value="true" :checked="submitForm.hasHypertension" @tap="handleCheckboxChange('hasHypertension')" />
						            高血压
						        </label>
						        <label class="checkbox-item">
						            <checkbox color="#285AC8" value="true" :checked="submitForm.hasDiabetes" @tap="handleCheckboxChange('hasDiabetes')" />
						            糖尿病
						        </label>
						        <label class="checkbox-item">
						            <checkbox color="#285AC8" value="true" :checked="submitForm.hasHyperlipidemia" @tap="handleCheckboxChange('hasHyperlipidemia')" />
						            高血脂症
						        </label>
								</view>
								<view class="order-options">
						        <label class="checkbox-item">
						          <checkbox color="#285AC8" value="true" :checked="submitForm.hasHyperuricemia" @tap="handleCheckboxChange('hasHyperuricemia')" />
						          高尿酸
						        </label>
						        <label class="checkbox-item">
						          <checkbox color="#285AC8" value="true" :checked="submitForm.hasCoronaryHeartDisease" @tap="handleCheckboxChange('hasCoronaryHeartDisease')" />
						          冠心病
						        </label>
						        <label class="checkbox-item">
						          <checkbox color="#285AC8" value="true" :checked="submitForm.hasStroke" @tap="handleCheckboxChange('hasStroke')"  />
						          脑卒中
						        </label>
								</view>
								<view class="order-options">
						        <label class="checkbox-item">
						          <checkbox color="#285AC8" value="true" :checked="submitForm.hasOtherCardiovascularDiseases" @tap="handleCheckboxChange('hasOtherCardiovascularDiseases')" />
						          其他心脑血管疾病
						        </label>
						        <label class="checkbox-item">
						          <checkbox color="#285AC8" value="true" :checked="submitForm.hasAsthma" @tap="handleCheckboxChange('hasAsthma')" />
						          哮喘
						        </label>
								</view>
								<view class="order-options">
								<label class="checkbox-item">
								  <checkbox color="#285AC8" value="true" :checked="submitForm.hasCOPD" @tap="handleCheckboxChange('hasCOPD')" />
								  慢性阻塞性肺部疾病
								</label>
								<label class="checkbox-item">
								  <checkbox color="#285AC8" value="true" :checked="submitForm.hasPepticUlcer" @tap="handleCheckboxChange('hasPepticUlcer')" />
								  消化性溃疡
								</label>
								</view>
								<view class="order-options">
								<label class="checkbox-item">
								  <checkbox color="#285AC8" value="true" :checked="submitForm.hasMalignantTumor" @tap="handleCheckboxChange('hasMalignantTumor')"  />
								  恶性肿瘤
								</label>
								</view>
								<view class="order-options" style="margin-left: 30px;" v-show="submitForm.hasMalignantTumor">
									<checkbox-group>
										<label class="checkbox-item">
										  <checkbox color="#285AC8" value="true" :checked="submitForm.hasLungCancer" @tap="handleCheckboxChange('hasLungCancer')"  />
										  肺癌
										</label>
										<label class="checkbox-item">
										  <checkbox color="#285AC8" value="true" :checked="submitForm.hasOtherCancer" @tap="handleCheckboxChange('hasOtherCancer')"  />
										  其他
										</label>
										<label class="checkbox-item">
										  <!-- <input type="text" placeholder="请输入" v-model="submitForm.otherCancerName" class="text-input" style="float: right;"/> -->
										</label>
									</checkbox-group>
								</view>
								<view class="order-options">
									<label class="checkbox-item">
										<checkbox color="#285AC8" value="true" :checked="submitForm.hasSevereMentalDisorders" @tap="handleCheckboxChange('hasSevereMentalDisorders')" />
										严重精神障碍
									</label>
									<label class="checkbox-item">
										<checkbox color="#285AC8" value="true" :checked="submitForm.hasTuberculosis" @tap="handleCheckboxChange('hasTuberculosis')" />
										结核病
									</label>
								<label class="checkbox-item">
								  <checkbox color="#285AC8" value="true" :checked="submitForm.hasHepatitis" @tap="handleCheckboxChange('hasHepatitis')"  />
								  肝炎
								</label>
								</view>
								<view class="order-options">
								<label class="checkbox-item">
								  <checkbox color="#285AC8" value="true" :checked="submitForm.hasOccupationalDisease" @tap="handleCheckboxChange('hasOccupationalDisease')" />
								  职业病
								</label>
								<label class="checkbox-item">
								  <checkbox color="#285AC8" value="true" :checked="submitForm.hasChronicKidneyDisease" @tap="handleCheckboxChange('hasChronicKidneyDisease')" />
								  慢性肾病
								</label>
								<label class="checkbox-item">
								  <checkbox color="#285AC8" value="true" :checked="submitForm.hasChronicLiverDisease" @tap="handleCheckboxChange('hasChronicLiverDisease')" />
								  慢性肝病
								</label>
								</view>
								<view class="order-options">
								<label class="checkbox-item">
								  <checkbox color="#285AC8" value="true" :checked="submitForm.hasImmunodeficiency" @tap="handleCheckboxChange('hasImmunodeficiency')" />
								  免疫缺陷类疾病
								</label>
								<label class="checkbox-item">
								  <checkbox color="#285AC8" value="true" :checked="submitForm.hasTyphus" @tap="handleCheckboxChange('hasTyphus')" />
								  斑疹伤寒
								</label>
								</view>
								<view class="order-options">
								<label class="checkbox-item">
								  <checkbox color="#285AC8" value="true" :checked="submitForm.isPostpartumInSixWeeks" @tap="handleCheckboxChange('isPostpartumInSixWeeks')"/>
								  产后（6周以内）
								</label>
								<label class="checkbox-item">
								  <checkbox color="#285AC8" value="true" :checked="submitForm.hasDustExposure" @tap="handleCheckboxChange('hasDustExposure')" />
								  粉尘接触史
								</label>
								</view>
								<label class="checkbox-item">
								  <checkbox color="#285AC8" value="true" :checked="submitForm.hasOtherDiseases" @tap="handleCheckboxChange('hasOtherDiseases')" />
								  其他
								</label>
								<label class="checkbox-item">
								  <input type="text" placeholder="请输入" v-model="submitForm.otherDiseasesName" class="text-input" style="float: right;"/>
								</label>
						      </checkbox-group>
						    </view>
						<view class="row">
							<view class="form-item">
								<text style="float: left;">吸烟史</text>
								<picker mode="selector" :range="staticForm.smokingStatusRange" @change="handleSmokingStatusChange">
									<view style="float: right;">{{ submitForm.smokingStatus || '请选择' }}</view>
								</picker>
							</view>
						</view>
						<view class="row">
							<view class="form-item">
								<text style="float: left;">饮酒史</text>
								<picker mode="selector" :range="staticForm.drinkingStatusRange" @change="handleDrinkingStatusChange">
									<view style="float: right;">{{ submitForm.drinkingStatus || '请选择' }}</view>
								</picker>
							</view>
						</view>
					</view>
					<view class="button-container">
					<button class="return-button" @tap="handlereturn">上一步</button>
					<button class="submit-button" @tap="handleSubmit">下一步</button>
					</view>
		</view>
			<!-- 第三个 -->
			<view class="section" v-show="pages == 4">
						<view class="sections1">
							<view class="section-title">
								<text class="label">身高体重及疫苗接种情况</text>
							</view>
						</view>
						<view class="row">
							<view class="form-item">
								<text style="float: left;">身高</text>
								<input type="text" placeholder="请输入身高(cm)" v-model="submitForm.height" class="text-input" style="float: right;"/>
							</view>
						</view>
						<view class="row">
							<view class="form-item">
								<text style="float: left;">体重</text>
								<input type="text" placeholder="请输入体重(kg)" v-model="submitForm.weight" class="text-input" style="float: right;"/>
							</view>
						</view>
						<text class="label" style="margin: 24px;">疫苗接种情况</text>
						<view class="checkboxs" >
						      <checkbox-group>
								  <view class="order-options">
						        <label >
						          <checkbox color="#285AC8" value="true" :checked="submitForm.isVaccinatedForCOVID" @tap="handleCheckboxChange('isVaccinatedForCOVID')"/>
						          新冠疫苗
						        </label>
						        <label>
						          <checkbox color="#285AC8" value="true" :checked="submitForm.isVaccinatedForFlu" @tap="handleCheckboxChange('isVaccinatedForFlu')"/>
						          流感疫苗
						        </label>
						        <label>
						          <checkbox color="#285AC8" value="true" :checked="submitForm.isVaccinatedForPlague" @tap="handleCheckboxChange('isVaccinatedForPlague')"/>
						          鼠疫疫苗
						        </label>
								</view>
								<view class="order-options">
						        <label>
						          <checkbox color="#285AC8" value="true" :checked="submitForm.isVaccinatedForBCG" @tap="handleCheckboxChange('isVaccinatedForBCG')" />
						          卡介疫苗
						        </label>
						        <label>
						          <checkbox color="#285AC8" value="true" :checked="submitForm.isVaccinatedForHepatitis" @tap="handleCheckboxChange('isVaccinatedForHepatitis')" />
						          肝炎疫苗
						        </label>
								</view>
						      </checkbox-group>
						    </view>
							<view class="button-container">
							<button class="return-button" @tap="handlereturn">上一步</button>
							<button class="submit-button" @tap="handleSubmit">提交</button>
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
  padding: 12px;
  border-radius: 8px;
  color: #333333;
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
	margin: 7px 30px;
}
.order-options{
	display: block; 
	margin: 15px 0; 
	display: flex; 
	justify-content: space-between; 
}
.inner-checkboxs{
	margin-left: 4vh;
}
.label {
  font-size: 16px;
  color: #333333;
  display: block;
}

.options {
  display: flex;
}

.radio, .checkbox {
  display: flex;
  align-items: center;
  margin-right: 10px;
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
.tabbar{
	z-index: 999;
}
.form-item{
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
.text-input{
	text-align: right;
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
input{
	opacity: 0.7;
}
</style>