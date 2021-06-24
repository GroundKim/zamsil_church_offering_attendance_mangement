import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import router from './router'
import vuetify from './plugins/vuetify'
import store from './store'
import errorHandler from './assets/common.js'
import VueTheMask from 'vue-the-mask'
import Vuelidate from 'vuelidate'

Vue.config.productionTip = false
Vue.prototype.$http = axios
//const protocol = 'http'
const protocol = process.env.VUE_APP_PROTOCOL
const address = process.env.VUE_APP_SERVERADDRESS
const port = process.env.VUE_APP_PORT

Vue.prototype.$serverAddress = `${protocol}://${address}:${port}`
Vue.prototype.$address = `${address}`

Vue.use(VueTheMask)
Vue.use(Vuelidate)

Vue.mixin({
  methods: {
    errorHandler
  }
})

new Vue({
  store,
  router,
  vuetify,
  render: h => h(App),
}).$mount('#app')
