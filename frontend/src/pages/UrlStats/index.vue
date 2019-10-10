<template>
    <div>
        <h1>Stats for {{url_short}}</h1>
        <p>Original URL: {{url}}</p>
        <p>Number of requests: {{qnt}}</p>
        <line-chart :discrete="true" :data=this.graphDt></line-chart>
    </div>
</template>

<script>
    export default {
        name: 'app',
        data() {
            return {
                page: null,
                url: null,
                url_short: null,
                stats: null,
                times: {},
                qnt: null,
                graphDt: '',
            }
        },
        created() {
            this.page = this.$route.params.page;
            this.axios.get(this.page + "/stats")
                .then(response => {
                    this.url = response.data.url;
                    this.url_short = response.data.url_short;
                    this.stats = response.data.data;
                    this.stats.forEach(this.timeParse);
                    this.graphDt = JSON.parse(JSON.stringify(this.times));
                })
        },
        methods: {
            timeParse(item) {
                this.qnt += 1;
                let date = new Date(item.time * 1000).toISOString();
                if (date in this.times) {
                    this.times[date] += 1
                } else {
                    this.times[date] = 1
                }
            },
        }
    }
</script>

<style scoped>
</style>