<script setup>
  //Importing the navbar component 
  import NavBar from '../components/NavBar.vue'
</script>

<template>
  <NavBar></NavBar>
  <h2 class="section-subtitle">Popular Comics</h2>
    
  <div class="" id="popular">
    <div class="container comics-section">
      <div class="row">
        <div class="col-xs-12 col-sm-12 col-lg-4">
          <div class="comic-card">
          <img class="rounded mx-auto d-block img-thumbnail comic-thumbnail" src="https://media.discordapp.net/attachments/231829187776741377/1271109748997099560/IMG_7191.png?ex=66b624b4&is=66b4d334&hm=e1b3e386a245c7d02e37f1b1bc93e61aae230204df060ece9607708c5ae7c0c1&=&format=webp&quality=lossless&width=354&height=350" alt="">
          <p class="comic-title">As vozes capitulo 100</p>          
          </div>
        </div>
    
        <div class="col-xs-12 col-sm-12 col-lg-4">
          <div class="comic-card">
          <img class="rounded mx-auto d-block img-thumbnail comic-thumbnail" src="https://cdn.discordapp.com/attachments/231829187776741377/1270820527304343653/tq9jch4n6g5b1.png?ex=66b51758&is=66b3c5d8&hm=cdc96b322063ca97ca586b1f2efecf2db195738abeeb93a5cd2a941f76229d4c&" alt="">
          <p class="comic-title">A lenda do dlagao neglo e a volta do lei lucio da silva capitulo 400</p>          
          </div>
        </div>
    
        <div class="col-xs-12 col-sm-12 col-lg-4">
          <div class="comic-card">
          <img class="rounded mx-auto d-block img-thumbnail comic-thumbnail" src="https://media.discordapp.net/attachments/231829187776741377/1271065332550729783/20240808_081750.jpg?ex=66b5fb56&is=66b4a9d6&hm=eb5834d099e900d2b1f3416c5f933cec00ca0893db6d6c966db618a1e64cf06c&=&format=webp&width=285&height=452" alt="">
          <p class="comic-title">Twitter da Silva cap 10</p>          
          </div>
        </div>
      </div>
    </div>
  </div>
  <hr class="section-line">  
  <h2 class="section-subtitle"> Recently Added</h2>
  <div class="container" id="recent">
    <div class="row">
      <div class="col">
        <img class="rounded mx-auto d-block img-thumbnail comic-thumbnail" src="https://cdn.discordapp.com/attachments/1244083068017446932/1271267902389223454/img_or2805211934_0002.png?ex=66b6b7fe&is=66b5667e&hm=518f0100c8435c35df8cd211f7ba5f982e74f427ea1d40e54fc266e974d234f1&" alt="">
        <p class="comic-title">Kaiju n 8 eu acho</p>
      </div>
      <div class="col">
        <img class="rounded mx-auto d-block img-thumbnail comic-thumbnail" src="https://cdn.discordapp.com/attachments/1244083068017446932/1271267902389223454/img_or2805211934_0002.png?ex=66b6b7fe&is=66b5667e&hm=518f0100c8435c35df8cd211f7ba5f982e74f427ea1d40e54fc266e974d234f1&" alt="">
        <p class="comic-title">Kaiju n 8 eu acho</p>
      </div>
      <div class="col">
        <img class="rounded mx-auto d-block img-thumbnail comic-thumbnail" src="https://cdn.discordapp.com/attachments/1244083068017446932/1271267902389223454/img_or2805211934_0002.png?ex=66b6b7fe&is=66b5667e&hm=518f0100c8435c35df8cd211f7ba5f982e74f427ea1d40e54fc266e974d234f1&" alt="">
        <p class="comic-title">Kaiju n 8 eu acho</p>
      </div>    
    </div>
  </div>
</template>
