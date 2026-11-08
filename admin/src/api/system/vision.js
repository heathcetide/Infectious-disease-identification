import axiosInstance from '@/utils/axiosInstance.js'

export const pageVisionSamples = (params) =>
  axiosInstance.get('/api/v1/admin/vision/samples', { params })

export const visionSummary = () => axiosInstance.get('/api/v1/admin/vision/summary')

export const exportVisionManifest = () => axiosInstance.get('/api/v1/admin/vision/export')

export const promoteVisionPhoto = (data) => axiosInstance.post('/api/v1/admin/vision/promote', data)

export const createVisionSample = (formData) =>
  axiosInstance.post('/api/v1/admin/vision/samples', formData, {
    headers: { 'Content-Type': 'multipart/form-data' },
  })

export const deleteVisionSample = (id) =>
  axiosInstance.delete(`/api/v1/admin/vision/samples/${id}`)

export const trialVisionClassify = (formData) =>
  axiosInstance.post('/api/v1/admin/vision/classify', formData, {
    headers: { 'Content-Type': 'multipart/form-data' },
  })
