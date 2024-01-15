import { configureStore } from "@reduxjs/toolkit";

import { mainSlice } from "./main/reducer";
import { commonSlice } from "./common/reducer";

export default configureStore({
    reducer: {
        main: mainSlice.reducer,
        common: commonSlice.reducer
    }
});