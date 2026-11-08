// 3.创建router实例，引入
import {createRouter,createWebHashHistory} from 'vue-router'

// 2.引入组件
// 这个main是自定义的名称，用什么都可以
import layout from '@/views/main.vue'
import login from '@/views/login/index.vue'
import { useUserStore } from '@/store/userrole'; 
// import dashboard from '@/views/dashboard/index.vue'
// import user from '@/views/user/index.vue'
// import alluser from '@/views/user/alluser/alluser.vue'
const user = () => import('@/views/user/index.vue')
const alluser = () => import('@/views/user/alluser/alluser.vue')
const profile = () => import('@/views/user/profile/profile.vue')

const workenv = () => import('@/views/work-env/index.vue')

const report = () => import('@/views/report/index.vue')
const checkin = () => import('@/views/report/checkin/checkin.vue')
const diseasereport = () => import('@/views/report/diseasereport/diseasereport.vue')
const screen = () => import('@/views/report/screen/screen.vue')

const diagnosis = () => import('@/views/diagnosis/index.vue') 
const check = () => import('@/views/check/index.vue') 

const system = () =>import('@/views/system/index.vue')
const weight = () =>import('@/views/system/weight/weight.vue')
const autoweight = () =>import('@/views/system/autoweight/autoweight.vue')
const expertAI = () =>import('@/views/system/expertAI/expertAI.vue')
const hospital = () =>import('@/views/system/hospital/hospital.vue')

const feedback = () =>import('@/views/feedback/index.vue')
const propaganda = () => import('@/views/propaganda/index.vue')
const satisfaction = () => import('@/views/satisfaction/index.vue')
const accuracy = () => import('@/views/report/accuracy/accuracy.vue')

const accuracymanage = () => import('@/views/system/accuracymanage/accuracymanage.vue')
const llmChat = () => import('@/views/system/llmChat/llmChat.vue')
const visionDataset = () => import('@/views/system/visionDataset/index.vue')
const evalLoop = () => import('@/views/system/eval/index.vue')
const riskAlerts = () => import('@/views/system/riskAlerts/index.vue')
// 1.定义路由数据
const routes=[
    {
        path: '/login',
        component: login,
        name: 'login',
    },
    {
        path: '/',
        component: layout,
        name: 'main',
        redirect: () => {
            try {
                const raw = localStorage.getItem('user')
                const user = raw ? JSON.parse(raw) : null
                if (!user?.token) return '/login'
                if (user.role === '系统管理员' || user.role === '疾控中心工作人员') return '/report/screen'
                if (user.role === '专职医护人员') return '/check'
                return '/user/profile'
            } catch {
                return '/login'
            }
        },
        children: [
            {
                path: 'dashboard',
                component: alluser,
                meta: {}
            },

            // 用户管理
            {
                path: 'user',
                component: user,
                name: 'user',
                children: [
                    {
                        path: 'alluser',
                        component: alluser,
                        name: 'alluser',
                    },
                    {
                        path: 'profile',
                        component: profile,
                        name: 'profile',
                    },
                    {
                        path: 'password',
                        redirect: { path: '/user/profile', query: { tab: 'password' } },
                        name: 'password',
                    },
                ]
            },

            // 工作环境管理
            {
                path: 'work-env',
                component: workenv,
                name: 'work-env'
            },

            // 统计报表
            {
                path: 'report',
                component: report,
                name: 'report',
                children: [
                    {
                        path: 'check-in',
                        component: checkin,
                        name: 'checkin'
                    },
                    {
                        path: 'disease-report',
                        component: diseasereport,
                        name: 'diseasereport'
                    },
                    {
                        path: 'screen',
                        component: screen,
                        name: 'screen'
                    },
                    {
                        path:'accuracy',
                        component:accuracy,
                        name:'accuracy'
                    }
                ]
            },

            // 诊断信息管理
            {
                path: 'diagnosis',
                component: diagnosis,
                name: 'diagnosis'
            },

            // 检查信息管理
            {
                path: 'check',
                component: check,
                name: 'check'
            },

            // 系统管理
            {
                path: 'system',
                component: system,
                name: 'system',
                children: [
                    {
                        path: 'dis-data',
                        redirect: '/system/weight',
                        name: 'disdata'
                    },
                    {
                        path: 'weight',
                        component: weight,
                        name: 'weight'
                    },
                    {
                        path: 'auto-weight',
                        component: autoweight,
                        name: 'autoweight'
                    },
                    {
                        path: 'expert-AI',
                        component: expertAI,
                        name: 'expertAI'
                    },
                    {
                        path: 'hospital',
                        component: hospital,
                        name: 'hospital'
                    },
                    {
                        path: 'accuracy-manage',
                        component: accuracymanage,
                        name: 'accuracymanage'
                    },
                    {
                        path: 'llm-chat',
                        component: llmChat,
                        name: 'llmChat'
                    },
                    {
                        path: 'vision-dataset',
                        component: visionDataset,
                        name: 'visionDataset'
                    },
                    {
                        path: 'eval',
                        component: evalLoop,
                        name: 'evalLoop'
                    },
                    {
                        path: 'risk-alerts',
                        component: riskAlerts,
                        name: 'riskAlerts'
                    }
                ]
            },

            // 用户反馈
            {
                path: 'feedback',
                component: feedback,
                name: 'feedback'
            },
            // 宣传材料
            {
                path: 'propaganda',
                component: propaganda,
                name: 'propaganda'
            },
            // 满意度调查
            {
                path: 'satisfaction',
                component: satisfaction,
                name: 'satisfaction'
            }
        ]
    },
]


// 4.创建router实例
const router = createRouter({
    // 把路由数据routes传递进去
    routes,
    // 路由匹配模式
    // hash模式 直接更新当前路由和页面内容，不会访问后台
    history: createWebHashHistory()
})

router.beforeEach((to, from, next) => {
    const userStore = useUserStore();
    const isLoggedIn = !!userStore.user;

    if (to.meta.requiresAuth && !isLoggedIn) {
        next('/login');
    } else {
        next();
    }
});
// 5.导出
export default router