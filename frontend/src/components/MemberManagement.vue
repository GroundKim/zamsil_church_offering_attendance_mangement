<template>
  <div class="">Member Management
    <v-container fill-height>
      <v-layout align-center justify-center>
        <v-flex xs12 sm8 md4>
          <v-card
            v-for="classItem in members"
            :key="classItem"
          >
            {{ classItem }}

          </v-card>          
        </v-flex>
      </v-layout>
    </v-container>

    
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'MemberManagement',
  data() {
    return {
      members: null
    }
  },

  created: async function () {
    await axios
      .get(`${this.$serverAddress}/Youth/members`, { withCredentials: true })
      .then((res) => {
        this.members = res.data
      })
      .catch((err) => {
        this.showError(err)
      })

    // split up the members with department name. Department ID 와 Deparmartment name을 분간 할것 attendance Info 컴포넌트와 서버에서도 
    this.members.forEach(member => {
      if (member.class.departmentId === 1)
      member 
    })

  },

    

    

  
}
</script>

<style>

</style>