import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import router from './router'
import vuetify from './plugins/vuetify'
import store from './store'
import VueCookies from 'vue-cookies'

Vue.use(VueCookies)

Vue.config.productionTip = false
Vue.prototype.$http = axios
const protocol = 'http'
const address = process.env.VUE_APP_SERVERADDRESS
const port = process.env.VUE_APP_PORT

Vue.prototype.$serverAddress = `${protocol}://${address}:${port}`
Vue.prototype.$address = `${address}`
// methods
function alertError (err) {
  let errStatusCode = err.response.status
  if (errStatusCode === 404) {
    alert(err.message + "데이터를 불러오는 도중 오류가 발생하였습니다 관리자에게 문의하십시오")
  }

  if (errStatusCode === 403) {
    alert("로그인을 해주세요")
    this.$router.push("/login")
  }
}

new Vue({
  store,
  router,
  vuetify,

  methods: {
    alertError: alertError
  },

  render: h => h(App),
}).$mount('#app')
