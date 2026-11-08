import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia';
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import axios from 'axios'
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import 'element-plus/theme-chalk/dark/css-vars.css';
import './styles/index.scss';
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import '@/styles/global.scss';
import '@/styles/NextContainer.css';
import { useThemeStore } from '@/store/theme'

// AMap is lazy-loaded from map views via loadAMap() — do not rewrite/preload
// on boot (login page was hitting /_AMapProxy 502 when Node DNS cannot resolve amap.com).

const app = createApp(App)

// 配置 axios
axios.defaults.baseURL = 'http://127.0.0.1:8080';

app.config.globalProperties.$http = axios;

// 路由挂载
app.use(router)
app.use(createPinia());
useThemeStore().init();

// 只需调用一次 ElementPlus
app.use(ElementPlus, {
  locale: zhCn,  
})

// 导入 element-plus icon
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

// 挂载 Vue 应用
app.mount('#app')

// 立刻去掉首屏占位，路由切换不再出现全屏 loading
const splash = document.getElementById('boot-splash')
if (splash) splash.remove()


