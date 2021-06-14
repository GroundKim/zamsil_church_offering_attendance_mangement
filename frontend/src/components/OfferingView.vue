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
          년 헌금 기록
        </h1>
      </v-col>
    </v-row>
    <v-row
      class="d-flex ma-5"
      v-for="o in offeredAtsByMonths"
      :key=o.month
    >
      <v-col
        
      >
        <div class="d-flex pa-2">
          <h2>{{ o.month }}월</h2>
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
                v-for="date in [...o.offeredAts].sort()"
                :key=date
              >
                <td><h3>{{ date.substring(5, 7)}}월 {{ date.substring(8, 10) }}일 ({{ getWeekOfMonth(date) }} 주차)</h3></td>
                <td>
                  <v-btn 
                    @click="downloadExcelByDate(date)"
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
                    @click="showOfferingDiaryDetail(date)"
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
  name: 'OfferingView',
  data () {
    return {
      year: null,
      offeredAts: [],
      offeredAtsByMonths: [],
    }
  },

  methods: {
    getWeekOfMonth(date) {
      const weekOfMonth = (m) => m.week() - moment(m).startOf('month').week() + 1
      return weekOfMonth(moment(date))
    },

    getOfferedAtsByYear (year) {
      axios
      .get(`${this.$serverAddress}/Youth/offering/view/list?year=${year}`, {withCredentials: true})
      .then(res => {
        this.offeredAts = res.data.offeredAts
        this.distingushWithMonth(this.offeredAts)
      })
      .catch(err => {
        this.alertError(err)
      })
    },

    showOfferingDiaryDetail (offeredAt) {
      if (this.$route.name.includes('simple')) this.$router.push(`/simple/offering/view/detail?date=${offeredAt.substring(0,10)}`)
      else this.$router.push(`/offering/view/detail?date=${offeredAt.substring(0,10)}`)
    },

    plusYear () {
      this.year++
    },

    minusYear () {
      this.year--
    },

    downloadExcelByDate (date) {
      let yearMonthDate = date.substring(0,10)
      axios({
        url: `${this.$serverAddress}/Youth/offering/view/excel?date=${yearMonthDate}`,
        method: 'GET',
        responseType: 'blob', // important
        withCredentials: true,
      })
      .then((response) => {
        const url = window.URL.createObjectURL(new Blob([response.data]))
        const link = document.createElement('a')
        link.href = url
        link.setAttribute('download', `헌금통게표_${yearMonthDate}.xlsx`)
        document.body.appendChild(link)
        link.click()
      })
      .catch(err => {
        this.alertError(err)
      })
    },

    distingushWithMonth (dates) {
      dates.forEach(date => {
        const month = date.substring(5, 7)
        let hasFound = false
        this.offeredAtsByMonths.forEach(o => {
          if (month === o.month) {
            hasFound = true
            o.offeredAts.push(date)
          } 
        })
        
        if (!hasFound) {
          let offeredAtsByMonth = {
            month: month,
            offeredAts: [date]
          }
          this.offeredAtsByMonths.push(offeredAtsByMonth)
        }
      })

      // sorting ascending
      this.offeredAtsByMonths.sort((a, b) => {
        return a.month < b.month ? -1 : a.month > b.month ? 1 : 0
      })
    },
  },

  watch: {
    year: function() {
      this.offeredAtsByMonths = []
      this.getOfferedAtsByYear(this.year)
    }
  },

  created () {
    this.year = moment().year()
  },
}
</script>
