<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script setup>
</script>

<style>
/* 知乎风格：简洁、专业、蓝色主调 */
* { margin: 0; padding: 0; box-sizing: border-box; }
body {
  font-family: -apple-system, BlinkMacSystemFont, 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', sans-serif;
  font-size: 15px;
  color: #1a1a1a;
  background: #f6f6f6;
  line-height: 1.6;
}
#app { min-height: 100vh; }
a { color: #0066ff; text-decoration: none; }
a:hover { color: #0052cc; text-decoration: underline; }
input, textarea, button { font-family: inherit; font-size: inherit; }
</style>
