import axios from 'axios';

const instance = axios.create({
  baseURL: 'http://47.242.39.174:8080',
});

export const flask = axios.create({
  baseURL: 'http://47.242.39.174:8081', // 后端 B 的地址
});

instance.interceptors.request.use(config => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = token;
    console.log(token)
  }
  return config;
});

export default instance;