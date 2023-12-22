import React, { useState } from 'react';

import SearchInput from '../common/searchInput/SearchInput';
import styles from './Search.module.css';

const Search = () => {
    const [searchString, setSearchString] = useState('');

    const handleChange = (e) => setSearchString(e.target.value);
    const handleClear = (e) => setSearchString('');

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