import request from '@/utils/request'

const dataApi = {
    ListData: "/data/listData",
}

export function listData(params: any) {
    console.log(params)
    return request({
        url: dataApi.ListData,
        method: 'post',
        data: {
            "tab": params.tab
        }
    })
}