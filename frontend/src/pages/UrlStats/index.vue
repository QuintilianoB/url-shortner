<template>
    <div>
        <p>{{page}}</p>
        <p>{{url}}</p>
        <p>{{url_short}}</p>
        <p>{{stats}}</p>
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
            }
        },
        created() {
            this.page = this.$route.params.page;
            this.axios.get('/' + this.page + "/stats")
                .then(response => {
                    this.url = response.data.url,
                    this.url_short = response.data.url_short,
                    this.stats = response.data.data
                })
        },
    }
</script>

<style scoped>
</style>