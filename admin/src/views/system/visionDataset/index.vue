<template>
  <div class="page">
    <div class="title-row">
      <div>
        <h3>症状视觉训练库（YOLO）</h3>
        <p class="sub">管理皮疹/溃疡等训练样本；打卡上传皮肤图时优先走 YOLO，未配置则回退多模态 LLM。</p>
      </div>
      <div class="btns">
        <el-button type="primary" plain @click="openTrial">试分类工作台</el-button>
        <el-button :loading="exporting" @click="exportManifest">导出训练 Manifest</el-button>
        <el-button type="primary" @click="showUpload = true">上传标注样本</el-button>
      </div>
    </div>

    <el-row :gutter="12" class="stats">
      <el-col :span="6"><el-card shadow="never">样本总数 <b>{{ summary.total || 0 }}</b></el-card></el-col>
      <el-col :span="18">
        <el-card shadow="never">
          <span v-for="(n, k) in summary.byClass || {}" :key="k" class="chip">{{ classLabelZh(k) }}: {{ n }}</span>
          <span v-if="!Object.keys(summary.byClass || {}).length" class="muted">暂无分类统计</span>
        </el-card>
      </el-col>
    </el-row>

    <el-form inline class="filters">
      <el-form-item label="症状类别">
        <el-select v-model="query.classLabel" clearable placeholder="全部类别" style="width: 200px" @change="load">
          <el-option v-for="c in classNames" :key="c" :label="classLabelZh(c)" :value="c" />
        </el-select>
      </el-form-item>
      <el-form-item label="数据集划分">
        <el-select v-model="query.split" clearable placeholder="全部" style="width: 140px" @change="load">
          <el-option label="训练集 train" value="train" />
          <el-option label="验证集 val" value="val" />
          <el-option label="测试集 test" value="test" />
        </el-select>
      </el-form-item>
      <el-button @click="load">刷新</el-button>
    </el-form>

    <el-table :data="list" v-loading="loading" style="width: 100%" class="vision-table">
      <el-table-column prop="sampleId" label="ID" width="70" />
      <el-table-column label="预览" width="88">
        <template #default="{ row }">
          <button
            v-if="thumbUrl(row)"
            type="button"
            class="thumb-btn"
            @click="openPreview(row)"
          >
            <img :src="thumbUrl(row)" alt="" class="thumb" />
          </button>
          <span v-else class="muted">—</span>
        </template>
      </el-table-column>
      <el-table-column label="症状类别" width="160">
        <template #default="{ row }">{{ classLabelZh(row.classLabel) }}</template>
      </el-table-column>
      <el-table-column label="症状字段" width="180">
        <template #default="{ row }">{{ symptomFieldOf(row) }}</template>
      </el-table-column>
      <el-table-column label="划分" width="100">
        <template #default="{ row }">{{ splitZh(row.split) }}</template>
      </el-table-column>
      <el-table-column prop="source" label="来源" width="90" />
      <el-table-column prop="fileName" label="文件" min-width="140" show-overflow-tooltip />
      <el-table-column prop="notes" label="备注" min-width="100" show-overflow-tooltip />
      <el-table-column label="操作" width="100">
        <template #default="{ row }">
          <el-button link type="danger" @click="remove(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog
      v-model="previewVisible"
      :title="previewTitle"
      width="min(860px, 92vw)"
      append-to-body
      destroy-on-close
      class="vision-preview-dialog"
      @closed="onPreviewClosed"
    >
      <div class="preview-wrap">
        <img v-if="previewSrc" :src="previewSrc" alt="preview" class="preview-img" />
      </div>
    </el-dialog>

    <el-pagination
      v-model:current-page="query.page"
      v-model:page-size="query.pageSize"
      :total="total"
      layout="total, prev, pager, next"
      style="margin-top: 12px"
      @current-change="load"
    />

    <el-dialog v-model="showUpload" title="上传 YOLO 训练样本" width="480px">
      <el-form label-width="90px">
        <el-form-item label="类别" required>
          <el-select v-model="form.classLabel" style="width: 100%">
            <el-option v-for="c in classNames" :key="c" :label="classLabelZh(c)" :value="c" />
          </el-select>
        </el-form-item>
        <el-form-item label="划分">
          <el-select v-model="form.split" style="width: 100%">
            <el-option label="训练集 train" value="train" />
            <el-option label="验证集 val" value="val" />
            <el-option label="测试集 test" value="test" />
          </el-select>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="form.notes" />
        </el-form-item>
        <el-form-item label="图片" required>
          <el-upload :auto-upload="false" :limit="1" accept="image/*" :on-change="onPick">
            <el-button>选择图片</el-button>
          </el-upload>
          <div v-if="form.fileName" class="muted">{{ form.fileName }}</div>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showUpload = false">取消</el-button>
        <el-button type="primary" :loading="uploading" @click="submitUpload">上传</el-button>
      </template>
    </el-dialog>

    <!-- YOLO 试分类工作台 -->
    <el-drawer
      v-model="trialOpen"
      title="YOLO 试分类工作台"
      size="min(560px, 96vw)"
      append-to-body
      destroy-on-close
      @closed="onTrialClosed"
    >
      <div class="trial-body">
        <p class="trial-help">
          上传一张皮肤/皮疹照片，查看当前模型（YOLO 优先，否则 LLM vision）给出的 Top 类别与置信度。
          置信度表示「模型认为像该类」的概率，总和约 100%。
        </p>

        <div
          class="dropzone"
          :class="{ active: trialDragging, has: !!trialPreview }"
          @dragenter.prevent="trialDragging = true"
          @dragover.prevent="trialDragging = true"
          @dragleave.prevent="trialDragging = false"
          @drop.prevent="onTrialDrop"
          @click="pickTrial"
        >
          <img v-if="trialPreview" :src="trialPreview" alt="" class="trial-img" />
          <div v-else class="drop-placeholder">
            <strong>点击或拖入图片</strong>
            <span>JPG / PNG / WebP</span>
          </div>
        </div>
        <input ref="trialInput" type="file" accept="image/*" style="display: none" @change="onTrialFile" />

        <div class="trial-actions">
          <el-button @click="pickTrial">重新选择</el-button>
          <el-button type="primary" :loading="trying" :disabled="!trialFile" @click="runTrial">
            开始分类
          </el-button>
        </div>

        <div v-if="trialError" class="trial-error">{{ trialError }}</div>

        <template v-if="trialLabels.length">
          <div class="trial-meta">
            <el-tag size="small" :type="trialClassifier === 'yolo' ? 'success' : 'warning'">
              {{ trialClassifier === 'yolo' ? 'YOLO 本地模型' : trialClassifier || '未知引擎' }}
            </el-tag>
            <span v-if="trialTop" class="top-line">
              最可能：{{ classLabelZh(trialTop.label) }}
              <b>{{ (trialTop.confidence * 100).toFixed(1) }}%</b>
            </span>
          </div>
          <el-alert
            v-if="trialTop && trialTop.confidence < 0.45"
            type="warning"
            :closable="false"
            show-icon
            title="置信度偏低：模型不太确定，结果仅供参考（训练数据仍偏少/类别不均）"
            style="margin-bottom: 10px"
          />

          <div class="charts">
            <div ref="pieRef" class="chart pie" />
            <div ref="barRef" class="chart bar" />
          </div>

          <el-table :data="trialLabels" size="small" style="width: 100%; margin-top: 12px">
            <el-table-column label="排名" width="60">
              <template #default="{ $index }">{{ $index + 1 }}</template>
            </el-table-column>
            <el-table-column label="类别">
              <template #default="{ row }">{{ classLabelZh(row.label) }}</template>
            </el-table-column>
            <el-table-column label="英文 key" prop="label" width="140" />
            <el-table-column label="置信度" width="100">
              <template #default="{ row }">{{ (row.confidence * 100).toFixed(1) }}%</template>
            </el-table-column>
          </el-table>
        </template>
      </div>
    </el-drawer>
  </div>
