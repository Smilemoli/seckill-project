import request from './request'
import type { AxiosResponse } from 'axios'

// 定义请求和响应的数据类型
export interface AuthRequest {
    username: string
    password: string
}

interface LoginResponse {
    message: string
    token: string
}

// 登录 API
export const loginAPI = (data: AuthRequest): Promise<LoginResponse> => {
    return request({
        url: '/api/v1/login',
        method: 'POST',
        data
    })
}

// 注册 API
export const registerAPI = (data: AuthRequest) => {
    return request({
        url: '/api/v1/register',
        method: 'POST',
        data
    })
}