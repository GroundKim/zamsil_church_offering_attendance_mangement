<script>
import { Line } from 'vue-chartjs'

export default ({
  extends: Line,
  props: {
    offeringByMonth: {
      type: Array,
      default:null 
    },
  },
  data() {
    return {
      chartdata: {
      labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'],
      datasets: [
        {
          label: '총 헌금',
          data: [],
          backgroundColor: '#f87979',
          fill: false,
          tension: 0,
          borderColor: '#f87979',
          pointBorderWidth: 7,
          pointHoverBorderWidth: 10,
        },
      ]
      },

      options: {
        scales: {
          yAxes: [{
            ticks: {
              beginAtZero: true,
              max: 0,
              callback: function (label) {
                label = label.toString()
                const result = label.replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ",")
                return '₩ ' + result
              }
            },
          }],
        },

        legend: {
          position: 'bottom'
        },

        title: {
          display: true,
          text: '소년부 월 별 총 헌금',
          font: { weight: 'bold'},
          fontSize: 30,
          padding: 15
        },

        plugins: {
          datalabels: {
            font: {
              weight: 'bold',
              size: 15
            },

            formatter: function (value) {
              if (value === 0) {
                return null
              } else {
                value = value.toString()
                const result = value.replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ",")
                return '₩ ' + result
              } 
            }
          }
        },

        responsive: true,
      }
    }
  },

  watch: {
    offeringByMonth: async function () {
      let largestNumber = 0
      await this.offeringByMonth.forEach(offering => {
        if (largestNumber < offering.totalCost) largestNumber = offering.totalCost
        this.chartdata.datasets[0].data.push(offering.totalCost)
      })
      
      let firstDigit = parseInt(largestNumber.toString()[0]) + 1
      let remainedDigits = []
      for (let index = 0; index < largestNumber.toString().length - 1; index++) {
        remainedDigits.push(0)
      }
      let maxNumber = firstDigit.toString() + remainedDigits.join('')
      this.options.scales.yAxes[0].ticks.max = parseInt(maxNumber)
      this.renderChart(this.chartdata, this.options)
    }
  },
})
</script>
