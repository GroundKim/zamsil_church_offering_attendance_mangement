<template>
	<v-container grid-list-xs>
		{{ offeringData }}
		<v-data-table
			:headers="totalOfferingTableHeaders"
			class="elevation-1"
			loading="true"
			hide-default-footer
		>
			
		</v-data-table>

		<v-data-table
			:headers="individualOfferingTableHeaders"
			:items="offeringData"
			class="elevation-1"
			hide-default-footer
			loading="true"
		>
			<template v-slot:[`item.offeringCost`]="{ item }">
				{{item}}
				<!-- ₩ {{changeCostWithDelimeter(item.offeringCost)}} -->
			</template>
			
			<template v-slot:[`item.action`]="{ item }">
				<v-icon
					@click="showOfferingEditDialog(item)"
				>
					mdi-pencil
				</v-icon>

				<v-icon
					color="red lighten-2"
					@click="deleteOfferingDiary(item)"
				>mdi-delete</v-icon>
			</template>
		</v-data-table>
		<v-dialog
			v-model="dialog"
			max-width="500px"
			transition="dialog-transition"
		>
			<v-card>
				<v-card-title primary-title>
					{{ offeringDiaryDialog.departmentName }} 부 {{ offeringDiaryDialog.className }}반 {{ offeringDiaryDialog.studentName }} 수정
				</v-card-title>
				<v-card-text>
					<v-list-item-group
						v-model="offeringDiaryDialog.offeringTypeId"
						mandatory
					>
						<v-list-item
							v-for="type in offeringTypes"
							:key="type.offeringTypeId"
							:value="type.offeringTypeId"
						>{{ type.offeringTypeName }}</v-list-item>
					</v-list-item-group>

					<v-text-field
						v-model="offeringDiaryDialog.offeringCost"
						label="헌금 액수"
					></v-text-field>
				</v-card-text>
				<v-card-actions>
					<v-btn color="fail">취소</v-btn>
					<v-btn color="success">저장</v-btn>
				</v-card-actions>
			</v-card>	
		</v-dialog>
	</v-container>
</template>
<script>
import axios from 'axios'
export default {
	name: 'OfferingViewDetail',
	data() {
		return {
			date: null,
			offeringData: [],
			offeringTypes: [],
			dialog: false,
			offeringDiaryDialog: {
				action: null,
				offeringDiaryId: null,
				studentName: null,
				departmentName: null,
				className: null,
				offeringTypeId: null,
				offeringCost: null
			},

			individualOfferingTableHeaders: [
				{
					text: '이름',
					align: 'start',
					value: 'student.name',
				},

				{
					text: '부서',
					value: 'departmentId'
				},

				{
					text: '반',
					value: 'student.class.name'
				},

				{
					text: '헌금 종류',
					value: 'offeringType.offeringTypeName'
				},

				{
					text: '헌금 액수',
					value: 'offeringCost',
				},

				{
					text: '수정',
					value: 'action',
				}
			],

			totalOfferingTableHeaders: [
				{
					text: '1부',
					align: 'start',
					value: 'departmentOne'
				},

				{
					text: '2부',
					value: ' departmentTwo'
				}
			],

			offeringTableItems: null,
			offerings: [
				// departmentOne
				{
					totalWeekOfferingCost: 0,
					totalTithOfferingCost: 0,
					totalThanksOfferingCost: 0,
					totalSeasonalOfferingCost: 0,
					totalEtcOfferingCost: 0,
				},

				// departmentTwo
				{
					totalWeekOfferingCost: 0,
					totalTithOfferingCost: 0,
					totalThanksOfferingCost: 0,
					totalSeasonalOfferingCost: 0,
					totalEtcOfferingCost: 0,
				}
			]
		}
	},

	methods: {
		changeCostWithDelimeter (value) {
			value = value.toString()
      const result = value.replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ",")
			return result
		},

		showOfferingEditDialog (item) {
			this.dialog = true
			this.offeringDiaryDialog.offeringDiaryId = item.offeringDiaryId
			this.offeringDiaryDialog.studentName = item.student.name
			this.offeringDiaryDialog.departmentName = item.departmentId
			this.offeringDiaryDialog.className = item.student.class.name
		},

		editOfferingDiary (item) {
			axios
				.put(`${this.$serverAddress}/Youth/offering/${item.offeringId}`)
				.then((res) => {
					let index = this.offeringdata.findIndex(o => o.offeringId === res.data.offeringId)
					this.offeringData[index] = res.data
				})
				.catch((err) => {
					this.alertError(err)
				})
		},

		delteOfferingDiary (item) {
			axios
				.delete(`${this.$serverAddress}/Youth/offering/${item.offeringId}`)
				.then((res) => {
					this.offeringdata.remove(o => o.offeringId === res.data.offeringId)
				})
				.catch((err) => {
					this.alertError(err)
				})
		}
	},

	created: async function () {
		// get date from URL query
		this.date = this.$route.query.date

		// get offeringTypes
		axios
			.get(`${this.$serverAddress}/Youth/offering/types`, { withCredentials: true })
			.then((res) => {
				this.offeringTypes = res.data
			})
			.catch((err) => {
				this.alertError(err)
			})

		// get offering Diary data with date from server
		await axios
			.get(`${this.$serverAddress}/Youth/offering/view?date=${this.date}`, { withCredentials: true })
			.then((res) => {
				this.offeringData = res.data
			})
			.catch((err) => {
				this.alertError(err)
			})

		// split offerings into ...
		this.offeringData.forEach(offering => {
			//split with department 
			if (offering.departmentId === 1) {
				console.log(offering['offeringType'])
			}

			if (offering.departmentId === 2) {
				console.log('123')
			}
		})
	},
}
</script>
