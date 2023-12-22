import React from 'react';
import { Route, Routes, Navigate } from "react-router-dom";

import './App.css';
import MainPage from './pages/mainPage/MainPage';

function App() {
    return (
        <React.Fragment>
            <Routes>
                <Route exact path="/" element={<Navigate replace to="/main" />} />
                <Route path="/main" element={<MainPage/>} />
            </Routes>
        </React.Fragment> 
    );
}

export default App;
