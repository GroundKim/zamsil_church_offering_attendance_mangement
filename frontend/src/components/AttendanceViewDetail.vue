<template>
  <v-container>
    <h1 class="ml-5 mt-5">{{ date }}</h1>
    
    <v-simple-table>
      <template v-slot:default>
        <thead>
          <tr>
            <th width="10%">출석 총원</th>
            <th>1부</th>
            <th>2부</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>{{ getNumberOfAttendedStudentsByDepartmentName(1) + getNumberOfAttendedStudentsByDepartmentName(2) }} 명</td>
            <td>{{ getNumberOfAttendedStudentsByDepartmentName(1) }} 명</td>
            <td>{{ getNumberOfAttendedStudentsByDepartmentName(2) }} 명</td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>
    <v-row
      v-for="(classInfo, classIndex) in classesWithStudents"
      :key="classIndex"
    >
      <v-container>
        <h2 class="ml-5 my-5">{{ classInfo.class.department.name }}부 {{ classInfo.class.name }}반</h2>
        <v-col>
          <v-data-table
            :headers="studentHeaders()"
            :items="getAttendanceDiaryTableItems(classInfo)"
            class="elevation-1"
            loading="true"
            hide-default-footer
            :footer-props="{
						itemsPerPageOptions: [-1]
					}"
          >
            <!-- 출석 여부 in data table -->
            <template v-slot:[`item.attendanceType`]="{ item, index }">
                <v-edit-dialog>
                  <div>{{ item.attendanceType }}</div>
                  <template v-slot:input>
                    <v-list-item-group
                      v-model=item.attendanceType
                      :mandatory=true
                    >
                      <v-list-item value="출석" color="primary" @click="changeAttendance(classIndex, index, true)">출석</v-list-item>
                      <v-list-item value="결석" color="red" @click="changeAttendance(classIndex, index, false)">결석</v-list-item>
                    </v-list-item-group>
                  </template>
                </v-edit-dialog>
            </template>

            <!-- 결석 종류 in data table -->
            <template v-slot:[`item.absence.absenceTypeName`] = "{ item }" >
              <v-edit-dialog>
                <div>{{ item.absence.absenceTypeName }}</div>
                <template v-slot:input>
                  <v-list-item-group
                    v-model="item.absence.absenceTypeId"
                    :mandatory="true"
                  >
                    <v-list-item
                      v-for="absenceType in absenceTypes"
                      :key=absenceType.id
                      :value="absenceType.absenceTypeId"
                      @click="changeAbsenceTypeName(absenceType, item)"
                    >{{ absenceType.name }}</v-list-item>
                  </v-list-item-group>
                </template>
              </v-edit-dialog>
            </template>
            
            <!-- 결석 사유 in data table -->
            <template v-slot:[`item.absence.absenceReason`] = "{ item }">
              <v-edit-dialog
                @close="changeAbsenceReason(item)"
              >
                <div v-if="item.absence.absenceTypeId != 1">{{ item.absence.absenceReason }}</div>
                <template v-slot:input v-if="item.absence.absenceTypeId != 1">
                  <v-text-field
                    v-model="item.absence.absenceReason"
                    label="사유"
                  ></v-text-field>
                </template>
                <template v-slot:input v-else>
                  <div>
                    일반 결석은 결석 사유를 적을 수 없습니다. 다른 결석으로 바꿔주세요.
                  </div>
                </template>
              </v-edit-dialog>
            </template>
          </v-data-table> 
          </v-col>
      </v-container>
      </v-row>
      <!-- edit dialog -->
  </v-container>
</template>

