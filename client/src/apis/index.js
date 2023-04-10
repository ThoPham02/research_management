import axios from 'axios';
import store from '../store/store';

const client = axios.create({
    baseURL: "localhost:8080",
    //baseURL: "https://nckh-be.onrender.com",
    headers: {
        'Accept': "application/json",
        'Content-Type': "application/json",
    }
})

client.interceptors.request.use(config => {
    const authToken = store.getState().login.token.accessToken;
    if (authToken) {
        config.headers['Authorization'] = "Bearer " + authToken
    }
    return config
})

export default client