import { mainSlice } from "./reducer";
import { makeHttpRequest } from "../service/actions";
import { getFeedRequestDefault, getSearchRequestDefault } from "./utils";
import { HTTP_METHOD, SERVICE_ENDPOINTS } from "../service/constants";
import { mainStateSelectors } from "./selectors";
import { ANON_CHAT_ID, ANON_CHAT_NAME } from "./constants";

const mainStateGenericActions = mainSlice.actions;

const searchRequestThunk = () => async(dispatch, getState) => {
    const searhString = mainStateSelectors.selectSearchString(getState());
    if (!searhString) {
        return;
    }
    dispatch(makeHttpRequest(
            { endpoint: SERVICE_ENDPOINTS.SEARCH, body: getSearchRequestDefault(searhString) },
            (res) => dispatch(mainStateGenericActions.setSearchResults(res)
        )
    ));
}

const selectChatThunk = (chatId) => async(dispatch) => {
    const successCallback = (res) => dispatch(
        mainStateGenericActions.setSelectedChat({ feed: res.body, chatId })
    );
    
    dispatch(makeHttpRequest(
        { 
            endpoint: SERVICE_ENDPOINTS.MESSAGES,
            body: getFeedRequestDefault(chatId) 
        },
        successCallback
    ));
}

const sendMessageThunk = (message) => async(dispatch, getState) => {
    const chatId = mainStateSelectors.selectSelectedChat(getState());

    dispatch(makeHttpRequest(
        { 
            endpoint: SERVICE_ENDPOINTS.MESSAGES, 
            method: HTTP_METHOD.POST,
            body: { body: message, chatId },
        }
    ));
}

const getChatListThunk = (payload) => (dispatch) => {
    // TODO: Http request should be here
    const mockChatsData = [
        {
            name: ANON_CHAT_NAME,
            id: ANON_CHAT_ID
        }
    ];
    dispatch(mainStateGenericActions.setChatList(mockChatsData)); 
}

export const mainStateActions = {
    ...mainStateGenericActions,
    searchRequest: searchRequestThunk,
    selectChat: selectChatThunk,
    sendMessage: sendMessageThunk,
    getChatList: getChatListThunk
}