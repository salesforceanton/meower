import styles from './AppHeader.module.css';

const AppHeader = (props) => {
    return (
        <div className={`${props.className} ${styles.header}`}>
            <h2>Meower</h2>
        </div>
    )
}

export default AppHeader;