import request from '../index'

const token = localStorage.getItem("token");

// 排行榜
export const ListRankingAPI = () => {
    return request("api/v1/rankings", {
        method: 'get', headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};


export const ListElecRankingAPI = () => {
    return request("api/v1/elec-rankings", {
        method: 'get', headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
}


export const ListAcceRankingAPI = () => {
    return request("api/v1/acce-rankings", {
        method: 'get', headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};