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
                @click="addYear()"
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


        <component v-for="item, i in attendanceAndOfferingComponents" :key="i" :is="item"></component>
    </v-container>
</template>

<script>
import AttendanceAndOffering from './dynamic/AttendanceAndOffering'

import axios from 'axios'
import moment from 'moment'
export default {
    name: 'AttendanceAndOfferingView',
    component: {
        AttendanceAndOffering,
    },

    data: () => ({
        year: null,
        attendanceAndOfferingComponents: [],

    }),

    methods: {
        getExcelByYear(year) {
            axios
            .get(`http://localhost:8080/Youth/offering/view?year=${year}`)
        },

        addYear() {
            this.year++
        },

        minusYear() {
            this.year--
        }

    },

    watch: {
        year: function() {
            this.getExcelByYear(this.year)
        }
    },

    created() {
        this.year = moment().year()
    },



}



</script>