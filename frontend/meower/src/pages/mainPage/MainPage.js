import React, { useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { mainStateActions } from '../../store/main/actions';

import AppHeader from '../../layout/appHeader/AppHeader';
import styles from './MainPage.module.css';
import FeedsSidebar from '../../layout/feedsSidebar/FeedsSidebar';
import Feed from '../../feed/Feed';

const MainPage = () => {
    const dispatch = useDispatch();

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