import axiosInstance from "@/utils/axiosInstance.js";

export const fetchWorkenvData = async (params) => {
  const response = await axiosInstance.get("/api/v1/cdc/work-environments", {
    params,
    headers: { "Content-Type": "application/json" },
  });
  return response;
};

export const checkWorkenvDataById = async (workId) => {
  const response = await axiosInstance.get(
    `/api/v1/cdc/work-environments/${workId}`
  );
  return response;
};

export const exportWorkenvData = async () => {
  const response = await axiosInstance.get(
    "/api/v1/cdc/work-environments/export",
    {
      responseType: "blob",
    }
  );
  return response;
};
