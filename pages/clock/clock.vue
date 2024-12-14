<!-- 引入 sockjs-client 和 stompjs -->
<script src="https://cdn.jsdelivr.net/npm/sockjs-client@1/dist/sockjs.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/stompjs@2.3.3/lib/stomp.min.js"></script>
<script>	
import { clockSubmitService,clockSituationService } from '../../services/clock';
import store from '../../store/index';
export default {
  data() {
    return {
	  clockForm:{
		  longitude: 0,
		  latitude: 0,
		  city: '',
		  county: '',
		  health: true,
		  locationName: '',
		  province: '',
	  },
      location: '正在获取位置...', // 默认显示文字
      markers: [],
	  dateForm:{
		date: '',
	  },
	  clockInfo:{
		dailyhealthstatus: {
			checkInDate: "2024-10-08",
			checkInTime: {
				hour: 0,
				minute: 0,
				nano: 0,
				second: 0
			},
			city: "北京市",
			county: "东城区",
			diagnosisDiseaseTypeName: "流感",
			diseaseTypeName: "感冒",
			health: true,
			latitude: 39.916527,
			locationName: "天安门广场",
			longitude: 116.397128,
			province: "北京市",
			statusId: 0,
			userId: 0
		},
		hasSignedToday: true,
	  },
    };
  },
  mounted(){
	const now = new Date();
	const formattedDate = this.formatDate(now);
	// 设置到 feedBackForm 中
	this.dateForm.date = formattedDate;
	clockSituationService(this.dateForm)
		.then(res => this.checkSuccess(res))
		.catch(error => this.checkError(error));
  },
  created() {
    this.getLoca();
  },
  methods: {
	formatDate(date) {
	    const year = date.getFullYear();
	    const month = String(date.getMonth() + 1).padStart(2, '0'); // 月份从0开始，所以加1
	    const day = String(date.getDate()).padStart(2, '0');
	    return `${year}-${month}-${day}`;
	},
    getLoca() {
      let that = this;
      uni.getLocation({
        type: 'gcj02', // 坐标系类型
        success: function(res) {
		  that.clockForm.latitude = res.latitude;
		  that.clockForm.longitude = res.longitude
          // 逆地理编码获取地名
          that.reverseGeocode(res.latitude, res.longitude);
          // 在位置更新后手动刷新地图
          that.refreshMap();
        },
        fail: function(res) {
          console.log('获取定位失败：' + res.errMsg);
        }
      });
    },
    reverseGeocode(latitude, longitude) {
      let that = this;
      uni.request({
        url: `https://apis.map.qq.com/ws/geocoder/v1/?location=${latitude},${longitude}&key=NPCBZ-DW5C7-FKWXD-PEJJY-U7AAJ-Q2FVF&get_poi=1`,
        success: function(res) {
          if (res.data.status === 0) {
            that.location = res.data.result.address; // 获取地名信息
          } else {
            that.location = '获取地名失败';
          }
        },
        fail: function(res) {
          console.log('获取定位失败：' + res.errMsg);
          uni.showToast({
            title: '无法获取位置信息，请检查设置。',
            icon: 'none'
          });
        }
      });
    },
    refreshMap() {
      this.clockForm.latitude += 0.000001; // 通过微小的变动来触发重新渲染
	  this.clockForm.longitude += 0.000001;
    },
    refreshLocation() {
      uni.showToast({
        title: '刷新位置',
        icon: 'none'
      });
      this.getLoca();
    },
	checkSuccess(response) {
	  this.clockInfo = response.data
	  if (response.code === 1 && !response.data.hasMonthSigned ) {
		uni.showModal({
		    title: '注意',
		    content: '检测到您未填写当月工作环境基本信息，请前往填写',
		    showCancel: false,
		    confirmText: '确定',
		    success: function (res) {
		        if (res.confirm) {
		           uni.redirectTo({
		               url: '/pages/workEnvironment/workEnvironment'
		           });
		        }
		    }
		});
	  }else if (response.code === 1 && response.data.hasSignedToday) {
		// 先显示toast
		uni.showToast({
		  title: '今日已打卡',
		  icon: 'success',
		  duration: 1500 // toast显示1.5秒，与延迟跳转时间一致
		});
		if(response.data.dailyhealthstatus.isHealth == true){
			// 使用setTimeout来延迟执行跳转操作
			setTimeout(() => {
			  uni.navigateTo({
			    url: '/pages/clockSuccess/clockSuccess'
			  });
			}, 1500); // 等待1.5秒后跳转
		}else{
			// 使用setTimeout来延迟执行跳转操作
			setTimeout(() => {
			  uni.navigateTo({
			    url: '/pages/clockFail/clockFail'
			  });
			}, 1500); // 等待1.5秒后跳转
		}
	  }
	},
	checkError(response){
		console.log(response)
	},
	submitError(response) {
		console.error('密码修改失败',response);
		uni.showToast({ 
			title: '密码修改失败，请检查手机号和密码',
			icon: 'none', // 可选，默认为none，还有success, loading, warning, error等
			duration: 2000, // 显示持续时间，单位为毫秒，默认为1500
		});
	},
	submitSuccess(response) {
		console.log(response)
	    if (response.code == 1) {
	        uni.showToast({
	            title: '打卡成功',
	            icon: 'success',
	            duration: 8000,
	            complete: () => {
	                setTimeout(() => {
	                    // 跳转到指定页面
	                    uni.navigateTo({
	                        url: '/pages/clockSuccess/clockSuccess'
	                    });
	                }, 1500); // 延迟 1.5 秒后跳转
	            }
	        });
	    } else {
	        uni.showModal({
	            title: '密码修改失败',
	            content: '密码修改失败，请检查您的密码是否正确。',
	            showCancel: false,
	            confirmText: '确定',
	            success: function (res) {
	                if (res.confirm) {
	                    console.log('用户点击确定');
	                }
	            }
	        });
	    }
	},
	submitSuccess1(response) {
		console.log(response)
	    if (response.code == 1) {
	        uni.showToast({
	            title: '有症状打卡成功',
	            icon: 'success',
	            duration: 3000,
	        });
	    } else {
	        uni.showModal({
	            title: '密码修改失败',
	            content: '密码修改失败，请检查您的密码是否正确。',
	            showCancel: false,
	            confirmText: '确定',
	            success: function (res) {
	                if (res.confirm) {
	                    console.log('用户点击确定');
	                }
	            }
	        });
	    }
	},
    handleHealth() {
		const now = new Date();
		const formattedDate = this.formatDate(now);
		// 设置到 feedBackForm 中
		this.dateForm.date = formattedDate;
		clockSituationService(this.dateForm)
			.then(res => this.checkClock1(res))
			.catch(error => this.checkError(error));
    },
	handleSymptom() {
	  this.clockForm.health = false;
	  const now = new Date();
	  const formattedDate = this.formatDate(now);
	  // 设置到 feedBackForm 中
	  store.dispatch('setClockFormParams', this.clockForm);
	  console.log("666666",store.state.clockFormParams);
	  this.dateForm.date = formattedDate;
	  clockSituationService(this.dateForm)
	  	.then(res => this.checkClock2(res))
	  	.catch(error => this.checkError(error));
	},
	checkClock2(response){
		if(response.code === 1 && !response.data.hasSignedToday){
			uni.navigateTo({
				url: '/pages/todaySymptoms/todaySymptoms'
			});
		}else{
			uni.showToast({
			    title: '今日已打卡',
			    icon: 'success',
			    duration: 1500,
			    complete: () => {
			        setTimeout(() => {
			            if(response.data.dailyhealthstatus.isHealth == true){
			            	// 使用setTimeout来延迟执行跳转操作
			            	setTimeout(() => {
			            	  uni.navigateTo({
			            	    url: '/pages/clockSuccess/clockSuccess'
			            	  });
			            	}, 300); // 等待1.5秒后跳转
			            }else{
			            	// 使用setTimeout来延迟执行跳转操作
			            	setTimeout(() => {
			            	  uni.navigateTo({
			            	    url: '/pages/clockFail/clockFail'
			            	  });
			            	}, 300); // 等待1.5秒后跳转
			            }
			        }, 300); // 延迟 1.5 秒后跳转
			    }
			});
		}
	},
	checkClock1(response){
		if(response.data.dailyhealthstatus != null){
			console.log(response.data.dailyhealthstatus.isHealth,'66666666')
		}
		if(response.code === 1 && !response.data.hasSignedToday){
			this.clockForm.health = true;
			clockSubmitService(this.clockForm)
				.then(res => this.submitSuccess(res))
				.catch(error => this.submitError(error));
		}else{
			uni.showToast({
			    title: '今日已打卡',
			    icon: 'success',
			    duration: 1500,
			    complete: () => {
			        setTimeout(() => {
			            if(response.data.dailyhealthstatus.isHealth == true){
			            	// 使用setTimeout来延迟执行跳转操作
			            	setTimeout(() => {
			            	  uni.navigateTo({
			            	    url: '/pages/clockSuccess/clockSuccess'
			            	  });
			            	}, 300); // 等待1.5秒后跳转
			            }else{
			            	// 使用setTimeout来延迟执行跳转操作
			            	setTimeout(() => {
			            	  uni.navigateTo({
			            	    url: '/pages/clockFail/clockFail'
			            	  });
			            	}, 300); // 等待1.5秒后跳转
			            }
			        }, 300); // 延迟 1.5 秒后跳转
			    }
			});
		}
	},

  }
};
</script>
<template>
  <view class="container">
    <image src="../../static/other-background.png" mode="widthFix" class="background-image"></image>
    <view class="main-container">
      <view class="location-info">
        <text class="location-text">
          <image src="/static/pin-map.svg"></image>
        </text>
      </view>
      <!-- 地图 -->
      <view class="map">
        <text>区域地图</text>
        <p>您的位置：{{ location }}</p>
        <view class="map-container">
          <map style="width: 100%; height: 300px;" :latitude="clockForm.latitude" :longitude="clockForm.longitude" :show-compass="true"
            :enable-overlooking="true" :show-location="true" :show-scale="true" :min-scale="3" :max-scale="20"
            :enable-poi="true" :enable-scroll="true" :enable-rotate="true"></map>
        </view>
      </view>
      <button class="refresh-button" @tap="refreshLocation">
        <image class="refresh-icon" src="/static/arrow-repeat.svg" mode="aspectFit"></image>刷新位置
      </button>
      <view class="calendar-card">
        <text class="calendarText">请点击打卡</text>
        <image src="/static/calendar.png" mode="aspectFit" class="calendarImg"></image>
      </view>
    </view>
    <view class="options">
      <button @tap="handleHealth" class="option-button healthy">
        <image src="../../static/icon-health.png" mode="aspectFit" class="clockIcon"></image>
        <view class="text-container1"><text>健康</text></view>
      </button>
      <button @tap="handleSymptom" class="option-button symptom">
        <image src="../../static/icon-Symptomatic.png" mode="aspectFit" class="clockIcon"></image>
        <view class="text-container2"><text>有症状</text></view>
      </button>
    </view>
    <uni-custom-tabbar class="tabbar"></uni-custom-tabbar>
  </view>
