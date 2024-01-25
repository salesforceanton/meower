import React from 'react';

import styles from './Message.module.css';

const Message = (props) => {
    const { body, time } = props.message;
    return (
        <React.Fragment>
            
            <div className={`${styles['message']} ${styles['droplet']}`}>
                <div className={styles['message__text']}>
                    <div className={styles['message__text__content']}>{body}
                        <div className={styles['message__time']}>{time}</div>
                    </div>
                </div>
            </div>
            
            <svg height="0" width="0">
                <defs>
                    <clipPath id="droplet">
                        <path d="M 10,0 A 10,10 0 0 1 0,10 H 16 V 0 Z"/>
                    </clipPath>
                </defs>
            </svg>
        </React.Fragment>
    )
}

export default Message;