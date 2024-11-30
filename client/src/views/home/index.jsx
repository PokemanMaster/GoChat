import React from "react";
import './style.less';
import { useNavigate } from "react-router-dom";
import { useUserId } from "../user";  // 导入自定义 Hook

// 商品的首页
export default function Home() {
    const userId = useUserId();  // 使用自定义 Hook 获取 userId
    console.log(userId);
    const navigateTo = useNavigate(); // 路由

    return (
        <div className="home">
            <img src="/picture/home.jpg" alt="Home" />
            <div className="button-container">
                <button className="button" onClick={() => navigateTo(`/layout/products/sort?userId=${userId}`)}>
                    挑选商品
                </button>
            </div>
        </div>
    );
}
