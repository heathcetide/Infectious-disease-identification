import axiosInstance from "@/utils/axiosInstance.js";

export const addSatisfactionSurvey = async (data) => {
  const response = await axiosInstance.post(
    "/api/v1/admin/satisfaction/surveys",
    data
  );
  return response;
};

export const closeSatisfaction = async (surveyID) => {
  const response = await axiosInstance.put(
    `/api/v1/admin/satisfaction/surveys/${surveyID}/close`
  );
  return response;
};

export const openSatisfaction = async (surveyID) => {
  const response = await axiosInstance.put(
    `/api/v1/admin/satisfaction/surveys/${surveyID}/open`
  );
  return response;
};

export const getSatisfactionNumberInfo = async (surveyID) => {
  const response = await axiosInstance.get(
    `/api/v1/admin/satisfaction/surveys/${surveyID}/stats`
  );
  return response;
};

export const pageSelectSatisfactionSurvey = async (params) => {
  const response = await axiosInstance.get(
    "/api/v1/admin/satisfaction/surveys",
    { params }
  );
  return response;
};

export const pageSelectSatisfactionRating = async (params) => {
  const response = await axiosInstance.get(
    "/api/v1/admin/satisfaction/ratings",
    { params }
  );
  return response;
};

export const updateSatisfaction = async (data) => {
  const surveyId = data.surveyID ?? data.id;
  const response = await axiosInstance.put(
    `/api/v1/admin/satisfaction/surveys/${surveyId}`,
    data
  );
  return response;
};
