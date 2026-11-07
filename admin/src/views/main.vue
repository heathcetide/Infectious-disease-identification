<template>
  <div class="global-header">
    <div class="header-banner" :style="{ backgroundImage: `url(${bannerUrl})` }">
      <div class="header-shade" />
      <div class="header-copy">
        <div class="header-title">突发传染病症状筛查与预警系统</div>
        <div class="header-sub">自主判别 · 现场调查 · 数据决策</div>
      </div>
    </div>
  </div>
  <div class="common-layout">
    <el-container>
      <el-aside class="aside">
        <Aside />
      </el-aside>
      <el-container class="main-container">
        <el-header class="main-header">
          <Navheader />
        </el-header>
        <el-main class="main-content">
          <router-view v-slot="{ Component }">
            <keep-alive :max="16" :exclude="['Screen']">
              <component :is="Component" />
            </keep-alive>
          </router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import Aside from '@/components/aside.vue'
import Navheader from '@/components/navheader.vue'
import bannerUrl from '@/assets/banner.png'

export default {
  components: {
    Aside,
    Navheader,
  },
  data() {
    return { bannerUrl }
  },
}
</script>

<style lang="less" scoped>
.global-header {
  position: fixed;
  top: 0;
  width: 100%;
  height: 64px;
  z-index: 10;
  background-color: var(--admin-header-bg, #0b1a2e);
}

.header-banner {
  position: relative;
  width: 100%;
  height: 100%;
  background-size: cover;
  background-position: center;
  overflow: hidden;
}

.header-shade {
  position: absolute;
  inset: 0;
  background: linear-gradient(
    90deg,
    rgba(6, 16, 32, 0.82) 0%,
    rgba(6, 16, 32, 0.45) 42%,
    rgba(6, 16, 32, 0.15) 100%
  );
}

.header-copy {
  position: relative;
  z-index: 1;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: 0 28px;
  color: #fff;
}

.header-title {
  font-size: 18px;
  font-weight: 650;
  letter-spacing: 0.04em;
  line-height: 1.2;
  text-shadow: 0 1px 8px rgba(0, 0, 0, 0.35);
}

.header-sub {
  margin-top: 2px;
  font-size: 12px;
  opacity: 0.82;
  letter-spacing: 0.12em;
}

.common-layout {
  display: flex;
  height: 100vh;
  margin-top: 64px;
}

.aside {
  position: fixed;
  width: 180px;
  top: 64px;
  bottom: 0;
  overflow: hidden;
  z-index: 10;
  background-color: var(--admin-card-bg, #f4f4f4);
  border-right: 1px solid var(--el-border-color-light, #ebeef5);
}

.main-container {
  margin-left: 180px;
  height: calc(100vh - 64px);
  display: flex;
  flex-direction: column;
}

.main-header {
  height: 56px;
  background: var(--admin-header-bg, #fff);
  border-bottom: 1px solid var(--el-border-color-light, #ebeef5);
}

.main-content {
  flex: 1;
  overflow-y: auto;
  background: var(--el-bg-color-page, #f2f3f5);
  padding: 8px 10px;
}
</style>
