import "./style.less"
import {useLocation, useParams} from "react-router-dom";
import React, {useEffect, useState} from "react";
import {ShowProductParamAPI} from "../../../api/products";
import {CreateCartAPI} from "../../../api/carts";
import {Button, Col, Modal, Row} from 'antd';
import {ShowAddressesAPI} from "../../../api/addresses";

export default function Details() {
    const {id} = useParams();     // 获取浏览器的URL的:id
    const location = useLocation();
    const {item} = location.state;
    const queryParams = new URLSearchParams(location.search);
    const userId = queryParams.get('userId');

    const [Product, setProduct] = useState([]); // 商品
    const [ProductParam, setProductParam] = useState([]); // 商品参数
    const [ProductParamImage, setProductParamImage] = useState([]); // 主图 + 所有次图
    const [ProductParamMainImage, setProductParamMainImage] = useState(""); // 主图
    const [ProductParamInfo, setProductParamInfo] = useState([]); // 款式信息


    const [Address, setAddress] = useState([]);
    const [Message, setMessage] = useState("");

    useEffect(() => {
        setProduct(item)
        ShowProductParamAPI(id).then(res => {
            const productParamImages = res.data.map(param => param.Image);
            const combinedImages = [item.image, ...productParamImages];
            setProductParam(res.data);
            setProductParamMainImage(item.image)
            setProductParamImage(combinedImages);
            setProductParamInfo(res.data[0])
        }).catch(error => {
            console.error("Error in useEffect:", error);
        });

        ShowAddressesAPI().then((res) => {
            setAddress(res.data[0]);
        }).catch(error => {
            console.log(error);
        })
    }, [id])

    // 切换图片
    const ShowParamImgs = (Image) => {
        setProductParamMainImage(Image);
    }

    // 选择不同款式
    function SelectSize(item) {
        setProductParamInfo(item)
    }


    // 点击添加到购物车
    function CreateCart(item) {
        CreateCartAPI({"UserID": parseInt(userId, 10), "ProductID": item.ID}).then(res => {
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
        <div>
            <div className="product-details">
                {/*图片*/}
                <div className="product-details-image">
                    <img src={ProductParamMainImage || ""} alt=""/>
                    <div className="image-thumbnails">
                        {ProductParamImage && ProductParamImage.length > 0 ? (ProductParamImage.map((item, index) => (
                            <img
                                key={index}
                                src={item}
                                alt={""}
                                onMouseEnter={() => {
                                    ShowParamImgs(item);
                                }}
                            />
                        ))) : (<div></div>)}
                    </div>
                </div>

                {/*介绍*/}
                <div className={"product-details-info"}>
                    <h1 className="info-title">{item.name}</h1>
                    <p className="info-subtitle">高质量 | 高品质 | 二合一结构 | 安全安心</p>
                    <div className={"info-select"}>
                        选择款式
                        {ProductParam && ProductParam.length > 0 ? (ProductParam.map((item, index) => (
                            <div key={index} className={"info-select-box"} onClick={() => SelectSize(item)}>
                                {item.Size || ""}
                            </div>
                        ))) : (<div></div>)}
                    </div>
                    <div className="info-price">
                        <span className="price-current">{ProductParamInfo.DiscountPrice} 元</span>
                        <span className="price-old">{ProductParamInfo.Price} 元</span>
                    </div>
                    <div className="info-stock">
                        <p>{Address.address}<span className="stock-modify">修改</span></p>
                        <p className="stock-status">有现货</p>
                    </div>
                    <div className="info-actions">
                        <Button className="actions-cart" type="primary" onClick={()=>CreateCart(ProductParamInfo)}>加入购物车</Button>
                    </div>
                </div>
            </div>
            {/*弹出层*/}
            <Modal open={isModalOpen} onOk={handleOk} onCancel={handleCancel}>
                <p>{Message}</p>
            </Modal>
        </div>
    )
}