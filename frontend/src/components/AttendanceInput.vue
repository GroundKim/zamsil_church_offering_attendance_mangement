<template>
  <v-container class="grey lighten-5">

  
    <v-row no-gutters>
      <template >
        <v-row >
          <v-col v-for="n in attendanceCount" :key="n">
          <v-card class="pa-2 text-center" >
            
          </v-card> 
          </v-col>
        </v-row>
      </template>
    </v-row>
  </v-container>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      attendanceData: [],
      attendanceCount: 0,

      a : null,
    }
  },

   created() {
     axios.get("http://localhost:8080/Youth/resource/attendance/department/1")
     .then(function (response) {
       this.a = response
     })


      
    
    
  
      
  }
}
</script>