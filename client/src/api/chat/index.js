import request from '../index'

const token = localStorage.getItem("token");
const queryParams = new URLSearchParams(window.location.search);
const userId = queryParams.get('userId');

// 展示收藏夹 (Show favorite items)
export const ChatMessageAPI = (data) => {
    return request("api/v1/chat/message", {
        method: 'post',
        data: data,
        headers: {
            'Authorization': `Bearer ${token}`,  // Corrected string interpolation
            'Content-Type': 'application/json',
        },
    });
};

export const SearchFriendAPI = (data) => {
    return request("api/v1/friends/search", {
        method: 'post',
        data: data,
        headers: {
            'Authorization': `Bearer ${token}`,  // Corrected string interpolation
            'Content-Type': 'application/json',
        },
    });
};

export const FriendListsAPI = () => {
    return request(`api/v1/friends/${userId}`, {
        method: 'get',
        headers: {
            'Authorization': `Bearer ${token}`,  // Corrected string interpolation
            'Content-Type': 'application/json',
        },
    });
};

export const CreateFriendAPI = (data) => {
    return request("api/v1/friend", {
        method: 'post',
        data: data,
        headers: {
            'Authorization': `Bearer ${token}`,  // Corrected string interpolation
            'Content-Type': 'application/json',
        },
    });
};
