<template>
  <el-card>
    <template #header>
      <div class="card-header">
        <span>进行中的秒杀活动</span>
      </div>
    </template>
    <el-table :data="activities" style="width: 100%" v-loading="loading" stripe border>
      <el-table-column prop="ID" label="活动ID" width="100" align="center" />
      <el-table-column prop="product_id" label="商品ID" width="100" align="center" />
      <el-table-column prop="seckill_price" label="秒杀价 (元)" width="150" align="center" />
      <el-table-column prop="stock" label="剩余库存" width="120" align="center" />
      <el-table-column prop="start_time" label="开始时间" :formatter="formatTime" align="center" />
      <el-table-column prop="end_time" label="结束时间" :formatter="formatTime" align="center" />
      <el-table-column label="操作" width="150" fixed="right" align="center">
        <template #default="scope">
          <el-button type="danger" size="small" @click="handleSeckill(scope.row.ID)">
            立即秒杀
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { getActivityListAPI, seckillAPI, type Activity } from '@/api/activity'
import { ElMessage } from 'element-plus'

const loading = ref(false)
const activities = ref<Activity[]>([])

const formatTime = (row: Activity, column: any, cellValue: string) => {
  if (!cellValue) return ''
  return new Date(cellValue).toLocaleString('zh-CN', { hour12: false })
}

const fetchActivities = async () => {
  loading.value = true
  try {
    const res = await getActivityListAPI()
    activities.value = res.activities || []
  } catch (error) {
    ElMessage.error('获取活动列表失败')
    console.error(error)
  } finally {
    loading.value = false
  }
}

const handleSeckill = async (activityId: number) => {
  try {
    const res = await seckillAPI(activityId)
    ElMessage.success(res.message || '抢购成功，订单处理中！')
    setTimeout(() => {
      fetchActivities()
    }, 500)
  } catch (error: any) {
    ElMessage.error(error.response?.data?.error || '秒杀失败，请重试')
  }
}

onMounted(() => {
  fetchActivities()
})
</script>

<style scoped>
.card-header {
  font-size: 18px;
  font-weight: bold;
}
</style>