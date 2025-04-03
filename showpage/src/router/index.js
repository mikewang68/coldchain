import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import ItemsView from '../views/ItemsView.vue'
import GoodsView from '../views/Goods.vue'
import MortgageView from '../views/Mortgage.vue'
import OrdersView from '../views/OrdersView.vue'
import UserLogin from '../views/Login.vue'
import UserRegister from '../views/Register.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/items',
    name: 'Items',
    component: ItemsView
  },
  {
    path: '/goods',
    name: 'Goods',
    component: GoodsView
  },
  {
    path: '/mortgage',
    name: 'Mortgage',
    component: MortgageView
  },
  {
    path: '/orders',
    name: 'Orders',
    component: OrdersView
  },
  {
    path: '/login',
    name: 'Login',
    component: UserLogin
  },
  {
    path: '/register',
    name: 'Register',
    component: UserRegister
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
  const isAuthenticated = localStorage.getItem('token')

  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!isAuthenticated) {
      next({
        path: '/login',
        query: { redirect: to.fullPath }
      })
    } else {
      next()
    }
  } else {
    next()
  }
})

export default router
