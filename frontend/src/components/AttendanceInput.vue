<template>
  <div>
    <v-container fill-height>
      <v-layout align-center justify-center>
        <v-flex xs12 sm8 md4>
          <form @submit.prevent="sendPost">
            <v-card class="mt-10" elevation="0" shaped outlined>
              <v-card-text>
                <v-radio-group v-model="department" row>
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
                  required
                ></v-text-field>
              </v-card-text>

              <v-card-actions>
                <v-btn
                  color="primary"
                  block
                  rounded
                  outlined
                  x-large
                  @click="absentDialog=true; getAbsentStudents(); "
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

    <v-container>
      <v-row>
        <v-col
          v-for="data in currentStudents"
          v-bind:key="data.class.classId"
          cols="4"
        >
          <v-card height="100%" class="d-flex flex-column">
            <v-card-title>Class {{ data.class.name }}</v-card-title>
            <v-card-subtitle>{{ getTeacherNames(data.teachers) }}</v-card-subtitle>
            <v-card-text>
              <v-row class="pa-2">
                <v-col
                  v-for="student in data.students"
                  v-bind:key="student.studentId"
                  cols="4"
                >
                  <v-checkbox
                    v-model="attendedStudents"
                    :value="student.studentId"
                    :label="`${student.name}`"
                  ></v-checkbox>
                </v-col>
              </v-row>
            </v-card-text>
            <v-spacer></v-spacer>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
    <v-dialog
      v-model="absentDialog"
      transition="dialog-transition"
      max-width="700"
    >
      <v-card class="justify-center">
        <v-card-title primary-title class="justify-center">
          <div>
          
            <h3 class="headline ma-5">{{ department }}부 결석 처리</h3>
          </div>
        </v-card-title>
        <v-card-text
          v-for="(absentStudent, index) in absentStudents" 
          :key=index
        >
          <v-card>
            <v-card-title primary-title>
              {{ absentStudent.class.name }} 반
            </v-card-title>
            <div
              v-for="(student) in absentStudent.students"
              :key=student.studentId
              class="pa-1 ma-5"
            >
  
            {{ student.name }}
            
            </div>
          </v-card>
        </v-card-text>
        
        <v-card-actions>
          <v-btn 
            color="primary"
            rounded
            outlined
            x-large
            @click="closeAbsentDialog()"
          >취소</v-btn>
          <v-btn
            color="primary"
            rounded
            outlined
            x-large
            @click="sendPost()"
          >제출</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import axios from "axios"
import moment from "moment"

export default {
  data() {
    return {
      attendanceData: null,
      createdBy: null,
      department: null,
      departmentsLabel: ["1부", "2부"],
      menu: false,
      date: moment().format("yyyy-MM-DD"),
      value: null,
      selectAbsentReason: "일반결석",
      absentReason: ["일반결석", "특별결석"],

      classes: [],
      absentDialog: false,
      attendedStudents: [],
      members: [],
      departmentOneStudents: [],
      departmentTwoStudents: [],
      currentStudents: [],
      absentStudents: [],
    }
  },
  methods: {
    getTeacherNames (teachers) {
      let teacherNames = []
      teachers.forEach(t => {
        teacherNames.push(t.name)
      })
      return teacherNames.join(',')
    },

    getAbsentStudents () {
      let currentStudents = JSON.parse(JSON.stringify(this.currentStudents))

      currentStudents.forEach(cs => {
        this.attendedStudents.forEach(attendedStudentId => {
          for (let i = 0; i < cs.students.length; i++) {
            const student = cs.students[i]
            if (student.studentId === attendedStudentId) {
              cs.students.splice(i, 1)
              break
            }
          }
        })
      })
      this.absentStudents = currentStudents
    },

    closeAbsentDialog() {
      this.absentDialog = false
      this.absentStudentsByClasses = []
    },

    sendPost() {
      let payload = []

      // handling attended students
      this.date += moment().format().substr(10)
      this.attendedStudents.forEach((element) => {
        let data = {
          studentId: element,
          attendedAt: this.date,
          createdBy: this.createdBy,
        }
        payload.push(data)
      })

      // handling absent students
      const headers = {
        "Content-Type": "application/json",
      }

      axios
        .post(
          `${this.$serverAddress}/Youth/attendances`,
          JSON.stringify(payload),
          { withCredentials: true, headers: headers }
        )
        .then(() => {
          alert("등록 완료!")
          location.reload()
        })
        .catch((err) => {
          this.alertError(err)
        })
    },
  },

  watch: {
    department: function () {
      if (this.department == '1') {
        this.currentStudents = this.departmentOneStudents
      }
      if (this.department == '2') {
        this.currentStudents = this.departmentTwoStudents
      }
    },
  },

  created: async function() {
    await axios
      .get(`${this.$serverAddress}/Youth/members`, { withCredentials: true })
      .then((res) => {
        this.members = res.data
      })
      .catch((err) => {
        this.alertError(err)
      })

    // split students up with department
    await this.members.forEach(member => {
      if (member.class.department.name == 1) {
        this.departmentOneStudents.push(member)
      }
      
      if (member.class.department.name == 2) {
        this.departmentTwoStudents.push(member)
      }
    })

    this.members.forEach(member => {
      this.classes.push(member.class)
    })
    this.$store.commit("changeHeaderName", "출석부 기입")
    this.department = '1'
  },
}
</script>
