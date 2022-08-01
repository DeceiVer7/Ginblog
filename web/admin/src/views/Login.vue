<template>
  <div class="container">
    <particles></particles>
    <nav id="Title">Simplicity博客系统</nav>
    <div class="loginBox">
      <a-form :model="formdata" class="loginForm">
        <a-form-item
          field="username"
          :rules="[
            { required: true, message: '账号不能为空' },
            { minLength: 5, maxLength: 15, message: '账号长度在5~15位' }
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-input
            v-model="formdata.username"
            :style="{ width: '320px' }"
            placeholder="请输入您的账号"
          >
            <template #prefix>
              <icon-user />
            </template>
          </a-input>
        </a-form-item>

        <a-form-item
          field="password"
          :rules="[
            { required: true, message: '密码不能为空' },
            { minLength: 6, message: '密码长度应大于6位' }
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-input-password
            v-model="formdata.password"
            :style="{ width: '320px' }"
            placeholder="请输入您的密码"
            allow-clear
            ><template #prefix>
              <icon-lock />
            </template>
          </a-input-password>
        </a-form-item>

        <a-form-item class="loginBtn">
          <!-- eslint-disable-next-line vue/no-parsing-error -->
          <a-button type="primary" shape="round" :style="{ margin: '10px' }"
            >登录</a-button
          >
          <a-button type="outline" shape="round">注册</a-button>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<style scoped>
.container {
  height: 100%;
  background-color: #121a2a;
}
.loginBox {
  width: 425px;
  height: 260px;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: #fff;
  border-radius: 5px;
}
.loginForm {
  width: 100%;
  position: absolute;
  bottom: 10%;
  padding: 0 20px;
  box-sizing: border-box;
  right: 35px;
}
.loginBtn {
  position: relative;
  justify-content: center;
  left: 17%;
}
#Title {
  color: #fff;
  position: absolute;
  top: 18%;
  left: 50%;
  transform: translate(-50%, -50%);
  font-family: Arial, Helvetica, sans-serif;
  font-size: 60px;
}
/* #backgroundImg {
  background: url('');
  width: 100%;
  height: 100%;
  position: fixed;
  background-size: 100% 100%;
} */
</style>

<script>
import { IconUser, IconLock } from '@arco-design/web-vue/es/icon'
import particles from '@/components/particles/Particles.vue'

export default {
  components: { IconUser, IconLock, particles },
  data() {
    return {
      formdata: {
        username: '',
        password: ''
      }
    }
  }
}
</script>
