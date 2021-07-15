<template>
  <v-container>
    <v-row>
      <h1 class="mt-5">{{ this.year }}년 헌금 통계자료</h1>
    </v-row>

    <v-row>
      <v-col>
         <h2 class="ma-3">통계</h2>
				<v-simple-table
				fixed-header
				>
					<template v-slot:default>
						<thead>
							<tr>
								<th>월</th>
								<th>주일헌금</th>
								<th>십일조헌금</th>
								<th>감사헌금</th>
								<th>절기헌금</th>
								<th>기타헌금</th>
								<th>총합</th>
							</tr>
						</thead>
						<tbody>
							<tr
                v-for="data in offeringData"
                :key="data.month"
              >
                <td>{{ data.month }} 월</td>
                <td>₩ {{ changeCostWithDelimeter(data.weekTotalCost) }}</td>
                <td>₩ {{ changeCostWithDelimeter(data.titheTotalCost) }}</td>
                <td>₩ {{ changeCostWithDelimeter(data.thanksTotalCost) }}</td>
                <td>₩ {{ changeCostWithDelimeter(data.seasonalTotalCost) }}</td>
                <td>₩ {{ changeCostWithDelimeter(data.etcTotalCost) }}</td>
                <td>₩ {{ changeCostWithDelimeter(data.totalCost) }}</td>
              </tr>

              <tr>
                <td> 총합 </td>
                <td>₩ {{ changeCostWithDelimeter(total.weekTotalCost) }} </td>
                <td>₩ {{ changeCostWithDelimeter(total.titheTotalCost) }} </td>
                <td>₩ {{ changeCostWithDelimeter(total.thanksTotalCost) }} </td>
                <td>₩ {{ changeCostWithDelimeter(total.seasonalTotalCost) }} </td>
                <td>₩ {{ changeCostWithDelimeter(total.etcTotalCost) }} </td>
                <td>₩ {{ changeCostWithDelimeter(total.totalCost) }} </td>
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
export default {
  name: 'OfferingStat',
  components:{
  },

  data() {
    return {
      year: null,
      offeringData: null,
      offeringTypes: null,
      total: {
        weekTotalCost: 0,
        titheTotalCost: 0,
        thanksTotalCost: 0,
        seasonalTotalCost: 0,
        etcTotalCost: 0,
        totalCost: 0,
      }
    }
  },

  methods: {
    changeCostWithDelimeter (value) {
			value = value.toString()
			const result = value.replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ",")
			return result
		},
  },

  mounted: async function() {    
    //get year from url parameter

    this.year = this.$route.query.year
    await axios
      .get(`${this.$serverAddress}/Youth/offering/view/statistic/summary?year=${this.year}`, {withCredentials: true})
      .then((res) => {
        this.offeringData = res.data
      })
      .catch((err) => {
        this.errorHandler(err)
      })

    // get total cost
    this.offeringData.forEach(offering => {
      this.total.weekTotalCost += offering.weekTotalCost
      this.total.titheTotalCost += offering.titheTotalCost
      this.total.thanksTotalCost += offering.thanksTotalCost
      this.total.seasonalTotalCost += offering.seasonalTotalCost
      this.total.etcTotalCost += offering.etcTotalCost
      this.total.totalCost += offering.totalCost
    })


  },
}
</script>