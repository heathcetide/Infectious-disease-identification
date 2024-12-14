import { http } from '@/utils/http'

export const satisfyService = () => {
  return http({
    method: 'GET',
    url: '/railwayemployee/satisfaction/getIsNeedRating',
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
		url: '/railwayemployee/satisfaction/submissionSatisfaction',
		data,
	})
}