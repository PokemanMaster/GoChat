import React, {useEffect, useState} from "react";
import {useNavigate} from "react-router-dom";
import {Col, Layout, Menu, Row} from "antd";
import Search from "antd/es/input/Search";
import {ListProductsParamsAPI, SearchProductsAPI} from "../../../api/products";
import {Link as MuiLink} from "@mui/material";
import Sider from "antd/es/layout/Sider";
import "./style.less";
import Categories from "../../../components/product_sort/categories";
import CarouselsComponent from "../../../components/product_carousels/carousels";

// 商品分类页
export default function Goods() {
    const navigateTo = useNavigate();
    const [ListProducts, setListProducts] = useState([]); // Ensure it starts as an empty array
    const [activeTab, setActiveTab] = useState(1);
    const queryParams = new URLSearchParams(window.location.search);
    const userId = queryParams.get('userId');

    // 初始化商品列表
    useEffect(() => {
        ListProductsParamsAPI({"Limit": 10, "Start": 0, "CategoryID": 1}).then(res => {
            console.log(res.data.items)
            setListProducts(res.data.items || []);
        }).catch(err => {
            console.error(err);
            setListProducts([]);
        });
    }, []);

    // 处理分类选择
    const ListCategoriesSelectFunc = (category_id) => {
        setActiveTab(0);
        category_id = parseInt(category_id, 10);
        ListProductsParamsAPI({CategoryID: category_id}).then(res => {
            setListProducts(res.data.items || []);
        }).catch(err => {
            console.error(err);
            setListProducts([]);
        });
    };

    // 商品详情
    function ShowProduct(value) {
        navigateTo(`/layout/product/details/${value.id}?userId=${userId}`, {
            state: {
                item : value
            }
        });
    }

    // 搜索
    function SearchProducts(value) {
        SearchProductsAPI({"search": value}).then(res => {
            setListProducts(res.data || []); // Ensure it's an array
        }).catch(err => {
            console.log(err);
            setListProducts([]); // If API fails, set to empty array
        });
    }

    return (
        <div className={"gochat-sort-body"}>
            <div className={"gochat-sort"}>
                {/*商品头部*/}
                <div className={"gochat-sort-header"}>
                    <img className={"header-logo"} src="/logo/logo.png" alt=""/>
                    <div className={"header-search"}>
                        <Search
                            placeholder="请输入要搜索的商品名字"
                            enterButton="Search"
                            size="large"
                            onSearch={SearchProducts}
                            maxLength={20}
                        />
                    </div>
                </div>

                {/* 分类栏 & 轮播图 */}
                <Layout className={"gochat-sort-layout"}>
                    <Categories onSelectCategory={ListCategoriesSelectFunc}/>
                    <CarouselsComponent/>
                </Layout>

                {/* 商品列表 */}
                <div className={"gochat-sort-products"}>
                    {ListProducts && ListProducts.length > 0 ? (
                        ListProducts.map((item) => {
                            return (
                                <div onClick={() => ShowProduct(item)} className="products-layout" key={item.id}>
                                    <div className="products-layout-image">
                                        <img src={item.image} alt={""}/>
                                    </div>
                                    <div className="products-layout-introduce">
                                        <div className="introduce-name">{item.name}</div>
                                        <div className={"introduce-value"}>
                                            <div className="introduce-price">¥{item.price}</div>
                                            <div className="introduce-count">已售{item.sold_count}+</div>
                                        </div>
                                    </div>
                                </div>
                            );
                        })
                    ) : (
                        <div>No products available</div> // Fallback message when there are no products
                    )}
                </div>

                {/* 备案号 */}
                <Row style={{textAlign: 'center'}}>
                    <Col xs={24} sm={24} md={24} lg={24} xl={24}>
                        <MuiLink href="https://beian.miit.gov.cn/" underline="none" style={{color: "#333"}}>
                            桂ICP备2023004200号-2
                        </MuiLink>
                    </Col>
                </Row>
            </div>
        </div>
    );
}
