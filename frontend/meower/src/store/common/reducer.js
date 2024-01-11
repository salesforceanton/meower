import { createSlice } from '@reduxjs/toolkit';

const initialState = {
    error: null,
    isLoading: false,
}

export const commonSlice = createSlice({
    name: 'common',
    initialState,
    reducers: {
        setError: (state, action) => {
            state.error = action.payload
        },
        setIsLoading: (state, action) => {
            state.isLoading = action.payload
        }
    }
});
