import React, {useEffect, useState} from 'react';
import {ShowFavoritesAPI} from "../../api/favorites";
import "./style.less"
import {Button, Col, Empty} from "antd";
import {Link} from "@mui/joy";
import {useNavigate} from "react-router-dom";

export default function Favorite() {

    const [Favorites, setFavorites] = useState([])
    const navigateTo = useNavigate();
    const [UserInfo] = useState(JSON.parse(localStorage.getItem("user")));

    useEffect(() => {
        if (UserInfo) {
            ShowFavoritesAPI().then(res => {
                setFavorites(res.data.items)
            }).catch(err => {
                console.log(err)
            });
        }
    }, [UserInfo]);

    // 商品详情 ShowProduct
    function ShowProduct(value) {
        navigateTo(`/layout/product/${value.id}`);
    }

    // 解析 JSON 格式图片
    function JsonParseFacade(value) {
        const parsedValue = JSON.parse(value); // 解析 JSON 字符串
        return parsedValue.facade
    }

    return (
        <>
            {UserInfo ? <div className={"Body"}>
                    {/* 头部 */}
                    <div className={"TopHeader"}>
                        <div className={"CartHeader"}>
                            <div className={"Logo"}>
                                <Link to="/">
                                </Link>
                            </div>
                            <div className={"CartHeaderContent"}>
                                <p>我的收藏夹</p>
                            </div>
                        </div>
                    </div>
                    {/*收藏*/}
                    <div className={"Favorite"}>
                        <div className={"Products"}>
                            {Favorites && Favorites.length > 0 ? (
                                Favorites.map((item, index) => (
                                    <Col xs={12} sm={8} md={6} lg={4} xl={4} key={item.id}>
                                        <div
                                            onClick={() => ShowProduct(item)}
                                            className={"Product"}
                                            key={index}
                                        >
                                            <div className={"ProductImage"}>
                                                <img src={JsonParseFacade(item.images)} alt={"loading..."}/>
                                            </div>
                                            <div className={"ProductIntroduce"}>
                                                <div className={"ProductIntroduceName"}>{item.title}</div>
                                                <div
                                                    className={"ProductIntroduceValue"}>¥{item.price}</div>
                                            </div>
                                        </div>
                                    </Col>
                                ))
                            ) : (
                                <div className={"EmptyCart"}>
                                    {/* 此处的图片不能直接写路径，只能通过import的方式将它引入进来 */}
                                    {/*<img src={emptyCart} alt="" className={"EmptyCartImg"}/>*/}
                                    <div className={"EmptyCartText1"}>收藏竟然是空的！</div>
                                    <div className={"EmptyCartText2"}>赶紧收藏一下把~</div>
                                </div>
                            )}
                        </div>
                    </div>
                </div> :
                // 用户没登录就显示
                <div className={"Empty"}>
                    <Empty
                        image="https://gw.alipayobjects.com/zos/antfincdn/ZHrcdLPrvN/empty.svg"
                        imageStyle={{
                            height: 160,
                        }}
                        description={<span>你还没有 <a href="/">登录？</a></span>}>
                        <Button type="primary" onClick={() => {
                            navigateTo("/login")
                        }}>点击登录</Button>
                    </Empty>
                </div>}
        </>
    );
}
