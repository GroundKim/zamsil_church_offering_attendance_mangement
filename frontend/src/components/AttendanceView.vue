<template>
    <v-container>
        <div class="float-md-left">
            <h1> {{ year }}년 기록 </h1>
        </div>
        <v-container>
            <v-btn
                class="mx-2"
                fab
                dark
                x-small
                color="indigo"
                @click="plusYear()"
            >
                <v-icon dark>
                    mdi-plus
                </v-icon>
            </v-btn>

            <v-btn
                class="mx-2"
                fab
                dark
                x-small
                color="primary"
                @click="minusYear()"
            >
                <v-icon dark>
                    mdi-minus
                </v-icon>
            </v-btn>
        </v-container>

        <v-sheet v-for="(attendedAt, i) in attendedAts" :key="i">
            {{ attendedAt }}
            <v-btn @click="downloadExcelByDate(attendedAt)">
                excel download
            </v-btn>
        </v-sheet>
    </v-container>
</template>

<script>
import axios from 'axios'
import moment from 'moment'
export default {
    name: 'AttendanceView',

    data () {
        return {
            year: null,
            attendedAts: [],
        }
    },

    methods: {
        getAttendedAtsByYear (year) {
            axios
            .get(`${this.$serverAddress}/Youth/attendance/view?year=${year}`)
            .then(res => {
                this.attendedAts = res.data.attendedAts
                alert(this.attendedAts)
            })
            .catch(err => {
                alert(err.message + "에러 발생")
            })
        },

        plusYear () {
            this.year++
        },

        minusYear () {
            this.year--
        },

        downloadAttendanceExcelByDate (date) {
            let yearMonthDate = date.substring(0,10)
            axios({
                url: `${this.$serverAddress}/Youth/attendance/view/excel?date=${yearMonthDate}`,
                method: 'GET',
                responseType: 'blob', // important
            }).then((response) => {
                const url = window.URL.createObjectURL(new Blob([response.data]))
                const link = document.createElement('a')
                link.href = url
                link.setAttribute('download', `헌금통게표_${yearMonthDate}.xlsx`)
                document.body.appendChild(link)
                link.click()
            })
            
        }
    },

    watch: {
        year: function() {
            this.getAttendedAtsByYear(this.year)
        }
    },

    created () {
        this.year = moment().year()
    },
}
</script>