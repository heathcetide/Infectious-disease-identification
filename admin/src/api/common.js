import axiosInstance from "@/utils/axiosInstance.js";

export const uploadFile = async (formData) => {
  const response = await axiosInstance.post("/api/v1/files/upload", formData, {
    headers: { "Content-Type": "multipart/form-data" },
  });
  return response;
};
