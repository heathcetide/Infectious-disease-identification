<script>
import store from '../../store/index.ts'
import { postLoginService } from '../../services/login'
import { tip } from '../../utils/tip'
import { clockSituationService } from '../../services/clock'
import { redirectIfSignedToday, todayStr } from '../../utils/clockRoute'

export default {
  data() {
    return {
      loginForm: {
        phoneNumber: '13800000000',
        password: 'Passw0rd!',
      },
      submitting: false,
    }
  },
  methods: {
    async loginSuccess(response) {
      const ok = response && (response.code === 200 || response.code === 1)
      if (!ok) {
        await tip.alert(
          (response && (response.msg || response.message)) || '手机号或密码不正确',
          '登录失败'
        )
        return
      }
      const data = response.data || {}
      store.dispatch('setJwt', data.accessToken || data.token)
      const userName =
        data.userName || (data.user && data.user.name) || ''
      store.dispatch('setUserName', userName)
      try {
        const phone =
          data.phoneNumber ||
          (data.user && data.user.phoneNumber) ||
          (this.loginForm.phoneNumber || '').trim()
        if (phone) uni.setStorageSync('phoneNumber', phone)
      } catch (e) {
        /* ignore */
      }
      tip.success('登录成功')
      setTimeout(async () => {
        if (data.informationComplete === false) {
          const go = await tip.confirm('检测到您未完善信息，请前往完善', '完善资料')
          if (go) {
            uni.redirectTo({ url: '/pages/surveyInfo/surveyInfo' })
            return
          }
        }
        await this.goHomeByClockStatus()
      }, 450)
    },
    async goHomeByClockStatus() {
      try {
        const res = await clockSituationService({ date: todayStr() })
        const ok = res && (res.code === 200 || res.code === 1)
        if (ok && redirectIfSignedToday(res.data)) {
          return
        }
      } catch (e) {
        console.log('check signed status failed', e)
      }
      uni.redirectTo({ url: '/pages/clock/clock' })
    },
    async loginError(error) {
      console.error('登录失败', error)
      const msg =
        (error && (error.msg || error.message)) ||
        '登录失败，请检查手机号和密码'
      await tip.alert(msg, '登录失败')
    },
    Login() {
      const phone = (this.loginForm.phoneNumber || '').trim()
      const password = this.loginForm.password || ''
      if (!phone || !password) {
        tip.fail('请输入手机号和密码')
        return
      }
      if (!/^1\d{10}$/.test(phone)) {
        tip.fail('请输入正确的11位手机号')
        return
      }
      if (this.submitting) return
      this.submitting = true
      tip.loading('登录中')
      postLoginService({ phoneNumber: phone, password })
        .then((res) => this.loginSuccess(res))
        .catch((error) => this.loginError(error))
        .finally(() => {
          tip.hide()
          this.submitting = false
        })
    },
  },
}
</script>

<template>
  <view class="page">
    <image class="bg" src="/static/login-bg.jpg" mode="aspectFill" />
    <view class="shade" />

    <view class="hero">
      <image class="hero-logo" src="/static/logo.png" mode="aspectFit" />
      <text class="hero-title">突发传染病症状筛查与预警系统</text>
      <text class="hero-sub">每日打卡 · 风险预警 · 防护宣教</text>
    </view>

    <view class="panel">
      <view class="field">
        <text class="label">手机号</text>
        <view class="input-wrap">
          <input
            class="input"
            type="number"
            maxlength="11"
            placeholder="请输入手机号"
            placeholder-class="ph"
            v-model="loginForm.phoneNumber"
          />
        </view>
      </view>
      <view class="field">
        <text class="label">密码</text>
        <view class="input-wrap">
          <input
            class="input"
            password
            maxlength="32"
            placeholder="请输入密码"
            placeholder-class="ph"
            v-model="loginForm.password"
            @confirm="Login"
          />
        </view>
      </view>
      <button class="btn" :loading="submitting" :disabled="submitting" @tap="Login">
        登录
      </button>
      <text class="hint">高原工区职工健康筛查入口</text>
    </view>
  </view>
</template>

<style>
.page {
  position: relative;
  width: 100%;
  height: 100vh;
  overflow: hidden;
  background: #0b1a2e;
  display: flex;
  flex-direction: column;
}
.bg {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  z-index: 0;
}
.shade {
  position: absolute;
  inset: 0;
  z-index: 1;
  background: linear-gradient(
    180deg,
    rgba(6, 18, 36, 0.35) 0%,
    rgba(6, 18, 36, 0.15) 38%,
    rgba(6, 18, 36, 0.55) 70%,
    rgba(6, 18, 36, 0.82) 100%
  );
  pointer-events: none;
}
.hero {
  position: relative;
  z-index: 2;
  flex: 1 1 auto;
  min-height: 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 10vh 28px 0;
  box-sizing: border-box;
}
.hero-logo {
  width: 72px;
  height: 72px;
  margin-bottom: 14px;
  border-radius: 50%;
  overflow: hidden;
  background: transparent;
}
.hero-title {
  display: block;
  text-align: center;
  color: #fff;
  font-size: 20px;
  font-weight: 700;
  letter-spacing: 0.04em;
  line-height: 1.4;
  text-shadow: 0 2px 16px rgba(0, 0, 0, 0.35);
  margin-bottom: 8px;
}
.hero-sub {
  display: block;
  text-align: center;
  color: rgba(255, 255, 255, 0.82);
  font-size: 13px;
  letter-spacing: 0.08em;
}
.panel {
  position: relative;
  z-index: 3;
  flex: 0 0 auto;
  width: 100%;
  box-sizing: border-box;
  background: #ffffff;
  padding: 28px 24px calc(28px + env(safe-area-inset-bottom));
  border-radius: 22px 22px 0 0;
  box-shadow: 0 -10px 36px rgba(8, 24, 48, 0.22);
}
.field {
  margin-bottom: 16px;
}
.label {
  display: block;
  font-size: 13px;
  color: #64748b;
  font-weight: 600;
  margin-bottom: 8px;
  letter-spacing: 0.02em;
}
.input-wrap {
  background: #f1f5f9;
  border-radius: 12px;
  border: 1px solid #e2e8f0;
  overflow: hidden;
}
.input {
  width: 100%;
  height: 50px;
  padding: 0 14px;
  box-sizing: border-box;
  font-size: 16px;
  color: #0f172a;
  background: transparent;
}
.ph {
  color: #94a3b8;
}
.btn {
  margin-top: 8px;
  width: 100%;
  height: 50px;
  line-height: 50px;
  border-radius: 12px;
  background: #285ac8;
  color: #fff;
  font-size: 16px;
  font-weight: 600;
  border: none;
}
.btn::after {
  border: none;
}
.hint {
  display: block;
  margin-top: 14px;
  text-align: center;
  font-size: 12px;
  color: #94a3b8;
}
</style>
