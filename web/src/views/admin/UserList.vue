<template>
  <AdminLayout>
    <div class="user-list">
      <h1>用户管理</h1>
      <div class="stats" v-if="stats">
        <div class="stat">用户总数: {{ stats.total_count }}</div>
        <div class="stat">今日新增: {{ stats.today_new_count }}</div>
        <div class="stat">今日活跃: {{ stats.today_active_count }}</div>
      </div>
      <div v-if="loading">加载中...</div>
      <table v-else class="table">
        <thead>
          <tr><th>ID</th><th>用户名</th><th>邮箱</th><th>角色</th><th>注册时间</th></tr>
        </thead>
        <tbody>
          <tr v-for="u in list" :key="u.id">
            <td>{{ u.id }}</td>
            <td>{{ u.username }}</td>
            <td>{{ u.email }}</td>
            <td>{{ u.role }}</td>
            <td>{{ formatDate(u.created_at) }}</td>
          </tr>
        </tbody>
      </table>
      <div class="pagination">
        <button :disabled="page <= 1" @click="page--">上一页</button>
        <span>{{ page }} / {{ totalPages }}</span>
        <button :disabled="page >= totalPages" @click="page++">下一页</button>
      </div>
    </div>
  </AdminLayout>
</template>

<script setup>
import { ref, watch, computed } from 'vue'
import AdminLayout from './AdminLayout.vue'
import { listUsers, getUserStats } from '../../api/admin'

const list = ref([])
const stats = ref(null)
const total = ref(0)
const page = ref(1)
const size = 10
const loading = ref(true)

const totalPages = computed(() => Math.ceil(total.value / size) || 1)

const fetch = async () => {
  loading.value = true
  try {
    const [usersRes, statsRes] = await Promise.all([
      listUsers({ page: page.value, size }),
      getUserStats(),
    ])
    list.value = usersRes.data?.list || []
    total.value = usersRes.data?.total || 0
    stats.value = statsRes.data
  } finally {
    loading.value = false
  }
}

watch(page, fetch, { immediate: true })

const formatDate = (d) => d ? new Date(d).toLocaleString('zh-CN') : ''
</script>

<style scoped>
.user-list h1 { margin-bottom: 24px; }
.stats { display: flex; gap: 24px; margin-bottom: 24px; }
.stat { padding: 16px 24px; background: #fff; border-radius: 8px; }
.table { width: 100%; padding: 24px; background: #fff; border-radius: 8px; }
.table th, .table td { padding: 12px; text-align: left; border-bottom: 1px solid #eee; }
.pagination { margin-top: 24px; display: flex; gap: 16px; align-items: center; }
</style>
