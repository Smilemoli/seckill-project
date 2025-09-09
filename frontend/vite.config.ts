import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  // 配置 Vite 代理解决跨域
  server: {
    proxy: {
      // 将所有 /api 开头的请求代理到 http://localhost:8080
      '/api': {
        target: 'http://localhost:8080', // 你的 Go 后端服务地址
        changeOrigin: true, // 必须设置为 true
      }
    }
  }
})