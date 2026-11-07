<template>
  <div class="container">
    <div class="chart-container">
      <div class="chart-title">重点人员跟踪</div>
      <div class="toolbar">
        <el-input
          v-model="phoneKeyword"
          placeholder="手机号筛选"
          clearable
          size="small"
          style="width: 140px"
          @clear="applyFilter"
          @keyup.enter="applyFilter"
        />
        <el-button size="small" type="primary" plain @click="applyFilter">搜索</el-button>
      </div>
      <div class="table">
        <el-table
          :data="displayedData"
          :header-cell-style="{ height: '15px', background: '#FFFFFF', color: '#333333' }"
          style="width: 100%; height: 100%"
          :show-overflow-tooltip="true"
          size="small"
        >
          <el-table-column label="" width="42">
            <template #default="scope">
              <img :src="getAvatarUrl(scope.row.gender)" alt="头像" class="avatar" />
            </template>
          </el-table-column>
          <el-table-column prop="name" label="姓名" min-width="56" />
          <el-table-column prop="department" label="部门/工种" min-width="78" />
          <el-table-column prop="statu" label="状态" min-width="72">
            <template #default="scope">
              <el-tag :type="scope.row.statu === '健康' ? 'success' : 'danger'" size="small">
                {{ scope.row.statu }}
              </el-tag>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, onBeforeUnmount, ref } from 'vue';
import maleAvatar from '@/assets/avatar/male.png';
import femaleAvatar from '@/assets/avatar/female.png';
import { getImportantUserInfo, getChickInInfoByText } from '@/api/report/screen.js';

const getAvatarUrl = (gender) => (gender === '男' ? maleAvatar : femaleAvatar);

const allRows = ref([]);
const displayedData = ref([]);
const phoneKeyword = ref('');
let rotateTimer = null;

const mapImportant = (item) => ({
  id: item.userId,
  name: item.name || '',
  department: item.department || '—',
  phoneNumber: item.phoneNumber || '',
  gender: item.gender || '男',
  statu:
    item.isHealth === true
      ? '健康'
      : item.diagnosisDiseaseTypeName || item.reason || '异常',
});

const fetchImportantUserData = async () => {
  try {
    const response = await getImportantUserInfo();
    if (response.data.code === 1) {
      allRows.value = (response.data.data || []).map(mapImportant);
      applyFilter();
    }
  } catch (error) {
    console.error('请求出错:', error);
  }
};

const applyFilter = async () => {
  const kw = (phoneKeyword.value || '').trim();
  if (!kw) {
    displayedData.value = [...allRows.value];
    return;
  }
  // Prefer server search when keyword provided
  try {
    const response = await getChickInInfoByText({
      infoNumber: 20,
      phoneNumber: kw,
      q: kw,
    });
    if (response.data.code === 1 && Array.isArray(response.data.data)) {
      displayedData.value = response.data.data.map((item) => ({
        id: item.userId,
        name: item.name || '',
        department: item.department || '—',
        phoneNumber: item.phoneNumber || '',
        gender: item.gender || '男',
        statu:
          item.isHealth === false
            ? item.diagnosisDiseaseTypeName || item.diseaseTypeName || '异常'
            : '健康',
      }));
      return;
    }
  } catch (_) {
    /* fall through to local filter */
  }
  displayedData.value = allRows.value.filter(
    (row) =>
      String(row.phoneNumber).includes(kw) ||
      String(row.name).includes(kw) ||
      String(row.department).includes(kw)
  );
};

const rotateData = () => {
  if (displayedData.value.length < 2) return;
  const first = displayedData.value.shift();
  displayedData.value.push(first);
};

onMounted(() => {
  fetchImportantUserData();
  rotateTimer = setInterval(rotateData, 3000);
});

onBeforeUnmount(() => {
  if (rotateTimer) clearInterval(rotateTimer);
});
</script>

<style scoped>
.container {
  margin-top: 8px;
  width: 100%;
  height: auto;
}

.chart-container {
  width: 100%;
  max-width: 250px;
  height: 200px;
  background-color: #fafafa;
  border-radius: 5px;
  position: relative;
  padding: 8px;
  overflow: hidden;
  box-sizing: border-box;
}

.chart-title {
  font-size: 13px;
  font-weight: bold;
  color: #4a4a4a;
  margin-bottom: 6px;
}

.toolbar {
  display: flex;
  gap: 6px;
  margin-bottom: 6px;
}

.table {
  height: calc(100% - 58px);
  overflow: auto;
}

.avatar {
  width: 22px;
  height: 22px;
  border-radius: 50%;
  object-fit: cover;
}
</style>
