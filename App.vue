<script>
	export default {
		onLaunch: function() {
			console.log('App Launch')
			uni.getLocation({
			    type: 'gcj02', // 坐标系类型
			    success: function (res) {
			        var latitude = res.latitude; // 维度
			        var longitude = res.longitude; // 经度
			        console.log('经度：' + longitude + '，纬度：' + latitude);
			    },
			    fail: function (res) {
			        console.log('获取定位失败：' + res.errMsg);
			    }
			});
		},
		onShow: function() {
			console.log('App Show')
		},
		onHide: function() {
			console.log('App Hide')
		}
	}
</script>

<style lang="scss">
	/*每个页面公共css */
	@import "uview-ui/index.scss";
</style>
