<template>
   <v-app id="inspire">
		<v-container fluid fill-height>
			<v-layout align-center justify-center>
				<v-flex xs12 sm8 md4>
					<form @submit.prevent="sendPost">
						<v-card class="elevation-12">
							<v-toolbar dark color="primary">
								<v-toolbar-title>Login form</v-toolbar-title>
							</v-toolbar>
							<v-card-text>
								<v-text-field
									v-model="id"
									@keyup.enter="sendPost"
									label="Login"
								></v-text-field>
								<v-text-field
									v-model="password"
									@keyup.enter="sendPost"
									label="Password"
								></v-text-field>
							</v-card-text>
							<v-card-text
							>
								{{ errorMessage }}
								
							</v-card-text>
							<v-card-actions>
								<v-spacer></v-spacer>
								<v-btn @click="sendPost" color="primary">Login</v-btn>
							</v-card-actions>
						</v-card>
					</form>
				</v-flex>
			</v-layout>
		</v-container>
	</v-app>
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
			axios.post(`${this.$serverAddress}/Youth/login`, JSON.stringify(payload))
				.then(res => {
					console.log(this.$serverAddress)
				})
				.catch(err => {
					console.log(err)
					if (err.response.status === 401) {
						this.password = null
						this.errorMessage = "잘못된 ID와 PASSWORD를 입력하였습니다"
					}
				})
		}
    },

	created() {
		console.log(this.$serverAddress)
	}
}
</script>

<style>

</style>