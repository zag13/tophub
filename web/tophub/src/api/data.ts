import request from '@/utils/request'

const dataApi = {
    GetData: "",
}

export function getData(param: any) {
    return request({
        url: dataApi.GetData,
        method: 'get',
        params: {
            "tab": param.tab,
        }
    })
}