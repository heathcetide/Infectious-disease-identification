import axiosInstance from '@/utils/axiosInstance.js'

export const listEvalVersions = () => axiosInstance.get('/api/v1/admin/eval/versions')
export const createEvalVersion = (data) => axiosInstance.post('/api/v1/admin/eval/versions', data)
export const getEvalVersion = (id) => axiosInstance.get(`/api/v1/admin/eval/versions/${id}`)
export const downloadEvalVersion = (id) =>
  axiosInstance.get(`/api/v1/admin/eval/versions/${id}/download`, { responseType: 'blob' })
export const restoreEvalVersion = (id) =>
  axiosInstance.post(`/api/v1/admin/eval/versions/${id}/restore`)
export const listEvalRuns = (params) => axiosInstance.get('/api/v1/admin/eval/runs', { params })
export const runEval = (data) => axiosInstance.post('/api/v1/admin/eval/runs', data)
export const listEvalCases = (runId, params) =>
  axiosInstance.get(`/api/v1/admin/eval/runs/${runId}/cases`, { params })
export const latestEvalCompare = () => axiosInstance.get('/api/v1/cdc/eval/latest')
