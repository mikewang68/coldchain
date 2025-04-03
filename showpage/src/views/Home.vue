<template>
  <div class="home-container">
    <!-- 顶部流程图 -->
    <div class="process-flow">
      <div class="flow-item" @click="handleNavClick('goods')" style="cursor: pointer;">
        <i class="el-icon-user"></i>
        <span>海鲜入库</span>
      </div>
      <i class="el-icon-arrow-right"></i>
      <div class="flow-item" @click="handleNavClick('mortgage')" style="cursor: pointer;">
        <i class="el-icon-box"></i>
        <span>银行/监管方</span>
      </div>
      <i class="el-icon-arrow-right"></i>
      <div class="flow-item" @click="handleNavClick('orders')" style="cursor: pointer;">
        <i class="el-icon-unlock"></i>
        <span>解押出库</span>
      </div>
      <i class="el-icon-arrow-right"></i>
      <div class="flow-item" @click="handleNavClick('items')" style="cursor: pointer;">
        <i class="el-icon-box"></i>
        <span>冷冻海鲜</span>
      </div>
      <i class="el-icon-arrow-right"></i>
      <div class="flow-item">
        <i class="el-icon-video-camera"></i>
        <span>温度监控</span>
      </div>
      <i class="el-icon-arrow-right"></i>
      <div class="flow-item">
        <i class="el-icon-warning"></i>
        <span>温度告警</span>
      </div>
      <i class="el-icon-arrow-right"></i>
      <div class="flow-item">
        <i class="el-icon-data-line"></i>
        <span>海鲜流通</span>
      </div>
    </div>

    <!-- 主要内容区域 -->
    <div class="main-content">
      <!-- 左侧数据统计 -->
      <div class="left-panel">
        <div class="data-section">
          <div class="data-title">冷库总览</div>
          <div class="data-row">
            <div class="data-item">
              <div class="data-value">128.5</div>
              <div class="data-label">总库存量（吨）</div>
            </div>
            <div class="data-item">
              <div class="data-value">1,280.5</div>
              <div class="data-label">总货值（万元）</div>
            </div>
          </div>
        </div>

        <div class="data-section">
          <div class="data-title">今日出入库</div>
          <div class="data-row">
            <div class="data-item">
              <div class="data-value">12.5</div>
              <div class="data-label">出库重量（吨）</div>
            </div>
            <div class="data-item">
              <div class="data-value">125.4</div>
              <div class="data-label">出库货值（万元）</div>
            </div>
          </div>
          <div class="data-row">
            <div class="data-item">
              <div class="data-value">15.8</div>
              <div class="data-label">入库重量（吨）</div>
            </div>
            <div class="data-item">
              <div class="data-value">158.6</div>
              <div class="data-label">入库货值（万元）</div>
            </div>
          </div>
        </div>

        <div class="data-section">
          <div class="data-title">监管情况</div>
          <div class="data-row">
            <div class="data-item">
              <div class="data-value">98.5</div>
              <div class="data-label">监管总重量（吨）</div>
            </div>
            <div class="data-item">
              <div class="data-value">985.4</div>
              <div class="data-label">监管总货值（万元）</div>
            </div>
          </div>
          <div class="data-row">
            <div class="data-item">
              <div class="data-value">8.5</div>
              <div class="data-label">今日新增重量（吨）</div>
            </div>
            <div class="data-item">
              <div class="data-value">85.4</div>
              <div class="data-label">今日新增货值（万元）</div>
            </div>
          </div>
        </div>

        <div class="data-section">
          <div class="data-title">温度告警</div>
          <div class="alarm-list">
            <div class="alarm-item" v-for="(item, index) in alarmList" :key="index">
              <span class="alarm-time">{{ item.time }}</span>
              <span class="alarm-content">{{ item.content }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- 中间3D展示区域 -->
      <div class="center-panel">
        <div class="warehouse-3d">
          <div class="camera-preview">
            <el-image :src="activeCameraUrl" fit="cover">
              <template #error>
                <div class="image-slot">暂无画面</div>
              </template>
            </el-image>
            <div class="camera-info">
              <div class="info-row">
                <span class="label">监控时间：</span>
                <span>2023/07/28 09:25:50</span>
              </div>
              <div class="info-row">
                <span class="label">冷库区域：</span>
                <span>A区-1号冷库</span>
              </div>
              <div class="info-row">
                <span class="label">当前温度：</span>
                <span>-18°C</span>
              </div>
            </div>
          </div>
          <div class="camera-list">
            <div class="camera-item" v-for="(camera, index) in cameraList" :key="index" @click="selectCamera(camera)">
              <el-image :src="camera.url" fit="cover">
                <template #error>
                  <div class="image-slot">暂无画面</div>
                </template>
              </el-image>
            </div>
          </div>
        </div>
      </div>

      <!-- 右侧物品详情 -->
      <div class="right-panel">
        <div class="item-detail">
          <div class="detail-header">
            <h3>货品详情</h3>
            <el-button type="text" @click="closeDetail">×</el-button>
          </div>
          <div class="detail-content">
            <div class="item-image">
              <el-image src="/path/to/seafood-image.jpg" fit="contain">
                <template #error>
                  <div class="image-slot">
                    <i class="el-icon-picture-outline"></i>
                  </div>
                </template>
              </el-image>
            </div>
            <div class="item-info">
              <div class="info-row">
                <span class="label">批次编号：</span>
                <span>SF20230728001</span>
              </div>
              <div class="info-row">
                <span class="label">货品名称：</span>
                <span>阿根廷红虾</span>
              </div>
              <div class="info-row">
                <span class="label">规格等级：</span>
                <span>L1级 2000-3000头/kg</span>
              </div>
              <div class="info-row">
                <span class="label">库存重量：</span>
                <span>25.5吨</span>
              </div>
              <div class="info-row">
                <span class="label">当前温度：</span>
                <span>-18°C</span>
              </div>
              <div class="info-row">
                <span class="label">入库时间：</span>
                <span>2023-07-20</span>
              </div>
              <div class="info-row">
                <span class="label">监管状态：</span>
                <span class="status">抵押中</span>
              </div>
              <div class="info-row">
                <span class="label">货值评估：</span>
                <span>255万元</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HomePage',
  data() {
    return {
      alarmList: [
        { time: '07/28 20:34:23', content: 'A区-1号冷库温度异常，当前温度-15°C，超过预警值' },
        { time: '07/28 19:45:12', content: 'B区-2号冷库温度波动，建议检查制冷设备' },
        { time: '07/28 18:22:45', content: 'C区-3号冷库门禁异常，检测到非授权人员进入' },
        { time: '07/28 17:15:30', content: 'A区-2号冷库温度恢复正常，当前温度-18°C' }
      ],
      activeCameraUrl: '/path/to/camera1.jpg',
      cameraList: [
        { id: 1, url: '/path/to/camera1.jpg', name: 'A区-1号冷库' },
        { id: 2, url: '/path/to/camera2.jpg', name: 'B区-2号冷库' },
        { id: 3, url: '/path/to/camera3.jpg', name: 'C区-3号冷库' }
      ]
    }
  },
  methods: {
    selectCamera(camera) {
      this.activeCameraUrl = camera.url
    },
    closeDetail() {
      // 实现关闭详情面板的逻辑
    },
    handleNavClick(route) {
      this.$router.push(`/${route}`)
    }
  }
}
</script>

