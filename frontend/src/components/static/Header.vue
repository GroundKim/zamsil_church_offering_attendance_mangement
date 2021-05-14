<template>
  <div>
    <v-app-bar
      color="blue-grey darken-4"
      dense
      dark
    >
      <v-tabs>
        <v-tab
          v-for="(link, index) in links"
          :key="index"
          @click="movePage(link)"
        >
          <v-icon>mdi-{{ link.icon }}</v-icon>
          <span class="d-none d-md-flex ml-2">{{ link.title }}</span>
        </v-tab>
      </v-tabs>
    </v-app-bar>
  </div>
</template>

<script>
export default {
  data: () => ({
    links: [
      {
        icon: 'clipboard-edit',
        title: '출석부 기입',
        target: 'attendanceInput'
      },
      {
        icon: 'clipboard-outline',
        title: '출석부 보기',
        target: 'attendanceView'
      },
      {
        icon: 'book-edit',
        title: '헌금 기입',
        target: 'offeringInput'
      },
      {
        icon: 'book-outline',
        title: '헌금 보기',
        target: 'offeringView'
      },
    ]
  }),

  methods: {
      movePage(link) {
        // avoid catch error for same url, but it is not a good way
        this.headerName = link.title
        this.$router.push({name: link.target}).catch(() => {})
      },
    },

    mounted () {
      this.headerName = this.$store.state.headerName
    },
}
</script>
