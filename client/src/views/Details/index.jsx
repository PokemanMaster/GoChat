import "./style.less"
import {useParams} from "react-router-dom";
import React, {useEffect, useState} from "react";
import {ShowProductParamAPI} from "../../api/products";
import {CreateFavoriteAPI} from "../../api/favorites";
import {CreateCartAPI} from "../../api/carts";
import {Button, Col, Modal, Row} from 'antd';
import {ShowAddressesAPI} from "../../api/addresses";

export default function Details() {
    const user = JSON.parse(localStorage.getItem("user")); // 获取用户数据
    const UserId = user ? user.id : null; // 检查用户数据是否存在
    const {id} = useParams();     // 获取浏览器的URL的:id

    const [ProductParam, setProductParam] = useState([]);
    const [ProductParamImagesDesc, setProductParamImagesDesc] = useState([]);
    const [ProductParamImagesFacade, setProductParamImagesFacade] = useState([]);

    const [Address, setAddress] = useState([]);
    const [Message, setMessage] = useState("");

    useEffect(() => {
        ShowProductParamAPI(id).then(res => {
            setProductParam(res.data)
            const images = JSON.parse(res.data[0].images);
            setProductParamImagesDesc(images.desc) // 存储展示图片
            setProductParamImagesFacade(images.facade) // 存储大图片
        }).catch(error => {
            console.error("Error in useEffect:", error);
        })

        ShowAddressesAPI().then((res) => {
            console.log(res.data[0]);
            setAddress(res.data[0]);
        }).catch(error => {
            console.log(error);
        })


    }, [id])

    // 切换图片
    const ShowParamImgs = (paramImg) => {
        console.log(paramImg)
        setProductParamImagesFacade(paramImg);
    }



    // 点击添加到购物车
    function CreateCart() {
        CreateCartAPI({"UserID": UserId, "ProductID": JSON.parse(id)}).then(res => {
            setMessage(res.msg)
            showModal()
        }).catch(error => {
            console.log(error)
        })
    }

    // 点击收藏
    function CreateFavorite() {
        CreateFavoriteAPI({"UserID": UserId, "ProductID": JSON.parse(id)}).then(res => {
            setMessage(res.msg)
            showModal()
        }).catch(error => {
            console.log(error)
        })
    }

    // modal对话框
    const [isModalOpen, setIsModalOpen] = useState(false);
    const showModal = () => {
        setIsModalOpen(true);
    };
    const handleOk = () => {
        setIsModalOpen(false);
    };
    const handleCancel = () => {
        setIsModalOpen(false);
    };

    return (
        <>

            <div className="product-page">
                <div className="product-container">
                    <div className="product-image">
                        <img src={ProductParamImagesFacade} alt="产品图片"/>
                        <div className="image-thumbnails">
                            {ProductParamImagesDesc && ProductParamImagesDesc.length > 0 ? (ProductParamImagesDesc.map((item, index) => (
                                <img
                                    key={index}
                                    src={item}
                                    alt={`缩略图 ${index + 1}`}
                                    className={ProductParamImagesFacade === item ? 'active' : ''}
                                    onClick={() => setProductParamImagesFacade(item)}
                                    onMouseEnter={() => {
                                        ShowParamImgs(item); // 这里的 `item` 是图片的 URL
                                    }}
                                />
                            ))) : (<div></div>)}
                        </div>
                    </div>
                    {ProductParam && ProductParam.length > 0 ? (ProductParam.map((item, index) => (
                        <div className="product-details">
                            <h1 className="product-title">{item.title}</h1>
                            <p className="product-subtitle">高质量 | 高品质 | 二合一结构 | 安全安心</p>
                            <div className="product-price">
                                <span className="price-current">{item.price} 元</span>
                                <span className="price-old">{item.price + 100} 元</span>
                            </div>
                            <div className="product-stock">
                                <p>{Address.address}<span className="modify">修改</span></p>
                                <p className="stock-status">有现货</p>
                            </div>
                            <div className="product-actions">
                                <Button className="btn-cart" type="primary" onClick={CreateCart}>加入购物车</Button>
                                <Button className="btn-like" type="default" onClick={CreateFavorite}>喜欢</Button>
                            </div>
                        </div>
                    ))) : (<div></div>)}
                </div>
            </div>

            {ProductParam ?
                <div key={ProductParam.id}>
                    {/*商品详情介绍*/}
                    <div className={"header"}>
                        <div className={"detail"}>
                            <Row className={"Row"}>
                                <Col xs={24} sm={24} md={24} lg={24} xl={12}>

                                </Col>
                            </Row>
                        </div>
                    </div>
                </div> : null}

            <Modal open={isModalOpen} onOk={handleOk} onCancel={handleCancel}>
                <p>{Message}</p>
            </Modal>
        </>
    )
}