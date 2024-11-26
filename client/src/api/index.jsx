import axios from "axios";

// 创建axios的实例
// https://www.lvyouwang.xyz/
const instance = axios.create({
    // 确保使用 HTTPS 协议
    baseURL: 'http://localhost:9000',
    // 这个时间是你每次请求过期的时间，请求10秒之后就是失败的
    timeout: 10000,
});

// 请求拦截器
instance.interceptors.request.use(config => {
    return config
}, error => {
    return Promise.reject(error)
})

// 响应拦截器
instance.interceptors.response.use(res => {
    return res.data
}, error => {
    return Promise.reject(error)
})

export default instance;
