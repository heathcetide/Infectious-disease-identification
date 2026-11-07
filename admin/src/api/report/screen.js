import axiosInstance from "@/utils/axiosInstance.js";

export const getDiseaseDataToday = (date) => {
  return axiosInstance.get("/api/v1/cdc/dashboard/disease-today", {
    params: { date },
  });
};

export const getCheckInDailyNumber = async (dateList) => {
  const [startDate, endDate] = Array.isArray(dateList) ? dateList : [];
  const response = await axiosInstance.post(
    "/api/v1/cdc/dashboard/checkin-daily-series",
    { startDate, endDate }
  );
  return response;
};

export const getStatisticsData = async (date) => {
  const response = await axiosInstance.get(
    "/api/v1/cdc/dashboard/statistics",
    { params: { date } }
  );
  return response;
};

export const getChickInInfoByText = async (params) => {
  const response = await axiosInstance.get(
    "/api/v1/cdc/dashboard/checkin-search",
    { params }
  );
  return response;
};

export const getImportantUserInfo = async () => {
  const response = await axiosInstance.get(
    "/api/v1/cdc/dashboard/important-users"
  );
  return response;
};

export const getUserStation = async (date) => {
  const response = await axiosInstance.get(
    "/api/v1/cdc/dashboard/user-stations",
    { params: { date } }
  );
  return response;
};

export const getDiseaseStatisticsListInfo = async (
  dateBegin,
  dateEnd,
  diseaseList,
  sourceType
) => {
  const response = await axiosInstance.get(
    "/api/v1/cdc/disease/statistics",
    {
      params: {
        dateBegin,
        dateEnd,
        diseaseList: diseaseList.join(","),
        sourceType,
      },
    }
  );
  return response;
};

export const getstrike = async (params) => {
  const response = await axiosInstance.get("/api/v1/cdc/accuracy/strike", {
    params,
  });
  return response;
};

export const JudgmentsPredictorNumber = async (range) => {
  const response = await axiosInstance.put(
    `/api/v1/cdc/accuracy/predictor-range/${range}`
  );
  return response;
};
