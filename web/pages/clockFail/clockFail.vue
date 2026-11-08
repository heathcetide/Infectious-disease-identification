<script>
import store from '../../store/index.ts'
import { tip } from '../../utils/tip'
import { clockSituationService } from '../../services/clock'

const HIGH_RISK = 80

export default {
  data() {
    return {
      dateForm: { date: '' },
      diseases: [],
      hybrid: null,
      loading: false,
      highRisk: false,
      highRiskDisease: '',
      nearestHospital: '',
      alertShown: false,
    }
  },
  mounted() {
    const diseasescoring = store.state.diseasescoring || []
    this.diseases = this.formatScores(diseasescoring)
    this.hybrid = store.state.scoringHybrid || null
    this.applyHighRiskFromList(this.diseases, store.state.scoringHybrid)

    const now = new Date()
    this.dateForm.date = this.formatDate(now)
    this.loading = true
    clockSituationService(this.dateForm)
      .then((res) => this.checkSuccess(res))
      .catch((error) => this.checkError(error))
      .finally(() => {
        this.loading = false
        this.maybePromptHighRisk()
      })
  },
  methods: {
    formatDate(date) {
      const year = date.getFullYear()
      const month = String(date.getMonth() + 1).padStart(2, '0')
      const day = String(date.getDate()).padStart(2, '0')
      return `${year}-${month}-${day}`
    },
    formatScores(list) {
      return (list || [])
        .map((item) => ({
          diseaseTypeName: item.diseaseTypeName,
          score: Math.round(Number(item.score) || 0),
          reason: item.reason || '',
          evidence: item.evidence || [],
          link:
            '/pages/publicityMaterial/publicityMaterial?diseaseTypeName=' +
            encodeURIComponent(item.diseaseTypeName || ''),
        }))
        .sort((a, b) => b.score - a.score)
    },
    applyHighRiskFromList(list, hybrid) {
      const top = (list || [])[0]
      const fromFlag = !!(hybrid && hybrid.highRisk)
      const fromScore = !!(top && top.score >= HIGH_RISK)
      this.highRisk = fromFlag || fromScore
      this.highRiskDisease = top ? top.diseaseTypeName : ''
      this.nearestHospital = (hybrid && hybrid.nearestHospital) || ''
    },
    maybePromptHighRisk() {
      if (!this.highRisk || this.alertShown) return
      this.alertShown = true
      const name = this.highRiskDisease || '相关疾病'
      const hosp = this.nearestHospital ? `建议优先前往：${this.nearestHospital}。` : ''
      uni.showModal({
        title: '高风险预警',
        content: `「${name}」可能性已超过 ${HIGH_RISK}%，系统已上报后台预警。${hosp}是否查看就近医院？`,
        confirmText: '查看医院',
        cancelText: '稍后',
        success: (res) => {
          if (res.confirm) this.goNearbyHospital()
        },
      })
    },
    goNearbyHospital() {
      uni.navigateTo({
        url:
          '/pages/nearbyHospital/nearbyHospital?disease=' +
          encodeURIComponent(this.highRiskDisease || ''),
      })
    },
    checkSuccess(response) {
      const data = (response && response.data) || {}
      const list = data.diseasescoringList || []
      if (list.length) {
        this.diseases = this.formatScores(list)
        store.dispatch('setDiseasescoring', list)
      }
      if (data.hybrid) {
        this.hybrid = data.hybrid
        store.dispatch('setScoringHybrid', data.hybrid)
      }
      if (data.highRisk || data.nearestHospital) {
        this.highRisk = !!data.highRisk || this.highRisk
        this.nearestHospital = data.nearestHospital || this.nearestHospital
      }
      this.applyHighRiskFromList(this.diseases, {
        highRisk: this.highRisk,
        nearestHospital: this.nearestHospital,
      })
    },
    checkError(response) {
      console.log(response)
      tip.fail('加载打卡结果失败')
    },
    openProtection(item) {
      uni.navigateTo({ url: item.link })
    },
  },
}
</script>