</template>
<style>
.container {
	display: flex;
	flex-direction: column;
	justify-content: space-between;
	align-items: center;
	width: 100%;
	height: 100%;
	position: relative; /* 设置为相对定位，使内部元素能够相对于容器定位 */
}
	
.background-image {
	position: absolute; /* 设置为绝对定位，使其占据整个容器 */
	top: 0;
	right: 0;
	bottom: 0;
	left: 0;
	width: 100%;
	height: 100%;
	z-index: -1; /* 将图片置于内容下方 */
}
.map-container {
  width: 90%;
  margin-top: 30vh;
  }
  .location-info {
  width: 80%;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 10px;
  border-radius: 5px;
  }

  .location-text {
  font-size: 14px; /* 调整字体大小 */
  color: #666; /* 调整字体颜色为浅色 */
  text-align: center; /* 文字水平居中 */
  width: 100%; /* 使文本容器宽度为 100% */
  }

  .location-button {
  text-align: center;
  padding: 10px 20px;
  background-color: #1E90FF;
  color: #ffffff;
  border-radius: 5px;
  align-self: flex-end;
  }
.main-container{
	width: 80%;
	padding: 0;
	margin: 0 auto;
	margin-top: 10vh;
	border-radius: 100%;
}
  .map-container {
  width: 100%;
  margin: 0 auto;
  border-radius: 100%;
  }

  .refresh-button {
	  vertical-align: middle;
  padding: 5px;
  margin-top: 10px;
  font-size: 14px;
  color: #333;
  background-color: transparent;
  text-align: center;
  align-self: center;
  border: none;
  border-radius: 100%;
  }
    .refresh-icon {
		vertical-align: middle;
    width: 20px;
    height: 20px;
    margin-right: 5px;
    }
	.calendar-card{
		position: relative;
		width: 100%;
		height: 80px;
	}
	.calendarText{
  font-size: 16px;
  line-height: 43px;
  margin-left: 6px;
  position: absolute;
  bottom: 0;
	color: #285AC8;
	}
	.calendarImg{
		width: 200px;
		height: 200px;
		position: absolute;
		top: -45px;
		right: -80px;
	}
 .options {
  width: 90%;
  margin-top: 15px;
  margin-bottom: 10vh;
  }

  .option-button {
  width: 85%;
  height: 60px;
  border-radius: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  color: #ffffff;
  background-color: #405ac8;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: background-color 0.3s ease;
  margin-top: 0px;
  margin-bottom: 30px;
  }

  .option-button.healthy {
  font-weight: 700;
  background-color: white;
  color: #405ac8;
  }

  .option-button.symptom {
  font-weight: 700;
  background-color: white;
  color: orangered;
  }
  .clockIcon{
	  width: 24px;
  }
  .text-container1{
	  margin-left: 18px;
  }
  .text-container2{
	  margin-left: 12px;
  }

.tabbar{
	z-index: 999;
}
input{
	opacity: 0.7;
}
</style>