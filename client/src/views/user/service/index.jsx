import * as React from "react";
import {useEffect, useState} from "react";
import {useNavigate} from "react-router-dom";
import "./style.less"
import {UserTokenAPI} from "../../../api/users";
import {Form, Input, Button, Row, Col, Upload, Empty, Typography} from "antd";
import {
    LockOutlined,
    MailOutlined,
    PhoneOutlined,
    SafetyOutlined,
    UserOutlined
} from "@ant-design/icons";
import cookie from 'react-cookies'

export default function UserService() {
    const queryParams = new URLSearchParams(window.location.search);
    const userId = queryParams.get('userId');
    const navigateTo = useNavigate();
    const [form] = Form.useForm();
    const [UserInfo, setUserInfo] = useState(() => JSON.parse(localStorage.getItem("user")) || {});
    const token = localStorage.getItem("token");

    // Fetch user information on mount
    useEffect(() => {
        UserTokenAPI(token)
            .then(res => {
                if (res.msg === "ok") {
                    const user = JSON.parse(localStorage.getItem("user"));
                    if (user) {
                        setUserInfo(user);
                        form.setFieldsValue({nickname: user.nickname, username: user.user_name});
                    }
                }
            })
            .catch(err => console.log(err));
    }, [token, form]);

    const onFinish = (values) => {
        // Your logic for saving the form
    };

    function Secure(values) {
        cookie.remove("MyCookie", {path: "/"});
        const removedCookie = cookie.load("MyCookie");
        console.log(removedCookie)
        window.location.href = '/login'  // 用户登出跳转到登录页面
    }

    return (
        userId ? (
            <div className="account-security">
                <Typography.Title level={4} className="section-title">登录方式</Typography.Title>
                <Row className="item-row" align="middle">
                    <Col span={2}><PhoneOutlined className="icon"/></Col>
                    <Col span={16}><Typography.Text>安全手机</Typography.Text></Col>
                    <Col span={6} className="right-text"
                         onClick={() => navigateTo("/layout/user/service/telephone")}>点击绑定</Col>
                </Row>
                {/*<Row className="item-row" align="middle">*/}
                {/*    <Col span={2}><MailOutlined className="icon"/></Col>*/}
                {/*    <Col span={16}><Typography.Text>安全邮箱</Typography.Text></Col>*/}
                {/*    <Col span={6} className="right-text">未设置</Col>*/}
                {/*</Row>*/}
                <Row className="item-row" align="middle">
                    <Col span={2}><LockOutlined className="icon"/></Col>
                    <Col span={16}><Typography.Text>修改密码</Typography.Text></Col>
                    <Col span={6} className="right-text"
                         onClick={() => navigateTo("/layout/user/service/password")}>点击重置</Col>
                </Row>
                {/*<Row className="item-row" align="middle">*/}
                {/*    <Col span={2}><UserOutlined className="icon"/></Col>*/}
                {/*    <Col span={16}><Typography.Text>第三方账号</Typography.Text></Col>*/}
                {/*    <Col span={6} className="right-text">已绑定</Col>*/}
                {/*</Row>*/}

                {/*<Typography.Title level={4} className="section-title">账号安全</Typography.Title>*/}
                {/*<Row className="item-row" align="middle">*/}
                {/*    <Col span={2}><SafetyOutlined className="icon"/></Col>*/}
                {/*    <Col span={16}><Typography.Text>密保问题</Typography.Text></Col>*/}
                {/*    <Col span={6} className="right-text">未设置</Col>*/}
                {/*</Row>*/}
                {/*<Row className="item-row" align="middle">*/}
                {/*    <Col span={2}><SafetyOutlined className="icon"/></Col>*/}
                {/*    <Col span={16}><Typography.Text>登录设备管理</Typography.Text></Col>*/}
                {/*    <Col span={6} className="right-text">未设置</Col>*/}
                {/*</Row>*/}
            </div>
        ) : (
            <div className="Empty">
                <Empty
                    image="https://gw.alipayobjects.com/zos/antfincdn/ZHrcdLPrvN/empty.svg"
                    imageStyle={{height: 160}}
                    description={<span>你还没有 <a href="/client/src/public">登录？</a></span>}
                >
                    <Button type="primary" onClick={() => navigateTo("/login")}>
                        点击登录
                    </Button>
                </Empty>
            </div>
        )
    );
}
