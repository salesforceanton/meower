export const getSearchRequestDefault = (query) => ({
    skip: 0,
    take: 10,
    query
});

export const getFeedRequestDefault = (chatId) => ({
    skip: 0,
    take: 10,
    chatId
});