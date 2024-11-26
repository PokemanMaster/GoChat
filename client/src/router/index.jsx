import React, {lazy} from 'react';
import {Navigate} from "react-router-dom";
import LayoutView from "../layout";

const About = lazy(() => import("../views/About"));
// const CallbackQQ = lazy(() => import("../views/CallbackQQ"));
const Cart = lazy(() => import("../views/Cart"));
const Center = lazy(() => import("../views/Center"));
const ConfirmOrder = lazy(() => import("../views/ConfirmOrder"));
const Details = lazy(() => import("../views/Details"));
const Favorite = lazy(() => import("../views/Favorite"));
const Goods = lazy(() => import("../views/Goods"));
const Home = lazy(() => import("../views/Home"));
const Login = lazy(() => import("../views/Login"));
const Order = lazy(() => import("../views/Order"));
const OrderDetails = lazy(() => import("../views/Payment"));
const Register = lazy(() => import("../views/Register"));
const UserAddress = lazy(() => import("../views/UserAddress"));
const UserAccount = lazy(() => import("../views/UserAccount"));
const UserService = lazy(() => import("../views/UserService"));
const UserServicePassword = lazy(() => import("../views/ValidPassword"));
const UserServiceTelephone = lazy(() => import("../views/ValidTelephone"));


// 懒加载组件加载时的 Loading 界面
const withLoadingComponent = (component) => (
    <React.Suspense fallback={<div>loading...</div>}>
        {component}
    </React.Suspense>
);

const router = [
    {
        path: "/",
        element: <Navigate to="/login"/>
    },
    {
        path: "/layout",
        element: <LayoutView/>,
        children: [
            {
                path: "home",
                element: withLoadingComponent(<Home/>)
            },
            {
                path: "goods",
                element: withLoadingComponent(<Goods/>)
            },
            {
                path: "personal",
                children: [
                    {
                        path: "center",
                        element: withLoadingComponent(<Center/>)
                    },
                    {
                        path: "favorite",
                        element: withLoadingComponent(<Favorite/>)
                    },
                    {
                        path: "order",
                        element: withLoadingComponent(<Order/>)
                    },
                    {
                        path: "cart",
                        element: withLoadingComponent(<Cart/>)
                    },
                ]
            },
            {
                path: "user",
                children: [
                    {
                        path: "service",
                        element: withLoadingComponent(<UserService/>)
                    },
                    {
                        path: "service/password",
                        element: withLoadingComponent(<UserServicePassword/>)
                    },
                    {
                        path: "service/telephone",
                        element: withLoadingComponent(<UserServiceTelephone/>)
                    },
                    {
                        path: "account",
                        element: withLoadingComponent(<UserAccount/>)
                    },
                    {
                        path: "address",
                        element: withLoadingComponent(<UserAddress/>)
                    },
                ]
            },
            {
                path: "about",
                element: withLoadingComponent(<About/>)
            },
            {
                path: "order/details/",
                element: withLoadingComponent(<OrderDetails/>)
            },
            {
                path: "order/confirm/:id",
                element: withLoadingComponent(<ConfirmOrder/>)
            },
            {
                path: "product/:id",
                element: withLoadingComponent(<Details/>)
            },
        ]
    },
    {
        path: "/register",
        element: withLoadingComponent(<Register/>)
    },
    {
        path: "/login",
        element: withLoadingComponent(<Login/>)
    },
    {
        path: "*",
        element: <Navigate to="/layout/home"/>
    }
];


export default router;
