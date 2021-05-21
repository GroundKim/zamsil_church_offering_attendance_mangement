<template>
  <div class="">Member Management</div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'MemberManagement',
  data() {
    return {
      classes: [],
      students: [],
    }
  },

  created () {
    axios
      .get(`${this.$serverAddress}/Youth/classes`, { withCredentials: true })
      .then((res) => {
        this.classes = res.data
      })
      .catch((err) => {
        this.alert(err)
      })

    axios
      .get(`${this.$serverAddress}/Youth/students`, { withCredentials: true })
      .then((res) => {
        this.students = res.data
      })

    axios
      .get(`${this.$serverAddress}/Youth/teachers`)
  },

    

    

  
}
</script>

<style>

</style>