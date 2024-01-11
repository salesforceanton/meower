import { commonStateActions } from "../common/actions";
import { HTTP_METHOD, DEFAULT_HEADERS } from "./constants";

export const makeHttpRequest = ({ endpoint, method, headers, body }, successCallback, errorCallback) => async(dispatch) => {
    dispatch(commonStateActions.setError(null));
    dispatch(commonStateActions.setIsLoading(true));

    try {
        const response = await fetch(endpoint, {
            method: method ?? HTTP_METHOD.GET,
            headers: headers ?? DEFAULT_HEADERS,
            body: body ? JSON.stringify(body) : null
        });

        const responseData = await response.json();
        successCallback(responseData);
    } catch (error) {
        dispatch(commonStateActions.setError(error));
        errorCallback && errorCallback(error);
    } finally {
        dispatch(commonStateActions.setIsLoading(false));
    }
}