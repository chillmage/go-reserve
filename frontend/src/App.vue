<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Generate a thumbnail of a website</h1>

        <form v-on:submit.prevent="makeWebsiteThumbnail">
          <div class="form-group">
            <input v-model="websiteUrl" type="text" id="website-input" placeholder="Enter a website" class="form-control">
          </div>
          <div class="form-group">
            <button class="btn btn-primary">Generate!</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>


<script>
import axios from "axios"

export default {
  name: 'App',

  data() { 
    return {
      websiteUrl: '',
    } 
  },

  methods: {
    makeWebsiteThumbnail() {
      axios.post("https://screenshotapi.net/api/v1/screenshot", {
        token: "SCREENSHOTAPI_TOKEN",
        url: this.websiteUrl,
        width: 1920,
        height: 1080,
        output: 'json',
        thumbnail_width: 300
      })
      .then((response) => {
        this.thumbnailUrl = response.data.screenshot;
      })
      .catch((error) => {
        window.alert(`The API returned an error: ${error}`);
      })    
    }
  }
}
</script>