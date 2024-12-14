<script>
import { satisfyService, submitService } from '../../services/satisfy.ts';
export default {
  data() {
    return {
      surveyData: null,
      feedBackForm: {
        feedbackDate: '',
        feedbackText: '',
        feedbackTitle: '',
        name: '',
        phoneNumber: '',
        satisfactionLevel: '非常满意',
      },
    };
  },
  mounted() {
    // Call satisfyService on page load
    satisfyService().then(response => {
	  console.log('666666',response)
      if (response.code === 1 && response.data) {
        this.surveyData = response.data;
        // Bind title and content to feedback form
        this.feedBackForm.feedbackTitle = this.surveyData.surveyTitle;
        this.feedBackForm.feedbackText = this.surveyData.surveyContent;
      } else {
        this.surveyData = null; // No data available
      }
    }).catch(error => {
      console.error('Failed to fetch survey data:', error);
      this.surveyData = null;
    });
  },
  methods: {
	changeLevel(date){
		this.feedBackForm.satisfactionLevel = date;
		console.log('选择',this.feedBackForm.satisfactionLevel);
	},
    // Format date method
    formatDate(date) {
      const year = date.getFullYear();
      const month = String(date.getMonth() + 1).padStart(2, '0');
      const day = String(date.getDate()).padStart(2, '0');
      const hours = String(date.getHours()).padStart(2, '0');
      const minutes = String(date.getMinutes()).padStart(2, '0');
      const seconds = String(date.getSeconds()).padStart(2, '0');
      return `${year}-${month}-${day}T${hours}:${minutes}:${seconds}`;
    },
    // Handle form submission
    handleSubmit() {
      const now = new Date();
      const formattedDate = this.formatDate(now);

      const submitData = {
        phoneNumber: this.feedBackForm.phoneNumber,
        ratingID: this.surveyData ? this.surveyData.surveyID : null,
        ratingTime: formattedDate,
        satisfactionLevel: this.feedBackForm.satisfactionLevel,
        surveyID: this.surveyData ? this.surveyData.surveyID : null,
        userID: 2001, // replace with actual user ID if available
      };
      submitService(submitData).then(this.submitSuccess).catch(this.submitError);
    },
    submitError(error) {
      console.error('反馈提交失败', error);
      uni.showToast({
        title: '反馈提交失败',
        icon: 'none',
        duration: 2000,
      });
    },
    submitSuccess(response) {
      if (response.code == 1) {
        uni.showToast({
          title: '反馈提交成功',
          icon: 'success',
		  duration: 2000,
		  success() {
			// 延迟3秒再跳转，3秒 = 3000毫秒
			setTimeout(() => {
				uni.navigateBack({
					delta: 1 // 返回上一级页面
				});
			}, 3000); // 延迟3秒
		  }
        });
      } else {
        uni.showModal({
          title: '反馈提交失败',
          content: '反馈提交失败',
          showCancel: false,
          confirmText: '确定',
          success: function (res) {
            if (res.confirm) {
              console.log('用户点击确定');
            }
          },
        });
      }
    },
  },
};
</script>

<template>
  <view class="container">
    <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
    <view class="content" v-if="surveyData">
      <!-- Survey form -->
      <view class="section">
        <view class="sections1">
          <view class="section-title">
            <text class="label">满意度调查标题：</text>
          </view>
          <view class="form-item" style="margin-top: 9px;">
            <input disabled="true" type="text" v-model="feedBackForm.feedbackTitle" class="input" readonly />
          </view>
        </view>
      </view>
      
      <view class="section">
        <view class="sections1">
          <view class="section-title">
            <text class="label">满意度调查内容：</text>
          </view>
          <view class="form-content" style="margin-top: 9px; margin-bottom: 0px;">
            <textarea disabled="true" v-model="feedBackForm.feedbackText" class="textarea" readonly></textarea>
          </view>
        </view>
      </view>

      <view>
        <view class="row" style="margin: 0 12px;">
          <radio-group class="options">
            <view class="order-options">
              <label class="radio">
                <radio color="#285AC8" v-model="feedBackForm.satisfactionLevel" value="非常满意" @click="changeLevel('非常满意')"/>非常满意
              </label>
              <label class="radio">
                <radio color="#285AC8" v-model="feedBackForm.satisfactionLevel" value="满意" @click="changeLevel('满意')"/>满意
              </label>
            </view>
            <view class="order-options">
			  <label class="radio">
				<radio color="#285AC8" v-model="feedBackForm.satisfactionLevel" value="一般" @click="changeLevel('一般')"/>一般
			  </label>
              <label class="radio">
                <radio color="#285AC8" v-model="feedBackForm.satisfactionLevel" value="不满意" @click="changeLevel('不满意')"/>不满意
              </label>
            </view>
			<view class="order-options">
			  <label class="radio">
			    <radio color="#285AC8" v-model="feedBackForm.satisfactionLevel" value="非常不满意" @click="changeLevel('非常不满意')"/>非常不满意
			  </label>
			</view>
          </radio-group>
        </view>
      </view>

      <button class="single-submit-button" @tap="handleSubmit">提交</button>
    </view>

    <!-- Message if no survey data available -->
    <view v-else class="none-page">
      <h1>暂无评分表</h1>
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
.radio{
	margin: 0 20px;
}
.content{
	position: relative;
	padding-top: 13vh;
	border-top-left-radius: 40px;
	border-top-right-radius: 40px;
	padding-bottom: 3vh;
}
.icons{
	width: 160px;
	position: absolute;
	right: 6px;
	top: 10vh;
}
.section {
  margin-bottom: 10px;
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
.form-item{
	height: 50px;
	background-color: #ffffff;
	display: flex;      
	align-items: center;    
	border-radius: 10px;
	padding-left: 11px;
}
.form-content{
  background-color: #ffffff;
  border-radius: 10px;
}
.textarea {
  width: 100%;
  padding: 8px;
  font-size: 14px;
  height: 260px;
  resize: none;
}
.label{
	font-weight: 600;
	font-size: 14px;
	color: rgba(51, 51, 51, 1);
}
.single-submit-button{
	height: 44px;
	background-color: #285AC8;
	color: #ffffff;
	text-align: center;
	border-radius: 8px;
	font-size: 16px;
	margin: 0 24px 100px 24px;
}
.order-options{
	display: block;
	margin: 15px 5px;
	display: flex; 
	justify-content: space-between;
}
.none-page {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f2f2f2;
  color: #285AC8;
  font-size: 25px;
  font-weight: 650;
}
</style>