import Search from '../search/Search';
import styles from './FeedsSidebar.module.css';

const FeedsSidebar = () => {
    return (
        <div className={styles['feeds-sidebar__wrapper']}>
            <Search/>
        </div>
    )
}
export default FeedsSidebar;