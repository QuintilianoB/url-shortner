import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import VueAxios from 'vue-axios'

Vue.config.productionTip = false;

new Vue({
  router,
  VueAxios,
  axios,
  render: h => h(App),
}).$mount('#app');

const axiosConfig = {
  baseURL: 'http://127.0.0.1:8000/',
  timeout: 30000,
};

Vue.prototype.$axios = axios.create(axiosConfig);
