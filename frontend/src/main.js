import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'

// import axios from 'axios'

// axios.defaults.baseURL = 'http://127.0.0.1:8080/';


loadFonts()

createApp(App)
  .use(router)
  .use(store)
  .use(vuetify)
  .mount('#app')