</template>

<script>
import * as echarts from 'echarts'
import {
  pageVisionSamples,
  visionSummary,
  createVisionSample,
  deleteVisionSample,
  exportVisionManifest,
  trialVisionClassify,
} from '@/api/system/vision.js'

const DEFAULT_CLASSES = [
  'painful_red_rash',
  'skin_ulcer',
  'itchy_rash',
  'petechial_rash',
  'blood_blister',
  'cyanosis',
  'normal_skin',
  'other',
]

const CLASS_ZH = {
  painful_red_rash: '疼痛性红疹',
  skin_ulcer: '皮肤溃疡',
  itchy_rash: '瘙痒性皮疹',
  petechial_rash: '淤点/淤斑疹',
  blood_blister: '血疱',
  cyanosis: '紫绀',
  normal_skin: '正常皮肤',
  other: '其他',
}

const CLASS_FIELD = {
  painful_red_rash: 'HasPainfulRedRash',
  skin_ulcer: 'HasSkinUlcer',
  itchy_rash: 'HasItchyRash',
  petechial_rash: 'HasCongestiveOrPetechialRash',
  blood_blister: 'HasBloodBlisters',
  cyanosis: 'HasCyanosis',
  normal_skin: '—',
  other: '—',
}

const SPLIT_ZH = {
  train: '训练集',
  val: '验证集',
  test: '测试集',
}

