<template>
  <div id="app" class="container">

    <!-- Stage it like a show. What is the sequence of events? -->
    <header class="branding">
      <img src="../public/logo.svg"/>
      <h1>Kaitak Racing</h1> 
    </header>

    <nav>
      <a class="site-links"><h2>Equipment/Reservations</h2></a>
      <p>Free scheduled delivery, installation and retrieval.</p>
    </nav>
  
    <article class="kit">
      <header>
        <h3>Ignition Kit</h3>
        <p><em>(Xbox and PC)</em></p>
        <p>$200 per day</p>
        <p>$500 per day</p>
      </header>
      <main>
        <p>Ignition Kit includes:</p>
        <ul>
          <li>
            <img class="product" src="https://pisces.bbystatic.com/image2/BestBuy_US/images/products/5434/5434000_sd.jpg"/>
            <p>Thrustmaster TMX Wheel</p>
          </li>
          <li>
            <img class="product" src="https://www.beracer.com/shop/pub/media/catalog/product/cache/2a6b0744b87cbe1990f7a65c1fd3659e/t/h/thrustmaster-tx-ferrari-458italia-edition-2.jpg"/>
            <p>Thrustmaster TX Alumininum Pedalset</p>
          </li>
          <li>
            <img class="product" src="https://3pgkb2jl8pg2smexondiy9e4-wpengine.netdna-ssl.com/wp-content/uploads/2020/10/gt_lite_website_gallery_3-1.png" alt="Cockpit Illustration"/>
            <p>Next Level Racing Folding Sim Cockpit</p>
          </li>
        </ul>
      </main>
    </article>

    <article id="faq">
      <h3> Q & A </h3>
    </article>

    <footer>
      <p><em>Full send from 2022</em></p>  
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

    },
    pullInteview() {

    }

  }
}
</script>

<style>

html {
  font-size: 16px;
}

* {
  box-sizing: border-box;
  font-family: "Overpass Mono", sans-serif;
  margin-top: 5vh;
}

h1, h2, h3 {
  font-weight: 900;
}

h1 {
  font-kerning: -50px;
  font-size: 4em;
}

.branding img {
  width: 150px;
  margin-left: -20px;
  margin-top: 50px;
}

nav a, nav a h2 {
  margin-right: 10px;
  font-weight: 700;
}

.kit img {
  height: 200px;
}

li {
  margin: 20px;
}

ul {
  display: flex;
  list-style: none;
}

</style>


