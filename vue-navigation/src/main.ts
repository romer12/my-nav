import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import { setupScrollbarStyle } from './plugins'

import '@/styles/global.scss'

const app = createApp(App)

// 设置顶部进度条样式
setupScrollbarStyle()

app.use(createPinia())
app.use(router)

app.mount('#app')
