<script>
export default {
  data() {
    return {
      currentPath: '',
      navBarTitle: '',
      showHeader: true,
    };
  },
  mounted() {
    this.setNavBarTitle();
    this.currentPath = getCurrentPages().pop().route;
  },
  computed: {
    headerTitle() {
      return this.navBarTitle || '';
    },
    healthProtectionIcon() {
      const paths = [
        'pages/healthProtection/healthProtection',
        'pages/clock/clock',
        'pages/clockSuccess/clockSuccess',
        'pages/clockFail/clockFail',
        'pages/todaySymptoms/todaySymptoms',
        'pages/publicityMaterial/publicityMaterial',
      ];
      return paths.includes(this.currentPath)
        ? '/static/tabs/tabicon-1-healthprotect-selected.png'
        : '/static/tabs/tabicon-1-healthprotect-Notselected.png';
    },
    healthStatusIcon() {
      const paths = [
        'pages/clockRecord/clockRecord',
        'pages/diagnosis/diagnosis',
        'pages/experimentalInput/experimentalInput',
        'pages/workEnvironment/workEnvironment',
        'pages/environmentInfo/environmentInfo',
        'pages/registerSituation/registerSituation',
      ];
      return paths.includes(this.currentPath)
        ? '/static/tabs/tabicon-2-healthsolution-selected.png'
        : '/static/tabs/tabicon-1-healthsolution-Notselected.png';
    },
    myInfoIcon() {
      const paths = [
        'pages/changePassword/changePassword',
        'pages/surveyInfo/surveyInfo',
        'pages/scoring/scoring',
        'pages/feedback/feedback',
      ];
      return paths.includes(this.currentPath)
        ? '/static/tabs/tabicon-3-me-selected.png'
        : '/static/tabs/tabicon-3-me-Notselected.png';
    },
  },
  methods: {
    returnBack() {
      if (getCurrentPages().length > 1) {
        uni.navigateBack({ delta: 1 });
      } else {
        uni.redirectTo({ url: '/pages/clock/clock' });
      }
    },
    setNavBarTitle() {
      const pages = getCurrentPages();
      const currentPage = pages[pages.length - 1];
      this.currentPath = currentPage.route;

      if (this.currentPath === 'pages/clock/clock') {
        this.navBarTitle = '';
        this.showHeader = false;
        return;
      }
      this.showHeader = true;

      const pageConfigs = [
        { path: 'pages/publicityMaterial/publicityMaterial', title: '宣传资料' },
        { path: 'pages/workEnvironment/workEnvironment', title: '当月工作环境基本信息' },
        { path: 'pages/experimentalInput/experimentalInput', title: '录入实验资料' },
        { path: 'pages/healthProtection/healthProtection', title: '健康防护' },
        { path: 'pages/clockSuccess/clockSuccess', title: '打卡成功' },
        { path: 'pages/clockFail/clockFail', title: '症状打卡' },
        { path: 'pages/diagnosis/diagnosis', title: '录入诊断信息' },
        { path: 'pages/changePassword/changePassword', title: '修改密码' },
        { path: 'pages/feedback/feedback', title: '意见反馈' },
        { path: 'pages/surveyInfo/surveyInfo', title: '个人信息' },
        { path: 'pages/clockRecord/clockRecord', title: '历次打卡' },
        { path: 'pages/todaySymptoms/todaySymptoms', title: '今日症状' },
        { path: 'pages/scoring/scoring', title: '用户打分' },
        { path: 'pages/environmentInfo/environmentInfo', title: '工作环境基本信息' },
        { path: 'pages/registerSituation/registerSituation', title: '本月情况登记' },
      ];
      const currentConfig = pageConfigs.find((c) => c.path === this.currentPath);
      this.navBarTitle = currentConfig ? currentConfig.title : '';
    },
    go(url) {
      if (url === '/' + this.currentPath) return;
      uni.navigateTo({ url });
    },
    handleTab(name) {
      if (name === '健康防护') {
        this.go('/pages/healthProtection/healthProtection');
        return;
      }
      if (name === '健康状况') {
        uni.showActionSheet({
          itemList: ['历次打卡', '录入诊断信息', '录入实验资料', '当月工作环境基本信息'],
          success: (res) => {
            const map = [
              '/pages/clockRecord/clockRecord',
              '/pages/diagnosis/diagnosis',
              '/pages/experimentalInput/experimentalInput',
              '/pages/workEnvironment/workEnvironment',
            ];
            this.go(map[res.tapIndex]);
          },
        });
        return;
      }
      if (name === '我的信息') {
        uni.showActionSheet({
          itemList: ['修改密码', '个人信息', '意见反馈', '用户打分'],
          success: (res) => {
            const map = [
              '/pages/changePassword/changePassword',
              '/pages/surveyInfo/surveyInfo',
              '/pages/feedback/feedback',
              '/pages/scoring/scoring',
            ];
            this.go(map[res.tapIndex]);
          },
        });
      }
    },
  },
};
</script>
<template>
  <view>
    <view v-if="showHeader" class="header">
      <image
        v-show="navBarTitle != ''"
        src="../../static/arrow-left.svg"
        mode="widthFix"
        class="returnIcon"
        @tap="returnBack()"
      />
      <text>{{ headerTitle }}</text>
    </view>
    <view class="footer">
      <view class="footer-item" @tap="handleTab('健康防护')">
        <image :src="healthProtectionIcon" mode="aspectFit" />
        <text>健康防护</text>
      </view>
      <view class="footer-item" @tap="handleTab('健康状况')">
        <image :src="healthStatusIcon" mode="aspectFit" />
        <text>健康状况</text>
      </view>
      <view class="footer-item" @tap="handleTab('我的信息')">
        <image :src="myInfoIcon" mode="aspectFit" />
        <text>我的信息</text>
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
  z-index: 5;
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
  z-index: 20;
  box-sizing: border-box;
}
.footer-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  font-size: 12px;
  color: #000000;
  flex: 1;
}
.footer-item image {
  width: 22px;
  height: 22px;
  margin-bottom: 4px;
}
.returnIcon {
  position: absolute;
  left: 15px;
  width: 20px;
}
</style>
