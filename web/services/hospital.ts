import { http } from '../utils/http'

export const nearbyHospitalsService = (params: { lat: number; lng: number; limit?: number }) => {
  const q = `lat=${encodeURIComponent(String(params.lat))}&lng=${encodeURIComponent(String(params.lng))}&limit=${params.limit || 8}`
  return http({
    method: 'GET',
    url: `/api/v1/employee/hospitals/nearby?${q}`,
  })
}
