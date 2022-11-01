import axios from "axios";

const request = axios.create({
    baseURL: '',
    timeout: 3000
})

const errorHandler = (error: any) => {
    return Promise.reject(error)
}

request.interceptors.request.use((config) => {
    return config
}, errorHandler)

request.interceptors.response.use((response) => {
    return response.data
}, errorHandler)

export default request