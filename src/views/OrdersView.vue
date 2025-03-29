<template>
  <div class="orders-page">
    <nav-bar></nav-bar>
    <div class="orders-container">
      <el-card class="orders-card">
        <!-- 搜索过滤区域 -->
        <div class="filter-container">
          <el-form :inline="true" :model="filterForm" class="form-inline">
            <el-form-item label="批次编号">
              <el-input v-model="filterForm.orderId" placeholder="请输入批次编号" clearable></el-input>
            </el-form-item>
            <el-form-item label="海鲜类型">
              <el-select v-model="filterForm.seafoodType" placeholder="请选择" clearable>
                <el-option label="阿根廷红虾" value="阿根廷红虾"></el-option>
                <el-option label="加拿大北极甜虾" value="加拿大北极甜虾"></el-option>
                <el-option label="智利三文鱼" value="智利三文鱼"></el-option>
                <el-option label="波士顿龙虾" value="波士顿龙虾"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="状态">
              <el-select v-model="filterForm.status" placeholder="请选择" clearable>
                <el-option label="待审核" value="待审核"></el-option>
                <el-option label="已审核" value="已审核"></el-option>
                <el-option label="已出库" value="已出库"></el-option>
                <el-option label="已取消" value="已取消"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="handleSearch">查询</el-button>
              <el-button @click="resetForm">重置</el-button>
            </el-form-item>
          </el-form>
        </div>

        <!-- 操作按钮区域 -->
        <div class="operation-container">
          <el-button type="primary" @click="handleApply">申请解押</el-button>
          <el-button type="success" @click="handleExport">导出记录</el-button>
        </div>

        <!-- 表格区域 -->
        <el-table :data="ordersList" style="width: 100%" border>
          <el-table-column prop="orderId" label="批次编号" width="180"></el-table-column>
          <el-table-column prop="seafoodName" label="海鲜名称" width="150"></el-table-column>
          <el-table-column prop="weight" label="重量(吨)" width="100"></el-table-column>
          <el-table-column prop="mortgageAmount" label="抵押金额(万元)" width="150"></el-table-column>
          <el-table-column prop="applyDate" label="申请日期" width="180"></el-table-column>
          <el-table-column prop="releaseDate" label="解押日期" width="180"></el-table-column>
          <el-table-column prop="status" label="状态" width="100">
            <template #default="scope">
              <el-tag :type="getStatusType(scope.row.status)">{{ scope.row.status }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="操作" fixed="right" width="200">
            <template #default="scope">
              <el-button type="text" size="small" @click="handleDetail(scope.row)">详情</el-button>
              <el-button type="text" size="small" @click="handleEdit(scope.row)" v-if="scope.row.status === '待审核'">编辑</el-button>
              <el-button type="text" size="small" @click="handleCancel(scope.row)" v-if="scope.row.status === '待审核'">取消</el-button>
              <el-button type="text" size="small" @click="handleRelease(scope.row)" v-if="scope.row.status === '已审核'">出库</el-button>
            </template>
          </el-table-column>
        </el-table>

        <!-- 分页区域 -->
        <div class="pagination-container">
          <el-pagination
            background
            layout="total, sizes, prev, pager, next, jumper"
            :total="total"
            :page-size="pageSize"
            :page-sizes="[10, 20, 50, 100]"
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange">
          </el-pagination>
        </div>
      </el-card>
    </div>
  </div>
</template>

<script>
import NavBar from '../components/Layout/NavBar.vue'

export default {
  name: 'OrdersView',
  components: {
    NavBar
  },
  data() {
    return {
      filterForm: {
        orderId: '',
        seafoodType: '',
        status: ''
      },
      ordersList: [
        {
          orderId: 'SF20240320001',
          seafoodName: '阿根廷红虾',
          weight: 25.5,
          mortgageAmount: 255,
          applyDate: '2024-03-20 10:00:00',
          releaseDate: '',
          status: '待审核'
        },
        {
          orderId: 'SF20240319002',
          seafoodName: '加拿大北极甜虾',
          weight: 18.2,
          mortgageAmount: 182,
          applyDate: '2024-03-19 15:30:00',
          releaseDate: '2024-03-19 16:30:00',
          status: '已审核'
        },
        {
          orderId: 'SF20240318003',
          seafoodName: '智利三文鱼',
          weight: 15.8,
          mortgageAmount: 158,
          applyDate: '2024-03-18 09:15:00',
          releaseDate: '2024-03-18 14:20:00',
          status: '已出库'
        },
        {
          orderId: 'SF20240317004',
          seafoodName: '波士顿龙虾',
          weight: 12.5,
          mortgageAmount: 125,
          applyDate: '2024-03-17 14:20:00',
          releaseDate: '',
          status: '已取消'
        }
      ],
      total: 100,
      pageSize: 10,
      currentPage: 1
    }
  },
  methods: {
    getStatusType(status) {
      const types = {
        pending: 'warning',
        approved: 'success',
        rejected: 'danger'
      }
      return types[status] || 'info'
    },
    handleSearch() {
      // 实现搜索逻辑
      console.log('搜索条件：', this.filterForm)
    },
    resetForm() {
      this.filterForm = {
        orderId: '',
        seafoodType: '',
        status: ''
      }
    },
    handleApply() {
      // 实现申请解押逻辑
      console.log('申请解押')
    },
    handleExport() {
      // 实现导出记录逻辑
      console.log('导出记录')
    },
    handleDetail(row) {
      // 实现查看详情逻辑
      console.log('查看详情：', row)
    },
    handleEdit(row) {
      // 实现编辑逻辑
      console.log('编辑：', row)
    },
    handleCancel(row) {
      // 实现取消逻辑
      console.log('取消：', row)
    },
    handleRelease(row) {
      // 实现出库逻辑
      console.log('出库：', row)
    },
    handleSizeChange(val) {
      this.pageSize = val
      // 重新加载数据
    },
    handleCurrentChange(val) {
      this.currentPage = val
      // 重新加载数据
    }
  }
}
</script>

<style scoped>
.orders-page {
  width: 100%;
  height: 100vh;
  overflow: hidden;
}

.orders-container {
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

.orders-card {
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

/* 响应式布局 */
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
</style>
