<script>
import { rePasswordService } from '../../services/login'
import { tip } from '../../utils/tip'

export default {
  data() {
    return {
      formData: {
        oldPassword: '',
        newPassword: '',
        confirmPassword: '',
      },
      isPasswordValid: false,
      showPasswordRule: false,
      submitting: false,
    }
  },
  computed: {
    canSubmit() {
      return (
        this.isPasswordValid &&
        this.formData.newPassword === this.formData.confirmPassword &&
        this.formData.oldPassword.trim() !== ''
      )
    },
  },
  methods: {
    goBack() {
      if (getCurrentPages().length > 1) {
        uni.navigateBack({ delta: 1 })
      } else {
        uni.redirectTo({ url: '/pages/clock/clock' })
      }
    },
    async submitError(error) {
      console.error('密码修改失败', error)
      await tip.alert(
        (error && (error.msg || error.message)) || '密码修改失败，请检查原密码是否正确',
        '修改失败'
      )
    },
    async submitSuccess(response) {
      if (response && (response.code === 200 || response.code === 1)) {
        tip.success('密码修改成功')
        this.formData.oldPassword = ''
        this.formData.newPassword = ''
        this.formData.confirmPassword = ''
        this.isPasswordValid = false
        this.showPasswordRule = false
        return
      }
      await tip.alert(
        (response && (response.msg || response.message)) || '密码修改失败，请检查原密码是否正确',
        '修改失败'
      )
    },
    onNewPasswordInput() {
      this.showPasswordRule = true
      this.validatePassword()
    },
    validatePassword() {
      const password = this.formData.newPassword
      const regex = /^\d{8,}$/
      this.isPasswordValid = regex.test(password)
    },
    changePassword() {
      if (!this.canSubmit) {
        if (this.formData.oldPassword.trim() === '') {
          tip.fail('请输入原始密码')
        } else if (this.formData.newPassword !== this.formData.confirmPassword) {
          tip.fail('两次输入的密码不一致')
        } else if (!this.isPasswordValid) {
          tip.fail('密码不符合规则')
        } else {
          tip.fail('请完善信息')
        }
        return
      }
      if (this.submitting) return
      this.submitting = true
      tip.loading('提交中')
      rePasswordService(this.formData)
        .then((res) => this.submitSuccess(res))
        .catch((error) => this.submitError(error))
        .finally(() => {
          tip.hide()
          this.submitting = false
        })
    },
  },
}
</script>

<template>
  <view class="container">
    <image src="../../static/other-background.png" mode="widthFix" class="background-image" />
    <view class="form-container">
      <view class="form-item">
        <text class="label">原密码</text>
        <input type="password" placeholder="请输入原始密码" v-model="formData.oldPassword" class="input" />
      </view>
      <view class="form-item">
        <text class="label">新密码</text>
        <input
          type="password"
          placeholder="请输入新密码"
          v-model="formData.newPassword"
          class="input"
          @input="onNewPasswordInput"
        />
        <text
          v-if="showPasswordRule"
          class="password-rule"
          :class="{ invalid: !isPasswordValid }"
        >密码规则：至少8位数字</text>
      </view>
      <view class="form-item">
        <text class="label">再次输入新密码</text>
        <input
          type="password"
          placeholder="请再次输入新密码"
          v-model="formData.confirmPassword"
          class="input"
        />
        <text
          v-if="formData.newPassword !== formData.confirmPassword && formData.confirmPassword"
          class="error-text"
        >两次输入的密码不一致</text>
      </view>
      <button class="submit-btn" :disabled="submitting" @tap="changePassword">确认修改</button>
    </view>
    <uni-custom-tabbar class="tabbar" />
  </view>
</template>

<style>
.container {
  width: 100%;
  min-height: 100vh;
  position: relative;
  overflow: hidden;
}
.background-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 0;
}
.form-container {
  position: relative;
  z-index: 1;
  padding: 12vh 24px 14vh;
}
.form-item {
  margin-bottom: 18px;
}
.label {
  display: block;
  font-size: 14px;
  color: #333;
  margin-bottom: 8px;
  font-weight: 600;
}
.input {
  height: 44px;
  padding: 0 12px;
  border-radius: 8px;
  background: #fff;
  font-size: 15px;
}
.error-text {
  display: block;
  margin-top: 6px;
  color: #e53e3e;
  font-size: 12px;
}
.password-rule {
  display: block;
  margin-top: 6px;
  font-size: 12px;
  color: #718096;
}
.password-rule.invalid {
  color: #e53e3e;
}
.submit-btn {
  margin-top: 24px;
  height: 48px;
  line-height: 48px;
  border-radius: 10px;
  background: #285ac8;
  color: #fff;
  font-size: 16px;
  border: none;
}
.submit-btn::after {
  border: none;
}
.tabbar {
  z-index: 999;
}
</style>
