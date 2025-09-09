// file: frontend/src/router/index.ts
import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import AppLayout from '../layouts/AppLayout.vue' // 导入布局组件
import { useUserStore } from '@/stores/user'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'layout', // 给布局一个名字
      component: AppLayout,
      children: [ // 子路由
        {
          path: '', // 默认子路由
          name: 'home',
          component: HomeView
        }
      ]
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    }
  ]
})

// 导航守卫不变
router.beforeEach((to, from, next) => {
  const userStore = useUserStore()
  if (!userStore.token && to.name !== 'login') {
    next({ name: 'login' })
  } else if (userStore.token && to.name === 'login') {
    next({ name: 'home' })
  } else {
    next()
  }
})

export default router