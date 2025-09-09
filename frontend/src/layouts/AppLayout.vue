<template>
  <el-container class="app-layout">
    <el-header class="app-header">
      <div class="header-content">
        <div class="logo">社区秒杀系统</div>
        <div class="user-info">
          <el-dropdown>
            <span class="el-dropdown-link">
              欢迎您
              <el-icon class="el-icon--right"><arrow-down /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="handleLogout">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </div>
    </el-header>
    <el-main class="app-main">
      <router-view />
    </el-main>
  </el-container>
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/user'
import { useRouter } from 'vue-router'
import { ArrowDown } from '@element-plus/icons-vue'

const userStore = useUserStore()
const router = useRouter()

const handleLogout = () => {
  userStore.logout()
  router.push('/login')
}
</script>

<style scoped>
.app-layout {
  min-height: 100vh;
}

.app-header {
  background-color: #ffffff;
  /* 改为白色背景 */
  border-bottom: 1px solid #e7e7e7;
  color: #333;
  position: fixed;
  width: 100%;
  top: 0;
  left: 0;
  z-index: 1000;
  display: flex;
  justify-content: center;
  /* 让内部的 header-content 居中 */
}

/* 新增 header-content 容器 */
.header-content {
  width: 100%;
  max-width: 1200px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo {
  font-size: 20px;
  font-weight: bold;
}

.el-dropdown-link {
  cursor: pointer;
  color: #333;
  display: flex;
  align-items: center;
}

.app-main {
  /*
    这里的 router-view 会被 HomeView.vue 替代。
    HomeView.vue 内部的最外层是 el-card，它是一个块级元素。
    我们将通过 router-view 的 :deep 选择器来设置它的宽度和居中。
  */
  padding-top: 80px;
  /* 留出 header 的空间 */
  padding-bottom: 40px;
}

/* 使用 :deep 选择器来为子路由（HomeView）的根元素设置样式 */
:deep(.el-card) {
  max-width: 1200px;
  width: 90%;
  margin: 0 auto;
  /* 水平居中 */
}
</style>