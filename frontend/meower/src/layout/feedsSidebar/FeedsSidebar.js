import ChatList from '../../components/chatList/ChatList';
import Search from '../../components/search/Search';

import styles from './FeedsSidebar.module.css';

const FeedsSidebar = () => {
    return (
        <div className={styles['feeds-sidebar__wrapper']}>
            <Search/>
            <ChatList/>
        </div>
    )
}
export default FeedsSidebar;