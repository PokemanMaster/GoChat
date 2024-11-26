import request from '../index'


const token = localStorage.getItem("token");


// 获取所有商品分类
export const ListCategoriesAPI = () => {
    return request("/api/v1/products/categories", {
        method: 'get', headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};

// 获取所有商品分类
export const ListProductsAPI = () => {
    return request("/api/v1/products", {
        method: 'get', headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};

// 获取某个商品详情
export const ShowProductAPI = (id) => {
    return request(`/api/v1/products/${id}`, {
        method: 'get', headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};

// 获取所有商品图片参数 (首页展示的商品图片)
export const ListProductsParamsAPI = (params) => {
    return request("api/v1/products/param", {
        method: 'get', params: params,
    });
};


// 获取某个商品参数 (点击商品进入详情页)
export const ShowProductParamAPI = (id) => {
    return request(`/api/v1/products/${id}/param`, {
        method: 'get', headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};


// 搜索商品的接口
export const SearchProductsAPI = (data) => {
    return request("api/v1/searches", {
        method: 'post', data: data, headers: {
            'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json',
        },
    });
};