import Vue from 'vue'
import Vuex from 'vuex'
import specificOfferingStudents from './modules/specificOfferingStudents'
Vue.use(Vuex)

const store = new Vuex.Store({
    modules: {
      specificOfferingStudents,
    },

    state: {
      isLogin: false,
      headerActiveTabName: null
    },

    mutations: {
      changeLoginStatus: (state, status) => {
        state.isLogin = status
      },

      changeHeaderActiveTabName: (state, activeTabName) => {
        state.headerActiveTabName = activeTabName
      }
    }
  })

export default store
  