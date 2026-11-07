import axiosInstance from "@/utils/axiosInstance.js";

export const pageSelectDiagnosis = async (params) => {
  const response = await axiosInstance.post(
    "/api/v1/medical/diagnosis/page",
    params
  );
  return response;
};

export const getExcelClinicalInformationExportForm = async () => {
  const response = await axiosInstance.get(
    "/api/v1/medical/diagnosis/export",
    { responseType: "blob" }
  );
  return response;
};

export const saveDiagnosis = async (data, userId) => {
  const response = await axiosInstance.post("/api/v1/medical/diagnosis", data, {
    params: { userId },
  });
  return response;
};

export const selectDiagnosis = async (diagnosisResultsID) => {
  const response = await axiosInstance.get(
    `/api/v1/medical/diagnosis/${diagnosisResultsID}`
  );
  return response;
};

export const updateDiagnosis = async (params) => {
  const id =
    params.diagnosisResultsID ?? params.id ?? params.diagnosisId;
  const response = await axiosInstance.put(
    `/api/v1/medical/diagnosis/${id}`,
    params
  );
  return response;
};

export const getDiagnosisReport = async (reportId) => {
  const response = await axiosInstance.get(
    `/api/v1/infect-files/diagnosis/${reportId}`,
    { responseType: "blob" }
  );
  return response;
};
