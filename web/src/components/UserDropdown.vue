<template>
  <div class="user-dropdown" ref="dropdownRef">
    <button class="avatar-btn" @click="open = !open">
      <img v-if="user?.avatar_url" :src="user.avatar_url" class="avatar" alt="" />
      <span v-else class="avatar-placeholder">{{ user?.username?.charAt(0) || '?' }}</span>
    </button>
    <Transition name="dropdown">
      <div v-if="open" class="dropdown-menu">
        <div class="dropdown-user">
          <img v-if="user?.avatar_url" :src="user.avatar_url" class="avatar-lg" alt="" />
          <span v-else class="avatar-placeholder-lg">{{ user?.username?.charAt(0) || '?' }}</span>
          <div class="user-info">
            <div class="username">{{ user?.username }}</div>
            <div class="email">{{ user?.email || '未设置邮箱' }}</div>
          </div>
        </div>
        <div class="dropdown-divider"></div>
        <button class="dropdown-item" @click="showEdit = true; open = false">编辑资料</button>
        <button class="dropdown-item" @click="showPwd = true; open = false">修改密码</button>
        <div class="dropdown-divider"></div>
        <button class="dropdown-item logout" @click="handleLogout">退出登录</button>
      </div>
    </Transition>
    <Modal v-model="showEdit" title="编辑资料">
      <form @submit.prevent="saveProfile" class="form">
        <div class="form-item">
          <label>真实姓名</label>
          <input v-model="editForm.real_name" placeholder="真实姓名" />
        </div>
        <div class="form-item">
          <label>邮箱</label>
          <input v-model="editForm.email" type="email" placeholder="邮箱" />
        </div>
        <div class="form-item">
          <label>手机</label>
          <input v-model="editForm.phone" placeholder="手机号" />
        </div>
        <div class="form-item">
          <label>个人简介</label>
          <textarea v-model="editForm.intro" placeholder="个人简介" rows="3"></textarea>
        </div>
        <div class="form-item">
          <label>头像URL</label>
          <input v-model="editForm.avatar_url" placeholder="头像图片链接" />
        </div>
        <p v-if="editErr" class="err">{{ editErr }}</p>
        <div class="form-actions">
          <button type="button" class="btn-secondary" @click="showEdit = false">取消</button>
          <button type="submit" class="btn-primary" :disabled="editLoading">保存</button>
        </div>
      </form>
    </Modal>
    <Modal v-model="showPwd" title="修改密码">
      <form @submit.prevent="savePassword" class="form">
        <div class="form-item">
          <label>原密码</label>
          <input v-model="pwdForm.old_password" type="password" placeholder="请输入原密码" required />
        </div>
        <div class="form-item">
          <label>新密码</label>
          <input v-model="pwdForm.new_password" type="password" placeholder="请输入新密码（至少6位）" required minlength="6" />
        </div>
        <p v-if="pwdErr" class="err">{{ pwdErr }}</p>
        <div class="form-actions">
          <button type="button" class="btn-secondary" @click="showPwd = false">取消</button>
          <button type="submit" class="btn-primary" :disabled="pwdLoading">确认修改</button>
        </div>
      </form>
    </Modal>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, onBeforeUnmount, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '../stores/user'
import { getProfile, updateProfile, changePassword } from '../api/user'
import Modal from './Modal.vue'

const userStore = useUserStore()
const router = useRouter()
const dropdownRef = ref(null)
const open = ref(false)
const showEdit = ref(false)
const showPwd = ref(false)
const user = ref(null)
const editForm = reactive({ real_name: '', email: '', phone: '', intro: '', avatar_url: '' })
const pwdForm = reactive({ old_password: '', new_password: '' })
const editErr = ref('')
const pwdErr = ref('')
const editLoading = ref(false)
const pwdLoading = ref(false)

const fetchProfile = async () => {
  try {
    const res = await getProfile()
    user.value = res.data
    userStore.user = res.data
    localStorage.setItem('user', JSON.stringify(res.data))
    Object.assign(editForm, {
      real_name: user.value.real_name || '',
      email: user.value.email || '',
      phone: user.value.phone || '',
      intro: user.value.intro || '',
      avatar_url: user.value.avatar_url || '',
    })
  } catch (_) {}
}

