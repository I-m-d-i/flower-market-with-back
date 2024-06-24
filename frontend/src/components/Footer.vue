<script>
export default {
  name: "Footer",
  methods: {
    scrollTo(id) {
      this.$router.push({name: "home", hash: `#${id}`})
    }
  }
}
</script>
<template>
<footer>
  <div id="footer" class="footer-container" >
    <div class="footer-logo" style=" position: relative;
       margin-right: 30px;
   overflow:hidden;
    width: 200px;
        height: 150px;">
      <img class="logo" alt="logo" :src="require('@/assets/logo.png')" style="position: absolute;
   top:50%;
   left:50%;
   transform:translate(-50%,-50%);
    width: 200px;
    height: 250px;
   object-fit:cover;">
    </div>
    <div class="footer-menu">
      <ul>
        <li><a @click="scrollTo('popular')">Популярное</a></li>
        <li><a @click="scrollTo('catalog')">Каталог</a></li>
        <li><a @click="scrollTo('about')">О нас</a></li>
        <li><a @click="scrollTo('reviews')">Отзывы</a></li>
        <li><a @click="scrollTo('footer')">Контакты</a></li>
      </ul>
  </div>
    <div class="footer-contact">
      <ul>
        <li>Адрес: город Красноярск, ул.Свердловская 15б.</li>
        <li>часы работы: с 9:00 - 21:00</li>
        <li>Телефон: 89082192323</li>
      </ul>
    </div>
  </div>
</footer>
</template>

<style scoped>
footer {
  padding-bottom: 200px;
}
li {
  list-style-type: none;
}
.footer-container{
  background-color: #81637D;
  display: flex;
  flex-flow: row wrap;
  justify-content: space-around;
  align-items: center;
  padding-block: 80px;
  gap: 20px;
}
.footer-menu li {
  mix-blend-mode: soft-light;
  margin-bottom: 30px;
  color: #FFF;
  font-family: Anonymous Pro;
  font-size: 36px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  margin-inline: 10px;
}
.footer-contact li {
  margin-bottom: 1em;
  color: #FFF;
  font-family: Anonymous Pro;
  font-size: 36px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  margin-inline: 10px;
}
.footer-logo {
  align-self: flex-start;
  width:175px;
  height:163px;
  background-color: #620255
}
</style>