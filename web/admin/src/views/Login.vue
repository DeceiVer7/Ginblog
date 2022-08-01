<template>
  <div class="container">
    <particles></particles>
    <div class="loginBox">
      <a-form :model="form" class="loginForm">
        <a-form-item>
          <a-input :style="{ width: '320px' }" placeholder="请输入您的账号">
            <template #prefix>
              <icon-user />
            </template>
          </a-input>
        </a-form-item>

        <a-form-item>
          <a-input-password
            :style="{ width: '320px' }"
            placeholder="请输入您的密码"
            allow-clear
            ><template #prefix>
              <icon-lock />
            </template>
          </a-input-password>
        </a-form-item>

        <a-form-item class="loginBtn">
          <a-button type="primary" shape="round" :style="{ margin: '10px' }"
            >登录
          </a-button>
          <a-button type="outline" shape="round">注册</a-button>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<style scoped>
.container {
  height: 100%;
  background-color: #282c34;
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
</style>

<script>
import { IconUser, IconLock } from '@arco-design/web-vue/es/icon'
import particles from '@/components/particles/Particles.vue'
export default {
  components: { IconUser, IconLock, particles }
}
</script>
