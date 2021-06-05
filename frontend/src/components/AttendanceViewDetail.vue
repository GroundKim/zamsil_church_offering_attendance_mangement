<template>
  <v-container>
    <h2>{{ date }}</h2>
    <v-row
      v-for="(classInfo, classIndex) in classesWithStudents"
      :key="classIndex"
    >
      <v-container>
        <h2>{{ classInfo.class.department.name }}부 {{ classInfo.class.name }}반</h2>
        <v-col>
          <!-- {{ getAttendanceDiaryTableItems(classInfo)}} -->
          <v-data-table
            :headers="studentHeaders()"
            :items="getAttendanceDiaryTableItems(classInfo)"
            class="elevation-1"
            loading="true"
            hide-default-footer
          >
            <template v-slot:[`item.attendanceType`]="{ item, index }">
                <v-edit-dialog
                  large
                  @save="changeAttendance(classIndex, index, item.attendanceType)"
                >
                  <div>{{ item.attendanceType }}</div>
                  <template v-slot:input>
                    <v-list-item-group
                      v-model="item.attendanceType"
                      :mandatory=true
                    >
                      <v-list-item value="참석">참석</v-list-item>
                      <v-list-item value="결석">결석</v-list-item>
                    </v-list-item-group>
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

export default {
	data() {
		return {
			date: null,
      studentsWithAttendance: [],
			classesWithStudents: [], 
      attendanceDiaries: {
        class: null,
        students: null,
      },


      attendanceClasses: [],

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
          value: 'absence.reason'
        }
      ]
    },

    changeAttendance(classIndex, studentIndex, value) {
      if (value === '참석') {
        axios.post()
        classIndex
        studentIndex
      } 
      
      if (value === '결석'){
        axios.delete()
      }
    }
  },

  watch: {

  },

	created: async function () {
		this.date = this.$route.query.date
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

    await axios
      .get(`${this.$serverAddress}/Youth/attendance/view?date=${this.date}`, { withCredentials: true })
      .then((res) => {
        this.studentsWithAttendance = res.data
      })
      .catch(() => {
      })

    // split up the students with divided by class
    this.studentsWithAttendance.forEach(s => {
      let classIndex = this.classesWithStudents.findIndex(c => c.class.classId == s.student.classId)
      let studentAttendance = {
        studentId: s.student.studentId,
        classId: s.student.classId,
        name: s.student.name,
        attendanceType: null,
        absence: {
          isAbsent: false,
          absenceTypeId: null,
          absenceTypeName: null,
          reason:null,
        }
      }
      
      if (s.absence !== null) {
        studentAttendance.attendanceType = '결석'
        studentAttendance.absence.isAbsent = true
        studentAttendance.absence.absenceTypeId = s.absence.absenceType.absenceTypeId
        studentAttendance.absence.absenceTypeName = s.absence.absenceType.name
        studentAttendance.absence.reason = s.absence.reason
      } 
      else {
        studentAttendance.attendanceType = '출석'
      }
      this.classesWithStudents[classIndex].studentAttendances.push(studentAttendance)
    })


    // // handling members
		// await axios
		// 	.get(`${this.$serverAddress}/Youth/members`, { withCredentials: true })
		// 	.then((res) => {
		// 		this.classes = res.data
		// 	})

    // await this.classes.forEach(c => {
    //   let attendanceClass = {
    //     attendanceDiaries: [],
    //   }
    //   c.students.forEach(student => {
    //     let attendanceDiary = {
    //       studentId: student.studentId,
    //       name: student.name,
    //       isAttended: '결석',
    //       absenceType: '일반결석',
    //       absenceReason: null
    //     }
    //     attendanceClass.attendanceDiaries.push(attendanceDiary)
    //   })    
    
    //   this.attendanceClasses.push(attendanceClass)
    // })

    // // handing attendance
		// await axios
		// 	.get(`${this.$serverAddress}/Youth/attendance/view?date=${this.date}`, { withCredentials: true })
		// 	.then((res) => {
    //     let attendedDiaries = res.data
    //     attendedDiaries.forEach(diary => {
    //     let classId = diary.student.classId
    //     let studentId = diary.studentId  
        
    //     let classIndex = this.classes.findIndex(c => c.class.classId == classId)
    //     let studentIndex = this.attendanceClasses[classIndex].attendanceDiaries.findIndex(a => a.studentId == studentId)
    //     this.attendanceClasses[classIndex].attendanceDiaries[studentIndex].isAttended = '출석'
    //     this.attendanceClasses[classIndex].attendanceDiaries[studentIndex].absenceType = null
    //     this.attendanceClasses[classIndex].attendanceDiaries[studentIndex].absenceReason = null          
    //     })
		// 	})

    // // handling absence
    
	},
}
</script>