import axios from 'axios';

import { commonStateActions } from "../common/actions";
import { HTTP_METHOD, DEFAULT_HEADERS } from "./constants";

export const makeHttpRequest = ({ endpoint, method, headers = DEFAULT_HEADERS, body }, successCallback, errorCallback) => async(dispatch) => {
    dispatch(commonStateActions.setError(null));
    dispatch(commonStateActions.setIsLoading(true));

    // try {
    //     const response = await fetch(endpoint, {
    //         mode: 'no-cors',
    //         method: method ?? HTTP_METHOD.GET,
    //         headers: headers ?? DEFAULT_HEADERS,
    //         body: body ? JSON.stringify(body) : null
    //     });

    //     const responseData = await response.json();
    //     successCallback(responseData);
    // } catch (error) {
    //     dispatch(commonStateActions.setError(error));
    //     errorCallback && errorCallback(error);
    // } finally {
    //     dispatch(commonStateActions.setIsLoading(false));
    // }

    let callout;

    switch (method) {
        case HTTP_METHOD.POST:
            callout = axios.post(endpoint, body, { headers })
            break
        default:
            callout = axios.get(endpoint, { headers, params: body })
            break
    }

    callout
        .then((res) => {
            successCallback && successCallback(res.data)
        })
        .catch((err) => errorCallback && errorCallback(err))
        .finally(() => {
            dispatch(commonStateActions.setIsLoading(false))
        });
}