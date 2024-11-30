import "./style.less"
import React, {useEffect, useState} from "react";
import {UpdateAPI} from "../../../api/users";
import {Button, Form, Input, Upload, message, Empty} from "antd";
import {UploadOutlined} from "@ant-design/icons";
import {useNavigate} from "react-router-dom";
import {Typography} from "antd";

export default function UserDetails() {
    const queryParams = new URLSearchParams(window.location.search);
    const userId = queryParams.get('userId');
    const [UserInfo, setUserInfo] = useState(() => JSON.parse(localStorage.getItem("user" + userId)) || {});
    const [form] = Form.useForm();
    const navigateTo = useNavigate();
    const {Text, Title} = Typography;


    // 展示用户信息
    useEffect(() => {
        const user = JSON.parse(localStorage.getItem("user"));
        if (userId) {
            setUserInfo(user + userId);
            form.setFieldsValue({nickname: UserInfo.nickname, username: UserInfo.user_name});
        }
    }, [form]);

    // 用户信息存储
    const [avatar, setAvatar] = useState(UserInfo.avatar || '');
    const [nickName, setNickName] = useState(UserInfo.nickname || '');
    const [userName, setUserName] = useState(UserInfo.user_name || '');

    // 提交修改
    const onFinish = (values) => {
        UpdateAPI({"id": UserInfo.id, "nickName": nickName, "user_name": userName, "avatar": avatar})
            .then(res => {
                success();
            }).catch(err => {
            console.log(err);
        });
    };

    // 修改头像
    const UpdateAvatar = ({file}) => {
        console.log(file)
        setAvatar(file.name);
    };

    // 提示消息
    const [messageApi, contextHolder] = message.useMessage();
    const success = () => {
        messageApi.open({
            type: 'success', content: '修改成功',
        }).then();
    };

    return (<>
        {userId  ? (<div className={"my_account"}>
            <div className={"UserDetailsContent"}>
                <div className={"Extra"}></div>
                <div className={"UserDetailsTitle"}>
                    <p>个人信息</p>
                </div>
                {contextHolder}
                <div className={"UserDetailsForm"}>
                    <Form
                        form={form}
                        name="user-details-form"
                        labelCol={{span: 8}}
                        wrapperCol={{span: 16}}
                        initialValues={{remember: true}}
                        onFinish={onFinish}
                        autoComplete="off"
                    >
                        <Form.Item label="头像:">
                            <Upload
                                listType="picture-card"
                                customRequest={UpdateAvatar}
                                showUploadList={false}
                            >
                                {avatar ? (<img src={"https://p3-pc-sign.douyinpic.com/tos-cn-i-0813c001/o8APGIcQn00NfE8aADb1AxAAxemkAKBBLIQ70f~tplv-dy-aweme-images-v2:3000:3000:q75.webp?biz_tag=aweme_images&from=327834062&s=PackSourceEnum_SEARCH&sc=image&se=false&x-expires=1727269200&x-signature=02VjKo6frbF4fMqDJhG0Hfh2e84%3D"}
                                                alt={""}
                                                style={{width: '100%'}}/>) : (<div>
                                    <UploadOutlined/>
                                    <div>点击上传头像,只能上传png/jpg文件，且不超过2M</div>
                                </div>)}
                            </Upload>
                        </Form.Item>
                        <Form.Item label="用户名:" name="username"
                                   rules={[{required: true, message: '请输入用户名'}]}>
                            <Input/>
                        </Form.Item>
                        <Form.Item wrapperCol={{offset: 8, span: 16}}>
                            <Button type="primary" htmlType="submit"
                                    style={{marginBottom: '83px'}}>保存</Button>
                        </Form.Item>
                    </Form>
                </div>
            </div>
        </div>) : (<div className={"Empty"}>
            <Empty
                image="https://gw.alipayobjects.com/zos/antfincdn/ZHrcdLPrvN/empty.svg"
                imageStyle={{
                    height: 160,
                }}
                description={<span>你还没有 <a href="/client/src/public">登录？</a></span>}>
                <Button type="primary" onClick={() => {
                    navigateTo("/login");
                }}>点击登录</Button>
            </Empty>
        </div>)}
    </>);
}
