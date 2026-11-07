import axiosInstance from "@/utils/axiosInstance.js";

export const fetchDiseaseData = async (params) => {
  const response = await axiosInstance.post(
    "/api/v1/admin/diseases/page",
    params,
    {
      headers: { "Content-Type": "application/json" },
    }
  );
  return response;
};

export const exportDiseaseBasicData = async () => {
  const response = await axiosInstance.get("/api/v1/admin/diseases/export", {
    responseType: "blob",
  });
  return response;
};
