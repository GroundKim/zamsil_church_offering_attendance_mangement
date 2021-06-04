<template>
  <v-container>
    <h2>{{ date }}</h2>
    <v-row
      v-for="(classInfo, index) in classesWithStudents"
      :key="index"
    >
      <v-container>
        
        <h2>{{ classInfo.class.department.name }}부 {{ classInfo.class.name }}반</h2>
        <v-col>
          
          <v-data-table
            :headers="studentHeaders"
            :items="getAttendanceDiaryTableItems(classInfo)"
            class="elevation-1"
            loading="true"
            hide-default-footer
          >
            
          </v-data-table>
        </v-col>
        </v-container>
      </v-row>
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
        absenceType: null,
        absenceReason: null,
      },

      attendanceClass: {
        attendanceDiaries: [],
      },


      // studentHeaders: [
      //   {
      //     text: '학생 이름',
      //     align: 'start',
      //     value: 'student.name',
      //     sortable: false,
      //   },

      //   {
      //     text: '출석 여부',
      //     value: 'absence',

      //   },

      //   {
      //     text: '결석 종류',
      //     value: 'absence.absenceType.name',

      //   },

      //   {
      //     text: '결석 이유',
      //     value: 'absence.reason'
      //   }
      // ]
		}
	},

  computed: {
    studentHeaders () {
      return [
        {
          text: '학생 이름',
          align: 'start',
          value: 'name',
        },

        {
          text: '출석 여부',
          value: 'absence.isAbsent',

        },

        {
          text: '결석 종류',
          value: 'absence.absenceTypeName',

        },

        {
          text: '결석 이유',
          value: 'absence.reason'
        }
      ]
    }
  },


  methods: {
    getAttendanceDiaryTableItems(classInfo) {
      let index = this.classesWithStudents.findIndex(c => c.class.classId == classInfo.class.classId)
      return this.classesWithStudents[index].studentAttendances
    }
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

    // split up the students which is divided by class
    this.studentsWithAttendance.forEach(s => {
      let classIndex = this.classesWithStudents.findIndex(c => c.class.classId == s.student.classId)
      let studentAttendance = {
        studentId: s.student.studentId,
        classId: s.student.classId,
        name: s.student.name,
        absence: {
          isAbsent: 'O',
          absenceType: null,
          reason:null,
        }
      }
      
      if (s.absence !== null) {
        studentAttendance.absence.isAbsent = 'X'
        studentAttendance.absence.absenceTypeName = s.absence.absenceType.name
        studentAttendance.absence.reason = s.absence.reason
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