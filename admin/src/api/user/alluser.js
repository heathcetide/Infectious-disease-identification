import axiosInstance from "@/utils/axiosInstance.js";

export const fetchUserData = async (params) => {
  const response = await axiosInstance.post("/api/v1/admin/users/page", params, {
    headers: { "Content-Type": "application/json" },
  });
  return response;
};

export const toggleUserStatus = async (userId) => {
  const response = await axiosInstance.put(
    `/api/v1/admin/users/${userId}/active`
  );
  return response;
};

export const initializeUserPassword = async (userId) => {
  const response = await axiosInstance.put(
    `/api/v1/admin/users/${userId}/password/reset`
  );
  return response;
};

export const exportUserData = async () => {
  const response = await axiosInstance.get("/api/v1/admin/users/export", {
    responseType: "blob",
  });
  return response;
};

export const importUsersByExcel = async (formData) => {
  const response = await axiosInstance.post(
    "/api/v1/admin/users/import",
    formData,
    {
      headers: { "Content-Type": "multipart/form-data" },
    }
  );
  return response;
};

export const fetchUserInfoById = async (userId) => {
  const response = await axiosInstance.get(`/api/v1/admin/users/${userId}`);
  return response;
};

export const updateUserInfo = async (requestData) => {
  const userId = requestData.userId ?? requestData.id;
  const response = await axiosInstance.put(
    `/api/v1/admin/users/${userId}`,
    requestData,
    {
      headers: { "Content-Type": "application/json" },
    }
  );
  return response;
};

export const queryAllUserToExcel = async () => {
  const response = await axiosInstance.get("/api/v1/admin/users/export", {
    responseType: "blob",
  });
  return response;
};

export const addOneUser = async (requestData) => {
  const response = await axiosInstance.post(
    "/api/v1/admin/users",
    requestData,
    {
      headers: { "Content-Type": "application/json" },
    }
  );
  return response;
};
