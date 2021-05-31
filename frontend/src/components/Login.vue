<template>
  <v-container fill-height>
    <v-layout align-center justify-center>
      <v-flex xs12 sm8 md4>
        <form @submit.prevent="sendPost">
          <v-card elevation="12" shaped>
            <v-card-title><h3>로그인</h3></v-card-title>
            <v-card-text>
              <v-text-field
                v-model="id"
                @keyup.enter="sendPost"
                label="아이디"
              ></v-text-field>
              <v-text-field
                v-model="password"
                type="password"
                @keyup.enter="sendPost"
                label="비밀번호"
              ></v-text-field>
            </v-card-text>
            <v-card-text v-if="errorMessage">
              {{ errorMessage }}
            </v-card-text>
            <v-card-actions>
              <v-btn @click="sendPost" block large rounded color="primary" elevation="0"><strong>로그인</strong></v-btn>
            </v-card-actions>
          </v-card>
        </form>
      </v-flex>
    </v-layout>
  </v-container>
</template>


<script>
import axios from 'axios'
export default {
  data() {
    return {
      id: null,
      password: null,
      errorMessage: null,
    }
  },

  methods: {
    sendPost() {
      let payload = {
        clientId: this.id,
        password: this.password,
      }

      axios.post(`${this.$serverAddress}/Youth/login`, JSON.stringify(payload), {withCredentials: true})
        .then(()=> {
          alert("로그인 되었습니다!")
          this.$router.push('/attendance/input')
          }
        )
        .catch(err => {
          this.alertError(err)
        })
    }
  },

	created() {
	}
}
</script>
