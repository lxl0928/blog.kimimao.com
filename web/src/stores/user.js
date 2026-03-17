import { defineStore } from 'pinia'
import { getProfile } from '../api/user'
import { login as authLogin } from '../api/auth'

export const useUserStore = defineStore('user', {
  state: () => ({
    token: localStorage.getItem('token') || '',
    user: JSON.parse(localStorage.getItem('user') || 'null'),
  }),
  getters: {
    isLoggedIn: (s) => !!s.token,
    isAdmin: (s) => s.user?.role === 'admin',
  },
  actions: {
    async login(user, pwd) {
      const res = await authLogin({ username: user, password: pwd })
      if (res.code !== 0) throw new Error(res.message || '登录失败')
      this.token = res.data.token
      this.user = res.data.user
      localStorage.setItem('token', this.token)
      localStorage.setItem('user', JSON.stringify(this.user))
      return res
    },
    logout() {
      this.token = ''
      this.user = null
      localStorage.removeItem('token')
      localStorage.removeItem('user')
    },
    async fetchProfile() {
      const res = await getProfile()
      this.user = res.data
      localStorage.setItem('user', JSON.stringify(this.user))
      return res
    },
  },
})
