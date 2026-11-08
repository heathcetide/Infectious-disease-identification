<template>
  <div class="container">
    <div class="title-container">
      <div class="blue-box"></div>
      <span class="title-text">专家 / AI 诊断模式</span>
    </div>

    <p class="desc">
      专家模式：人工标定权重求和。AI 模式：LLM 输出证据链（症状→病种理由），并与专家权重做 Top 一致性校验；分歧时职工端标黄并需人工确认。LLM 失败时降级到 AI 权重表。
    </p>

    <div class="content">
      <el-form :model="form" label-width="88px">
        <el-form-item label="诊断模式">
          <el-radio-group v-model="form.diagnosisMethod">
            <el-radio value="expert">专家诊断</el-radio>
            <el-radio value="ai">AI 诊断</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
    </div>

    <div class="status">
      当前生效：
      <span>{{ activeLabel }}</span>
    </div>

    <div class="hint">
      <template v-if="form.diagnosisMethod === 'expert'">
        专家诊断：按现有专家权重表计算患病率。
      </template>
      <template v-else>
        AI 诊断：大模型打分并给出证据理由；与专家 Top 结果比对，不一致时标黄待确认。LLM 不可用时降级到 AI 权重表。
      </template>
    </div>

    <div class="button">
      <el-button type="primary" @click="handleAction">确认切换</el-button>
    </div>
  </div>
</template>

<script>
import { ElMessage, ElMessageBox } from 'element-plus'
import { systemChange, systemChangeSearch } from '@/api/system/expertAI'

export default {
  data() {
    return {
      form: {
        diagnosisMethod: 'expert',
      },
      activeMode: 'expert',
    }
  },
  computed: {
    activeLabel() {
      return this.activeMode === 'ai' ? 'AI 诊断' : '专家诊断'
    },
  },
  methods: {
    async handleQuery() {
      const response = await systemChangeSearch()
      const payload = response?.data?.data
      const isAI =
        payload === 'AI诊断' ||
        payload === true ||
        payload?.isAI === true
      this.activeMode = isAI ? 'ai' : 'expert'
      this.form.diagnosisMethod = this.activeMode
    },
    async handleAction() {
      ElMessageBox.confirm('是否确认切换诊断模式？', '确认', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          try {
            const toAI = this.form.diagnosisMethod === 'ai'
            const response = await systemChange(toAI)
            if (response.data.code === 1 || response.data.code === 200) {
              ElMessage.success('切换成功')
              this.handleQuery()
            } else {
              ElMessage.error('切换失败')
            }
          } catch (error) {
            console.error('Error toggling status:', error)
            ElMessage.error('切换失败，请重试')
          }
        })
        .catch(() => {})
    },
  },
  mounted() {
    this.handleQuery()
  },
}
</script>

<style scoped>
.container {
  width: 640px;
  margin: 12vh auto 0;
  background-color: #ffffff;
  border-radius: 10px;
  padding: 24px 28px 28px;
  box-shadow: 0 8px 28px rgba(40, 90, 200, 0.08);
}
.desc {
  margin: 0 0 18px;
  font-size: 13px;
  line-height: 1.6;
  color: #64748b;
}
.content {
  font-size: 14px;
  margin-bottom: 8px;
}
.title-container {
  display: flex;
  align-items: center;
  margin-bottom: 14px;
}
.blue-box {
  width: 6px;
  height: 18px;
  background-color: #285ac8;
  margin-right: 10px;
}
.title-text {
  font-size: 15px;
  font-weight: 700;
  color: #1f2a44;
}
.status {
  margin-bottom: 8px;
  font-size: 14px;
  color: #334155;
}
.status span {
  color: #285ac8;
  font-weight: 700;
}
.hint {
  margin-bottom: 20px;
  font-size: 12px;
  color: #94a3b8;
  line-height: 1.5;
}
.button {
  display: flex;
  justify-content: flex-end;
}
</style>
