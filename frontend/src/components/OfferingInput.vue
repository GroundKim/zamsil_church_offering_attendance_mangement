<template>
  <v-container fill-height>
    <v-layout align-center justify-center>
      <v-flex xs12 sm8 md8>
        <form @submit.prevent="sendPost">
          <v-card elevation="0" shaped outlined>
            <v-card-text>
              <v-radio-group v-model="departmentId" row>
                <v-radio label="1부" value="1"></v-radio>
                <v-radio label="2부" value="2"></v-radio>
              </v-radio-group>

              <v-menu
                ref="menu"
                v-model="menu"
                :close-on-content-click="false"
                :return-value.sync="date"
                transition="scale-transition"
                offset-y
                min-width="auto"
              >
                <template v-slot:activator="{ on, attrs }">
                  <v-text-field
                    v-model="date"
                    label="날짜"
                    prepend-icon="mdi-calendar"
                    readonly
                    v-bind="attrs"
                    v-on="on"
                  ></v-text-field>
                </template>

                <v-date-picker v-model="date" no-title scrollable>
                  <v-spacer></v-spacer>
                  <v-btn text color="primary" @click="menu = false">
                    Cancel
                  </v-btn>
                  <v-btn text color="primary" @click="$refs.menu.save(date)">
                    OK
                  </v-btn>
                </v-date-picker>
              </v-menu>

              <v-text-field
                v-model="createdBy"
                label="작성자"
                :ruels="[rules.required]"
                required
              ></v-text-field>

              <v-text-field
                id="offeringCost"
                v-model="weekOfferingCost"
                label="주일헌금"
                prefix="₩"
                :rules="[rules.offering, rules.required]"
                required
              >
              </v-text-field>

              <div class="d-flex justify-center">
                <v-btn
                  class="mx-2"
                  fab
                  outlined
                  small
                  color="primary"
                  @click="addOffering()"
                >
                  <v-icon dark> mdi-pencil-outline </v-icon>
                </v-btn>
              </div>

              <v-divider class="mt-5 mb-5"></v-divider>

              <component
                v-for="(offering, index) in offerings"
                :is="offering.type"
                :key="index"
                :offeringId="offering.offeringId"
                @delete="deleteOffering"
              ></component>
            </v-card-text>

            <v-card-actions>
              <v-btn
                type="submit"
                color="primary"
                block
                rounded
                outlined
                x-large
              >
                제출
              </v-btn>
            </v-card-actions>
          </v-card>
        </form>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import SpecificOfferingInput from './SpecificOfferingInput'

