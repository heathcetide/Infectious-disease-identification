<template>
  <div class="aside-container">
    <div class="brand">
      <img class="brand-logo" src="/logo.png" alt="logo" />
      <p class="logo-lg">{{ getTitle() }}</p>
    </div>
    <TreeMenu />
  </div>
</template>

<script setup>
import TreeMenu from './treemenu.vue';
import { useUserStore } from '@/store/userrole';

const userStore = useUserStore();

const getTitle = () => {
  const name = userStore.name || '';
  if (userStore.role === '系统管理员') return `系统管理员-${name}`;
  if (userStore.role === '专职医护人员') return `专职医护人员-${name}`;
  if (userStore.role === '疾控中心工作人员') return `疾控中心工作人员-${name}`;
  return `用户-${name}`;
};
</script>


<style lang="less" scoped>
.aside-container {
  height: 100%;
  width: 180px;
  overflow-y: auto !important;
  overflow-x: hidden !important;
  border-right: none !important;
  background: var(--admin-card-bg, #f4f4f4) !important;
}

.aside-container::-webkit-scrollbar {
  display: none;
}

.brand {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  height: 48px;
  padding: 0 8px;
  box-sizing: border-box;
}

.brand-logo {
  width: 28px;
  height: 28px;
  flex-shrink: 0;
  object-fit: contain;
}

.logo-lg {
  margin: 0;
  font-style: normal;
  font-size: 12px;
  text-align: left;
  line-height: 1.3;
  color: var(--el-text-color-primary, #333);
  font-weight: 600;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  max-width: 128px;
}
</style>
