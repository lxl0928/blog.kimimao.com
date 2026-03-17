<template>
  <Layout>
    <div class="edit-box">
      <h1>{{ isEdit ? '编辑文章' : '写文章' }}</h1>
      <form @submit.prevent="handleSubmit">
        <input v-model="form.title" placeholder="标题" required />
        <input v-model="form.tags" placeholder="标签（逗号分隔）" />
        <textarea v-model="form.summary" placeholder="简介" rows="2"></textarea>
        <textarea v-model="form.content" placeholder="正文（Markdown）" rows="15" required></textarea>
        <p v-if="err" class="err">{{ err }}</p>
        <div class="btns">
          <button type="submit" :disabled="loading">{{ isEdit ? '保存' : '发布' }}</button>
          <router-link to="/blogs">取消</router-link>
        </div>
      </form>
    </div>
  </Layout>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Layout from '../../components/Layout.vue'
import { useUserStore } from '../../stores/user'
import { getBlog, createBlog, updateBlog } from '../../api/blog'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()
const id = computed(() => route.params.id)
const isEdit = computed(() => !!id.value)
const loading = ref(false)
const err = ref('')

const form = reactive({ title: '', content: '', summary: '', tags: '' })

onMounted(async () => {
  if (isEdit.value) {
    try {
      const res = await getBlog(id.value, false)
      const blog = res.data
      if (userStore.user?.id !== blog.user_id) {
        router.push('/blogs')
        return
      }
      Object.assign(form, {
        title: blog.title,
        content: blog.content,
        summary: blog.summary,
        tags: blog.tags,
      })
    } catch {
      router.push('/blogs')
    }
  }
})

const handleSubmit = async () => {
  err.value = ''
  loading.value = true
  try {
    if (isEdit.value) {
      await updateBlog(id.value, form)
    } else {
      await createBlog(form)
    }
    router.push('/blogs')
  } catch (e) {
    err.value = e?.message || '保存失败'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.edit-box { max-width: 720px; margin: 0 auto; padding: 24px; background: #fff; border-radius: 8px; }
.edit-box h1 { margin-bottom: 24px; }
.edit-box input, .edit-box textarea { width: 100%; padding: 12px; margin-bottom: 16px; border: 1px solid #ddd; border-radius: 6px; }
.edit-box .err { color: #e74c3c; margin-bottom: 12px; }
.edit-box .btns { display: flex; gap: 12px; }
.edit-box .btns button { padding: 10px 24px; background: #0066ff; color: #fff; border: none; border-radius: 6px; cursor: pointer; }
.edit-box .btns button:hover:not(:disabled) { background: #0052cc; }
</style>
