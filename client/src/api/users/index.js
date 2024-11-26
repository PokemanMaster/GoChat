import request from '../index'

const token = localStorage.getItem("token");

// 检测用户token
export const UserTokenAPI = (token) => {
    return request("/api/v1/ping", {
        method: 'get', headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    })
}

// 用户注册
export const RegisterAPI = (data) => {
    return request("api/v1/user/register", {
        method: "post", data: data, credentials: 'include', // Ensure cookies are included
    });
}

// 登录
export const LoginAPI = (data) => {
    return request("api/v1/user/login", {
        method: "post", data: data, credentials: 'include', // Ensure cookies are included
    });
}

// 用户修改信息操作
export const UpdateAPI = (data) => {
    return request("api/v1/user", {
        method: 'put', data: data, headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    })
}

// 用户修改密码接口
export const UpdatePasswordAPI = (data) => {
    return request("api/v1/user/password", {
        method: 'put', data: data, headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    })
}

// 用户绑定手机号接口
export const UpdateTelephoneAPI = (data) => {
    return request("api/v1/user/telephone", {
        method: 'post', data: data, headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    })
}

// 验证码
export const CategoryAPI = () => {
    return request("/api/v1/captcha/image", {
        method: "get",
    });
}