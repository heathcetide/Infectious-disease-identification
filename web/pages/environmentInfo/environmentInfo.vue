<script>
import { tip } from '../../utils/tip'
import { clockEnvironmentService, environmentInfoService } from '../../services/clock'

const VENTILATION_OPTIONS = ['不开窗通风', '不定时开窗通风', '每日开窗通风']

export default {
  data() {
    return {
      ventilationOptions: VENTILATION_OPTIONS,
      clockEnvironmentForm: {
        dormitoryArea: '',
        dormitoryBuildingNumber: '',
        dormitoryOccupants: '',
        dormitoryRoomNumber: '',
        ventilationCondition: null,
        windowOpenArea: '',
        workStationAltitude: '',
        workStationName: '',
        submissionTime: '',
      },
      currentYearMonth: '',
      submitting: false,
    }
  },
  created() {
    this.getCurrentYearMonth()
  },
  methods: {
    handleVentilationCondition(value) {
      this.clockEnvironmentForm.ventilationCondition = value
    },
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
    getSuccess(response) {
      if (response && response.data != null) {
        this.clockEnvironmentForm = {
          ...this.clockEnvironmentForm,
          ...response.data,
        }
      }
    },
    getError(error) {
      console.error('获取环境信息失败', error)
      tip.fail('获取信息失败')
    },
    handleSubmit() {
      if (this.submitting) return
      const f = this.clockEnvironmentForm
      if (!f.workStationName && !f.workStationAltitude && !f.dormitoryBuildingNumber) {
        tip.fail('请至少填写工务段或宿舍信息')
        return
      }
      this.submitting = true
      f.submissionTime = this.formatDate(new Date())
      tip.loading('提交中')
      clockEnvironmentService(f)
        .then((res) => this.submitSuccess(res))
        .catch((error) => this.submitError(error))
        .finally(() => {
          tip.hide()
          this.submitting = false
        })
    },
    submitError(error) {
      console.error('信息提交失败', error)
      tip.fail('信息提交失败')
    },
    async submitSuccess(response) {
      if (response && response.code == 200) {
        tip.success('提交成功')
        const go = await tip.confirm('请继续填写本月情况登记', '注意')
        if (go) {
          uni.redirectTo({ url: '/pages/registerSituation/registerSituation' })
        }
        return
      }
      tip.fail((response && response.msg) || '信息提交失败')
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
        <view class="card-title">工务段信息</view>
        <view class="field">
          <text class="field-label">所在地海拔高度</text>
          <input
            class="field-input"
            type="digit"
            v-model="clockEnvironmentForm.workStationAltitude"
            placeholder="单位：米"
          />
        </view>
        <view class="field">
          <text class="field-label">工务段名称</text>
          <input
            class="field-input"
            type="text"
            v-model="clockEnvironmentForm.workStationName"
            placeholder="请输入工务段名称"
          />
        </view>
      </view>

      <view class="card">
        <view class="card-title">宿舍信息</view>
        <view class="field dual">
          <view class="field-half">
            <text class="field-label">幢号</text>
            <input
              class="field-input"
              type="text"
              v-model="clockEnvironmentForm.dormitoryBuildingNumber"
              placeholder="幢"
            />
          </view>
          <view class="field-half">
            <text class="field-label">室号</text>
            <input
              class="field-input"
              type="text"
              v-model="clockEnvironmentForm.dormitoryRoomNumber"
              placeholder="室"
            />
          </view>
        </view>
        <view class="field">
          <text class="field-label">同室同居人数</text>
          <input
            class="field-input"
            type="number"
            v-model="clockEnvironmentForm.dormitoryOccupants"
            placeholder="请输入人数"
          />
        </view>
        <view class="field">
          <text class="field-label">宿舍面积</text>
          <input
            class="field-input"
            type="digit"
            v-model="clockEnvironmentForm.dormitoryArea"
            placeholder="单位：平方米"
          />
        </view>
        <view class="field">
          <text class="field-label">窗户可打开面积</text>
          <input
            class="field-input"
            type="digit"
            v-model="clockEnvironmentForm.windowOpenArea"
            placeholder="单位：平方米"
          />
        </view>
      </view>

      <view class="card">
        <view class="card-title">宿舍通风情况</view>
        <view class="chip-row">
          <view
            v-for="opt in ventilationOptions"
            :key="opt"
            class="chip"
            :class="{ active: clockEnvironmentForm.ventilationCondition === opt }"
            @tap="handleVentilationCondition(opt)"
          >
            {{ opt }}
          </view>
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
  padding: 16px 14px 8px;
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
.field {
  margin-bottom: 12px;
}
.field.dual {
  display: flex;
  gap: 10px;
}
.field-half {
  flex: 1;
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
.chip-row {
  display: flex;
  flex-direction: column;
  gap: 8px;
  padding-bottom: 8px;
}
.chip {
  padding: 12px 14px;
  border-radius: 10px;
  background: #f5f7fb;
  font-size: 14px;
  color: #4a5568;
  border: 1px solid transparent;
}
.chip.active {
  background: rgba(40, 90, 200, 0.08);
  color: #285ac8;
  border-color: rgba(40, 90, 200, 0.35);
  font-weight: 600;
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