onMounted(() => {
  if (userStore.isLoggedIn) fetchProfile()
  document.addEventListener('click', handleClickOutside)
})
onBeforeUnmount(() => document.removeEventListener('click', handleClickOutside))
watch(() => userStore.user, (u) => { user.value = u || user.value }, { immediate: true })
watch(showPwd, (v) => { if (!v) { pwdForm.old_password = ''; pwdForm.new_password = ''; pwdErr.value = '' } })
watch(showEdit, (v) => { if (!v) editErr.value = '' })

const handleClickOutside = (e) => {
  if (dropdownRef.value && !dropdownRef.value.contains(e.target)) open.value = false
}

const saveProfile = async () => {
  editErr.value = ''
  editLoading.value = true
  try {
    await updateProfile(editForm)
    user.value = { ...user.value, ...editForm }
    userStore.user = user.value
    localStorage.setItem('user', JSON.stringify(user.value))
    showEdit.value = false
  } catch (e) {
    editErr.value = e?.message || '保存失败'
  } finally {
    editLoading.value = false
  }
}

const savePassword = async () => {
  pwdErr.value = ''
  pwdLoading.value = true
  try {
    await changePassword(pwdForm)
    pwdForm.old_password = ''
    pwdForm.new_password = ''
    showPwd.value = false
  } catch (e) {
    pwdErr.value = e?.message || '修改失败'
  } finally {
    pwdLoading.value = false
  }
}

const handleLogout = () => {
  userStore.logout()
  open.value = false
  router.push('/')
}
</script>

<style scoped>
.user-dropdown { position: relative; }
.avatar-btn {
  width: 36px;
  height: 36px;
  border: none;
  background: none;
  padding: 0;
  cursor: pointer;
  border-radius: 50%;
  overflow: hidden;
}
.avatar { width: 100%; height: 100%; object-fit: cover; }
.avatar-placeholder {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #0066ff;
  color: #fff;
  font-size: 16px;
  font-weight: 600;
}
.dropdown-menu {
  position: absolute;
  top: calc(100% + 8px);
  right: 0;
  min-width: 220px;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 4px 24px rgba(0,0,0,.12);
  padding: 12px 0;
  z-index: 100;
}
.dropdown-user {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 0 16px 12px;
}
.avatar-lg { width: 48px; height: 48px; border-radius: 50%; object-fit: cover; }
.avatar-placeholder-lg {
  width: 48px; height: 48px; border-radius: 50%; background: #0066ff; color: #fff;
  display: flex; align-items: center; justify-content: center; font-size: 20px; font-weight: 600;
}
.user-info { flex: 1; min-width: 0; }
.username { font-weight: 600; color: #1a1a1a; }
.email { font-size: 13px; color: #999; margin-top: 2px; }
.dropdown-divider { height: 1px; background: #f0f0f0; margin: 8px 0; }
.dropdown-item {
  display: block;
  width: 100%;
  padding: 10px 16px;
  border: none;
  background: none;
  text-align: left;
  font-size: 15px;
  color: #1a1a1a;
  cursor: pointer;
}
.dropdown-item:hover { background: #f6f6f6; }
.dropdown-item.logout { color: #999; }
.dropdown-enter-active, .dropdown-leave-active { transition: opacity .15s, transform .15s; }
.dropdown-enter-from, .dropdown-leave-to { opacity: 0; transform: translateY(-8px); }
.form-item { margin-bottom: 16px; }
.form-item label { display: block; margin-bottom: 6px; font-size: 14px; color: #666; }
.form-item input, .form-item textarea {
  width: 100%; padding: 10px 12px; border: 1px solid #ddd; border-radius: 6px;
}
.form-item input:focus, .form-item textarea:focus {
  outline: none; border-color: #0066ff; box-shadow: 0 0 0 2px rgba(0,102,255,.1);
}
.form .err { color: #e74c3c; font-size: 14px; margin-bottom: 12px; }
.form-actions { display: flex; gap: 12px; justify-content: flex-end; margin-top: 20px; }
.btn-primary { padding: 8px 20px; background: #0066ff; color: #fff; border: none; border-radius: 6px; cursor: pointer; }
.btn-primary:hover:not(:disabled) { background: #0052cc; }
.btn-primary:disabled { opacity: 0.6; cursor: not-allowed; }
.btn-secondary { padding: 8px 20px; background: #f6f6f6; color: #1a1a1a; border: none; border-radius: 6px; cursor: pointer; }
.btn-secondary:hover { background: #ebebeb; }
</style>
