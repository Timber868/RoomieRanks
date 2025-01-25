import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
import MoneyView from '@/views/MoneyView.vue'
import ProfileView from '@/views/ProfileView.vue'
import TradeView from '@/views/TradeView.vue'


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/login',
      name: 'Login',
      component: LoginView,
    },
    {
      path: '/money',
      name: 'Money',
      component: MoneyView,
    },    {
      path: '/trade',
      name: 'Trade',
      component: TradeView,
    },    {
      path: '/profile',
      name: 'Profile',
      component: ProfileView,
    },
  ],
})

export default router
