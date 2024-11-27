// src/components/Categories.js
import React, {useEffect, useState} from "react";
import {Menu} from "antd";
import {ListCategoriesAPI} from "../../api/products";
import "./style.less";  // 引入自定义样式

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
        <Menu className="custom-menu">
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
    );
}
