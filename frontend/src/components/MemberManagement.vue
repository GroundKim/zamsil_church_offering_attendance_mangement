<template>
  <div class="">Member Management
    <v-container fill-height>

      <v-radio-group v-model="department">
        <v-radio label="1부" value="1"></v-radio>
        <v-radio label="2부" value="2"></v-radio>
      </v-radio-group>


      <v-card
        v-for="classInfo in currentClasses"
        :key="classInfo.Class.classId"
      >
      <br>
      <br>

      </v-card>          
    </v-container>

    
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'MemberManagement',
  data() {
    return {
      classes: null,
      departmentOneMembers: [],
      departmentTwoMembers: [],
      currentClasses: [],

      listHeaders: [
        {
          classManagement: [
            text: '반',
            align: 'start',
            value: 'className'

          ]
          },

          teacherManagement: {

          },

          studentManagement: {

          }
        }
      ],
      department: null,

  
  },

  watch: {
    department: function() {
      if (this.department == 1) {
        this.currentClasses = this.departmentOneMembers
      }

      if (this.department == 2) {
        this.currentClasses = this.departmentTwoMembers
      }
    }
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
    this.classes.forEach(classInfo => {
      if (classInfo.Class.departmentId === 1) {
        this.departmentOneMembers.push(classInfo)
      }
      
      if (classInfo.Class.departmentId === 2) {
        this.departmentTwoMembers.push(classInfo)
      }
    })
    console.log(JSON.stringify(this.departmentOneMembers))
    
    this.department = '1'
  },

    

    

  
}
</script>

<style>

</style>