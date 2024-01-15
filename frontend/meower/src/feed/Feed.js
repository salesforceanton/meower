import { useSelector } from 'react-redux';

import styles from './Feed.module.css';
import { mainStateSelectors } from '../store/main/selectors';
import SendMessage from '../sendMessage/SendMessage';

const Feed = () => {
    const showSelectChatMessage = useSelector(mainStateSelectors.selectIsNoSelectedChat);
    return (
        <div className={styles['feed-wrapper']}>
            {showSelectChatMessage && 
                <div className={styles['empty-feed__message']}>
                    <div><span>Select chat to start messaging</span></div>
                </div>
            }
            <SendMessage/>
        </div>
    )
}

export default Feed;