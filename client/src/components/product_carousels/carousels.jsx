// 日历组件
import React, {useEffect, useState} from 'react';
import {ListCarouselsAPI} from "../../api/carousels";
import {Carousel, Layout} from "antd";
import "./style.less";
import {Content} from "antd/es/layout/layout";

const CarouselsComponent = () => {
    // 轮播图 CarouselImages
    const [carousel, setCarousel] = useState([]);

    useEffect(() => {
        ListCarouselsAPI()
            .then(res => {
                setCarousel(res.data);
                console.log(res.data);
            })
            .catch(err => {
                console.error(err);
            });
    }, []);

    return (
        <Layout>
            <Content>
                {/* 轮播图 */}
                <div className={"gochat-carousels"}>
                    <Carousel autoplay>
                        {carousel.map(item => (
                            <div key={item.id} className={"gochat-carousels-images"}>
                                <img src={item.img_path} alt=""/>
                            </div>
                        ))}
                    </Carousel>
                </div>
            </Content>
        </Layout>
    );
};


export default CarouselsComponent;
