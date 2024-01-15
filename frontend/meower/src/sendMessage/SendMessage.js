import SendMessageInput from '../common/sendMessageInput/SendMessageInput';
import styles from './SendMessage.module.css';

const SendMessage = (props) => {
    return (
        <div className={`${styles['send-message__wrapper']} ${props.className}`}>
            <SendMessageInput className={styles['send-message__input']} />
        </div>
    )
}

export default SendMessage; 