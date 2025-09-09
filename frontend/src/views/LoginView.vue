<template>
  <div class="login-container">
    <el-card class="login-card">
      <template #header>
        <div class="card-header">
          <h3>社区秒杀系统</h3>
        </div>
      </template>
      <el-tabs v-model="activeTab" stretch>
        <el-tab-pane label="登录" name="login">
          <el-form :model="form" ref="loginFormRef" @submit.prevent="handleLogin">
            <el-form-item prop="username">
              <el-input size="large" v-model="form.username" placeholder="用户名"></el-input>
            </el-form-item>
            <el-form-item prop="password">
              <el-input size="large" v-model="form.password" type="password" placeholder="密码" show-password></el-input>
            </el-form-item>
            <el-form-item>
              <el-button size="large" type="primary" native-type="submit" style="width: 100%">登 录</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>
        <el-tab-pane label="注册" name="register">
          <el-form :model="form" ref="registerFormRef" @submit.prevent="handleRegister">
            <el-form-item prop="username">
              <el-input size="large" v-model="form.username" placeholder="用户名"></el-input>
            </el-form-item>
            <el-form-item prop="password">
              <el-input size="large" v-model="form.password" type="password" placeholder="密码" show-password></el-input>
            </el-form-item>
            <el-form-item>
              <el-button size="large" type="primary" native-type="submit" style="width: 100%">注 册</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { ElMessage } from 'element-plus'

const router = useRouter()
const userStore = useUserStore()

const activeTab = ref('login')
const form = reactive({
  username: '',
  password: ''
})

const handleLogin = async () => {
  if (!form.username || !form.password) {
    ElMessage.warning('请输入用户名和密码');
    return;
  }
  try {
    await userStore.login(form)
    ElMessage.success('登录成功')
    router.push('/')
  } catch (error: any) {
    ElMessage.error(error.response?.data?.error || '登录失败')
  }
}

const handleRegister = async () => {
  if (!form.username || !form.password) {
    ElMessage.warning('请输入用户名和密码');
    return;
  }
  try {
    await userStore.register(form)
    ElMessage.success('注册成功，请登录')
    activeTab.value = 'login'
    form.password = ''
  } catch (error: any) {
    ElMessage.error(error.response?.data?.error || '注册失败')
  }
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 100%;
}

.login-card {
  width: 100%;
  max-width: 400px;
  min-width: 350px;
}

.card-header {
  text-align: center;
  font-size: 1.2em;
}
</style>