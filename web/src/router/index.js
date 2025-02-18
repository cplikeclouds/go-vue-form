import { createRouter, createWebHistory } from 'vue-router';
import FormPage from '../frontend/views/FormPages.vue'
import HomePage from '../frontend/views/Home.vue'
import Success from "../frontend/views/Success.vue";
import FormComponent from "../frontend/components/FormComponent.vue";

import AdminLayout from '../admin/layout/AdminLayout.vue';
import Dashboard from '../admin/views/Dashboard.vue';
import DataList from '../admin/views/DataList.vue';
import Analysis from '../admin/views/Analysis.vue';
import { checkAdminRole } from '../admin/check/AdminCheck.js';
import FrontendLayout from "../frontend/layout/FrontendLayout.vue";


const routes = [
    // 用户端路由
    {
        path: '/',
        component: FrontendLayout,
        children: [
            { path: '', redirect: 'home' },
            { path: 'home', component: HomePage },
            { path: 'form', component: FormPage },
            { path: 'success', component: Success,
                beforeEnter: (to, from, next) => {
                    // 检查是否从表单页面跳转
                    if (from.path === '/form') {
                        next(); // 允许访问
                    } else {
                        next('/'); // 重定向到首页
                    }
                },},
        ],
    },
    // 管理端路由
    {
        path: '/admin',
        component: AdminLayout,
        meta: { requiresAuth: true }, // 需要登录
        children: [
            { path: '', redirect: '/admin/dashboard' },
            { path: 'dashboard', component: Dashboard },
            { path: 'data-list', component: DataList },
            { path: 'analysis', component: Analysis },
        ],
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

router.beforeEach((to, from, next) => {
    const isAdmin = checkAdminRole(); // 检查用户是否是管理员
    if (to.meta.requiresAuth && !isAdmin) {
        next('/'); // 非管理员跳转到首页
    } else {
        next();
    }
});
export default router;