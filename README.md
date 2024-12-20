<img src="./client/public/logo/logo.png" alt="Logo" width="615"/>

## GoChat 项目概述 
GoChat是一个结合了电子商务功能和即时通讯（IM）的系统。

### 在线体验项目
https://www.lvyouwang.xyz/
```
账号：admin123
密码：123123

账号：admin1234
密码：123123
```

### 功能介绍
- 用户管理：用户注册、登录、信息修改等。
- 购物车管理：购物车的创建、修改、删除、展示等。
- 商品管理：商品的介绍、分类、详情、搜索等。
- 订单管理：订单的创建、修改、删除等。
- 支付管理：用户可以支付订单。
- 好友管理：用户可以添加好友、管理好友等。
- 聊天管理：支持与好友实时聊天功能。

### 前端项目结构（client）
- src
  - api：与后端 api 交互接口。
  - components：页面的组件。
  - layout：侧边栏菜单。
  - public：存放公共数据。
  - views：存放各个页面。

### 后端项目结构（server）
- app：存放应用的核心模块逻辑。
- cmd：项目的启动入口。
- common：包含数据库和缓存相关的定义与配置。
- pkg：存放第三方工具或包。
- resp：统一封装返回数据的格式。
- route ：路由的配置和逻辑处理。

### 启动项目
```
git clone https://github.com/PokemanMaster/GoChat.git
cd GoChat
docker-compose up -d
```