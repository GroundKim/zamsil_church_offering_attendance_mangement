<template>
  <v-container fill-height>   

    <v-row class="d-flex mt-10">
      <div class="ml-5">
        <h1>재적관리</h1>
      </div>
      <div class="ml-auto">
        <v-radio-group row v-model="department">
          <v-radio label="1부" value="1"></v-radio>
          <v-radio label="2부" value="2"></v-radio>
        </v-radio-group>
      </div>
    </v-row>

    <v-row
      v-for="(classInfo, index) in currentClasses"
      :key="index"
    >
      <v-col xs12 sm8 md4>
        <v-card
          elevation="0" outlined
          class="mt-10"
        >
          <div class="d-flex pa-5">
            <div>
              <h2>반: {{ classInfo.class.name }}</h2>
              <h3>선생님: {{ getTeacherNames(classInfo.teachers) }}</h3>
            </div>
            
            <div class="ml-auto">
              <v-btn
                fab
                outlined
                small
                color="primary"
                @click="showNewStudentDialog(classInfo)"
              >
                <v-icon>
                  mdi-plus
                </v-icon>
              </v-btn>
            </div>
          </div>

          <v-data-table
            :headers="studentHeaders"
            :items="formatStudent(classInfo.students, '없음')"
            hide-default-footer
            :footer-props="{
              itemsPerPageOptions: [-1]
            }"
            class="elevation-1"
          >
            <template v-slot:[`item.name`]="{ item }">
              <v-row>{{ item.name }} <div class="blue--text">{{ validateNewStudent(item) ? "(신입)":null }}</div></v-row>
            </template>
            <template v-slot:[`item.actions`]="{ item }">
              <v-icon
                @click="editStudent(item, classInfo.class)"
              >
                mdi-pencil
              </v-icon>

              <v-icon
                color="red lighten-2"
                @click="deleteStudent(item, index)"
              >mdi-delete</v-icon>
            </template>
          </v-data-table>
        </v-card>
      </v-col>
    </v-row>

    <v-dialog
      v-model="dialog"
      max-width="500px"
      persistent
    >
      <form @submit.prevent="saveNewStudent">
        <v-card>
          <v-card-title primary-title>
            <div v-if="isAdd">
              <h3>{{ dialogStudent.departmentName }} 부 {{ dialogStudent.className }} 반</h3>
              <p>학생 추가</p>
            </div>

            <div v-else>
              <h3>{{ dialogStudent.departmentName }} 부 {{ dialogStudent.className }} 반</h3>
              <p>{{ dialogStudent.name }} 수정</p>
            </div>

          </v-card-title>
          <v-card-text>
            <v-row>
              <v-col
                cols="12"
                sm="6"
                md="4"
              >
                <v-text-field
                  label="이름"
                  v-model="dialogStudent.name"
                  prepend-icon="mdi-account"
                  :rules="[() => !!dialogStudent.name || '이름을 반드시 입력해주세요']"
                  required
                  @input="$v.name.$touch()"
                  @blur="$v.name.$touch()"
                ></v-text-field>
              </v-col>

              <v-col
                cols="12"
                sm="6"
                md="4"
              >
                <v-text-field
                  v-model="dialogStudent.dayOfBirth"
                  prepend-icon="mdi-cake"
                  label="생일"
                  v-mask="'####-##-##'"
                  :rules="[rules.date]"
                >
                </v-text-field>  
              </v-col>

              <v-col
                cols="12"
                sm="6"
                md="4"
              >
                <v-text-field
                  prepend-icon="mdi-home"
                  v-model="dialogStudent.address"
                  label="주소"
                ></v-text-field>
              </v-col>

              <v-col
                cols="12"
                sm="6"
                md="4"
              >
                <v-text-field
                  prepend-icon="mdi-cellphone"
                  v-model="dialogStudent.phoneNumber"
                  label="학생 번호"
                  v-mask="'###-####-####'"
                ></v-text-field>
              </v-col>

              <v-col
                cols="12"
                sm="6"
                md="4"
              >
                <v-text-field
                  prepend-icon="mdi-human-female-boy"
                  v-model="dialogStudent.parentPhoneNumber"
                  label="부모님 번호"
                  v-mask="'###-####-####'"
                ></v-text-field>
                </v-col>

              <v-col
                cols="12"
                sm="6"
                md="4"
              >
              <v-text-field
                prepend-icon="mdi-school"
                v-model="dialogStudent.schoolName"
                label="학교 이름"
              ></v-text-field>
              </v-col>
            </v-row>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              color="black"
              outlined
              x-large
              @click="closeDialog()"  
            >
              취소
            </v-btn>
            <v-btn
              color="primary"   
              outlined
              x-large
              @click="isAdd ? saveNewStudent() : putStudent()"
              >
                저장
            </v-btn>
          </v-card-actions>
        </v-card>
      </form>
    </v-dialog>
  </v-container>
