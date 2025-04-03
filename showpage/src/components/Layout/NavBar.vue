<template>
  <div class="nav-container">
    <div class="nav-content">
      <div class="nav-left">
        <router-link to="/" class="nav-logo">
          区块链商品交易管理系统
        </router-link>
        <div class="nav-menu">
          <router-link to="/" class="nav-item">首页</router-link>
          <router-link to="/goods" class="nav-item">商品监管</router-link>
          <router-link to="/mortgage" class="nav-item">抵押管理</router-link>
          <router-link to="/orders" class="nav-item">解押出库</router-link>
          <router-link to="/items" class="nav-item">品类管理</router-link>
        </div>
      </div>
      <div class="nav-user">
        <template v-if="!isLoggedIn">
          <router-link to="/login">
            <el-button>登录</el-button>
          </router-link>
        </template>
        <template v-else>
          <el-dropdown @command="handleCommand">
            <span class="el-dropdown-link">
              {{ username }}<i class="el-icon-arrow-down el-icon--right"></i>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="logout">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </template>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'NavBar',
  data() {
    return {
      isLoggedIn: false,
      username: ''
    }
  },
  created() {
    const token = localStorage.getItem('token')
    const username = localStorage.getItem('username')
    if (token && username) {
      this.isLoggedIn = true
      this.username = username
    }
  },
  methods: {
    handleCommand(command) {
      if (command === 'logout') {
        localStorage.removeItem('token')
        localStorage.removeItem('username')
        this.isLoggedIn = false
        this.username = ''
        this.$router.push('/login')
      }
    }
  }
}
</script>

<style scoped>
.nav-container {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  height: 60px;
  background: linear-gradient(135deg, #409EFF 0%, #2c5282 100%);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  z-index: 1000;
}

.nav-content {
  height: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
}

.nav-left {
  display: flex;
  align-items: center;
}

.nav-logo {
  color: #fff;
  font-size: 20px;
  font-weight: bold;
  text-decoration: none;
  margin-right: 40px;
  background: linear-gradient(45deg, #ffffff, #e3f2fd);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.nav-menu {
  display: flex;
  align-items: center;
}

.nav-item {
  color: #fff;
  text-decoration: none;
  padding: 0 20px;
  font-size: 16px;
  transition: all 0.3s ease;
  position: relative;
  opacity: 0.9;
}

.nav-item:hover,
.nav-item.router-link-active {
  color: #e3f2fd;
  opacity: 1;
  transform: translateY(-1px);
}

.nav-item::after {
  content: '';
  position: absolute;
  bottom: -2px;
  left: 50%;
  width: 0;
  height: 2px;
  background: #e3f2fd;
  transition: all 0.3s ease;
  transform: translateX(-50%);
}

.nav-item:hover::after,
.nav-item.router-link-active::after {
  width: 80%;
}

.nav-user {
  display: flex;
  align-items: center;
}

.nav-user .el-button {
  background: transparent;
  border: 1px solid #e3f2fd;
  color: #fff;
  transition: all 0.3s ease;
  width: 40px;
  height: 40px;
  padding: 0;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
}

.nav-user .el-button:hover {
  background: rgba(227, 242, 253, 0.1);
  border-color: #fff;
  transform: translateY(-1px);
}

.el-dropdown-link {
  color: #fff;
  cursor: pointer;
  display: flex;
  align-items: center;
  padding: 8px 12px;
  border-radius: 4px;
  transition: all 0.3s ease;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.1);
  justify-content: center;
  font-size: 16px;
}

.el-dropdown-link:hover {
  background: rgba(227, 242, 253, 0.2);
}

.el-dropdown-link i {
  margin-left: 0;
  transition: transform 0.3s ease;
}

.el-dropdown-link:hover i {
  transform: translateY(1px);
}

/* 响应式布局调整 */
@media screen and (max-width: 1366px) {
  .nav-logo {
    font-size: 18px;
  }

  .nav-item {
    padding: 0 15px;
    font-size: 14px;
  }

  .nav-user .el-button,
  .el-dropdown-link {
    width: 36px;
    height: 36px;
    font-size: 14px;
  }
}

@media screen and (min-width: 1920px) {
  .nav-logo {
    font-size: 22px;
  }

  .nav-item {
    padding: 0 25px;
    font-size: 16px;
  }

  .nav-user .el-button,
  .el-dropdown-link {
    width: 44px;
    height: 44px;
    font-size: 18px;
  }
}
</style>
