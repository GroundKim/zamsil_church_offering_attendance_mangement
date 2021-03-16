<template>
  <v-app>
    <Header/>
    <v-main>
      <router-view></router-view>
    </v-main>
    <Footer/>
  </v-app>
</template>

<script>
import Header from './components/static/Header'
import Footer from './components/static/Footer'
export default {
  name: 'App',
  components: {
    Header,
    Footer,
  },

};
</script>
