import request from '../index'

let UserId = '';
const user = JSON.parse(localStorage.getItem("user"));
const token = localStorage.getItem("token");

if (user && user.ID) {
    UserId = user.ID;
}


// 创建收藏夹
export const CreateFavoriteAPI = (data) => {
    return request("api/v1/favorites", {
        method: 'post', data: data, headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};


// 展示收藏夹
export const ShowFavoritesAPI = () => {
    return request(`api/v1/favorites/${UserId}`, {
        method: 'get', headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};


// 删除收藏夹
export const DeleteFavoriteAPI = (data) => {
    return request("api/v1/favorites", {
        method: 'DELETE', data: data, headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};
