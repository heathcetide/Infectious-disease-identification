<script>
import { reportSubmitService } from '../../services/experiment'
import { getLocalReportService } from '../../services/clock'
import { baseURL } from '../../utils/http'
import { tip } from '../../utils/tip'
import ReportUpload from '../../components/report-upload/report-upload.vue'

export default {
  components: { ReportUpload },
  data() {
    return {
      formData: {
        date: '',
        stoolReport: [],
        vomitReport: [],
        bloodReport: [],
      },
      reportForm: {
        isSerologicalTestDone: false,
        isVirusAntigenTestDone: false,
        isVirusNucleicAcidTestDone: false,
        isVirusCultureIsolationDone: false,
        pathogenicTestResults: '',
        labTestFiles: [],
      },
      isViewMode: false,
      submitting: false,
    }
  },
  onLoad(options) {
    const { date } = options || {}
    if (date) {
      this.isViewMode = true
      this.getDiagnosisData(date)
    }
  },
  methods: {
    getDiagnosisData(date) {
      getLocalReportService(date)
        .then((res) => {
          this.reportForm = { ...this.reportForm, ...(res.data || {}) }
          const files = (res.data && res.data.labTestFiles) || []
          files.forEach((file) => {
            const field = file.fileId
            const newFile = {
              name: `${file.specimenType}_report_${field}.pdf`,
              url: `${baseURL}/api/v1/infect-files/lab/${field}`,
              path: `${baseURL}/api/v1/infect-files/lab/${field}`,
            }
            if (file.specimenType === '粪便') this.formData.stoolReport.push(newFile)
            else if (file.specimenType === '呕吐物') this.formData.vomitReport.push(newFile)
            else if (file.specimenType === '血液' || file.specimenType === '血') {
              this.formData.bloodReport.push(newFile)
            }
          })
        })
        .catch(() => tip.fail('加载失败'))
    },
    onDateChange(e) {
      this.formData.date = e.detail.value
    },
    handleReport(propName) {
      if (this.isViewMode) return
      this.reportForm[propName] = !this.reportForm[propName]
    },
    uploadFile(filePath, reportType, index) {
      return new Promise((resolve, reject) => {
        uni.uploadFile({
          url: '/api/v1/lab/file',
          filePath,
          name: 'file',
          formData: {
            labTestType: reportType,
            index: String(index),
          },
          success: (res) => {
            try {
              const body = JSON.parse(res.data)
              const fileId = body.data?.file?.fileId ?? body.data
              if (fileId != null) this.reportForm.labTestFiles.push(fileId)
              resolve(res)
            } catch (err) {
              reject(err)
            }
          },
          fail: reject,
        })
      })
    },
    handleSubmit() {
      if (this.isViewMode || this.submitting) return
      uni.showModal({
        content: '提交后不可修改，是否确认提交？',
        cancelText: '取消',
        confirmText: '确认提交',
        success: async (res) => {
          if (!res.confirm) return
          this.submitting = true
          tip.loading('上传中')
          this.reportForm.labTestFiles = []
          const uploadPromises = []
          this.formData.stoolReport.forEach((file, index) => {
            if (file.path && !String(file.path).startsWith('http')) {
              uploadPromises.push(this.uploadFile(file.path, '粪便', index))
            }
          })
          this.formData.vomitReport.forEach((file, index) => {
            if (file.path && !String(file.path).startsWith('http')) {
              uploadPromises.push(this.uploadFile(file.path, '呕吐物', index))
            }
          })
          this.formData.bloodReport.forEach((file, index) => {
            if (file.path && !String(file.path).startsWith('http')) {
              uploadPromises.push(this.uploadFile(file.path, '血', index))
            }
          })
          try {
            await Promise.all(uploadPromises)
            const response = await reportSubmitService(this.reportForm)
            tip.hide()
            if (response && (response.code == 200 || response.code == 1)) {
              tip.success('提交成功')
            } else {
              tip.fail((response && response.msg) || '提交失败')
            }
          } catch (error) {
            tip.hide()
            tip.fail('提交失败')
          } finally {
            this.submitting = false
          }
        },
      })
    },
  },
}
</script>

