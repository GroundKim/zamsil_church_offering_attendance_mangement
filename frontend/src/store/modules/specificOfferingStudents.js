const state = {
    students: null,
    teachers: null,
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

    //
    pushOfferingPayloads (state, offering) {
        state.offeringPayloads.push(offering)
    },
    
    setTeachers (state, teachers) {
        state.teachers = teachers
    },
    
    updateOfferingPayload (state, offeringPayload) {
        // state.offeringPayloads[offering.id].studentId = offering.studentId
        // state.offeringPayloads[offering.id].offeringTypeId = offering.offeringTypeId
        // state.offeringPayloads[offering.id].offeringCost = offering.offeringCost
        const index = state.offeringPayloads.findIndex(o => o.offeringId == offeringPayload.offeringId)
        state.offeringPayloads[index].studentId = offeringPayload.studentId
        state.offeringPayloads[index].offeringTypeId = offeringPayload.offeringTypeId
        state.offeringPayloads[index].offeringCost = offeringPayload.offeringCost
        state.offeringPayloads[index].departmentId = offeringPayload.departmentId
        state.offeringPayloads[index].offeredAt = offeringPayload.offeredAt
        state.offeringPayloads[index].createdBy = offeringPayload.createdBy
    },

    deleteOfferingPayload (state, id) {
        const index = state.offeringPayloads.findIndex(o => o.offeringId == id)
        state.offeringPayloads.splice(index, 1)

    },
    
    setCreatedBy(state, createdBy) {
        state.createdBy = createdBy
    },
    
    pushStudent(state, student) {
        state.payloadStudents.push(student)
    },

    deleteSpecificOffering(state) {
        state.payloadOfferings = []
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