import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import router from './router'
import vuetify from './plugins/vuetify'
import store from './store'
import VueCookies from 'vue-cookies'
import alertError from './assets/common.js'

Vue.use(VueCookies)

Vue.config.productionTip = false
Vue.prototype.$http = axios
const protocol = 'http'
const address = process.env.VUE_APP_SERVERADDRESS
const port = process.env.VUE_APP_PORT

Vue.prototype.$serverAddress = `${protocol}://${address}:${port}`
Vue.prototype.$address = `${address}`

Vue.mixin({
  methods: {
    alertError
  }
})

new Vue({
  store,
  router,
  vuetify,
  render: h => h(App),
}).$mount('#app')
