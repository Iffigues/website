
import { configureStore } from '@reduxjs/toolkit';
import dataSlice from '../reducer/reducer';

const store = configureStore({
    reducer: dataSlice,
});

export {store};
