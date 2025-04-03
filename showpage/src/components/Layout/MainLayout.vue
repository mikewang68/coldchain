<template>
  <div class="main-layout">
    <el-container>
      <el-header>
        <div class="header-content">
          <div class="logo">区块链商品交易管理系统</div>
          <el-menu
            mode="horizontal"
            :router="true"
            :default-active="activeMenu"
            background-color="#409EFF"
            text-color="#fff"
            active-text-color="#ffd04b"
          >
            <el-menu-item index="/">首页</el-menu-item>
            <el-menu-item index="/goods" v-if="isLogin">货品监管</el-menu-item>
            <el-menu-item index="/mortgage" v-if="isLogin">抵押管理</el-menu-item>
          </el-menu>
          <div class="user-info">
            <template v-if="isLogin">
              <el-dropdown @command="handleCommand">
                <span class="user-avatar">
                  <el-avatar :size="32" icon="el-icon-user"></el-avatar>
                  <span class="username">{{ username }}</span>
                </span>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item command="logout">退出登录</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </template>
            <template v-else>
              <el-button type="text" class="login-btn" @click="goToLogin">
                <el-avatar :size="32" icon="el-icon-user"></el-avatar>
                <span>登录</span>
              </el-button>
            </template>
          </div>
        </div>
      </el-header>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  name: 'MainLayout',
  data() {
    return {
      activeMenu: '/',
      isLogin: false,
      username: ''
    }
  },
  created() {
    // 检查登录状态
    this.checkLoginStatus()
  },
  methods: {
    checkLoginStatus() {
      const token = localStorage.getItem('token')
      const username = localStorage.getItem('username')
      this.isLogin = !!token
      this.username = username || ''
    },
    handleCommand(command) {
      if (command === 'logout') {
        this.handleLogout()
      }
    },
    handleLogout() {
      localStorage.removeItem('token')
      localStorage.removeItem('username')
      this.isLogin = false
      this.username = ''
      this.$message.success('退出成功')
      if (this.$route.meta.requiresAuth) {
        this.$router.push('/')
      }
    },
    goToLogin() {
      this.$router.push({
        path: '/login',
        query: {
          redirect: this.$route.fullPath
        }
      })
    }
  },
  watch: {
    $route: {
      immediate: true,
      handler(route) {
        this.activeMenu = route.path
      }
    }
  }
}
</script>

<style scoped>
.main-layout {
  height: 100vh;
}

.el-container {
  height: 100%;
}

.el-header {
  padding: 0;
  background-color: #409EFF;
}

.header-content {
  height: 60px;
  display: flex;
  align-items: center;
  padding: 0 20px;
}

.logo {
  color: white;
  font-size: 20px;
  font-weight: bold;
  margin-right: 40px;
}

.user-info {
  margin-left: auto;
  color: white;
  display: flex;
  align-items: center;
}

.user-avatar {
  display: flex;
  align-items: center;
  cursor: pointer;
}

.username {
  margin-left: 8px;
  color: white;
}

.login-btn {
  display: flex;
  align-items: center;
  color: white;
}

.login-btn span {
  margin-left: 8px;
}

.el-main {
  padding: 20px;
  background-color: #f0f2f5;
  overflow-y: auto;
}

:deep(.el-menu) {
  border-bottom: none;
  background-color: transparent;
}

:deep(.el-menu--horizontal > .el-menu-item) {
  height: 60px;
  line-height: 60px;
  border-bottom: none;
}

:deep(.el-menu-item.is-active) {
  background-color: rgba(255, 255, 255, 0.1);
}

:deep(.el-dropdown-menu__item) {
  padding: 8px 20px;
}
</style>
