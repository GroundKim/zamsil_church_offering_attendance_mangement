import Vue from 'vue'
import Vuex from 'vuex'
import specificOfferingStudents from './modules/specificOfferingStudents'

Vue.use(Vuex)

const store = new Vuex.Store({
    modules: {
        specificOfferingStudents,
    },
  })

export default store
  