<template>
  <div class="goods-page">
    <nav-bar></nav-bar>
    <div class="goods-container">
      <el-card>
        <div class="filter-container">
          <el-form :inline="true" :model="filterForm">
            <el-form-item label="批次编号">
              <el-input v-model="filterForm.goodsId" placeholder="请输入批次编号"></el-input>
            </el-form-item>
            <el-form-item label="海鲜类型">
              <el-select v-model="filterForm.type" placeholder="请选择海鲜类型">
                <el-option label="冷冻虾类" value="shrimp"></el-option>
                <el-option label="冷冻鱼类" value="fish"></el-option>
                <el-option label="冷冻贝类" value="shellfish"></el-option>
                <el-option label="冷冻蟹类" value="crab"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="状态">
              <el-select v-model="filterForm.status" placeholder="请选择状态">
                <el-option label="在库" value="in_stock"></el-option>
                <el-option label="已出库" value="out_stock"></el-option>
                <el-option label="抵押中" value="mortgaged"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="handleSearch">查询</el-button>
              <el-button @click="resetForm">重置</el-button>
            </el-form-item>
          </el-form>
        </div>

        <div class="operation-container">
          <el-button type="primary" @click="handleAdd">新增入库</el-button>
          <el-button type="warning" @click="handleOut">货物出库</el-button>
        </div>

        <el-table :data="goodsList" border style="width: 100%">
          <el-table-column prop="goodsId" label="批次编号" width="120"></el-table-column>
          <el-table-column prop="deviceId" label="设备标识" width="150">
            <template #default="scope">
              <el-tag type="info">{{ scope.row.deviceId || '未绑定' }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="name" label="海鲜名称"></el-table-column>
          <el-table-column prop="type" label="类型" width="100">
            <template #default="scope">
              <el-tag>{{ getTypeText(scope.row.type) }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="specification" label="等级" width="150"></el-table-column>
          <el-table-column prop="size" label="包装" width="100"></el-table-column>
          <el-table-column prop="weight" label="重量(吨)" width="100"></el-table-column>
          <el-table-column prop="temperature" label="温度(°C)" width="100"></el-table-column>
          <el-table-column prop="inTime" label="入库时间" width="180"></el-table-column>
          <el-table-column prop="status" label="状态" width="100">
            <template #default="scope">
              <el-tag :type="getStatusType(scope.row.status)">
                {{ getStatusText(scope.row.status) }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="200">
            <template #default="scope">
              <el-button size="mini" @click="handleEdit(scope.row)">编辑</el-button>
              <el-button size="mini" type="danger" @click="handleDelete(scope.row)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>

        <div class="pagination-container">
          <el-pagination
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :current-page="currentPage"
            :page-sizes="[10, 20, 30, 50]"
            :page-size="pageSize"
            layout="total, sizes, prev, pager, next, jumper"
            :total="total">
          </el-pagination>
        </div>
      </el-card>

      <!-- 新增/编辑对话框 -->
      <el-dialog :title="dialogTitle" v-model="dialogVisible" width="500px">
        <el-form :model="goodsForm" :rules="rules" ref="goodsForm" label-width="100px">
          <el-form-item label="海鲜名称" prop="name">
            <el-select v-model="goodsForm.name" placeholder="请选择海鲜名称">
              <el-option label="阿根廷红虾" value="阿根廷红虾"></el-option>
              <el-option label="加拿大北极甜虾" value="加拿大北极甜虾"></el-option>
              <el-option label="智利三文鱼" value="智利三文鱼"></el-option>
              <el-option label="大连扇贝" value="大连扇贝"></el-option>
              <el-option label="青岛鲅鱼" value="青岛鲅鱼"></el-option>
              <el-option label="波士顿龙虾" value="波士顿龙虾"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="设备标识" prop="deviceId">
            <el-input v-model="goodsForm.deviceId" placeholder="请输入NFC或条形码标识">
              <template #append>
                <el-button @click="scanDeviceId">扫描</el-button>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item label="海鲜类型" prop="type">
            <el-select v-model="goodsForm.type" placeholder="请选择海鲜类型">
              <el-option label="冷冻虾类" value="shrimp"></el-option>
              <el-option label="冷冻鱼类" value="fish"></el-option>
              <el-option label="冷冻贝类" value="shellfish"></el-option>
              <el-option label="冷冻蟹类" value="crab"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="等级" prop="specification">
            <el-select v-model="goodsForm.specification" placeholder="请选择等级">
              <el-option label="L1级 2000-3000头/kg" value="L1级 2000-3000头/kg"></el-option>
              <el-option label="L2级 2000-3000头/kg" value="L2级 2000-3000头/kg"></el-option>
              <el-option label="L3级 2000-3000头/kg" value="L3级 2000-3000头/kg"></el-option>
              <el-option label="L4级 2000-3000头/kg" value="L4级 2000-3000头/kg"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="包装" prop="size">
            <el-select v-model="goodsForm.size" placeholder="请选择包装大小">
              <el-option label="大包装" value="大包装"></el-option>
              <el-option label="中包装" value="中包装"></el-option>
              <el-option label="小包装" value="小包装"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="重量" prop="weight">
            <el-input-number v-model="goodsForm.weight" :precision="2" :step="0.1" :min="0"></el-input-number>
            <span class="unit">吨</span>
          </el-form-item>
          <el-form-item label="存储温度" prop="temperature">
            <el-input-number v-model="goodsForm.temperature" :precision="1" :step="0.5" :max="0"></el-input-number>
            <span class="unit">°C</span>
          </el-form-item>
          <el-form-item label="存储位置" prop="location">
            <el-select v-model="goodsForm.location" placeholder="请选择存储位置">
              <el-option label="A区-1号冷库" value="A区-1号冷库"></el-option>
              <el-option label="A区-2号冷库" value="A区-2号冷库"></el-option>
              <el-option label="B区-1号冷库" value="B区-1号冷库"></el-option>
              <el-option label="B区-2号冷库" value="B区-2号冷库"></el-option>
            </el-select>
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="submitForm">确定</el-button>
          </span>
        </template>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import NavBar from '../components/Layout/NavBar.vue'
import axios from 'axios'

export default {
  name: 'GoodsView',
  components: {
    NavBar
  },
  data() {
    return {
      filterForm: {
        goodsId: '',
        type: '',
        status: ''
      },
      goodsList: [],
      dialogVisible: false,
      dialogTitle: '新增入库',
      goodsForm: {
        name: '',
        deviceId: '',
        type: '',
        specification: '',
        size: '',
        weight: 0,
        temperature: 0,
        location: ''
      },
      rules: {
        name: [{ required: true, message: '请选择海鲜名称', trigger: 'change' }],
        deviceId: [{ required: true, message: '请输入设备标识', trigger: 'blur' }],
        type: [{ required: true, message: '请选择海鲜类型', trigger: 'change' }],
        specification: [{ required: true, message: '请选择等级', trigger: 'change' }],
        size: [{ required: true, message: '请选择包装大小', trigger: 'change' }],
        weight: [{ required: true, message: '请输入重量', trigger: 'blur' }],
        temperature: [{ required: true, message: '请输入存储温度', trigger: 'blur' }],
        location: [{ required: true, message: '请选择存储位置', trigger: 'change' }]
      },
      currentPage: 1,
      pageSize: 10,
      total: 0
    }
  },
  created() {
    this.fetchGoodsList()
  },
  methods: {
    async fetchGoodsList() {
      try {
        const response = await axios.get('http://localhost:8090/api/goods/list', {
          params: this.filterForm
        })
        this.goodsList = response.data.items
        this.total = response.data.total
      } catch (error) {
        this.$message.error('获取货物列表失败')
      }
    },
    handleSearch() {
      this.currentPage = 1
      this.fetchGoodsList()
    },
    resetForm() {
      this.filterForm = {
        goodsId: '',
        type: '',
        status: ''
      }
      this.handleSearch()
    },
    handleAdd() {
      this.dialogTitle = '新增入库'
      this.goodsForm = {
        name: '',
        deviceId: '',
        type: '',
        specification: '',
        size: '',
        weight: 0,
        temperature: 0,
        location: ''
      }
      this.dialogVisible = true
    },
    async submitForm() {
      try {
        await this.$refs.goodsForm.validate()
        if (this.dialogTitle === '新增入库') {
          await axios.post('http://localhost:8090/api/goods/create', this.goodsForm)
          this.$message.success('入库成功')
        } else {
          await axios.put(`http://localhost:8090/api/goods/update/${this.goodsForm.goodsId}`, this.goodsForm)
          this.$message.success('更新成功')
        }
        this.dialogVisible = false
        this.fetchGoodsList()
      } catch (error) {
        this.$message.error(error.response?.data?.error || '操作失败')
      }
    },
    async handleEdit(row) {
      this.dialogTitle = '编辑货物'
      this.goodsForm = { ...row }
      this.dialogVisible = true
    },
    async handleDelete(row) {
      try {
        await this.$confirm('确认删除该货物？', '提示', {
          type: 'warning'
        })
        await axios.delete(`http://localhost:8090/api/goods/delete/${row.goodsId}`)
        this.$message.success('删除成功')
        this.fetchGoodsList()
      } catch (error) {
        if (error !== 'cancel') {
          this.$message.error('删除失败')
        }
      }
    },
    async handleOut(row) {
      try {
        await this.$confirm('确认出库该货物？', '提示', {
          type: 'warning'
        })
        await axios.post(`http://localhost:8090/api/goods/out/${row.goodsId}`)
        this.$message.success('出库成功')
        this.fetchGoodsList()
      } catch (error) {
        if (error !== 'cancel') {
          this.$message.error(error.response?.data?.error || '出库失败')
        }
      }
    },
    handleSizeChange(val) {
      this.pageSize = val
      this.fetchGoodsList()
    },
    handleCurrentChange(val) {
      this.currentPage = val
      this.fetchGoodsList()
    },
    getTypeText(type) {
      const typeMap = {
        shrimp: '冷冻虾类',
        fish: '冷冻鱼类',
        shellfish: '冷冻贝类',
        crab: '冷冻蟹类'
      }
      return typeMap[type] || type
    },
    getStatusType(status) {
      const statusMap = {
        in_stock: 'success',
        out_stock: 'info',
        mortgaged: 'warning'
      }
      return statusMap[status] || 'info'
    },
    getStatusText(status) {
      const statusMap = {
        in_stock: '在库',
        out_stock: '已出库',
        mortgaged: '抵押中'
      }
      return statusMap[status] || status
    },
    scanDeviceId() {
      // 模拟扫描设备ID
      this.goodsForm.deviceId = 'NFC:' + Math.random().toString(16).slice(2, 8).toUpperCase()
    }
  }
}
</script>

<style scoped>
.goods-page {
  width: 100%;
  height: 100vh;
  overflow: hidden;
}

.goods-container {
  position: fixed;
  top: 60px;
  left: 0;
  right: 0;
  bottom: 0;
  padding: 0;
  margin: 0;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  background-color: #f0f2f5;
  overflow: hidden;
}

.el-card {
  height: 100%;
  display: flex;
  flex-direction: column;
  border-radius: 0;
  margin: 0;
}

:deep(.el-card__body) {
  flex: 1;
  padding: 15px;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.filter-container {
  margin-bottom: 15px;
  flex-shrink: 0;
  padding: 0 5px;
}

.operation-container {
  margin-bottom: 15px;
  flex-shrink: 0;
  padding: 0 5px;
}

.el-table {
  flex: 1;
  overflow: auto;
}

.pagination-container {
  margin-top: 15px;
  text-align: right;
  flex-shrink: 0;
  padding: 0 5px;
}

.dialog-footer {
  text-align: right;
}

.unit {
  margin-left: 8px;
  color: #666;
}

.el-tag {
  margin-right: 8px;
}

/* 响应式布局调整 */
@media screen and (max-width: 1366px) {
  :deep(.el-card__body) {
    padding: 10px;
  }

  .el-table {
    font-size: 13px;
  }

  .el-button--mini {
    padding: 6px 10px;
    font-size: 12px;
  }
}

@media screen and (min-width: 1920px) {
  :deep(.el-card__body) {
    padding: 20px;
  }

  .el-table {
    font-size: 15px;
  }
}

/* 表格滚动条样式 */
.el-table__body-wrapper::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

.el-table__body-wrapper::-webkit-scrollbar-thumb {
  background: #ddd;
  border-radius: 3px;
}

.el-table__body-wrapper::-webkit-scrollbar-track {
  background: #f5f5f5;
}

/* 对话框样式优化 */
:deep(.el-dialog) {
  margin: 0 !important;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  max-height: 90vh;
  display: flex;
  flex-direction: column;
}

:deep(.el-dialog__body) {
  flex: 1;
  overflow-y: auto;
}
</style>
