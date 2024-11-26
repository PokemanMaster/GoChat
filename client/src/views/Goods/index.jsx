import React, {useEffect, useState} from "react";
import {useNavigate} from "react-router-dom";
import {Col, Layout, Row} from "antd";
import Search from "antd/es/input/Search";
import {ListProductsParamsAPI, SearchProductsAPI} from "../../api/products";
import {Link as MuiLink} from "@mui/material";
import Sider from "antd/es/layout/Sider";
import CarouselsComponent from "../../components/Carousels/Carousels";
import "./style.less";
import Categories from "../../components/Categories/Categories";

// 商品分类页
export default function Goods() {
    const navigateTo = useNavigate();
    const [ListProducts, setListProducts] = useState([]);
    const [activeTab, setActiveTab] = useState(0);

    // 初始化商品列表
    useEffect(() => {
        ListProductsParamsAPI().then(res => {
            setListProducts(res.data.items);
        }).catch(err => {
            console.error(err);
        });
    }, []);

    // 处理分类选择
    const ListCategoriesSelectFunc = (category_id) => {
        setActiveTab(0);
        category_id = parseInt(category_id, 10);
        ListProductsParamsAPI({CategoryID: category_id}).then(res => {
            setListProducts(res.data.items);
        }).catch(err => {
            console.error(err);
        });
    };

    // 商品详情
    function ShowProduct(value) {
        navigateTo(`/layout/product/${value.id}`);
    }

    // 搜索
    function SearchProducts(value) {
        SearchProductsAPI({"search": value}).then(res => {
            setListProducts(res.data);
        }).catch(err => {
            console.log(err);
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
                    <Categories onSelectCategory={ListCategoriesSelectFunc}/>
                </Sider>
                <CarouselsComponent/>
            </Layout>
            {/* 商品 */}
            <div className="Products">
                <Row>
                    {ListProducts ? ListProducts.map((item) => {
                        const facade = JSON.parse(item.images).facade || [];
                        return (
                            <Col xs={12} sm={8} md={6} lg={4} xl={3} key={item.id}>
                                <div onClick={() => ShowProduct(item)} className="Product">
                                    <div className="ProductImage">
                                        {facade.map((url, index) => (
                                            <img key={index} src={url} alt={""}/>
                                        ))}
                                    </div>
                                    <div className="ProductIntroduce">
                                        <div className="ProductIntroduceName">{item.title}</div>
                                        <div className="ProductIntroduceValue">¥{item.price}</div>
                                    </div>
                                </div>
                            </Col>
                        );
                    }) : <div></div>}
                </Row>
            </div>
            {/* 备案号 */}
            <Row style={{textAlign: 'center'}}>
                <Col xs={24} sm={24} md={24} lg={24} xl={24}>
                    <MuiLink href="https://beian.miit.gov.cn/" underline="none" style={{color: "#333"}}>
                        桂ICP备2023004200号-2
                    </MuiLink>
                </Col>
            </Row>
        </>
    );
}
