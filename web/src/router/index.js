import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '../stores/user'

const routes = [
  { path: '/', redirect: '/blogs' },
  { path: '/blogs', component: () => import('../views/user/BlogList.vue'), meta: { title: '文章列表' } },
  { path: '/blogs/:id', component: () => import('../views/user/BlogDetail.vue'), meta: { title: '文章详情' } },
  { path: '/blogs/edit/:id?', component: () => import('../views/user/BlogEdit.vue'), meta: { title: '编辑文章', auth: true } },
  { path: '/login', component: () => import('../views/user/Login.vue'), meta: { title: '登录' } },
  { path: '/register', component: () => import('../views/user/Register.vue'), meta: { title: '注册' } },
  { path: '/admin', redirect: '/admin/users' },
  { path: '/admin/users', component: () => import('../views/admin/UserList.vue'), meta: { title: '用户管理', admin: true } },
  { path: '/admin/blogs', component: () => import('../views/admin/BlogList.vue'), meta: { title: '文章管理', admin: true } },
]

const router = createRouter({ history: createWebHistory(), routes })

router.beforeEach((to, from, next) => {
  document.title = to.meta.title ? `${to.meta.title} - 个人博客` : '个人博客'
  const store = useUserStore()
  if (to.meta.admin && !store.isAdmin) {
    next('/login')
  } else if (to.meta.auth && !store.isLoggedIn) {
    next('/login')
  } else {
    next()
  }
})

export default router
