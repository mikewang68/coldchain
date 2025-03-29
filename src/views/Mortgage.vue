<template>
  <div class="mortgage-page">
    <nav-bar></nav-bar>
    <div class="mortgage-container">
      <el-card>
        <div class="filter-container">
          <el-form :inline="true" :model="filterForm">
            <el-form-item label="订单编号">
              <el-input v-model="filterForm.orderId" placeholder="请输入订单编号"></el-input>
            </el-form-item>
            <el-form-item label="抵押状态">
              <el-select v-model="filterForm.status" placeholder="请选择状态">
                <el-option label="抵押中" value="mortgaged"></el-option>
                <el-option label="已赎回" value="redeemed"></el-option>
                <el-option label="已逾期" value="overdue"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="时间范围">
              <el-date-picker
                v-model="filterForm.dateRange"
                type="daterange"
                range-separator="至"
                start-placeholder="开始日期"
                end-placeholder="结束日期">
              </el-date-picker>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="handleSearch">查询</el-button>
              <el-button @click="resetForm">重置</el-button>
            </el-form-item>
          </el-form>
        </div>

        <div class="operation-container">
          <el-button type="primary" @click="handleAdd">新增抵押</el-button>
          <el-button type="success" @click="handleRedeem">办理赎回</el-button>
        </div>

        <el-table :data="mortgageList" border style="width: 100%">
          <el-table-column prop="orderId" label="订单编号" width="120"></el-table-column>
          <el-table-column prop="goodsName" label="抵押物品"></el-table-column>
          <el-table-column prop="mortgageAmount" label="抵押金额" width="120"></el-table-column>
          <el-table-column prop="mortgageDate" label="抵押日期" width="180"></el-table-column>
          <el-table-column prop="expireDate" label="到期日期" width="180"></el-table-column>
          <el-table-column prop="status" label="状态" width="100">
            <template #default="scope">
              <el-tag :type="getStatusType(scope.row.status)">
                {{ getStatusText(scope.row.status) }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="200">
            <template #default="scope">
              <el-button
                size="mini"
                type="success"
                @click="handleRedeem(scope.row)"
                v-if="scope.row.status === 'mortgaged'">
                赎回
              </el-button>
              <el-button size="mini" @click="handleDetail(scope.row)">详情</el-button>
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

      <!-- 新增抵押对话框 -->
      <el-dialog :title="dialogTitle" v-model="dialogVisible" width="600px">
        <el-form :model="mortgageForm" :rules="rules" ref="mortgageForm" label-width="100px">
          <el-form-item label="抵押物品" prop="goodsList">
            <el-select
              v-model="mortgageForm.goodsList"
              multiple
              filterable
              placeholder="请选择抵押物品"
              style="width: 100%">
              <el-option
                v-for="item in availableGoods"
                :key="item.goodsId"
                :label="item.name"
                :value="item.goodsId">
                <span>{{ item.name }}</span>
                <span style="float: right; color: #8492a6; font-size: 13px">
                  {{ item.weight }}吨 | {{ item.type }}
                </span>
              </el-option>
            </el-select>
          </el-form-item>

          <!-- 选中货物的重量编辑区域 -->
          <el-form-item label="货物重量" v-if="mortgageForm.goodsList.length > 0">
            <div v-for="goodsId in mortgageForm.goodsList" :key="goodsId" class="goods-weight-item">
              <span class="goods-name">{{ getGoodsName(goodsId) }}</span>
              <div class="weight-input-wrapper">
                <el-input-number
                  v-model="mortgageForm.goodsWeights[goodsId]"
                  :min="0"
                  :max="getGoodsMaxWeight(goodsId)"
                  :precision="2"
                  :step="0.1"
                  size="small"
                  @change="handleWeightChange">
                  <template #append>吨</template>
                </el-input-number>
                <span class="max-weight">最大: {{ getGoodsMaxWeight(goodsId) }}吨</span>
              </div>
            </div>
          </el-form-item>

          <el-form-item label="抵押金额" prop="mortgageAmount">
            <el-input-number v-model="mortgageForm.mortgageAmount" :min="1" :precision="2"></el-input-number>
          </el-form-item>
          <el-form-item label="抵押期限" prop="mortgagePeriod">
            <el-input-number v-model="mortgageForm.mortgagePeriod" :min="1" :max="36"></el-input-number>
            <span class="unit">月</span>
          </el-form-item>
          <el-form-item label="备注" prop="remark">
            <el-input type="textarea" v-model="mortgageForm.remark"></el-input>
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="submitForm">确定</el-button>
          </span>
        </template>
      </el-dialog>

      <!-- 赎回对话框 -->
      <el-dialog title="办理赎回" v-model="redeemDialogVisible" width="400px">
        <el-form :model="redeemForm" :rules="redeemRules" ref="redeemForm" label-width="100px">
          <el-form-item label="赎回金额" prop="redeemAmount">
            <el-input-number v-model="redeemForm.redeemAmount" :min="1" :precision="2"></el-input-number>
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="redeemDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="submitRedeem">确定</el-button>
          </span>
        </template>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import NavBar from '../components/Layout/NavBar.vue'

export default {
  name: 'MortgageView',
  components: {
    NavBar
  },
  data() {
    return {
      filterForm: {
        orderId: '',
        status: '',
        dateRange: []
      },
      mortgageList: [
        {
          orderId: 'M20240320001',
          goodsList: ['SF20240320001', 'SF20240319002'],
          goodsName: '阿根廷红虾、加拿大北极甜虾',
          mortgageAmount: 255.5,
          mortgageDate: '2024-03-20',
          expireDate: '2024-06-20',
          status: 'mortgaged'
        },
        {
          orderId: 'M20240319002',
          goodsList: ['SF20240318003'],
          goodsName: '智利三文鱼',
          mortgageAmount: 182.6,
          mortgageDate: '2024-03-19',
          expireDate: '2024-06-19',
          status: 'mortgaged'
        },
        {
          orderId: 'M20240318003',
          goodsList: ['SF20240317004'],
          goodsName: '大连扇贝',
          mortgageAmount: 125.8,
          mortgageDate: '2024-03-18',
          expireDate: '2024-06-18',
          status: 'mortgaged'
        },
        {
          orderId: 'M20240317004',
          goodsList: ['SF20240314007'],
          goodsName: '波士顿龙虾',
          mortgageAmount: 102.4,
          mortgageDate: '2024-03-17',
          expireDate: '2024-06-17',
          status: 'redeemed'
        },
        {
          orderId: 'M20240316005',
          goodsList: ['SF20240316005'],
          goodsName: '帝王蟹',
          mortgageAmount: 158.2,
          mortgageDate: '2024-03-16',
          expireDate: '2024-06-16',
          status: 'mortgaged'
        },
        {
          orderId: 'M20240315006',
          goodsList: ['SF20240315006'],
          goodsName: '青岛鲅鱼',
          mortgageAmount: 86.5,
          mortgageDate: '2024-03-15',
          expireDate: '2024-06-15',
          status: 'overdue'
        },
        {
          orderId: 'M20240314007',
          goodsList: ['SF20240314007'],
          goodsName: '波士顿龙虾',
          mortgageAmount: 205.3,
          mortgageDate: '2024-03-14',
          expireDate: '2024-06-14',
          status: 'redeemed'
        }
      ],
      availableGoods: [
        {
          goodsId: 'SF20240320001',
          name: '阿根廷红虾',
          type: '冷冻虾类',
          weight: 25.5,
          status: 'in_stock'
        },
        {
          goodsId: 'SF20240319002',
          name: '加拿大北极甜虾',
          type: '冷冻虾类',
          weight: 18.2,
          status: 'in_stock'
        },
        {
          goodsId: 'SF20240318003',
          name: '智利三文鱼',
          type: '冷冻鱼类',
          weight: 15.8,
          status: 'in_stock'
        },
        {
          goodsId: 'SF20240317004',
          name: '大连扇贝',
          type: '冷冻贝类',
          weight: 12.5,
          status: 'in_stock'
        },
        {
          goodsId: 'SF20240316005',
          name: '帝王蟹',
          type: '冷冻蟹类',
          weight: 8.6,
          status: 'in_stock'
        },
        {
          goodsId: 'SF20240315006',
          name: '青岛鲅鱼',
          type: '冷冻鱼类',
          weight: 20.5,
          status: 'in_stock'
        },
        {
          goodsId: 'SF20240314007',
          name: '波士顿龙虾',
          type: '冷冻贝类',
          weight: 10.2,
          status: 'in_stock'
        }
      ],
      currentPage: 1,
      pageSize: 10,
      total: 100,
      dialogVisible: false,
      dialogTitle: '新增抵押',
      mortgageForm: {
        goodsList: [],
        goodsWeights: {},
        mortgageAmount: 1000,
        mortgagePeriod: 3,
        remark: ''
      },
      rules: {
        goodsList: [{ required: true, message: '请选择抵押物品', trigger: 'change' }],
        mortgageAmount: [{ required: true, message: '请输入抵押金额', trigger: 'blur' }],
        mortgagePeriod: [{ required: true, message: '请输入抵押期限', trigger: 'blur' }],
        goodsWeights: [{ required: true, message: '请输入货物重量', trigger: 'change' }]
      },
      redeemDialogVisible: false,
      redeemForm: {
        redeemAmount: 0
      },
      redeemRules: {
        redeemAmount: [{ required: true, message: '请输入赎回金额', trigger: 'blur' }]
      }
    }
  },
  methods: {
    getStatusType(status) {
      const types = {
        mortgaged: 'warning',
        redeemed: 'success',
        overdue: 'danger'
      }
      return types[status] || 'info'
    },
    getStatusText(status) {
      const texts = {
        mortgaged: '抵押中',
        redeemed: '已赎回',
        overdue: '已逾期'
      }
      return texts[status] || status
    },
    handleSearch() {
      // TODO: 实现搜索逻辑
    },
    resetForm() {
      this.filterForm = {
        orderId: '',
        status: '',
        dateRange: []
      }
    },
    handleAdd() {
      this.dialogTitle = '新增抵押'
      this.dialogVisible = true
      this.mortgageForm = {
        goodsList: [],
        goodsWeights: {},
        mortgageAmount: 1000,
        mortgagePeriod: 3,
        remark: ''
      }
    },
    handleRedeem(row) {
      if (row) {
        this.redeemForm.redeemAmount = row.mortgageAmount
      }
      this.redeemDialogVisible = true
    },
    handleDetail(row) {
      // TODO: 实现查看详情逻辑
    },
    handleSizeChange(val) {
      this.pageSize = val
      // TODO: 重新加载数据
    },
    handleCurrentChange(val) {
      this.currentPage = val
      // TODO: 重新加载数据
    },
    getGoodsName(goodsId) {
      const goods = this.availableGoods.find(item => item.goodsId === goodsId)
      return goods ? goods.name : ''
    },
    getGoodsMaxWeight(goodsId) {
      const goods = this.availableGoods.find(item => item.goodsId === goodsId)
      return goods ? goods.weight : 0
    },
    handleWeightChange(value, goodsId) {
      // 更新货物重量
      this.mortgageForm.goodsWeights[goodsId] = value
      // 重新计算总抵押金额
      this.calculateTotalAmount()
      // 触发表单验证
      this.$refs.mortgageForm.validateField('goodsWeights')
    },
    calculateTotalAmount() {
      // 根据货物重量计算总金额（示例：每吨10000元）
      const totalWeight = Object.values(this.mortgageForm.goodsWeights).reduce((sum, weight) => sum + weight, 0)
      this.mortgageForm.mortgageAmount = totalWeight * 10000
    },
    submitForm() {
      // 验证所有选中货物是否都已设置重量
      const hasEmptyWeight = this.mortgageForm.goodsList.some(goodsId => {
        return !this.mortgageForm.goodsWeights[goodsId] || this.mortgageForm.goodsWeights[goodsId] <= 0
      })

      if (hasEmptyWeight) {
        this.$message.error('请为所有货物设置重量')
        return
      }

      this.$refs.mortgageForm.validate((valid) => {
        if (valid) {
          // 获取选中货物的名称和重量列表
          const selectedGoods = this.mortgageForm.goodsList.map(goodsId => {
            const goods = this.availableGoods.find(item => item.goodsId === goodsId)
            return {
              name: goods ? goods.name : '',
              weight: this.mortgageForm.goodsWeights[goodsId] || 0
            }
          }).filter(item => item.name)

          // TODO: 实现提交逻辑
          this.dialogVisible = false
          this.$message.success('操作成功')
        }
      })
    },
    submitRedeem() {
      this.$refs.redeemForm.validate((valid) => {
        if (valid) {
          // TODO: 实现赎回逻辑
          this.redeemDialogVisible = false
          this.$message.success('赎回成功')
        }
      })
    }
  }
}
</script>

<style scoped>
.mortgage-page {
  width: 100%;
  height: 100vh;
  overflow: hidden;
}

.mortgage-container {
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
  margin-left: 10px;
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

  .el-form--inline .el-form-item {
    margin-right: 10px;
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
  max-height: calc(90vh - 150px);
}

/* 货物重量编辑区域样式 */
.goods-weight-item {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
  padding: 8px;
  background-color: #f5f7fa;
  border-radius: 4px;
}

.goods-weight-item:last-child {
  margin-bottom: 0;
}

.goods-name {
  flex: 1;
  margin-right: 10px;
  color: #606266;
}

.weight-input-wrapper {
  display: flex;
  align-items: center;
  gap: 10px;
}

.max-weight {
  color: #909399;
  font-size: 13px;
}

:deep(.el-input-number) {
  width: 150px;
}

:deep(.el-input-number .el-input__inner) {
  text-align: right;
}

:deep(.el-input-number .el-input-group__append) {
  padding: 0 8px;
  background-color: #f5f7fa;
  color: #606266;
}
</style>
