<template>
	<view>
		<view class="tabbar">
			<u-tabbar @change="changeIndex" :list="list" :mid-button="true" active-color="#5098ff"></u-tabbar>
		</view>
	</view>
</template>

<script>
	// import store from '@/store/index.js';
	export default {
		name: "xtabbar",
		data() {
			return {

			};
		},
		computed: {
			list() {
				return store.state.list
			}
		},
		methods: {
			changeIndex(e) {
				// e 导航索引
				// e 为0时为首页 current == 0

			}
		},
	}
</script>

<style>

</style>