<script>
export default {
  name: 'App'
}
</script>

<template>
  <router-view></router-view>
</template>

<style>
html, body {
  margin: 0;
  padding: 0;
  height: 100%;
  width: 100%;
  overflow-x: hidden;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  overflow-x: hidden;
}

.nav-container {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  height: 60px;
  background-color: #409EFF;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.nav-content {
  height: 100%;
  display: flex;
  align-items: center;
  padding: 0;
  margin: 0;
}

.nav-logo {
  display: flex;
  align-items: center;
  font-size: 20px;
  font-weight: bold;
  color: #fff;
  padding-left: 20px;
}

.nav-menu {
  display: flex;
  align-items: center;
  margin: 0;
  padding: 0;
  list-style: none;
  padding-left: 40px;
}

.nav-menu li {
  padding: 0 15px;
}

.nav-menu a {
  color: #fff;
  text-decoration: none;
  font-size: 16px;
}

.nav-menu a:hover,
.nav-menu a.router-link-active {
  color: #e6f1ff;
  font-weight: bold;
}

.nav-user {
  position: absolute;
  right: 20px;
  top: 50%;
  transform: translateY(-50%);
  display: flex;
  align-items: center;
}

.nav-user .el-button {
  margin-left: 15px;
  background: transparent;
  border: 1px solid #fff;
  color: #fff;
}

.nav-user .el-button:hover {
  background: rgba(255, 255, 255, 0.1);
}

@media screen and (max-width: 768px) {
  .nav-logo {
    padding-left: 10px;
  }

  .nav-menu {
    display: none;
  }

  .nav-user {
    right: 10px;
  }
}

/* 滚动条样式 */
::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

::-webkit-scrollbar-thumb {
  background: #ddd;
  border-radius: 3px;
}

::-webkit-scrollbar-track {
  background: #f5f5f5;
}
</style>
