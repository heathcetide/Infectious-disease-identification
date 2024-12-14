import App from './App'

// #ifndef VUE3
import Vue from 'vue'
import './uni.promisify.adaptor'
import uView from "uview-ui";
import store from './store'
import xtabbar from './components/xtabbar.vue'

Vue.use(uView);
Vue.config.productionTip = false
// 注册全局组件
Vue.component('xtabbar', xtabbar);


App.mpType = 'app'
const app = new Vue({
	store,
	...App
})
app.$mount()
// #endif

// #ifdef VUE3
import {
	createSSRApp
} from 'vue'
export function createApp() {
	const app = createSSRApp(App)
	return {
		app
	}
}
// #endif