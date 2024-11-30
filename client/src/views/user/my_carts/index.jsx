import React, {useEffect, useState} from "react";
import {Button, Col, Modal, Row, Empty} from "antd";
import {DeleteCartAPI, ShowCartAPI, UpdateCartAPI} from "../../../api/carts";
import {useLocation, useNavigate} from "react-router-dom";
import {Link} from "@mui/joy";
import {DeleteOutlined} from "@ant-design/icons";
import "./style.less"
import {CreateOrderAPI} from "../../../api/orders";
import NotLoginComponent from "../../../components/not_login/not_login";
import NotDataComponent from "../../../components/not_data/not_data";

export default function Cart() {
    const navigateTo = useNavigate() // 路由
    const location = useLocation();
    const queryParams = new URLSearchParams(location.search);
    const userId = queryParams.get('userId') || '';
    const [cart, setCart] = useState([]); // 存放商品
    const [totalPrice, setTotalPrice] = useState(0); // 总价
    const [selectAll, setSelectAll] = useState(false); // 全选按钮

    // 展示购物车
    useEffect(() => {
        ShowCartAPI()
            .then(res => {
                setCart(res.data.items);
            }).catch(error => {
            console.log(error)
        });
    }, []);

    // 标记勾选状态
    function ToggleChecked(index) {
        const newShopCar = [...cart];
        newShopCar[index].check = !newShopCar[index].check;
        const allSelected = newShopCar.every(item => item.check); // 更新 selectAll 状态
        setSelectAll(allSelected);
        setCart(newShopCar); // 更新购物车状态
        const totalPrice = newShopCar.reduce((acc, item) => { // 计算总价
            if (item.check && item.num && item.price) {
                return acc + parseFloat(item.price) * item.num;
            }
            return acc;
        }, 0);
        setTotalPrice(parseFloat(totalPrice.toFixed(2)));
    }


    // 全选按钮
    function allClick() {
        const allSelected = !selectAll;
        const totalPrice = cart.reduce((acc, item) => {
            if (item.check !== allSelected) {
                item.check = allSelected;
            }
            if (item.check) {
                return acc + parseFloat(item.price) * item.num;
            }
            return acc;
        }, 0);
        setCart([...cart]); // 强制更新 Index 数组，以触发重新渲染
        setSelectAll(allSelected);
        setTotalPrice(parseFloat(totalPrice.toFixed(2)));
    }


    // 创建订单
    const ShopEverything = () => {
        const checkedCart = cart.filter(item => item.check === true);
        if (checkedCart != null && checkedCart.length > 0) {
            // 提交订单逻辑
            CreateOrderAPI({
                "ProductID": checkedCart[0].product_id,
                "Type": 2,
                "ShopID": 1,
                "UserID": parseInt(userId , 10),
                "Amount": checkedCart[0].price * checkedCart[0].num,
                "Postage": 0,
                "Weight": 1,
                "Price": checkedCart[0].price,
                "ActualPrice": checkedCart[0].price,
                "Num": checkedCart[0].num,
            }).then(res => {
                console.log(res)
                navigateTo(`/layout/my/orders?userId=${userId}`)
            }).catch(err => {
                console.log(err)
            })
        } else {
            alert("请选择要购买的商品！")
        }
    }


    // 选择数量按钮
    const decrement = (item) => {
        if (item.num > 1) {
            const newCart = cart.map(cartItem => {
                if (cartItem.product_id === item.product_id) {
                    return {...cartItem, num: cartItem.num - 1};
                }
                return cartItem;
            });
            setCart(newCart);

            UpdateCartAPI({
                "UserID": item.user_id, "ProductID": item.product_id, "Num": item.num - 1
            }).then(res => {
                console.log(res);
                updateTotalPrice(newCart); // 更新总金额
            });
        }
    };

    // 添加商品数量
    const increment = (item) => {
        if (item.num < 99) {
            const newCart = cart.map(cartItem => {
                if (cartItem.product_id === item.product_id) {
                    return {...cartItem, num: cartItem.num + 1};
                }
                return cartItem;
            });
            setCart(newCart);
            UpdateCartAPI({
                "UserID": item.user_id, "ProductID": item.product_id, "Num": item.num + 1
            }).then(res => {
                console.log(res);
                updateTotalPrice(newCart); // 更新总金额
            });
        }
    };

    // 更新总金额函数
    const updateTotalPrice = (cartItems) => {
        const totalPrice = cartItems.reduce((acc, item) => {
            if (item.check && item.num && item.price) {
                return acc + parseFloat(item.price) * item.num;
            }
            return acc;
        }, 0);
        setTotalPrice(parseFloat(totalPrice.toFixed(2)));
    };


    // 对话框
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [TmpID, SetTmpID] = useState("")
    const showModal = (product_id) => {
        console.log(product_id)
        SetTmpID(product_id)
        setIsModalOpen(true);
    };
    const handleOk = () => {
        // 删除购物车操作
        DeleteCartAPI({"UserID": userId, "ProductID": TmpID})
            .then(res => {
                console.log(res)
            }).catch(err => {
            console.log(err)
        })
        setIsModalOpen(false);
    };
    const handleCancel = () => {
        setIsModalOpen(false);
    };

    // 解析 JSON 格式图片
    function JsonParseFacade(value) {
        const parsedValue = JSON.parse(value); // 解析 JSON 字符串
        return parsedValue.facade
    }


    return (<>
        {userId ? <div className={"Body"}>
            {/*头部*/}
            <div className={"TopHeader"}>
                <div className={"CartHeader"}>
                    <div className={"Logo"}>
                        <Link to="/">
                        </Link>
                    </div>
                    <div className={"CartHeaderContent"}>
                        <p>我的购物车</p>
                    </div>
                </div>
            </div>

            {/*商品*/}
            <div className={"cartContain"}>
                {/* 购物车表头开始 */}
                <Row>
                    <Col xs={24} sm={24} md={24} lg={24} xl={24}>
                        <div className={"ContentHeader"}>
                            <Col xs={6} sm={6} md={5} lg={5} xl={5} className={"ProCheck"}>
                                <div className={"checkBtn"} onClick={allClick}
                                     style={{backgroundColor: selectAll ? "red" : "white"}}></div>
                                <div>全选</div>
                            </Col>
                            <Col xs={4} sm={4} md={6} lg={4} xl={5}>
                                <div className={"ProName"}>商品名称</div>
                            </Col>
                            <Col xs={5} sm={5} md={4} lg={4} xl={5}>
                                <div className={"ProNum"}>数量</div>
                            </Col>
                            <Col xs={4} sm={4} md={4} lg={4} xl={5}>
                                <div className={"ProTotal"}>价格</div>
                            </Col>
                            <Col xs={2} sm={2} md={5} lg={4} xl={4}>
                                <div className={"ProAction"}>操作</div>
                            </Col>
                        </div>
                    </Col>
                </Row>
                {cart && cart.length > 0 ? (cart.map((item, index) => (<div key={index}>
                    {/* 购物车表头结束 */}
                    <Col xs={24} sm={24} md={24} lg={24} xl={24} key={item.id}>
                        <div className={"CartItem"}>
                            {/* 选中按钮 */}
                            <Col xs={6} sm={6} md={5} lg={6} xl={5} className={"CartInfo"}>
                                <div
                                    style={{backgroundColor: item.check ? "red" : "white"}}
                                    className={"CarItemCheck"}
                                    onClick={() => ToggleChecked(index)}
                                ></div>
                                <div className={"CarItemPicture"}>
                                    <img src={JsonParseFacade(item.images)} alt=""/>
                                </div>
                            </Col>
                            {/* 商品名称 */}
                            <Col xs={4} sm={4} md={5} lg={4} xl={5}>
                                <div className={"CarItemIntroduce"}>
                                    <span>{item.title}</span>
                                </div>
                            </Col>
                            {/* 数量 */}
                            <Col xs={5} sm={5} md={3} lg={5} xl={5}>
                                <div className={"CarItemNum"}>
                                    <button onClick={() => {
                                        decrement(item)
                                    }}>-
                                    </button>
                                    <span>{item.num}</span>
                                    <button onClick={() => increment(item)}>+</button>
                                </div>
                            </Col>
                            {/* 价格 */}
                            <Col xs={5} sm={5} md={3} lg={5} xl={5}>
                                <div className={"CarItemMoney"}>{item.price}</div>
                            </Col>
                            {/* 操作 */}
                            <Col xs={1} sm={1} md={5} lg={1} xl={4} className={"CarItemButton"}>
                                <DeleteOutlined onClick={() => showModal(item.product_id)}/>
                            </Col>
                        </div>
                    </Col>
                </div>))) : <NotDataComponent text="你还没有数据" />}
                {/*提交*/}
                <div className={"shopCarSubmit"}>
                    <div className={"tap"}>
                        <span>继续购物</span>
                        <span>已经选择 <b>1</b> 件</span>
                    </div>
                    <div className={"settle"}>
                        <div className={"totalMoney"}>
                            <span>总共:</span><span>{totalPrice.toFixed(2)}</span></div>
                        <Button type="primary" className={"shopCarSubmitButton"}
                                onClick={ShopEverything}>
                            <span>购买商品</span>
                        </Button>
                    </div>
                </div>
            </div>

            <Modal title="Basic Modal" open={isModalOpen} onOk={handleOk} onCancel={handleCancel}>
                <p>确定要删除吗?</p>
            </Modal>
        </div> :<NotLoginComponent/>
        }
    </>)
}