<script>
import {clockEnvironmentService, environmentInfoService}from '../../services/clock'
export default {
  data() {
    return {
	  clockEnvironmentForm :{
		dormitoryArea: '',
		dormitoryBuildingNumber: '',
		dormitoryOccupants: '',
		dormitoryRoomNumber: '',
		ventilationCondition: null,
		windowOpenArea: '',
		workStationAltitude: '',
		workStationName: '',
	  },
      currentYearMonth: '', // 存储当前年份和月份
    };
  },
  created() {
    this.getCurrentYearMonth(); // 组件创建时获取当前年份和月份
  },
  methods: {
	  handleVentilationCondition(value){
		  this.clockEnvironmentForm.ventilationCondition = value;
		  console.log('选择', this.clockEnvironmentForm.ventilationCondition);
	  },
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
      // 处理提交逻辑
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
		console.error('信息提交失败', error);
		uni.showToast({ 
			title: '信息提交失败',
			icon: 'none', // 可选，默认为none，还有success, loading, warning, error等
			duration: 2000, // 显示持续时间，单位为毫秒，默认为1500
		});
	},
	submitSuccess(response) {
	    if (response.code == 1) {
	        uni.showToast({
	          title: '信息提交成功',
	          icon: 'success',
			  duration: 1000,
	        });
			uni.showModal({
			    title: '注意',
			    content: '请前往填写本月情况登记',
			    showCancel: false,
			    confirmText: '确定',
			    success: function (res) {
			        if (res.confirm) {
			           uni.redirectTo({
			               url: '/pages/registerSituation/registerSituation'
			           });
			        }
			    }
			});
	    } else {
	        uni.showModal({
	            title: '信息提交失败',
	            content: '信息提交失败，',
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
	    <view class="section">
	  	  <view class="sections1">
	  		  <view class="section-title">
	  		  	<text class="label">请填写信息</text>
	  		  </view>
			  <view class="row">
			  	<view class="form-item">
			  		<text style="float: left;">工务段所在地海拔高度</text>
			  		<input type="text" v-model="clockEnvironmentForm.workStationAltitude" placeholder="请输入所在地海拔高度(m)" class="text-input" style="float: right;"/>
			  	</view>
			  </view>
			  <view class="row">
			  	<view class="form-item">
			  		<text style="float: left;">工务段名称</text>
			  		<input type="text" v-model="clockEnvironmentForm.workStationName" placeholder="请输入工务段名称" class="text-input" style="float: right;"/>
			  	</view>
			  </view>
			  <view class="row">
			  	<view class="form-item">
			  		<text style="float: left;">宿舍</text>
			  		<input type="text" v-model="clockEnvironmentForm.dormitoryBuildingNumber" placeholder="请输入幢" class="text-input" style="float: right; margin-bottom: 6px;"/>
					<input type="text" v-model="clockEnvironmentForm.dormitoryRoomNumber" placeholder="请输入室" class="text-input" style="float: right;"/>
			  	</view>
			  </view>
			  <view class="row">
			  	<view class="form-item">
			  		<text style="float: left;">同室同居人数</text>
			  		<input type="text" placeholder="请输入人数" class="text-input" style="float: right;" v-model="clockEnvironmentForm.dormitoryOccupants"/>
			  	</view>
			  </view>
			  <view class="row">
			  	<view class="form-item">
			  		<text style="float: left;">宿舍面积</text>
			  		<input type="text" v-model="clockEnvironmentForm.dormitoryArea" placeholder="请输入宿舍面积(平方米)" class="text-input" style="float: right;"/>
			  	</view>
			  </view>
			  <view class="row">
			  	<view class="form-item">
			  		<text style="float: left;">宿舍窗户可打开面积</text>
			  		<input type="text" v-model="clockEnvironmentForm.windowOpenArea" placeholder="窗户可打开面积(平方米)" class="text-input" style="float: right;"/>
			  	</view>
			  </view>
			  <view class="section-title">
			  	<text>宿舍通风情况</text>
			  </view>
			  <radio-group>
			  <view style="margin: 5px 15px">
			    <radio color="#285AC8" @tap="handleVentilationCondition('不开窗通风')" :checked="clockEnvironmentForm.ventilationCondition == '不开窗通风'">不开窗通风</radio>
			  </view>
			  <view style="margin: 5px 15px">
			    <radio color="#285AC8" @tap="handleVentilationCondition('不定时开窗通风')" :checked="clockEnvironmentForm.ventilationCondition == '不定时开窗通风'">不定时开窗通风</radio>
			  </view>
			  <view style="margin: 5px 15px">
			    <radio color="#285AC8" @tap="handleVentilationCondition('每日开窗通风')" :checked="clockEnvironmentForm.ventilationCondition == '每日开窗通风'">每日开窗通风</radio>
			  </view>
			  </radio-group>
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
.tabbar{
	z-index: 999;
}
.content{
	position: relative;
	padding-top: 9vh;
	margin: 0 6px;
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
  margin-top: 30px;
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
	margin-left: 4vh;
}
.label {
  font-size: 16px;
  color: #333333;
}
input{
	opacity: 0.7;
}
.options {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between; /* 确保选项在一行内均匀分布 */
}
.text-input{
	text-align: right;
}
.checkbox {
  display: flex;
  align-items: center;
  margin: 5px; /* 添加四周的边距 */
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
.single-submit-button{
	height: 44px;
	background-color: #285AC8;
	color: #ffffff;
	text-align: center;
	border-radius: 8px;
	font-size: 16px;
	margin: 0 24px;
}
</style>