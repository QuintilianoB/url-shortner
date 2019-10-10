<template>
    <div>
        <h1>Redirecting to {{url}}</h1>
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
            }
        },
        created() {
            this.page = this.$route.params.page;
            this.axios.get(this.page + "/stats")
                .then(response => {
                    this.url = response.data.url;
                    this.url_short = response.data.url_short;
                    this.redirect()
                })
        },
        methods: {
            // https://stackoverflow.com/questions/951021/what-is-the-javascript-version-of-sleep
            async redirect(){
                // Wait 3 seconds
                await this.sleep(3000)
                window.location.href = this.url;
            },
            sleep(ms){
                return new Promise(resolve => setTimeout(resolve, ms));
            }
        }
    }
</script>

<style scoped>
</style>