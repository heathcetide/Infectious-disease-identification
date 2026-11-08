<template>
  <div class="page">
    <div class="title-row">
      <div>
        <h3>高风险预警记录</h3>
        <p class="sub">打卡评分 ≥ 80% 自动上报；可确认已处理。点击用户名查看联系方式等详情。</p>
      </div>
      <el-button @click="load">刷新</el-button>
    </div>
    <el-table :data="list" v-loading="loading" style="width: 100%">
      <el-table-column prop="alertId" label="ID" width="70" />
      <el-table-column label="用户" min-width="140">
        <template #default="{ row }">
          <el-button link type="primary" @click="openUser(row)">
            {{ displayName(row) }}
          </el-button>
          <div class="phone-sub" v-if="row.phoneNumber">{{ row.phoneNumber }}</div>
        </template>
      </el-table-column>
      <el-table-column prop="diseaseTypeName" label="病种" min-width="140" />
      <el-table-column prop="score" label="评分" width="80" />
      <el-table-column prop="nearestHospital" label="建议医院" min-width="160" />
      <el-table-column prop="message" label="说明" min-width="200" />
      <el-table-column label="时间" min-width="160">
        <template #default="{ row }">{{ formatDateTime(row.createdAt) }}</template>
      </el-table-column>
      <el-table-column label="状态" width="100">
        <template #default="{ row }">
          <el-tag v-if="row.acknowledged" type="success" size="small">已确认</el-tag>
          <el-tag v-else type="danger" size="small">待处理</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="100">
        <template #default="{ row }">
          <el-button v-if="!row.acknowledged" link type="primary" @click="ack(row)">确认</el-button>
        </template>
      </el-table-column>
    </el-table>

    <Checkuserdata ref="Checkuserdata" />
  </div>
</template>

<script>
import axiosInstance from '@/utils/axiosInstance.js'
import { formatDateTime } from '@/utils/date.js'
import Checkuserdata from '@/views/user/alluser/components/checkuserdata.vue'

export default {
  name: 'RiskAlerts',
  components: { Checkuserdata },
  data() {
    return { list: [], loading: false }
  },
  mounted() {
    this.load()
  },
  methods: {
    formatDateTime,
    displayName(row) {
      return row.name || row.username || String(row.userId || '未知用户')
    },
    openUser(row) {
      const id = row.userId
      if (!id) {
        this.$message.warning('缺少用户 ID')
        return
      }
      this.$refs.Checkuserdata.showDrawer(String(id))
    },
    async load() {
      this.loading = true
      try {
        const res = await axiosInstance.get('/api/v1/admin/risk-alerts', { params: { limit: 100 } })
        if (res.data?.code === 1) this.list = res.data.data || []
      } finally {
        this.loading = false
      }
    },
    async ack(row) {
      await axiosInstance.put(`/api/v1/admin/risk-alerts/${row.alertId}/ack`)
      this.$message.success('已确认')
      this.load()
    },
  },
}
</script>

<style scoped>
.page { padding: 8px 12px; }
.title-row { display: flex; justify-content: space-between; margin-bottom: 12px; }
h3 { margin: 0; color: var(--el-text-color-primary); }
.sub { margin: 4px 0 0; font-size: 12px; color: var(--el-text-color-secondary); }
.phone-sub {
  font-size: 12px;
  color: var(--el-text-color-secondary);
  line-height: 1.3;
  margin-top: 2px;
}
</style>
