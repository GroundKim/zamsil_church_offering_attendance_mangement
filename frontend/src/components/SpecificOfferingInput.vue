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
import axios from 'axios'

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
  
  methods: {
    makePayload: async function() {
      let offeringPayload = {
        studentId: this.offerorValue['studentId'],
        offeringTypeId: this.offeringTypeValue['offeringTypeId'],
        offeringCost: parseInt(this.offeringCost),
        departmentId: parseInt(await this.$store.getters.getDepartmentId),
        offeredAt: await this.$store.getters.getOfferedAt + moment().format().substr(10),
        createdAt: moment().format(),
        createdBy: await this.$store.getters.getCreatedBy,
      }

      return offeringPayload
      // await this.$store.dispatch('pushOffering', offeringPayload)
      // console.log(JSON.stringify(this.$store.getters.getOfferingPayload))
      // console.log("push finished")
    },

    sendPost: async function() {
      let payload = []
      payload.push(await this.makePayload())
      console.log(JSON.stringify(payload))
      const headers = {
        "Content-Type": "application/json",
      }
      axios.post(
        `${this.$serverAddress}/Youth/offering`,
        JSON.stringify(payload),
        {withCredentials: true, headers: headers}
      )
      .then()
      .catch((err) => {
        this.$store.commit('addOfferingError', err)
      })


    }
  },

  computed: {
    // cannot use with method, not a computed because update cycle is not proccessed in methods
    getStudents() {
      return this.$store.getters.getStudents
    },

    getPostTrigger() {
      return this.$store.getters.getPostTrigger
    }
  },

  watch: {
    getStudents() {
      this.students = this.$store.getters.getStudents
    },

    getPostTrigger: async function() {
      this.sendPost()
    }
  },

  created() {
    this.students = this.$store.getters.getStudents
    this.offeringType = this.$store.getters.getOfferingType
  },
}
</script>