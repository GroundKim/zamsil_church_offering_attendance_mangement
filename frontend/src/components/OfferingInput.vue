<template>
    <v-container>
        <h1>헌금통계표</h1>
        <br>
        <v-form>
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
                    <v-text-field label="작성자" :ruels="[rules.required]"></v-text-field>
                </v-col>

                <v-col
                    cols="2"
                    class="mx-10"
                >
                    <v-slider
                        v-model="department"
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
                        label="주일헌금"
                        outlined
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

        <component v-for="item in offerings" :key="item" :is="item"></component>
        <v-btn type="submit" width="700">제출</v-btn>

        </v-form>
        <span>{{ department }}</span>
    </v-container>
</template>

<script>
import SpecificOfferingInput from "./SpecificOfferingInput"

import axios from "axios";
export default {
    components: {
        SpecificOfferingInput,
    },

    data: () => ({
        date: new Date().toISOString().substr(0, 10),
        menu: false,
        modal: false,
        menu2: false,
        department: null,
        studentData: [],
        departmentsLabel: ['1부', '2부'],
        students: [],
        offerings: [],
        rules: {
            reuqired: value => !!value || 'Required.',
            offering: value => {
                const pattern = /^[0-9]+$/
                return pattern.test(value)
            }  
        }
    }),
    methods: {
        addOffering() {
            this.offerings.push('SpecificOfferingInput')
        },

        deleteOffering() {
            this.offerings.pop()
        }
    },
    watch: {
        department : async function () {
            let getURL = `http://localhost:8080/Youth/students?department_id=${this.department + 1}`
            await axios
            .get(getURL)
            .then((response) => {
                this.studentData = response.data
                })
            .catch(err => {
                alert(err.message + " 학생정보를 불러오는 도중 오류가 발생했습니다 관리자에게 문의하십시오")
            })  
        }
    },

    created() {
        this.offerings.push('SpecificOfferingInput')
    },
  }
</script>
