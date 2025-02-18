<template>
  <div class="success-page">
    <h1>提交成功！</h1>
    <router-link to="/">返回首页</router-link>
  </div>
</template>

<script setup lang="ts">
// 提交成功页逻辑（可选）
</script>

<style scoped>
.success-page {
  text-align: center;
  margin-top: 50px;
}

h1 {
  color: #42b983;
}

a {
  display: inline-block;
  margin-top: 20px;
  color: #42b983;
  text-decoration: none;
}

a:hover {
  text-decoration: underline;
}
</style>