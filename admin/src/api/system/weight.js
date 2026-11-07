import axiosInstance from "@/utils/axiosInstance.js";

export const getAllWeightScoringInfo = async (params) => {
  const response = await axiosInstance.get("/api/v1/admin/weights", {
    params,
  });
  return response;
};

export const getExcelInitialWeightScoringTableForInfectiousDiseases =
  async () => {
    const response = await axiosInstance.get(
      "/api/v1/admin/weights/export-template",
      {
        responseType: "blob",
      }
    );
    return response;
  };

export const updateBatchWeightScoring = async (params) => {
  const response = await axiosInstance.put(
    "/api/v1/admin/weights/batch",
    params
  );
  return response;
};

export const updateBatchWeightScoringAndDiagnosisSoring = async (params) => {
  const response = await axiosInstance.put(
    "/api/v1/admin/weights/batch-with-diagnosis",
    params
  );
  return response;
};
