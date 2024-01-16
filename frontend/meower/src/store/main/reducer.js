import { createSlice } from '@reduxjs/toolkit';

const initialState = {
    searchString: '',
    selectedChat: null,
    chatList: [],
    searchResults: {},
    feed: [],
}

export const mainSlice = createSlice({
    name: 'main',
    initialState,
    reducers: {
        setChatList: (state, action) => {
            state.chatList = action.payload;
        },
        setSelectedChat: (state, action) => {
            state.selectedChat = action.payload.chatId;
            state.feed = action.payload.feed;
        },
        setSearchString: (state, action) => {
            state.searchString = action.payload;
        },
        clearSearch: (state) => {
            state.searchString = '';
            state.searchResults = {};
        },
        setSearchResults: (state, action) => {
            state.searchResults = action.payload;
        }
    }
});