<style scoped>
.home-container {
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: #f0f2f5;
  overflow: hidden;
  position: fixed;
  top: 0;
  left: 0;
}

.process-flow {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1vh 2vw;
  background-color: #1e2128;
  color: white;
  height: 10vh;
}

.flow-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.5vh;
}

.flow-item i {
  font-size: calc(1.2rem + 0.5vw);
}

.flow-item span {
  font-size: calc(0.7rem + 0.2vw);
  white-space: nowrap;
}

.main-content {
  flex: 1;
  display: flex;
  gap: 1vw;
  padding: 1vh 1vw;
  height: 90vh;
}

.left-panel {
  width: 25vw;
  display: flex;
  flex-direction: column;
  gap: 1vh;
  overflow-y: auto;
}

.data-section {
  background-color: white;
  border-radius: 8px;
  padding: 1.5vh 1vw;
  flex-shrink: 0;
}

.data-title {
  font-size: calc(0.9rem + 0.2vw);
  color: #333;
  margin-bottom: 1.5vh;
  border-left: 4px solid #409EFF;
  padding-left: 0.8vw;
}

.data-row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 1.5vh;
}

.data-value {
  font-size: calc(1.2rem + 0.3vw);
  color: #409EFF;
  font-weight: bold;
}

.data-label {
  font-size: calc(0.7rem + 0.1vw);
  color: #666;
  margin-top: 0.5vh;
}

.alarm-list {
  max-height: 20vh;
  overflow-y: auto;
}

.alarm-item {
  display: flex;
  gap: 0.8vw;
  padding: 1vh 0;
  border-bottom: 1px solid #eee;
}

.alarm-time {
  color: #999;
  font-size: calc(0.7rem + 0.1vw);
  white-space: nowrap;
}

.alarm-content {
  color: #666;
  font-size: calc(0.7rem + 0.1vw);
}

.center-panel {
  flex: 1;
  min-width: 45vw;
  background-color: white;
  border-radius: 8px;
  padding: 1.5vh 1vw;
  display: flex;
  flex-direction: column;
}

.warehouse-3d {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 1vh;
}

.camera-preview {
  flex: 1;
  position: relative;
  background-color: #000;
  border-radius: 4px;
  overflow: hidden;
  min-height: 50vh;
}

.camera-preview .el-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.camera-info {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 1.5vh 1vw;
  background: rgba(0, 0, 0, 0.7);
  color: white;
}

.camera-list {
  height: 15vh;
  display: flex;
  gap: 1vw;
}

.camera-item {
  flex: 1;
  height: 100%;
  border-radius: 4px;
  overflow: hidden;
  cursor: pointer;
}

.camera-item .el-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.right-panel {
  width: 20vw;
  background-color: white;
  border-radius: 8px;
  padding: 1.5vh 1vw;
  overflow-y: auto;
}

.item-detail .detail-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2vh;
}

.item-detail h3 {
  margin: 0;
  color: #333;
  font-size: calc(1rem + 0.2vw);
}

.item-image {
  width: 100%;
  height: 25vh;
  margin-bottom: 2vh;
}

.item-info .info-row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 1.5vh;
  font-size: calc(0.8rem + 0.1vw);
}

.info-row .label {
  color: #666;
}

.image-slot {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  background: #f5f7fa;
  color: #909399;
}

:deep(.el-icon-arrow-right) {
  color: white;
  font-size: calc(1.2rem + 0.3vw);
}

/* 滚动条样式优化 */
::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

::-webkit-scrollbar-thumb {
  background: #888;
  border-radius: 3px;
}

::-webkit-scrollbar-thumb:hover {
  background: #555;
}

.flow-item:hover {
  opacity: 0.8;
  transform: scale(1.05);
  transition: all 0.3s ease;
}

.flow-link {
  text-decoration: none;
  color: inherit;
  display: flex;
  flex-direction: column;
  align-items: center;
  cursor: pointer;
}

.flow-link:hover {
  color: #409EFF;
}
</style>