<script>
import axios from 'axios'
import moment from 'moment'
export default {
	data() {
		return {
			date: null,
      numberOfAttendedStudentDepartmentOne: 0,
      numberOfAttendedStudentDepartmentTwo: 0,
      studentsWithAttendance: [],
			classesWithStudents: [], 
      attendanceDiaries: {
        class: null,
        students: null,
      },

      absenceTypes: null,
      attendanceClasses: [],
      attendanceSummary: [],

      absenceDiary: {
        studentId: null,
        isAttendned: '사유결석',
        absenceType: '일반결석',
        absenceReason: null,
      },

      attendanceClass: {
        attendanceDiaries: [],
      },

		}
	},

  methods: {
    getAttendanceDiaryTableItems(classInfo) {
      let index = this.classesWithStudents.findIndex(c => c.class.classId == classInfo.class.classId)
      return this.classesWithStudents[index].studentAttendances
    },

    getNumberOfAttendedStudentsByDepartmentName(deparmtentName) {
      let numberOfStudents = 0
      this.classesWithStudents.forEach(c => {
        let cDepartmentName = c.class.department.name
        if (cDepartmentName === deparmtentName) c.studentAttendances.forEach(sa => {
          if (sa.attendanceType === '출석') numberOfStudents++
        })
      })

      return numberOfStudents
    },

    changeAbsenceReason (student) {
      if (student.absence.absenceTypeId !== 1) {
        axios
          .patch(`${this.$serverAddress}/Youth/absence/${student.absence.absenceDiaryId}?reason=${student.absence.absenceReason}`, null, { withCredentials: true })
          .then(() => {})
          .catch((err) => {
            this.alertError(err)
          })
      }
    },

    studentHeaders () {
      return [
        {
          text: '학생 이름',
          align: 'start',
          value: 'name',
        },

        {
          text: '출석 여부',
          value: 'attendanceType',

        },

        {
          text: '결석 종류',
          value: 'absence.absenceTypeName',
        },

        {
          text: '결석 사유',
          value: 'absence.absenceReason'
        }
      ]
    },

    changeAbsenceTypeName (absenceType, item) {
      // firstly, do delete attendance method when the student is attended
      if (!item.absence.isAbsent) {
        this.deleteAttendance(item)
      }

      if (absenceType.absenceTypeId === 1) {
        // if type id 1, which means '일반 출석': 다른 타입 -> 일반
        // delete not 일반 absence diary
        this.deleteAbsenceDiary(item.absence.absenceDiaryId)

      } else {
        let payload = []
        item.absentAt = this.date + moment().format().substr(10)
        item.absenceTypeId = absenceType.absenceTypeId
        payload.push(item)
        const headers = {
          "content-type": "application/json"
        }
        // 일반 - > 다른 타입
        axios
          .post(`${this.$serverAddress}/Youth/absence`, JSON.stringify(payload), {withCredentials: true, headers: headers})
          .then(() => {
          })
          .catch((err) => {
            this.alertError(err)
          })
      }


      // axios
      //   .patch(`${this.$serverAddress}/Youth/absence/type`, , {})

      // axios then
      item.absence.absenceTypeId = absenceType.absenceTypeId
      item.absence.absenceTypeName = absenceType.name

    },

    changeAttendance: async function (classIndex, studentIndex, isAttended) {
      let student = this.classesWithStudents[classIndex].studentAttendances[studentIndex]
      let absenceDiaryId = student.absence.absenceDiaryId
      student.attendedAt = this.date + moment().format().substr(10)
      student.createdBy = null
      if (isAttended) {
        student.attendanceType = "출석"
        let payload = []
        payload.push(student)

        const headers = {
          "Content-Type": "application/json"
        }
        await axios
          .post(`${this.$serverAddress}/Youth/attendances`, JSON.stringify(payload), { withCredentials: true, headers: headers })
          .then(()=> {
            for(const key in student.absence) {
              student.absence[key] = null
            }
          })
          .catch((err) => {
            this.alertError(err)
          })
        
        if (absenceDiaryId !== null) {
          this.deleteAbsenceDiary(absenceDiaryId)
        }
      } 
      
      if (!isAttended){
        student.attendanceType = "결석"
        // delete attendance diary
        this.deleteAttendance(student)

      }
    },

    deleteAttendance (student) {
      axios
        .delete(`${this.$serverAddress}/Youth/attendance/${this.date}?student_id=${student.studentId}`, { withCredentials:true })
        .then(() => {
          student.attendanceType = '결석'
          if (student.absence.absenceTypeName == null){
            student.absence.absenceTypeName = '일반결석'
            student.absence.absenceTypeId = 1
          }
          student.absence.isAbsent = true
        })
        .catch((err) => {
          this.alertError(err)
        })
      },
    
    deleteAbsenceDiary: async function (absenceDiaryId) {
      let error = false
      await axios
      .delete(`${this.$serverAddress}/Youth/absence/${absenceDiaryId}`, {withCredentials: true})
      .then(() => {})
      .catch((err) => {
        error = true
        this.alertError(err)
      })
      return error 
    }
  },

	created: async function () {
		this.date = this.$route.query.date
    await axios
      .get(`${this.$serverAddress}/Youth/absence/types`, { withCredentials: true })
      .then((res) => {
        this.absenceTypes = res.data
      })
      .catch((err) => {
        this.alertError(err)
      })

    // get all classes
    await axios
      .get(`${this.$serverAddress}/Youth/classes`, { withCredentials: true })
      .then((res) => {
        let classes = res.data
        classes.forEach(c => {
          let classWithAttendanceDiaries = {
            class: c,
            studentAttendances: []
          }
        this.classesWithStudents.push(classWithAttendanceDiaries) 
        })
      })
      .catch((err) => {
        this.alertError(err)
      })

    await axios
      .get(`${this.$serverAddress}/Youth/attendance/view?date=${this.date}`, { withCredentials: true })
      .then((res) => {
        this.studentsWithAttendance = res.data
      })
      .catch((err) => {
        this.alertError(err)
      })

    // split up the students with class
    this.studentsWithAttendance.forEach(s => {
      let classIndex = this.classesWithStudents.findIndex(c => c.class.classId == s.student.classId)
      let studentAttendance = {
        studentId: s.student.studentId,
        classId: s.student.classId,
        name: s.student.name,
        attendanceType: null,
        absence: {
          absenceDiaryId: null,
          isAbsent: false,
          absenceTypeId: null,
          absenceTypeName: null,
          absenceReason:null,
        }
      }
      
      if (s.absence !== null) {
        studentAttendance.attendanceType = '결석'
        studentAttendance.absence.absenceDiaryId = s.absence.absenceDiaryId
        studentAttendance.absence.isAbsent = true
        studentAttendance.absence.absenceTypeId = s.absence.absenceType.absenceTypeId
        studentAttendance.absence.absenceTypeName = s.absence.absenceType.name
        studentAttendance.absence.absenceReason = s.absence.absenceReason
      }

      else {
        studentAttendance.attendanceType = '출석'
      }
      

      this.classesWithStudents[classIndex].studentAttendances.push(studentAttendance)
    })

	},
}
</script>