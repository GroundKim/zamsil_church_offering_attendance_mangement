const state = {
    students: null,
    payloadStudents: [],
    payloadOfferings: [],
    sendPost: false,
    
}

const mutations = {
    setStudents(state, newStudents) {
        state.students = newStudents
    },
    
    pushStudent(state, student) {
        state.payloadStudents.push(student)
    },

    pushOffering(state, newOffering) {
        state.payloadOfferings.push(newOffering)
    },

    changeSendPost(state, bool) {
        state.sendPost = bool
    }
}

const actions = {

}

const getters = {
    getStudents (state) {
        return state.students
    },

    getOfferingPayload(state) {
        return state.payloadOfferings
    }
}

export default {
    state,
    mutations,
    actions,
    getters
}