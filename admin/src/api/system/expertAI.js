import axiosInstance from "@/utils/axiosInstance.js";

export const systemChange = async (isAI) => {
  const response = await axiosInstance.put("/api/v1/admin/settings/ai-mode", {
    isAI,
  });
  return response;
};

export const systemChangeSearch = async () => {
  const response = await axiosInstance.get("/api/v1/admin/settings/ai-mode");
  return response;
};
