import React from 'react';
import { useSelector } from 'react-redux';

import { mainStateSelectors } from '../../store/main/selectors';

import styles from './Feed.module.css';
import SendMessage from '../sendMessage/SendMessage';
import MessagesList from '../messagesList/MessagesList';

const Feed = () => {
    const showSelectChatMessage = useSelector(mainStateSelectors.selectIsNoSelectedChat);

    return (
        <div className={styles['feed-wrapper']}>
            {showSelectChatMessage 
            ?
                <div className={styles['empty-feed__message']}>
                    <div><span>Select chat to start messaging</span></div>
                </div>
            :
            <React.Fragment>
                <MessagesList/>
                <SendMessage/>
            </React.Fragment>
            }
        </div>
    )
}

export default Feed;