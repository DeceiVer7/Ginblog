// import { createApp } from 'vue'
// import App from './App.vue'
// import router from './router'
// import { Button } from 'ant-design-vue'

// const app = createApp(App)
// app.config.productionTip = false
// app.use(Button)
//   .use(router)
//   .mount('#app')

import { createApp } from 'vue'

import router from './router'
import axios from 'axios'
import ArcoVue from '@arco-design/web-vue'
import App from './App.vue'
import '@arco-design/web-vue/dist/arco.css'
import './assets/css/style.css'

const app = createApp(App)
axios.defaults.baseURL = 'http://localhost:3000/api/v1'
app.config.globalProperties.$http = axios

app.config.productionTip = false

app.use(ArcoVue).use(router).mount('#app')
