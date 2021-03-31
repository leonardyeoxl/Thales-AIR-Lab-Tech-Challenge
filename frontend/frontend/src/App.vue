<template>
  <div id="app">
    <STARSANDSIDS msg="Thales AIR Lab Tech Challenge"/>
  </div>
</template>

<script>
import STARSANDSIDS from './components/STARSANDSIDS.vue'

export default {
  name: 'App',
  components: {
    STARSANDSIDS
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
