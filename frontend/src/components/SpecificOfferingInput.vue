<template>
    <v-row>
        <v-col
            cols="2"
        >
            <v-autocomplete 
                v-model="offeringTypeValue"
                label="offering type"
                :items="offeringType"
                item-text="offeringName"
                required
                return-object
            ></v-autocomplete>
        </v-col>

        <v-col
            cols="2"
        >
            <v-autocomplete
                v-model="offerorValue"
                :items="students"
                label="who"
                item-text="name"
                required
                return-object
            ></v-autocomplete>
        </v-col>

        <v-col
            cols="4"
        >
            <v-text-field
                v-model="offeringAmountValue"
                label ="what cost" 
                required
                :rules="[rules.offering]"
            ></v-text-field>
        </v-col>
    </v-row>
</template>
<script>

export default {
    name: 'SpecificOfferingInput',
    data: () => ({
        offeringType: [], //... backend 서버에서 get offering type
        offeringTypeValue: null,
        offerorValue: null,
        offeringAmountValue: null,
        students: null,

        rules: {
            offering: value => {
                const pattern = /^[0-9]+$/
                return pattern.test(value)
            }  
        }
    }),
    
    props:['testProp'],
    methods: {
        getStudents() {
            return this.$store.getters.getStudents
        },

        watchSendPost() {
            const {sendPost} = this
            return sendPost
        },

        makePayload() {
            let offeringPayload = {
                offeringTypeId: this.offeringTypeValue['offeringTypeId'],
                studentId: this.offerorValue['studentId'],
                offeringAmount: this.offeringAmountValue,
            }
            this.$store.commit('pushOffering', offeringPayload)
        }

        
    },

    watch: {
        getStudents() {
            this.students = this.$store.getters.getStudents
        },

        testProp: function() {
            this.makePayload()
        }

    },

    created() {
        this.students = this.$store.getters.getStudents
        this.offeringType = this.$store.getters.getOfferingType
    },
}
</script>