const PIE_COLORS = ['#285ac8', '#0d9488', '#c2410c', '#7c3aed', '#059669', '#b45309', '#64748b', '#db2777']

export default {
  name: 'VisionDataset',
  data() {
    return {
      list: [],
      total: 0,
      loading: false,
      summary: {},
      classNames: DEFAULT_CLASSES,
      query: { page: 1, pageSize: 20, classLabel: '', split: '' },
      showUpload: false,
      uploading: false,
      exporting: false,
      form: { classLabel: 'painful_red_rash', split: 'train', notes: '', file: null, fileName: '' },
      previewVisible: false,
      previewSrc: '',
      previewTitle: '',
      trialOpen: false,
      trying: false,
      trialDragging: false,
      trialFile: null,
      trialPreview: '',
      trialLabels: [],
      trialClassifier: '',
      trialError: '',
      pieChart: null,
      barChart: null,
      _blobCache: {},
    }
  },
  computed: {
    trialTop() {
      return this.trialLabels[0] || null
    },
  },
  beforeUnmount() {
    this.disposeCharts()
    if (this.trialPreview) URL.revokeObjectURL(this.trialPreview)
    Object.values(this._blobCache || {}).forEach((url) => {
      try { URL.revokeObjectURL(url) } catch { /* ignore */ }
    })
  },
  mounted() {
    this.refreshSummary()
    this.load()
  },
  methods: {
    classLabelZh(key) {
      return CLASS_ZH[key] || key || '—'
    },
    symptomFieldOf(row) {
      const raw = (row && row.symptomField) || ''
      if (raw) return raw
      const mapped = CLASS_FIELD[row && row.classLabel]
      return mapped || '—'
    },
    splitZh(s) {
      return SPLIT_ZH[s] || s || '—'
    },
    thumbUrl(row) {
      return row._blobUrl || row.publicUrl || ''
    },
    openPreview(row) {
      const src = this.thumbUrl(row)
      if (!src) return
      this.previewSrc = src
      this.previewTitle = `${this.classLabelZh(row.classLabel)} · ${row.fileName || row.sampleId || ''}`
      this.previewVisible = true
    },
    onPreviewClosed() {
      this.previewSrc = ''
      this.previewTitle = ''
    },
    async hydrateThumbs(rows) {
      const axiosInstance = (await import('@/utils/axiosInstance.js')).default
      await Promise.all(
        (rows || []).map(async (row) => {
          if (!row.previewUrl || row._blobUrl) return
          if (this._blobCache[row.previewUrl]) {
            row._blobUrl = this._blobCache[row.previewUrl]
            return
          }
          try {
            const res = await axiosInstance.get(row.previewUrl, { responseType: 'blob' })
            const url = URL.createObjectURL(res.data)
            this._blobCache[row.previewUrl] = url
            row._blobUrl = url
          } catch {
            /* ignore missing local file */
          }
        }),
      )
    },
    async refreshSummary() {
      try {
        const res = await visionSummary()
        if (res.data?.code === 1) {
          this.summary = res.data.data || {}
          if (Array.isArray(this.summary.classNames) && this.summary.classNames.length) {
            this.classNames = this.summary.classNames
          }
        }
      } catch (e) {
        console.error(e)
      }
    },
    async load() {
      this.loading = true
      try {
        const res = await pageVisionSamples(this.query)
        if (res.data?.code === 1) {
          this.list = res.data.data?.list || []
          this.total = res.data.data?.total || 0
          await this.hydrateThumbs(this.list)
        }
      } finally {
        this.loading = false
      }
    },
    onPick(file) {
      this.form.file = file.raw
      this.form.fileName = file.name
    },
    async submitUpload() {
      if (!this.form.file || !this.form.classLabel) {
        this.$message.warning('请选择类别和图片')
        return
      }
      const fd = new FormData()
      fd.append('file', this.form.file)
      fd.append('classLabel', this.form.classLabel)
      fd.append('split', this.form.split)
      fd.append('notes', this.form.notes || '')
      fd.append('source', 'manual')
      this.uploading = true
      try {
        const res = await createVisionSample(fd)
        if (res.data?.code === 1 || res.data?.code === 201 || res.status === 201) {
          this.$message.success('已入库')
          this.showUpload = false
          this.form.file = null
          this.form.fileName = ''
          this.refreshSummary()
          this.load()
        } else {
          this.$message.error(res.data?.msg || '上传失败')
        }
      } catch (e) {
        this.$message.error(e.message || '上传失败')
      } finally {
        this.uploading = false
      }
    },
    async remove(row) {
      await this.$confirm('确认软删除该样本？', '提示')
      await deleteVisionSample(row.sampleId)
      this.$message.success('已删除')
      this.refreshSummary()
      this.load()
    },
    openTrial() {
      this.trialOpen = true
    },
    onTrialClosed() {
      this.disposeCharts()
      this.trialLabels = []
      this.trialClassifier = ''
      this.trialError = ''
      this.trialDragging = false
    },
    pickTrial() {
      this.$refs.trialInput?.click()
    },
    setTrialFile(file) {
      if (!file || !file.type?.startsWith('image/')) {
        this.$message.warning('请选择图片文件')
        return
      }
      if (this.trialPreview) URL.revokeObjectURL(this.trialPreview)
      this.trialFile = file
      this.trialPreview = URL.createObjectURL(file)
      this.trialLabels = []
      this.trialError = ''
      this.trialClassifier = ''
      this.disposeCharts()
    },
    onTrialFile(e) {
      const file = e.target.files?.[0]
      e.target.value = ''
      if (file) this.setTrialFile(file)
    },
    onTrialDrop(e) {
      this.trialDragging = false
      const file = e.dataTransfer?.files?.[0]
      if (file) this.setTrialFile(file)
    },
    async runTrial() {
      if (!this.trialFile) return
      const fd = new FormData()
      fd.append('file', this.trialFile)
      this.trying = true
      this.trialError = ''
      try {
        const res = await trialVisionClassify(fd)
        const data = res.data?.data || {}
        const labels = (data.labels || [])
          .map((l) => ({
            label: l.label,
            confidence: Math.max(0, Math.min(1, Number(l.confidence) || 0)),
          }))
          .filter((l) => l.label)
          .sort((a, b) => b.confidence - a.confidence)
        this.trialClassifier = data.classifier || (data.yoloEnv ? 'yolo' : 'llm_vision')
        this.trialLabels = labels
        if (!labels.length) {
          this.trialError = data.error || res.data?.msg || '未得到分类结果'
          this.disposeCharts()
          return
        }
        await this.$nextTick()
        this.renderCharts()
      } catch (err) {
        this.trialError = err.message || '试分类失败'
        this.trialLabels = []
        this.disposeCharts()
      } finally {
        this.trying = false
      }
    },
    disposeCharts() {
      this.pieChart?.dispose()
      this.barChart?.dispose()
      this.pieChart = null
      this.barChart = null
    },
    renderCharts() {
      const labels = this.trialLabels
      if (!labels.length || !this.$refs.pieRef || !this.$refs.barRef) return
      this.disposeCharts()
      this.pieChart = echarts.init(this.$refs.pieRef)
      this.barChart = echarts.init(this.$refs.barRef)

      // Use absolute confidence (model softmax). Do NOT use {d}% — that re-normalizes slices to 100%.
      this.pieChart.setOption({
        color: PIE_COLORS,
        tooltip: {
          trigger: 'item',
          formatter: (p) => `${p.name}<br/>置信度 ${Number(p.value).toFixed(1)}%`,
        },
        series: [
          {
            type: 'pie',
            radius: ['42%', '70%'],
            center: ['50%', '52%'],
            label: {
              formatter: (p) => `${p.name}\n${Number(p.value).toFixed(1)}%`,
              fontSize: 11,
              color: '#475569',
            },
            data: labels.map((l) => ({
              name: this.classLabelZh(l.label),
              value: Number((l.confidence * 100).toFixed(2)),
            })),
          },
        ],
      })

      this.barChart.setOption({
        color: ['#285ac8'],
        grid: { left: 88, right: 28, top: 16, bottom: 28 },
        tooltip: {
          trigger: 'axis',
          axisPointer: { type: 'shadow' },
          formatter: (items) => {
            const p = items?.[0]
            return p ? `${p.name}: ${p.value.toFixed(1)}%` : ''
          },
        },
        xAxis: {
          type: 'value',
          max: 100,
          axisLabel: { formatter: '{value}%', color: '#94a3b8' },
          splitLine: { lineStyle: { color: '#e2e8f0' } },
        },
        yAxis: {
          type: 'category',
          data: labels.map((l) => this.classLabelZh(l.label)).reverse(),
          axisLabel: { color: '#334155', fontSize: 12 },
          axisTick: { show: false },
          axisLine: { show: false },
        },
        series: [
          {
            type: 'bar',
            barWidth: 14,
            data: labels.map((l) => Number((l.confidence * 100).toFixed(2))).reverse(),
            itemStyle: { borderRadius: [0, 6, 6, 0] },
            label: {
              show: true,
              position: 'right',
              formatter: '{c}%',
              color: '#64748b',
              fontSize: 11,
            },
          },
        ],
      })
    },
    async exportManifest() {
      this.exporting = true
      try {
        const res = await exportVisionManifest()
        if (res.data?.code !== 1) {
          this.$message.error(res.data?.msg || '导出失败')
          return
        }
        const blob = new Blob([JSON.stringify(res.data.data, null, 2)], { type: 'application/json' })
        const url = URL.createObjectURL(blob)
        const a = document.createElement('a')
        a.href = url
        a.download = `vision-manifest-${Date.now()}.json`
        a.click()
        URL.revokeObjectURL(url)
        this.$message.success(`已导出 ${res.data.data?.count || 0} 条`)
      } catch (e) {
        this.$message.error(e.message || '导出失败')
      } finally {
        this.exporting = false
      }
    },
  },
}
</script>

