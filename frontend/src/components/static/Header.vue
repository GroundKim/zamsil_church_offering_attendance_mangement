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
    activeTab: 0,
    links: [
      {
        icon: 'clipboard-edit',
        title: '출석부 기입',
        target: 'attendanceInput',
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
      },
    ],

    simpleLinks: [
      {
        icon: 'clipboard-outline',
        title: '출석부 보기',
        target: 'simpleAttendanceView'
      },

      {
        icon: 'book-outline',
        title: '헌금 보기',
        target: 'simpleOfferingView'
      },
    ],

    loginLink: {
      icon: 'login',
      title: '로그인',
      target: 'login'
    },
    
    logoutLink: {
      icon: 'logout',
      title: '로그아웃',
    }
    
  }),

  watch: {
    $route() { 
      this.setHeader()
    },

    getLoginStatus: function () {
      if (this.getLoginStatus) {
        let index = this.links.findIndex(l => l.title === '로그인')
        this.links.splice(index, 1)
        this.links.push(this.logoutLink)
      }
    }
  },

  computed: {
    getLoginStatus() {
      return this.$store.state.isLogin
    }
  },

  methods: {
    movePage(link) {
      // if user want to logout
      if (link.title === '로그아웃') {
        // change login status
        this.$store.commit('changeLoginStatus', false)

        // remove auth cookie
        this.$cookies.remove('auth_token')

        // remove logout activetab and push login
        let logoutIndex = this.links.findIndex(l => l.title === '로그아웃')
        this.links.splice(logoutIndex, 1)
        this.links.push(this.loginLink)

        // move to login page
        this.$router.push({name: 'login'}).catch(() => {})
        return
      }
      // avoid catch error for same url, but it is not a good way
      this.$router.push({name: link.target}).catch(() => {})
    },

    setHeader() {
      let headerTabName = null
      if (this.$store.state.headerActiveTabName !== null) {
        headerTabName = this.$store.state.headerActiveTabName

        // set state.headerActiveTabName null
        this.$store.commit('changeHeaderActiveTabName', null)
      } else {
        headerTabName = this.$router.currentRoute.name

        if (headerTabName.includes('attendanceView')) headerTabName = 'attendanceView'
        if (headerTabName.includes('offeringView')) headerTabName = 'offeringView'
      }

      const index = this.links.findIndex(l => l.target === headerTabName)
      this.activeTab = index
    }
  },

  mounted () {
    if (this.$route.name.includes('simple')) {
      this.links = this.simpleLinks
    }

    if (this.$cookies.isKey('auth_token')) {
      this.links.push(this.logoutLink)
    } else {
      this.links.push(this.loginLink)
    }
    
    this.setHeader()
  },
}
</script>
