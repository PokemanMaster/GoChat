import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Col, Layout, Row } from "antd";
import Search from "antd/es/input/Search";
import { ListProductsParamsAPI, SearchProductsAPI } from "../../../api/products";
import { Link as MuiLink } from "@mui/material";
import Sider from "antd/es/layout/Sider";
import "./style.less";
import Categories from "../../../components/product_sort/categories";

// 商品分类页
export default function Goods() {
    const navigateTo = useNavigate();
    const [ListProducts, setListProducts] = useState([]); // Ensure it starts as an empty array
    const [activeTab, setActiveTab] = useState(0);
    const queryParams = new URLSearchParams(window.location.search);
    const userId = queryParams.get('userId');

    // 初始化商品列表
    useEffect(() => {
        ListProductsParamsAPI().then(res => {
            setListProducts(res.data.items || []); // Ensure it's an array, fallback to empty array
        }).catch(err => {
            console.error(err);
            setListProducts([]); // If API fails, set to empty array
        });
    }, []);

    // 处理分类选择
    const ListCategoriesSelectFunc = (category_id) => {
        setActiveTab(0);
        category_id = parseInt(category_id, 10);
        ListProductsParamsAPI({ CategoryID: category_id }).then(res => {
            setListProducts(res.data.items || []);
        }).catch(err => {
            console.error(err);
            setListProducts([]);
        });
    };

    // 商品详情
    function ShowProduct(value) {
        navigateTo(`/layout/product/${value.id}?userId=${userId}`);
    }

    // 搜索
    function SearchProducts(value) {
        SearchProductsAPI({ "search": value }).then(res => {
            setListProducts(res.data || []); // Ensure it's an array
        }).catch(err => {
            console.log(err);
            setListProducts([]); // If API fails, set to empty array
        });
    }

    return (
        <>
            {/* 搜索 */}
            <Row>
                <Col xs={24} sm={24} md={24} lg={24} xl={24}>
                    <div className="search">
                        <Search
                            placeholder="请输入要搜索的商品名字"
                            enterButton="Search"
                            size="large"
                            onSearch={SearchProducts}
                            maxLength={20}
                        />
                    </div>
                </Col>
            </Row>
            {/* 分类栏 & 轮播图 */}
            <Layout>
                <Sider width={200}>
                    <Categories onSelectCategory={ListCategoriesSelectFunc} />
                </Sider>
                {/*<CarouselsComponent />*/}
                <div className="Products">
                    <Row>
                        {ListProducts && ListProducts.length > 0 ? (
                            ListProducts.map((item) => {
                                const facade = JSON.parse(item.images).facade || [];
                                return (
                                    <Col xs={12} sm={8} md={6} lg={4} xl={3} key={item.id}>
                                        <div onClick={() => ShowProduct(item)} className="Product">
                                            <div className="ProductImage">
                                                {facade.map((url, index) => (
                                                    <img key={index} src={url} alt={""} />
                                                ))}
                                            </div>
                                            <div className="ProductIntroduce">
                                                <div className="ProductIntroduceName">{item.title}</div>
                                                <div className="ProductIntroduceValue">¥{item.price}</div>
                                            </div>
                                        </div>
                                    </Col>
                                );
                            })
                        ) : (
                            <div>No products available</div> // Fallback message when there are no products
                        )}
                    </Row>
                </div>
            </Layout>
            {/* 商品 */}

            {/* 备案号 */}
            <Row style={{ textAlign: 'center' }}>
                <Col xs={24} sm={24} md={24} lg={24} xl={24}>
                    <MuiLink href="https://beian.miit.gov.cn/" underline="none" style={{ color: "#333" }}>
                        桂ICP备2023004200号-2
                    </MuiLink>
                </Col>
            </Row>
        </>
    );
}
