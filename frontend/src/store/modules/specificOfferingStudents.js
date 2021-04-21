const state = {
    students: null,
    offeringType: [],
    payloadStudents: [],
    payloadOfferings: [],
    sendPost: false,
    offeredAt: null,
    createdBy: null,
    departmentId: null,
}

const mutations = {
    setStudents(state, newStudents) {
        state.students = newStudents
    },

    setOfferingType(state, newOfferingType) {
        state.offeringType = newOfferingType
    },

    setCreatedBy(state, createdBy) {
        state.createdBy = createdBy
    },

    setOfferedAt(state, offeredAt) {
        state.offeredAt = offeredAt
    },

    setDepartmentId(state, departmentId) {
        state.departmentId = departmentId
    },
    
    pushStudent(state, student) {
        state.payloadStudents.push(student)
    },

    pushOffering(state, newOffering) {
        state.payloadOfferings.push(newOffering)
    },

    deleteSpecificOffering(state) {
        state.payloadOfferings = []
    }
}

const actions = {

}

const getters = {
    getStudents (state) {
        return state.students
    },

    getCreatedBy (state) {
        return state.createdBy
    },

    getOfferingPayload(state) {
        return state.payloadOfferings
    },

    getOfferingType(state) {
        return state.offeringType
    },
    
    getDepartmentId(state) {
        return state.departmentId
    },

    getOfferedAt(state) {
        return state.offeredAt
    }
}

export default {
    state,
    mutations,
    actions,
    getters
}