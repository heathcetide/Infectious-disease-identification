<script>
import { tip } from '../../utils/tip'
import { clockEnvironmentService, environmentInfoService } from '../../services/clock'

const PRESSURE_OPTIONS = [
  { label: '70 kPa 以下', value: 65 },
  { label: '70–80 kPa', value: 75 },
  { label: '80–90 kPa', value: 85 },
  { label: '90–100 kPa', value: 95 },
  { label: '100 kPa 以上', value: 105 },
]

const TEMP_OPTIONS = [
  { label: '0℃ 以下', value: -5 },
  { label: '0–10℃', value: 5 },
  { label: '11–20℃', value: 15 },
  { label: '21–25℃', value: 23 },
  { label: '26–30℃', value: 28 },
  { label: '31–35℃', value: 33 },
  { label: '36–40℃', value: 38 },
  { label: '40℃ 以上', value: 42 },
]

function nearestOptionIndex(options, raw) {
  const n = Number(raw)
  if (!Number.isFinite(n) || raw === '' || raw == null) return -1
  let best = 0
  let bestDiff = Math.abs(options[0].value - n)
  for (let i = 1; i < options.length; i++) {
    const diff = Math.abs(options[i].value - n)
    if (diff < bestDiff) {
      best = i
      bestDiff = diff
    }
  }
  return best
}

export default {
  data() {
    return {
      clockEnvironmentForm: {
        avgWorkRoomPeople: '',
        avgWorkFaceToFace: '',
        avgWorkPhysicalContact: '',
        avgBreakfastRoomPeople: '',
        avgBreakfastFaceToFace: '',
        avgBreakfastPhysicalContact: '',
        avgLunchRoomPeople: '',
        avgLunchFaceToFace: '',
        avgLunchPhysicalContact: '',
        avgDinnerRoomPeople: '',
        avgDinnerFaceToFace: '',
        avgDinnerPhysicalContact: '',
        collectiveActivitiesCount: '',
        avgActivityRoomPeople: '',
        avgActivityFaceToFace: '',
        avgActivityPhysicalContact: '',
        avgMonthlyPressure: '',
        maxTemperature: '',
        minTemperature: '',
        submissionTime: '',
      },
      pressureOptions: PRESSURE_OPTIONS,
      tempOptions: TEMP_OPTIONS,
      pressureIndex: -1,
      maxTempIndex: -1,
      minTempIndex: -1,
      currentYearMonth: '',
      submitting: false,
    }
  },
  computed: {
    pressureLabel() {
      return this.pressureIndex >= 0
        ? this.pressureOptions[this.pressureIndex].label
        : '请选择气压区间'
    },
    maxTempLabel() {
      return this.maxTempIndex >= 0
        ? this.tempOptions[this.maxTempIndex].label
        : '请选择最高气温'
    },
    minTempLabel() {
      return this.minTempIndex >= 0
        ? this.tempOptions[this.minTempIndex].label
        : '请选择最低气温'
    },
  },
  created() {
    this.getCurrentYearMonth()
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
    getCurrentYearMonth() {
      const now = new Date()
      const year = now.getFullYear()
      const month = String(now.getMonth() + 1).padStart(2, '0')
      this.currentYearMonth = `${year}-${month}`
      environmentInfoService(this.currentYearMonth)
        .then((res) => this.getSuccess(res))
        .catch((error) => this.getError(error))
    },
    syncClimatePickers() {
      this.pressureIndex = nearestOptionIndex(
        this.pressureOptions,
        this.clockEnvironmentForm.avgMonthlyPressure
      )
      this.maxTempIndex = nearestOptionIndex(
        this.tempOptions,
        this.clockEnvironmentForm.maxTemperature
      )
      this.minTempIndex = nearestOptionIndex(
        this.tempOptions,
        this.clockEnvironmentForm.minTemperature
      )
    },
    onPressurePick(e) {
      const idx = Number(e.detail.value)
      this.pressureIndex = idx
      this.clockEnvironmentForm.avgMonthlyPressure = this.pressureOptions[idx].value
    },
    onMaxTempPick(e) {
      const idx = Number(e.detail.value)
      this.maxTempIndex = idx
      this.clockEnvironmentForm.maxTemperature = this.tempOptions[idx].value
    },
    onMinTempPick(e) {
      const idx = Number(e.detail.value)
      this.minTempIndex = idx
      this.clockEnvironmentForm.minTemperature = this.tempOptions[idx].value
    },
    getSuccess(response) {
      if (response && response.data != null) {
        this.clockEnvironmentForm = {
          ...this.clockEnvironmentForm,
          ...response.data,
        }
        this.syncClimatePickers()
      }
    },
    getError(error) {
      console.error('获取本月情况失败', error)
      tip.fail('获取信息失败')
    },
    handleSubmit() {
      if (this.submitting) return
      const allEmpty = Object.keys(this.clockEnvironmentForm)
        .filter((k) => k !== 'submissionTime')
        .every((k) => this.clockEnvironmentForm[k] === '' || this.clockEnvironmentForm[k] == null)
      if (allEmpty) {
        tip.fail('请填写本月情况')
        return
      }
      if (
        this.maxTempIndex >= 0 &&
        this.minTempIndex >= 0 &&
        this.tempOptions[this.maxTempIndex].value < this.tempOptions[this.minTempIndex].value
      ) {
        tip.fail('最高气温不能低于最低气温')
        return
      }
      this.submitting = true
      this.clockEnvironmentForm.submissionTime = this.formatDate(new Date())
      tip.loading('提交中')
      clockEnvironmentService(this.clockEnvironmentForm)
        .then((res) => this.submitSuccess(res))
        .catch((error) => this.submitError(error))
        .finally(() => {
          tip.hide()
          this.submitting = false
        })
    },
    submitError(error) {
      console.error('提交失败', error)
      tip.fail('提交失败')
    },
    submitSuccess(response) {
      if (response && response.code == 200) {
        tip.success('提交成功')
        setTimeout(() => {
          uni.redirectTo({ url: '/pages/clock/clock' })
        }, 800)
        return
      }
      tip.fail((response && response.msg) || '提交失败')
    },
  },
}
</script>

