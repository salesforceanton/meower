import { createSelector } from "@reduxjs/toolkit";

const selectMainState = ({ main: output = {}}) => output;
const selectSelectedChat = createSelector(selectMainState, (state) => state.selectedChat);
const selectIsNoSelectedChat = createSelector(selectSelectedChat, (state) => !state);
const selectChatListData = createSelector(selectMainState, (state) => state.chatList);
const selectChatList = createSelector(
    selectChatListData,
    selectSelectedChat,
    (list, id) => list.map((e) => ({ ...e, isSelected: id === e.id }))
);
const selectSearchResults = createSelector(selectMainState, (state) => state.searchResults);
const selectSearchString = createSelector(selectMainState, (state) => state.searchString);

const selectSearchResultMessages = createSelector(
    selectSearchResults,
    (res) => res.messages || []
);
const selectShowSearchResults = createSelector(
    selectSearchResultMessages, (messages) => !!messages.length
);
const selectFeed = createSelector(selectMainState, (state) => state.feed);

export const mainStateSelectors = {
    selectSelectedChat,
    selectIsNoSelectedChat,
    selectChatList,
    selectSearchResultMessages,
    selectShowSearchResults,
    selectSearchString,
    selectFeed
}