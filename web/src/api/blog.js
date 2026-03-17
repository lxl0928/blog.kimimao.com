import request from '../utils/request'

export const listBlogs = (params) => request.get('/blogs', { params })
export const getBlog = (id, incrView = true) =>
  request.get(`/blogs/${id}`, { params: { incr_view: incrView } })
export const listMyBlogs = (params) => request.get('/blogs/my', { params })
export const createBlog = (data) => request.post('/blogs', data)
export const updateBlog = (id, data) => request.put(`/blogs/${id}`, data)
export const deleteBlog = (id) => request.delete(`/blogs/${id}`)
export const likeBlog = (id) => request.post(`/blogs/${id}/like`)
export const unlikeBlog = (id) => request.delete(`/blogs/${id}/like`)
export const favoriteBlog = (id) => request.post(`/blogs/${id}/favorite`)
export const unfavoriteBlog = (id) => request.delete(`/blogs/${id}/favorite`)
