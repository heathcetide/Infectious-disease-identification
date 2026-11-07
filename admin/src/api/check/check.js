import axiosInstance from "@/utils/axiosInstance.js";

export const pageSelectLabTest = async (data) => {
  const response = await axiosInstance.post("/api/v1/medical/lab/page", data);
  return response;
};

export const getDetectionInformationExportFormExcel = async () => {
  const response = await axiosInstance.get("/api/v1/medical/lab/export", {
    responseType: "blob",
  });
  return response;
};

export const getUserByBaseInfo = async (data) => {
  const response = await axiosInstance.post(
    "/api/v1/medical/users/lookup",
    data
  );
  return response;
};

export const saveLabTestReport = async (data, userId) => {
  const response = await axiosInstance.post("/api/v1/medical/lab", data, {
    params: { userId },
  });
  return response;
};

export const selectLabTest = async (labTestReportID) => {
  const response = await axiosInstance.get(
    `/api/v1/medical/lab/${labTestReportID}`
  );
  return response;
};

export const saveLabTestFile = (formData) => {
  return axiosInstance.post("/api/v1/lab/file", formData, {
    headers: { "Content-Type": "multipart/form-data" },
  });
};

export const getLabTestFile = async (labTestId) => {
  const response = await axiosInstance.get(
    `/api/v1/infect-files/lab/${labTestId}`,
    { responseType: "blob" }
  );
  return response;
};

export const updateLabTest = async (data) => {
  const id = data.labTestReportID ?? data.id ?? data.labTestId;
  const response = await axiosInstance.put(`/api/v1/medical/lab/${id}`, data);
  return response;
};
