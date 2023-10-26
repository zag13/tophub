import axios, {AxiosRequestConfig} from 'axios';
import {Configuration, TopHubServiceApiFactory} from '@/openapi';

const defaultConfig: AxiosRequestConfig = {
    timeout: 5000,
    headers: {
        "Content-Type": "application/json",
    },
    baseURL: window.location.protocol + "//" + window.location.host +"/api"
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
    return response.data;
}, error => {
    return Promise.reject(error);
});

export const customApiFactory = (configuration?: Configuration, basePath?: string) => {
    return TopHubServiceApiFactory(configuration, basePath, customAxiosInstance);
};
