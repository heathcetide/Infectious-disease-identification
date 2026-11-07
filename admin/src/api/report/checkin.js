import axiosInstance from "@/utils/axiosInstance.js";

export const pageSelectCheckin = async (params) => {
  const response = await axiosInstance.post(
    "/api/v1/cdc/checkin/stats/page",
    params,
    {
      headers: { "Content-Type": "application/json" },
    }
  );
  return response;
};

export const getEmployeeCheckInExcel = async () => {
  const response = await axiosInstance.get("/api/v1/cdc/checkin/stats/export", {
    responseType: "blob",
  });
  return response;
};

export const getEmployeeCheckInExcelAI = async () => {
  const response = await axiosInstance.get("/api/v1/cdc/checkin/stats/export", {
    responseType: "blob",
  });
  return response;
};

export const selectCheckinInfoById = async (dailyHealthId) => {
  const response = await axiosInstance.get(
    `/api/v1/cdc/checkin/stats/${dailyHealthId}/location`
  );
  return response;
};

export const selectCheckinById = async (dailyHealthId) => {
  const response = await axiosInstance.get(
    `/api/v1/cdc/checkin/stats/${dailyHealthId}`
  );
  return response;
};
