<template>
  <v-container class="grey lighten-5">
    <v-row no-gutters>
      <template >
        <v-row >
          <v-col v-for="n in attendanceCount" :key="n">
          <v-card class="pa-2 text-center" >
            1
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
    }
  },

  async created() {
    const response = await axios.get("http://localhost:8080/Youth/input/1")
    console.log(response.data)
    this.attendanceCount = response.data["count"]
      
      
  }
}
</script>