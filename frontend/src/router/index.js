import {createRouter, createWebHashHistory} from 'vue-router'
import axios from "axios";
import bus from '../plugins/eventBus'

// import store from "@/store";

const routes = [
    {
        path: '/',
        name: 'home',
        component: () => import('@/views/Home.vue')
    },
    {
        path: '/product/:id',
        name: 'CardPage',
        component: () => import('@/views/CardPage.vue') ,
    },
    {
        path: '/basket',
        name: 'Basket',
        component: () => import('@/views/Basket.vue')
    },
    {
        path: '/payment',
        name: 'PaymentPage',
        component: () => import('@/views/PaymentPage.vue'),
        meta: {
            requiresAuth: true
        }
    },
    {
        path: '/profile',
        name: 'ProfilePage',
        component: () => import('@/views/ProfilePage.vue'),
        meta: {
            requiresAuth: true
        }
    },
    {
        path: '/profile/notifications',
        name: 'Настройка уведомлений',
        component: () => import('@/views/NotificationsPage.vue'),
        meta: {
            requiresAuth: true
        }
    },
    {
        path: '/profile/editName',
        name: 'Редактировать имя',
        component: () => import('@/views/EditNamePage.vue'),
        meta: {
            requiresAuth: true
        },
    },
    {
        path: '/authPage',
        name: 'AuthPage',
        component: () => import('@/views/AuthPage.vue'),
    },
    {
        path: '/potsPayment/:id',
        name: 'PotsPaymentPage',
        component: () => import('@/views/PotsPaymentPage.vue'),
        meta: {
            requiresAuth: true
        },
        props: true
    },
    {
        path: '/admin',
        name: 'MainAdminPage',
        component: () => import('@/views/MainAdminPage.vue'),
        meta: {
            requiresAuth: false,
            admin: true
        },
    }

]

const router = createRouter({
    scrollBehavior(to,from,savedPosition) {
        if (savedPosition) {
            return savedPosition;
        } else if (to.hash) {
            return {
                el: to.hash,
                behavior: 'smooth',
            };
        } else {
            return { x: 0, y: 0 };
        }
    },
    history: createWebHashHistory(),
    routes
})

router.beforeEach( (to, from, next) => {
    if (to.matched.some(record => record.meta.requiresAuth)) {
        axios.get('/api/session/check').then(()=>{
            next()
        }).catch(()=>{
            bus.$emit('openRegistration', to)
        })
    }else{
        next()
    }
})

export default router
