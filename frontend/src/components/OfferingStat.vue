<template>
  <v-container>
    <v-row>
      <h1 class="mt-5">{{ this.year }}년 헌금 통계자료</h1>
    </v-row>

    <v-row>
      <v-col>
      </v-col>

      <v-col>

      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from 'axios'
export default {
  name: 'OfferingStat',
  components:{
  },

  data() {
    return {
      year: null,
      offeringData: null,
    }
  },

  beforeMount () {    
    //get year from url parameter
    this.year = this.$route.query.year
    axios
      .get(`${this.$serverAddress}/Youth/offering/view/statistic/summary?year=${this.year}`, {withCredentials: true})
      .then((res) => {
        this.offeringData = res.data
      })
      .catch((err) => {
        this.errorHandler(err)
      })
  },
}
</script>