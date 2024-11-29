import React, {useEffect, useState, useRef} from 'react';
import './style.less';
import {useLocation} from "react-router-dom";
import {ChatMessageAPI, CreateFriendAPI, FriendListsAPI, SearchFriendAPI} from "../../api/chat";
import Modal from 'react-modal';
import {SendOutlined} from "@ant-design/icons"; // Importing react-modal

function Chat() {
    const [message, setMessage] = useState('');
    const [messages, setMessages] = useState([]);
    const [isWsOpen, setIsWsOpen] = useState(false);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [targetId, setTargetId] = useState(''); // 新增状态来存储 target_id


    const wsRef = useRef(null);
    const location = useLocation();
    const queryParams = new URLSearchParams(location.search);
    const userId = queryParams.get('userId');

    const openModal = () => setIsModalOpen(true);
    const closeModal = () => {
        setIsModalOpen(false);
        setFriendName('');
    };

    useEffect(() => {
        Modal.setAppElement('#root'); // Ensure this matches your root element's ID
    }, []);

    useEffect(() => {
        console.log('Messages updated:', messages);
    }, [messages]); // 每当 messages 更新时，会执行这个 effect

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
        setTargetId(id)
        // 获取聊天记录
        ChatMessageAPI({
            "UserIdA": parseInt(userId, 10),
            "UserIdB": parseInt(id, 10),
            "Start": 0,
            "End": -1,
            "IsRev": true
        }).then(res => {
            const messages = res.data.map((msgStr) => {
                const msgObj = JSON.parse(msgStr);
                return {
                    UserId: msgObj.UserId,
                    TargetId: msgObj.TargetId,
                    Type: msgObj.Type,
                    Media: msgObj.Media,
                    Content: msgObj.Content,
                    CreateTime: msgObj.CreateTime,
                    ReadTime: msgObj.ReadTime,
                    Pic: msgObj.Pic,
                    Url: msgObj.Url,
                    Desc: msgObj.Desc,
                    Amount: msgObj.Amount
                };
            });
            setUserFriendMsg(messages); // 更新聊天信息
        }).catch(err => console.log(err));

        // 建立 WebSocket 连接
        const socket = new WebSocket(`ws://localhost:9000/api/v1/chat/send?userId=${userId}`);

        socket.onopen = () => {
            setIsWsOpen(true);
        };

        socket.onmessage = (event) => {
            console.log('Received message:', event.data);
            const receivedMsg = JSON.parse(event.data);
            setMessages((prevMessages) => [...prevMessages, receivedMsg].sort((a, b) => a.CreateTime - b.CreateTime));
        };

        socket.onclose = () => {
            console.log('WebSocket connection closed');
            setIsWsOpen(false);
        };
        socket.onerror = (error) => {
            console.error('WebSocket error:', error);
        };
        wsRef.current = socket;
    };



    const handleSend = () => {
        if (isWsOpen && wsRef.current && message) {
            const msg = {
                "UserId": parseInt(userId, 10), // 用户ID
                "TargetId": parseInt(targetId, 10), // 目标ID
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
            wsRef.current.send(JSON.stringify(msg));
            setMessages((prevMessages) => [...prevMessages, msg].sort((a, b) => a.CreateTime - b.CreateTime));
            setMessage('');
            ToChat(targetId); // 发送消息后立即获取聊天信息
        } else {
            console.warn('WebSocket is not open or message is empty');
        }
    };


    // 搜索好友
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

    // 添加好友
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
                             className={`chat-message ${msg.UserId === parseInt(userId, 10) ? 'chat-message-right' : 'chat-message-left'}`}>
                            <div className="message-content">{msg.Content}</div>
                            <div className={`avatar ${msg.UserId === parseInt(userId, 10) ? 'avatar-right' : 'avatar-left'}`}>{msg.UserId}</div>
                        </div>
                    ))}
                </div>
                <div className="chat-input">
                    <input
                        type="text"
                        value={message}
                        onChange={(e) => setMessage(e.target.value)}
                        placeholder="Write something..."
                    />
                    <button onClick={handleSend} className="send-btn">
                        <SendOutlined />
                    </button>
                </div>
            </div>
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
                        zIndex: 999,
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
