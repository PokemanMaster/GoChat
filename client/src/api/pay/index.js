import request from '../index'

const token = localStorage.getItem("token");

// 支付
export const CreatePayAPI = (data) => {
    return request("api/v1/pay", {
        method: 'post', data: data, headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};


