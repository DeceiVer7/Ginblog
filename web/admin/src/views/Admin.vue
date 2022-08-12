<template>
  <a-layout class="container">
    <headerTemplate></headerTemplate>
    <a-button class="headerBtn" type="primary" status="danger" @click="loginOut"
      >退出</a-button
    >
    <navTemplate></navTemplate>
    <footerTemplate></footerTemplate>
  </a-layout>
</template>

<script>
import footerTemplate from '@/components/admin/Footer.vue'
import headerTemplate from '@/components/admin/Header.vue'
import navTemplate from '@/components/admin/Nav.vue'

export default {
  components: {
    footerTemplate,
    headerTemplate,
    navTemplate
  },
  methods: {
    loginOut() {
      window.sessionStorage.clear('token')
      this.$message.success('退出成功')
      this.$router.push('/login')
    }
  }
}
</script>

<style scoped>
.container {
  height: 100%;
}
.content {
  position: absolute;
  margin: 15px;
  top: 30px;
  width: 100px;
  left: 270px;
  color: #fff;
  font-size: 16px;
  font-stretch: condensed;
  text-align: center;
  background-color: rgb(194, 127, 41);
}

.headerBtn {
  position: absolute;
  right: 4px;
  top: 5px;
}
</style>