<template>
  <view class="container">
    <image src="../../static/other-background.png" mode="widthFix" class="background-image" />
    <scroll-view scroll-y class="scroll" :show-scrollbar="false">
      <view class="meta">登记月份 · {{ currentYearMonth }}</view>

      <view class="card">
        <view class="card-title">上周白天/晚上工作时间段</view>
        <view class="metric">
          <text class="metric-text">每次平均大约和</text>
          <input class="metric-input" type="number" v-model="clockEnvironmentForm.avgWorkRoomPeople" placeholder="0" />
          <text class="metric-text">人处于同一房间/空间</text>
        </view>
        <view class="metric">
          <text class="metric-text">大约和</text>
          <input class="metric-input" type="number" v-model="clockEnvironmentForm.avgWorkFaceToFace" placeholder="0" />
          <text class="metric-text">人进行面对面交流（1m 内）</text>
        </view>
        <view class="metric">
          <text class="metric-text">大约和</text>
          <input class="metric-input" type="number" v-model="clockEnvironmentForm.avgWorkPhysicalContact" placeholder="0" />
          <text class="metric-text">人有肢体接触</text>
        </view>
      </view>

      <view class="card">
        <view class="card-title">上周早餐时间段</view>
        <view class="metric">
          <text class="metric-text">每次平均大约和</text>
          <input class="metric-input" type="number" v-model="clockEnvironmentForm.avgBreakfastRoomPeople" placeholder="0" />
          <text class="metric-text">人处于同一房间/空间</text>
        </view>
        <view class="metric">
          <text class="metric-text">大约和</text>
          <input class="metric-input" type="number" v-model="clockEnvironmentForm.avgBreakfastFaceToFace" placeholder="0" />
          <text class="metric-text">人进行面对面交流（1m 内）</text>
        </view>
        <view class="metric">
          <text class="metric-text">大约和</text>
          <input class="metric-input" type="number" v-model="clockEnvironmentForm.avgBreakfastPhysicalContact" placeholder="0" />
          <text class="metric-text">人有肢体接触</text>
        </view>
      </view>

      <view class="card">
        <view class="card-title">上周午餐时间段</view>
        <view class="metric">
          <text class="metric-text">每次平均大约和</text>
          <input class="metric-input" type="number" v-model="clockEnvironmentForm.avgLunchRoomPeople" placeholder="0" />
          <text class="metric-text">人处于同一房间/空间</text>
        </view>
        <view class="metric">
          <text class="metric-text">大约和</text>
          <input class="metric-input" type="number" v-model="clockEnvironmentForm.avgLunchFaceToFace" placeholder="0" />
          <text class="metric-text">人进行面对面交流（1m 内）</text>
        </view>
        <view class="metric">
          <text class="metric-text">大约和</text>
          <input class="metric-input" type="number" v-model="clockEnvironmentForm.avgLunchPhysicalContact" placeholder="0" />
          <text class="metric-text">人有肢体接触</text>
        </view>
      </view>

      <view class="card">
        <view class="card-title">上周晚餐时间段</view>
        <view class="metric">
          <text class="metric-text">每次平均大约和</text>
          <input class="metric-input" type="number" v-model="clockEnvironmentForm.avgDinnerRoomPeople" placeholder="0" />
          <text class="metric-text">人处于同一房间/空间</text>
        </view>
        <view class="metric">
          <text class="metric-text">大约和</text>
          <input class="metric-input" type="number" v-model="clockEnvironmentForm.avgDinnerFaceToFace" placeholder="0" />
          <text class="metric-text">人进行面对面交流（1m 内）</text>
        </view>
        <view class="metric">
          <text class="metric-text">大约和</text>
          <input class="metric-input" type="number" v-model="clockEnvironmentForm.avgDinnerPhysicalContact" placeholder="0" />
          <text class="metric-text">人有肢体接触</text>
        </view>
      </view>

      <view class="card">
        <view class="card-title">集体活动</view>
        <view class="metric">
          <text class="metric-text">上周共参加</text>
          <input class="metric-input" type="number" v-model="clockEnvironmentForm.collectiveActivitiesCount" placeholder="0" />
          <text class="metric-text">次集体活动（如开会）</text>
        </view>
        <view class="metric">
          <text class="metric-text">每次平均大约和</text>
          <input class="metric-input" type="number" v-model="clockEnvironmentForm.avgActivityRoomPeople" placeholder="0" />
          <text class="metric-text">人处于同一房间/空间</text>
        </view>
        <view class="metric">
          <text class="metric-text">大约和</text>
          <input class="metric-input" type="number" v-model="clockEnvironmentForm.avgActivityFaceToFace" placeholder="0" />
          <text class="metric-text">人进行面对面交流（1m 内）</text>
        </view>
        <view class="metric">
          <text class="metric-text">大约和</text>
          <input class="metric-input" type="number" v-model="clockEnvironmentForm.avgActivityPhysicalContact" placeholder="0" />
          <text class="metric-text">人有肢体接触</text>
        </view>
      </view>

      <view class="card">
        <view class="card-title">气压气温</view>
        <view class="field">
          <text class="field-label">当月平均气压</text>
          <picker
            mode="selector"
            :range="pressureOptions"
            range-key="label"
            :value="pressureIndex < 0 ? 0 : pressureIndex"
            @change="onPressurePick"
          >
            <view class="field-picker" :class="{ placeholder: pressureIndex < 0 }">
              {{ pressureLabel }}
            </view>
          </picker>
        </view>
        <view class="field">
          <text class="field-label">当月最高气温</text>
          <picker
            mode="selector"
            :range="tempOptions"
            range-key="label"
            :value="maxTempIndex < 0 ? 0 : maxTempIndex"
            @change="onMaxTempPick"
          >
            <view class="field-picker" :class="{ placeholder: maxTempIndex < 0 }">
              {{ maxTempLabel }}
            </view>
          </picker>
        </view>
        <view class="field">
          <text class="field-label">当月最低气温</text>
          <picker
            mode="selector"
            :range="tempOptions"
            range-key="label"
            :value="minTempIndex < 0 ? 0 : minTempIndex"
            @change="onMinTempPick"
          >
            <view class="field-picker" :class="{ placeholder: minTempIndex < 0 }">
              {{ minTempLabel }}
            </view>
          </picker>
        </view>
      </view>

      <button class="submit-btn" :disabled="submitting" @tap="handleSubmit">提交</button>
      <view class="bottom-space" />
    </scroll-view>
    <uni-custom-tabbar class="tabbar" />
  </view>
