<template>
  <v-container>
    <form @submit.prevent="sendPost">
      <v-row>
        <v-col>
          <v-text-field
            v-model="createdBy"
            label="Who are you ?"
            required
          ></v-text-field>
        </v-col>

        <v-col>
          <v-card
            flat
            color="transparent"
          >
            <v-card-text>
              <v-slider
                v-model="department"
                :tick-labels="departmentsLabel"
                :max="1"
                step="1"
                ticks="always"
                tick-size="4"
              ></v-slider>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>

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
            class="mb-15"
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


      <v-row>
        <v-col
          v-for="data in attendanceData"
          v-bind:key="data.classId"
          cols="4"
        >

          <v-card height="100%" class="d-flex flex-column">
            <h2 class="pa-3">Class {{ data.className }}</h2>
            <h3 class="pa-3">{{ data.teacherName.join(', ') }}</h3>

            <v-row class="pa-2">
              <v-col
                v-for="student in data.studentsIdandName"
                v-bind:key="student.studentId"
                cols="5"
              >
              <v-checkbox
                v-model="attendedStudents"
                :value="student.studentId"
                :label="`${student.studentName}`"
              ></v-checkbox>
              </v-col>
            </v-row>
            <v-spacer></v-spacer>
            
          </v-card>
        </v-col>
        
      </v-row>
      
      <div class="text-center ma-15">
        <v-btn
          type="submit"
          style="width: 50%"
          class="white--text"
          color="indigo"
          elevation="4"
          x-large
        ><h3>제출</h3></v-btn>
      </div>
    </form>

    <span>{{ attendedStudents }}</span>
    <br>
    <span>{{ department }}</span>
    <br>
    <span> {{ value }} </span>
    <br>
    <span> {{ test }} </span>
    
  </v-container>
</template>

<script>
import axios from "axios";
import moment from "moment";
export default {
  data() {
    return {
      attendanceData: null,
      attendedStudents: [],
      createdBy: null,
      department: null,
      departmentsLabel: ['1부', '2부'],
      menu: false,
      date: moment().format('yyyy-MM-DD'),
      value: null,
      selectAbsentReason: "일반결석",
      absentReason: ["일반결석", "특별결석"],
      absentStudents: [],
      
      test: process.env.VUE_APP_SERVERADDRESS
    };
  },
  methods: {
    sendPost() {
      let payload = [];
      this.date += moment().format().substr(10, )
      this.attendedStudents.forEach((element) => {
        let data = {
          studentId: element,
          attendedAt: this.date,
          createdBy: this.createdBy
        };
        payload.push(data);
      });
      const headers = {
      'Content-Type': 'application/json',
      }

      axios
        .post(
          `${this.$serverAddress}/Youth/attendances`, JSON.stringify(payload), {headers: headers}, {withCredentials: true} 
        )
        .then(res => {
          console.log(res.data)
          alert("등록 완료!")
          location.reload()
        })
        .catch(err => {
          alert(err.message + ' 등록중 오류 발생 관리자에게 문의하십시오')
        })
    },
  },

  watch: {
    department: function () {
      let getURL = `${this.$serverAddress}/Youth/attendances?department_id=${this.department + 1}`
      axios
        .get(getURL, {withCredentials: true})
        .then((response) => {
          this.attendanceData= response.data
          })
        .catch(err => {
          let errStatusCode = err.response.status
          if (errStatusCode === 404) {
            alert(err.message + ": 출석부를 불러오는 도중 오류가 발생했습니다 관리자에게 문의하십시오")
          }

          if (errStatusCode === 403) {
            alert('로그인을 해주세요')
            this.$router.push('/login')
          }

        })
      },
      
    attendedStudents: function() {      
      console.log(this.absentStudents)
      console.log(this.attendedStudents)
    }
  },



  created () {
    this.$store.commit('changeHeaderName', '출석부 기입')
    this.absentStudents = this.attendedStudents
  }
};
</script>
