import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import { resolve } from 'path';
import AutoImport from 'unplugin-auto-import/vite';
import Components from 'unplugin-vue-components/vite';
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers';
import { amapProxyPlugin } from './vite.amap-proxy.js';

// 已全量引入 element-plus 样式，关闭按需 style 注入，避免 Vite
// “optimized dependencies changed. reloading” 导致点几个菜单就整页刷新。
const epResolver = ElementPlusResolver({ importStyle: false });

export default defineConfig({
  plugins: [
    vue(),
    amapProxyPlugin(),
    AutoImport({
      resolvers: [epResolver],
    }),
    Components({
      resolvers: [epResolver],
    }),
  ],
  resolve: {
    alias: {
      '@': resolve(__dirname, './src'),
    },
  },
  optimizeDeps: {
    include: [
      'vue',
      'vue-router',
      'pinia',
      'axios',
      'element-plus',
      'element-plus/es',
      'element-plus/es/locale/lang/zh-cn',
      '@element-plus/icons-vue',
      'echarts',
      'dayjs',
      'qs',
      'xlsx',
    ],
  },
  server: {
    warmup: {
      clientFiles: [
        './src/views/**/*.vue',
        './src/components/**/*.vue',
      ],
    },
    proxy: {
      '/admin': {
        target: 'https://ds.sccdc.cn/',
        changeOrigin: true,
      },
      '/cdcStaff': {
        target: 'https://ds.sccdc.cn/',
        changeOrigin: true,
      },
      '/medicalStaff': {
        target: 'https://ds.sccdc.cn/',
        changeOrigin: true,
      },
      '/common': {
        target: 'https://ds.sccdc.cn/',
        changeOrigin: true,
      },
      '/user': {
        target: 'https://ds.sccdc.cn/',
        changeOrigin: true,
      },
      '/ai': {
        target: 'https://ds.sccdc.cn/',
        changeOrigin: true,
      },
    },
  },
});
