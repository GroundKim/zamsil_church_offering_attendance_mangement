<template>
  <v-container>
    <v-row>
      <h1 class="mt-5">{{ this.year }}년 통계자료</h1>
    </v-row>
    
  </v-container>
</template>

<script>
import axios from 'axios'
export default {
  data() {
    return {
      year: null,
      data: null,
    }
  },
  
  mounted () {
    //get year from url parameter
    this.year = this.$route.query.year
    axios
      .get(`${this.$serverAddress}/Youth/attendance/statistic/${this.year}`)
      .then((res) => {
        this.data = res.data
      })
      .catch((err) => {
        err
      })
  },
}
</script>