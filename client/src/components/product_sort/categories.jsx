// src/components/product_sort.jsx
import React, {useEffect, useState} from "react";
import {Menu} from "antd";
import {ListCategoriesAPI} from "../../api/products";
import "./style.less";
import Sider from "antd/es/layout/Sider";

export default function Categories({onSelectCategory}) {
    const [CategoriesParent, setCategoriesParent] = useState([]); // 父类分类栏

    useEffect(() => {
        // 调用API获取商品分类数据
        ListCategoriesAPI().then(res => {
            const filteredCategoriesParent = res.data.filter(item => item.if_parent === true);
            setCategoriesParent(filteredCategoriesParent);
        }).catch(err => {
            console.error(err);
        });
    }, []);

    return (
        <Sider width={200} >
            <Menu  className={"gochat-categories-sider"}>
                {CategoriesParent && CategoriesParent.length > 0 ? (
                    CategoriesParent.map((item) => (
                        <Menu.Item key={item.id} onClick={() => onSelectCategory(item.id)}>
                            <span style={{paddingRight: '10px'}}>{item.name}</span>
                            <span style={{float: 'right'}}>></span>
                        </Menu.Item>
                    ))
                ) : (
                    <div></div>
                )}
            </Menu>
        </Sider>
    );
}
