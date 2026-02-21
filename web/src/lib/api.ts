import axios from 'axios';

export const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

export const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
}, (error) => {
  return Promise.reject(error);
});

api.interceptors.response.use((response) => {
  return response;
}, (error) => {
  if (error.response?.status === 401) {
    // Redirect to login if unauthorized
    // But we are outside React component, so we can't use useNavigate
    // We can clear token and reload or window.location.href
    localStorage.removeItem('token');
    if (!window.location.pathname.includes('/login') && !window.location.pathname.includes('/register')) {
       window.location.href = '/login';
    }
  }
  return Promise.reject(error);
});
