import { useDispatch } from 'react-redux';

import SendMessageInput from '../common/sendMessageInput/SendMessageInput';
import styles from './SendMessage.module.css';
import { mainStateActions } from '../../store/main/actions';

const SendMessage = (props) => {
    const dispatch = useDispatch();
    const handleSendMessage = (message) => dispatch(mainStateActions.sendMessage(message));
    
    return (
        <div className={`${styles['send-message__wrapper']} ${props.className}`}>
            <SendMessageInput className={styles['send-message__input']} onSend={handleSendMessage}/>
        </div>
    )
}

export default SendMessage; 