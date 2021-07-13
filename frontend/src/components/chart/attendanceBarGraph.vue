<script>
import { Bar } from 'vue-chartjs'

export default {
  extends: Bar,
  props: {
    numberData: {
      type: Array,
      default: null
    },
  },
  name: 'attendanceBarGraph',
  data: () => ({
    chartdata: {
      labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'],
      datasets: [
        {
          label: '1부 인원',
          data: [],
          backgroundColor: '#f87979',
        },

        {
          label: '2부 인원',
          data: [],
          backgroundColor: '#4B89DC',
        },
      ]
    },
  
    options: {
      scales: {
        yAxes: [{
          ticks: {
            beginAtZero: true
          },
          stacked: true,
        }],

        xAxes: [{
          stacked: true,
        }],
       
      },

      legend: {
        position: 'bottom'
      },

      title: {
        display: true,
        text: '소년부 월 별 출석 인원수',
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
            }
          }
        }
      },

      responsive: true,
    }
  }),

  watch: {
    numberData: async function() {
      let largestNumber = 0
      // split data
      await this.numberData.forEach(d => {
        if (largestNumber < d.numberOfDepartmentOne) largestNumber = d.numberOfDepartmentOne
        if (largestNumber < d.numberOfDepartmentTwo) largestNumber = d.numberOfDepartmentTwo
        this.chartdata.datasets[0].data.push(d.numberOfDepartmentOne)
        this.chartdata.datasets[1].data.push(d.numberOfDepartmentTwo)
      })
      this.options.scales.yAxes[0].ticks.max = largestNumber + 10

      // render graph 
      this.renderChart(this.chartdata, this.options)
    }
  },
}
</script>
