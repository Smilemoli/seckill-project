import axios from 'axios'
import { useUserStore } from '@/stores/user'

// 创建 axios 实例
const service = axios.create({
    // 我们的请求路径会自己带上 /api，所以这里baseURL设为 / 即可
    baseURL: '/',
    timeout: 10000 // 请求超时时间
})

// 请求拦截器
service.interceptors.request.use(
    (config) => {
        // 在发送请求前做些什么
        // 比如，从 Pinia 中获取 token，然后添加到请求头里
        const userStore = useUserStore()
        if (userStore.token) {
            config.headers.Authorization = `Bearer ${userStore.token}`
        }
        return config
    },
    (error) => {
        // 对请求错误做些什么
        console.log(error) // for debug
        return Promise.reject(error)
    }
)

// 响应拦截器 (可选，但推荐)
service.interceptors.response.use(
    (response) => {
        // 对响应数据做点什么
        // 后端返回的数据在 response.data 中
        return response.data
    },
    (error) => {
        // 对响应错误做点什么
        console.log('err' + error) // for debug
        return Promise.reject(error)
    }
)

export default service