import moment from "moment"
import axios from'axios'
export default {
  name: "OfferingInput",
  components: {
    SpecificOfferingInput,
  },

  data: () => ({
    date: moment().format("yyy-MM-DD"),
    menu: false,
    modal: false,
    menu2: false,
    departmentId: null,
    departmentsLabel: ["1부", "2부"],
    weekOfferingCost: 0,
    createdBy: null,
    offerings: [],
    offeringCount: 0,
    students: [],
    postTrigger: 0,

    rules: {
      required: (value) => !!value || "Required.",
      offering: (value) => {
        const pattern = /^(\d+|\d{1,3}(,\d{3})*)(\.\d+)?$/
        return pattern.test(value)
      },
    },
  }),

  methods: {
    // sendPost
    sendPost: async function() {
      let offeringPayload = []
      let date = this.date + moment().format().substr(10)
      let weekOfferingPayload = {
        studentId: null,
        teacherId: null,
        offeringTypeId: 1,
        offeringCost: parseInt(this.weekOfferingCost.replace(',','')),
        departmentId: parseInt(this.departmentId),
        offeredAt: date,
        createdAt: moment().format(),
        createdBy: this.createdBy,
      }
      
      // push week offering
      offeringPayload.push(weekOfferingPayload)

      // push specific offerings
      offeringPayload = offeringPayload.concat(this.$store.getters.getOfferingPayloads)

      // set header for JSON post
      const headers = {
       'content-type': 'application/json' 
      }

      // axios post
      await axios
        .post(
          `${this.$serverAddress}/Youth/offering`,
          JSON.stringify(offeringPayload),
          { withCredentials: true, headers: headers}
        )
        .then(() => {
          alert('등록완료')
          location.reload()   
        })
        .catch((err) => {
          this.errorHandler(err)
        })
    },

    // add offering component
    addOffering: function () {
      this.offeringCount++
      this.offerings.push({
        'type': SpecificOfferingInput,
        offeringId: this.offeringCount
      })

    },

    deleteOffering(id) {
      let index = this.offerings.findIndex(o => o.offeringId == id)
      this.$store.commit('deleteOfferingPayload', id)
      this.offerings.splice(index, 1)
    },

    // utilities
    formatCurrency(input) {
      return input.replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ",")
    }
  },

  computed: {
    getDeleteOfferingId: function () {
      return this.$store.getters.getDeleteOfferingId
    },
  },

  watch: {
    departmentId: function () {
      this.$store.commit('setDepartmentId', this.departmentId)
    },

    date: function () {
      this.$store.commit('setOfferedAt', this.date)
    },

    createdBy: function () {
      this.$store.commit('setCreatedBy', this.createdBy)
    },

    // weekOfferingCost: function(newValue) {
    //   const inputElem = document.getElementById('weekOfferingCost')
    //   const result = this.formatCurrency(newValue)
    //   const prevSelection = inputElem.selectionStart

    //   this.$nextTick(() => {
    //     this.weekOfferingCost = result

    //     let selectionStart = prevSelection

    //     if (inputElem.value.length == inputElem.selectionStart) {
    //       selectionStart += 1
    //     }
        
    //     inputElem.focus()
    //     inputElem.setSelectionRange(selectionStart, selectionStart)
    //   })          
    // }
  },
  
  created: async function () {
    // get offering Type
    await axios
      .get(`${this.$serverAddress}/Youth/offering/types`, { withCredentials: true})
      .then((res) => {
        this.$store.commit('setOfferingType', res.data)
      })
      .catch((err) => {
        this.errorHandler(err)
      })
    
    // get teachers
    await axios
      .get(`${this.$serverAddress}/Youth/teachers`, { withCredentials: true})
      .then((res) => {
        this.$store.commit('setTeachers', res.data)
      })
      .catch((err) => {
        this.errorHandler(err)
      })

    // get all students
    await axios
      .get(`${this.$serverAddress}/Youth/students`, { withCredentials: true })
      .then((res) => {
        this.students = res.data
      })
      .catch((err) => {
        this.errorHandler(err)
      })

    //split students into department 1 and 2
    await this.students.forEach(student => {
      if (student.departmentId === 1) {
        this.$store.commit('pushDepartmentOneStudent', student)
      }

      if (student.departmentId === 2) {
        this.$store.commit('pushDepartmentTwoStudent', student)
      }
    })
    
    // store offeredAt in vuex with proper date formate for server
    this.$store.commit('setOfferedAt', this.date)
    this.departmentId = '1'
    this.offerings.push({
      'type':SpecificOfferingInput,
      offeringId: this.offeringCount
    })

    // offering cost delim
    const weekOfferingCost = document.getElementById('offeringCost')
    weekOfferingCost.addEventListener('blur', () => {
      this.weekOfferingCost = this.formatCurrency(this.weekOfferingCost)
    })
    
    // weekOfferingCost.addEventListener('keydown', (e) => {
    //   if (e.key === 'Backspace' || e.key === 'Delete') {
    //     e.preventDefault()

    //     let value = weekOfferingCost.value.replace(/,/g, '').split('')
    //     let selectionStart = weekOfferingCost.selectionStart
    //     let selectionEnd = weekOfferingCost.selectionEnd

    //     if (selectionStart == selectionEnd) {
    //       if (value.length >= 3) {
    //         selectionStart -= parseInt(value.length / 3)
    //       }
  
    //       if (value.length % 3 !== 0) {
    //         selectionStart -= 1
    //       }
  
    //       if (weekOfferingCost.value[weekOfferingCost.selectionStart] === ',') {
    //         selectionStart += 1
    //       }
  
    //       if (e.key === 'Delete') {
    //         selectionStart += 1
    //       } 
  
    //       if (selectionStart > -1) {
    //         value.splice(selectionStart, 1)
    //       }
  
    //       weekOfferingCost.value = this.formatCurrency(value.join(''))
  
    //       if (weekOfferingCost.value.replace(/,/g, '').length >= 3) {
    //         selectionStart += parseInt(weekOfferingCost.value.replace(/,/g, '').length / 3)
    //       }
  
    //       weekOfferingCost.setSelectionRange(selectionStart, selectionStart)
    //     } else {
    //       value.splice(selectionStart, selectionEnd)
    //       weekOfferingCost.value = this.formatCurrency(value.join(''))
    //     }
    //   }
    // })
  }
}
</script>
