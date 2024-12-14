<template>
	<view>
		<view class="container u-skeleton">
			<view class="userinfo">
				<block>
					<!--u-skeleton-circle 绘制圆形-->
					<image class="userinfo-avatar u-skeleton-circle" :src="userInfo.avatarUrl"></image>
					<!--u-skeleton-fillet 绘制圆角矩形-->
					<text class="u-skeleton-fillet">{{userInfo.nickName}}</text>
				</block>
			</view>
			<view style="margin: 20px 0; text-align: left;">
				<view v-for="(item,index) in lists" :key="index" class="lists">
					<!--u-skeleton-rect 绘制矩形-->
					<text class="u-skeleton-rect">{{index+1+":"}} {{item}}</text>
				</view>
			</view>
		</view>
		<!--引用组件-->
		<u-skeleton :loading="loading" :animation="true" bgColor="#FFF"></u-skeleton>
	</view>
</template>

<script>
	export default {
		name: "xskeleton",
		data() {
			return {
				// 模拟数据
				userInfo: {
					avatarUrl: 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fc-ssl.duitang.com%2Fuploads%2Fblog%2F202110%2F31%2F20211031215023_805c2.thumb.1000_0.png&refer=http%3A%2F%2Fc-ssl.duitang.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1702049213&t=28b34148c808e94f0b263770e7feee86',
					nickName: 'uni-App 模板'
				},
				lists: [
					'配备好: uView, 自定义TabBar, Store, 常见个人中心 的uniapp模板',
					'含有组件列表：Loading, Skeleton, Swiper, Empty, MessageInput, Upload',
					'集成腾讯地图功能，可扩展',
				],
				loading: true, // 是否显示骨架屏组件
			};
		},
		mounted() {
			// 通过延时模拟向后端请求数据，调隐藏骨架屏
			setTimeout(() => {
				this.loading = false;
			}, 2000)
		}
	}
</script>

<style lang="scss" scoped>
	.container {
		padding: 80rpx 60rpx;
	}

	.userinfo {
		display: flex;
		flex-direction: column;
		align-items: center;
		font-weight: 900;
	}

	.userinfo-avatar {
		width: 128rpx;
		height: 128rpx;
		margin: 20rpx;
		border-radius: 50%;
	}

	.lists {
		margin: 10px 0;
		color: #fff;
		background-color: #5098ff;
		padding: 10rpx;
		font-size: 36rpx;
	}
</style>