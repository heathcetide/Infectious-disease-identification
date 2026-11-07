<!-- 引入 sockjs-client 和 stompjs -->
<script src="https://cdn.jsdelivr.net/npm/sockjs-client@1/dist/sockjs.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/stompjs@2.3.3/lib/stomp.min.js"></script>
<script>
import { tip } from '../../utils/tip'
import {
  clockSubmitService,
  clockSituationService,
  reverseGeocodeService,
} from '../../services/clock'
import { redirectIfSignedToday, todayStr } from '../../utils/clockRoute'
import store from '../../store/index'

export default {
  data() {
    return {
      clockForm: {
        longitude: 94.9053,
        latitude: 36.4016,
        city: '',
        county: '',
        health: true,
        locationName: '',
        province: '',
      },
      location: '正在获取位置...',
      dateForm: {
        date: '',
      },
      clockInfo: {
        hasSignedToday: false,
      },
      pageReady: false,
      locating: false,
    }
  },
  mounted() {
    this.dateForm.date = todayStr()
    clockSituationService(this.dateForm)
      .then((res) => this.checkSuccess(res))
      .catch((error) => this.checkError(error))
  },
  created() {
    this.getLoca()
  },
  methods: {
    getLoca() {
      if (this.locating) return
      this.locating = true
      const that = this
      uni.getLocation({
        type: 'gcj02',
        success(res) {
          that.clockForm.latitude = res.latitude
          that.clockForm.longitude = res.longitude
          that.reverseGeocode(res.latitude, res.longitude)
        },
        fail(res) {
          console.log('获取定位失败：' + res.errMsg)
          that.location = '定位失败，请开启定位权限后刷新'
          that.locating = false
        },
      })
    },
    reverseGeocode(latitude, longitude) {
      reverseGeocodeService(latitude, longitude)
        .then((res) => {
          const data = (res && res.data) || {}
          const name =
            data.locationName ||
            [data.province, data.city, data.county].filter(Boolean).join('')
          if (name) {
            this.location = name
            this.clockForm.locationName = data.locationName || name
            this.clockForm.province = data.province || ''
            this.clockForm.city = data.city || ''
            this.clockForm.county = data.county || ''
            if (typeof data.latitude === 'number') this.clockForm.latitude = data.latitude
            if (typeof data.longitude === 'number') this.clockForm.longitude = data.longitude
          } else {
            this.location = `${latitude.toFixed(5)}, ${longitude.toFixed(5)}`
          }
        })
        .catch((err) => {
          console.log('逆地理编码失败', err)
          this.location = `${latitude.toFixed(5)}, ${longitude.toFixed(5)}`
        })
        .finally(() => {
          this.locating = false
        })
    },
    refreshLocation() {
      tip.fail('正在刷新定位')
      this.location = '正在获取位置...'
      this.getLoca()
    },
    checkSuccess(response) {
      const ok = response && (response.code === 200 || response.code === 1)
      const data = (response && response.data) || {}
      this.clockInfo = data
      if (!ok) {
        this.pageReady = true
        return
      }
      if (!data.hasMonthSigned) {
        uni.showModal({
          title: '注意',
          content: '检测到您未填写当月工作环境基本信息，请前往填写',
          showCancel: false,
          confirmText: '确定',
          success: (res) => {
            if (res.confirm) {
              uni.redirectTo({ url: '/pages/workEnvironment/workEnvironment' })
            }
          },
        })
        return
      }
      if (redirectIfSignedToday(data)) {
        return
      }
      this.pageReady = true
    },
    checkError(response) {
      console.log(response)
      this.pageReady = true
    },
    submitError(response) {
      console.error('打卡失败', response)
      tip.fail((response && (response.msg || response.message)) || '打卡失败，请稍后重试')
    },
    submitSuccess(response) {
      if (response && (response.code == 200 || response.code == 1)) {
        tip.success('打卡成功')
        setTimeout(() => {
          uni.redirectTo({ url: '/pages/clockSuccess/clockSuccess' })
        }, 600)
      } else {
        tip.fail((response && (response.msg || response.message)) || '打卡失败')
      }
    },
    handleHealth() {
      this.dateForm.date = todayStr()
      clockSituationService(this.dateForm)
        .then((res) => this.checkClock1(res))
        .catch((error) => this.checkError(error))
    },
    handleSymptom() {
      this.clockForm.health = false
      store.dispatch('setClockFormParams', this.clockForm)
      this.dateForm.date = todayStr()
      clockSituationService(this.dateForm)
        .then((res) => this.checkClock2(res))
        .catch((error) => this.checkError(error))
    },
    checkClock2(response) {
      const data = (response && response.data) || {}
      if ((response.code === 200 || response.code === 1) && !data.hasSignedToday) {
        uni.navigateTo({
          url: '/pages/todaySymptoms/todaySymptoms',
        })
        return
      }
      redirectIfSignedToday(data)
    },
    checkClock1(response) {
      const data = (response && response.data) || {}
      if ((response.code === 200 || response.code === 1) && !data.hasSignedToday) {
        this.clockForm.health = true
        clockSubmitService(this.clockForm)
          .then((res) => this.submitSuccess(res))
          .catch((error) => this.submitError(error))
        return
      }
      redirectIfSignedToday(data)
    },
  },
}
</script>

