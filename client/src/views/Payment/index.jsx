import React, {useEffect, useState} from "react";
import {Button, Col, Drawer, Row} from "antd";
import {useLocation, useNavigate} from "react-router-dom";
import {ShowAddressesAPI} from "../../api/addresses";
import "./style.less"
import {CreatePayAPI} from "../../api/pay";

export default function OrderDetails() {
    const navigateTo = useNavigate()
    const location = useLocation();


    // 地址
    const [address, setAddress] = useState("");
    const [addressData, setAddressData] = useState([]);

    useEffect(() => {
        ShowAddressesAPI().then(res => {
            setAddress(res.data[0]);
        }).catch(err => {
            console.log(err);
        });
    }, []);


    // 展示收货地址信息
    const cart = location.state.Cart
    const WeChatPay = () => {
        CreatePayAPI({
            "ProductID": cart.product_id,
            "Code" : cart.code,
            "UserID": cart.user_id,
            "OrderID": cart.id,
            "QAID": 1,
            "DEID": 1,
            "PostID": 1,
            "Price": cart.price,
            "AddressID": address.id,
            "ECP": 1,
            "PaymentType": 3,
            "Status": 2,
        }).then(res => {
            console.log(res)
            navigateTo(`/layout/personal/order/`)
        }).catch(err => {
            console.log(err)
        })
    }


    // 打开/关闭地址 抽屉
    const [OpenDrawer, setOpenDrawer] = useState(false);
    const DrawerModal = (bool) => {
        ShowAddressesAPI().then(res => {
            setAddressData(res.data)
        }).catch(err => {
            console.log(err);
        });
        setOpenDrawer(bool);
    };


    function selectAddress(data) {
        setAddress(data)
        setOpenDrawer(false)
    }

    return (
        <div className={"body"}>
            <h1>线上支付</h1>
            <div className={"address"} onClick={() => DrawerModal(true)}>
                {address ? (<div className={"addressItem"} key={address.id}>
                    <div className={"itemProfile"}>
                        <span>{address.name}&nbsp;,</span>
                        <span className={"itemTelephone"}>{address.phone}</span>
                        <p>{address.address}</p>
                    </div>
                    <div className={"itemButton"}>
                        <div className={"default"}></div>
                        <div className={"selectButton"}>
                            <Button className={"updateButton"}
                                    type="primary">修改</Button>
                        </div>
                    </div>
                </div>) : <div className={"EmptyCart"}>
                    {/* 此处的图片不能直接写路径，只能通过import的方式将它引入进来 */}
                    {/*<img src={emptyCart} alt="" className={"EmptyCartImg"}/>*/}
                    <div className={"EmptyCartText1"}>购物车竟然是空的！</div>
                    <div className={"EmptyCartText2"}>再忙，也要记得买点什么犒劳自己~</div>
                </div>}
            </div>
            {/* 状态： 1未付款、    2已付款、   3已发货、   4已签收 */}
            {/* 支付方式： 1借记卡、  2信用卡、   3微信、    4支付宝、   5现金 */}
            <div className={"Contain"}>
                <Row>
                    <Col xs={24} sm={24} md={24} lg={24} xl={24}>
                        <div className={"ContainPay"}>
                            <h4>选择支付方式</h4>
                            <h5>平台支付</h5>
                            <div className={"Platform"}>
                                <Button onClick={WeChatPay}>微信支付</Button>
                                <Button onClick={WeChatPay}>支付宝支付</Button>
                                <Button onClick={WeChatPay}>云闪付</Button>
                                <Button onClick={WeChatPay}>PayPal</Button>
                            </div>
                            <h5>银联支付</h5>
                            <div className={"UnionPay"}>
                                <Button onClick={WeChatPay}>江苏银行</Button>
                                <Button onClick={WeChatPay}>中国银行</Button>
                                <Button onClick={WeChatPay}>招商银行</Button>
                                <Button onClick={WeChatPay}>工商银行</Button>
                                <Button onClick={WeChatPay}>建设银行</Button>
                                <Button onClick={WeChatPay}>农业银行</Button>
                            </div>
                        </div>
                    </Col>
                </Row>
            </div>

            {/* 选择抽屉地址栏 */}
            <Drawer
                placement={"bottom"}
                closable={false}
                onClose={() => DrawerModal(false)}
                open={OpenDrawer}
                key={"bottom"}
            >
                <h2>选择地址栏</h2>
                <div className={"address"}>
                    {addressData ? (addressData.map((item) => (
                        <div className={"addressItem"} onClick={() => selectAddress(item)} key={item.id}>
                            <div className={"itemProfile"}>
                                <span>{item.name}&nbsp;,</span>
                                <span className={"itemTelephone"}>{item.phone}</span>
                                <p>{item.address}</p>
                            </div>
                        </div>))) : <div className={"EmptyCart"}>
                        {/* 此处的图片不能直接写路径，只能通过import的方式将它引入进来 */}
                        {/*<img src={emptyCart} alt="" className={"EmptyCartImg"}/>*/}
                        <div className={"EmptyCartText1"}>购物车竟然是空的！</div>
                        <div className={"EmptyCartText2"}>再忙，也要记得买点什么犒劳自己~</div>
                    </div>}
                </div>
            </Drawer>
        </div>
    )
}