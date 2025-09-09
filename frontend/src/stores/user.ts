import { defineStore } from 'pinia'
import { ref } from 'vue'
import { loginAPI, registerAPI } from '@/api/user'
import type { AuthRequest } from '@/api/user'

export const useUserStore = defineStore(
  'user',
  () => {
    const token = ref('')

    // 登录 action
    const login = async (userData: AuthRequest) => {
      const res = await loginAPI(userData)
      token.value = res.token
    }

    // 注册 action
    const register = async (userData: AuthRequest) => {
      await registerAPI(userData)
    }

    // 退出登录 action
    const logout = () => {
      token.value = ''
    }

    return { token, login, register, logout }
  },
  {
    persist: true // 开启当前 store 的持久化
  }
)