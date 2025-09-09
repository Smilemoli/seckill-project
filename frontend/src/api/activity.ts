import request from './request'

// 定义单个活动的数据结构，要和后端返回的 JSON 字段匹配
export interface Activity {
    ID: number
    product_id: number
    seckill_price: number
    stock: number
    start_time: string
    end_time: string
}

// 定义 API 响应的数据结构
interface ActivityListResponse {
    activities: Activity[]
}

// 获取活动列表的 API 函数
export const getActivityListAPI = (): Promise<ActivityListResponse> => {
    return request({
        url: '/api/v1/activities',
        method: 'GET'
    })
}

// 秒杀请求的 API 函数
// 它接收一个活动 ID，并且后端成功时只返回一个 message，没有复杂的 data
export const seckillAPI = (activityId: number) => {
    return request({
        url: `/api/v1/seckill/${activityId}`, // 使用模板字符串拼接 URL
        method: 'POST'
    })
}