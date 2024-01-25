import { useSelector } from 'react-redux';

import ChatList from '../../components/chatList/ChatList';
import Search from '../../components/search/Search';
import SearchResults from '../../components/searchResults/SearchResults';

import styles from './FeedsSidebar.module.css';

import { mainStateSelectors } from '../../store/main/selectors';

const FeedsSidebar = () => {
    const searchString = useSelector(mainStateSelectors.selectSearchString);

    return (
        <div className={styles['feeds-sidebar__wrapper']}>
            <Search/>
            {searchString 
                ? <SearchResults/>
                : <ChatList/>
            }
        </div>
    )
}
export default FeedsSidebar;