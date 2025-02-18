<template>
  <div class="login-page">
    <h1>管理员登录</h1>
    <form @submit.prevent="login">
      <input v-model="username" placeholder="用户名" required />
      <input v-model="password" type="password" placeholder="密码" required />
      <button type="submit">登录</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const username = ref('');
const password = ref('');

const login = async () => {
  // 调用登录接口
  const response = await api.post('/admin/login', {
    username: username.value,
    password: password.value,
  });

  // 保存用户信息
  localStorage.setItem('user', JSON.stringify(response.data));

  // 跳转到管理端首页
  router.push('/admin/dashboard');
};
</script>