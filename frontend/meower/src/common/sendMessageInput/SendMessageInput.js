import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPaw } from '@fortawesome/free-solid-svg-icons';

import { useState } from 'react';
import styles from './SendMessageInput.module.css';
import { KeyCodes } from '../utils/Utils';

const SendMessageInput = (props) => {
    const [ message, setMessage ] = useState('');

    const handleKeyDown = (e) => {
        if (e.key === KeyCodes.ENTER) {
            sendMessage();
        }
    }
    const handleChange = (e) => setMessage(e.target.value);
    const sendMessage = () => {
        props.onSend(message)
        setMessage('')
    };

    return (
        <div className={styles['send-message-input__container']}>
            <input 
                className={`${styles['send-message-input']} ${props.className}`} 
                placeholder='Write a message...'
                type="text"
                onKeyDown={handleKeyDown}
                onChange={handleChange}
                value={message}
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