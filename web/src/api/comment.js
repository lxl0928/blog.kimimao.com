import request from '../utils/request'

export const listComments = (blogId) => request.get(`/blogs/${blogId}/comments`)
export const createComment = (blogId, data) => request.post(`/blogs/${blogId}/comments`, data)
