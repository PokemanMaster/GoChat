import * as React from 'react';
import {useEffect, useState} from 'react';
import {useNavigate} from "react-router-dom";
import "./style.less"
import {CategoryAPI, RegisterAPI} from "../../api/users";
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

export default function Register() {
    // 路由跳转
    const navigateTo = useNavigate()

    function toLogin() {
        navigateTo("/login")
    }


    // 初始化验证码图片
    const [codeId, setCodeId] = useState("")
    const [base64, setBase64] = useState("")
    useEffect(() => {
        CategoryAPI().then(res => {
            setCodeId(res.data["code_id"])
            setBase64(res.data["base64"])
        })
    }, [])


    // 点击切换验证码图片
    const changeCategory = () => {
        CategoryAPI().then(res => {
            setCodeId(res.data["code_id"])
            setBase64(res.data["base64"])
        })
    }


    // userName处理
    const [name, setName] = useState('');
    const handleNameChange = (event) => {
        setName(event.target.value);
    };


    // password处理
    const [password, setPassword] = useState('');
    const handlePasswordChange = (event) => {
        setPassword(event.target.value);
    };
    const [showPassword, setShowPassword] = useState(false);
    const handleClickShowPassword = () => setShowPassword((show) => !show);
    const handleMouseDownPassword = (event) => {
        event.preventDefault();
    };

    // password处理
    const [RePassword, setRePassword] = useState('');
    const handleRePasswordChange = (event) => {
        setRePassword(event.target.value);
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
        RegisterAPI({
            "UserName": name,
            "Password": password,
            "RePassword" : RePassword,
            "Code": code,
            "CodeId": codeId,
        }).then(res => {
            handleOpen(res.msg)
            navigateTo("/login")
        }).catch(err => {
            handleOpen(err.response.data.msg)
        })
    }


    return (
        <div className={"body"}>
            <Row justify="center" align="top">
                <Col xs={24} sm={24} md={24} lg={24} xl={24}>
                    <h1>模拟商城旅游网</h1>
                </Col>
                <Col xs={24} sm={24} md={24} lg={24} xl={24}>
                    <p>心态特好&nbsp;&nbsp;&nbsp;从不摆烂</p>
                </Col>
            </Row>
            <Row>
                <Col xs={3} sm={4} md={6} lg={8} xl={9}></Col>
                <Col xs={18} sm={16} md={12} lg={8} xl={6}>
                    <h2>用户注册</h2>
                </Col>
            </Row>

            {/*用户输入*/}
            <Row>
                <Col xs={3} sm={4} md={6} lg={8} xl={9}></Col>
                <Col xs={18} sm={16} md={12} lg={8} xl={6}>
                    <FormControl variant="standard" style={{width: '100%'}}>
                        <InputLabel htmlFor="standard-adornment-password" style={{color: '#3E3F4C'}}>用户名</InputLabel>
                        <Input
                            id="standard-adornment-name"
                            value={name}
                            onChange={handleNameChange}
                            inputProps={{maxLength: 11}} // 限制最大长度为11
                        />
                    </FormControl>
                </Col>
                <Col xs={3} sm={4} md={6} lg={8} xl={9}></Col>
            </Row>

            {/*密码输入*/}
            <Row>
                <Col xs={3} sm={4} md={6} lg={8} xl={9}></Col>
                <Col xs={18} sm={16} md={12} lg={8} xl={6}>
                    <FormControl variant="standard" style={{width: '100%'}}>
                        <InputLabel htmlFor="standard-adornment-password"
                                    style={{color: '#3E3F4C'}}>密码</InputLabel>
                        <Input
                            id="standard-adornment-password"
                            value={password}
                            onChange={handlePasswordChange}
                            type={showPassword ? 'text' : 'password'}
                            inputProps={{maxLength: 20}} // 限制最大长度为11
                            endAdornment={
                                <InputAdornment position="end">
                                    <IconButton
                                        aria-label="toggle password visibility"
                                        onClick={handleClickShowPassword}
                                        onMouseDown={handleMouseDownPassword}
                                    >
                                        {showPassword ? <VisibilityOffIcon/> : <VisibilityIcon/>}
                                    </IconButton>
                                </InputAdornment>
                            }
                        />
                    </FormControl>
                </Col>
                <Col xs={3} sm={4} md={6} lg={8} xl={9}></Col>
            </Row>

            {/*再次密码输入*/}
            <Row>
                <Col xs={3} sm={4} md={6} lg={8} xl={9}></Col>
                <Col xs={18} sm={16} md={12} lg={8} xl={6}>
                    <FormControl variant="standard" style={{width: '100%'}}>
                        <InputLabel htmlFor="standard-adornment-password"
                                    style={{color: '#3E3F4C'}}>再次输入密码</InputLabel>
                        <Input
                            id="standard-adornment-repassword"
                            value={RePassword}
                            onChange={handleRePasswordChange}
                            type={showPassword ? 'text' : 'password'}
                            inputProps={{maxLength: 20}} // 限制最大长度为11
                            endAdornment={
                                <InputAdornment position="end">
                                    <IconButton
                                        aria-label="toggle password visibility"
                                        onClick={handleClickShowPassword}
                                        onMouseDown={handleMouseDownPassword}
                                    >
                                        {showPassword ? <VisibilityOffIcon/> : <VisibilityIcon/>}
                                    </IconButton>
                                </InputAdornment>
                            }
                        />
                    </FormControl>
                </Col>
                <Col xs={3} sm={4} md={6} lg={8} xl={9}></Col>
            </Row>

            {/*验证码*/}
            <Row>
                <Col xs={3} sm={4} md={6} lg={8} xl={9}></Col>
                <Col xs={18} sm={16} md={12} lg={8} xl={6}>
                    <FormControl sx={{display: 'flex', flexDirection: 'row', width: '100%'}} variant="standard">
                        <InputLabel htmlFor="standard-code"
                                    style={{color: '#3E3F4C', paddingTop: '7px'}}>验证码</InputLabel>
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
                </Col>
                <Col xs={3} sm={4} md={6} lg={8} xl={9}></Col>
            </Row>

            {/*提交*/}
            <Row>
                <Col xs={3} sm={4} md={6} lg={8} xl={9}></Col>
                <Col xs={18} sm={16} md={12} lg={8} xl={6}>
                    <Button variant="contained" style={{width: '100%'}}
                            onClick={submitForm}>登录</Button>
                </Col>
                <Col xs={3} sm={4} md={6} lg={8} xl={9}></Col>
            </Row>
            {/*遗忘处理*/}
            <Row style={{textAlign: 'center'}}>
                <Col xs={6} sm={6} md={6} lg={8} xl={10} style={{textAlign: 'right'}}>
                </Col>
                <Col xs={12} sm={12} md={12} lg={8} xl={4}>
                </Col>
                <Col xs={6} sm={6} md={6} lg={8} xl={10} style={{textAlign: 'left', cursor: 'pointer'}}>
                    <div
                        style={{
                            color: 'blue',
                            fontWeight: '700',
                            fontSize: '18px' // 你可以根据需要调整大小
                        }}
                        onClick={toLogin}
                    >
                        点击登录
                    </div>
                </Col>
            </Row>

            <Row style={{textAlign: 'center'}}>
                <Col xs={24} sm={24} md={24} lg={24} xl={24}>
                    <MuiLink href="https://beian.miit.gov.cn/" underline="none" style={{color: "#333"}}>
                        桂ICP备2023004200号-1
                    </MuiLink>
                </Col>
            </Row>


            {/*弹出层*/}
            <div>
                {/*npm install react-modal*/}
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
        </div>
    )
}