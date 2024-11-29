import React, {useEffect, useState} from "react";
import "./style.less"
import {Button, Col, Empty, Row} from "antd";
import {Typography} from 'antd';
import {useNavigate} from "react-router-dom";
import {ListOrdersAPI} from "../../../api/orders";

export default function Center() {
    const queryParams = new URLSearchParams(window.location.search);
    const userId = queryParams.get('userId');
    const {Text} = Typography;
    const [UserInfo] = useState(JSON.parse(localStorage.getItem("user" + userId)));
    const navigateTo = useNavigate()
    const [OrderNum, setOrderNum] = useState(0)
    const [FavoritesNum, setFavoritesNum] = useState(0)


    useEffect(() => {
        if (UserInfo) {
            ListOrdersAPI().then(res => {
                console.log("ListOrders", res.data.items);
                setOrderNum(res.data.items.length);
            }).catch(err => {
                console.log(err)
            });
        }
    }, [UserInfo]);


    return (<div className={"CenterLayout"}>
        {userId ? (<div className={"CenterContent"}>
                {/* 个人信息头部 */}
                <div className={"CenterUser"}>
                    <Row>
                        <Col xs={24} sm={24} md={12} lg={12} xl={12}>
                            <div className={"User"}>
                                <div className={"UserAvator"}>
                                    <img src={UserInfo.avatar} alt=""/>
                                </div>
                                <div className={"UserInfo"}>
                                    <p style={{fontSize: '25px', fontWeight: 100, color: '#464547'}}>
                                        {UserInfo.user_name}
                                    </p>
                                    <span>下午好~</span>
                                    <p>
                                        <i
                                            onClick={() => {
                                                navigateTo("/layout/user/account")
                                            }}
                                            style={{fontSize: '13px', color: '#ff6700'}}>修改个人信息 &gt;
                                        </i>
                                    </p>
                                </div>
                            </div>
                        </Col>
                        <Col xs={24} sm={24} md={12} lg={12} xl={12}>
                            <div className={"MSG"}>
                                <div className="account-security">
                                    <div className="security-level">
                                        <span>账户安全：</span>
                                        <Text type="success" className="high-security">较高</Text>
                                    </div>
                                    <div className="bound-info">
                                        <span>绑定手机：</span>
                                        <Text>{UserInfo.telephone}</Text>
                                    </div>
                                    <div className="bound-info">
                                        <span>绑定邮箱：</span>
                                        <Text>{UserInfo.telephone}@qq.com</Text>
                                    </div>
                                </div>
                            </div>
                        </Col>
                    </Row>
                </div>
                <div>
                    <Row>
                        <Col xs={24} sm={24} md={12} lg={12} xl={12}>
                            <div className={"UserDetails"} onClick={() => {
                                navigateTo("/layout/personal/order")
                            }}>
                                <div>
                                    <img src="https://s01.mifile.cn/i/user/portal-icon-1.png" alt={""}/>
                                </div>
                                <div className={"Operate"}>
                                    <p className={"OperTitle"}>待支付订单：{OrderNum}</p>
                                    <p>
                                        <router-link to="/order?type=1"
                                                     className={"OperHref"}>查看待支付订单 &gt;</router-link>
                                    </p>
                                </div>
                            </div>
                        </Col>
                        <Col xs={24} sm={24} md={12} lg={12} xl={12}>
                            <div className={"UserDetails"} onClick={() => {
                                navigateTo("/layout/personal/order")
                            }}>
                                <div>
                                    <img src="https://s01.mifile.cn/i/user/portal-icon-2.png" alt={""}/>
                                </div>
                                <div className={"Operate"}>
                                    <p className={"OperTitle"}>已付款订单：{OrderNum}</p>
                                    <p>
                                        <router-link to="/order?type=2"
                                                     className={"OperHref"}>查看已付款订单 &gt;</router-link>
                                    </p>
                                </div>
                            </div>
                        </Col>
                        <Col xs={24} sm={24} md={12} lg={12} xl={12}>
                            <div className={"UserDetails"} onClick={() => {
                                navigateTo("/layout/personal/order")
                            }}>
                                <div>
                                    <img src="https://s01.mifile.cn/i/user/portal-icon-3.png" alt={""}/>
                                </div>
                                <div className={"Operate"}>
                                    <p className={"OperTitle"}>待评价商品：{OrderNum}</p>
                                    <p>
                                        <router-link to="/"
                                                     className={"OperHref"}>查看待评价订单 &gt;</router-link>
                                    </p>
                                </div>
                            </div>
                        </Col>
                    </Row>
                </div>
            </div>) : // 用户没登录就显示
            <div className={"Empty"}>
                <Empty
                    image="https://gw.alipayobjects.com/zos/antfincdn/ZHrcdLPrvN/empty.svg"
                    imageStyle={{
                        height: 160,
                    }}
                    description={<span>你还没有 <a href="/client/src/public">登录？</a></span>}>
                    <Button type="primary" onClick={() => {
                        navigateTo("/login")
                    }}>点击登录</Button>
                </Empty>
            </div>}
    </div>)
}
