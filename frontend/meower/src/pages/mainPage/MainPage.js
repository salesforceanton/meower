import React, { useEffect } from 'react';
import useWebSocket from "react-use-websocket"
import { useDispatch } from 'react-redux';
import { mainStateActions } from '../../store/main/actions';

import AppHeader from '../../layout/appHeader/AppHeader';
import styles from './MainPage.module.css';
import FeedsSidebar from '../../layout/feedsSidebar/FeedsSidebar';
import Feed from '../../components/feed/Feed';

import { SERVICE_ENDPOINTS } from '../../store/service/constants';

const MainPage = () => {
    const dispatch = useDispatch();

    // Websocket connector
    const { lastJsonMessage, readyState } = useWebSocket(
        SERVICE_ENDPOINTS.WEBSOKET,
        {
            share: false,
            shouldReconnect: () => true,
        },
    )
    
    // Run when the connection state (readyState) changes
    useEffect(() => {
        console.log("Connection state changed" + readyState)
    }, [readyState])
    
    // Run when a new WebSocket message is received (lastJsonMessage)
    useEffect(() => {
        dispatch(mainStateActions.addNewMessage(lastJsonMessage));
    }, [lastJsonMessage, dispatch])

    // Chat list initialization
    useEffect(() => {
        const mockChatsData = [
            {
                name: 'Anonymous',
                id: 'anon1231'
            }
        ];
        dispatch(mainStateActions.setChatList(mockChatsData)); 
        //dispatch(mainStateActions.getChatList())
    }, [dispatch]);

    return (
        <React.Fragment>
            <AppHeader className={styles['app-header']}/>
            <div className={styles['content-wrapper']}>
                <FeedsSidebar/>
                <Feed/>
            </div>
        </React.Fragment>
    )
}

export default MainPage;