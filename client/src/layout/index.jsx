import React from 'react';
import {Layout, Menu} from 'antd';
import {AppstoreOutlined, InsertRowAboveOutlined} from '@ant-design/icons';
import "./layout.less";
import {Outlet, useLocation, useNavigate} from "react-router-dom";

const {Sider, Content} = Layout;

export default function LayoutView() {
    const navigateTo = useNavigate();
    const currentRoute = useLocation();
    const queryParams = new URLSearchParams(currentRoute.search);
    const userId = queryParams.get('userId');

    return (
        <Layout className={"layout"}>
            <Sider breakpoint="lg" collapsedWidth="0" className={"sider"}>
                <img className={"logo"} src="/logo/logo.png" alt=""/>
                <Menu
                    defaultSelectedKeys={[currentRoute.pathname]}
                    mode="inline"
                >
                    <Menu.Item
                        key="/layout/products/sort"
                        icon={<InsertRowAboveOutlined/>}
                        onClick={() => navigateTo(`/layout/products/sort?userId=${userId}`)}
                    >
                        商品分类页
                    </Menu.Item>

                    <Menu.Item
                        key="/layout/my/center"
                        icon={<AppstoreOutlined/>}
                        onClick={() => navigateTo(`/layout/my/center?userId=${userId}`)}
                    >
                        我的个人中心
                    </Menu.Item>

                    <Menu.Item
                        key="/layout/my/orders"
                        icon={<AppstoreOutlined/>}
                        onClick={() => navigateTo(`/layout/my/orders?userId=${userId}`)}
                    >
                        我的订单
                    </Menu.Item>

                    <Menu.Item
                        key="/layout/my/carts"
                        icon={<AppstoreOutlined/>}
                        onClick={() => navigateTo(`/layout/my/carts?userId=${userId}`)}
                    >
                        我的购物车
                    </Menu.Item>

                    <Menu.Item
                        key="/layout/my/service"
                        icon={<AppstoreOutlined/>}
                        onClick={() => navigateTo(`/layout/my/service?userId=${userId}`)}
                    >
                        登录与安全
                    </Menu.Item>

                    <Menu.Item
                        key="/layout/my/account"
                        icon={<AppstoreOutlined/>}
                        onClick={() => navigateTo(`/layout/my/account?userId=${userId}`)}
                    >
                        个人信息
                    </Menu.Item>

                    <Menu.Item
                        key="/layout/my/addresses"
                        icon={<AppstoreOutlined/>}
                        onClick={() => navigateTo(`/layout/my/addresses?userId=${userId}`)}
                    >
                        收货地址
                    </Menu.Item>

                    <Menu.Item
                        key="/layout/my/chat"
                        icon={<AppstoreOutlined/>}
                        onClick={() => navigateTo(`/layout/my/chat?userId=${userId}`)}
                    >
                        聊天
                    </Menu.Item>
                </Menu>
            </Sider>
            <Layout className="content">
                <Content>
                    <div>
                        <Outlet/>
                    </div>
                </Content>
            </Layout>
        </Layout>
    );
}
