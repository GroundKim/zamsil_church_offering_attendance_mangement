<template>
  <div>
    <v-app-bar
      color="blue-grey darken-4"
      dense
      dark
    >
      <v-tabs class="d-flex justify-center" v-model="activeTab">
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
    activeTab: 1,
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

      {
        icon: 'domain',
        title: '재적 관리',
        target: 'memberManagement'
      }
    ]
  }),

  methods: {
      movePage(link) {
        // avoid catch error for same url, but it is not a good way
        this.headerName = link.title
        this.$router.push({name: link.target}).catch(() => {})
      },

      selectHeader(routeName) {
        const index = this.links.findIndex(l => l.target === routeName)
        this.activeTab = index
      }
    },

    mounted () {
      this.headerName = this.$store.state.headerName
      // avoid selecting wrong header v-tab when refreishing from client
      this.selectHeader(this.$router.currentRoute.name)
      document.title = '잠실교회 소년부 관리'
    },


}
</script>
