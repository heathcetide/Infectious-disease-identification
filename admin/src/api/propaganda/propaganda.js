import axiosInstance from "@/utils/axiosInstance.js";

export const pageSelectDiseaseLearningMaterials = async (params) => {
  const response = await axiosInstance.get("/api/v1/cdc/materials", {
    params,
    headers: { "Content-Type": "application/json" },
  });
  return response;
};

export const startOrStopMaterial = async (materialId, isDelete) => {
  const response = await axiosInstance.put(
    `/api/v1/cdc/materials/${materialId}/status`,
    { disabled: isDelete }
  );
  return response;
};

export const getMaterialById = async (materialId) => {
  const response = await axiosInstance.get(
    `/api/v1/cdc/materials/${materialId}`
  );
  return response;
};

export const getLearningMaterialFile = async (filename) => {
  const response = await axiosInstance.get(
    `/api/v1/infect-files/learning/${filename}`,
    { responseType: "blob" }
  );
  return response;
};

export const saveMaterial = async (data) => {
  const response = await axiosInstance.post("/api/v1/cdc/materials", data, {
    headers: { "Content-Type": "multipart/form-data" },
  });
  return response;
};

export const updateMaterialById = async (data) => {
  const materialId =
    (typeof data?.get === "function" ? data.get("materialId") : null) ??
    data.materialId;
  const response = await axiosInstance.put(
    `/api/v1/cdc/materials/${materialId}`,
    data,
    {
      headers:
        typeof data?.get === "function"
          ? { "Content-Type": "multipart/form-data" }
          : { "Content-Type": "application/json" },
    }
  );
  return response;
};
