<template>
  <div class="admin-layout">
    <!-- 侧边栏 -->
    <aside class="sidebar">
      <div class="sidebar-header">
        <h1>活动报名管理系统</h1>
      </div>
      <nav>
        <router-link to="/admin/dashboard" class="nav-item">
          <i class="fas fa-tachometer-alt"></i>
          <span>数据总览</span>
        </router-link>
        <router-link to="/admin/data-list" class="nav-item">
          <i class="fas fa-list"></i>
          <span>数据列表</span>
        </router-link>
        <router-link to="/admin/analysis" class="nav-item">
          <i class="fas fa-chart-bar"></i>
          <span>深度分析</span>
        </router-link>
        <div class="user-info">
          <i class="fas fa-user-circle"></i>
          <span>管理员</span>
        </div>
      </nav>
    </aside>

    <!-- 主内容区 -->
    <main class="main-content">
      <router-view></router-view>
    </main>
  </div>
</template>

<style scoped>
.admin-layout {
  display: flex;
  height: 100vh;
  background: #f5f5f5;
}

/* 侧边栏 */
.sidebar {
  width: 250px;
  background: linear-gradient(rgba(0,0,0,0.6), rgba(0,0,0,0.6)),
  url('../../../public/home_background.png') center/cover;
  color: white;
  box-shadow: 2px 0 5px rgba(0, 0, 0, 0.1);
}

.sidebar-header {
  padding: 20px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.sidebar-header h1 {
  margin: 0;
  font-size: 1.5rem;
}

/* 导航项 */
.nav-item {
  display: flex;
  align-items: center;
  padding: 15px 20px;
  color: white;
  text-decoration: none;
  transition: background 0.3s;
}

.nav-item:hover {
  background: rgba(255, 255, 255, 0.1);
}

.nav-item i {
  margin-right: 10px;
}

/* 用户信息 */
.user-info {
  display: flex;
  align-items: center;
  padding: 15px 20px;
  margin-top: auto;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.user-info i {
  margin-right: 10px;
  font-size: 1.2rem;
}

/* 主内容区 */
.main-content {
  flex: 1;
  padding: 20px;
  background: white;
  overflow-y: auto;
}
</style>