<template>
  <div class="home-container">
    <el-row :gutter="20">
      <el-col :span="8">
        <el-card class="stat-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <span>物品总数</span>
              <el-icon><Box /></el-icon>
            </div>
          </template>
          <div class="stat-value">{{ stats.totalItems }}</div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card class="stat-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <span>待签收订单</span>
              <el-icon><Clock /></el-icon>
            </div>
          </template>
          <div class="stat-value">{{ stats.pendingOrders }}</div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card class="stat-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <span>已签收订单</span>
              <el-icon><Check /></el-icon>
            </div>
          </template>
          <div class="stat-value">{{ stats.signedOrders }}</div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" class="mt-20">
      <el-col :span="12">
        <el-card shadow="hover" class="data-card">
          <template #header>
            <div class="card-header">
              <span>最近物品</span>
              <el-button type="primary" link @click="$router.push('/items')">
                查看全部
              </el-button>
            </div>
          </template>
          <el-table :data="recentItems" style="width: 100%" :show-header="false">
            <el-table-column prop="item_number" label="物品编号">
              <template #default="scope">
                <div class="table-cell">
                  <el-icon><Box /></el-icon>
                  <span>{{ scope.row.item_number }}</span>
                </div>
              </template>
            </el-table-column>
            <el-table-column prop="department" label="所属单位">
              <template #default="scope">
                <div class="table-cell">
                  <el-icon><OfficeBuilding /></el-icon>
                  <span>{{ scope.row.department }}</span>
                </div>
              </template>
            </el-table-column>
            <el-table-column prop="status" label="状态" width="100">
              <template #default="scope">
                <el-tag :type="getStatusType(scope.row.status)" size="small">
                  {{ getStatusText(scope.row.status) }}
                </el-tag>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card shadow="hover" class="data-card">
          <template #header>
            <div class="card-header">
              <span>最近订单</span>
              <el-button type="primary" link @click="$router.push('/orders')">
                查看全部
              </el-button>
            </div>
          </template>
          <el-table :data="recentOrders" style="width: 100%" :show-header="false">
            <el-table-column prop="order_number" label="订单编号">
              <template #default="scope">
                <div class="table-cell">
                  <el-icon><Document /></el-icon>
                  <span>{{ scope.row.order_number }}</span>
                </div>
              </template>
            </el-table-column>
            <el-table-column prop="item.item_number" label="物品编号">
              <template #default="scope">
                <div class="table-cell">
                  <el-icon><Box /></el-icon>
                  <span>{{ scope.row.item.item_number }}</span>
                </div>
              </template>
            </el-table-column>
            <el-table-column prop="status" label="状态" width="100">
              <template #default="scope">
                <el-tag :type="getOrderStatusType(scope.row.status)" size="small">
                  {{ getOrderStatusText(scope.row.status) }}
                </el-tag>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Box, Clock, Check, OfficeBuilding, Document } from '@element-plus/icons-vue'
import axios from 'axios'

const stats = ref({
  totalItems: 0,
  pendingOrders: 0,
  signedOrders: 0
})

const recentItems = ref([])
const recentOrders = ref([])

const API_BASE_URL = 'http://localhost:8080/api'

// 获取统计数据
const fetchStats = async () => {
  try {
    const [itemsRes, ordersRes] = await Promise.all([
      axios.get(`${API_BASE_URL}/items`),
      axios.get(`${API_BASE_URL}/orders`)
    ])

    stats.value.totalItems = itemsRes.data.length
    stats.value.pendingOrders = ordersRes.data.filter(order => order.status === 'pending').length
    stats.value.signedOrders = ordersRes.data.filter(order => order.status === 'signed').length
  } catch (error) {
    ElMessage.error('获取统计数据失败')
  }
}

// 获取最近物品
const fetchRecentItems = async () => {
  try {
    const response = await axios.get(`${API_BASE_URL}/items`)
    recentItems.value = response.data.slice(0, 5)
  } catch (error) {
    ElMessage.error('获取最近物品失败')
  }
}

// 获取最近订单
const fetchRecentOrders = async () => {
  try {
    const response = await axios.get(`${API_BASE_URL}/orders`)
    recentOrders.value = response.data.slice(0, 5)
  } catch (error) {
    ElMessage.error('获取最近订单失败')
  }
}

// 获取物品状态类型
const getStatusType = (status) => {
  const types = {
    in_stock: 'info',
    confirmed: 'warning',
    signed: 'success'
  }
  return types[status] || 'info'
}

// 获取物品状态文本
const getStatusText = (status) => {
  const texts = {
    in_stock: '入库',
    confirmed: '确权',
    signed: '签收'
  }
  return texts[status] || status
}

// 获取订单状态类型
const getOrderStatusType = (status) => {
  const types = {
    pending: 'warning',
    signed: 'success'
  }
  return types[status] || 'info'
}

// 获取订单状态文本
const getOrderStatusText = (status) => {
  const texts = {
    pending: '待签收',
    signed: '已签收'
  }
  return texts[status] || status
}

onMounted(() => {
  fetchStats()
  fetchRecentItems()
  fetchRecentOrders()
})
</script>

<style scoped>
.home-container {
  padding: 20px;
}

.stat-card {
  text-align: center;
  transition: all 0.3s;
}

.stat-card:hover {
  transform: translateY(-5px);
}

.stat-value {
  font-size: 36px;
  font-weight: bold;
  color: #409eff;
  margin: 20px 0;
}

.mt-20 {
  margin-top: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.data-card {
  height: 100%;
}

.table-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.table-cell .el-icon {
  font-size: 16px;
  color: #909399;
}

:deep(.el-card__header) {
  padding: 15px 20px;
  border-bottom: 1px solid #ebeef5;
}

:deep(.el-table) {
  --el-table-border-color: #ebeef5;
}

:deep(.el-table__row) {
  height: 50px;
}

:deep(.el-table__cell) {
  padding: 8px 0;
}
</style>
