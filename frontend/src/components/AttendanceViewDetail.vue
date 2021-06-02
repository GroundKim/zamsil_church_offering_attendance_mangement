<template>
	<v-container>
			<v-row
				v-for="(classInfo, index) in classes"
				:key="index"
			>
				<v-col>
					{{classInfo}}
				</v-col>
			</v-row>
	</v-container>
</template>

<script>
import axios from 'axios'

export default {
	data() {
		return {
			date: null,
			classes: null,
		}
	},

	created () {
		let date = this.$route.query.date
		axios
			.get(`${this.$serverAddress}/Youth/members`, { withCredentials: true })
			.then((res) => {
				this.classes = res.data
				console.log(JSON.stringify(this.classes))
			})

		axios
			.get(`${this.$serverAddress}/Youth/attendnace/view?date=${date}`, { withCredentials: true })
			.then(() => {
			})
	},
}
</script>