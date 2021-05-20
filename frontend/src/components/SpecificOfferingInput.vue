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
  props: ['offeringId', 'postTrigger'],
  data: () => ({
    offeringType: [],
    offeringTypeValue: null,
    offerorId: null,
    offeringCost: null,
    students: null,
    departmentId: null,
    offeringPayload: {},

    rules: {
      offering: value => {
        const pattern = /^[0-9]+$/
        return pattern.test(value)
      }  
    }
  }),

  computed: {
    getDepartmentId() {
      return this.$store.getters.getDepartmentId
    },

    getOfferedAt() {
      return this.$store.getters.getOfferedAt
    }
  },

  methods: {
    deleteOffering: function (offeringId) {
      this.$emit('delete', offeringId)

    }
  },

  watch: {
    getDepartmentId: function () {
      let departmentId = this.getDepartmentId
      // it doesn't look so good
      if (departmentId == 1) this.students = this.$store.getters.getDepartmentOneStudents
      if (departmentId == 2) this.students = this.$store.getters.getDepartmentTwoStudents
      
    },

    postTrigger: function () {
      this.offeringPayload = {
        studentId: null,
        offeringTypeId: null,
        offeringCost: null,
        offeredAt: this.$store.getters.getOfferedAt,
        createdAt: moment().format(),
        createdBy: this.$store.getters.getCreatedBy,
      }
      this.$store.commit('pushOfferingPayload', this.offeringPayload)
    },

    getOfferedAt: function () {
      
    },

    offeringTypeValue: function() {
    },

    offerorId: function() {
    },

    offeringCost: function() {
    }


  },

  created() {
    this.offeringType = this.$store.getters.getOfferingType
    this.departmentId = this.getDepartmentId
    

    // it doesn't look so good
    if (this.departmentId == 1) this.students = this.$store.getters.getDepartmentOneStudents
    if (this.departmentId == 2) this.students = this.$store.getters.getDepartmentTwoStudents
    

  }
  
}
</script>