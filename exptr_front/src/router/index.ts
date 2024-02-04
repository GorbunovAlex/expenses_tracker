import { createRouter, createWebHistory } from 'vue-router'
import DefaultLayout from '@/layouts/DefaultLayout.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: DefaultLayout,
      children: [
        {
          path: '',
          name: 'Home',
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
    }
  ]
})

export default router
