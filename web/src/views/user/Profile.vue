<template>
  <Layout>
    <div class="profile">
      <h1>个人中心</h1>
      <div v-if="mode === 'view'" class="card">
        <p><strong>用户名</strong>: {{ user?.username }}</p>
        <p><strong>真实姓名</strong>: {{ user?.real_name || '-' }}</p>
        <p><strong>邮箱</strong>: {{ user?.email || '-' }}</p>
        <p><strong>手机</strong>: {{ user?.phone || '-' }}</p>
        <p><strong>简介</strong>: {{ user?.intro || '-' }}</p>
        <p><strong>头像</strong>: {{ user?.avatar_url || '-' }}</p>
        <div class="actions">
          <button @click="mode = 'edit'">编辑资料</button>
          <button @click="mode = 'pwd'">修改密码</button>
        </div>
      </div>
      <form v-else-if="mode === 'edit'" @submit.prevent="saveProfile" class="card">
        <input v-model="editForm.real_name" placeholder="真实姓名" />
        <input v-model="editForm.email" type="email" placeholder="邮箱" />
        <input v-model="editForm.phone" placeholder="手机" />
        <textarea v-model="editForm.intro" placeholder="个人简介" rows="3"></textarea>
        <input v-model="editForm.avatar_url" placeholder="头像URL" />
        <div class="actions">
          <button type="submit">保存</button>
          <button type="button" @click="mode = 'view'">取消</button>
        </div>
      </form>
      <form v-else @submit.prevent="savePassword" class="card">
        <input v-model="pwdForm.old_password" type="password" placeholder="原密码" required />
        <input v-model="pwdForm.new_password" type="password" placeholder="新密码" required />
        <div class="actions">
          <button type="submit">修改密码</button>
          <button type="button" @click="mode = 'view'">取消</button>
        </div>
      </form>
    </div>
  </Layout>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import Layout from '../../components/Layout.vue'
import { useUserStore } from '../../stores/user'
import { getProfile, updateProfile, changePassword } from '../../api/user'

const userStore = useUserStore()
const user = ref(null)
const mode = ref('view')
const editForm = reactive({ real_name: '', email: '', phone: '', intro: '', avatar_url: '' })
const pwdForm = reactive({ old_password: '', new_password: '' })

onMounted(async () => {
  const res = await getProfile()
  user.value = res.data
  Object.assign(editForm, {
    real_name: user.value.real_name,
    email: user.value.email,
    phone: user.value.phone,
    intro: user.value.intro,
    avatar_url: user.value.avatar_url,
  })
})

const saveProfile = async () => {
  await updateProfile(editForm)
  user.value = { ...user.value, ...editForm }
  mode.value = 'view'
}

const savePassword = async () => {
  await changePassword(pwdForm)
  pwdForm.old_password = ''
  pwdForm.new_password = ''
  mode.value = 'view'
}
</script>

<style scoped>
.profile h1 { margin-bottom: 24px; }
.card { padding: 24px; background: #fff; border-radius: 8px; max-width: 480px; }
.card p { margin-bottom: 12px; }
.card input, .card textarea { width: 100%; padding: 12px; margin-bottom: 16px; border: 1px solid #ddd; border-radius: 6px; }
.card .actions { display: flex; gap: 12px; margin-top: 16px; }
.card .actions button { padding: 10px 20px; cursor: pointer; }
</style>
