<template>
  <div id="app" class="container">

    <!-- Stage it like a show. What is the sequence of events? -->
    <stanza>
      <h1 class="branding">kaitak racing</h1> 
      <p class="branding"><em>from 2022</em></p>  

    </stanza>

    <nav>
      <a class="">equipment/reservations</a>
      <a class="">league/events</a>
    </nav>
    <ol>
      <li class="kit">
        <header>
          <h3>ignition kit</h3>
        </header>
        <p>can be configured for formula racing or gt racing</p>
        <p>gt-lite cockpit</p>
        <p>thrustmaster tmx wheel</p>
        <p>thrustmaster tx 300 alumininum pedals</p>

        <h4>Q & A</h4> 
        <h5>will this fit?</h5>
      
      </li>
    </ol>
    <!-- This should be  a horizontal scrolling element. -->
    <nav>
      <article>
        <h2></h2>
      </article>
    </nav>
    <footer>
      <h2>shift up</h2>
    </footer>
  </div>
</template>



<script>
export default {
  name: 'App',

  data() { 
    return {
      selection: 1,
    } 
  },

  methods: {
    makeWebsiteThumbnail() {

    }
  }
}
</script>

<style>

* {
  box-sizing: border-box;
  font-family: "SF Mono", sans-serif;
  margin-top: 5vh;
}



</style>


