import { createSelector } from "@reduxjs/toolkit";

const selectRootState = ({ main: output = {}}) => output;
const selectSelectedChat = createSelector(selectRootState, (state) => state.selectedChat);
const selectChatListData = createSelector(selectRootState, (state) => state.chatList);
const selectChatList = createSelector(
    selectChatListData,
    selectSelectedChat,
    (list, id) => list.map((e) => ({ ...e, isSelected: id === e.id }))
);
const selectSearchResults = createSelector(selectRootState, (state) => state.searchResults);
const selectSearchResultMessages = createSelector(
    selectSearchResults,
    (res) => res.messages
);
const selectFeed = createSelector(selectRootState, (state) => state.feed)

export const mainStateSelectors = {
    selectSelectedChat,
    selectChatList,
    selectSearchResultMessages,
    selectFeed
}