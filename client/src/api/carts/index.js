import request from '../index'

const queryParams = new URLSearchParams(window.location.search);
const userId = parseInt(queryParams.get('userId') , 10);
const token = localStorage.getItem("token");


// 创建购物车
export const CreateCartAPI = (data) => {
    return request("api/v1/carts", {
        method: 'post', data: data, headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};

// 展示购物车
export const ShowCartAPI = () => {
    return request(`api/v1/carts/${userId}`, {
        method: 'get', headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};

// 修改购物车
export const UpdateCartAPI = (data) => {
    return request("api/v1/carts", {
        method: 'put', data: data, headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};

// 删除购物车
export const DeleteCartAPI = (data) => {
    return request("api/v1/carts", {
        method: 'delete', data: data, headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};
