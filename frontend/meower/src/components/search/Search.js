import React, { useEffect } from 'react';

import SearchInput from '../common/searchInput/SearchInput';
import styles from './Search.module.css';

import { SEARCH_DEBOUNCE_TIME } from '../../store/main/constants';
import { mainStateActions } from '../../store/main/actions';
import { useDispatch, useSelector } from 'react-redux';
import { mainStateSelectors } from '../../store/main/selectors';

const Search = () => {
    const dispatch = useDispatch();
    
    const searchString = useSelector(mainStateSelectors.selectSearchString);

    const handleChange = (e) => dispatch(mainStateActions.setSearchString(e.target.value));
    const handleClear = (e) => dispatch(mainStateActions.clearSearch());

    useEffect(() => {
        let searchTimeout = setTimeout(() => {
            dispatch(mainStateActions.searchRequest())
        }, SEARCH_DEBOUNCE_TIME)
        
        
        return () => searchTimeout && clearTimeout(searchTimeout);
    }, [searchString, dispatch])

    return (
        <div className={styles['search-wrapper']}>
            <SearchInput 
                className={styles['search-input']}
                value={searchString}
                onChange={handleChange}
                onClear={handleClear}
            />
        </div>
    )
}

export default Search;