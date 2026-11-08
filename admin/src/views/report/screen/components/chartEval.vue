<template>
  <div class="wrap" v-if="run">
    <div class="head" @click="expanded = !expanded">
      <div class="title">专家 vs AI（Top-K 回算）</div>
      <div class="head-right">
        <span class="brief" v-if="!expanded">
          AI@1 {{ pct(run.aiHitAt1) }} · 专家@1 {{ pct(run.expertHitAt1) }} · {{ run.caseCount || 0 }} 例
        </span>
        <span class="toggle">{{ expanded ? '收起' : '展开' }}</span>
      </div>
    </div>
    <div v-show="expanded">
      <div class="grid">
        <div class="cell">
          <div class="k">专家 Hit@1</div>
          <div class="v">{{ pct(run.expertHitAt1) }}</div>
        </div>
        <div class="cell">
          <div class="k">AI Hit@1</div>
          <div class="v ai">{{ pct(run.aiHitAt1) }}</div>
        </div>
        <div class="cell">
          <div class="k">专家 Hit@{{ run.k || 3 }}</div>
          <div class="v">{{ pct(run.expertHitAtK) }}</div>
        </div>
        <div class="cell">
          <div class="k">AI Hit@{{ run.k || 3 }}</div>
          <div class="v ai">{{ pct(run.aiHitAtK) }}</div>
        </div>
      </div>
      <div class="meta">病例 {{ run.caseCount || 0 }} · {{ run.name || '' }}</div>
    </div>
  </div>
</template>

<script>
import { latestEvalCompare } from '@/api/system/eval.js'

export default {
  name: 'ChartEval',
  data() {
    return { run: null, expanded: false }
  },
  mounted() {
    this.load()
  },
  methods: {
    pct(v) {
      if (v == null) return '—'
      return (Number(v) * 100).toFixed(1) + '%'
    },
    async load() {
      try {
        const res = await latestEvalCompare()
        if (res.data?.code === 1) this.run = res.data.data?.run || null
      } catch (e) {
        console.error(e)
      }
    },
  },
}
</script>

<style scoped>
.wrap {
  margin: 0 10px 6px;
  padding: 8px 12px;
  background: var(--admin-card-bg, #fafafa);
  border: 1px solid var(--el-border-color-lighter, transparent);
  border-radius: 8px;
  flex: 0 0 auto;
}
.head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  cursor: pointer;
  user-select: none;
}
.title {
  font-size: 12px;
  font-weight: 700;
  color: var(--el-text-color-primary, #4a4a4a);
}
.head-right {
  display: flex;
  align-items: center;
  gap: 10px;
  min-width: 0;
}
.brief {
  font-size: 11px;
  color: var(--el-text-color-secondary, #888);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.toggle {
  font-size: 11px;
  color: #285ac8;
  flex-shrink: 0;
}
.grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 8px;
  margin-top: 8px;
}
.cell {
  background: var(--el-fill-color-blank, #fff);
  border-radius: 6px;
  padding: 8px;
}
.k {
  font-size: 11px;
  color: var(--el-text-color-secondary, #888);
}
.v {
  margin-top: 4px;
  font-size: 18px;
  font-weight: 700;
  color: var(--el-text-color-primary, #333);
}
.v.ai {
  color: #285ac8;
}
.meta {
  margin-top: 6px;
  font-size: 11px;
  color: var(--el-text-color-secondary, #999);
}
</style>
