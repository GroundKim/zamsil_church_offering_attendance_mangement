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
                    <v-text-field label="작성자"></v-text-field>
                </v-col>
            </v-row>

            <v-row>
                <v-col
                    cols="2"
                >
                    <v-radio-group v-model="radioGroup" mandatory> 
                        <v-radio label="1부" value="1">1부</v-radio>
                        <v-radio label="2부" value="2">2부</v-radio>
                    </v-radio-group>
                </v-col>
            </v-row>
        </v-form>
        



    </v-container>
</template>

<script>
  export default {
    data: () => ({
      date: new Date().toISOString().substr(0, 10),
      menu: false,
      modal: false,
      menu2: false,
    }),
  }
</script>
