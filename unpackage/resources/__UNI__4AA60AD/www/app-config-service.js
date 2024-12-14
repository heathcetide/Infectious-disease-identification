
var isReady=false;var onReadyCallbacks=[];
var isServiceReady=false;var onServiceReadyCallbacks=[];
var __uniConfig = {"pages":["pages/login/login","pages/workEnvironment/workEnvironment","pages/experimentalInput/experimentalInput","pages/healthProtection/healthProtection","pages/diagnosis/diagnosis","pages/changePassword/changePassword","pages/feedback/feedback","pages/surveyInfo/surveyInfo","pages/clockRecord/clockRecord","pages/clock/clock","pages/clockFail/clockFail","pages/clockSuccess/clockSuccess","pages/todaySymptoms/todaySymptoms","pages/environmentInfo/environmentInfo","pages/registerSituation/registerSituation","pages/scoring/scoring","pages/publicityMaterial/publicityMaterial","pages/webview/webview"],"window":{"navigationBarTextStyle":"black","navigationBarTitleText":"uni-app","navigationBarBackgroundColor":"#FFFFFF","navigationStyle":"custom","backgroundColor":"#F8F8F8"},"darkmode":false,"nvueCompiler":"uni-app","nvueStyleCompiler":"uni-app","renderer":"auto","splashscreen":{"alwaysShowBeforeRender":true,"autoclose":false},"appname":"高原中","compilerVersion":"4.24","entryPagePath":"pages/login/login","networkTimeout":{"request":60000,"connectSocket":60000,"uploadFile":60000,"downloadFile":60000}};
var __uniRoutes = [{"path":"/pages/login/login","meta":{"isQuit":true},"window":{"navigationBarTitleText":"登录"}},{"path":"/pages/workEnvironment/workEnvironment","meta":{},"window":{"navigationBarTitleText":"当月工作环境基本信息"}},{"path":"/pages/experimentalInput/experimentalInput","meta":{},"window":{"navigationBarTitleText":"录入实验资料"}},{"path":"/pages/healthProtection/healthProtection","meta":{},"window":{"navigationBarTitleText":"健康防护"}},{"path":"/pages/diagnosis/diagnosis","meta":{},"window":{"navigationBarTitleText":"录入诊断信息"}},{"path":"/pages/changePassword/changePassword","meta":{},"window":{"navigationBarTitleText":"修改密码"}},{"path":"/pages/feedback/feedback","meta":{},"window":{"navigationBarTitleText":"意见反馈"}},{"path":"/pages/surveyInfo/surveyInfo","meta":{},"window":{"navigationBarTitleText":"调查信息"}},{"path":"/pages/clockRecord/clockRecord","meta":{},"window":{"navigationBarTitleText":"历次打卡"}},{"path":"/pages/clock/clock","meta":{},"window":{"navigationBarTitleText":"每日打卡","enablePullDownRefresh":true}},{"path":"/pages/clockFail/clockFail","meta":{},"window":{"navigationBarTitleText":"打卡失败"}},{"path":"/pages/clockSuccess/clockSuccess","meta":{},"window":{"navigationBarTitleText":"打卡成功","enablePullDownRefresh":true}},{"path":"/pages/todaySymptoms/todaySymptoms","meta":{},"window":{"navigationBarTitleText":"今日症状"}},{"path":"/pages/environmentInfo/environmentInfo","meta":{},"window":{"navigationBarTitleText":"工作环境基本信息"}},{"path":"/pages/registerSituation/registerSituation","meta":{},"window":{"navigationBarTitleText":"本月情况登记"}},{"path":"/pages/scoring/scoring","meta":{},"window":{"navigationBarTitleText":"打分"}},{"path":"/pages/publicityMaterial/publicityMaterial","meta":{},"window":{"navigationBarTitleText":"宣传资料"}},{"path":"/pages/webview/webview","meta":{},"window":{"navigationBarTitleText":""}}];
__uniConfig.onReady=function(callback){if(__uniConfig.ready){callback()}else{onReadyCallbacks.push(callback)}};Object.defineProperty(__uniConfig,"ready",{get:function(){return isReady},set:function(val){isReady=val;if(!isReady){return}const callbacks=onReadyCallbacks.slice(0);onReadyCallbacks.length=0;callbacks.forEach(function(callback){callback()})}});
__uniConfig.onServiceReady=function(callback){if(__uniConfig.serviceReady){callback()}else{onServiceReadyCallbacks.push(callback)}};Object.defineProperty(__uniConfig,"serviceReady",{get:function(){return isServiceReady},set:function(val){isServiceReady=val;if(!isServiceReady){return}const callbacks=onServiceReadyCallbacks.slice(0);onServiceReadyCallbacks.length=0;callbacks.forEach(function(callback){callback()})}});
service.register("uni-app-config",{create(a,b,c){if(!__uniConfig.viewport){var d=b.weex.config.env.scale,e=b.weex.config.env.deviceWidth,f=Math.ceil(e/d);Object.assign(__uniConfig,{viewport:f,defaultFontSize:Math.round(f/20)})}return{instance:{__uniConfig:__uniConfig,__uniRoutes:__uniRoutes,global:void 0,window:void 0,document:void 0,frames:void 0,self:void 0,location:void 0,navigator:void 0,localStorage:void 0,history:void 0,Caches:void 0,screen:void 0,alert:void 0,confirm:void 0,prompt:void 0,fetch:void 0,XMLHttpRequest:void 0,WebSocket:void 0,webkit:void 0,print:void 0}}}});