<template>
  <view class="page">
    <image src="/static/other-background.png" mode="aspectFill" class="background-image" />
    <view v-if="pageReady" class="content">
      <view class="location-row">
        <image class="pin-icon" src="/static/pin-map.svg" mode="aspectFit" />
        <text class="location-address">您的位置：{{ location }}</text>
      </view>

      <view class="map-shell">
        <map
          class="clock-map"
          :latitude="clockForm.latitude"
          :longitude="clockForm.longitude"
          :scale="14"
          :show-location="true"
          :enable-scroll="false"
          :enable-zoom="false"
          :enable-poi="true"
        />
      </view>

      <button class="refresh-button" :disabled="locating" @tap="refreshLocation">
        <image class="refresh-icon" src="/static/arrow-repeat.svg" mode="aspectFit" />
        {{ locating ? '定位中…' : '刷新位置' }}
      </button>

      <view class="action-block">
        <view class="calendar-card">
          <text class="calendarText">请点击打卡</text>
          <image src="/static/calendar.png" mode="aspectFit" class="calendarImg" />
        </view>
        <view class="options">
          <button @tap="handleHealth" class="option-button healthy">
            <image src="/static/icon-health.png" mode="aspectFit" class="clockIcon" />
            <text>健康</text>
          </button>
          <button @tap="handleSymptom" class="option-button symptom">
            <image src="/static/icon-Symptomatic.png" mode="aspectFit" class="clockIcon" />
            <text>有症状</text>
          </button>
        </view>
      </view>
    </view>
    <uni-custom-tabbar class="tabbar" />
  </view>
</template>

<style>
.page {
  width: 100%;
  height: 100vh;
  overflow: hidden;
  position: relative;
  box-sizing: border-box;
}
.background-image {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  z-index: 0;
}
.content {
  position: relative;
  z-index: 1;
  height: calc(100vh - 11vh);
  padding: 7vh 16px 8px;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}
.location-row {
  display: flex;
  align-items: center;
  gap: 6px;
  margin-bottom: 14px;
  flex-shrink: 0;
}
.pin-icon {
  width: 16px;
  height: 16px;
}
.location-address {
  flex: 1;
  font-size: 13px;
  color: #555;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.map-shell {
  width: 100%;
  flex: 1 1 auto;
  min-height: 180px;
  border-radius: 12px;
  overflow: hidden;
  background: #eef2f7;
}
.clock-map {
  width: 100%;
  height: 100%;
}
.refresh-button {
  margin-top: 8px;
  margin-bottom: 0;
  padding: 0;
  font-size: 13px;
  color: #333;
  background: transparent;
  border: none;
  line-height: 1.4;
  flex-shrink: 0;
}
.refresh-button::after {
  border: none;
}
.refresh-icon {
  width: 16px;
  height: 16px;
  margin-right: 4px;
  vertical-align: middle;
}
.action-block {
  flex-shrink: 0;
  margin-top: 6px;
  padding-bottom: 4px;
}
.calendar-card {
  position: relative;
  width: 100%;
  height: 44px;
  margin-bottom: 4px;
}
.calendarText {
  font-size: 15px;
  color: #285ac8;
  font-weight: 600;
  line-height: 44px;
}
.calendarImg {
  width: 92px;
  height: 92px;
  position: absolute;
  top: -28px;
  right: -12px;
  opacity: 0.92;
  pointer-events: none;
}
.options {
  margin-top: 0;
}
.option-button {
  width: 100%;
  height: 50px;
  border-radius: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  font-size: 16px;
  font-weight: 700;
  background-color: #fff;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.08);
  margin-bottom: 10px;
  border: none;
}
.option-button::after {
  border: none;
}
.option-button.healthy {
  color: #405ac8;
}
.option-button.symptom {
  color: orangered;
}
.clockIcon {
  width: 22px;
  height: 22px;
}
.tabbar {
  z-index: 999;
}
</style>
