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
              <el-option label="阿根廷红虾" value="1"></el-option>
              <el-option label="加拿大北极甜虾" value="2"></el-option>
              <el-option label="智利三文鱼" value="3"></el-option>
              <el-option label="大连扇贝" value="4"></el-option>
              <el-option label="青岛鲅鱼" value="5"></el-option>
              <el-option label="波士顿龙虾" value="6"></el-option>
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
              <el-option label="L1级 2000-3000头/kg" value="first"></el-option>
              <el-option label="L2级 2000-3000头/kg" value="second"></el-option>
              <el-option label="L3级 2000-3000头/kg" value="third"></el-option>
              <el-option label="L4级 2000-3000头/kg" value="forth"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="包装" prop="size">
            <el-select v-model="goodsForm.size" placeholder="请选择包装大小">
              <el-option label="大包装" value="big"></el-option>
              <el-option label="中包装" value="mid"></el-option>
              <el-option label="小包装" value="small"></el-option>
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
              <el-option label="A区-1号冷库" value="A1"></el-option>
              <el-option label="A区-2号冷库" value="A2"></el-option>
              <el-option label="B区-1号冷库" value="B1"></el-option>
              <el-option label="B区-2号冷库" value="B2"></el-option>
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
      goodsList: [
        {
          goodsId: 'SF20240320001',
          deviceId: 'NFC:04:E7:52:B9:D4:80',
          name: '阿根廷红虾',
          type: 'shrimp',
          specification: 'L1级 2000-3000头/kg',
          size: '大包装',
          weight: 25.5,
          temperature: -18,
          location: 'A1',
          inTime: '2024-03-20 10:00:00',
          status: 'in_stock'
        },
        {
          goodsId: 'SF20240319002',
          deviceId: 'NFC:04:E7:52:C1:A3:F5',
          name: '加拿大北极甜虾',
          type: 'shrimp',
          specification: 'L2级 2000-3000头/kg',
          size: '中包装',
          weight: 18.2,
          temperature: -20,
          location: 'A2',
          inTime: '2024-03-19 15:30:00',
          status: 'mortgaged'
        },
        {
          goodsId: 'SF20240318003',
          deviceId: '978020137962',
          name: '智利三文鱼',
          type: 'fish',
          specification: 'L3级 2000-3000头/kg',
          size: '小包装',
          weight: 15.8,
          temperature: -22,
          location: 'B1',
          inTime: '2024-03-18 09:15:00',
          status: 'in_stock'
        },
        {
          goodsId: 'SF20240317004',
          deviceId: 'NFC:04:E7:52:D8:E2:B1',
          name: '大连扇贝',
          type: 'shellfish',
          specification: 'L4级 2000-3000头/kg',
          size: '大包装',
          weight: 12.5,
          temperature: -18,
          location: 'B2',
          inTime: '2024-03-17 14:20:00',
          status: 'mortgaged'
        },
        {
          goodsId: 'SF20240316005',
          deviceId: '978159683254',
          name: '帝王蟹',
          type: 'crab',
          specification: 'L1级 2000-3000头/kg',
          size: '中包装',
          weight: 8.6,
          temperature: -20,
          location: 'A1',
          inTime: '2024-03-16 11:40:00',
          status: 'in_stock'
        },
        {
          goodsId: 'SF20240315006',
          deviceId: 'NFC:04:E7:52:F4:C7:A9',
          name: '青岛鲅鱼',
          type: 'fish',
          specification: 'L2级 2000-3000头/kg',
          size: '小包装',
          weight: 20.5,
          temperature: -19,
          location: 'A2',
          inTime: '2024-03-15 16:50:00',
          status: 'out_stock'
        },
        {
          goodsId: 'SF20240314007',
          deviceId: '978314529687',
          name: '波士顿龙虾',
          type: 'shellfish',
          specification: 'L3级 2000-3000头/kg',
          size: '大包装',
          weight: 10.2,
          temperature: -20,
          location: 'B1',
          inTime: '2024-03-14 13:25:00',
          status: 'mortgaged'
        }
      ],
      currentPage: 1,
      pageSize: 10,
      total: 100,
      dialogVisible: false,
      dialogTitle: '新增入库',
      goodsForm: {
        name: '',
        type: '',
        specification: '',
        size: '',
        weight: 1,
        temperature: -18,
        location: '',
        deviceId: ''
      },
      rules: {
        name: [{ required: true, message: '请输入海鲜名称', trigger: 'blur' }],
        deviceId: [{ required: true, message: '请输入设备标识', trigger: 'blur' }],
        type: [{ required: true, message: '请选择海鲜类型', trigger: 'change' }],
        specification: [{ required: true, message: '请选择等级', trigger: 'blur' }],
        size: [{ required: true, message: '请选择包装', trigger: 'blur' }],
        weight: [{ required: true, message: '请输入重量', trigger: 'blur' }],
        temperature: [{ required: true, message: '请输入存储温度', trigger: 'blur' }],
        location: [{ required: true, message: '请选择存储位置', trigger: 'change' }]
      }
    }
  },
  methods: {
    getTypeText(type) {
      const types = {
        shrimp: '冷冻虾类',
        fish: '冷冻鱼类',
        shellfish: '冷冻贝类',
        crab: '冷冻蟹类'
      }
      return types[type] || type
    },
    getStatusType(status) {
      const types = {
        normal: 'success',
        warning: 'warning',
        danger: 'danger'
      }
      return types[status] || 'info'
    },
    getStatusText(status) {
      const texts = {
        in_stock: '在库',
        out_stock: '已出库',
        mortgaged: '抵押中'
      }
      return texts[status] || status
    },
    handleSearch() {
      // TODO: 实现搜索逻辑
    },
    resetForm() {
      this.filterForm = {
        goodsId: '',
        type: '',
        status: ''
      }
    },
    handleAdd() {
      this.dialogTitle = '新增入库'
      this.dialogVisible = true
      this.goodsForm = {
        name: '',
        type: '',
        specification: '',
        size: '',
        weight: 1,
        temperature: -18,
        location: '',
        deviceId: ''
      }
    },
    handleOut() {
      // TODO: 实现出库逻辑
    },
    handleEdit(row) {
      this.dialogTitle = '编辑货物'
      this.dialogVisible = true
      this.goodsForm = { ...row }
    },
    handleDelete(row) {
      this.$confirm('确认删除该货物记录?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        // TODO: 实现删除逻辑
        this.$message.success('删除成功')
      }).catch(() => {})
    },
    handleSizeChange(val) {
      this.pageSize = val
      // TODO: 重新加载数据
    },
    handleCurrentChange(val) {
      this.currentPage = val
      // TODO: 重新加载数据
    },
    submitForm() {
      this.$refs.goodsForm.validate((valid) => {
        if (valid) {
          // TODO: 实现提交逻辑
          this.dialogVisible = false
          this.$message.success('操作成功')
        }
      })
    },
    scanDeviceId() {
      // TODO: 实现扫描NFC或条形码的逻辑
      this.$message.info('请将设备靠近NFC读取器或对准条形码扫描器')
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
