export const HTTP_METHOD = Object.freeze({
    POST: 'POST',
    GET: 'GET'
});

export const DEFAULT_HEADERS = {
    'Content-Type': 'application/json',
};

// TODO: add getting this endpoint from env
export const BACKEND_URL = 'localhost:8080';

export const SERVICE_ENDPOINTS = Object.freeze({
    SEARCH: `http://${BACKEND_URL}/search`,
    MESSAGES: `http://${BACKEND_URL}/meows`,
    WEBSOKET: `ws://${BACKEND_URL}/pusher`
});
