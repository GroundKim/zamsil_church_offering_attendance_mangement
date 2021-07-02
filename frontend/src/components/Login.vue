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
              <p class="red--text" :class="{'shake' : shakeAnimated}">{{ errorMessage }}</p>
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
      errorMessage: false,
      shakeAnimated: false,
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
          this.$store.commit('changeLoginStatus', true)
          }
        )
        .catch(err => {
          if (err.response.status === 401) {
            this.errorMessage = "아이디 및 암호를 잘못 입력하셨습니다."
            this.shakeErrorMessage()
          }
        })
    },

    shakeErrorMessage() {
      this.shakeAnimated = true
      setTimeout(() => {
        this.shakeAnimated = false
      }, 1000)
    }    
  },

	created() {
    // check the client auth token is valid
    axios
      .get(`${this.$serverAddress}/Youth/login`, { withCredentials: true })
      .then(() => {
        alert('이미 로그인 되어있습니다')
        // distinguish simple or not
        if (this.$route.name.includes('simple')) this.$router.push({ path: "/simple/attendance/view"});
        else this.$router.push({ path: "/attendance/input"})

        // change login status to true
        this.$store.commit('changeLoginStatus', true)
        
      })
      .catch(() => {})
	}
}
</script>

<style>
.shake {
  animation: shake 0.82s cubic-bezier(.36,.07,.19,.97) both;
  transform: translate3d(0, 0, 0);
}

@keyframes shake {
  10%, 90% {
    transform: translate3d(-1px, 0, 0);
  }
  20%, 80% {
    transform: translate3d(2px, 0, 0);
  }
  30%, 50%, 70% {
    transform: translate3d(-4px, 0, 0);
  }
  40%, 60% {
    transform: translate3d(4px, 0, 0);
  }
}
</style>
