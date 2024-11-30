import * as React from "react";
import {useEffect, useState} from "react";
import {useNavigate} from "react-router-dom";
import "./style.less";
import {CategoryAPI, UpdatePasswordAPI} from "../../../api/users";
import {Form, Input, Button, Modal, Typography} from "antd";

export default function UserService() {
    const navigateTo = useNavigate();
    const [form] = Form.useForm();
    const token = localStorage.getItem("token");
    const [UserInfo, setUserInfo] = useState(() => JSON.parse(localStorage.getItem("user")) || {});

    // 初始化验证码图片
    const [codeId, setCodeId] = useState("");
    const [base64, setBase64] = useState("");
    useEffect(() => {
        CategoryAPI()
            .then((res) => {
                setCodeId(res.data["code_id"]);
                setBase64(res.data["base64"]);
            })
            .catch((err) => {
                console.log("CategoryAPI:", err);
            });
    }, []);

    // 点击切换验证码图片
    const changeCategory = () => {
        CategoryAPI()
            .then((res) => {
                setCodeId(res.data["code_id"]);
                setBase64(res.data["base64"]);
            })
            .catch((err) => {
                console.log(err);
            });
    };

    // 密码处理
    const [password, setPassword] = useState("");
    const handlePasswordChange = (event) => {
        setPassword(event.target.value);
    };

    // 验证码处理
    const [code, setCode] = useState("");
    const handleCodeChange = (event) => {
        setCode(event.target.value);
    };

    // 弹出层处理
    const [message, setMessage] = useState("");
    const [open, setOpen] = useState(false);
    const handleOpen = (msg) => {
        setOpen(true);
        setMessage(msg);
    };
    const handleClose = () => setOpen(false);

    const handleSubmit = () => {
        UpdatePasswordAPI({
            "ID": UserInfo.id,
            "Password": password,
            "Code": code,
            "CodeId": codeId,  // codeId 是在 useEffect 中设置的
        }).then(res => {
            handleOpen(res.msg);
            console.log(res);
        }).catch(err => {
            console.log(err)
        });
    };


    return (
        <div className="reset-password-page">
            <div className="logo">
                <img
                    src={
                        "https://p3-sign.bdxiguaimg.com/tos-cn-i-0026/oYIgAWbvsAAiwmAaZtHZqBzh7giEMLAIsClhi~tplv-pk90l89vgd-crop-center:864:486.jpeg?appId=0&channelId=0&customType=custom%2Fnone&from=0_large_image_list&imageType=video1609&isImmersiveScene=0&is_stream=0&lk3s=9d3f5bff&logId=202408252157412FF2ACF086057837939F&requestFrom=0&x-expires=1756130261&x-signature=Es5SF30gsHnZ8ZWFLAtsN2Obnl8%3D"
                    }
                    alt="Xiaomi Logo"
                />
            </div>
            <h2>重置密码</h2>
            <Form
                name="reset_password"
                onFinish={handleSubmit}
                layout="vertical"
                className="reset-form"
            >
                <Form.Item
                    name="password"
                    label="请输入新密码："
                    rules={[{required: true, message: "请输入新密码"}]}
                >
                    <Input.Password
                        placeholder="新密码"
                        value={password}
                        onChange={handlePasswordChange}
                    />
                </Form.Item>

                <Form.Item
                    name="captcha"
                    label="图片验证码"
                    rules={[{required: true, message: "请输入图片验证码"}]}
                >
                    <div className="captcha-container">
                        <Input placeholder="图片验证码" value={code} onChange={handleCodeChange}/>
                        <img
                            src={base64}
                            onClick={changeCategory}
                            alt="Captcha"
                            className="captcha-image"
                        />
                    </div>
                </Form.Item>

                <Form.Item>
                    <Button type="primary" htmlType="submit" className="submit-button">
                        修改
                    </Button>
                </Form.Item>
            </Form>

            {/*弹出层*/}
            <Modal title="提示" visible={open} onCancel={handleClose} footer={null}>
                <Typography>{message}</Typography>
            </Modal>
        </div>
    );
}
