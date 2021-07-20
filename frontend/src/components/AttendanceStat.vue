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

    <!-- class -->
    <v-row
      v-for="data in classData"
      :key="data.class.classId"
    >
      <v-col>
        <v-simple-table>
          <template v-slot:default>
            <thead>
              <tr>
                <th width="15%">반 개요</th>
                <th width="10%">재적</th>
                <th v-for="numberByDate, nbdIndex in data.numberByDates" :key="nbdIndex">{{ numberByDate.date.substr(5,10) }}</th>
              </tr>
            </thead>

            <tbody>
              <tr>
                <td>{{ data.class.department.name }}부 {{ data.class.name }}반 {{ getTeachersName(data.teachers) }}</td>
                <td>{{ data.classHeadcount }} 명</td>
                <td v-for="numberByDate, nbdIndex in data.numberByDates" :key="nbdIndex">{{ numberByDate.number }} 명</td>
              </tr>
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
      classData: null,
    }
  },

  methods: {
    getTeachersName(teachers) {
      let teachersName = ""
      teachers.forEach(teacher => {
        teachersName += teacher.name
      })
      return teachersName
    }
  },

  beforeMount: async function () {    
    //get year from url parameter
    this.year = this.$route.query.year
    await axios
      .get(`${this.$serverAddress}/Youth/attendance/view/statistic/number/month?year=${this.year}`, {withCredentials: true})
      .then((res) => {
        this.numberData = res.data
      })
      .catch((err) => {
        err
      })
    
    await axios
      .get(`${this.$serverAddress}/Youth/attendance/view/statistic?year=${this.year}`, { withCredentials: true })
      .then((res) => {
        let data = res.data
        data.sort((a, b) => {
          return new Date(a.date) - new Date(b.date)
        })
        this.attendanceDataForLine = data
      })
      .catch((err) => {
        err
      })
    
    await axios
      .get(`${this.$serverAddress}/Youth/attendance/view/statistic/summary?year=${this.year}`, { withCredentials: true })
      .then((res) => {
        let classInfo = res.data
        for (let i = 0; i < classInfo.length; i++) {
          const numberByDates = classInfo[i].numberByDates;
          if (numberByDates !== null) {
            numberByDates.sort((a, b) => {
              return new Date(a.date) - new Date(b.date) 
            })
          }
        }
        this.classData = classInfo
      })
      .catch((err) => {
        err
      })
  },
}
</script>