<script>
import { feedBackService } from '../../services/login';
export default {
  data() {
    return {
      feedBackForm: {
        feedbackDate: '',
        feedbackText: '',
        feedbackTitle: '',
        name: '',
        phoneNumber: '',
      },
    };
  },
  methods: {
    // 自定义格式化日期的方法
    formatDate(date) {
      const year = date.getFullYear();
      const month = String(date.getMonth() + 1).padStart(2, '0'); // 月份从0开始，所以加1
      const day = String(date.getDate()).padStart(2, '0');
      const hours = String(date.getHours()).padStart(2, '0');
      const minutes = String(date.getMinutes()).padStart(2, '0');
      const seconds = String(date.getSeconds()).padStart(2, '0');

      return `${year}-${month}-${day}T${hours}:${minutes}:${seconds}`;
    },
    handleSubmit() {
      // 校验标题和内容的长度
      if (this.feedBackForm.feedbackTitle.length < 3) {
        uni.showToast({
          title: '标题不得少于3个字',
          icon: 'none',
          duration: 2000,
        });
        return;
      }
      if (this.feedBackForm.feedbackText.length < 10) {
        uni.showToast({
          title: '内容不得少于10个字',
          icon: 'none',
          duration: 2000,
        });
        return;
      }

      // 获取当前时间并格式化
      const now = new Date();
      const formattedDate = this.formatDate(now);

      // 设置到 feedBackForm 中
      this.feedBackForm.feedbackDate = formattedDate;

      feedBackService(this.feedBackForm)
        .then(res => this.submitSuccess(res))
        .catch(error => this.submitError(error));
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
		<view class="content">
		
		<!-- 第一个 -->
		<view class="section">
					<view class="sections1">
						<view class="section-title">
							<text class="label">您反馈意见标题：</text>
						</view>
						<image src="../../static/feedback.png" mode="widthFix" class="icons"></image>
						<view class="form-item" style="margin-top: 9px;">
							<input type="text" placeholder="您反馈意见的标题内容..." v-model="feedBackForm.feedbackTitle" class="input"/>
						</view>
					</view>
		</view>
		<!-- 第二个 -->
		<view class="section">
					<view class="sections1">
						<view class="section-title">
							<text class="label">您反馈意见的内容：</text>
						</view>
						<view class="form-content" style="margin-top: 9px; margin-bottom: 32px;">
							<textarea placeholder="请输入反馈内容" v-model="feedBackForm.feedbackText" class="textarea"></textarea>
						</view>
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
	padding-top: 18vh;
	border-top-left-radius: 40px;
	border-top-right-radius: 40px;
	padding-bottom: 13vh;
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
	margin: 0 24px;
}
</style>