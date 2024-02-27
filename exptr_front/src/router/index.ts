import { createRouter, createWebHistory } from 'vue-router'
import DefaultLayout from '@/layouts/DefaultLayout.vue'
import LoginLayout from '@/layouts/LoginLayout.vue'
import { getToken } from '@/helpers/funcs/auth-utils'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: DefaultLayout,
      children: [
        {
          path: '',
          name: 'Dashboard',
          component: () => import('@/views/Home.vue')
        },
        {
          path: '/analytics',
          name: 'Analytics',
          component: () => import('@/views/Analytics.vue')
        },
        {
          path: '/history',
          name: 'History',
          component: () => import('@/views/History.vue')
        },
      ]
    },
    {
      path: '/login',
      component: LoginLayout,
      children: [
        {
          path: '',
          name: 'Login',
          component: () => import('@/views/Login.vue')
        }
      ]
    }
  ]
})

router.beforeEach(async (to) => {
  const isAuthenticated = getToken()
  if (
    !isAuthenticated &&
    // ❗️ Avoid an infinite redirect
    to.name !== 'Login'
  ) {
    return { name: 'Login' }
  }
})

export default router
