
import { http } from '@/utils/http'

/** 小程序登录 登录用户信息 */
type LoginResult =  {
  /** 手机号 */
  mobile: string
  /** 登录凭证 */
  token: string
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
    url: '/user/login',
    data,
  })
}

type rePasswordParams = {
	newPassword: string
	oldPassword: string
	confirmPassword: string
}

export const rePasswordService = (data: rePasswordParams) =>{
	return http<LoginResult>({
		method: 'POST',
		url: '/railwayemployee/user/password',
		data,
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
	return http<LoginResult>({
		method: 'POST',
		url: '/railwayemployee/user/submit',
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
	return http<LoginResult>({
		method: 'POST',
		url: '/railwayemployee/user/information',
		data,
	})
}

export const getUserInfoService = () =>{
	return http({
		method: 'GET',
		url: '/railwayemployee/user/getUserInfo',
	})
}
