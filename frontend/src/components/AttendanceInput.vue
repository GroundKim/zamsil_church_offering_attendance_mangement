<template>
  <v-container>
    <form @submit.prevent="sendPost">
      <v-row>
        <v-col>
          <v-text-field
            v-model="createdBy"
            label="Who are you ?"
            width="20px"
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
          <v-card height="400">
            <h2 class="pa-5">Class {{ data.classId }}</h2>
            <h3>{{ data.teacherName }}</h3>

            <v-row>
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
          </v-card>
        </v-col>
      </v-row>
      <v-container class="text-center ma-10">
        <v-btn type="submit" width="700">제출</v-btn>
      </v-container>
    </form>

    <span>{{ attendedStudents }}</span>
    <br>
    <span>{{ department }}</span>
    <br>
    <span> {{ value }} </span>
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
      date: moment().format(),
      value: null,
    };
  },
  methods: {
    sendPost() {
      let payload = [];
      this.attendedStudents.forEach((element) => {
        let data = {
          studentId: element,
          date: this.date,
          createdBy: this.createdBy
        };
        payload.push(data);
      });
      const headers = {
      'Content-Type': 'application/json',
      }

      axios
        .post(
          "http://localhost:8080/Youth/attendances", JSON.stringify(payload), {headers: headers}
        )
        .then(res => {
          console.log(res.data)
        })
        .catch(err => {
          alert(err.message + ' 등록중 오류 발생 관리자에게 문의하십시오')
        })
    },
  },

  watch: {
  department : function () {
    let getURL = `http://localhost:8080/Youth/students?department_id=${this.department + 1}`
    axios
      .get(getURL)
      .then((response) => {
        this.attendanceData= response.data
        })
      .catch(err => {
        alert(err.message + " 출석부를 불러오는 도중 오류가 발생했습니다 관리자에게 문의하십시오")
      })
    }
  },
};
</script>
