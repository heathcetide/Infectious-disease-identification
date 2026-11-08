<template>
  <div class="page">
    <div class="title-row">
      <div>
        <h3>P5 评测闭环 · 专家 vs AI</h3>
        <p class="sub">用确诊病例回算 Top-K；权重/提示词可版本化快照，支持查看、下载与恢复。</p>
      </div>
      <div class="btns">
        <el-button @click="snapshot">保存当前版本快照</el-button>
        <el-button type="primary" :loading="running" @click="run">运行 Top-K 评测</el-button>
      </div>
    </div>

    <el-form inline>
      <el-form-item label="K">
        <el-input-number v-model="form.k" :min="1" :max="10" />
      </el-form-item>
      <el-form-item label="日期">
        <el-date-picker
          v-model="form.range"
          type="daterange"
          value-format="YYYY-MM-DD"
          start-placeholder="开始"
          end-placeholder="结束"
        />
      </el-form-item>
    </el-form>

    <el-row :gutter="12" v-if="latest">
      <el-col :span="6"><el-card>病例数 <b>{{ latest.caseCount }}</b></el-card></el-col>
      <el-col :span="6"><el-card>专家 Hit@1 <b>{{ pct(latest.expertHitAt1) }}</b></el-card></el-col>
      <el-col :span="6"><el-card>AI Hit@1 <b>{{ pct(latest.aiHitAt1) }}</b></el-card></el-col>
      <el-col :span="6"><el-card>AI Hit@{{ latest.k || 3 }} <b>{{ pct(latest.aiHitAtK) }}</b></el-card></el-col>
    </el-row>

    <h4>评测记录</h4>
    <el-table :data="runs" style="width: 100%" highlight-current-row @row-click="openCases">
      <el-table-column prop="runId" label="Run" width="70" />
      <el-table-column prop="name" label="名称" min-width="140" />
      <el-table-column prop="caseCount" label="病例" width="80" />
      <el-table-column label="专家@1" width="90">
        <template #default="{ row }">{{ pct(row.expertHitAt1) }}</template>
      </el-table-column>
      <el-table-column label="专家@K" width="90">
        <template #default="{ row }">{{ pct(row.expertHitAtK) }}</template>
      </el-table-column>
      <el-table-column label="AI@1" width="90">
        <template #default="{ row }">{{ pct(row.aiHitAt1) }}</template>
      </el-table-column>
      <el-table-column label="AI@K" width="90">
        <template #default="{ row }">{{ pct(row.aiHitAtK) }}</template>
      </el-table-column>
      <el-table-column label="时间" min-width="160">
        <template #default="{ row }">{{ formatDateTime(row.createdAt) }}</template>
      </el-table-column>
    </el-table>

    <h4 v-if="activeRunId">病例明细 · Run {{ activeRunId }}</h4>
    <el-table v-if="activeRunId" :data="cases" style="width: 100%" max-height="360">
      <el-table-column prop="goldDisease" label="确诊" min-width="120" />
      <el-table-column label="专家Top" min-width="160">
        <template #default="{ row }">{{ prettyTop(row.expertTopJson) }}</template>
      </el-table-column>
      <el-table-column label="AI Top" min-width="160">
        <template #default="{ row }">{{ prettyTop(row.aiTopJson) }}</template>
      </el-table-column>
      <el-table-column label="专家@1" width="80">
        <template #default="{ row }">{{ row.expertHitAt1 ? '✓' : '—' }}</template>
      </el-table-column>
      <el-table-column label="AI@1" width="80">
        <template #default="{ row }">{{ row.aiHitAt1 ? '✓' : '—' }}</template>
      </el-table-column>
      <el-table-column label="AI@K" width="80">
        <template #default="{ row }">{{ row.aiHitAtK ? '✓' : '—' }}</template>
      </el-table-column>
    </el-table>

    <h4 style="margin-top: 20px">版本快照</h4>
    <el-table :data="versions" style="width: 100%">
      <el-table-column prop="versionId" label="ID" width="70" />
      <el-table-column prop="name" label="名称" min-width="140" />
      <el-table-column prop="mode" label="模式" width="90" />
      <el-table-column prop="promptHash" label="prompt" width="100" />
      <el-table-column prop="weightsHash" label="weights" width="100" />
      <el-table-column prop="isActive" label="当前" width="80">
        <template #default="{ row }">
          <el-tag v-if="row.isActive" type="success" size="small">active</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="notes" label="备注" min-width="120" />
      <el-table-column label="操作" width="220" fixed="right">
        <template #default="{ row }">
          <el-button link type="primary" @click="viewVersion(row)">查看</el-button>
          <el-button link type="primary" @click="downloadVersion(row)">下载</el-button>
          <el-button link type="warning" @click="restoreVersion(row)">恢复</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-drawer v-model="drawerVisible" title="快照详情" size="520px">
      <div v-if="detailLoading" class="detail-loading">加载中…</div>
      <template v-else-if="detail">
        <el-descriptions :column="1" border size="small">
          <el-descriptions-item label="名称">{{ detail.version?.name }}</el-descriptions-item>
          <el-descriptions-item label="模式">{{ detail.version?.mode }}</el-descriptions-item>
          <el-descriptions-item label="Prompt Hash">{{ detail.version?.promptHash || '—' }}</el-descriptions-item>
          <el-descriptions-item label="Weights Hash">{{ detail.version?.weightsHash || '—' }}</el-descriptions-item>
          <el-descriptions-item label="专家权重条数">{{ detail.summary?.expertWeightCount ?? 0 }}</el-descriptions-item>
          <el-descriptions-item label="AI 权重条数">{{ detail.summary?.aiWeightCount ?? 0 }}</el-descriptions-item>
          <el-descriptions-item label="可恢复">
            <el-tag :type="detail.summary?.restorable ? 'success' : 'info'" size="small">
              {{ detail.summary?.restorable ? '是' : '否（旧快照无权重）' }}
            </el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="备注">{{ detail.version?.notes || '—' }}</el-descriptions-item>
          <el-descriptions-item label="时间">{{ formatDateTime(detail.version?.createdAt) || '—' }}</el-descriptions-item>
        </el-descriptions>
        <h4 style="margin: 16px 0 8px">Prompt</h4>
        <pre class="prompt-box">{{ detail.snapshot?.promptText || detail.version?.promptText || '（空）' }}</pre>
        <div class="drawer-actions">
          <el-button @click="downloadVersion({ versionId: detail.version?.versionId })">下载 JSON</el-button>
          <el-button
            type="warning"
            :disabled="!detail.summary?.restorable"
            @click="restoreVersion({ versionId: detail.version?.versionId, name: detail.version?.name })"
          >恢复到此快照</el-button>
        </div>
      </template>
    </el-drawer>
  </div>
