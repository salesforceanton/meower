import { useSelector } from 'react-redux';
import styles from './SearchResults.module.css';
import { mainStateSelectors } from '../../store/main/selectors';
import SearchResultTile from '../searchResultTile/SearchResultTile';

const SearchResults = (props) => {
    const messages = useSelector(mainStateSelectors.selectSearchResultMessages)

    return (
        <div className={`${styles['results-list__wrapper']} ${props.className}`}>
            <div className={styles['messages-header']}>
                <p>Found {messages.lenght} messages</p>
            </div>
            {messages.map((e) => 
                <SearchResultTile 
                    result={e} 
                    className={styles['result-tile']}
                    key={e.id}
                />
            )}
        </div>
    )
}

export default SearchResults;