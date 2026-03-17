<template>
  <AdminLayout>
    <div class="blog-list">
      <h1>文章管理</h1>
      <div class="stats" v-if="stats">
        <div class="stat">文章总数: {{ stats.total_count }}</div>
        <div class="stat">今日新增: {{ stats.today_new_count }}</div>
      </div>
      <div v-if="stats?.rank_list?.length" class="rank">
        <h3>阅读排行</h3>
        <ul>
          <li v-for="(b, i) in stats.rank_list" :key="b.id">
            {{ i + 1 }}. {{ b.title }} ({{ b.view_count }})
          </li>
        </ul>
      </div>
      <div v-if="loading">加载中...</div>
      <table v-else class="table">
        <thead>
          <tr><th>ID</th><th>标题</th><th>作者</th><th>阅读</th><th>时间</th></tr>
        </thead>
        <tbody>
          <tr v-for="b in list" :key="b.id">
            <td>{{ b.id }}</td>
            <td><router-link :to="`/blogs/${b.id}`">{{ b.title }}</router-link></td>
            <td>{{ b.author_name }}</td>
            <td>{{ b.view_count }}</td>
            <td>{{ formatDate(b.created_at) }}</td>
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
import { listAdminBlogs, getBlogStats } from '../../api/admin'

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
    const [blogsRes, statsRes] = await Promise.all([
      listAdminBlogs({ page: page.value, size }),
      getBlogStats(),
    ])
    list.value = blogsRes.data?.list || []
    total.value = blogsRes.data?.total || 0
    stats.value = statsRes.data
  } finally {
    loading.value = false
  }
}

watch(page, fetch, { immediate: true })

const formatDate = (d) => d ? new Date(d).toLocaleString('zh-CN') : ''
</script>

<style scoped>
.blog-list h1 { margin-bottom: 24px; }
.stats { display: flex; gap: 24px; margin-bottom: 24px; }
.stat { padding: 16px 24px; background: #fff; border-radius: 8px; }
.rank { margin-bottom: 24px; padding: 16px; background: #fff; border-radius: 8px; }
.rank ul { list-style: none; }
.rank li { padding: 8px 0; }
.table { width: 100%; padding: 24px; background: #fff; border-radius: 8px; }
.table th, .table td { padding: 12px; text-align: left; border-bottom: 1px solid #eee; }
.pagination { margin-top: 24px; display: flex; gap: 16px; align-items: center; }
</style>
