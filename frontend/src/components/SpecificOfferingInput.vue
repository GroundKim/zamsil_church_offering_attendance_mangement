<template>
  <v-row>
    <v-col class="py-0">
      <v-autocomplete 
        v-model="offeringTypeValue"
        label="종류"
        :items="offeringType"
        required
        item-text="offeringTypeName"
        item-value="offeringTypeId"
      ></v-autocomplete>
    </v-col>

    <v-col class="py-0">
      <v-autocomplete
        v-model="offeror"
        label="이름"
        :items="offerors"
        required
        item-text="name"
        return-object
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

    <v-col 
      class="py-0"
      cols="4"
    >
      <v-text-field
        v-model="offeringNote"
        label ="헌금 사유"
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
    offeror: null,
    offeringCost: null,
    students: [],
    teachers: [],
    offerors: [],
    departmentId: null,

    offeringNote: null,

    offeringPayload: {
      offeringId: null,
      studentId: null,
      teacherId: null,
      offeringTypeId: null,
      offeringCost: null,
      departmentId: null,
      offeredAt: null,
      createdAt: null,
      createdBy: null,
      offeringNote: null,
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
    },
  },

  watch: {
    getDepartmentId: function () {
      let departmentId = parseInt(this.getDepartmentId)
      // it doesn't look so good
      if (departmentId == 1) this.offerors = this.$store.getters.getDepartmentOneStudents
      if (departmentId == 2) this.offerors = this.$store.getters.getDepartmentTwoStudents
      
      this.offerors = this.offerors.concat(this.$store.state.specificOfferingStudents.teachers)
      this.offeringPayload.departmentId = departmentId
    },

    // repetition is not good
    getCreatedBy: function () {
      this.offeringPayload.createdBy = this.getCreatedBy
      this.$store.commit('updateOfferingPayload', this.offeringPayload)
    },

    getOfferedAt: function () {
      let date = this.getOfferedAt + moment().format().substr(10)
      this.offeringPayload.offeredAt = date
      this.$store.commit('updateOfferingPayload', this.offeringPayload)
    },

    offeringTypeValue: function() {
      this.offeringPayload.offeringTypeId = this.offeringTypeValue
      this.$store.commit('updateOfferingPayload', this.offeringPayload)

    },

    offeror: {
      deep: true,
      handler() {
        if (this.offeror.studentId != null) {
          this.offeringPayload.studentId = this.offeror.studentId
          this.offeringPayload.teacherId = null
        }

        if (this.offeror.teacherId != null) {
          this.offeringPayload.studentId = null
          this.offeringPayload.teacherId = this.offeror.teacherId
        }
        this.$store.commit('updateOfferingPayload', this.offeringPayload)
      }
    },

    offeringCost: function(newValue) {
      const result = newValue.replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ",")
      this.$nextTick(() => this.offeringCost = result)
      this.offeringPayload.offeringCost = parseInt(this.offeringCost.replace(',', ''))
      this.$store.commit('updateOfferingPayload', this.offeringPayload)
    },

    offeringNote: function() {
      this.offeringPayload.offeringNote = this.offeringNote
      this.$store.commit('updateOfferingPayload', this.offeringPayload)
    }
  },

  created() {
    this.offeringType = this.$store.getters.getOfferingType
    this.departmentId = this.getDepartmentId
    this.offerors = this.offerors.concat(this.$store.state.specificOfferingStudents.teachers)

    this.students.push({
      'studentId': null,
      'name': '무명'
    })

    if (this.departmentId == 1) {
      this.offerors = this.offerors.concat(this.$store.getters.getDepartmentOneStudents)
    }

    if (this.departmentId == 2) {
      this.offerors = this.offerors.concat(this.$store.getters.getDepartmentTwoStudents)
    }

    this.offeringPayload.offeringId = this.offeringId
    this.offeringPayload.departmentId = parseInt(this.$store.getters.getDepartmentId)
    this.offeringPayload.offeredAt = this.$store.getters.getOfferedAt + moment().format().substring(10)
    this.offeringPayload.createdAt = moment().format()
    this.offeringPayload.createdBy = this.$store.getters.getCreatedBy

    this.$store.commit('pushOfferingPayloads', this.offeringPayload)
  }
  
}
</script>