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

        {
          label: 'total',
          data: [],
          display: false,
          backgroundColor: 'rgba(24,91,62,0)',
          datalabels: {
            backgroundColor: () => 'white',
            formatter: (item) => {
              if (item == 0) return null
              else return '총 ' + item + '명'
            },
            align: 'top',
            anchor: 'start',
            padding: 0,
          },
        },
      ]
    },
  
    options: {
      scales: {
        yAxes: [{
          ticks: {
            beginAtZero: true,
            max: 0,
          },
          stacked: true,
        }],

        xAxes: [{
          stacked: true,
        }],
      },

      tooltips: {
        filter: (tooltipItem) => {
          if (tooltipItem.datasetIndex === 2) {
            return false
          } else {
            return true
          }
        }
      },

      legend: {
        position: 'bottom',
        labels: {
          filter: (item) => {
            return !item.text.includes('total')
          }
        }
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
            } else {
              return value
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
        if (largestNumber < (d.numberOfDepartmentOne + d.numberOfDepartmentTwo)) largestNumber = d.numberOfDepartmentOne + d.numberOfDepartmentTwo
        this.chartdata.datasets[0].data.push(d.numberOfDepartmentOne)
        this.chartdata.datasets[1].data.push(d.numberOfDepartmentTwo)
        this.chartdata.datasets[2].data.push((d.numberOfDepartmentOne + d.numberOfDepartmentTwo))
        
      })
      this.options.scales.yAxes[0].ticks.max = largestNumber + 15

      // render graph 
      this.renderChart(this.chartdata, this.options)
    }
  },
}
</script>
