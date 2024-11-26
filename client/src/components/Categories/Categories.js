// src/components/Categories.js
import React, {useEffect, useState} from "react";
import {Menu} from "antd";
import {ListCategoriesAPI} from "../../api/products";

export default function Categories({onSelectCategory}) {
    const [CategoriesParent, setCategoriesParent] = useState([]); // 父类分类栏

    useEffect(() => {
        // 调用API获取商品分类数据
        ListCategoriesAPI().then(res => {
            console.log(res)
            // 筛选父类分类栏
            const filteredCategoriesParent = res.data.filter(item => item.if_parent === true);
            setCategoriesParent(filteredCategoriesParent);
        }).catch(err => {
            console.error(err);
        });
    }, []);

    return (
        <Menu
            mode="vertical"
            style={{
                height: '100%',
                borderRight: 0,
                display: 'flex',
                flexDirection: 'column',
                justifyContent: 'space-around',
                paddingTop: '27px',
                paddingBottom: '27px',
            }}
        >
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
