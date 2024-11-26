import "./style.less"
import {Button, Drawer, Empty, Form, Input, Modal} from "antd";
import React, {useEffect, useState} from "react";
import {CreateAddressAPI, DeleteAddressAPI, ShowAddressesAPI, UpdateAddressAPI} from "../../api/addresses";
import emptyCart from "../../public/images/cart_empty.png"
import {Link} from "@mui/joy";
import clogo from "../../public/images/clogo.png";
import {cityData} from "../../public/data/city";
import {useNavigate} from "react-router-dom";

export default function UserAddress() {
    const navigateTo = useNavigate();
    const [UserInfo] = useState(JSON.parse(localStorage.getItem("user")));
    const [form] = Form.useForm();

    // 展示收货地址信息
    const [ShowAddresses, setShowAddresses] = useState([])
    useEffect(() => {
        ShowAddressesAPI().then(res => {
            console.log("展示收货地址信息", res)
            setShowAddresses(res.data)
        }).catch(err => {
            console.log(err)
        });
    }, [])

    // 删除收货地址
    const delectAddress = (id) => {
        DeleteAddressAPI({"AddressID": id})
            .then(res => {
                console.log("删除收货地址", res)
            }).catch(err => {
            console.log(err)
        })
        window.location.reload();
    }

    // 添加收货地址
    const [isCreateAddressModalOpen, setIsCreateAddressModalOpen] = useState(false);
    const CreateAddressModal = (bool) => {
        setIsCreateAddressModalOpen(bool);
    };
    const CreateAddress = () => { // 表单提交
        form.validateFields().then(values => {
            // 在这里执行您的提交逻辑
            const addressesValue = values["addresses"];
            const telephoneValue = values["telephone"];
            const drawerValue = values["drawer"];
            const detailsValue = values["details"];
            const addressValue = drawerValue + '-' + detailsValue;
            // 添加收货地址
            CreateAddressAPI({
                "UserID": UserInfo.id, "Name": addressesValue, "Telephone": telephoneValue, "Address": addressValue
            }).then(res => {
                console.log("添加收货地址", res)
            }).catch(err => {
                console.log(err)
            })
            CreateAddressModal(false);
            window.location.reload();
        }).catch(errorInfo => {
            console.log('表单验证失败:', errorInfo);
        });
    };

    // 修改收货地址
    const [isUpdateAddressModalOpen, setIsUpdateAddressModalOpen] = useState(false);
    const [UpdateAddressId, setUpdateAddressId] = useState("");
    const UpdateAddressModal = (bool, id, name, phone) => {
        form.setFieldsValue({addresses: name, telephone: phone});
        setUpdateAddressId(id);
        setIsUpdateAddressModalOpen(bool);
    };
    const UpdateAddress = () => {
        form.validateFields().then(values => {
            // 在这里执行您的提交逻辑
            const addressesValue = values["addresses"];
            const telephoneValue = values["telephone"];
            const drawerValue = values["drawer"];
            const detailsValue = values["details"];
            const addressValue = drawerValue + '-' + detailsValue;
            // 添加收货地址
            UpdateAddressAPI({
                "UserID": UserInfo.id,
                "ID": UpdateAddressId,
                "Name": addressesValue,
                "Telephone": telephoneValue,
                "Address": addressValue,
                "Prime": false,
            }).then(res => {
                console.log("添加收货地址", res)
            }).catch(err => {
                console.log(err)
            })
            UpdateAddressModal(false);
            window.location.reload();
        }).catch(errorInfo => {
            console.log('表单验证失败:', errorInfo);
        });
    }

    // 打开/关闭地址 抽屉
    const [OpenDrawer, setOpenDrawer] = useState(false);
    const DrawerModal = (bool) => {
        setOpenDrawer(bool);
    };

    // 地址选择(省)
    const [selectedProvince, setSelectedProvince] = useState('');
    const [Province, setProvince] = useState(true);
    const ProvinceClick = () => {
        setCity(false);
        setDistrict(false);
        setProvince(true);
    };
    const SelectedProvince = (Province_Label) => {
        setSelectedProvince(Province_Label)
        setProvince(false);
        setCity(true);
    }
    // 地址选择(市)
    const [selectedCity, setSelectedCity] = useState('');
    const [City, setCity] = useState(false);
    const CityClick = () => {
        setProvince(false);
        setDistrict(false);
        setCity(true);
    };
    const SelectedCity = (City_Label) => {
        setSelectedCity(City_Label)
        setCity(false);
        setDistrict(true);
    }
    // 地址选择(区)
    const [selectedDistrict, setSelectedDistrict] = useState('');
    const [District, setDistrict] = useState(false);
    const DistrictClick = () => {
        setProvince(false);
        setCity(false);
        setDistrict(true);
    };
    const SelectedDistrict = (District_Label) => {
        setSelectedDistrict(District_Label)
        const selectedArea = `${selectedProvince || ''}${selectedCity ? '-' + selectedCity : ''}${District_Label ? '-' + District_Label : ''}`
        form.setFieldsValue({drawer: selectedArea})
        DrawerModal(false)
    }

    return (<>
        {UserInfo ? <div>
                {/*头部*/}
                <div className={"TopHeader"}>
                    <div className={"CartHeader"}>
                        <div className={"Logo"}>
                            <Link to="/">
                                <img src={clogo} alt=""/>
                            </Link>
                        </div>
                        <div className={"CartHeaderContent"}>
                            <p>我的地址</p>
                        </div>
                    </div>
                </div>

                <div className={"address"}>
                    {ShowAddresses ? (ShowAddresses.map((item) => (<div className={"addressItem"} key={item.id}>
                        <div className={"itemProfile"}>
                            <span>{item.name}&nbsp;,</span>
                            <span className={"itemTelephone"}>{item.Telephone}</span>
                            <span className={"itemDelete"} onClick={() => delectAddress(item.id)}>X</span>
                            <p>{item.address}</p>
                        </div>
                        <div className={"itemButton"}>
                            <div className={"default"}></div>
                            <div className={"selectButton"}>
                                <Button className={"updateButton"}
                                        onClick={() => UpdateAddressModal(true, item.id, item.name, item.phone)}
                                        type="primary">修改</Button>
                            </div>
                        </div>
                    </div>))) : <div className={"EmptyCart"}>
                        {/* 此处的图片不能直接写路径，只能通过import的方式将它引入进来 */}
                        <img src={emptyCart} alt="" className={"EmptyCartImg"}/>
                        <div className={"EmptyCartText1"}>收货地址竟然是空的！</div>
                        <div className={"EmptyCartText2"}>赶紧创建一个收货地址把</div>
                    </div>}
                    <Button className={"addressButton"} onClick={CreateAddressModal} type="primary">添加收货地址</Button>
                </div>
            </div> :                 // 用户没登录就显示
            <div className={"Empty"}>
                <Empty
                    image="https://gw.alipayobjects.com/zos/antfincdn/ZHrcdLPrvN/empty.svg"
                    imageStyle={{
                        height: 160,
                    }}
                    description={<span>你还没有 <a href="/">登录？</a></span>}>
                    <Button type="primary" onClick={() => {
                        navigateTo("/login")
                    }}>点击登录</Button>
                </Empty>
            </div>}


        {/* 添加收货地址models虚拟栏 */}
        <Modal title="添加收货地址" centered={true} open={isCreateAddressModalOpen} onOk={CreateAddress}
               onCancel={() => CreateAddressModal(false)}>
            <Form form={form}>
                <Form.Item label="收货人:" name="addresses" style={{marginTop: '37px'}}
                           rules={[{required: true, message: '收货人不能为空'}]}>
                    <Input/>
                </Form.Item>
                <Form.Item label="电话:" name="telephone"
                           rules={[{required: true, message: '电话不能为空'}]}>
                    <Input/>
                </Form.Item>
                <Form.Item label="地区:" name="drawer" onClick={() => DrawerModal(true)}
                           rules={[{required: true, message: '地区不能为空'}]}>
                    <Input readOnly onClick={() => DrawerModal(true)} id="drawer-input"/>
                </Form.Item>
                <Form.Item label="详情地址（如：街道/小区/乡镇/村）:" name="details"
                           rules={[{required: true, message: '详情地址不能为空'}]}>
                    <Input/>
                </Form.Item>
            </Form>
        </Modal>

        {/* 修改收货地址models虚拟栏 */}
        <Modal title="修改收货地址" centered={true} open={isUpdateAddressModalOpen} onOk={UpdateAddress}
               onCancel={() => UpdateAddressModal(false)}>
            <Form form={form}>
                <Form.Item label="收货人:" name="addresses" style={{marginTop: '37px'}}
                           rules={[{required: true, message: '收货人不能为空'}]}>
                    <Input/>
                </Form.Item>
                <Form.Item label="电话:" name="telephone"
                           rules={[{required: true, message: '电话不能为空'}]}>
                    <Input/>
                </Form.Item>
                <Form.Item label="地区:" name="drawer" onClick={() => DrawerModal(true)}
                           rules={[{required: true, message: '地区不能为空'}]}>
                    <Input readOnly onClick={() => DrawerModal(true)} id="drawer-input"/>
                </Form.Item>
                <Form.Item label="详情地址（如：街道/小区/乡镇/村）:" name="details"
                           rules={[{required: true, message: '详情地址不能为空'}]}>
                    <Input/>
                </Form.Item>
            </Form>
        </Modal>

        {/* 选择抽屉地址栏 */}
        <Drawer
            placement={"bottom"}
            closable={false}
            onClose={() => DrawerModal(false)}
            open={OpenDrawer}
            key={"bottom"}
        >
            <div className={"DrawerHeader"}>
                <span onClick={() => ProvinceClick(selectedProvince)}>
                    {selectedProvince ? selectedProvince : '请选择'}
                </span>
                <span onClick={() => CityClick(selectedCity)}>
                    {selectedCity ? selectedCity : '请选择'}
                </span>
                <span onClick={() => DistrictClick(selectedDistrict)}>
                    {selectedDistrict ? selectedDistrict : '请选择'}
                </span>
            </div>
            <div>
                {Province && (<div>
                    {cityData.map(province => (<p key={province.value} onClick={() => SelectedProvince(province.label)}>
                        {province.label}
                    </p>))}
                </div>)}
                {City && (<div>
                    {selectedProvince && cityData.find(province => province.label === selectedProvince)
                        .children.map(city => (
                            <p key={city.value} onClick={() => SelectedCity(city.label)}>{city.label}</p>))}
                </div>)}
                {District && (<div>
                    {selectedCity && cityData.find(province => province.label === selectedProvince)
                        .children.find(city => city.label === selectedCity)
                        .children.map(district => (<p key={district.value}
                                                      onClick={() => SelectedDistrict(district.label)}>{district.label}</p>))}
                </div>)}
            </div>
        </Drawer>


    </>)
}