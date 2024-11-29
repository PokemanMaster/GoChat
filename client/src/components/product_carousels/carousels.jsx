// 日历组件
import React, {useEffect, useState} from 'react';
import {ListCarouselsAPI} from "../../api/carousels";
import {Carousel, Layout} from "antd";
import {Content} from "antd/es/layout/layout";

const CarouselsComponent = () => {
    // 轮播图 CarouselImages
    const [CarouselImages, setCarouselImages] = useState([])
    useEffect(() => {
        ListCarouselsAPI().then(res => {
            setCarouselImages(res.data);
        }).catch(err => {
            console.error(err);
        });
    }, []);

    return (
        <Layout>
            {/*<div>*/}
            {/*    <Content*/}
            {/*        style={{*/}
            {/*            margin: 0, minHeight: 280,*/}
            {/*        }}*/}
            {/*    >*/}
            {/*        <Carousel autoplay>*/}
            {/*            {CarouselImages.map((item, index) => (<div key={index}>*/}
            {/*                <img src={item.img_path} alt={""} style={{width: '100%', height: 'auto'}}/>*/}
            {/*            </div>))}*/}
            {/*        </Carousel>*/}
            {/*    </Content>*/}
            {/*</div>*/}
        </Layout>
    );
};

export default CarouselsComponent;
