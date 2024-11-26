import request from '../index'

let UserId = '';
const user = JSON.parse(localStorage.getItem("user"));
const token = localStorage.getItem("token");

if (user && user.id) {
    UserId = user.id;
}

// 创建订单
export const CreateOrderAPI = (data) => {
    return request("api/v1/orders", {
        method: 'post', data: data, headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};

// 获取订单
export const ShowOrderAPI = (num) => {
    return request(`api/v1/orders/${num}`, {
        method: 'get', headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};

// 获取某个用户所有订单
export const ListOrdersAPI = () => {
    return request(`api/v1/user/${UserId}/orders`, {
        method: 'get', headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};