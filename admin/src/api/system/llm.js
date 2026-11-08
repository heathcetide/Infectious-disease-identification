import axiosInstance from '@/utils/axiosInstance.js'

export const getLLMInfo = () => axiosInstance.get('/api/v1/admin/llm/info')

export const postLLMChat = (data) => axiosInstance.post('/api/v1/admin/llm/chat', data)
