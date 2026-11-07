import axiosInstance from "@/utils/axiosInstance.js";

export const loginUser = async (params) => {
  const response = await axiosInstance.post("/api/v1/auth/login", params);
  return response;
};

export const changePassword = async (data) => {
  const response = await axiosInstance.put("/api/v1/medical/password", data);
  return response;
};

export const updateProfile = async (data) => {
  const response = await axiosInstance.put("/api/v1/employee/profile", data);
  return response;
};
