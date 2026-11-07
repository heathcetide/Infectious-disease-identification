import axiosInstance from "@/utils/axiosInstance.js";

export const startAI = async () => {
  const response = await axiosInstance.post("/api/v1/ai/auto-weight/run");
  return response;
};

export const updateAutoAIWeightUpdater = async (autoAIWeightUpdater) => {
  const response = await axiosInstance.put("/api/v1/ai/auto-weight", {
    enabled: autoAIWeightUpdater,
  });
  return response;
};

export const getAutoAIWeightUpdater = async () => {
  const response = await axiosInstance.get("/api/v1/ai/auto-weight");
  return response;
};
