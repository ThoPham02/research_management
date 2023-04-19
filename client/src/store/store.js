import { configureStore } from '@reduxjs/toolkit'
import { loginReducer } from '../pages/Login/LoginSlice'

const store = configureStore({
  reducer: {
    login: loginReducer
  }
})

export default store