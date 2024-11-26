import {Button, Empty} from "antd";
import React, {useEffect, useState} from "react";
import "./style.less";
import {ListOrdersAPI} from "../../api/orders";
import {useNavigate} from "react-router-dom";
import emptyCart from "../../public/images/cart_empty.png";
import {Link} from "@mui/joy";
import clogo from "../../public/images/clogo.png";

export default function Order() {
    const [UserInfo] = useState(JSON.parse(localStorage.getItem("user")));
    const navigateTo = useNavigate();
    const [order, setOrder] = useState([]);
    const [currentTab, setCurrentTab] = useState("all");

    useEffect(() => {
        if (UserInfo) {
            ListOrdersAPI(UserInfo.id).then((res) => {
                if (res.data && res.data.items) {
                    setOrder(res.data.items);
                } else {
                    console.error("No items found in the response", res.data);
                    setOrder([]); // 如果没有 items，就设置为空数组，避免空值错误
                }
            }).catch((error) => {
                console.error("Error fetching orders:", error);
                setOrder([]); // 请求出错时处理
            });
        }
    }, [UserInfo]);


    function ViewTheOrder(item) {
        if (item.status === 1) {
            navigateTo("/layout/order/details/", {
                state: {Cart: item},
            });
        } else {
            navigateTo(`/layout/order/confirm/${item.code}`, {
                state: {orderNum: item.code},
            });
        }

    }

    function JsonParseFacade(value) {
        const parsedValue = JSON.parse(value);
        return parsedValue.facade;
    }

    function filterOrdersByTab() {
        if (currentTab === "all") {
            return order;
        } else if (currentTab === "unpaid") {
            return order.filter(item => item.status === 1);
        } else if (currentTab === "paid") {
            return order.filter(item => item.status === 2);
        } else if (currentTab === "shipped") {
            return order.filter(item => item.status === 3);
        } else if (currentTab === "received") {
            return order.filter(item => item.status === 4);
        }
    }

    function getOrderStatusText(status) {
        switch (status) {
            case 1:
                return "未付款";
            case 2:
                return "已付款";
            case 3:
                return "已发货";
            case 4:
                return "已签收";
            default:
                return "未知状态";
        }
    }

    return (
        <>
            {UserInfo ? (
                <div>
                    <div className="TopHeader">
                        <div className="CartHeader">
                            <div className="Logo">
                                <Link to="/">
                                    <img src={clogo} alt=""/>
                                </Link>
                            </div>
                            <div className="CartHeaderContent">
                                <p>我的订单</p>
                            </div>
                        </div>
                    </div>

                    <div className="order-header">
                        <div className="header">
                            <h2>我的订单</h2>
                            <p>请耐心等待，或者致电了解更多</p>
                        </div>
                        <div className="tabs">
                            <span className={`tab ${currentTab === "all" ? "active" : ""}`}
                                  onClick={() => setCurrentTab("all")}>全部有效订单</span>
                            <span className={`tab ${currentTab === "unpaid" ? "active" : ""}`}
                                  onClick={() => setCurrentTab("unpaid")}>待支付</span>
                            <span className={`tab ${currentTab === "paid" ? "active" : ""}`}
                                  onClick={() => setCurrentTab("paid")}>待收货</span>
                            <span className={`tab ${currentTab === "shipped" ? "active" : ""}`}
                                  onClick={() => setCurrentTab("shipped")}>已发货</span>
                            <span className={`tab ${currentTab === "received" ? "active" : ""}`}
                                  onClick={() => setCurrentTab("received")}>已签收</span>
                        </div>

                        {filterOrdersByTab() && filterOrdersByTab().length > 0 ? (
                            filterOrdersByTab().map(item => (
                                <div className="order-detail" key={item.id}>
                                    <div className="order-summary">
                                        <p>{getOrderStatusText(item.status)}</p>
                                        <div className="order-info">
                                            <p>2016年05月04日 18:56 | 米兔 | 订单号：{item.code} | 在线支付</p>
                                            <div className="order-total">
                                                <span>订单金额：</span>
                                                <span className="total-amount">
                                                    {(item.actualPrice).toFixed(2)}
                                                </span>&nbsp;元
                                            </div>
                                        </div>
                                    </div>
                                    <div className="order-items">
                                        <div className="item">
                                            <img src={JsonParseFacade(item.images)} alt="Phone"/>
                                            <div className="item-info">
                                                <p>{item.title}</p>
                                                <p>{item.price}元 × {item.num}</p>
                                            </div>
                                        </div>
                                    </div>
                                    <div className="order-actions">
                                        <Button type="primary"
                                                onClick={() => ViewTheOrder(item, item.status)}>订单详情</Button>
                                        <Button>申请售后</Button>
                                    </div>
                                </div>
                            ))
                        ) : (
                            <div className="EmptyCart">
                                <img src={emptyCart} alt="" className="EmptyCartImg"/>
                                <div className="EmptyCartText1">订单竟然是空的！</div>
                                <div className="EmptyCartText2">再忙，也要记得买点什么犒劳自己~</div>
                            </div>
                        )}
                    </div>
                </div>
            ) : (
                <div className="Empty">
                    <Empty
                        image="https://gw.alipayobjects.com/zos/antfincdn/ZHrcdLPrvN/empty.svg"
                        imageStyle={{height: 160}}
                        description={<span>你还没有 <a href="/">登录？</a></span>}
                    >
                        <Button type="primary" onClick={() => {
                            navigateTo("/login")
                        }}>点击登录</Button>
                    </Empty>
                </div>
            )}
        </>
    );
}
