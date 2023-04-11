<template>
  <a-layout-sider breakpoint="lg" v-model="collapsed">
    <div class="log">
      <span>{{ collapsed ? 'Blog' : 'My Blog' }}</span>
    </div>

    <a-menu theme="dark" mode="inline" @click="goToPage">
      <a-menu-item key="index">
        <a-icon type="dashboard"></a-icon>
        <span>仪表盘</span>
      </a-menu-item>

        <a-menu-item key="artlist">
          <a-icon type="snippets" />
          <span>文章管理</span>
        </a-menu-item>
      
      <!-- <a-sub-menu>
        <span slot="title">
          <a-icon type="file" />
          <span>内容审核</span>
        </span>
        <a-menu-item key="artaudit">
          <a-icon type="snippets" />
          <span>文章审核</span>
        </a-menu-item>
      </a-sub-menu> -->

      <a-menu-item key="commentlist">
        <a-icon type="message" />
        <span>评论审核</span>
      </a-menu-item>

      <a-menu-item key="catelist">
        <a-icon type="book" />
        <span>标签列表</span>
      </a-menu-item>

      <a-menu-item key="userlist">
        <a-icon type="user" />
        <span>用户列表</span>
      </a-menu-item>  

      <a-menu-item key="logsystem">
        <a-icon type="area-chart" />
        <span>日志系统</span>
      </a-menu-item>

    </a-menu>
  </a-layout-sider>
</template>

<script>
export default {
  data() {
    return {
      collapsed: false,
    }
  },
  methods: {
    goToPage(item) {
      this.$router.push('/' + item.key).catch((err) => err)
    },
  },
}
</script>

<style scoped>
.log {
  height: 32px;
  margin: 16px;
  background-color: #fff;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 17px;
}
</style>
