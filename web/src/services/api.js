import axios from 'axios';

// 定义请求前缀
const baseURL = '/api/v1';

// 创建 Axios 实例
const api = axios.create({
    baseURL: baseURL
});

export default api;