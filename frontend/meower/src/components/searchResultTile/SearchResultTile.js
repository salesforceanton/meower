import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPaw } from '@fortawesome/free-solid-svg-icons';

import { useDispatch } from 'react-redux';

import styles from './SearchResultTile.module.css';
import { mainStateActions } from '../../store/main/actions';
import { ANON_CHAT_ID, ANON_CHAT_SENDER_NAME } from '../../store/main/constants';

const SearchResultTile = (props) => {
    const dispatch = useDispatch();
    
    const { sender, chatId, body, createdAt } = props.result;

    const createAtDatetime = new Date(createdAt);
    const formattedTime = createAtDatetime.toLocaleTimeString();

    const handleResultSelect = () => dispatch(mainStateActions.selectChat(chatId || ANON_CHAT_ID));

    return (
        <div className={`${styles['result-tile__wrapper']} ${props.className}`} onClick={handleResultSelect}>
            <div className={styles['tile-avatar']}>
                <FontAwesomeIcon icon={faPaw} title="meow"/>
            </div>
            <div className={styles['tile-main-content']}>
                <p className={styles['sender-name']}>{sender || ANON_CHAT_SENDER_NAME}</p>
                <p className={styles['message']}>{body}</p>
            </div>
            <div className={styles['message-time__container']}>
                <p className={styles['message-time']}>{formattedTime}</p>
            </div>
        </div>
    )
}

export default SearchResultTile;