import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import router from './router'
import vuetify from './plugins/vuetify'
import store from './store'


Vue.config.productionTip = false
Vue.prototype.$http = axios

const protocol = 'http'
const host = 'localhost'
const port = '8080'

Vue.prototype.$serverAddress = `${protocol}://${host}:${port}`

new Vue({
  store,
  router,
  vuetify,
  render: h => h(App),
}).$mount('#app')
