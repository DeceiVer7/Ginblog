<template>
  <div class="container">
    <particles></particles>
    <nav id="Title">Simplicity博客系统</nav>
    <div class="loginBox">
      <a-form :model="formdata" class="loginForm" ref="loginFormRef">
        <a-form-item
          field="username"
          :rules="[
            { required: false, message: '账号不能为空' },
            {
              minLength: 5,
              maxLength: 15,
              message: '账号长度必须在5~15个字符之间'
            }
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
            { required: false, message: '密码不能为空' },
            {
              minLength: 6,
              maxLength: 20,
              message: '密码长度必须在6~20个字符之间'
            }
          ]"
          :validate-trigger="['change', 'input']"
        >
          <a-input-password
            v-model="formdata.password"
            :style="{ width: '320px' }"
            placeholder="请输入您的密码"
            v-on:keyup.enter="login"
            allow-clear
            ><template #prefix>
              <icon-lock />
            </template>
          </a-input-password>
        </a-form-item>

        <a-form-item class="loginBtn">
          <!-- eslint-disable-next-line vue/no-parsing-error -->
          <a-button
            type="primary"
            shape="round"
            :style="{ margin: '10px' }"
            @click="login"
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
  background-color: #003472;
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
import { getCurrentInstance } from 'vue'

export default {
  components: { IconUser, IconLock, particles },
  data() {
    return {
      formdata: {
        username: '',
        password: ''
      }
    }
  },
  methods: {
    login() {
      const proxy = getCurrentInstance()
      this.$refs.loginFormRef.validate(async (valid) => {
        if (valid) {
          this.$message.error('输入非法数据，请重新输入')
        }

        const { data: res } = await proxy.$http.post('login', this.formdata)
        if (res.status !== 200) return this.$message.error(res.message)
        window.sessionStorage.setItem('token', res.token)
        this.$router.push('admin')
      })
    }
  }
}
</script>
