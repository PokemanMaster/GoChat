import React from 'react';
import {Layout, Menu} from 'antd';
import {AppstoreOutlined, InsertRowAboveOutlined, UserOutlined} from '@ant-design/icons';
import "./layout.less"
import {Outlet, useLocation, useNavigate} from "react-router-dom";

const {Sider, Content} = Layout;

export default function LayoutView() {
    const navigateTo = useNavigate();
    const currentRoute = useLocation();

    return (
        <Layout className={"layout"}>
            <Sider breakpoint="lg" collapsedWidth="0" className={"sider"}>
                <div className={"logo"}/>
                <Menu theme="dark" defaultSelectedKeys={[currentRoute.pathname]} mode="inline"
                      defaultOpenKeys={['/layout/personal', '/layout/user']}>
                    <Menu.Item key="/layout/home" icon={<AppstoreOutlined/>}
                               onClick={() => navigateTo('/layout/home')}>
                        商品的首页
                    </Menu.Item>

                    <Menu.Item key="/layout/goods" icon={<InsertRowAboveOutlined/>}
                               onClick={() => navigateTo('/layout/goods')}>
                        商品分类页
                    </Menu.Item>

                    <Menu.SubMenu key="/layout/personal" icon={<UserOutlined/>} title="个人中心">
                        <Menu.Item key="/layout/personal/center" onClick={() => navigateTo('/layout/personal/center')}>
                            我的个人中心
                        </Menu.Item>
                        <Menu.Item key="/layout/personal/favorite"
                                   onClick={() => navigateTo('/layout/personal/favorite')}>
                            我的收藏
                        </Menu.Item>
                        <Menu.Item key="/layout/personal/order" onClick={() => navigateTo('/layout/personal/order')}>
                            我的订单
                        </Menu.Item>
                        <Menu.Item key="/layout/personal/cart" onClick={() => navigateTo('/layout/personal/cart')}>
                            我的购物车
                        </Menu.Item>
                    </Menu.SubMenu>

                    <Menu.SubMenu key="/layout/user" icon={<UserOutlined/>} title="账号管理">
                        <Menu.Item key="/layout/user/pass" onClick={() => navigateTo('/layout/user/service')}>
                            登录与安全
                        </Menu.Item>
                        <Menu.Item key="/layout/user/account" onClick={() => navigateTo('/layout/user/account')}>
                            个人信息
                        </Menu.Item>
                        <Menu.Item key="/layout/user/address" onClick={() => navigateTo('/layout/user/address')}>
                            收货地址
                        </Menu.Item>
                    </Menu.SubMenu>

                    <Menu.SubMenu key="/layout/about" icon={<UserOutlined/>} title="关于网站">
                        <Menu.Item key="/layout/about1" onClick={() => navigateTo('/layout/about')}>
                            网站地址
                        </Menu.Item>
                        <Menu.Item key="/layout/about2" onClick={() => navigateTo('/layout/about')}>
                            项目地址(前端)
                        </Menu.Item>
                        <Menu.Item key="/layout/about3" onClick={() => navigateTo('/layout/about')}>
                            项目地址(后端)
                        </Menu.Item>
                    </Menu.SubMenu>
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