</template>

<script>
import axios from 'axios'
import moment from 'moment'
import { validationMixin } from 'vuelidate'
import { required } from 'vuelidate/lib/validators'

export default {
  name: 'MemberManagement',
  mixins: [validationMixin],
  validations: {
    name: { required }  
  },

  data() {
    return {
      dialog: false,
      classes: null,
      newStudentDialog: [],
      departmentOneMembers: [],
      departmentTwoMembers: [],
      currentClasses: [],
      dialogClass: null,
      isAdd: false,

      dialogStudent: {
        studentId: null,
        departmentName: null,
        classId: null,
        className: null,
        name: '',
        dayOfBirth: null,
        address: null,
        phoneNumber: null,
        parentPhoneNumber: null,
        schoolName: null,
      },    
      

      rules: {
        date: value => {
          const pattern = /^(?:\d{4}-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])|null|)$/
          return pattern.test(value) || '날짜를 정확히 입력해주세요'
        }

      },
      department: null,
      studentHeaders: [
        {
          text: '학생 이름',
          align: 'start',
          value: 'name',
          sortable: false,
        },

        {
          text: '생일',
          value: 'dayOfBirth',
        },

        {
          text: '주소',
          value: 'address',
        },

        {
          text: '학교 이름',
          value: 'schoolName'
        },

        {
          text:'전화 번호',
          value: 'phoneNumber',
        },

        {
          text: '부모님 전화번호',
          value: 'parentPhoneNumber',
        },

        {
          text: '수정',
          value: 'actions',
          sortable: false
        }
      ],

    }
  },

  methods: {
    getTeacherNames (teachers) {
      let teacherNames = []
      teachers.forEach(teacher => {
        teacherNames.push(teacher.name)
      })
      return teacherNames.toString()
    },

    formatStudent (students) {
      this.studentsForTable = students
      this.studentsForTable.forEach(student => {
        for (const key in student) {
          if (student[key] === null) student[key] = '없음'
        }
        if (student.dayOfBirth !== null) student.dayOfBirth = student.dayOfBirth.substr(0, 10)
      })

      return this.studentsForTable
    },

    deleteStudent (item, classIndex){
      if (confirm(`${item.name}을(를) 정말로 삭제하시겠습니까?`) !== true) {
        return
      }

      axios
        .delete(`${this.$serverAddress}/Youth/students/${item.studentId}`,  
         { withCredentials: true }
        )
        .then(() => {
          let index = this.currentClasses[classIndex].students.findIndex(s => s.studentId == item.studentId)
          this.currentClasses[classIndex].students.splice(index, 1)
          alert(`${item.name}(이)가 삭제되었습니다`)
        })
        .catch((err)=> {
          this.alertError(err)
        })
    },

    saveNewStudent () {
      let dialogStudent = JSON.parse(JSON.stringify(this.dialogStudent))
      if (dialogStudent.dayOfBirth !== null)  dialogStudent.dayOfBirth += moment().format().substr(10)
      let payload = []
      payload.push(dialogStudent)
      
      const headers = {
        'Content-type': "application/json"
      }
      axios
        .post(`${this.$serverAddress}/Youth/students`, JSON.stringify(payload), { withCredentials: true, Headers: headers })
        .then((res) => {
          let student = res.data
          let index = this.currentClasses.findIndex(c => c.class.classId == this.dialogStudent.classId)
          this.currentClasses[index].students = this.currentClasses[index].students.concat(student)
          this.closeDialog()
        })
        .catch((err) => {
          this.alertError(err)
        })
    },

    putStudent () {
      let dialogStudent = JSON.parse(JSON.stringify(this.dialogStudent))
      if (dialogStudent.dayOfBirth !== null)  dialogStudent.dayOfBirth += moment().format().substr(10)
      const headers = {
        'Content-type': 'application/json'
      }
      axios
        .put(`${this.$serverAddress}/Youth/students/${dialogStudent.studentId}`, JSON.stringify(dialogStudent), { withCredentials: true, Headers: headers})
        .then((res) => {
          let student = res.data
          let classIndex = this.currentClasses.findIndex(c => c.class.classId == dialogStudent.classId)
          let studentIndex = this.currentClasses[classIndex].students.findIndex(s => s.studentId == student.studentId)
          Object.assign(this.currentClasses[classIndex].students[studentIndex], student)
          
          this.closeDialog()

        })
        .catch((err) => {
          this.alertError(err)
        })
    },

    editStudent (studentInfo, classInfo) {
      this.isAdd = false
      this.dialog = true
      
      // change '없음' to null
      for (const key in studentInfo) {
        if (studentInfo[key] === '없음') studentInfo[key] = null
      }
      this.dialogStudent.studentId = studentInfo.studentId
      this.dialogStudent.name = studentInfo.name
      this.dialogStudent.dayOfBirth = studentInfo.dayOfBirth
      this.dialogStudent.address = studentInfo.address
      this.dialogStudent.phoneNumber = studentInfo.phoneNumber
      this.dialogStudent.parentPhoneNumber = studentInfo.parentPhoneNumber
      this.dialogStudent.schoolName = studentInfo.schoolName
      this.dialogStudent.departmentName = classInfo.department.name
      this.dialogStudent.className = classInfo.name
      this.dialogStudent.classId = classInfo.classId
    }, 

    closeDialog () {
      this.dialog = false
      this.dialogStudent.name = ''
      this.dialogStudent.dayOfBirth = null
      this.dialogStudent.address = null
      this.dialogStudent.phoneNumber = null
      this.dialogStudent.parentPhoneNumber = null
      this.dialogStudent.schoolName = null
    },

    showNewStudentDialog (classInfo) {
      this.dialogStudent.classId = classInfo.class.classId
      this.dialogStudent.className = classInfo.class.name
      this.dialogStudent.departmentName = classInfo.class.department.name
      this.isAdd = true
      this.dialog = true
    },

    validateNewStudent(item) {
      let isNewStudent = moment.duration(moment().diff(item.enrolledAt)).asDays() <= 14 ? true : false
      return isNewStudent
    }
  },

  watch: {
    department: function() {

      if (this.department == 1) {
        this.currentClasses = this.departmentOneMembers
      }

      if (this.department == 2) {
        this.currentClasses = this.departmentTwoMembers
      }  
    },    
  },

  created: async function () {
    await axios
      .get(`${this.$serverAddress}/Youth/members`, { withCredentials: true })
      .then((res) => {
        this.classes = res.data
      })
      .catch((err) => {
        this.showError(err)
      })

    // split up the members with department name. Department ID 와 Deparmartment name을 분간 할것 attendance Info 컴포넌트와 서버에서도 
    await this.classes.forEach(classInfo => {
      if (classInfo.class.departmentId === 1) {
        this.departmentOneMembers.push(classInfo)
      }
      
      if (classInfo.class.departmentId === 2) {
        this.departmentTwoMembers.push(classInfo)
      }
    })
    this.dialogClass = this.classes[0]
    this.department = '1'
  },
}
</script>

<style>

</style>