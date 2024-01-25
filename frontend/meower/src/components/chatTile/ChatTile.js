import { useDispatch } from 'react-redux';

import { mainStateActions } from '../../store/main/actions';
import styles from './ChatTile.module.css';

const ChatTile = (props) => {
    const dispatch = useDispatch();

    const chatAvatarSymbol = props.chat.name[0];

    const handleChatSelect = () => {
        dispatch(mainStateActions.selectChat(props.chat.id))
    }

    const chatTileStyles = `
        ${styles['chat-tile__wrapper']} 
        ${props.className} 
        ${props.chat.isSelected ? styles['selected'] : ''}
    `;

    return (
        <div className={chatTileStyles} onClick={handleChatSelect}>
            <div className={styles['tile-avatar']}><p>{chatAvatarSymbol}</p></div>
            <div className={styles['tile-main-content']}>
                <p className={styles['chat-name']}>{props.chat.name}</p>
                <p className={styles['last-message']}>Click here to read the messages feed...</p>
            </div>
            <p className={styles['last-message-date']}></p>
        </div>
    )
}

export default ChatTile;