<template>
  <Layout>
    <div class="blog-list">
      <h1 class="page-title">文章</h1>
      <div v-if="loading" class="loading">加载中...</div>
      <div v-else-if="list.length === 0" class="empty">暂无文章</div>
      <div v-else class="card-grid">
        <article v-for="b in list" :key="b.id" class="card">
          <router-link :to="`/blogs/${b.id}`" class="card-link">
            <h3 class="card-title">{{ b.title }}</h3>
            <p class="card-summary">{{ b.summary || b.content?.slice(0, 150) + (b.content?.length > 150 ? '...' : '') }}</p>
            <div class="card-meta">
              <span class="author">{{ b.author_name }}</span>
              <span class="dot">·</span>
              <span class="date">{{ formatDate(b.created_at) }}</span>
            </div>
          </router-link>
        </article>
      </div>
      <div v-if="total > size" class="pagination">
        <button class="page-btn" :disabled="page <= 1" @click="page--">上一页</button>
        <span class="page-info">{{ page }} / {{ Math.ceil(total / size) }}</span>
        <button class="page-btn" :disabled="page >= Math.ceil(total / size)" @click="page++">下一页</button>
      </div>
    </div>
  </Layout>
</template>

<script setup>
import { ref, watch } from 'vue'
import Layout from '../../components/Layout.vue'
import { listBlogs } from '../../api/blog'

const list = ref([])
const total = ref(0)
const page = ref(1)
const size = 10
const loading = ref(true)

const fetch = async () => {
  loading.value = true
  try {
    const res = await listBlogs({ page: page.value, size })
    list.value = res.data.list || []
    total.value = res.data.total || 0
  } finally {
    loading.value = false
  }
}

watch(page, fetch, { immediate: true })

const formatDate = (d) => d ? new Date(d).toLocaleDateString('zh-CN') : ''
</script>

<style scoped>
.blog-list { padding: 0; }
.page-title { font-size: 22px; font-weight: 600; margin-bottom: 24px; color: #1a1a1a; }
.loading, .empty { padding: 48px 0; text-align: center; color: #999; }
.card-grid { display: flex; flex-direction: column; gap: 16px; }
.card {
  background: #fff;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 1px 3px rgba(0,0,0,.06);
  transition: box-shadow .2s;
}
.card:hover { box-shadow: 0 2px 12px rgba(0,0,0,.08); }
.card-link { display: block; padding: 24px; color: inherit; }
.card-link:hover { color: inherit; }
.card-title {
  font-size: 18px;
  font-weight: 600;
  color: #1a1a1a;
  margin-bottom: 12px;
  line-height: 1.4;
}
.card-title:hover { color: #0066ff; }
.card-summary {
  font-size: 15px;
  color: #646464;
  line-height: 1.6;
  margin-bottom: 16px;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
.card-meta {
  font-size: 14px;
  color: #999;
}
.card-meta .dot { margin: 0 6px; }
.pagination {
  margin-top: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 16px;
}
.page-btn {
  padding: 8px 20px;
  border: 1px solid #e0e0e0;
  background: #fff;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  color: #646464;
}
.page-btn:hover:not(:disabled) { border-color: #0066ff; color: #0066ff; }
.page-btn:disabled { opacity: 0.5; cursor: not-allowed; }
.page-info { font-size: 14px; color: #999; }
</style>