</template>

<script>
import {
  listEvalVersions,
  createEvalVersion,
  getEvalVersion,
  downloadEvalVersion,
  restoreEvalVersion,
  listEvalRuns,
  runEval,
  listEvalCases,
} from '@/api/system/eval.js'
import { formatDateTime } from '@/utils/date.js'

export default {
  name: 'EvalLoop',
  data() {
    return {
      form: { k: 3, range: [] },
      runs: [],
      versions: [],
      cases: [],
      activeRunId: null,
      running: false,
      latest: null,
      drawerVisible: false,
      detailLoading: false,
      detail: null,
    }
  },
  mounted() {
    this.reload()
  },
  methods: {
    formatDateTime,
    pct(v) {
      if (v == null || Number.isNaN(Number(v))) return '—'
      return (Number(v) * 100).toFixed(1) + '%'
    },
    prettyTop(raw) {
      try {
        const arr = typeof raw === 'string' ? JSON.parse(raw || '[]') : raw
        return Array.isArray(arr) ? arr.join(' / ') : '—'
      } catch (e) {
        return raw || '—'
      }
    },
    async openCases(row) {
      this.activeRunId = row.runId
      const res = await listEvalCases(row.runId, { limit: 200 })
      if (res.data?.code === 1) this.cases = res.data.data || []
    },
    async reload() {
      const [r1, r2] = await Promise.all([listEvalRuns({ limit: 30 }), listEvalVersions()])
      if (r1.data?.code === 1) {
        this.runs = r1.data.data || []
        this.latest = this.runs[0] || null
        if (this.latest?.runId) this.openCases(this.latest)
      }
      if (r2.data?.code === 1) this.versions = r2.data.data || []
    },
    async snapshot() {
      const { value } = await this.$prompt('版本名称', '保存快照', {
        inputValue: `v_${new Date().toISOString().slice(0, 16)}`,
      }).catch(() => ({ value: null }))
      if (!value) return
      await createEvalVersion({
        name: value,
        notes: 'manual snapshot',
        promptText: 'hybrid scoring + fusion priors + vision assist',
      })
      this.$message.success('已保存版本（含完整权重，可下载/恢复）')
      this.reload()
    },
    async viewVersion(row) {
      this.drawerVisible = true
      this.detailLoading = true
      this.detail = null
      try {
        const res = await getEvalVersion(row.versionId)
        if (res.data?.code === 1) this.detail = res.data.data
        else this.$message.error(res.data?.msg || '加载失败')
      } catch (e) {
        this.$message.error(e.message || '加载失败')
      } finally {
        this.detailLoading = false
      }
    },
    async downloadVersion(row) {
      const id = row?.versionId
      if (!id) return
      try {
        const res = await downloadEvalVersion(id)
        const blob = new Blob([res.data], { type: 'application/json' })
        const url = URL.createObjectURL(blob)
        const a = document.createElement('a')
        a.href = url
        a.download = `scoring_version_${id}.json`
        a.click()
        URL.revokeObjectURL(url)
        this.$message.success('已开始下载')
      } catch (e) {
        this.$message.error(e.message || '下载失败')
      }
    },
    async restoreVersion(row) {
      const id = row?.versionId
      if (!id) return
      try {
        await this.$confirm(
          `将把当前评分模式与权重表恢复为快照「${row.name || id}」，是否继续？`,
          '恢复快照',
          { type: 'warning' }
        )
      } catch {
        return
      }
      try {
        const res = await restoreEvalVersion(id)
        if (res.data?.code === 1) {
          this.$message.success('已恢复到该快照')
          this.drawerVisible = false
          this.reload()
        } else {
          this.$message.error(res.data?.msg || '恢复失败')
        }
      } catch (e) {
        this.$message.error(e.response?.data?.msg || e.message || '恢复失败')
      }
    },
    async run() {
      this.running = true
      try {
        const payload = {
          k: this.form.k,
          name: `topk_${Date.now()}`,
          dateFrom: this.form.range?.[0] || '',
          dateTo: this.form.range?.[1] || '',
        }
        const res = await runEval(payload)
        if (res.data?.code === 1 || res.status === 201) {
          this.$message.success('评测完成')
          this.reload()
        } else {
          this.$message.error(res.data?.msg || '失败')
        }
      } catch (e) {
        this.$message.error(e.message || '失败')
      } finally {
        this.running = false
      }
    },
  },
}
</script>

<style scoped>
.page { padding: 8px 12px; }
.title-row { display: flex; justify-content: space-between; gap: 12px; margin-bottom: 12px; }
h3 { margin: 0; color: var(--el-text-color-primary); }
.sub { margin: 4px 0 0; font-size: 12px; color: var(--el-text-color-secondary); }
.btns { display: flex; gap: 8px; }
h4 { margin: 16px 0 8px; color: var(--el-text-color-primary); }
.prompt-box {
  margin: 0;
  padding: 12px;
  background: var(--el-fill-color-light, #f5f7fa);
  border-radius: 8px;
  font-size: 12px;
  white-space: pre-wrap;
  word-break: break-word;
  max-height: 240px;
  overflow: auto;
}
.drawer-actions {
  margin-top: 16px;
  display: flex;
  gap: 8px;
}
.detail-loading {
  padding: 24px 0;
  color: var(--el-text-color-secondary);
}
</style>
