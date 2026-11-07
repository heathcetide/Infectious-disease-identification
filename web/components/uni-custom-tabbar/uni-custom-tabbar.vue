<script>
export default {
  data() {
    return {
      currentPath: '', // 用于存储当前路径
      navBarTitle: '',
      showDropdown: false, // 控制下拉框的显示
      dropdownItems: [] // 存储下拉框的选项
    };
  },
  mounted() {
    this.setNavBarTitle();
    this.currentPath = getCurrentPages().pop().route;
  },
  computed: {
    healthProtectionIcon() {
      const heathProtectPaths = [
        'pages/healthProtection/healthProtection',
        'pages/clock/clock',
        'pages/clockSuccess/clockSuccess',
        'pages/clockFail/clockFail',
        'pages/todaySymptoms/todaySymptoms',
      ];
      return heathProtectPaths.includes(this.currentPath)
        ? '/static/tabs/tabicon-1-healthprotect-selected.png'
        : '/static/tabs/tabicon-1-healthprotect-Notselected.png';
    },
    healthStatusIcon() {
      const healthStatusPaths = [
        'pages/clockRecord/clockRecord',
        'pages/diagnosis/diagnosis',
        'pages/experimentalInput/experimentalInput',
        'pages/workEnvironment/workEnvironment',
        'pages/environmentInfo/environmentInfo',
        'pages/registerSituation/registerSituation',
      ];
      return healthStatusPaths.includes(this.currentPath)
        ? '/static/tabs/tabicon-2-healthsolution-selected.png'
        : '/static/tabs/tabicon-1-healthsolution-Notselected.png';
    },
    myInfoIcon() {
      const myInfoPaths = [
        'pages/changePassword/changePassword',
        'pages/surveyInfo/surveyInfo',
        'pages/scoring/scoring',
        'pages/feedback/feedback',
      ];
      return myInfoPaths.includes(this.currentPath)
        ? '/static/tabs/tabicon-3-me-selected.png'
        : '/static/tabs/tabicon-3-me-Notselected.png';
    }
  },
  methods: {
    returnBack() {
      if (getCurrentPages().length > 1) {
        wx.navigateBack({
          delta: 1
        });
      }
    },
    setNavBarTitle() {
      const pages = getCurrentPages();
      const currentPage = pages[pages.length - 1]; // 获取当前页面实例
      this.currentPath = currentPage.route; // 获取当前页面的路径

      // 页面配置信息
      const pageConfigs = [
		{ path: "pages/publicityMaterial/publicityMaterial", title: "宣传资料"},
        { path: "pages/index/index", title: "首页" },
        { path: "pages/save/save", title: "省钱" },
        { path: "pages/scan/scan", title: "扫码" },
        { path: "pages/money/money", title: "钱包" },
        { path: "pages/user/user", title: "个人中心" },
        { path: "pages/workEnvironment/workEnvironment", title: "当月工作环境基本信息" },
        { path: "pages/experimentalInput/experimentalInput", title: "录入实验资料" },
        { path: "pages/healthProtection/healthProtection", title: "健康防护" },
        { path: "pages/clockSuccess/clockSuccess", title: "打卡成功" },
        { path: "pages/clockFail/clockFail", title: "症状打卡" },
        { path: "pages/diagnosis/diagnosis", title: "录入诊断信息" },
        { path: "pages/changePassword/changePassword", title: "修改密码" },
        { path: "pages/feedback/feedback", title: "意见反馈" },
        { path: "pages/surveyInfo/surveyInfo", title: "个人信息" },
        { path: "pages/clockRecord/clockRecord", title: "历次打卡" },
        { path: "pages/todaySymptoms/todaySymptoms", title: "今日症状" },
        { path: "pages/scoring/scoring", title: "用户打分" }
      ];

      // 查找当前路径的配置
      const currentConfig = pageConfigs.find(config => config.path === this.currentPath);
      if (currentConfig) {
        this.navBarTitle = currentConfig.title;
      } else {
        this.navBarTitle = ''; // 如果没有匹配，使用默认值
      }
    },
    handleTextClick(text) {
      if (text === '我的信息') {
        this.dropdownItems = ['修改密码', '个人信息', '意见反馈', '用户打分'];
        this.showDropdown = true;
      } else if (text === '健康状况') {
        this.dropdownItems = ['历次打卡', '录入诊断信息', '录入实验资料', '当月工作环境基本信息'];
        this.showDropdown = true;
      } else if (text === '健康防护') {
        const targetUrl = '/pages/healthProtection/healthProtection';
        if (targetUrl !== '/' + this.currentPath) {
          uni.navigateTo({
            url: targetUrl,
          });
        }
      }
    },
    handleOption(action) {
      let url = '';
      switch (action) {
        case '用户打分':
          url = '/pages/scoring/scoring';
          break;
        case '修改密码':
          url = '/pages/changePassword/changePassword';
          break;
        case '个人信息':
          url = '/pages/surveyInfo/surveyInfo';
          break;
        case '意见反馈':
          url = '/pages/feedback/feedback';
          break;
        case '历次打卡':
          url = '/pages/clockRecord/clockRecord';
          break;
        case '录入诊断信息':
          url = '/pages/diagnosis/diagnosis';
          break;
        case '录入实验资料':
          url = '/pages/experimentalInput/experimentalInput';
          break;
        case '当月工作环境基本信息':
          url = '/pages/workEnvironment/workEnvironment';
          break;
        default:
          break;
      }
      if (url !== '/' + this.currentPath) {
        uni.navigateTo({ url });
      }
      this.showDropdown = false; // 关闭下拉框
    },
    closeDropdown() {
      this.showDropdown = false;
    }
  }
};
</script>
<template>
  <view>
    <view class="header">
      <image v-show="navBarTitle!=''" src="../../static/arrow-left.svg" mode="widthFix" class="returnIcon" @tap="returnBack()"></image>
      <text>{{navBarTitle}}</text>
    </view>
    <view class="footer">
      <view class="footer-item" @tap="handleTextClick('健康防护')">
        <image :src="healthProtectionIcon" mode="aspectFit"></image>
        <text>健康防护</text>
      </view>
      <view class="footer-item" @tap="handleTextClick('健康状况')">
        <image :src="healthStatusIcon" mode="aspectFit"></image>
        <text>健康状况</text>
      </view>
      <view class="footer-item" @tap="handleTextClick('我的信息')">
        <image :src="myInfoIcon" mode="aspectFit"></image>
        <text>我的信息</text>
      </view>
    </view>
	<!-- 自定义下拉框 -->
	<view v-if="showDropdown" class="dropdown">
	  <view class="dropdown-overlay" @tap="closeDropdown"></view>
	  <view class="dropdown-content">
	    <view v-for="(item, index) in dropdownItems" :key="index" class="dropdown-item" @tap="handleOption(item)">
	      <button class="dropdown-button">{{ item }}</button>
	    </view>
	  </view>
	</view>
  </view>

</template>

<style scoped>
.header {
  position: absolute;
  top: 0;
  width: 100%;
  height: 9.5vh;
  text-align: center;
  padding-top: 6vh;
  font-size: 17px;
  font-weight: 500;
}
.footer {
  height: 11vh;
  display: flex;
  justify-content: space-around;
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  width: 100%;
  background-color: #ffffff;
  padding: 10px 0;
  box-shadow: 0 -2px 4px rgba(0, 0, 0, 0.1);
}

.footer-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  font-size: 14px;
  color: #000000;
}

.footer-item image {
  width: 24px;
  height: 24px;
  margin-bottom: 5px;
}
.returnIcon {
  position: absolute;
  left: 15px;
  width: 20px;
}
/* 下拉框样式 */
.dropdown {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 10;
}
.dropdown-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
}
.dropdown-content {
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  background: #ffffff;
  border-top-left-radius: 8px;
  border-top-right-radius: 8px;
}
.dropdown-item {
  padding: 10px;
  text-align: center;
}
.dropdown-button {
  width: 100%;
  padding: 10px;
  background: none;
  border: none;
  font-size: 16px;
  color: #000;
}
</style>
