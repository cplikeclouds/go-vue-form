<template>
  <div class="dashboard">
    <h2>数据总览</h2>
    <div class="stats-grid">
      <div class="stat-card">
        <i class="fas fa-users"></i>
        <h3>总报名人数</h3>
        <p>1,234</p>
      </div>
      <div class="stat-card">
        <i class="fas fa-building"></i>
        <h3>企业代表</h3>
        <p>567</p>
      </div>
      <div class="stat-card">
        <i class="fas fa-graduation-cap"></i>
        <h3>学生代表</h3>
        <p>678</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.dashboard {
  padding: 20px;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
  margin-top: 20px;
}

.stat-card {
  background: white;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 3px 10px rgba(0, 0, 0, 0.1);
  text-align: center;
}

.stat-card i {
  font-size: 2rem;
  color: var(--primary-color);
  margin-bottom: 10px;
}

.stat-card h3 {
  margin: 10px 0;
  font-size: 1.2rem;
  color: var(--secondary-color);
}

.stat-card p {
  font-size: 1.5rem;
  font-weight: bold;
  color: var(--accent-color);
}
</style>