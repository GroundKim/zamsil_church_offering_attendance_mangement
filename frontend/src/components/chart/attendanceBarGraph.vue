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
        }
      ]
    },
  
    options: {
      scales: {
        yAxes: [{
          ticks: {
            max: Math.max(40),
            beginAtZero: true
          }
        }],

        xAxes: [{
          stacked:false,
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
      responsive: true,
    }
  }),

  watch: {

  },

  mounted () {
    console.log(JSON.stringify(this.numberData));
    this.renderChart(this.chartdata, this.options)
  }
}

</script>
