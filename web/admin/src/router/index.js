import { createRouter, createWebHashHistory } from 'vue-router'
import Login from '../views/Login.vue'
import Admin from '../views/Admin.vue'
import Particles from '../components/particles/Particles.vue'

const routes = [
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/admin',
    name: 'admin',
    component: Admin
  },
  {
    path: '/particles',
    name: 'particles',
    component: Particles
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
