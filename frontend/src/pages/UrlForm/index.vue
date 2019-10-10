<template>
    <div>
        <input type="text" v-model="url"/>
        <button @click="submit()">Submit</button>
        <p>{{url_short}}</p>
    </div>
</template>

<script>

    export default {
        name: 'app',
        data() {
            return {
                url: null,
                url_short: null,
            }
        },
        created() {
        },
        methods: {
            submit() {
                let data = {url: this.url}
                this.$axios.post('/create', data, {
                    headers: {
                        "Content-Type": "application/json"
                    },
                }).then(response => {
                    this.url_short = response.data.url_short
                }).catch(error => {
                    this.errors.push(error)
                })
            },
            validURL(myURL) {
                let pattern = new RegExp('^(https?:\\/\\/)?' + // protocol
                    '((([a-z\\d]([a-z\\d-][a-z\\d]))\\.?)+[a-z]{2,}|' + // domain name
                    '((\\d{1,3}\\.){3}\\d{1,3}))' + // ip (v4) address
                    '(\\:\\d+)?(\\/[-a-z\\d%_.~+])' + //port
                    '(\\?[;&amp;a-z\\d%_.~+=-]*)?' + // query string
                    '(\\#[-a-z\\d_]*)?$', 'i');
                return pattern.test(myURL);
            },
        }
    }
</script>

<style scoped>
</style>