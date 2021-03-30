<template>
  <div class="hello">
    <ul id="app">
      <li v-for="item in airport_list" :key="item.SIDs_airport_list">
        {{ item.SIDs_airport_list }}
      </li>
      <li v-for="item in airport_list" :key="item.STARs_airport_list">
        {{ item.STARs_airport_list }}
      </li>
    </ul>
  </div>
</template>

<script>
import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'

Vue.config.productionTip = false
Vue.use(VueAxios, axios)

export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  el: '#app',
  data() {
    return {
      airport_list: null
    }
  },
  mounted () {
    axios
      .get('http://127.0.0.1:8000/STARS-AND-SIDS')
      .then(response => (this.airport_list = response))
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
