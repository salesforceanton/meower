import { commonSlice } from "./reducer";

const commonStateGenericActions = commonSlice.actions;

export const commonStateActions = {
    ...commonStateGenericActions
}