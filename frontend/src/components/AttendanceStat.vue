<template>
  <v-container>
    <v-row>
      <h1 class="mt-5">{{ this.year }}년 통계자료</h1>
    </v-row>

    <v-row>
      <v-col>
        <attendance-bar-graph :numberData="numberData" ></attendance-bar-graph>
      </v-col>
    </v-row>

    <v-row>
      <v-col>
        <attendance-line-graph :attendanceDataForLine="attendanceDataForLine"></attendance-line-graph>
      </v-col>
    </v-row>

    <v-row>
      <v-col>
        <v-simple-table>
          <template v-slot:default>
            <thead>
              <tr>
                <th>1부</th>
                <th>재적</th>
                <th 
                  v-for="data in attendanceDataForLine"
                  :key="data.date"
                >
                  {{ data.date.substring(5, 7) }}월 {{ data.date.substring(8,10) }}일
                </th>
              </tr>
            </thead>

            <tbody>


            </tbody>
          </template>
        </v-simple-table>
      </v-col>
    </v-row>
  </v-container>
</template>
<script>
import axios from 'axios'
import attendanceBarGraph from './chart/attendanceBarGraph.vue'
import attendanceLineGraph from './chart/attendanceLineGraph.vue'

export default {
  name: 'AttendanceStat',
  components:{
    attendanceBarGraph,
    attendanceLineGraph,
  },

  data() {
    return {
      year: null,
      numberData: null,
      attendanceDataForLine: null,
    }
  },

  beforeMount () {    
    //get year from url parameter
    this.year = this.$route.query.year
    axios
      .get(`${this.$serverAddress}/Youth/attendance/view/statistic/number/month?year=${this.year}`, {withCredentials: true})
      .then((res) => {
        this.numberData = res.data
      })
      .catch((err) => {
        err
      })
    
    axios
      .get(`${this.$serverAddress}/Youth/attendance/view/statistic?year=${this.year}` , { withCredentials: true })
      .then((res) => {
        let data = res.data
        data.sort((a, b) => {
          return new Date(a.date) - new Date(b.date);
        })
        this.attendanceDataForLine = data
      })
      .catch((err) => {
        err
      })
  },
}
</script>