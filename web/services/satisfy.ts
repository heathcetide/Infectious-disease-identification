import { http } from '@/utils/http'

export const satisfyService = () => {
  return http({
    method: 'GET',
    url: '/api/v1/employee/satisfaction/need-rating',
  })
}
type submitParams ={
	phoneNumber: number,
	ratingID: number,
	ratingTime: string,
	satisfactionLevel: string,
	surveyID: number,
	userID: number,
}

export const submitService = (data: submitParams) =>{
	return http({
		method: 'POST',
		url: '/api/v1/employee/satisfaction',
		data,
	})
}
