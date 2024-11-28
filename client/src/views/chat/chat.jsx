import React, {useEffect, useState, useRef} from 'react';
import './style.less';
import {useLocation} from "react-router-dom";
import {ChatMessageAPI, CreateFriendAPI, FriendListsAPI, SearchFriendAPI} from "../../api/chat";
import Modal from 'react-modal'; // Importing react-modal

function Chat() {
    const [message, setMessage] = useState('');
    const [messages, setMessages] = useState([]);
    const [isWsOpen, setIsWsOpen] = useState(false);
    const [isModalOpen, setIsModalOpen] = useState(false);

    const location = useLocation();
    const wsRef = useRef(null);
    const queryParams = new URLSearchParams(location.search);
    const userId = queryParams.get('userId');

    useEffect(() => {
        console.log('Messages updated:', messages);
    }, [messages]); // 每当 messages 更新时，会执行这个 effect

    useEffect(() => {
        Modal.setAppElement('#root'); // Ensure this matches your root element's ID
    }, []);

    // 用户好友列表
    const [userFriendData, setUserFriendData] = useState([]);
    useEffect(() => {
        FriendListsAPI().then(res => {
            setUserFriendData(res.data)
        }).catch(err => console.log(err));
    }, []);

    // 和好友聊天
    const [userFriendMsg, setUserFriendMsg] = useState([]);
    const ToChat = (id) => {
        console.log(userId, id)
        // 获取聊天记录
        ChatMessageAPI({
            "UserIdA": parseInt(userId, 10),
            "UserIdB": parseInt(id, 10),
            "Start": 0,
            "End": -1,
            "IsRev": true
        }).then(res => {
            console.log(res.data);
            // 假设 res.data 是一个包含 JSON 字符串的数组
            const messages = res.data.map((msgStr) => {
                // 解析每个 JSON 字符串为 JavaScript 对象
                const msgObj = JSON.parse(msgStr);
                // 返回一个新的对象，提取需要的数据
                return {
                    UserId: msgObj.UserId,
                    TargetId: msgObj.TargetId,
                    Type: msgObj.Type,
                    Media: msgObj.Media,
                    Content: msgObj.Content, // 获取 Content 信息
                    CreateTime: msgObj.CreateTime,
                    ReadTime: msgObj.ReadTime,
                    Pic: msgObj.Pic,
                    Url: msgObj.Url,
                    Desc: msgObj.Desc,
                    Amount: msgObj.Amount
                };
            });
            // 将解析后的消息设置到 state 中
            setUserFriendMsg(messages);
        }).catch(err => console.log(err));


        // 建立 WebSocket 连接
        const socket = new WebSocket(`ws://localhost:9000/api/v1/chat/send?userId=${userId}`);

        socket.onopen = () => { // 连接打开后，设置 WebSocket 状态
            console.log('WebSocket connected');
            setIsWsOpen(true);
            socket.send('Hello, WebSocket!'); // 你可以发送一条初始消息（如果需要）
        };

        socket.onmessage = (event) => { // 监听收到的消息
            console.log('Received message:', event.data);
            const receivedMsg = JSON.parse(event.data);
            setMessages((prevMessages) => [...prevMessages, receivedMsg].sort((a, b) => a.CreateTime - b.CreateTime));
        };

        socket.onclose = () => { // 连接关闭时
            console.log('WebSocket connection closed');
            setIsWsOpen(false);
        };

        socket.onerror = (error) => { // 错误处理
            console.error('WebSocket error:', error);
        };

        wsRef.current = socket; // 保存 WebSocket 实例
    };

    const handleSend = () => {
        if (isWsOpen && wsRef.current && message) {
            const msg = {
                "UserId": parseInt(userId, 10), // 用户ID
                "TargetId": 2, // 目标ID（假设你发送给的好友ID）
                "Type": 1, // 消息类型
                "Media": 1, // 媒体类型
                "Content": message, // 消息内容
                "CreateTime": Date.now(), // 当前时间戳
                "ReadTime": 0, // 阅读时间
                "Pic": null, // 图片（如果有的话）
                "Url": null, // URL（如果有的话）
                "Desc": null, // 描述（如果有的话）
                "Amount": 3 // 数量（如果有的话）
            };

            // 发送消息
            wsRef.current.send(JSON.stringify(msg));

            // 更新消息列表，立刻将新消息添加到当前聊天窗口
            setMessages((prevMessages) => [...prevMessages, msg].sort((a, b) => a.CreateTime - b.CreateTime));

            // 清空输入框
            setMessage('');
        } else {
            console.warn('WebSocket is not open or message is empty');
        }
    };


    // Add Friend Modal
    const openModal = () => setIsModalOpen(true);
    const closeModal = () => {
        setIsModalOpen(false);
        setFriendName('');
    };

    const [friendName, setFriendName] = useState('');
    const [friendData, setFriendData] = useState('');
    const handleSearch = () => {
        SearchFriendAPI({
            "UserID": parseInt(userId, 10),
            "FriendName": friendName,
        }).then(res => {
            setFriendData(res.data)
        }).catch(err => {
            console.log(err)
        })
    };

    const CreateFriend = (id) => {
        CreateFriendAPI({
            "UserID": parseInt(userId, 10),
            "TargetID": id,
        }).then(res => {
            closeModal();
        }).catch(err => {
            console.log(err)
        })
    };

    return (
        <div className="chat-app">
            <div className="chat-sidebar">
                <div className="sidebar-top">
                    <div className="user-info">
                        <div className="user-name">添加好友</div>
                    </div>
                    <div className="add-group-btn">
                        <button className="add-btn" onClick={openModal}>+</button>
                    </div>
                </div>
                {userFriendData && userFriendData.length > 0 ? userFriendData.map((item, index) => (
                    <div className="menu" key={item.ID} onClick={() => ToChat(item.ID)}>
                        <div className="contact-item">
                            <div className="contact-avatar">AB</div>
                            <div className="contact-info">
                                <div className="contact-name">{item.user_name}</div>
                                <div className="contact-last-message">等级：{item.level_id}</div>
                                <div className="contact-time">{item.heartbeat_time}</div>
                            </div>
                        </div>
                    </div>
                )) : <div></div>}
            </div>

            <div className="chat-window">
                <div className="chat-messages">
                    {userFriendMsg.map((msg, index) => (
                        <div key={index}
                             className={`chat-message ${msg.UserId === parseInt(userId, 10) ? 'outgoing' : 'incoming'}`}>
                            <div className="avatar">{msg.UserId}</div>
                            <div className="message-content">{msg.Content}</div>
                            {/* 显示 Content */}
                        </div>
                    ))}
                </div>
                <div className="chat-input">
                    <input type="text" value={message} onChange={(e) => setMessage(e.target.value)}
                           placeholder="Write something..."/>
                    <button onClick={handleSend}>Send</button>
                </div>
            </div>


            {/* Add Friend Modal */}
            <Modal
                isOpen={isModalOpen}
                onRequestClose={closeModal}
                contentLabel="Add Friend"
                overlayClassName="overlay"
                shouldCloseOnOverlayClick={true}
                style={{
                    content: {
                        height: 'auto',   // 高度自适应
                        margin: 'auto auto',  // 居中显示
                        border: 'none',   // 去除边框
                        backgroundColor: 'transparent',  // 设置背景透明
                        padding: '20px',  // 为内容添加内边距
                    },
                    overlay: {
                        backgroundColor: 'rgba(0, 0, 0, 0.5)',  // 遮罩背景色，半透明黑色
                        display: 'flex',
                        justifyContent: 'center',
                        alignItems: 'center',
                        position: 'fixed',
                        top: '0',
                        left: '0',
                        right: '0',
                        bottom: '0',
                        zIndex: 999,  // 确保弹窗层级高于其他内容
                    }
                }}
            >
                <div className="modal-main">
                    <h2>请输入要查询的好友名字</h2>
                    <div className="modal-content">
                        <input
                            type="text"
                            value={friendName}
                            onChange={(e) => setFriendName(e.target.value)}
                            placeholder="Enter friend's name"
                            className="modal-input"
                        />
                        <button className="search-btn" onClick={handleSearch}>搜索</button>
                    </div>
                    {friendData && friendData.length > 0 ? friendData.map((item, index) => (
                        <div className="menu" key={item.ID}>
                            <div className="contact-item">
                                <div className="contact-avatar">AB</div>
                                <div className="contact-info">
                                    <div className="contact-name">{item.user_name}</div>
                                    <div className="contact-last-message">等级：{item.level_id}</div>
                                </div>
                                <div className="unread-messages" onClick={() => CreateFriend(item.ID)}>添加好友</div>
                            </div>
                        </div>
                    )) : <div></div>}
                    <button className="close-btn" onClick={closeModal}>关闭</button>
                </div>
            </Modal>
        </div>
    );
}

export default Chat;
