<template>
  <div class="home-container">
    <!-- 头部背景图 -->
    <header class="hero-banner">
      <div class="hero-content">
        <h1 class="event-title">{{ eventTitle }}</h1>
        <p class="event-date-location">
          <i class="fas fa-calendar-alt"></i> {{ eventDate }}
          <i class="fas fa-map-marker-alt"></i> {{ eventLocation }}
        </p>
        <button class="cta-button" @click="navigateToForm">立即报名</button>
      </div>
    </header>

    <!-- 活动亮点 -->
    <section class="highlights">
      <div class="highlight-card">
        <i class="fas fa-users"></i>
        <h3>行业大咖云集</h3>
        <p>来自腾讯、阿里、哈佛、斯坦福等顶尖企业和院校的嘉宾</p>
      </div>
      <div class="highlight-card">
        <i class="fas fa-handshake"></i>
        <h3>精准资源对接</h3>
        <p>根据您的身份（学生/在职）匹配交流对象</p>
      </div>
      <div class="highlight-card">
        <i class="fas fa-chart-line"></i>
        <h3>数据驱动洞察</h3>
        <p>会后提供行业人才趋势分析报告</p>
      </div>
    </section>

    <!-- 适合人群 -->
    <section class="target-audience">
      <h2 class="section-title">适合参与人群</h2>
      <div class="audience-grid">
        <div class="audience-item employee">
          <i class="fas fa-briefcase"></i>
          <h3>在职人员</h3>
          <ul>
            <li>寻找行业合作机会</li>
            <li>了解前沿技术趋势</li>
            <li>拓展职业人脉网络</li>
          </ul>
        </div>
        <div class="audience-item student">
          <i class="fas fa-graduation-cap"></i>
          <h3>在校学生</h3>
          <ul>
            <li>获取实习就业机会</li>
            <li>接触行业真实案例</li>
            <li>寻找导师与合作伙伴</li>
          </ul>
        </div>
      </div>
    </section>

    <!-- 底部CTA -->
    <div class="bottom-cta">
      <h2>立即加入，开启您的行业连接之旅！</h2>
      <button class="cta-button" @click="navigateToForm">马上报名</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';

const router = useRouter();
const eventTitle = "2025 行业创新人才交流会";
const eventDate = "2025年11月15-16日";
const eventLocation = "上海国际会议中心 | 线上同步";

const navigateToForm = () => {
  router.push('/form');
};
</script>

<style scoped>
/* 自定义CSS变量 */
:root {
  --primary-color: #42b983;  /* 主品牌色 */
  --secondary-color: #2c3e50; /* 次级颜色 */
  --accent-color: #ff6b6b;   /* 强调色 */
}

/* 基础布局 */
.home-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

/* 头部背景 */
.hero-banner {
  background: linear-gradient(rgba(0,0,0,0.6), rgba(0,0,0,0.6)),
  url('../../../public/home_background.png') center/cover;
  border-radius: 15px;
  padding: 80px 20px;
  text-align: center;
  color: white;
  margin-bottom: 40px;
}

.event-title {
  font-size: 2.5rem;
  margin-bottom: 20px;
  text-shadow: 2px 2px 4px rgba(0,0,0,0.5);
}

.event-date-location {
  font-size: 1.2rem;
  margin-bottom: 30px;
}

.event-date-location i {
  margin: 0 10px;
}

/* CTA按钮 */
.cta-button {
  background: var(--primary-color);
  color: white;
  padding: 15px 40px;
  border: none;
  border-radius: 30px;
  font-size: 1.1rem;
  cursor: pointer;
  transition: transform 0.3s, box-shadow 0.3s;
}

.cta-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 5px 15px rgba(66, 185, 131, 0.4);
}

/* 亮点卡片 */
.highlights {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 30px;
  margin: 50px 0;
}

.highlight-card {
  background: white;
  padding: 30px;
  border-radius: 10px;
  box-shadow: 0 5px 15px rgba(0,0,0,0.1);
  text-align: center;
  transition: transform 0.3s;
}

.highlight-card:hover {
  transform: translateY(-5px);
}

.highlight-card i {
  font-size: 2.5rem;
  color: var(--primary-color);
  margin-bottom: 20px;
}

/* 适合人群 */
.target-audience {
  text-align: center;
  background: #f9f9f9;
  padding: 50px 20px;
  border-radius: 15px;
  margin: 40px 0;
}

.section-title {
  text-align: center;
  color: var(--secondary-color);
  margin-bottom: 40px;
}

.audience-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 30px;
  max-width: 1000px;
  margin: 0 auto;
}

.audience-item {
  background: white;
  padding: 30px;
  border-radius: 10px;
  box-shadow: 0 3px 10px rgba(0,0,0,0.1);
  text-align: center;
}

.audience-item i {
  font-size: 2rem;
  color: var(--accent-color);
  margin-bottom: 15px;
}

.audience-item ul {
  list-style: none;
  padding: 0;
  text-align: left;
}

.audience-item li {
  padding: 8px 0;
  border-bottom: 1px solid #eee;
}

/* 底部CTA */
.bottom-cta {
  text-align: center;
  padding: 60px 20px;
  background: var(--secondary-color);
  color: white;
  border-radius: 15px;
  margin-top: 50px;
}

.bottom-cta h2 {
  margin-bottom: 30px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .event-title {
    font-size: 2rem;
  }

  .highlight-card {
    padding: 20px;
  }

  .audience-item {
    padding: 20px;
  }
}
</style>