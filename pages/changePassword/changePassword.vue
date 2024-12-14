<script>
import { rePasswordService } from '../../services/login'

export default {
  data() {
    return {
      formData: {
        oldPassword: '',
        newPassword: '',
        confirmPassword: '',
      },
      isPasswordValid: false,
    };
  },
  computed: {
    canSubmit() {
      // 确保所有条件都符合
      return this.isPasswordValid &&
             this.formData.newPassword === this.formData.confirmPassword &&
             this.formData.oldPassword.trim() !== '';
    }
  },
  methods: {
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
	          title: '密码修改成功',
	          icon: 'success'
	        });
	    } else {
	        uni.showModal({
	            title: '密码修改失败',
	            content: '密码修改失败，请检查您的密码是否正确。',
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
    validatePassword() {
      const password = this.formData.newPassword;
      const regex = /^\d{8,}$/;
      this.isPasswordValid = regex.test(password);
    },
    changePassword() {
      if (!this.canSubmit) {
        let message = '请完善信息';

        if (this.formData.oldPassword.trim() === '') {
          message = '请输入原始密码';
        } else if (this.formData.newPassword !== this.formData.confirmPassword) {
          message = '两次输入的密码不一致';
        } else if (!this.isPasswordValid) {
          message = '密码不符合规则';
        }
	  }
		rePasswordService(this.formData)
			.then(res => this.submitSuccess(res))
			.catch(error => this.submitError(error));
      // 清空表单数据
      this.formData.oldPassword = '';
      this.formData.newPassword = '';
      this.formData.confirmPassword = '';
      this.isPasswordValid = false;
    }
  }
}
</script>

<template>
  <view class="container">
    <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
    <view class="form-container">
      <view class="form-item">
        <text class="label">原密码：</text>
        <input type="password" placeholder="请输入原始密码" v-model="formData.oldPassword" class="input"/>
      </view>
      <view class="form-item">
        <text class="label">新密码：</text>
        <input type="password" placeholder="请输入新密码" v-model="formData.newPassword" class="input" @input="validatePassword"/>
      </view>
      <view class="form-item">
        <text class="label">再次输入新密码：</text>
        <input type="password" placeholder="请再次输入新密码" v-model="formData.confirmPassword" class="input"/>
        <text v-if="formData.newPassword !== formData.confirmPassword && formData.confirmPassword" class="error-text">两次输入的密码不一致</text>
        <view>
          <text class="password-rule" :class="{ invalid: !isPasswordValid }">密码规则：数字和字母组合，至少8位</text>
        </view>
      </view>
      <button class="submit-button" :disabled="!canSubmit" @tap="changePassword">提交</button>
    </view>
    <uni-custom-tabbar class="tabbar"></uni-custom-tabbar>
  </view>
</template>

<style scoped>
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

.tabbar {
  position: fixed; /* 固定到页面底部 */
  bottom: 0;
  left: 0;
  right: 0;
  height: 11vh; /* 与原来的 footer 样式相匹配 */
  z-index: 999;
  background-color: #ffffff;
  box-shadow: 0 -2px 4px rgba(0, 0, 0, 0.1);
}

.form-container {
  position: relative; /* 设置为相对定位，避免固定位缓冲与 tabbar 重叠 */
  bottom: 0;
  padding: 24vh 20px; /* 下移形式，使表单位置下移 */
  padding-bottom: 15vh; /* 增加底部间距，避免与 tabbar 重叠 */
  overflow-y: auto; /* 如果内容超出容器，则显示滚动条 */
  z-index: 20;
}

.form-item {
  margin: 30px 20px;
}

.label {
  font-size: 16px;
  color: #333333;
  margin-bottom: 5px;
  display: block;
}

.input {
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 5px;
  font-size: 14px;
  margin-bottom: 5px;
  border: none;
  background-color: #FFFFFF;
  height: 33px;
}
input {
  opacity: 0.7;
}
.password-rule {
  font-size: 12px;
  color: #999;
}

.password-rule.invalid {
  color: #A6A6A6;
}
.error-text {
  font-size: 12px;
  color: #ff4d4f;
  margin-top: 5px;
}

.submit-button {
  padding: 10px;
  background-color: #285AC8;
  color: #ffffff;
  border-radius: 5px;
  text-align: center;
  font-size: 16px;
  margin: 60px 20px;
}
</style>
