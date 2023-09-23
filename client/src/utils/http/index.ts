import axios, { AxiosRequestConfig } from "axios";

const defaultConfig: AxiosRequestConfig = {
  timeout: 5000,
  headers: {
    "Content-Type": "application/json",
  },
};

const request = axios.create(defaultConfig);

const errorHandler = (error: any) => {
  return Promise.reject(error);
};

request.interceptors.request.use((config) => {
  return config;
}, errorHandler);

request.interceptors.response.use((response) => {
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
}, errorHandler);

export default request;
