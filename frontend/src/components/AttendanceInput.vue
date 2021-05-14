<template>
  <div>
    <v-container fill-height>
			<v-layout align-center justify-center>
				<v-flex xs12 sm8 md4>
					<form @submit.prevent="sendPost">
            <v-card class="mt-10" elevation="0" shaped outlined>
              <v-card-text>
                <v-radio-group
                  v-model="department"
                  row
                >
                  <v-radio
                    label="1부"
                    value="1"
                    
                  ></v-radio>
                  <v-radio
                    label="2부"
                    value="2"
                  ></v-radio>
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
                  label="Who are you ?"
                  required
                ></v-text-field>
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
      <!--
      <span>{{ attendedStudents }}</span>
      <br>
      <span>{{ department }}</span>
      <br>
      <span> {{ value }} </span>
      <br>
      <span> {{ test }} </span>
      -->
    </v-container>

    <v-divider v-if="attendanceData"></v-divider>

    <v-container>
      <v-row>
          <v-col
            v-for="data in attendanceData"
            v-bind:key="data.classId"
          >
            <v-card height="100%" class="d-flex flex-column">
              <v-card-title>Class {{ data.className }}</v-card-title>
              <v-card-subtitle>{{ data.teacherName.join(', ') }}</v-card-subtitle>
              <v-card-text>
                <v-row class="pa-2">
                  <v-col
                    v-for="student in data.studentsIdandName"
                    v-bind:key="student.studentId"
                  >
                    <v-checkbox
                      v-model="attendedStudents"
                      :value="student.studentId"
                      :label="`${student.studentName}`"
                    ></v-checkbox>
                  </v-col>
                </v-row>
              </v-card-text>
              <v-spacer></v-spacer>
            </v-card>
          </v-col>
        </v-row>
    </v-container>
  </div>
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
      department: "1",
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
      let getURL = `${this.$serverAddress}/Youth/attendances?department_id=${this.department}`
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
