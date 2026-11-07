import axiosInstance from "@/utils/axiosInstance.js";

export const fetchHospitalData = async (params) => {
  const response = await axiosInstance.post(
    "/api/v1/admin/hospitals/page",
    params,
    {
      headers: { "Content-Type": "application/json" },
    }
  );
  return response;
};

export const exportHospitalData = async () => {
  const response = await axiosInstance.get("/api/v1/admin/hospitals/export", {
    responseType: "blob",
  });
  return response;
};

export const getExcelHospitalTemplate = async () => {
  const response = await axiosInstance.get(
    "/api/v1/admin/hospitals/import/template",
    { responseType: "blob" }
  );
  return response;
};

export const fetchHospitalDataById = async (hospitalId) => {
  const response = await axiosInstance.get(
    `/api/v1/admin/hospitals/${hospitalId}`
  );
  return response;
};

export const editHospitalDataById = async (params) => {
  const hospitalId = params.hospitalId ?? params.id;
  const response = await axiosInstance.put(
    `/api/v1/admin/hospitals/${hospitalId}`,
    params,
    {
      headers: { "Content-Type": "application/json" },
    }
  );
  return response;
};

export const putHospitalDataByExcel = async (formData) => {
  const response = await axiosInstance.post(
    "/api/v1/admin/hospitals/import",
    formData,
    {
      headers: { "Content-Type": "multipart/form-data" },
    }
  );
  return response;
};

export const deleteHospitalInfo = async (hospitalId) => {
  const response = await axiosInstance.delete(
    `/api/v1/admin/hospitals/${hospitalId}`
  );
  return response;
};
