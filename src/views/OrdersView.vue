<template>
  <div class="orders-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>订单管理</span>
          <el-button type="primary" @click="showAddDialog">创建订单</el-button>
        </div>
      </template>

      <el-table :data="orders" style="width: 100%">
        <el-table-column prop="order_number" label="订单编号" />
        <el-table-column prop="item.item_number" label="物品编号" />
        <el-table-column prop="item.department" label="所属单位" />
        <el-table-column prop="status" label="状态">
          <template #default="scope">
            <el-tag :type="getStatusType(scope.row.status)">
              {{ getStatusText(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="receiver" label="签收人" />
        <el-table-column prop="sign_time" label="签收时间">
          <template #default="scope">
            {{ formatDate(scope.row.sign_time) }}
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template #default="scope">
            <el-button-group>
              <el-button type="primary" @click="handleEdit(scope.row)">编辑</el-button>
              <el-button type="success" @click="handleSign(scope.row)">签收</el-button>
              <el-button type="danger" @click="handleDelete(scope.row)">删除</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 添加/编辑订单对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogType === 'add' ? '创建订单' : '编辑订单'"
      width="500px"
    >
      <el-form :model="form" label-width="100px">
        <el-form-item label="物品">
          <el-select v-model="form.item_id" placeholder="请选择物品">
            <el-option
              v-for="item in availableItems"
              :key="item.id"
              :label="item.item_number"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="form.status">
            <el-option label="待签收" value="pending" />
            <el-option label="已签收" value="signed" />
          </el-select>
        </el-form-item>
        <el-form-item label="签收人">
          <el-input v-model="form.receiver" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import axios from 'axios'

const orders = ref([])
const availableItems = ref([])
const dialogVisible = ref(false)
const dialogType = ref('add')
const form = ref({
  item_id: '',
  status: 'pending',
  receiver: ''
})

const API_BASE_URL = 'http://localhost:8080/api'

// 获取订单列表
const fetchOrders = async () => {
  try {
    const response = await axios.get(`${API_BASE_URL}/orders`)
    orders.value = response.data
  } catch (error) {
    ElMessage.error('获取订单列表失败')
  }
}

// 获取可用物品列表
const fetchAvailableItems = async () => {
  try {
    const response = await axios.get(`${API_BASE_URL}/items`)
    availableItems.value = response.data
  } catch (error) {
    ElMessage.error('获取物品列表失败')
  }
}

// 显示添加对话框
const showAddDialog = () => {
  dialogType.value = 'add'
  form.value = {
    item_id: '',
    status: 'pending',
    receiver: ''
  }
  dialogVisible.value = true
}

// 处理编辑
const handleEdit = (row) => {
  dialogType.value = 'edit'
  form.value = { ...row }
  dialogVisible.value = true
}

// 处理签收
const handleSign = async (row) => {
  try {
    await axios.put(`${API_BASE_URL}/orders/${row.id}`, {
      ...row,
      status: 'signed',
      sign_time: new Date().toISOString()
    })
    ElMessage.success('签收成功')
    fetchOrders()
  } catch (error) {
    ElMessage.error('签收失败')
  }
}

// 处理删除
const handleDelete = (row) => {
  ElMessageBox.confirm('确定要删除该订单吗？', '提示', {
    type: 'warning'
  }).then(async () => {
    try {
      await axios.delete(`${API_BASE_URL}/orders/${row.id}`)
      ElMessage.success('删除成功')
      fetchOrders()
    } catch (error) {
      ElMessage.error('删除失败')
    }
  })
}

// 处理表单提交
const handleSubmit = async () => {
  try {
    if (dialogType.value === 'add') {
      await axios.post(`${API_BASE_URL}/orders`, form.value)
      ElMessage.success('创建成功')
    } else {
      await axios.put(`${API_BASE_URL}/orders/${form.value.id}`, form.value)
      ElMessage.success('更新成功')
    }
    dialogVisible.value = false
    fetchOrders()
  } catch (error) {
    ElMessage.error(dialogType.value === 'add' ? '创建失败' : '更新失败')
  }
}

// 获取状态类型
const getStatusType = (status) => {
  const types = {
    pending: 'warning',
    signed: 'success'
  }
  return types[status] || 'info'
}

// 获取状态文本
const getStatusText = (status) => {
  const texts = {
    pending: '待签收',
    signed: '已签收'
  }
  return texts[status] || status
}

// 格式化日期
const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleString()
}

onMounted(() => {
  fetchOrders()
  fetchAvailableItems()
})
</script>

<style scoped>
.orders-container {
  max-width: 1200px;
  margin: 0 auto;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}
</style> 