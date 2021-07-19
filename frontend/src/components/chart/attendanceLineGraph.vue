<script>
import { Line } from 'vue-chartjs'

export default ({
  extends: Line,
  props: {
    attendanceDataForLine: {
      type: Array,
      default:null 
    },
  },
  data() {
    return {
      chartdata: {
      labels: [],
      datasets: [
        {
          label: '명 수',
          data: [],
          backgroundColor: '#f87979',
          fill: false,
          tension: 0,
          borderColor: '#f87979',
          pointBorderWidth: 7,
          pointHoverBorderWidth: 10,
          datalabels: {
            backgroundColor: () => 'white',
            formatter: (item) => {
              return item + '명'
            },
            anchor: 'end',
            align: 'end',
            padding: 0,
          }
        },
      ]
      },

      options: {
        scales: {
          yAxes: [{
            ticks: {
              beginAtZero: false,
              max: 5,
            },
          }],
        },

        legend: {
          position: 'bottom'
        },

        title: {
          display: true,
          text: '소년부 출석 통계',
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
          }
        },

        responsive: true,
      }
    }
  },

  watch: {
    attendanceDataForLine: async function () {
      let [...attendanceDataForLine] = this.attendanceDataForLine
      
      let largestNumber = 0
      attendanceDataForLine.forEach(diaryData => {
        this.chartdata.labels.push(diaryData.date.substr(0, 10))
        let totalNumber = 0
        diaryData.attendanceInfo.forEach(info => {
          totalNumber = totalNumber + info.students.length
        })
        if (largestNumber < totalNumber) largestNumber = totalNumber
        this.chartdata.datasets[0].data.push(totalNumber)
      })
      this.options.scales.yAxes[0].ticks.max = largestNumber + 30
      this.renderChart(this.chartdata, this.options)
    }
  },
})
</script>