<template>
  <view class="page">
    <image class="bg" src="/static/other-background.png" mode="aspectFill" />
    <scroll-view scroll-y class="scroll" :show-scrollbar="false">
      <view class="card">
        <view class="card-title">病原学及血清学检查</view>
        <text class="card-sub">可多选</text>
        <view class="chip-grid">
          <view
            class="chip"
            :class="{ on: reportForm.isVirusAntigenTestDone }"
            @tap="handleReport('isVirusAntigenTestDone')"
          >病毒抗原检测</view>
          <view
            class="chip"
            :class="{ on: reportForm.isVirusNucleicAcidTestDone }"
            @tap="handleReport('isVirusNucleicAcidTestDone')"
          >病毒核酸检测</view>
          <view
            class="chip"
            :class="{ on: reportForm.isVirusCultureIsolationDone }"
            @tap="handleReport('isVirusCultureIsolationDone')"
          >病毒培养分离</view>
          <view
            class="chip"
            :class="{ on: reportForm.isSerologicalTestDone }"
            @tap="handleReport('isSerologicalTestDone')"
          >血清学检测</view>
        </view>
      </view>

      <view class="card" v-if="!isViewMode">
        <view class="card-title">提交时间</view>
        <picker mode="date" :value="formData.date || '2024-01-01'" start="2024-01-01" end="2099-12-31" @change="onDateChange">
          <view class="picker">{{ formData.date || '请选择日期' }}</view>
        </picker>
      </view>

      <view class="card">
        <view class="card-title">检查结果</view>
        <input
          class="input"
          type="text"
          :disabled="isViewMode"
          placeholder="请输入病原学及血清学检查结果"
          v-model="reportForm.pathogenicTestResults"
        />
      </view>

      <view class="card">
        <view class="card-title">粪便检查报告</view>
        <report-upload
          :files="formData.stoolReport"
          :disabled="isViewMode"
          @update:files="formData.stoolReport = $event"
        />
      </view>

      <view class="card">
        <view class="card-title">呕吐物检查报告</view>
        <report-upload
          :files="formData.vomitReport"
          :disabled="isViewMode"
          @update:files="formData.vomitReport = $event"
        />
      </view>

      <view class="card">
        <view class="card-title">血液检查报告</view>
        <report-upload
          :files="formData.bloodReport"
          :disabled="isViewMode"
          @update:files="formData.bloodReport = $event"
        />
      </view>

      <button
        v-if="!isViewMode"
        class="submit-btn"
        :disabled="submitting"
        @tap="handleSubmit"
      >提交</button>
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
  height: 36vh;
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
  padding-left: 8px;
  border-left: 3px solid #285ac8;
  line-height: 1.2;
  margin-bottom: 6px;
}
.card-sub {
  display: block;
  font-size: 12px;
  color: #94a3b8;
  margin-bottom: 12px;
  padding-left: 11px;
}
.chip-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}
.chip {
  padding: 8px 12px;
  border-radius: 999px;
  background: #f3f6fb;
  color: #4a5568;
  font-size: 13px;
  border: 1px solid transparent;
}
.chip.on {
  background: rgba(40, 90, 200, 0.1);
  color: #285ac8;
  border-color: rgba(40, 90, 200, 0.35);
  font-weight: 600;
}
.picker {
  height: 42px;
  line-height: 42px;
  padding: 0 12px;
  border-radius: 10px;
  background: #f5f7fb;
  font-size: 15px;
  color: #1a202c;
}
.input {
  height: 42px;
  padding: 0 12px;
  border-radius: 10px;
  background: #f5f7fb;
  font-size: 15px;
  color: #1a202c;
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
