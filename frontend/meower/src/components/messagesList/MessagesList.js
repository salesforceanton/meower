import { useSelector } from 'react-redux';

import styles from './MessagesList.module.css';
import { mainStateSelectors } from '../../store/main/selectors';
import Message from '../message/Message';

const MessagesList = () => {
    const feed = useSelector(mainStateSelectors.selectFeed);
    return (
        <div className={styles['messages-list__wrapper']}>
            {feed.map((e) =>
                <Message message={e} key={e.id} /> 
            )}
        </div>
    )
}

export default MessagesList;