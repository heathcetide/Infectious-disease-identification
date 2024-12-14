<script>
import {clockEnvironmentService, environmentInfoService}from '../../services/clock'
export default {
  data() {
    return {
	  clockEnvironmentForm :{
		avgWorkRoomPeople:'',
		avgWorkFaceToFace: '',
		avgWorkPhysicalContact:'',
		avgBreakfastRoomPeople:'',
		avgBreakfastFaceToFace:'',
		avgBreakfastPhysicalContact:'',
		avgLunchRoomPeople:'',
		avgLunchFaceToFace:'',
		avgLunchPhysicalContact:'',
		avgDinnerRoomPeople:'',
		avgDinnerFaceToFace:'',
		avgDinnerPhysicalContact:'',
		collectiveActivitiesCount:'',
		avgActivityRoomPeople:'',
		avgActivityFaceToFace:'',
		avgActivityPhysicalContact:'',
		avgMonthlyPressure:'',
		maxTemperature:'',
		minTemperature:'',
		submissionTime:'',
	  },
      currentYearMonth: '' // 存储当前年份和月份
    };
  },
  created() {
    this.getCurrentYearMonth(); // 组件创建时获取当前年份和月份
  },
  methods: {
    formatDate(date) {
      const year = date.getFullYear();
      const month = String(date.getMonth() + 1).padStart(2, '0'); // 月份从0开始，所以加1，并格式化为两位数
      const day = String(date.getDate()).padStart(2, '0');
      const hours = String(date.getHours()).padStart(2, '0');
      const minutes = String(date.getMinutes()).padStart(2, '0');
      const seconds = String(date.getSeconds()).padStart(2, '0');
      return `${year}-${month}-${day}T${hours}:${minutes}:${seconds}`;
    },
    getCurrentYearMonth() {
      const now = new Date();
      const year = now.getFullYear(); // 获取当前年份
      const month = String(now.getMonth() + 1).padStart(2, '0'); // 获取当前月份（0-11），因此需要 +1，并格式化为两位数
      this.currentYearMonth = `${year}-${month}`; // 格式化为 'YYYY-MM'
      environmentInfoService(this.currentYearMonth)
        .then(res => this.getSuccess(res))
        .catch(error => this.getError(error));
    },
	getSuccess(response) {
		if(response.data != null){
			this.clockEnvironmentForm = response.data
			console.log(this.clockEnvironmentForm)
		}
	},
	goError(error) {
	    console.error('登录失败', error);
	},
    handleSubmit() {
	  // Check if all fields are empty
	  const allFieldsEmpty = Object.values(this.clockEnvironmentForm).every(value => value === '');
	  
	  if (allFieldsEmpty) {
	    uni.showToast({
	      title: '请填写基本信息',
	      icon: 'none',
	      duration: 2000,
	    });
	    return; // Stop the submission process
	  }
	  // 获取当前时间并格式化
	  const now = new Date();
	  const formattedDate = this.formatDate(now);
	  this.clockEnvironmentForm.submissionTime = formattedDate;
	  console.log("clockEnvironment",this.clockEnvironmentForm)
      clockEnvironmentService(this.clockEnvironmentForm)
		.then(res => this.submitSuccess(res))
		.catch(error => this.submitError(error));	
    },
	submitError(response) {
		console.error('密码修改失败', error);
		uni.showToast({ 
			title: '密码修改失败，请检查手机号和密码',
			icon: 'none', // 可选，默认为none，还有success, loading, warning, error等
			duration: 2000, // 显示持续时间，单位为毫秒，默认为1500
		});
	},
	submitSuccess(response) {
	    if (response.code == 1) {
	        uni.showToast({
	          title: '提交成功',
	          icon: 'success',
			  success:function(res){
			  	uni.redirectTo({
			  		url: '/pages/clock/clock'
			  	})
			  }
	        });
	    } else {
	        uni.showModal({
	            title: '提交失败',
	            content: '提交失败，请检查。',
	            showCancel: false,
	            confirmText: '确定',
	            success: function (res) {
	                if (res.confirm) {
	                    console.log('用户点击确定');
	                }
	            }
	        });
	    }
	},
  },
};
</script>

