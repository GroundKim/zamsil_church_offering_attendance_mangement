<template>
  <div>
    <v-app-bar
      color="blue-grey darken-4"
      dense
      dark
    >
    <v-menu
        left
        bottom
    >
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          icon
          v-bind="attrs"
          v-on="on"
        >
          <v-app-bar-nav-icon></v-app-bar-nav-icon>
        </v-btn>
      </template>

      <v-list>
        <v-list-item
          v-for="(link, index) in links"
          :key="index"
          @click="movePage(link)"
        >
          <v-list-item-title>{{ link.title }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>

      <v-toolbar-title>잠실교회 소년부 출석부 기입</v-toolbar-title>

      <v-spacer></v-spacer>

      
    </v-app-bar>
  </div>
</template>

<script>
  export default {
    data: () => ({
      links: [
        { title: '출석부 기입',
          target: 'attendanceInput'
        },
        { title: '출석부 보기',
          target: 'attendanceView'
        },
        { title: '헌금 기입',
          target: 'offeringInput'
        },
        { title: '헌금 보기',
          target: 'offeringView'
        },
      ]
    }),

    methods: {
      movePage(link) {
        // avoid catch error for same url, but it is not good way
        this.$router.push({name: link.target}).catch(() => {})
      }
    }
  }
</script>
