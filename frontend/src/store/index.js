import Vue from 'vue'
import Vuex from 'vuex'
import specificOfferingStudents from './modules/specificOfferingStudents'
Vue.use(Vuex)

const store = new Vuex.Store({
    modules: {
        specificOfferingStudents,
    },

    state: {
      headerName: null
    },

    mutations: {
      changeHeaderName (state, headerName) {
        state.headerName = headerName
      }
    }
  })

export default store
  