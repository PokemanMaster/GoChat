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

    // Set the app element for accessibility
    useEffect(() => {
        Modal.setAppElement('#root'); // Ensure this matches your root element's ID
    }, []);

    // WebSocket connection
    // useEffect(() => {
    //     wsRef.current = new WebSocket(`ws://localhost:9000/api/v1/chat/send?userId=${userId}`);
    //     wsRef.current.onopen = () => setIsWsOpen(true);
    //     wsRef.current.onclose = () => setIsWsOpen(false);
    //     wsRef.current.onerror = (error) => console.error('WebSocket error:', error);
    //     wsRef.current.onmessage = (event) => {
    //         try {
    //             const receivedMessage = JSON.parse(event.data);
    //             if (receivedMessage.Type === 2) {
    //                 setMessages((prevMessages) => [...prevMessages, receivedMessage].sort((a, b) => a.CreateTime - b.CreateTime));
    //             }
    //         } catch (e) {
    //             console.error('Error parsing message:', e);
    //         }
    //     };
    // }, [userId]);

    // 用户好友列表
    const [userFriendData, setUserFriendData] = useState([]);
    useEffect(() => {
        FriendListsAPI().then(res => {
            console.log("FriendListsAPI", res)
            setUserFriendData(res.data)
        }).catch(err => console.log(err));
    }, []);

    const handleSend = () => {
        if (isWsOpen && wsRef.current && message) {
            const msg = {
                "TargetId": 1,
                "Type": 2,
                "CreateTime": Date.now(),
                "userId": 1,
                "Media": 1,
                "Content": message,
            };
            wsRef.current.send(JSON.stringify(msg));
            setMessages((prevMessages) => [...prevMessages, msg].sort((a, b) => a.CreateTime - b.CreateTime));
            setMessage('');
        } else {
            console.warn('WebSocket is not open or message is empty');
        }
    };

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
            // 显示操作
            closeModal()
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
                    <div className="menu" key={item.ID}>
                        <div className="contact-item">
                            <div className="contact-avatar">AB</div>
                            <div className="contact-info">
                                <div className="contact-name">{item.user_name}</div>
                                <div className="contact-last-message">等级：{item.level_id}
                                </div>
                                <div className="contact-time">{item.heartbeat_time}</div>
                            </div>
                            {/*<div className="unread-messages"></div>*/}
                        </div>
                    </div>
                )) : <div></div>}
            </div>

            <div className="chat-window">
                <div className="chat-messages">
                    <div className="chat-message incoming">
                        <div className="avatar">M</div>
                        <div className="message-content">Hi Marie, are you all right lately? I miss you very much. Do
                            you know?
                        </div>
                    </div>
                    <div className="chat-message outgoing">
                        <div className="avatar">Y</div>
                        <div className="message-content">I'm happy to hear you say that. I've been very good lately.
                        </div>
                    </div>
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
                                    <div className="contact-last-message">等级：{item.level_id}
                                    </div>
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
