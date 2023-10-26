import axios, {AxiosRequestConfig} from 'axios';
import {Configuration, TopHubServiceApi} from '@/openapi';

const defaultConfig: AxiosRequestConfig = {
    timeout: 5000,
    headers: {
        "Content-Type": "application/json",
    },
    baseURL: window.location.protocol + "//" + window.location.host + "/api"
};

const customAxiosInstance = axios.create(defaultConfig);

customAxiosInstance.interceptors.request.use(config => {
    return config;
}, error => {
    return Promise.reject(error);
});

customAxiosInstance.interceptors.response.use((response) => {
    if (response.status !== 200) {
        const err = {
            code: response.data.code,
            message: response.data.message,
            toString: function () {
                return response.data.message;
            },
        };
        return Promise.reject(err);
    }
    return response;
}, error => {
    return Promise.reject(error);
});

// 创建继承自 TopHubServiceApi 的新类
class CustomApi extends TopHubServiceApi {
    constructor(configuration?: Configuration, basePath?: string) {
        // 将自定义的 Axios 实例传递给父类
        super(configuration, basePath, customAxiosInstance);
    }
}

export const customApiOOP = new CustomApi();
