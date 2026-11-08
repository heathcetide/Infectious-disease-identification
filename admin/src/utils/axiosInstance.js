import axios from 'axios';

// Embedded admin (prod build on :8080): same-origin empty baseURL.
// Vite DEV: default to local Go API. Override anytime with VITE_API_BASE.
const apiBase =
  (typeof import.meta !== 'undefined' && import.meta.env?.VITE_API_BASE) ||
  (typeof import.meta !== 'undefined' && import.meta.env?.DEV
    ? 'http://127.0.0.1:8080'
    : '');

const axiosInstance = axios.create({
    baseURL: apiBase,
    timeout: 300000,
});

function getStoredToken() {
    try {
        const raw = localStorage.getItem('user');
        if (!raw) return null;
        const user = JSON.parse(raw);
        return user?.token ?? null;
    } catch {
        return null;
    }
}

axiosInstance.interceptors.request.use(
    (config) => {
      const token = getStoredToken();
      if (token) {
        config.headers['Authorization'] = `Bearer ${token}`;
      }
      return config;
    },
    (error) => Promise.reject(error)
);

axiosInstance.interceptors.response.use(
    (response) => {
      const res = response;
      const data = res.data;
      if (
        data &&
        typeof data === 'object' &&
        !(data instanceof Blob) &&
        !(data instanceof ArrayBuffer)
      ) {
        if (data.code === 200) {
          data.code = 1;
        }
        if (data.data && data.data.list && !data.data.records) {
          data.data.records = data.data.list;
        }
        if (
          data.data &&
          typeof data.data === 'object' &&
          typeof data.data.isAI === 'boolean' &&
          data.data.isAI !== undefined &&
          Object.keys(data.data).length <= 2
        ) {
          data.data = data.data.isAI ? 'AI诊断' : '专家诊断';
        }
      }
      return response;
    },
    (error) => Promise.reject(error)
);

export default axiosInstance;
