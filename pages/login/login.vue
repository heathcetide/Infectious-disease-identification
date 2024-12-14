<script>
import store from '../../store/index.ts'
import {postLoginService} from '../../services/login';
export default{
	  data() {
	    return {
	      loginForm: {
	        phoneNumber: '',
	        password: '',
	      },
	    };
	  },
	  methods: {
	    reset() {
	      this.loginForm.phoneNumber = '';
	      this.loginForm.password = '';
	    },
		loginSuccess(response) {
		    if (response.code == 1) {
				store.dispatch('setJwt', response.data.token);
				if(response.data.informationComplete == false){
					uni.showModal({
					    title: '注意',
					    content: '检测到您未完善信息，请前往完善',
					    showCancel: false,
					    confirmText: '确定',
					    success: function (res) {
					        if (res.confirm) {
					           uni.redirectTo({
					               url: '/pages/surveyInfo/surveyInfo'
					           });
					        }
					    }
					});
				}else{
					uni.redirectTo({
					    url: '/pages/clock/clock'
					});
				}

		    } else {
		        console.log('登录失败', response);
		        uni.showModal({
		            title: '登录失败',
		            content: '登录失败，请检查您的用户名和密码。',
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
		loginError(error) {
		    console.error('登录失败', error);
			uni.showToast({ 
				title: '登录失败，请检查手机号和密码',
				icon: 'none', // 可选，默认为none，还有success, loading, warning, error等
				duration: 2000, // 显示持续时间，单位为毫秒，默认为1500
			});
		},
	    Login() {
	      console.log("username:", this.loginForm.phoneNumber);
	      if (this.loginForm.phoneNumber && this.loginForm.password) {
			// uni.redirectTo({
			//   url: '/pages/clock/clock'
			// });
	        console.log('登录成功');
			postLoginService(this.loginForm)
			    .then(res => this.loginSuccess(res))
			    .catch(error => this.loginError(error));
	      } else {
	        console.log('请输入手机号和密码');
			uni.showToast({
				title: '请输入手机号和密码',
			    icon: 'none', // 可选，默认为none，还有success, loading, warning, error等
			    duration: 2000, 
			});
	      }
	    },
	  },
}
</script>
<template>
	<view class="container">
	  <image src="30fbc998-34bd-448b-bcad-ed2b4446e5a3.png" mode="widthFix" class="background-image"></image>
	  <view class="content">

		<view class="inputs">
			<view class="clock-info">
				<view class="clock-text">
					<text class="text-global">川藏铁路突发传染病</text>
				</view>
				<text class="text-global">自主判别现场调查软件</text>
			</view>
			<view class="input-box">
				<text class="text-global">手机号</text>
				<input type="text" placeholder="请输入手机号" v-model="loginForm.phoneNumber" />
			</view>
			<view class="input-box">
				<text class="text-global">密码</text>
				<input type="password" placeholder="默认认为身份证后六位" v-model="loginForm.password" />
			</view>
		</view>
	    <button class="submit-btn" @tap="Login">提交</button>
	  </view>
	</view>
</template>
<style>
.container {
	width: 100%;
	height: 100%;
	position: relative; /* 设置为相对定位，使内部元素能够相对于容器定位 */
}
.background-image {
	width: 100%;
	z-index: -1;
}
.content {
	position: fixed;
	bottom: 0;
	left: 0;
	right: 0;
	width: 100%;
	height: 68vh;
	overflow-y: auto;
	background-color: #FAFAFA;
	z-index: 20;
}
.input-box {
	margin: 30px 35px;
}
input {
	padding: 15px;
	font-size: 18px;
	border: none;
	border-radius: 5px;
	outline: none;
	background-color: #ffffff;
	margin-top: 10px;
}
.submit-btn {
	display: flex;
	justify-content: center;
	align-items: center;
	height: 53px;
	background-color: #285AC8;
	color: white;
	font-size: 16px;
	border: none;
	border-radius: 10px;
	margin: 0 45px;
	margin-top: 60px;
	cursor: pointer;
}
.reset-password {
	display: flex;
	justify-content: center;
	align-items: center;
	height: 53px;
	margin-top: 20px;
	cursor: pointer;
	margin: 0 45px;
	margin-top: 10px;
	font-size: 16px;
	border-radius: 10px;
	background-color: #FAFAFA;
	border: none;
}
.clock-info {
	padding: 20px 0;
	background-color: #ffffff;
	width: 100%;
	font-size: 20px;
	color: #646464;
	text-align: center;
}
.clock-text {
	text-align: center;
}
.text-global{
	color: #000000;
	font-weight: 800;
}
</style>