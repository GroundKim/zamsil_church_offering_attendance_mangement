const state = {
    students: null,
    offeringType: [],
    payloadStudents: [],
    payloadOfferings: [],
    sendPost: false,
    
}

const mutations = {
    setStudents(state, newStudents) {
        state.students = newStudents
    },

    setOfferingType(state, newOfferingType) {
        state.offeringType = newOfferingType
    },
    
    pushStudent(state, student) {
        state.payloadStudents.push(student)
    },

    pushOffering(state, newOffering) {
        state.payloadOfferings.push(newOffering)
    },
}

const actions = {

}

const getters = {
    getStudents (state) {
        return state.students
    },

    getOfferingPayload(state) {
        return state.payloadOfferings
    },

    getOfferingType(state) {
        return state.offeringType
    }

}

export default {
    state,
    mutations,
    actions,
    getters
}