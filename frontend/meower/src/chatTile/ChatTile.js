import styles from './ChatTile.module.css';

const ChatTile = (props) => {
    const chatAvatarSymbol = props.chat.name[0];

    return (
        <div className={`${styles['chat-tile__wrapper']} ${props.className}`}>
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