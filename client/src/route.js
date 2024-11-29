import React, { lazy, Suspense } from 'react';
import { Navigate, useRoutes } from 'react-router-dom';
import LayoutView from "../src/layout";

// 懒加载组件
const Chat = lazy(() => import("../src/views/chat/chat"));
const Cart = lazy(() => import("../src/views/cart"));
const Center = lazy(() => import("../src/views/user/center"));
const ConfirmOrder = lazy(() => import("./views/order/order_details"));
const Details = lazy(() => import("../src/views/product/product_details"));
const Goods = lazy(() => import("./views/product/product_sort"));
const Home = lazy(() => import("../src/views/home"));
const Login = lazy(() => import("../src/views/user/login"));
const Order = lazy(() => import("./views/order/my_orders"));
const OrderDetails = lazy(() => import("./views/order/order_pay"));
const Register = lazy(() => import("../src/views/user/register"));
const UserAddress = lazy(() => import("../src/views/user/address"));
const UserAccount = lazy(() => import("../src/views/user/account"));
const UserService = lazy(() => import("../src/views/user/service"));
const UserServicePassword = lazy(() => import("../src/views/user/service_password"));
const UserServiceTelephone = lazy(() => import("../src/views/user/service_telephone"));

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
                path: "chat",
                element: withLoadingComponent(<Chat />)
            },
            {
                path: "home",
                element: withLoadingComponent(<Home />)
            },
            {
                path: "product",
                element: withLoadingComponent(<Goods />)
            },
            {
                path: "personal",
                children: [
                    {
                        path: "center",
                        element: withLoadingComponent(<Center />)
                    },
                    {
                        path: "order",
                        element: withLoadingComponent(<Order />)
                    },
                    {
                        path: "cart",
                        element: withLoadingComponent(<Cart />)
                    },
                ]
            },
            {
                path: "user",
                children: [
                    {
                        path: "service",
                        element: withLoadingComponent(<UserService />)
                    },
                    {
                        path: "service/password",
                        element: withLoadingComponent(<UserServicePassword />)
                    },
                    {
                        path: "service/telephone",
                        element: withLoadingComponent(<UserServiceTelephone />)
                    },
                    {
                        path: "account",
                        element: withLoadingComponent(<UserAccount />)
                    },
                    {
                        path: "address",
                        element: withLoadingComponent(<UserAddress />)
                    },
                ]
            },
            {
                path: "order/product_details/",
                element: withLoadingComponent(<OrderDetails />)
            },
            {
                path: "order/confirm/:id",
                element: withLoadingComponent(<ConfirmOrder />)
            },
            {
                path: "product_sort/:id",
                element: withLoadingComponent(<Details />)
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
