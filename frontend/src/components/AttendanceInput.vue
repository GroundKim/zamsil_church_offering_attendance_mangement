<template>
  <v-container>
    <form @submit.prevent="submit">
      <v-row>
        <v-col
          v-for="data in attendanceData"
          v-bind:key="data.classId"
          cols="4"
        >
          <v-card height="400">
            <h2 class="pa-5">Class {{ data.classId }}</h2>
            <h3> {{ data.teacherName }} </h3>
            
            <v-row>
              <v-col v-for="student in data.studentsIdandName" v-bind:key="student.studentId" cols="4">
                <v-checkbox v-model="checkedNames" :value=student.studentId :label="`${student.studentName}`"></v-checkbox>
              </v-col>
            </v-row>
            

            
          </v-card>
        </v-col>
      </v-row>
      <v-container class="text-center ma-10">
        <v-btn type="submit" width="700">제출</v-btn>
      </v-container>
      
    </form>

    <span>{{checkedNames}}</span>
  </v-container>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      checkedNames: [],
      attendanceData: null
    }
  },
  methods: {
    submit() {

      axios.post("http://localhost:8080/Youth/post/attendance/department/1")
    }

  },
   created() {
    axios.get("http://localhost:8080/Youth/resource/attendance/department/1")
     .then(response => {
       this.attendanceData = response.data
     })
    
  }
}
</script>