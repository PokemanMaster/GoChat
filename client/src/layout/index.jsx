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
                        key="/layout/home"
                        icon={<AppstoreOutlined/>}
                        onClick={() => navigateTo(`/layout/home?userId=${userId}`)}
                    >
                        商品的首页
                    </Menu.Item>

                    <Menu.Item
                        key="/layout/product"
                        icon={<InsertRowAboveOutlined/>}
                        onClick={() => navigateTo(`/layout/product?userId=${userId}`)}
                    >
                        商品分类页
                    </Menu.Item>

                    <Menu.Item
                        key="/layout/personal/center"
                        icon={<AppstoreOutlined/>}
                        onClick={() => navigateTo(`/layout/personal/center?userId=${userId}`)}
                    >
                        我的个人中心
                    </Menu.Item>

                    <Menu.Item
                        key="/layout/personal/order"
                        icon={<AppstoreOutlined/>}
                        onClick={() => navigateTo(`/layout/personal/order?userId=${userId}`)}
                    >
                        我的订单
                    </Menu.Item>

                    <Menu.Item
                        key="/layout/personal/cart"
                        icon={<AppstoreOutlined/>}
                        onClick={() => navigateTo(`/layout/personal/cart?userId=${userId}`)}
                    >
                        我的购物车
                    </Menu.Item>

                    <Menu.Item
                        key="/layout/user/pass"
                        icon={<AppstoreOutlined/>}
                        onClick={() => navigateTo(`/layout/user/service?userId=${userId}`)}
                    >
                        登录与安全
                    </Menu.Item>

                    <Menu.Item
                        key="/layout/user/account"
                        icon={<AppstoreOutlined/>}
                        onClick={() => navigateTo(`/layout/user/account?userId=${userId}`)}
                    >
                        个人信息
                    </Menu.Item>

                    <Menu.Item
                        key="/layout/user/address"
                        icon={<AppstoreOutlined/>}
                        onClick={() => navigateTo(`/layout/user/address?userId=${userId}`)}
                    >
                        收货地址
                    </Menu.Item>

                    <Menu.Item
                        key="/layout/chat"
                        icon={<AppstoreOutlined/>}
                        onClick={() => navigateTo(`/layout/chat?userId=${userId}`)}
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
