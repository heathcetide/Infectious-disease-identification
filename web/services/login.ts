
import { http } from '@/utils/http'
import { normalizeEmployeeProfilePayload } from '@/utils/form'

type LoginResult = {
  accessToken: string
  refreshToken: string
  expiresIn: number
  tokenType: string
  user: {
    id: number
    name: string
    phone: string
    role: string
  }
  informationComplete: boolean
  userName?: string
  userType?: string
}
type LoginParams = {
  phoneNumber: string
  password: string
}
/**
 * 传统登录-手机号+密码
 * @param data 请求参数
 */
export const postLoginService = (data: LoginParams) => {
  return http<LoginResult>({
    method: 'POST',
    url: '/api/v1/auth/login',
    data,
    quiet: true,
  })
}

type rePasswordParams = {
	newPassword: string
	oldPassword: string
	confirmPassword: string
}

export const rePasswordService = (data: rePasswordParams) => {
	return http({
		method: 'PUT',
		url: '/api/v1/employee/password',
		data,
		quiet: true,
	})
}

type feedBackParams ={
	feedbackDate: string
	feedbackText: string
	feedbackTitle: string
	name: string
	phoneNumber: string
}

export const feedBackService = (data: feedBackParams) =>{
	return http({
		method: 'POST',
		url: '/api/v1/employee/feedback',
		data,
	})
}
type submitUserInfoParams = {
	age: number,
	birthDate: string,
	department: string,
	drinkingStatus: string,
	educationLevel: string,
	emergencyContactName: string,
	emergencyContactPhoneNumber: number,
	emergencyContactRelation: string,
	ethnicity: string,
	gender: string,
	hasAsthma: boolean,
	hasCOPD: boolean,
	hasChronicKidneyDisease: boolean,
	hasChronicLiverDisease: boolean,
	hasCoronaryHeartDisease: boolean,
	hasDiabetes: boolean,
	hasDustExposure: boolean,
	hasHepatitis: boolean,
	hasHyperlipidemia: boolean,
	hasHypertension: boolean,
	hasHyperuricemia: boolean,
	hasImmunodeficiency: boolean,
	hasLungCancer: boolean,
	hasMalignantTumor: boolean,
	hasMedicalHistory: boolean,
	hasOccupationalDisease: boolean,
	hasOtherCancer: boolean,
	hasOtherCardiovascularDiseases: boolean,
	hasOtherDiseases: boolean,
	hasPepticUlcer: boolean,
	hasSevereMentalDisorders: boolean,
	hasStroke: boolean,
	hasTuberculosis: boolean,
	hasTyphus: boolean,
	height: number,
	idNumber: number,
	medicalPersonnelType: string,
	name: string,
	otherCancerName: string,
	otherDiseasesName: string,
	otherPhoneNumber: number,
	otherPositionName: string,
	phoneNumber: number,
	postpartumInSixWeeks: boolean,
	pregnancyWeeks: number,
	pregnant: boolean,
	smokingStatus: string,
	specificOccupation: string,
	vaccinatedForBCG: boolean,
	vaccinatedForCOVID: boolean,
	vaccinatedForFlu: boolean,
	vaccinatedForHepatitis: boolean,
	vaccinatedForPlague: boolean,
	weight: number,
	workOnPlateauStartDate: string
}


export const submitUserInfoService = (data: submitUserInfoParams) =>{
	const payload = normalizeEmployeeProfilePayload(
		data as unknown as Record<string, unknown>
	)
	return http({
		method: 'PUT',
		url: '/api/v1/employee/profile',
		data: payload,
	})
}

export const getUserInfoService = () =>{
	return http({
		method: 'GET',
		url: '/api/v1/employee/profile',
	})
}
