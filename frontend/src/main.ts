import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import piniaPluginPersistedstate from 'pinia-plugin-persistedstate' // 导入插件

// --- 添加 Element Plus ---
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import App from './App.vue'
import router from './router'

const app = createApp(App)
const pinia = createPinia()
pinia.use(piniaPluginPersistedstate) // 使用插件

app.use(pinia)
app.use(router)
app.use(ElementPlus) // --- 使用 Element Plus ---

app.mount('#app')
