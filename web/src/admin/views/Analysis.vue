<template>
  <div class="analysis">
    <h2>深度分析</h2>
    <div class="charts">
      <div class="chart-card">
        <h3>报名人数趋势</h3>
        <div class="chart-placeholder"></div>
      </div>
      <div class="chart-card">
        <h3>行业分布</h3>
        <div class="chart-placeholder"></div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.analysis {
  padding: 20px;
}

.charts {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
  gap: 20px;
  margin-top: 20px;
}

.chart-card {
  background: white;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 3px 10px rgba(0, 0, 0, 0.1);
}

.chart-card h3 {
  margin-bottom: 20px;
  color: var(--secondary-color);
}

.chart-placeholder {
  height: 300px;
  background: #f5f5f5;
  border-radius: 5px;
}
</style>