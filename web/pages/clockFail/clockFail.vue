<script>
import store from '../../store/index.ts'
import { clockSituationService } from '../../services/clock'

export default {
  data() {
    return {
      dateForm: { date: '' },
      diseases: [],
      loading: false,
    }
  },
  mounted() {
    const diseasescoring = store.state.diseasescoring || []
    this.diseases = this.formatScores(diseasescoring)

    const now = new Date()
    this.dateForm.date = this.formatDate(now)
    this.loading = true
    clockSituationService(this.dateForm)
      .then((res) => this.checkSuccess(res))
      .catch((error) => this.checkError(error))
      .finally(() => {
        this.loading = false
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
          link:
            '/pages/publicityMaterial/publicityMaterial?diseaseTypeName=' +
            encodeURIComponent(item.diseaseTypeName || ''),
        }))
        .sort((a, b) => b.score - a.score)
    },
    checkSuccess(response) {
      const list =
        (response && response.data && response.data.diseasescoringList) || []
      if (list.length) {
        this.diseases = this.formatScores(list)
      }
    },
    checkError(response) {
      console.log(response)
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
  left: 0;
  top: 0;
  width: 100%;
  height: 38vh;
  z-index: 0;
}
.body {
  position: relative;
  z-index: 1;
  height: calc(100vh - 11vh);
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
  padding: 9vh 0 0;
}
.hero {
  flex-shrink: 0;
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  padding: 0 20px 10px;
}
.hero-text {
  flex: 1;
  min-width: 0;
  padding-right: 8px;
  padding-bottom: 8px;
}
.hero-title-row {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}
.hero-title {
  font-size: 22px;
  font-weight: 700;
  color: #2a3348;
}
.status-icon {
  width: 22px;
  height: 22px;
  margin-left: 8px;
}
.hero-desc {
  display: block;
  font-size: 15px;
  line-height: 1.5;
  color: #4a5568;
  font-weight: 600;
}
.hero-art {
  width: 118px;
  height: 118px;
  flex-shrink: 0;
}
.panel {
  flex: 1;
  min-height: 0;
  background: #fff;
  border-radius: 28px 28px 0 0;
  box-shadow: 0 -6px 24px rgba(40, 90, 200, 0.08);
  padding: 18px 16px 8px;
  box-sizing: border-box;
}
.list-scroll {
  height: 100%;
}
.list {
  padding-bottom: 12px;
}
.list-item {
  display: flex;
  align-items: center;
  padding: 14px 10px;
  margin-bottom: 6px;
  border-radius: 12px;
  background: #f7f9fc;
}
.index {
  width: 28px;
  flex-shrink: 0;
  text-align: center;
  font-size: 18px;
  font-weight: 700;
  color: #c5cedb;
}
.main {
  flex: 1;
  min-width: 0;
  padding: 0 10px;
}
.disease-name {
  font-size: 15px;
  font-weight: 600;
  color: #1f2a44;
  line-height: 1.35;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  white-space: normal;
}
.protect-link {
  display: block;
  margin-top: 4px;
  font-size: 13px;
  color: #285ac8;
}
.rate-box {
  flex-shrink: 0;
  width: 64px;
  min-height: 52px;
  padding: 8px 4px;
  box-sizing: border-box;
  border-radius: 10px;
  background: #285ac8;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
.rate {
  font-size: 16px;
  font-weight: 700;
  color: #fff;
  line-height: 1.2;
  white-space: nowrap;
}
.rate-label {
  margin-top: 2px;
  font-size: 11px;
  color: rgba(255, 255, 255, 0.92);
  line-height: 1.2;
  white-space: nowrap;
}
.empty {
  padding: 48px 12px;
  text-align: center;
  color: #94a3b8;
  font-size: 14px;
}
.tabbar {
  position: relative;
  z-index: 10;
}
</style>
