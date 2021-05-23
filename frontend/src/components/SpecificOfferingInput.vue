<template>
  <v-row>
    <v-col class="py-0">
      <v-autocomplete 
        v-model="offeringTypeValue"
        label="종류"
        :items="offeringType"
        required
        item-text="offeringName"
        item-value="offeringTypeId"
      ></v-autocomplete>
    </v-col>

    <v-col class="py-0">
      <v-autocomplete
        v-model="offerorId"
        label="이름"
        :items="students"
        required
        item-text="name"
        item-value="studentId"
      ></v-autocomplete>
    </v-col>

    <v-col class="py-0">
      <v-text-field
        v-model="offeringCost"
        label ="금액"
        prefix="₩"
        required
        :rules="[rules.offering]"
      ></v-text-field>
    </v-col>
    <div class="ma-auto">
      <v-btn
        fab
        outlined
        small
        color="red"
        @click="deleteOffering(offeringId)"
      >
        <v-icon dark> mdi-trash-can-outline </v-icon>
      </v-btn>
    </div>
    {{offeringId}}
  </v-row>
</template>

<script>
import moment from 'moment'

export default {
  name: 'SpecificOfferingInput',
  props: ['offeringId'],
  data: () => ({
    offeringType: [],
    offeringTypeValue: null,
    offerorId: null,
    offeringCost: null,
    students: [],
    departmentId: null,
    offeringPayload: {
      offeringId: null,
      studentId: null,
      offeringTypeId: null,
      offeringCost: null,
      departmentId: null,
      offeredAt: null,
      createdAt: null,
      createdBy: null,
    },

    rules: {
      offering: value => {
        const pattern = /^(\d+|\d{1,3}(,\d{3})*)(\.\d+)?$/
        return pattern.test(value)
      }  
    }
  }),

  computed: {
    getDepartmentId() {
      return this.$store.getters.getDepartmentId
    },

    getCreatedBy() {
      return this.$store.getters.getCreatedBy
    },

    getOfferedAt() {
      return this.$store.getters.getOfferedAt
    },

  },

  methods: {
    deleteOffering: function (offeringId) {
      this.$emit('delete', offeringId)
    }
  },

  watch: {
    getDepartmentId: function () {
      let departmentId = parseInt(this.getDepartmentId)
      // it doesn't look so good
      if (departmentId == 1) this.students = this.$store.getters.getDepartmentOneStudents
      if (departmentId == 2) this.students = this.$store.getters.getDepartmentTwoStudents
      
      this.offeringPayload.departmentId = departmentId
    },

    // repetition is not good
    getCreatedBy: function () {
      this.offeringPayload.createdBy = this.getCreatedBy
      this.$store.commit('updateOfferingPayload', this.offeringPayload)
    },

    getOfferedAt: function () {
      this.offeringPayload.offeredAt = this.getOfferedAt
      this.$store.commit('updateOfferingPayload', this.offeringPayload)
    },

    offeringTypeValue: function() {
      this.offeringPayload.offeringTypeId = this.offeringTypeValue
      this.$store.commit('updateOfferingPayload', this.offeringPayload)

    },

    offerorId: function() {
      this.offeringPayload.studentId = this.offerorId
      this.$store.commit('updateOfferingPayload', this.offeringPayload)
    },

    offeringCost: function(newValue) {
      const result = newValue.replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ",")
      this.$nextTick(() => this.offeringCost = result)
      this.offeringPayload.offeringCost = parseInt(this.offeringCost.replace(',', ''))
      this.$store.commit('updateOfferingPayload', this.offeringPayload)
    }
  },

  created() {
    this.offeringType = this.$store.getters.getOfferingType
    this.departmentId = this.getDepartmentId
    
    this.students.push({
      'studentId': null,
      'name': '무명'
    })

    // it doesn't look so good
    if (this.departmentId == 1) this.students = this.students.concat(this.$store.getters.getDepartmentOneStudents)
    if (this.departmentId == 2) this.students = this.students.concat(this.$store.getters.getDepartmentTwoStudents)

    this.offeringPayload.offeringId = this.offeringId
    this.offeringPayload.departmentId = parseInt(this.$store.getters.getDepartmentId)
    this.offeringPayload.offeredAt = this.$store.getters.getOfferedAt
    this.offeringPayload.createdAt = moment().format()
    this.offeringPayload.createdBy = this.$store.getters.getCreatedBy

    this.$store.commit('pushOfferingPayloads', this.offeringPayload)
  }
  
}
</script>