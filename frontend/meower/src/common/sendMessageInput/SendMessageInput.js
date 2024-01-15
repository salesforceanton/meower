import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPaw } from '@fortawesome/free-solid-svg-icons';

import { useRef } from 'react';
import styles from './SendMessageInput.module.css';
import { KeyCodes } from '../utils/Utils';

const SendMessageInput = (props) => {
    const messageInputRef = useRef();

    const handleKeyDown = (e) => {
        if (e.key === KeyCodes.ENTER) {
            sendMessage();
        }
    }
    const sendMessage = () => props.onSend(messageInputRef.current.value);

    return (
        <div className={styles['send-message-input__container']}>
            <input 
                className={`${styles['send-message-input']} ${props.className}`} 
                placeholder='Write a message...'
                type="text"
                onKeyDown={handleKeyDown}
                ref={messageInputRef}
            />
            <FontAwesomeIcon 
                icon={faPaw} 
                className={styles['send-message__button']}
                title="meow"
                onClick={sendMessage}
            />
        </div>
    )
}

export default SendMessageInput;