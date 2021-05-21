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
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Noto+Sans+KR:wght@200;300;600&display=swap');

html, body, #app, .v-application {
  font-family: 'Noto Sans KR', sans-serif;
}
</style>
