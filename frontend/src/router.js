import Vue from 'vue';
import Router from 'vue-router';
import UrlForm from './pages/UrlForm/index.vue';
import UrlStats from './pages/UrlStats/index.vue';

Vue.use(Router);

export default new Router({
    mode: 'history',
    routes: [{
        path: '/create',
        name: 'UrlForm',
        component: UrlForm,
    },
        {
            path: '/:page/stats',
            name: 'Stats',
            component: UrlStats,
        },
        {
            path: '/',
            redirect: '/create',
        },
    ]
});