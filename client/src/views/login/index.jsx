import * as React from 'react';
import {useEffect, useState} from 'react';
import {useNavigate} from "react-router-dom";
import "./style.less"
import {CategoryAPI, LoginAPI} from "../../api/users";
import IconButton from '@mui/material/IconButton';
import Input from '@mui/material/Input';
import InputLabel from '@mui/material/InputLabel';
import InputAdornment from '@mui/material/InputAdornment';
import FormControl from '@mui/material/FormControl';
import VisibilityIcon from '@mui/icons-material/Visibility';
import VisibilityOffIcon from '@mui/icons-material/VisibilityOff';
import {Box, Button, Link as MuiLink} from "@mui/material";
import Typography from '@mui/material/Typography';
import Modal from '@mui/material/Modal';
import {Col, Row} from "antd";

export default function Login() {
    // 初始化验证码图片
    const [codeId, setCodeId] = useState("")
    const [base64, setBase64] = useState("")
    useEffect(() => {
        CategoryAPI().then(res => {
            setCodeId(res.data["code_id"])
            setBase64(res.data["base64"])
        }).catch(err => {
            console.log("CategoryAPI:", err)
        })
    }, [])

    // 点击切换验证码图片
    const changeCategory = () => {
        CategoryAPI().then(res => {
            setCodeId(res.data["code_id"])
            setBase64(res.data["base64"])
        }).catch(err => {
            console.log(err)
        })
    }

    // 路由跳转定义
    const navigateTo = useNavigate()

    // 用户名处理
    const [userName, setUserName] = useState('');
    const handleUserNameChange = (event) => {
        setUserName(event.target.value);
    };

    // 密码处理
    const [password, setPassword] = useState('');
    const handlePasswordChange = (event) => {
        setPassword(event.target.value);
    };
    const [showPassword, setShowPassword] = React.useState(false);
    const handleClickShowPassword = () => setShowPassword((show) => !show);
    const handleMouseDownPassword = (event) => {
        event.preventDefault();
    };

    // 验证码处理
    const [code, setCode] = useState('');
    const handleCodeChange = (event) => {
        setCode(event.target.value);
    };

    // 弹出层处理
    const style = {
        position: 'absolute',
        top: '50%',
        left: '50%',
        transform: 'translate(-50%, -50%)',
        width: 300,
        bgcolor: 'background.paper',
        boxShadow: 24,
        p: 3,
    };
    const [message, setMessage] = useState('');
    const [open, setOpen] = React.useState(false);
    const handleOpen = (msg) => {
        setOpen(true)
        setMessage(msg)
    };
    const handleClose = () => setOpen(false);

    // 提交处理
    const submitForm = () => {
        LoginAPI({
            "UserName": userName,
            "Password": password,
            "Code": code,
            "CodeId": codeId,
        }).then(res => {
            handleOpen(res.msg)
            const GoChatID = res.data.data.ID || "";
            localStorage.setItem('token', res.data.token);
            localStorage.setItem('user' + GoChatID, JSON.stringify(res.data.data));
            navigateTo(`/layout/products/sort?userId=${GoChatID}`)
        }).catch(err => {
            if (err.response.data.status === 40001) {
                handleOpen("账号或者密码不能为空")
            } else {
                handleOpen("系统错误")
            }
        })
    }

    // 路由跳转处理
    function toRegister() {
        navigateTo("/register")
    }

    return (
        <div className={"gochat"}>
            <div className={"gochat-app"}>
                <div className={"gochat-app-left"}></div>
                <div className={"gochat-app-right"}>
                    <img className={"right-logo"} src="/logo/logo.png" alt=""/>
                    <p>Sign in to your Account</p>
                    <div className={"right-form"}>
                        <div className={"right-form-username"}>
                            <FormControl variant="standard" style={{width: '100%'}}>
                                <InputLabel htmlFor="standard-adornment-password"
                                            style={{color: '#3E3F4C'}}>用户名：admin123</InputLabel>
                                <Input
                                    id="standard-adornment-userName"
                                    value={userName}
                                    onChange={handleUserNameChange}
                                    inputProps={{maxLength: 11}} // 限制最大长度为11
                                />
                            </FormControl>
                        </div>
                        <div className={"right-form-password"}>
                            <FormControl variant="standard" style={{width: '100%'}}>
                                <InputLabel htmlFor="standard-adornment-password"
                                            style={{color: '#3E3F4C'}}>密码：123123</InputLabel>
                                <Input
                                    id="standard-adornment-password"
                                    value={password}
                                    onChange={handlePasswordChange}
                                    type={showPassword ? 'text' : 'password'}
                                    inputProps={{maxLength: 20}} // 限制最大长度为11
                                    endAdornment={<InputAdornment position="end">
                                        <IconButton
                                            aria-label="toggle password visibility"
                                            onClick={handleClickShowPassword}
                                            onMouseDown={handleMouseDownPassword}
                                        >
                                            {showPassword ? <VisibilityOffIcon/> : <VisibilityIcon/>}
                                        </IconButton>
                                    </InputAdornment>}
                                />
                            </FormControl>
                        </div>
                        <div className={"right-form-code"}>
                            <FormControl sx={{display: 'flex', flexDirection: 'row', width: '100%'}} variant="standard">
                                <InputLabel htmlFor="standard-code"
                                            style={{color: '#3E3F4C'}}>验证码</InputLabel>
                                <Input
                                    id="standard-code"
                                    value={code}
                                    inputProps={{maxLength: 5}} // 限制最大长度为11
                                    onChange={handleCodeChange}
                                    variant="standard"
                                    fullWidth
                                />
                                <img onClick={changeCategory} src={base64} style={{width: '40%'}} alt=""/>
                            </FormControl>
                        </div>
                        <div className={"right-form-submit"}>
                            <Button variant="contained" style={{width: '100%'}}
                                    onClick={submitForm}>登录</Button>
                        </div>
                        <div className={"right-form-register"} onClick={toRegister}>
                            点击注册
                        </div>
                    </div>
                    <div className={"right-link"}>
                        <MuiLink href="https://beian.miit.gov.cn/" underline="none" style={{color: "#333"}}>
                            桂ICP备2023004200号-2
                        </MuiLink>
                    </div>
                </div>
            </div>

            {/*弹出层*/}
            <div>
            <Modal
                    open={open}
                    onClose={handleClose}
                    aria-labelledby="modal-modal-title"
                    aria-describedby="modal-modal-description"
                >
                    <Box sx={style}>
                        <Typography id="modal-modal-title" variant="h6" component="h2">
                            Text in a modal
                        </Typography>
                        <Typography id="modal-modal-description" sx={{mt: 2}}>
                            {message}
                        </Typography>
                    </Box>
                </Modal>
            </div>
        </div>)
}