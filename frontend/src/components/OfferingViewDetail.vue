<template>
	<v-container grid-list-xs>
    <v-simple-table
      fixed-header 
    >
      <template v-slot:default>
        <thead>
          <tr>
            <th>헌금 종류</th>
            <th>1부</th>
            <th>2부</th>
            <th>총합</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="type in offeringSummary"
            :key="type.offeringType.offeringTypeId"
          >
            <td>{{ type.offeringType.offeringTypeName }}</td>
            <td>₩ {{ changeCostWithDelimeter(type.offeringCost[0]) }}</td>
            <td>₩ {{ changeCostWithDelimeter(type.offeringCost[1]) }}</td>
            <td>₩ {{ changeCostWithDelimeter(type.offeringCost[0] + type.offeringCost[1]) }}</td>
          </tr>

          <tr>
            <td>총합</td>
            <td>₩ {{ changeCostWithDelimeter(getTotalCostByDepartmentId(1)) }}</td>
            <td>₩ {{ changeCostWithDelimeter(getTotalCostByDepartmentId(2)) }}</td>
            <td>₩ {{ changeCostWithDelimeter(getTotalCostByDepartmentId(1) + getTotalCostByDepartmentId(2)) }}</td>
          </tr>
        </tbody>
      </template>

    </v-simple-table>

		<v-data-table
			:headers="individualOfferingTableHeaders"
			:items="getOfferingData"
			class="elevation-1"
			hide-default-footer
		>
			<template v-slot:[`item.student.name`]= "{ item }">{{ item.studentId===null ? "무명" :item.student.name }}</template>
			<template v-slot:[`item.offeringCost`]="{ item }">
				₩ {{changeCostWithDelimeter(item.offeringCost)}}
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
					<div>
						<h3><span>{{ offeringDiaryDialog.departmentId }} 부</span> <span v-show="offeringDiaryDialog.studentId!==null">{{offeringDiaryDialog.className }} 반</span></h3>
						<p>{{ offeringDiaryDialog.studentName }} 수정</p>
					</div>
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
						prefix="₩"
						:rules="[rules.offering, rules.required]"
						required
					></v-text-field>
				</v-card-text>
				<v-card-actions>
					<v-btn color="fail">취소</v-btn>
					<v-btn color="success" @click="editOfferingDiary()">저장</v-btn>
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
      offeringItem: null,
			offeringData: [],
			offeringTypes: [],
			dialog: false,
			offeringDiaryDialog: {
				offeringDiaryId: null,
				studentName: null,
				departmentId: null,
				className: null,
				offeringTypeId: null,
				offeringCost: null
			},

			rules: {
				required: (value) => !!value || "Required.",
				offering: (value) => {
					const pattern = /^(\d+|\d{1,3}(,\d{3})*)(\.\d+)?$/
					return pattern.test(value)
				}
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

			offeringTableItems: null,
			offeringSummary: {
        createdBy: '',
      },
		}
	},
	
	computed: {
		getDialogOfferingCost() {
			return this.offeringDiaryDialog.offeringCost
		},

    getOfferingData () {
      return this.offeringData
    }
	},

	watch: {
		getDialogOfferingCost(value) {
			this.offeringDiaryDialog.offeringCost = this.changeCostWithDelimeter(value)
		}
	},

	methods: {
		changeCostWithDelimeter (value) {
			value = value.toString()
      const result = value.replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ",")
			return result
		},

    getTotalCostByDepartmentId(deparmtentId) {
      let totalCost = 0
      this.offeringSummary.forEach(o => {
        totalCost += o.offeringCost[deparmtentId - 1]
      }) 
      return totalCost
    },

		showOfferingEditDialog (item) {
      this.offeringItem = item
			this.dialog = true
      this.offeringDiaryDialog.offeringDiaryId = item.offeringDiaryId
			this.offeringDiaryDialog.studentId = item.studentId
			this.offeringDiaryDialog.departmentId = item.departmentId
			this.offeringDiaryDialog.offeringDiaryId = item.offeringDiaryId
			this.offeringDiaryDialog.offeringTypeId = item.offeringTypeId
			this.offeringDiaryDialog.offeringCost = item.offeringCost
			// handling no name offering
			if (item.studentId === null) {
				this.offeringDiaryDialog.studentName = "무명"
				
			} else {
				this.offeringDiaryDialog.studentName = item.student.name
				this.offeringDiaryDialog.className = item.student.class.name

			}
		},


		editOfferingDiary () {
      // remove all delimeter ',' from offeringCost, and conversion string to int
      this.offeringDiaryDialog.offeringCost = Number(this.offeringDiaryDialog.offeringCost.replace(',', ''))
			const headers = {
				'content-type': 'application/json'
			}
			axios
				.put(`${this.$serverAddress}/Youth/offering`, JSON.stringify(this.offeringDiaryDialog), { withCredentials: true, headers: headers })
				.then((res) => {
          this.dialog = false
					this.offeringItem.offeringTypeId = res.data.offeringTypeId
          // I don't know why copy the res data whole object into the ordinary offering diary, so that I put fit in property into each variable
          // get offering name by id
          this.offeringTypes.forEach(type => {
            if(type.offeringTypeId === res.data.offeringTypeId) this.offeringItem.offeringType.offeringTypeName = type.offeringTypeName
          })

          // change offering cost
          this.offeringItem.offeringCost = res.data.offeringCost
				})
				.catch((err) => {
					this.alertError(err)
				})
		},

		deleteOfferingDiary (item) {
      if (confirm(`${item.student.name}의 헌금 기록을 삭제할까요?`)) {
        axios 
          .delete(`${this.$serverAddress}/Youth/offering/${item.offeringDiaryId}`, { withCredentials: true })
          .then((res) => {
            const index = this.offeringData.findIndex(o => o.offeringDiaryId === res.data.offeringDiaryId )
            this.offeringData.splice(index, 1)
          })
          .catch((err) => {
            this.alertError(err)
          })
      }
		},

    getOfferingSummary() {
      axios
        .get(`${this.$serverAddress}/Youth/offering/summary/${this.date}`, { withCredentials: true })
        .then((res) => {
          this.offeringSummary = res.data
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
		await axios
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

    // get offering summary
    this.getOfferingSummary()
	},
}
</script>
