import request from '../index'

const queryParams = new URLSearchParams(window.location.search);
const userId = queryParams.get('userId');
const token = localStorage.getItem("token");

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
    return request(`api/v1/user/${userId}/orders`, {
        method: 'get', headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};