<script>
import { tip } from '../../utils/tip'
import { feedBackService } from '../../services/login'

export default {
  data() {
    return {
      submitting: false,
      feedBackForm: {
        feedbackDate: '',
        feedbackText: '',
        feedbackTitle: '',
        name: '',
        phoneNumber: '',
      },
    }
  },
  methods: {
    formatDate(date) {
      const year = date.getFullYear()
      const month = String(date.getMonth() + 1).padStart(2, '0')
      const day = String(date.getDate()).padStart(2, '0')
      const hours = String(date.getHours()).padStart(2, '0')
      const minutes = String(date.getMinutes()).padStart(2, '0')
      const seconds = String(date.getSeconds()).padStart(2, '0')
      return `${year}-${month}-${day}T${hours}:${minutes}:${seconds}`
    },
    handleSubmit() {
      if (this.feedBackForm.feedbackTitle.length < 3) {
        tip.fail('标题不得少于3个字')
        return
      }
      if (this.feedBackForm.feedbackText.length < 10) {
        tip.fail('内容不得少于10个字')
        return
      }
      if (this.submitting) return
      this.submitting = true
      this.feedBackForm.feedbackDate = this.formatDate(new Date())
      tip.loading('提交中')
      feedBackService(this.feedBackForm)
        .then((res) => this.submitSuccess(res))
        .catch((error) => this.submitError(error))
        .finally(() => {
          tip.hide()
          this.submitting = false
        })
    },
    submitError(error) {
      console.error('反馈提交失败', error)
      tip.fail((error && (error.msg || error.message)) || '反馈提交失败')
    },
    submitSuccess(response) {
      if (response && response.code == 200) {
        tip.success('反馈提交成功')
        this.feedBackForm.feedbackTitle = ''
        this.feedBackForm.feedbackText = ''
        return
      }
      tip.fail((response && response.msg) || '反馈提交失败')
    },
  },
}
</script>

<template>
  <view class="page">
    <image class="bg" src="/static/other-background.png" mode="aspectFill" />
    <scroll-view scroll-y class="scroll" :show-scrollbar="false">
      <view class="hero">
        <view class="hero-text">
          <text class="hero-kicker">意见反馈</text>
          <text class="hero-title">告诉我们您的建议</text>
          <text class="hero-desc">问题、体验或改进想法都可以写在这里</text>
        </view>
        <image class="hero-art" src="/static/feedback.png" mode="aspectFit" />
      </view>

      <view class="card">
        <text class="card-title">反馈标题</text>
        <input
          class="input"
          type="text"
          maxlength="50"
          placeholder="请用一句话概括（至少3个字）"
          placeholder-class="ph"
          v-model="feedBackForm.feedbackTitle"
        />
      </view>

      <view class="card">
        <text class="card-title">反馈内容</text>
        <textarea
          class="textarea"
          maxlength="500"
          placeholder="请详细描述您遇到的问题或建议（至少10个字）"
          placeholder-class="ph"
          v-model="feedBackForm.feedbackText"
        />
        <text class="count">{{ feedBackForm.feedbackText.length }}/500</text>
      </view>

      <button class="submit-btn" :disabled="submitting" @tap="handleSubmit">提交反馈</button>
      <view class="bottom-space" />
    </scroll-view>
    <uni-custom-tabbar class="tabbar" />
  </view>
</template>

<style>
.page {
  position: relative;
  width: 100%;
  min-height: 100vh;
  background: #eef3fb;
}
.bg {
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 42vh;
  z-index: 0;
}
.scroll {
  position: relative;
  z-index: 1;
  height: 100vh;
  box-sizing: border-box;
  padding: 11vh 16px 0;
}
.hero {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  margin-bottom: 14px;
  padding: 4px 4px 0;
}
.hero-text {
  flex: 1;
  padding-right: 8px;
}
.hero-kicker {
  display: block;
  font-size: 12px;
  color: #5b6b85;
  margin-bottom: 4px;
}
.hero-title {
  display: block;
  font-size: 24px;
  font-weight: 700;
  color: #1f2a44;
  margin-bottom: 6px;
}
.hero-desc {
  display: block;
  font-size: 13px;
  color: #718096;
}
.hero-art {
  width: 110px;
  height: 90px;
  flex-shrink: 0;
}
.card {
  background: #fff;
  border-radius: 14px;
  padding: 16px 14px;
  margin-bottom: 12px;
  box-shadow: 0 4px 16px rgba(40, 90, 200, 0.06);
}
.card-title {
  display: block;
  font-size: 15px;
  font-weight: 600;
  color: #1f2a44;
  margin-bottom: 12px;
  padding-left: 8px;
  border-left: 3px solid #285ac8;
  line-height: 1.2;
}
.input {
  width: 100%;
  height: 46px;
  padding: 0 12px;
  box-sizing: border-box;
  border-radius: 10px;
  background: #f5f7fb;
  font-size: 15px;
  color: #1a202c;
}
.textarea {
  width: 100%;
  min-height: 180px;
  padding: 12px;
  box-sizing: border-box;
  border-radius: 10px;
  background: #f5f7fb;
  font-size: 15px;
  color: #1a202c;
  line-height: 1.6;
}
.ph {
  color: #a0aec0;
}
.count {
  display: block;
  text-align: right;
  margin-top: 8px;
  font-size: 12px;
  color: #94a3b8;
}
.submit-btn {
  margin-top: 8px;
  height: 48px;
  line-height: 48px;
  border-radius: 12px;
  background: #285ac8;
  color: #fff;
  font-size: 16px;
  border: none;
}
.submit-btn::after {
  border: none;
}
.bottom-space {
  height: 14vh;
}
.tabbar {
  z-index: 999;
}
</style>
