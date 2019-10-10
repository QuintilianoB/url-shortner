import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import VueAxios from 'vue-axios'
import Chartkick from 'vue-chartkick'
import Chart from 'chart.js'

Vue.config.productionTip = false;

Vue.use(VueAxios, axios);
Vue.use(Chartkick.use(Chart));

axios.defaults.baseURL = 'http://127.0.0.1:8000/';

new Vue({
  router,
  render: h => h(App),
}).$mount('#app');