<style scoped>
.page { padding: 8px 12px; }
.title-row { display: flex; justify-content: space-between; align-items: flex-start; margin-bottom: 12px; gap: 12px; }
.btns { display: flex; gap: 8px; flex-shrink: 0; flex-wrap: wrap; }
h3 { margin: 0; color: var(--el-text-color-primary); }
.sub { margin: 4px 0 0; color: var(--el-text-color-secondary); font-size: 12px; }
.stats { margin-bottom: 12px; }
.chip { display: inline-block; margin-right: 10px; font-size: 12px; color: var(--el-text-color-regular); }
.muted { color: var(--el-text-color-secondary); font-size: 12px; }
.filters { margin-bottom: 8px; }
.vision-table :deep(.el-table__cell) { overflow: hidden; }
.thumb-btn {
  display: block;
  padding: 0;
  margin: 0;
  border: 0;
  background: transparent;
  cursor: zoom-in;
  line-height: 0;
}
.thumb {
  width: 56px;
  height: 56px;
  border-radius: 6px;
  object-fit: cover;
  border: 1px solid var(--el-border-color-lighter, #eee);
  display: block;
}
.preview-wrap {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 280px;
  max-height: min(72vh, 720px);
  background: #111;
  border-radius: 8px;
  overflow: auto;
}
.preview-img {
  max-width: 100%;
  max-height: min(72vh, 720px);
  object-fit: contain;
  display: block;
}

.trial-body { padding: 0 4px 16px; }
.trial-help {
  margin: 0 0 14px;
  font-size: 13px;
  line-height: 1.55;
  color: #64748b;
}
.dropzone {
  border: 1.5px dashed #cbd5e1;
  border-radius: 12px;
  min-height: 180px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  background: #f8fafc;
  overflow: hidden;
  transition: border-color 0.15s, background 0.15s;
}
.dropzone.active,
.dropzone:hover {
  border-color: #285ac8;
  background: #eef3ff;
}
.dropzone.has {
  background: #0f172a;
  border-style: solid;
  border-color: #0f172a;
}
.drop-placeholder {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
  color: #64748b;
}
.drop-placeholder strong {
  font-size: 14px;
  color: #334155;
}
.drop-placeholder span {
  font-size: 12px;
}
.trial-img {
  max-width: 100%;
  max-height: 240px;
  object-fit: contain;
  display: block;
}
.trial-actions {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
  margin-top: 12px;
}
.trial-error {
  margin-top: 12px;
  padding: 10px 12px;
  border-radius: 8px;
  background: #fff1f2;
  color: #be123c;
  font-size: 13px;
}
.trial-meta {
  display: flex;
  align-items: center;
  gap: 12px;
  margin: 16px 0 8px;
  flex-wrap: wrap;
}
.top-line {
  font-size: 13px;
  color: #334155;
}
.top-line b {
  color: #285ac8;
  margin-left: 4px;
}
.charts {
  display: grid;
  grid-template-columns: 1fr;
  gap: 8px;
}
.chart {
  width: 100%;
  height: 220px;
  background: #fff;
  border: 1px solid #e2e8f0;
  border-radius: 10px;
}
.chart.bar { height: 240px; }
</style>
