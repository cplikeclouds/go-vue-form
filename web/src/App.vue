<template>
  <!-- 仅保留根路由视图，布局由子路由决定 -->
  <router-view></router-view>
</template>

<script setup>
// 不需要路由判断逻辑，布局由路由配置控制
</script>

<style>
/* 全局基础样式 */
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  height: 100vh;
  display: flex;
  flex-direction: column;
}
</style>