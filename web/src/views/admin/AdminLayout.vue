<template>
  <Layout>
    <div class="admin">
      <aside class="sidebar">
        <router-link to="/admin/users">用户管理</router-link>
        <router-link to="/admin/blogs">文章管理</router-link>
      </aside>
      <main class="content">
        <slot />
      </main>
    </div>
  </Layout>
</template>

<style scoped>
.admin { display: flex; gap: 24px; }
.sidebar { width: 180px; flex-shrink: 0; }
.sidebar a { display: block; padding: 12px; color: #666; border-radius: 6px; }
.sidebar a.router-link-active { background: #42b983; color: #fff; }
.content { flex: 1; }
</style>
