import { useSelector } from 'react-redux';

import styles from './ChatList.module.css';

import { mainStateSelectors } from '../../store/main/selectors';
import ChatTile from '../chatTile/ChatTile';

const ChatList = (props) => {
    const chatListData = useSelector(mainStateSelectors.selectChatList);

    return (
        <div className={`${styles['chat-list__wrapper']} ${props.className}`}>
            {chatListData.map((e) => 
                <ChatTile 
                    chat={e} 
                    className={styles['chat-tile']}
                    key={e.id}
                />
            )}
        </div>
    )
}

export default ChatList;