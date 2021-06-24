<template>
  <v-container>
    <v-row class="d-flex mt-5">
      <v-col>
        <h1>
          <v-btn
            class="mx-2 mb-1"
            fab
            outlined
            x-small
            color="blue-grey"
            @click="minusYear()"
          >
            <v-icon>
              mdi-minus
            </v-icon>
          </v-btn>
          {{ year }}
          <v-btn
            class="mx-2 mb-1"
            fab
            outlined
            x-small
            color="primary"
            @click="plusYear()"
          >
            <v-icon>
              mdi-plus
            </v-icon>
          </v-btn>
          년 출석부 기록
        </h1>
      </v-col>
    </v-row>

    <v-row
      class="d-flex ma-5"
      v-for="a in attendedAtsByMonths"
      :key=a.month
    > 
      <v-col>
        <div class="d-flex pa-2">
          <h2>{{ a.month }}월</h2>
        </div>
        <v-simple-table>
          <template v-slot:default>
            <thead>
              <tr>
                <th>날짜</th>
                <th>엑셀 다운로드 및 세부 보기</th>
              </tr>
            </thead>
            
            <tbody>
              <tr
                v-for="(date, index) in [...a.attendedAts].sort()"
                :key=index
              >
                <td><h3>{{ date.substring(5, 7)}}월 {{ date.substring(8, 10) }}일 ({{ getWeekOfMonth(date) }} 주차)</h3></td>
                <td>
                  <v-btn 
                    @click="downloadAttendanceExcelByDate(date)"
                    fab
                    outlined
                    x-small
                    color="indigo"
                    class="ml-3"
                  >
                    <v-icon>
                      mdi-download
                    </v-icon>
                  </v-btn>

                  <v-btn
                    @click="showAttendanceDetail(date)"
                    fab
                    outlined
                    x-small
                    color="indigo"
                    class="ml-3"
                  >
                    <v-icon>
                      mdi-clipboard-list-outline
                    </v-icon>
                  </v-btn>
                </td>
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
import moment from 'moment'

export default {
  name: 'AttendanceView',

  data () {
    return {
      year: null,
      attendedAts: [],
      attendedAtsByMonths: [],
    }
  },
  
  methods: {
    getWeekOfMonth(date) {
      const weekOfMonth = (m) => m.week() - moment(m).startOf('month').week() + 1
      return weekOfMonth(moment(date))
    },

    getAttendedAtsByYear (year) {
      axios
      .get(`${this.$serverAddress}/Youth/attendance/view/list?year=${year}`, {withCredentials: true})
      .then(res => {
        this.attendedAts = res.data.attendedAts
        this.distingushWithMonth(this.attendedAts)
      })
      .catch(err => {
        this.errorHandler(err)
      })
    },

    showAttendanceDetail (attendedAt) {
      if (this.$route.name.includes('simple')) this.$router.push(`/simple/attendance/view/detail?date=${attendedAt.substring(0,10)}`)
      else this.$router.push(`/attendance/view/detail?date=${attendedAt.substring(0,10)}`)
      
    },

    plusYear () {
      this.year++
    },

    minusYear () {
      this.year--
    },

    downloadAttendanceExcelByDate (date) {
      let yearMonthDate = date.substring(0,10)
      axios({
        url: `${this.$serverAddress}/Youth/attendance/view/excel?date=${yearMonthDate}`, 
        method: 'GET',
        responseType: 'blob', // important
        withCredentials: true,
      }).then((response) => {
        const url = window.URL.createObjectURL(new Blob([response.data]))
        const link = document.createElement('a')
        link.href = url
        link.setAttribute('download', `출석부_${yearMonthDate}.xlsx`)
        document.body.appendChild(link)
        link.click()
      }).catch(err => {
        this.errorHandler(err)
      })
    },

    distingushWithMonth (dates) {
      dates.forEach(date => {
        const month = date.substring(5, 7)
        let hasFound = false
        this.attendedAtsByMonths.forEach(a => {
          if (month === a.month) {
            hasFound = true
            a.attendedAts.push(date)
          } 
        })
        
        if (!hasFound) {
          let attendedAtsByMonth = {
            month: month,
            attendedAts: [date]
          }
          this.attendedAtsByMonths.push(attendedAtsByMonth)
        }
      })

      // sorting ascending 
      this.attendedAtsByMonths.sort((a, b) => {
        return a.month < b.month ? -1 : a.month > b.month ? 1 : 0
      })
    },

    sortAttendedAts (attendedAts) {
      return attendedAts.sort()
    }
  },

  watch: {
    year: function() {
      this.attendedAtsByMonths = []
      this.getAttendedAtsByYear(this.year)
    }
  },

  created () {
    this.year = moment().year()
  },
}
</script>
