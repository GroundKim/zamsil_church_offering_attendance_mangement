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

        <v-sheet
            v-for="(attendedAt, i) in attendedAts" 
            :key="i"
            class="ma-5"
        >
            {{ attendedAt.substr(0, 10) }}
            <v-btn 
                @click="downloadAttendanceExcelByDate(attendedAt)"
                class="ma-10"
            >
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
            .get(`${this.$serverAddress}/Youth/attendance/view?year=${year}`, {withCredentials: true})
            .then(res => {
                this.attendedAts = res.data.attendedAts
            })
            .catch(err => {
                let errStatusCode = err.response.status
                if (errStatusCode === 404) {
                    alert(err.message + ": 출석부를 불러오는 도중 오류가 발생했습니다 관리자에게 문의하십시오")
                }

                if (errStatusCode === 403) {
                    alert('로그인을 해주세요')
                    this.$router.push('/login')
                }

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
                withCredentials: true,
            }).then((response) => {
                const url = window.URL.createObjectURL(new Blob([response.data]))
                const link = document.createElement('a')
                link.href = url
                link.setAttribute('download', `출석부_${yearMonthDate}.xlsx`)
                document.body.appendChild(link)
                link.click()
            }).catch(err => {
                let errStatusCode = err.response.status
                if (errStatusCode === 404) {
                    alert(err.message + ": 출석부를 불러오는 도중 오류가 발생했습니다 관리자에게 문의하십시오")
                }

                if (errStatusCode === 403) {
                    alert('로그인을 해주세요')
                    this.$router.push('/login')
                }

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
        this.$store.commit('changeHeaderName', '출석부 보기')
    },
}
</script>