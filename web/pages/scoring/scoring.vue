<script>
import { tip } from '../../utils/tip'
import { satisfyService, submitService } from '../../services/satisfy.ts'

const LEVELS = ['非常满意', '满意', '一般', '不满意', '非常不满意']

export default {
  data() {
    return {
      levels: LEVELS,
      surveys: [],
      selectedIndex: 0,
      loading: true,
      submitting: false,
      feedBackForm: {
        phoneNumber: '',
        satisfactionLevel: '非常满意',
      },
    }
  },
  computed: {
    selectedSurvey() {
      if (!this.surveys.length) return null
      const idx = Math.min(Math.max(this.selectedIndex, 0), this.surveys.length - 1)
      return this.surveys[idx]
    },
    surveyTitles() {
      return this.surveys.map((s) => s.surveyTitle || '未命名问卷')
    },
  },
  mounted() {
    this.loadSurveys()
  },
  methods: {
    loadSurveys() {
      this.loading = true
      satisfyService()
        .then((response) => {
          if (response.code !== 200 || !response.data) {
            this.surveys = []
            return
          }
          const data = response.data
          const list = Array.isArray(data.surveys) ? data.surveys : []
          if (list.length) {
            this.surveys = list
            this.selectedIndex = 0
            return
          }
          // legacy single-survey payload
          if (data.surveyID || data.surveyTitle) {
            this.surveys = [
              {
                surveyID: data.surveyID,
                surveyTitle: data.surveyTitle,
                surveyContent: data.surveyContent,
              },
            ]
            this.selectedIndex = 0
            return
          }
          this.surveys = []
        })
        .catch((error) => {
          console.error('Failed to fetch survey data:', error)
          this.surveys = []
          tip.fail('加载问卷失败')
        })
        .finally(() => {
          this.loading = false
        })
    },
    onSurveyPick(e) {
      this.selectedIndex = Number(e.detail.value) || 0
    },
    changeLevel(level) {
      this.feedBackForm.satisfactionLevel = level
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
    handleSubmit() {
      if (!this.selectedSurvey) {
        tip.fail('暂无可评价问卷')
        return
      }
      if (this.submitting) return
      this.submitting = true
      tip.loading('提交中')
      const submitData = {
        phoneNumber: this.feedBackForm.phoneNumber,
        ratingID: 0,
        ratingTime: this.formatDate(new Date()),
        satisfactionLevel: this.feedBackForm.satisfactionLevel,
        surveyID: this.selectedSurvey.surveyID,
      }
      submitService(submitData)
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
        return
      }
      tip.fail((response && response.msg) || '反馈提交失败')
    },
  },
}
</script>

<template>
  <view class="container">
    <image src="../../static/other-background.png" mode="widthFix" class="background-image" />

    <scroll-view v-if="!loading && selectedSurvey" scroll-y class="scroll" :show-scrollbar="false">
      <view class="card">
        <view class="card-title">选择问卷</view>
        <picker mode="selector" :range="surveyTitles" :value="selectedIndex" @change="onSurveyPick">
          <view class="picker-box">
            <text class="picker-text">{{ selectedSurvey.surveyTitle }}</text>
            <text class="picker-arrow">更换</text>
          </view>
        </picker>
      </view>

      <view class="card">
        <view class="card-title">问卷说明</view>
        <text class="content-text">{{ selectedSurvey.surveyContent || '暂无说明' }}</text>
      </view>

      <view class="card">
        <view class="card-title">您的评价</view>
        <view class="level-row">
          <view
            v-for="level in levels"
            :key="level"
            class="level-chip"
            :class="{ active: feedBackForm.satisfactionLevel === level }"
            @tap="changeLevel(level)"
          >
            {{ level }}
          </view>
        </view>
      </view>

      <button class="submit-btn" :disabled="submitting" @tap="handleSubmit">提交评价</button>
      <view class="bottom-space" />
    </scroll-view>

    <view v-else-if="loading" class="empty">
      <text>加载中…</text>
    </view>
    <view v-else class="empty">
      <text>暂无开放中的评分问卷</text>
    </view>

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
.card {
  background: #fff;
  border-radius: 14px;
  padding: 16px 14px;
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
.picker-box {
  display: flex;
  align-items: center;
  justify-content: space-between;
  min-height: 44px;
  padding: 10px 12px;
  border-radius: 10px;
  background: #f5f7fb;
}
.picker-text {
  flex: 1;
  font-size: 14px;
  color: #1a202c;
  padding-right: 8px;
}
.picker-arrow {
  font-size: 13px;
  color: #285ac8;
  flex-shrink: 0;
}
.content-text {
  display: block;
  font-size: 14px;
  color: #4a5568;
  line-height: 1.7;
  white-space: pre-wrap;
}
.level-row {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}
.level-chip {
  padding: 10px 14px;
  border-radius: 20px;
  background: #f5f7fb;
  font-size: 13px;
  color: #4a5568;
  border: 1px solid transparent;
}
.level-chip.active {
  background: rgba(40, 90, 200, 0.1);
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
.empty {
  position: relative;
  z-index: 1;
  height: 70vh;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #285ac8;
  font-size: 16px;
  font-weight: 600;
}
.tabbar {
  z-index: 999;
}
</style>
