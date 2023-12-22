import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faXmark } from '@fortawesome/free-solid-svg-icons' 

import styles from './SearchInput.module.css';

const SearchInput = (props) => {
    return (
        <div className={styles['search-input__container']}>
            <input 
                className={`${styles['search-input']} ${props.className}`} 
                placeholder='Search...'
                type="text"
                value={props.value}
                onChange={props.onChange}
            />
            {props.value && <FontAwesomeIcon 
                icon={faXmark} 
                className={styles['clear-input__button']}
                title="Clear input button"
                onClick={props.onClear}
            />}
        </div>
    )
}

export default SearchInput;