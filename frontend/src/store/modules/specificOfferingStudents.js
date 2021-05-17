const state = {
    students: null,
    offeringType: [],
    payloadStudents: [],
    payloadOfferings: [],
    sendPost: false,
    offeredAt: null,
    createdBy: null,
    departmentId: null,
    isFinishedMakingPayload: false,
    postTrigger: 0,
    offeringError: [],
    
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

    pushOffering (state, newOffering) {
        state.payloadOfferings.push(newOffering)
    },
    
    deleteSpecificOffering(state) {
        state.payloadOfferings = []
    },
    
    setFinishMakingPayload(state, bool) {
        state.isFinishedMakingPayload = bool
    },

    triggerPost(state) {
        state.postTrigger++
    },

    addOfferingError(state, error) {
        state.offeringError.push(error)
    }
}

const actions = {
    pushOffering: function (context, newOffering) {
        return context.commit('pushOffering', newOffering)
    },

    setFinishMakingPayload: function (context) {
        return context.commit('setFinishMakingPayload')
    },

    triggerPost: function (context) {
        return context.commit('triggerPost')
    }

    
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
    },

    getIsFinishedMakingPayload(state) {
        return state.isFinishedMakingPayload
    },

    getPostTrigger(state) {
        return state.postTrigger
    }

}

export default {
    state,
    mutations,
    actions,
    getters
}