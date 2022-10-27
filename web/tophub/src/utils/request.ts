import axios from "axios";
import {ElMessage} from "element-plus";

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

request.interceptors.response.use((res) => {
    const {data: result, meta: {msg, status}} = res.data
    ElMessage({message: msg, type: "success"})
    return result
}, errorHandler)

export default request