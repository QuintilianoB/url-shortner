<template>
    <div>
        <h1>Shortened URLs</h1>
        <p>Only URLs with protocol are accepted.</p>
        <input type="text" v-model="url"/>
        <button @click="submit()">Submit</button>
        <h3>Result</h3>
        <a v-bind:href=url_short>{{url_short}}</a>
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
                if (this.validURL(this.url)) {
                    let data = {url: this.url};
                    this.axios.post('/create', data, {
                        headers: {
                            "Content-Type": "application/json"
                        },
                    }).then(response => {
                        this.url_short = response.data.url_short
                    }).catch(error => {
                        this.errors.push(error)
                    })
                } else {
                    alert("Invalid URL.")
                }
            },
            validURL(myURL) {
                let pattern = /(ftp|http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-/]))?/;
                return pattern.test(myURL)
            },
        }
    }
</script>

<style scoped>
</style>