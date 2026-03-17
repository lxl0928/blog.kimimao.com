<template>
  <Layout>
    <div v-if="loading" class="detail">加载中...</div>
    <div v-else-if="blog" class="detail">
      <h1>{{ blog.title }}</h1>
      <p class="meta">{{ blog.author_name }} · {{ formatDate(blog.created_at) }} · {{ blog.view_count }} 阅读</p>
      <div class="tags" v-if="blog.tags">{{ blog.tags }}</div>
      <div class="content" v-html="renderedContent"></div>
      <div class="actions">
        <button v-if="userStore.isLoggedIn" @click="toggleLike">{{ liked ? '取消点赞' : '点赞' }}</button>
        <button v-if="userStore.isLoggedIn" @click="toggleFavorite">{{ favorited ? '取消收藏' : '收藏' }}</button>
        <button @click="copyLink">分享链接</button>
      </div>
      <div class="comments">
        <h3>评论</h3>
        <form v-if="userStore.isLoggedIn" @submit.prevent="submitComment">
          <textarea v-model="commentText" placeholder="写下评论..." rows="3"></textarea>
          <button type="submit">发表</button>
        </form>
        <ul v-if="comments.length">
          <li v-for="c in comments" :key="c.id">
            <strong>{{ c.username }}</strong>: {{ c.content }}
            <span class="time">{{ formatDate(c.created_at) }}</span>
          </li>
        </ul>
        <p v-else>暂无评论</p>
      </div>
    </div>
  </Layout>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { marked } from 'marked'
import Layout from '../../components/Layout.vue'
import { useUserStore } from '../../stores/user'
import { getBlog, likeBlog, unlikeBlog, favoriteBlog, unfavoriteBlog } from '../../api/blog'
import { listComments, createComment } from '../../api/comment'

const route = useRoute()
const userStore = useUserStore()
const blog = ref(null)
const loading = ref(true)
const comments = ref([])
const commentText = ref('')
const liked = ref(false)
const favorited = ref(false)

const renderedContent = computed(() =>
  blog.value?.content ? marked(blog.value.content) : ''
)

const fetch = async () => {
  const id = route.params.id
  const res = await getBlog(id)
  blog.value = res.data
}

const fetchComments = async () => {
  const res = await listComments(route.params.id)
  comments.value = res.data || []
}

onMounted(async () => {
  try {
    await fetch()
    await fetchComments()
  } finally {
    loading.value = false
  }
})

const submitComment = async () => {
  if (!commentText.value.trim()) return
  await createComment(route.params.id, { content: commentText.value })
  commentText.value = ''
  fetchComments()
}

const toggleLike = async () => {
  try {
    if (liked.value) await unlikeBlog(route.params.id)
    else await likeBlog(route.params.id)
    liked.value = !liked.value
  } catch (_) {}
}

const toggleFavorite = async () => {
  try {
    if (favorited.value) await unfavoriteBlog(route.params.id)
    else await favoriteBlog(route.params.id)
    favorited.value = !favorited.value
  } catch (_) {}
}

const copyLink = () => {
  navigator.clipboard.writeText(window.location.href)
  alert('链接已复制')
}

const formatDate = (d) => d ? new Date(d).toLocaleString('zh-CN') : ''
</script>

<style scoped>
.detail { padding: 24px; background: #fff; border-radius: 8px; }
.detail h1 { margin-bottom: 12px; }
.detail .meta { color: #888; margin-bottom: 16px; }
.detail .tags { font-size: 14px; color: #0066ff; margin-bottom: 16px; }
.detail .content { line-height: 1.8; margin-bottom: 24px; }
.detail .content :deep(pre) { background: #f6f6f6; padding: 12px; overflow-x: auto; border-radius: 6px; }
.detail .actions { display: flex; gap: 12px; margin-bottom: 32px; }
.detail .actions button {
  padding: 8px 20px;
  border: 1px solid #e0e0e0;
  background: #fff;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  color: #646464;
}
.detail .actions button:hover { border-color: #0066ff; color: #0066ff; }
.comments button[type="submit"] { background: #0066ff; color: #fff; border: none; padding: 8px 20px; border-radius: 6px; cursor: pointer; }
.comments button[type="submit"]:hover { background: #0052cc; }
.comments h3 { margin-bottom: 16px; }
.comments textarea { width: 100%; padding: 12px; margin-bottom: 8px; border: 1px solid #ddd; border-radius: 6px; }
.comments ul { list-style: none; }
.comments li { padding: 12px; border-bottom: 1px solid #eee; }
.comments .time { color: #888; font-size: 12px; margin-left: 8px; }
</style>
