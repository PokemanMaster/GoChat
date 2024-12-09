import React, { lazy, Suspense } from 'react';
import { Navigate, useRoutes } from 'react-router-dom';
import LayoutView from "../src/layout";

// 懒加载组件
const MyChat = lazy(() => import("../src/views/user/my_chat/chat"));
const MyCart = lazy(() => import("../src/views/user/my_carts"));
const MyCenter = lazy(() => import("../src/views/user/my_center"));
const OrderDetails = lazy(() => import("../src/views/order/order_details"));
const ProductDetails = lazy(() => import("../src/views/product/product_details"));
const ProductSort = lazy(() => import("../src/views/product/product_sort"));
const Login = lazy(() => import("../src/views/login"));
const MyOrders = lazy(() => import("../src/views/order/my_orders"));
const OrderPay = lazy(() => import("../src/views/order/order_pay"));
const Register = lazy(() => import("../src/views/register"));
const MyAddress = lazy(() => import("../src/views/user/my_addresses"));
const MyAccount = lazy(() => import("../src/views/user/my_account"));
const MyService = lazy(() => import("../src/views/user/my_service"));
const MyServicePassword = lazy(() => import("../src/views/user/my_service_password"));
const MyServiceTelephone = lazy(() => import("../src/views/user/my_service_telephone"));

// 懒加载包装组件
const withLoadingComponent = (component) => (
    <Suspense fallback={<div></div>}>
        {component}
    </Suspense>
);

// 定义路由表
const router = [
    {
        path: "/",
        element: <Navigate to="/login" />
    },
    {
        path: "/layout",
        element: <LayoutView />,
        children: [
            {
                path: "products/sort",  // 商品分类
                element: withLoadingComponent(<ProductSort />)
            },
            {
                path: "product/details/:id",  // 商品详情
                element: withLoadingComponent(<ProductDetails />)
            },
            {
                path: "my/center",  // 我的中心
                element: withLoadingComponent(<MyCenter />)
            },
            {
                path: "my/carts",  // 我的购物车
                element: withLoadingComponent(<MyCart />)
            },
            {
                path: "my/service",  // 我的服务
                element: withLoadingComponent(<MyService />)
            },
            {
                path: "my/service/password", // 我的服务=> 修改密码
                element: withLoadingComponent(<MyServicePassword />)
            },
            {
                path: "my/service/telephone", // 我的服务=> 修改手机号
                element: withLoadingComponent(<MyServiceTelephone />)
            },
            {
                path: "my/account", // 我的个人信息
                element: withLoadingComponent(<MyAccount />)
            },
            {
                path: "my/addresses", // 我的地址
                element: withLoadingComponent(<MyAddress />)
            },
            {
                path: "my/chat",  // 我的聊天
                element: withLoadingComponent(<MyChat />)
            },
            {
                path: "my/orders",  // 我的订单
                element: withLoadingComponent(<MyOrders />)
            },
            {
                path: "order/pay", // 订单支付
                element: withLoadingComponent(<OrderPay />)
            },
            {
                path: "order/details/:id", // 订单详情
                element: withLoadingComponent(<OrderDetails />)
            },
        ]
    },
    {
        path: "/register",
        element: withLoadingComponent(<Register />)
    },
    {
        path: "/login",
        element: withLoadingComponent(<Login />)
    },
    {
        path: "*",
        element: <Navigate to="/layout/home" />
    }
];

// 主组件
export default function Route() {
    const routes = useRoutes(router);
    return (
        <div>
            {routes}
        </div>
    );
}
