import React, {useEffect, useState} from "react";
import "./style.less"
import {ShowOrderAPI} from "../../../api/orders";
import {useLocation} from "react-router-dom";
import {Steps} from "antd";
import {ShowAddressesAPI} from "../../../api/addresses";
import {Step} from "@mui/joy";

export default function ConfirmOrder() {
    const location = useLocation();
    const [order, setOrder] = useState([])
    const [address, setAddress] = useState([])
    const orderNum = location.state && location.state.orderNum;

    useEffect(() => {
        ShowOrderAPI(orderNum).then(res => {
            setOrder(res.data);
            ShowAddressesAPI(res.data.user_id).then(res => {
                setAddress(res.data[0]);
                console.log("ShowAddressesAPI", res.data[0]);
            }).catch(err => {
                console.log(err);
            });
        });
    }, [location.state]); // 将 location.state 添加到依赖数组


    // 解析 JSON 格式图片
    function JsonParseFacade(value) {
        if (value) {
            try {
                const parsedValue = JSON.parse(value); // 解析 JSON 字符串
                return parsedValue.facade || ""; // 返回 facade 数组，如果不存在则返回空字符串
            } catch (error) {
                console.error("JSON解析错误:", error);
                return ""; // 返回空字符串或处理错误的方式
            }
        }
        return ""; // 如果 value 是 undefined，返回空字符串
    }

    const OrderDetails = [{
        id: "购物车总价",
        images: order.images,
        name: order.title,
        number: order.num,
        price: order.actualPrice * order.num
    },];


    return (
        <div className={"body"}>
            <div className="order-detail">
                <div className="order-header">
                    <h2>订单详情</h2>
                    <p>订单号：{order.code}</p>
                </div>
                <div className="order-status">
                    <p className="status-label">已支付</p>
                    <Steps size="small" current={3} className="order-steps">
                        <Step title="下单" description="06月29日 13:44"/>
                        <Step title="付款" description="06月29日 13:45"/>
                        <Step title="配货"/>
                        <Step title="出库"/>
                        <Step title="交易成功"/>
                    </Steps>
                </div>
                <div className="order-items">
                    <div className="item">
                        <img src={JsonParseFacade(order.images)} alt="Item 2"/>
                        <div className="item-info">
                            <p>{order.title}</p>
                            <p>{order.actualPrice}元 × {order.num}</p>
                        </div>
                    </div>
                </div>
                <div className="order-section">
                    <h3>收货信息</h3>
                    <p>姓名：{address.name}</p>
                    <p>联系电话：{address.Telephone}</p>
                    <p>收货地址：{address.address}</p>
                </div>
                <div className="order-section">
                    <h3>支付方式及送货时间</h3>
                    <p>支付方式：在线支付</p>
                    <p>送货时间：不限送货时间</p>
                    <p>送达时间：预计明天送达</p>
                </div>
            </div>
        </div>
    )
}