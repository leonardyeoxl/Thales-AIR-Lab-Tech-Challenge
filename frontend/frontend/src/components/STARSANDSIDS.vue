<template>
  <div class="hello">
    <h2>SIDs</h2>
    <ul id="app">
      <li v-for="SIDs_item in SIDs_airport_list" :key="SIDs_item">
        {{ SIDs_item.airport }}
        <h3>Waypoints</h3>
        <ul>
          <li v-for="SIDs_waypoint in SIDs_item.topWaypoints" :key="SIDs_waypoint">
          Waypoint: {{ SIDs_waypoint.name }} <br>
          Count: {{ SIDs_waypoint.count }}
          </li>
        </ul>
      </li>
    </ul>

    <h2>STARs</h2>
    <ul id="app">
      <li v-for="STARs_item in STARs_airport_list" :key="STARs_item">
        {{ STARs_item.airport }}
        <h3>Waypoints</h3>
        <ul>
          <li v-for="STARs_waypoint in STARs_item.topWaypoints" :key="STARs_waypoint">
          Waypoint: {{ STARs_waypoint.name }} <br>
          Count: {{ STARs_waypoint.count }}
          </li>
        </ul>
      </li>
    </ul>
  </div>
</template>

<script>
import Vue from 'vue'
import axios from 'axios'

Vue.config.productionTip = false

export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  el: '#app',
  data() {
    return {
      SIDs_airport_list: null,
      STARs_airport_list: null
    }
  },
  mounted () {
    axios
      .get('http://127.0.0.1:8000/STARS-AND-SIDS')
      .then(response => (this.SIDs_airport_list = response.data.SIDs_airport_list, this.STARs_airport_list = response.data.STARs_airport_list))
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
