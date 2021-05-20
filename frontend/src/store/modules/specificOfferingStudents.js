const state = {
    students: null,
    offeringType: [],
    offeringPayloads: [],
    sendPost: false,
    offeredAt: null,
    createdBy: null,
    departmentId: null,
    departmentOneStudents: [],
    departmentTwoStudents: [],

    offeringId: 0,
    deleteOfferingId: null,
}

const mutations = {
    //
    setStudents (state, newStudents) {
        state.students = newStudents
    },

    //
    setOfferingId (state, offeringId) {
        state.offeringId = offeringId
    },

    //
    setOfferingType (state, offeringType) {
        state.offeringType = offeringType
    },
    
    //
    setDepartmentId (state, departmentId) {
        state.departmentId = departmentId
    },

    //
    setDeleteOfferingId (state, offeringId) {
        state.deleteOfferingId = offeringId
    },

    //
    pushDepartmentOneStudent (state, student) {
        state.departmentOneStudents.push(student)
    },
    
    //
    pushDepartmentTwoStudent (state, student) {
        state.departmentTwoStudents.push(student)
    },
    
    //
    setOfferedAt (state, offeredAt) {
        state.offeredAt = offeredAt
    },

    updateOffering (state, newOffering) {
        state.offeringPayloads[newOffering.id].studentId = newOffering.studentId
        state.offeringPayloads[newOffering.id].offeringTypeId = newOffering.offeringTypeId
        state.offeringPayloads[newOffering.id].offeringCost = newOffering.offeringCost
    },
    
    setCreatedBy(state, createdBy) {
        state.createdBy = createdBy
    },
    
    pushStudent(state, student) {
        state.payloadStudents.push(student)
    },

    //
    pushOfferingPayload (state, offeringPayload) {
        state.offeringPayloads.push(offeringPayload)
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

    addOffering(state) {
        state.countOfferings++
    },

    addOfferingError(state, error) {
        state.offeringError.push(error)
    },

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

    //
    getOfferingType (state) {
        return state.offeringType
    },

    //
    getOfferedAt (state) {
        return state.offeredAt
    },
    
    //
    getDepartmentId (state) {
        return state.departmentId
    },

    // 
    getOfferingId (state) {
        return state.offeringId
    },

    //
    getDeleteOfferingId (state) {
        return state.deleteOfferingId
    },

    //
    getDepartmentOneStudents (state) {
        return state.departmentOneStudents
    },

    //
    getDepartmentTwoStudents (state) {
        return state.departmentTwoStudents
    },

    //
    getOfferingPayloads (state) {
        return state.offeringPayloads
    }

}

export default {
    state,
    mutations,
    actions,
    getters
}