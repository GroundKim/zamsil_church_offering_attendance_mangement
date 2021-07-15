<script>
import { Line } from 'vue-chartjs'

export default ({
  extends: Line,
  props: {
    dataByMonth: {
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
    }
  },
})
</script>
