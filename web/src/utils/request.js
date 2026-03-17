import axios from 'axios'

const request = axios.create({
  baseURL: '/api',
  timeout: 10000,
})

// 请求拦截：附加 token
request.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// 响应拦截：统一错误处理，返回 data 层
request.interceptors.response.use(
  (res) => {
    const data = res.data
    if (data?.code !== 0 && data?.code !== undefined) {
      const err = new Error(data.message || '请求失败')
      err.code = data.code
      return Promise.reject(err)
    }
    return data
  },
  (err) => {
    if (err.response?.status === 401) {
      localStorage.removeItem('token')
      localStorage.removeItem('user')
      window.location.href = '/login'
    }
    const data = err.response?.data
    const e = new Error(data?.message || err.message || '网络错误')
    e.code = data?.code
    return Promise.reject(e)
  }
)

export default request
