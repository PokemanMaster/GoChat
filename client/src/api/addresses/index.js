import request from '../index'

const token = localStorage.getItem("token");
const queryParams = new URLSearchParams(window.location.search);
const userId = queryParams.get('userId');

// 创建收货地址
export const CreateAddressAPI = (data) => {
    return request("api/v1/addresses", {
        method: 'post', data: data, headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};


// 展示收货地址
export const ShowAddressesAPI = () => {
    return request(`api/v1/addresses/${userId}`, {
        method: 'get', headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};


// 修改收货地址
export const UpdateAddressAPI = (data) => {
    return request("api/v1/addresses", {
        method: 'put', data: data, headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};


// 删除收货地址
export const DeleteAddressAPI = (data) => {
    return request("api/v1/addresses", {
        method: 'delete', data: data, headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};
