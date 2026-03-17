<template>
  <div class="layout">
    <header class="header">
      <router-link to="/" class="logo">个人博客</router-link>
      <nav class="nav">
        <router-link to="/blogs">首页</router-link>
        <template v-if="userStore.isLoggedIn">
          <router-link to="/blogs/edit">写文章</router-link>
          <router-link v-if="userStore.isAdmin" to="/admin/users">管理</router-link>
          <UserDropdown />
        </template>
        <template v-else>
          <router-link to="/login">登录</router-link>
          <router-link to="/register">注册</router-link>
        </template>
      </nav>
    </header>
    <main class="main">
      <slot></slot>
    </main>
  </div>
</template>

<script setup>
import { useUserStore } from '../stores/user'
import UserDropdown from './UserDropdown.vue'

const userStore = useUserStore()
</script>

<style scoped>
.layout { min-height: 100vh; display: flex; flex-direction: column; }
.header {
  background: #fff;
  padding: 0 24px;
  height: 52px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-bottom: 1px solid #f0f0f0;
}
.logo { font-size: 20px; font-weight: 600; color: #1a1a1a; }
.nav { display: flex; align-items: center; gap: 28px; }
.nav a {
  color: #646464;
  font-size: 15px;
}
.nav a:hover { color: #0066ff; }
.nav a.router-link-active { color: #0066ff; font-weight: 500; }
.main { flex: 1; padding: 24px; max-width: 1000px; margin: 0 auto; width: 100%; }
</style>
