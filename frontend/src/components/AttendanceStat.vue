<template>
  <v-container>
    <v-row>
      <h1 class="mt-5">{{ this.year }}년 통계자료</h1>
    </v-row>
    
    <v-row>
      <barGraph/>
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
      data: null,
    }
  },
  
  mounted () {
    this.renderChart()
    
    //get year from url parameter
    this.year = this.$route.query.year
    axios
      .get(`${this.$serverAddress}/Youth/attendance/view/statistic/number/month?year=${this.year}`, {withCredentials: true})
      .then((res) => {
        this.data = res.data
        console.log(this.data);
      })
      .catch((err) => {
        err
      })
  },
}
</script>