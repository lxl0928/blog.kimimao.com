import request from '../utils/request'

export const listUsers = (params) => request.get('/admin/users', { params })
export const getUserStats = () => request.get('/admin/users/stats')
export const listAdminBlogs = (params) => request.get('/admin/blogs', { params })
export const getBlogStats = () => request.get('/admin/blogs/stats')
