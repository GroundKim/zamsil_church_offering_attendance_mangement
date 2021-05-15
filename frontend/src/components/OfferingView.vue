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

    <v-row>
      <v-col>
        <v-sheet
          v-for="(offeredAt, i) in offeredAts" 
          :key="i"
          class="ma-3 pa-2"
        >
          <h3>{{ offeredAt.substr(0, 10) }}
            <v-btn 
              @click="downloadExcelByDate(offeredAt)"
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
  name: 'OfferingView',

  data () {
    return {
      year: null,
      offeredAts: [],
    }
  },

  methods: {
    getOfferedAtsByYear (year) {
      axios
      .get(`${this.$serverAddress}/Youth/offering/view?year=${year}`, {withCredentials: true})
      .then(res => {
        this.offeredAts = res.data.offeredAts
      })
      .catch(err => {
        let errStatusCode = err.response.status
        if (errStatusCode === 404) {
          alert(err.message + ": 오류 발생 관리자에게 문의하십시오")
        }

        if (errStatusCode === 403) {
          alert('로그인을 해주세요')
          this.$router.push('/login')
        }
      })
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
        let errStatusCode = err.response.status
        if (errStatusCode === 404) {
          alert(err.message + ": 오류 발생 관리자에게 문의하십시오")
        }

        if (errStatusCode === 403) {
          alert('로그인을 해주세요')
          this.$router.push('/login')
        }
      })
    }
  },

  watch: {
    year: function() {
      this.getOfferedAtsByYear(this.year)
    }
  },

  created () {
    this.$store.commit('changeHeaderName', '헌금 보기')
    this.year = moment().year()
  },
}
</script>
