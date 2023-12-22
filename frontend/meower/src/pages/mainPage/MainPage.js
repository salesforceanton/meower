import React from 'react'

import AppHeader from '../../layout/appHeader/AppHeader';
import styles from './MainPage.module.css';
import FeedsSidebar from '../../feedsSidebar/FeedsSidebar';
import Feed from '../../feed/Feed';

const MainPage = () => {
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