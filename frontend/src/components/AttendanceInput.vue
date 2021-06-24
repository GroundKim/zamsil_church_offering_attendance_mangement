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
    </v-container>

    <v-container>
      <v-row>
        <v-col
          v-for="data in currentStudents"
          v-bind:key="data.class.classId"
          cols="12"
          sm="12" md="4"
        >
          <v-card height="100%" class="d-flex flex-column">
            <v-card-title>Class {{ data.class.name }}</v-card-title>
            <v-card-subtitle>{{ getTeacherNames(data.teachers) }}</v-card-subtitle>
            <v-card-text>
              <div
                v-for="student in data.students"
                v-bind:key="student.studentId"
              >
                <v-checkbox
                  v-model="attendedStudents"
                  :value="student.studentId"
                  :label="`${student.name}`"
                ></v-checkbox>
              </div>
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
              <h3>{{ absentStudent.class.name }} 반</h3>
            </v-card-title>
              <v-row class="ma-5">
                <v-col><h4>이름</h4></v-col>
                <v-col><h4>결석종류</h4></v-col>
                <v-col><h4>결석사유</h4></v-col>
              </v-row>
              <v-row
                v-for="(student) in absentStudent.students"
                :key=student.studentId
                class="pa-0 ma-5"
              >
                <v-col>
                  {{student.name}}
                </v-col>

                <v-col>
                  <v-list-item-group
                    :mandatory=true
                    v-model="student.absenceTypeId"
                    color="primary"
                  >
                    <v-list-item
                      v-for="type in absenceTypes"
                      :key=type.absenceTypeId
                      :value="type.absenceTypeId"
                    >
                      {{type.name}}
                    </v-list-item>
                  </v-list-item-group>
                </v-col>

                <v-col>
                  <div
                    v-show="student.absenceTypeId !== 1"
                  >
                    <v-text-field
                      v-model="student.absenceReason"
                      label="사유"
                    ></v-text-field>
                  </div>
                </v-col>
              </v-row>
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

      absentDialog: false,
      attendedStudents: [],
      members: [],
      departmentOneStudents: [],
      departmentTwoStudents: [],
      currentStudents: [],
      absentStudents: [],
      absenceTypes: [],
      absentStudentDetails: [],
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
        for (let i = 0; i < cs.students.length; i++) {
          const student = cs.students[i];
          student.absenceTypeId = 1
          student.absenceReason = null
          
          for (let j = 0; j < this.attendedStudents.length; j++) {
            const as = this.attendedStudents[j];
            if (as.studentId === cs.studentId) {
              cs.students.splice(i, 1)
              break
            }
          }
        }
      })


      this.absentStudents = currentStudents
    },

    closeAbsentDialog() {
      this.absentDialog = false
      this.absentStudentsByClasses = []
    },

    sendPost: async function() {
      let payload = []
      let absentStudentspayload = []
      let postError = false

      // handling attended students
      let date = this.date + moment().format().substr(10)
      this.attendedStudents.forEach((element) => {
        let data = {
          studentId: element,
          attendedAt: date,
          createdBy: this.createdBy,
        }
        payload.push(data)
      })

      // handling absent students
      this.absentStudents.forEach((absentStudent) => {
        absentStudent.students.forEach(student => {
          if (student.absenceTypeId !== 1) {
            student.absentAt = date
            student.createdBy = this.createdBy
            absentStudentspayload.push(student)
          }
        })
      })

      // post 
      const headers = {
        "Content-Type": "application/json",
      }

      await axios
        .post(
          `${this.$serverAddress}/Youth/attendances`,
          JSON.stringify(payload),
          { withCredentials: true, headers: headers }
        )
        .then(() => {})
        .catch((err) => {
          postError = true
          this.errorHandler(err)
        })
      
      await axios
        .post(`${this.$serverAddress}/Youth/absence`, JSON.stringify(absentStudentspayload), { withCredentials: true, headers: headers})
        .then(() => {})
        .catch((err) => {
          postError = true
          this.errorHandler(err)
        })

      if (!postError) {
        alert('등록완료!') 
        location.reload()
      }

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
        this.errorHandler(err)
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

    await axios
      .get(`${this.$serverAddress}/Youth/absence/types`, { withCredentials: true })
      .then((res) => {
        this.absenceTypes = res.data
      })
      .catch((err) => {
        this.errorHandler(err)
      })


    this.department = '1'
  },
}
</script>