</template>

<style>
.container {
  width: 100%;
  min-height: 100vh;
  position: relative;
  background: #eef3fb;
}
.background-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 0;
}
.scroll {
  position: relative;
  z-index: 1;
  height: 100vh;
  box-sizing: border-box;
  padding: 11vh 16px 0;
}
.meta {
  font-size: 13px;
  color: #5b6b85;
  margin-bottom: 12px;
  padding-left: 4px;
}
.card {
  background: #fff;
  border-radius: 14px;
  padding: 16px 14px 10px;
  margin-bottom: 12px;
  box-shadow: 0 4px 16px rgba(40, 90, 200, 0.06);
}
.card-title {
  font-size: 15px;
  font-weight: 600;
  color: #1f2a44;
  margin-bottom: 12px;
  padding-left: 8px;
  border-left: 3px solid #285ac8;
  line-height: 1.2;
}
.metric {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 6px 4px;
  margin-bottom: 12px;
  padding: 10px 12px;
  background: #f5f7fb;
  border-radius: 10px;
}
.metric-text {
  font-size: 13px;
  color: #4a5568;
  line-height: 1.5;
}
.metric-input {
  width: 56px;
  height: 32px;
  text-align: center;
  border-radius: 8px;
  background: #fff;
  font-size: 14px;
  color: #1a202c;
  border: 1px solid #d9e2f2;
}
.field {
  margin-bottom: 12px;
}
.field-label {
  display: block;
  font-size: 13px;
  color: #4a5568;
  margin-bottom: 6px;
}
.field-input {
  height: 42px;
  padding: 0 12px;
  border-radius: 10px;
  background: #f5f7fb;
  font-size: 15px;
  color: #1a202c;
}
.field-picker {
  height: 42px;
  line-height: 42px;
  padding: 0 12px;
  border-radius: 10px;
  background: #f5f7fb;
  font-size: 15px;
  color: #1a202c;
}
.field-picker.placeholder {
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
