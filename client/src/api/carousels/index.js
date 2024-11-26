import request from '../index'

const token = localStorage.getItem("token");

// 轮播图
export const ListCarouselsAPI = () => {
    return request("api/v1/carousels", {
        method: 'get', headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};
