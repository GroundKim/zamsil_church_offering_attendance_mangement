<template>
    <v-row>
        <v-col
            cols="2"
        >
            <v-autocomplete 
                label="offering type"
                :items="offeringType"
            ></v-autocomplete>
        </v-col>

        <v-col
            cols="2"
        >
            <v-autocomplete
                label="who"
            ></v-autocomplete>
        </v-col>

        <v-col
            cols="4"
        >
            <v-text-field label ="what cost" :rules="[rules.offering]"></v-text-field>
        </v-col>
    </v-row>
</template>

<script>
export default {
    data: () => ({
        offeringType: ['tith', 'thanks', 'seasonal', 'etc'], //... backend 서버에서 get offering type
        offeringTypeValue: null,
        rules: {
            reuqired: value => !!value || 'Required.',
            offering: value => {
                const pattern = /^[0-9]+$/
                return pattern.test(value)
            }  
        }

    }),
}
</script>