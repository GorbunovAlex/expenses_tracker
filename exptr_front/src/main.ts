import { createApp } from 'vue'
import { createPinia } from 'pinia'

import './assets/main.scss';

import { Quasar, Notify } from 'quasar'

import '@quasar/extras/roboto-font-latin-ext/roboto-font-latin-ext.css'
import '@quasar/extras/material-icons/material-icons.css'
import '@quasar/extras/material-icons-outlined/material-icons-outlined.css'

import 'quasar/src/css/index.sass'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(Quasar, {
  plugins: {
    Notify
  }
})

app.mount('#app')
