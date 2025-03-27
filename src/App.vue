<script setup>
import { ElContainer, ElHeader, ElMain, ElMenu, ElMenuItem } from 'element-plus'
import { House, Box, List } from '@element-plus/icons-vue'
</script>

<template>
  <el-container class="layout-container">
    <el-header class="main-header">
      <div class="header-content">
        <div class="logo">
          <h1>物品管理系统</h1>
        </div>
        <el-menu mode="horizontal" router class="main-menu">
          <el-menu-item index="/">
            <el-icon><House /></el-icon>
            首页
          </el-menu-item>
          <el-menu-item index="/items">
            <el-icon><Box /></el-icon>
            物品管理
          </el-menu-item>
          <el-menu-item index="/orders">
            <el-icon><List /></el-icon>
            订单管理
          </el-menu-item>
        </el-menu>
      </div>
    </el-header>
    <el-main class="main-content">
      <router-view></router-view>
    </el-main>
  </el-container>
</template>

<style scoped>
.layout-container {
  min-height: 100vh;
  background-color: #f5f7fa;
}

.main-header {
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  padding: 0;
  position: fixed;
  width: 100%;
  z-index: 1000;
}

.header-content {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 100%;
  padding: 0 20px;
}

.logo h1 {
  margin: 0;
  font-size: 24px;
  color: #409eff;
  font-weight: 600;
}

.main-menu {
  border-bottom: none;
  background-color: transparent;
}

.main-menu :deep(.el-menu-item) {
  font-size: 16px;
  height: 60px;
  line-height: 60px;
  padding: 0 20px;
}

.main-menu :deep(.el-menu-item.is-active) {
  color: #409eff;
  border-bottom: 2px solid #409eff;
}

.main-menu :deep(.el-menu-item:hover) {
  background-color: #ecf5ff;
}

.main-content {
  padding-top: 80px;
  max-width: 1200px;
  margin: 0 auto;
  min-height: calc(100vh - 60px);
}

:deep(.el-menu-item .el-icon) {
  margin-right: 5px;
  font-size: 18px;
}
</style>
