<template>
  <Layout>
    <div class="form-box">
      <h1>登录</h1>
      <form @submit.prevent="handleLogin">
        <input v-model="form.username" placeholder="用户名" required />
        <input v-model="form.password" type="password" placeholder="密码" required />
        <p v-if="err" class="err">{{ err }}</p>
        <button type="submit" :disabled="loading">登录</button>
      </form>
      <p>还没有账号？<router-link to="/register">注册</router-link></p>
    </div>
  </Layout>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import Layout from '../../components/Layout.vue'
import { useUserStore } from '../../stores/user'

const router = useRouter()
const userStore = useUserStore()
const loading = ref(false)
const err = ref('')

const form = reactive({ username: '', password: '' })

const handleLogin = async () => {
  err.value = ''
  loading.value = true
  try {
    await userStore.login(form.username, form.password)
    router.push('/')
  } catch (e) {
    err.value = e?.message || '登录失败'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.form-box { max-width: 360px; margin: 60px auto; padding: 32px; background: #fff; border-radius: 8px; box-shadow: 0 2px 12px rgba(0,0,0,.08); }
.form-box h1 { margin-bottom: 24px; font-size: 24px; }
.form-box input { width: 100%; padding: 12px; margin-bottom: 16px; border: 1px solid #ddd; border-radius: 6px; }
.form-box button { width: 100%; padding: 12px; background: #0066ff; color: #fff; border: none; border-radius: 6px; cursor: pointer; }
.form-box button:hover:not(:disabled) { background: #0052cc; }
.form-box button:disabled { opacity: 0.6; cursor: not-allowed; }
.form-box .err { color: #e74c3c; margin-bottom: 12px; font-size: 14px; }
.form-box p { margin-top: 16px; text-align: center; color: #666; }
</style>