<template>
  <view class="page">
    <image class="bg" src="/static/other-background.png" mode="aspectFill" />

    <view class="body">
      <view class="hero">
        <view class="hero-text">
          <view class="hero-title-row">
            <text class="hero-title">您已打卡</text>
            <image
              class="status-icon"
              src="/static/icon-Symptomatic.png"
              mode="aspectFit"
            />
          </view>
          <text class="hero-desc">根据您提交的症状，您可能的疾病是：</text>
        </view>
        <image class="hero-art" src="/static/virus.png" mode="aspectFit" />
      </view>

      <view v-if="highRisk" class="risk-banner" @tap="goNearbyHospital">
        <text class="risk-title">高风险预警（≥{{ 80 }}%）</text>
        <text class="risk-desc">
          「{{ highRiskDisease }}」已上报后台。
          {{ nearestHospital ? '就近医院：' + nearestHospital + '。' : '' }}点击查看医院列表 ›
        </text>
      </view>

      <view class="panel">
        <scroll-view scroll-y class="list-scroll" :show-scrollbar="false">
          <view v-if="diseases.length" class="list">
            <view
              class="list-item"
              v-for="(item, index) in diseases"
              :key="index"
              @tap="openProtection(item)"
            >
              <text class="index">{{ index + 1 }}</text>
              <view class="main">
                <text class="disease-name">{{ item.diseaseTypeName }}</text>
                <text v-if="item.reason" class="reason">{{ item.reason }}</text>
                <text class="protect-link">健康防护 ›</text>
              </view>
              <view class="rate-box">
                <text class="rate">{{ item.score }}%</text>
                <text class="rate-label">患病率</text>
              </view>
            </view>
          </view>
          <view v-else class="empty">
            <text>{{ loading ? '加载中…' : '暂无疾病风险记录' }}</text>
          </view>
        </scroll-view>
      </view>
    </view>

    <uni-custom-tabbar class="tabbar" />
  </view>
</template>

<style>
.page {
  position: relative;
  width: 100%;
  height: 100vh;
  overflow: hidden;
  background: #eef3fb;
}
.bg {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  z-index: 0;
}
.body {
  position: relative;
  z-index: 1;
  height: 100%;
  box-sizing: border-box;
  padding: 10vh 16px 72px;
  display: flex;
  flex-direction: column;
}
.hero {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 12px;
}
.hero-title {
  font-size: 22px;
  font-weight: 700;
  color: #0f172a;
}
.hero-title-row {
  display: flex;
  align-items: center;
  gap: 8px;
}
.status-icon {
  width: 28px;
  height: 28px;
}
.hero-desc {
  margin-top: 6px;
  font-size: 13px;
  color: #64748b;
  display: block;
}
.hero-art {
  width: 88px;
  height: 88px;
}
.risk-banner {
  margin-bottom: 10px;
  padding: 12px 14px;
  border-radius: 12px;
  background: #fff1f0;
  border: 1px solid #fca5a5;
}
.risk-title {
  display: block;
  font-size: 14px;
  font-weight: 700;
  color: #b91c1c;
}
.risk-desc {
  display: block;
  margin-top: 4px;
  font-size: 12px;
  color: #7f1d1d;
  line-height: 1.45;
}
.panel {
  flex: 1;
  min-height: 0;
  background: rgba(255, 255, 255, 0.92);
  border-radius: 16px;
  padding: 8px 4px;
}
.list-scroll {
  height: 100%;
}
.list-item {
  display: flex;
  align-items: flex-start;
  padding: 12px 10px;
  border-bottom: 1px solid #e2e8f0;
}
.index {
  width: 22px;
  color: #285ac8;
  font-weight: 700;
  margin-top: 2px;
}
.main {
  flex: 1;
  min-width: 0;
  padding-right: 8px;
}
.disease-name {
  display: block;
  font-size: 15px;
  font-weight: 600;
  color: #0f172a;
}
.reason {
  display: block;
  margin-top: 4px;
  font-size: 12px;
  color: #64748b;
  line-height: 1.4;
}
.protect-link {
  display: block;
  margin-top: 4px;
  font-size: 12px;
  color: #285ac8;
}
.rate-box {
  text-align: right;
  min-width: 64px;
}
.rate {
  display: block;
  font-size: 18px;
  font-weight: 700;
  color: #dc2626;
}
.rate-label {
  display: block;
  font-size: 11px;
  color: #94a3b8;
}
.empty {
  padding: 40px 16px;
  text-align: center;
  color: #94a3b8;
}
.tabbar {
  position: relative;
  z-index: 2;
}
</style>
