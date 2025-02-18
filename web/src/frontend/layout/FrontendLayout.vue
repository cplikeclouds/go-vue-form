<template>
  <div class="frontend-layout">
    <!-- 用户端顶部导航 -->
    <el-menu mode="horizontal">
      <el-menu-item index="/home">
        <router-link to="/home">主页</router-link>
      </el-menu-item>
      <el-menu-item index="/form">
        <router-link to="/form">活动报名</router-link>
      </el-menu-item>
    </el-menu>

    <!-- 用户端内容区 -->
    <main class="content">
      <router-view></router-view>
    </main>
  </div>
</template>

<style scoped>
.frontend-layout {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.content {
  flex: 1;
  padding: 20px;
}
</style>