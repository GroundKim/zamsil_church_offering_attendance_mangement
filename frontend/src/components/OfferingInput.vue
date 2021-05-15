<template>
    <v-container>
        <br>
        <form @submit.prevent="sendPost">
            <v-row>
                <v-col
                    cols="12"
                    sm="6"
                    md="4"
                >
                    <v-menu
                        ref="menu"
                        v-model="menu"
                        :close-on-content-click="false"
                        :return-value.sync="date"
                        transition="scale-transition"
                        offset-y
                        min-width="auto"
                    >
                        <template v-slot:activator="{ on, attrs }">
                            <v-text-field
                                v-model="date"
                                label="날짜"
                                prepend-icon="mdi-calendar"
                                readonly
                                v-bind="attrs"
                                v-on="on"
                            ></v-text-field>
                        </template>

                        <v-date-picker
                        v-model="date"
                        no-title
                        scrollable
                        >
                        <v-spacer></v-spacer>
                        <v-btn
                            text
                            color="primary"
                            @click="menu = false"
                        >
                            Cancel
                        </v-btn>
                        <v-btn
                            text
                            color="primary"
                            @click="$refs.menu.save(date)"
                        >
                            OK
                        </v-btn>
                        </v-date-picker>
                    </v-menu>
                </v-col>

                <v-col
                    cols="4"
                >
                    <v-text-field
                        v-model="createdBy"
                        label="작성자"
                        :ruels="[rules.required]" 
                        required
                        ></v-text-field>
                </v-col>

                <v-col
                    cols="2"
                    class="mx-10"
                >
                    <v-slider
                        v-model="departmentId"
                        :tick-labels="departmentsLabel"
                        :max="1"
                        step="1"
                        ticks="always"
                        tick-size="4"
                    ></v-slider>
                </v-col>
            </v-row>

            <v-row>
                <v-col
                    cols="5"
                >
                    <v-text-field
                        v-model="weekOfferingCost"
                        label="주일헌금"
                        outlined
                        required
                        :rules="[rules.offering, rules.required]"
                    >
                    </v-text-field>
                </v-col>
            </v-row>

        <v-btn
            class="mx-2"
            fab
            dark
            color="indigo"
            @click="addOffering()"
        >
            <v-icon dark>
                mdi-plus
            </v-icon>
        </v-btn>

        <v-btn
            class="mx-2"
            fab
            dark
            color="primary"
            @click="deleteOffering()"
        >
            <v-icon dark>
                mdi-minus
            </v-icon>
        </v-btn>

        <component v-for="item, i in offerings" :key="i" :is="item" :offerPostTrigger="specificOfferingTrigger"></component>

        <div class="text-center ma-15">
            <v-btn
            type="submit"
            style="width: 50%"
            class="white--text"
            color="indigo"
            elevation="4"
            x-large
            ><h3>제출</h3></v-btn>
        </div>

        </form>
        <span>{{ departmentId }}</span>
    </v-container>
</template>

<script>
import SpecificOfferingInput from "./SpecificOfferingInput"

import axios from 'axios'
import moment from 'moment'
export default {
    name: 'OfferingInput',
    components: {
        SpecificOfferingInput,
    },

    data: () => ({
        date: moment().format('yyy-MM-DD'),
        menu: false,
        modal: false,
        menu2: false,
        departmentId: null,
        studentData: [],
        departmentsLabel: ['1부', '2부'],
        weekOfferingCost: null,
        createdBy: null,
        students: [],
        offerings: [],
        specificOfferingTrigger: 0,

        rules: {
            required: value => !!value || 'Required.',
            offering: value => {
                const pattern = /^[0-9]+$/
                return pattern.test(value)
            }  
        }
    }),


    methods: {
        async sendPost() {
            let hasPostError = false
            this.specificOfferingTrigger++
            let specificOfferingPayload = await this.$store.getters.getOfferingPayload
            this.date += moment().format().substr(10, )
            let weekOfferingPayload = {
                studentId: null,
                offeringTypeId: 1,
                offeringCost: parseInt(this.weekOfferingCost),
                departmentId: parseInt(this.departmentId + 1),
                offeredAt: this.date,
                createdAt: moment().format(),
                createdBy: this.createdBy,
            }
            console.log(specificOfferingPayload)
            specificOfferingPayload.push(weekOfferingPayload)
            console.log(JSON.stringify(specificOfferingPayload))
            
            const headers = {
                'Content-Type': 'application/json',
            }

            await axios
                .post(
                    `${this.$serverAddress}/Youth/offering`, JSON.stringify(specificOfferingPayload), {headers: headers}, {withCredentials: true}
                )
                .then(res => {
                    console.log(res.data)
                    alert("등록 완료!")
                    this.$store.commit('deleteSpecificOffering')
                })
                .catch(err => {
                    let errStatusCode = err.response.status
                    if (errStatusCode === 404) {
                        alert(err.message + ": 등록중 오류 발생 관리자에게 문의하십시오")
                    }

                    if (errStatusCode === 403) {
                        alert('로그인을 해주세요')
                        this.$router.push('/login')
                    }
                    hasPostError = true
                    this.$store.commit('deleteSpecificOffering')
                })
            
            if (!hasPostError) {
                window.location.reload()
            }
        },

        addOffering() {
            this.offerings.push('SpecificOfferingInput')
        },

        deleteOffering() {
            this.offerings.pop()
        },

        setStudents: async function() {
            let getURL = `${this.$serverAddress}/Youth/students?department_id=${this.departmentId + 1}`
            await axios
            .get(getURL, {withCredentials: true})
            .then((response) => {              
                this.$store.commit('setStudents', response.data)
                })
            .catch(err => {
                alert(err.message + " 학생정보를 불러오는 도중 오류가 발생했습니다 관리자에게 문의하십시오")
            })
        },

        setOfferingType: async function() {
            let getURL = `${this.$serverAddress}/Youth/offering/type`
            await axios
                .get(getURL, {withCredentials: true})
                .then((response) => {
                    this.$store.commit('setOfferingType', response.data)
                })
                .catch(err => {
                    alert(err.message + " offeriny type을 불러오는 도중 오류가 발생했습니다. 관리자에게 문의하십시오")
                })
        },

        setCreatedBy() {
            this.$store.commit('setCreatedBy', this.createdBy)
        },

        setOfferedAt() {
            this.$store.commit('setOfferedAt', this.date)
        },

        setDepartmentId() {
            this.$store.commit('setDepartmentId', this.departmentId + 1)
        }

    },

    watch: {
       departmentId: async function () {
           await this.setStudents()
           this.setDepartmentId()
       },

       createdBy: function () {
           this.setCreatedBy()
       },

       date: function() {
           this.setOfferedAt()
       }
    },

    async created() {
        await this.$store.commit('changeHeaderName', '헌금 기입')
        await this.setStudents()
        await this.setOfferingType()
        this.setOfferedAt()
        this.offerings.push('SpecificOfferingInput')
    },
  }
</script>