<template>
  <view class="container">
    <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
    <view class="content">
      <view class="header">
        <text class="title">基本</text>数据
      </view>
      <view style="margin-left: 10px;">
        <text>当前年月：{{ currentYearMonth }}</text> <!-- 显示动态获取的年月 -->
      </view>
      
      <!-- 添加新的内容区域 -->
      <view class="section">
        <view class="sections1">
          <view class="section-title">
            <text class="label">上周白天/晚上工作时间段</text>
          </view>
          <view class="row">
            <text class="label">每次平均大约和</text>
            <input type="text" v-model="clockEnvironmentForm.avgWorkRoomPeople" placeholder="请输入" class="input" />
            <text class="label">人处于同一房间/空间；</text>
          </view>
          <view class="row">
            <text class="label">大约和</text>
            <input type="text" v-model="clockEnvironmentForm.avgWorkFaceToFace" placeholder="请输入" class="input" />
            <text class="label">人进行了面对面语言交流（1m之内）；</text>
          </view>
          <view class="row">
            <text class="label">大约和</text>
            <input type="text" v-model="clockEnvironmentForm.avgWorkPhysicalContact" placeholder="请输入" class="input" />
            <text class="label">人进行了肢体接触（如握手）</text>
          </view>
        </view>
      </view>

      <!-- 添加新的内容区域 -->
      <view class="section">
        <view class="sections1">
          <view class="section-title">
            <text class="label">上周早餐时间段：</text>
          </view>
          <view class="row">
            <text class="label">每次平均大约和</text>
            <input type="text" v-model="clockEnvironmentForm.avgBreakfastRoomPeople" placeholder="请输入" class="input" />
            <text class="label">人处于同一房间/空间；</text>
          </view>
          <view class="row">
            <text class="label">大约和</text>
            <input type="text" v-model="clockEnvironmentForm.avgBreakfastFaceToFace" placeholder="请输入" class="input" />
            <text class="label">人进行了面对面语言交流（1m之内）；</text>
          </view>
          <view class="row">
            <text class="label">大约和</text>
            <input type="text" v-model="clockEnvironmentForm.avgBreakfastPhysicalContact" placeholder="请输入" class="input" />
            <text class="label">人进行了肢体接触（如握手）</text>
          </view>
        </view>
      </view>
      
	  <!-- 添加新的内容区域 -->
	  <view class="section">
	    <view class="sections1">
	      <view class="section-title">
	        <text class="label">上周午餐时间段：</text>
	      </view>
	      <view class="row">
	        <text class="label">每次平均大约和</text>
	        <input type="text" v-model="clockEnvironmentForm.avgLunchRoomPeople" placeholder="请输入" class="input" />
	        <text class="label">人处于同一房间/空间；</text>
	      </view>
	      <view class="row">
	        <text class="label">大约和</text>
	        <input type="text" v-model="clockEnvironmentForm.avgLunchFaceToFace" placeholder="请输入" class="input" />
	        <text class="label">人进行了面对面语言交流（1m之内）；</text>
	      </view>
	      <view class="row">
	        <text class="label">大约和</text>
	        <input type="text" v-model="clockEnvironmentForm.avgLunchPhysicalContact" placeholder="请输入" class="input" />
	        <text class="label">人进行了肢体接触（如握手）</text>
	      </view>
	    </view>
	  </view>
	  <!-- 添加新的内容区域 -->
	  <view class="section">
	    <view class="sections1">
	      <view class="section-title">
	        <text class="label">上周晚餐时间段：</text>
	      </view>
	      <view class="row">
	        <text class="label">每次平均大约和</text>
	        <input type="text" v-model="clockEnvironmentForm.avgDinnerRoomPeople" placeholder="请输入" class="input" />
	        <text class="label">人处于同一房间/空间；</text>
	      </view>
	      <view class="row">
	        <text class="label">大约和</text>
	        <input type="text" v-model="clockEnvironmentForm.avgDinnerFaceToFace" placeholder="请输入" class="input" />
	        <text class="label">人进行了面对面语言交流（1m之内）；</text>
	      </view>
	      <view class="row">
	        <text class="label">大约和</text>
	        <input type="text" v-model="clockEnvironmentForm.avgDinnerPhysicalContact" placeholder="请输入" class="input" />
	        <text class="label">人进行了肢体接触（如握手）</text>
	      </view>
	    </view>
	  </view>
	  <!-- 添加新的内容区域 -->
	  <view class="section">
	    <view class="sections1">
	      <view class="section-title">
			  <view class="row" style="width: 80%;">
	        <text class="label">上周共参加</text>
			<input type="text" v-model="clockEnvironmentForm.collectiveActivitiesCount" placeholder="请输入" class="input" />
			<text>次集体活动（如开会）</text>
			</view>
	      </view>
	      <view class="row">
	        <text class="label">每次平均大约和</text>
	        <input type="text" v-model="clockEnvironmentForm.avgActivityRoomPeople" placeholder="请输入" class="input" />
	        <text class="label">人处于同一房间/空间；</text>
	      </view>
	      <view class="row">
	        <text class="label">大约和</text>
	        <input type="text" v-model="clockEnvironmentForm.avgActivityFaceToFace" placeholder="请输入" class="input" />
	        <text class="label">人进行了面对面语言交流（1m之内）；</text>
	      </view>
	      <view class="row">
	        <text class="label">大约和</text>
	        <input type="text" v-model="clockEnvironmentForm.avgActivityPhysicalContact" placeholder="请输入" class="input" />
	        <text class="label">人进行了肢体接触（如握手）</text>
	      </view>
	    </view>
	  </view>
	  <view class="section-title">
	  	<text class="label">气压气温</text>
	  </view>
	  <view class="row">
	  	<view class="form-item">
	  		<text style="float: left;">当月平均气压</text>
	  		<input type="text" v-model="clockEnvironmentForm.avgMonthlyPressure" placeholder="请输入(kPa)" class="text-input" style="float: right;"/>
	  	</view>
	  </view>
	  <view class="row">
	  	<view class="form-item">
	  		<text style="float: left;">当月最高气温</text>
	  		<input type="text" v-model="clockEnvironmentForm.maxTemperature" placeholder="请输入(℃)" class="text-input" style="float: right;"/>
	  	</view>
	  </view>
	  <view class="row">
	  	<view class="form-item">
	  		<text style="float: left;">当月最低气温</text>
	  		<input type="text" v-model="clockEnvironmentForm.minTemperature" placeholder="请输入(℃)" class="text-input" style="float: right;"/>
	  	</view>
	  </view>
      <button class="single-submit-button" @tap="handleSubmit">提交</button>
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

.content {
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
.tabbar{
	z-index: 999;
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
  margin-top: 30px;
  margin-bottom: 20px;
  background-color: #ffffff;
  padding: 12px;
  border-radius: 8px;
}

.section-title {
  margin-left: -3px;
  padding-left: 4px;
  border-left: 8px solid #285AC8;
  margin-bottom: 16px;
}

.row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 6px;
  margin-bottom: 3px;
}

.label {
  font-size: 14px;
  color: #333333;
}

.input {
  border: 1px solid #ccc;
  border-radius: 4px;
  padding: 0px 8px;
  width: 40px;
  text-align: center;
}

.single-submit-button {
  height: 44px;
  background-color: #285AC8;
  color: #ffffff;
  text-align: center;
  border-radius: 8px;
  font-size: 16px;
  margin: 0 24px;
}
.text-input{
	text-align: right;
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
input{
	opacity: 0.7;
}
</style>
