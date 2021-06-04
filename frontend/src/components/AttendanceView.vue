<template>
  <v-container>
    <v-row class="mt-5">
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

    <v-row>
      <v-col>
        <v-sheet
          v-for="(attendedAt, i) in attendedAts" 
          :key="i"
          class="ma-3 pa-2"
        >
          <h3>{{ attendedAt.substr(0, 10) }}
            <v-btn 
              @click="downloadAttendanceExcelByDate(attendedAt)"
              fab
              outlined
              x-small
              color="indigo"
              class="ml-3"
            >
              <v-icon>
                mdi-download
              </v-icon>
              <h2>엑셀</h2>
            </v-btn>

            <v-btn
              @click="showAttendanceDetail(attendedAt)"
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
          </h3>
        </v-sheet>
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
    }
  },
  
  methods: {
    getAttendedAtsByYear (year) {
      axios
      .get(`${this.$serverAddress}/Youth/attendance/view/list?year=${year}`, {withCredentials: true})
      .then(res => {
        this.attendedAts = res.data.attendedAts
      })
      .catch(err => {
        this.alertError(err)
      })
    },

    showAttendanceDetail(attendedAt) {
      this.$router.push(`/attendance/view/detail?date=${attendedAt.substring(0,10)}`)
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
        this.alertError(err)
      })
    }
  },

  watch: {
    year: function() {
      this.getAttendedAtsByYear(this.year)
    }
  },

  created () {
    this.year = moment().year()
    this.$store.commit('changeHeaderName', '출석부 보기')
  },
}
</script>
