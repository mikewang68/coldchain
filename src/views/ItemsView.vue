<template>
  <div class="items-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>物品管理</span>
          <el-button type="primary" @click="showAddDialog">添加物品</el-button>
        </div>
      </template>

      <el-table :data="items" style="width: 100%">
        <el-table-column prop="item_number" label="物品编号" />
        <el-table-column prop="department" label="所属单位" />
        <el-table-column prop="status" label="状态">
          <template #default="scope">
            <el-tag :type="getStatusType(scope.row.status)">
              {{ getStatusText(scope.row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template #default="scope">
            <el-button-group>
              <el-button type="primary" @click="handleEdit(scope.row)">编辑</el-button>
              <el-button type="success" @click="handleConfirm(scope.row)">确权</el-button>
              <el-button type="danger" @click="handleDelete(scope.row)">删除</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 添加/编辑物品对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogType === 'add' ? '添加物品' : '编辑物品'"
      width="500px"
    >
      <el-form :model="form" label-width="100px">
        <el-form-item label="物品编号">
          <el-input v-model="form.item_number" />
        </el-form-item>
        <el-form-item label="所属单位">
          <el-input v-model="form.department" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="form.status">
            <el-option label="入库" value="in_stock" />
            <el-option label="确权" value="confirmed" />
            <el-option label="签收" value="signed" />
          </el-select>
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

const items = ref([])
const dialogVisible = ref(false)
const dialogType = ref('add')
const form = ref({
  item_number: '',
  department: '',
  status: 'in_stock'
})

const API_BASE_URL = 'http://localhost:8080/api'

// 获取物品列表
const fetchItems = async () => {
  try {
    const response = await axios.get(`${API_BASE_URL}/items`)
    items.value = response.data
  } catch (error) {
    ElMessage.error('获取物品列表失败')
  }
}

// 显示添加对话框
const showAddDialog = () => {
  dialogType.value = 'add'
  form.value = {
    item_number: '',
    department: '',
    status: 'in_stock'
  }
  dialogVisible.value = true
}

// 处理编辑
const handleEdit = (row) => {
  dialogType.value = 'edit'
  form.value = { ...row }
  dialogVisible.value = true
}

// 处理确权
const handleConfirm = async (row) => {
  try {
    await axios.put(`${API_BASE_URL}/items/${row.id}`, {
      ...row,
      status: 'confirmed'
    })
    ElMessage.success('确权成功')
    fetchItems()
  } catch (error) {
    ElMessage.error('确权失败')
  }
}

// 处理删除
const handleDelete = (row) => {
  ElMessageBox.confirm('确定要删除该物品吗？', '提示', {
    type: 'warning'
  }).then(async () => {
    try {
      await axios.delete(`${API_BASE_URL}/items/${row.id}`)
      ElMessage.success('删除成功')
      fetchItems()
    } catch (error) {
      ElMessage.error('删除失败')
    }
  })
}

// 处理表单提交
const handleSubmit = async () => {
  try {
    if (dialogType.value === 'add') {
      await axios.post(`${API_BASE_URL}/items`, form.value)
      ElMessage.success('添加成功')
    } else {
      await axios.put(`${API_BASE_URL}/items/${form.value.id}`, form.value)
      ElMessage.success('更新成功')
    }
    dialogVisible.value = false
    fetchItems()
  } catch (error) {
    ElMessage.error(dialogType.value === 'add' ? '添加失败' : '更新失败')
  }
}

// 获取状态类型
const getStatusType = (status) => {
  const types = {
    in_stock: 'info',
    confirmed: 'warning',
    signed: 'success'
  }
  return types[status] || 'info'
}

// 获取状态文本
const getStatusText = (status) => {
  const texts = {
    in_stock: '入库',
    confirmed: '确权',
    signed: '签收'
  }
  return texts[status] || status
}

onMounted(() => {
  fetchItems()
})
</script>

<style scoped>
.items-container {
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