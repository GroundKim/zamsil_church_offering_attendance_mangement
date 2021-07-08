<template>
  <v-container>
    <v-row>
      <h1 class="mt-5">{{ this.year }}년 통계자료</h1>
    </v-row>
    
    <v-row>
      <barGraph :numberData="numberData"></barGraph>
    </v-row>
  </v-container>
</template>
<script>
import axios from 'axios'
import barGraph from './chart/attendanceBarGraph.vue'
export default {
  name: 'AttendanceStat',
  components:{
    barGraph
  },

  data() {
    return {
      year: null,
      numberData: null,
    }
  },

  
  beforeMount: function () {    
    //get year from url parameter
    this.year = this.$route.query.year
    axios
      .get(`${this.$serverAddress}/Youth/attendance/view/statistic/number/month?year=${this.year}`, {withCredentials: true})
      .then((res) => {
        this.numberData = res.data
        console.log(JSON.stringify(this.numberData));
        
              
      this.numberData.forEach(d => {
        this.chartdata.datasets[0].data.push(d.numberOfDepartmentOne)
        this.chartdata.datasets[1].data.push(d.numberOfDepartmentTwo)
      })
      console.log(this.chartdata.datasets[0]);
    

      })
      .catch((err) => {
        err
      })
  },
}
</script>