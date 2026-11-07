import axios from 'axios';

const axiosInstance = axios.create({
    baseURL: 'http://127.0.0.1:8080',
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
