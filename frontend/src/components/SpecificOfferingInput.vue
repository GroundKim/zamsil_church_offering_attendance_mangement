<template>
  <v-row>
    <v-col class="py-0">
      <v-autocomplete 
        v-model="offeringTypeValue"
        label="종류"
        :items="offeringType"
        item-text="offeringName"
        required
        return-object
      ></v-autocomplete>
    </v-col>

    <v-col class="py-0">
      <v-autocomplete
        v-model="offerorValue"
        label="이름"
        :items="students"
        item-text="name"
        required
        return-object
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
  </v-row>
</template>

<script>
import moment from 'moment'

export default {
  name: 'SpecificOfferingInput',
  data: () => ({
    offeringType: [], //... backend 서버에서 get offering type
    offeringTypeValue: null,
    offerorValue: null,
    offeringCost: null,
    students: null,

    rules: {
      offering: value => {
        const pattern = /^[0-9]+$/
        return pattern.test(value)
      }  
    }
  }),
  
  props:['offerPostTrigger'],
  methods: {
    watchSendPost() {
      const {sendPost} = this
      return sendPost
    },

    makePayload() {
      let offeringPayload = {
        offeringTypeId: this.offeringTypeValue['offeringTypeId'],
        studentId: this.offerorValue['studentId'],
        offeringCost: parseInt(this.offeringCost),
        createdAt: moment().format(),
        createdBy: this.$store.getters.getCreatedBy,
        offeredAt: this.$store.getters.getOfferedAt,
        departmentId: this.$store.getters.getDepartmentId,
      }
      this.$store.commit('pushOffering', offeringPayload)
    }
  },

  computed: {
    // cannot use with method, not a computed because update cycle is not proccessed in methods
    getStudents() {
      return this.$store.getters.getStudents
    },
  },

  watch: {
    getStudents() {
      this.students = this.$store.getters.getStudents
    },

    offerPostTrigger: function() {
      this.makePayload()
    },
  },

  created() {
    this.students = this.$store.getters.getStudents
    this.offeringType = this.$store.getters.getOfferingType
  },
}
</script>