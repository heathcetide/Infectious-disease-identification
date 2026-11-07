import axiosInstance from "@/utils/axiosInstance.js";

export const fetchFeedbackData = async (data) => {
  const response = await axiosInstance.get("/api/v1/admin/feedback", {
    params: data,
  });
  return response;
};

export const fetchFeedbackDataById = async (feedbackId) => {
  const response = await axiosInstance.get(
    `/api/v1/admin/feedback/${feedbackId}`
  );
  return response;
};

export const exportFeedbackData = async () => {
  const response = await axiosInstance.get("/api/v1/admin/feedback/export", {
    responseType: "blob",
  });
  return response;
};
