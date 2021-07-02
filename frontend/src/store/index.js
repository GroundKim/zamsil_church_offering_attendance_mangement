import Vue from 'vue'
import Vuex from 'vuex'
import specificOfferingStudents from './modules/specificOfferingStudents'
Vue.use(Vuex)

const store = new Vuex.Store({
    modules: {
      specificOfferingStudents,
    },

    state: {
      isLogin: false
    },

    mutations: {
      changeLoginStatus: (state, status) => {
        state.isLogin = status
      }
    }
  })

export default store
  