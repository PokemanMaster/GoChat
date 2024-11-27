import React, { lazy, Suspense } from 'react';
import { Navigate, useRoutes } from 'react-router-dom';
import LayoutView from "../src/layout";

// 懒加载组件
const About = lazy(() => import("../src/views/About"));
const Chat = lazy(() => import("../src/views/chat/chat"));
const Cart = lazy(() => import("../src/views/Cart"));
const Center = lazy(() => import("../src/views/User/Center"));
const ConfirmOrder = lazy(() => import("../src/views/Order/ConfirmOrder"));
const Details = lazy(() => import("../src/views/Product/Details"));
const Favorite = lazy(() => import("../src/views/Favorite"));
const Goods = lazy(() => import("../src/views/Product/Goods"));
const Home = lazy(() => import("./views/home"));
const Login = lazy(() => import("../src/views/User/Login"));
const Order = lazy(() => import("../src/views/Order/list_order"));
const OrderDetails = lazy(() => import("../src/views/Payment"));
const Register = lazy(() => import("../src/views/User/Register"));
const UserAddress = lazy(() => import("../src/views/User/Address"));
const UserAccount = lazy(() => import("../src/views/User/Account"));
const UserService = lazy(() => import("../src/views/User/Service"));
const UserServicePassword = lazy(() => import("../src/views/User/ValidPassword"));
const UserServiceTelephone = lazy(() => import("../src/views/User/ValidTelephone"));

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
                path: "goods",
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
                        path: "favorite",
                        element: withLoadingComponent(<Favorite />)
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
                path: "about",
                element: withLoadingComponent(<About />)
            },
            {
                path: "order/details/",
                element: withLoadingComponent(<OrderDetails />)
            },
            {
                path: "order/confirm/:id",
                element: withLoadingComponent(<ConfirmOrder />)
            },
            {
                path: "product/:id",
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
