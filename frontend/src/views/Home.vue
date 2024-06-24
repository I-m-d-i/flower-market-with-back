<script>
import Popular from "@/views/popular.vue";
import Catalog from "@/views/Catalog.vue";
import HeaderComponent from "@/components/Header.vue";
import About from "@/views/About.vue";
import Reviews from "@/views/Reviews.vue";
import MyFooter from "@/components/Footer.vue";
import AboutMobile from "@/views/AboutMobile.vue";
import axios from "axios";
import bus from "../plugins/eventBus"

export default {
  name: "Home-component",
  components: {AboutMobile, Popular,Catalog, About,  MyFooter,  HeaderComponent, Reviews},
  methods: {
    scrollToCatalog() {
      document.getElementById("catalog").scrollIntoView({behavior: "smooth"});
    }
  },
  computed: {
    width() {
      return window.innerWidth
    }
  },
  async mounted() {
    axios.get('/api/session/check',{withCredentials: true}).then(()=>{
      if (!this.$store.getters.isValidUser){
        this.$router.push({name: "AuthPage"})
      }
    }).catch(()=>{
      setTimeout(() => {
        bus.$emit("openRegistration", this.$router.currentRoute.value)
      },1500)
    })
  }
}
</script>

<template>
  <header-component/>
<!--  <v-dialog v-model="registerWindow" max-width="400">-->
<!--    <registration-component @close="registerWindow = false"/>-->
<!--  </v-dialog>-->
  <div class="home">
    <div class="home-text">
      <div >
        <h1>ТВОЙ БУКЕТ<br/>ВСЁ СКАЖЕТ ЗА ТЕБЯ</h1>
        <h2>РАДУЙТЕ СВОИХ БЛИЗКИХ</h2>
        <div style="cursor: pointer" class="catalog-button">
          <div  class="catalog-text" @click="scrollToCatalog()">в каталог</div>
        </div>
      </div>
      <img style="width: 60vw;
    z-index: -1;
    position: absolute;
    right: 0;
    top: 0;
    max-width: 600px;" alt="flowers-home"
           :src="require('@/assets/flowers-home.png')"/>
    </div>

    <div class="guarantee">
      <div class="guarantee-h">
              <span>
        Гарантия качества
      </span>
      </div>
      <div class="guarantee-text">
           <span>
        КАЖДЫЕ 3 ДНЯ У НАС<br/> ПОСТАВКА САМЫХ СВЕЖИХ<br/> ЦВЕТОВ
      </span>
      </div>
    </div>
  </div>
  <div class="subscribe-form">
    <h2 class="subscribe-title">Подпишись и получи
      <svg xmlns="http://www.w3.org/2000/svg" width="135" height="115"
           viewBox="0 0 135 115" fill="none">
        <ellipse cx="67.5628" cy="57.4209" rx="73.127" ry="48.1633" transform="rotate(-33.3285 67.5628 57.4209)"
                 fill="white"/>
        <text x="28" y="70" font-size="55" font-family="Crimson Text" fill="#81637d" font-weight="normal"
              pointer-events="none">10%
        </text>
      </svg>
      скидку на первый заказ
    </h2>
    <div class="subscribe-form-wrapper">
      <input type="email" class="label" id="email" placeholder="E-mail"/>
      <button class="button">Подписаться</button>
    </div>
  </div>
  <div style="padding-block: 80px;">
    <popular></popular>
  </div>
  <div class="sale-form">
      <h2 class="sale-title">Закажите букет на<br/>
        день рождения со скидкой </h2>
        <button class="button">Подробнее</button>


    <svg class="percent" xmlns="http://www.w3.org/2000/svg" width="135" height="115"
         viewBox="0 0 135 115" fill="none">
      <ellipse cx="67.5628" cy="57.4209" rx="73.127" ry="48.1633" transform="rotate(-33.3285 67.5628 57.4209)"
               fill="white"/>
      <text x="35" y="70" font-size="40" font-family="Crimson Text" fill="#81637d" font-weight="normal"
            pointer-events="none">15%
      </text>
    </svg>
  </div>
 <catalog/>
  <about v-if="width > 950"/>
  <about-mobile v-else/>
  <reviews/>
  <my-footer/>
</template>

<style scoped>


.home {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  width: 100%;
}

.home-text{
  margin-left: 90px;
}

h1 {
  font-family: source serif pro, EXTRA-LIGHT;
  font-weight: 400;
  line-height: normal;
  font-style: normal;
  font-variant: small-caps;
  text-align: start;
  font-size: 55px;
  color: #4F3960;
  text-indent: 30px;
}

h1:first-letter {
  font-size: 65px;
}

.home h2 {
  font-family: source serif pro, EXTRA-LIGHT;
  font-size: 50px;
  color: #4F3960;
  font-weight: 400;
  font-style: normal;
  text-align: start;
  line-height: normal;
  font-variant: small-caps;
}

@media (max-width: 1800px) {
  .home h2 {
    font-size: 36px;
  }
  h1 {
    font-size: 48px;
  }
  h1:first-letter {
    font-size: 60px;
  }
}

.catalog-button {
  margin: 15% 0 0 290px;
  width: 405px;
  border-radius: 64px;
  background-color: #81637d;
  display: flex;
  flex-direction: column;
  height: 42px;
}

.catalog-text {
  color: #fff;
  align-self: center;
  max-width: 278px;
  font: small-caps 400 30px Source Serif Pro, -apple-system, Roboto, Helvetica, sans-serif;
}

@media (max-width: 991px) {
  .catalog-button {
    margin: 15% 0 0 0%;
  }
  .subscribe-title {
    flex-direction: column;
  }
  .subscribe-title {
    max-width: 100%;
  }
  .img {
    max-width: 100%;
  }
  h1 {
    font-size: 40px;
  }
  h1:first-letter {
    font-size: 50px;
  }
  .home-text{
    margin-left: 20px !important;
  }
  .catalog-button{
    width: 300px;
  }
}

.guarantee {
  margin-left: 110px;
  margin-top: 24px;
}

.guarantee .guarantee-h {
  font-family: playfair display, serif;
  font-size: 24px;
  font-style: normal;
  color: #000;
  font-weight: 400;
  text-align: start;
  line-height: normal;
  font-variant: small-caps;
  padding-bottom: 10px;
}

input {
  padding: 10px;
  border: none;
  border-radius: 4px;
  margin-bottom: 10px;
}

.guarantee .guarantee-text {
  font-family: playfair display, serif;
  font-size: 16px;
  font-style: normal;
  color: #000;
  font-weight: 400;
  text-align: start !important;
  line-height: normal;
  font-variant: small-caps;
}

.subscribe-form {
  background-color: #81637d;
  display: flex;
  margin-top: 60px;
  flex-direction: column;
  align-self: center;
  width: 100%;
  padding-bottom: 15px;
}

.subscribe-title {
  display: flex;
  gap: 20px;
  align-items: center;
  justify-content: center;
  color: #fff;
  text-align: center;
  font: small-caps 400 32px Source Serif Pro, -apple-system, Roboto, Helvetica, sans-serif;
  margin: 5px 0px 0px 0px;
}


.subscribe-form-wrapper {
  align-self: center;
  display: flex;
  flex-wrap: wrap;
  margin-top: 15px;
  width: 100%;
  align-items: flex-start;
  justify-content: space-around;
  gap: 20px;
}

.label {
  align-self: start;
  max-width: 300px;
  border-radius: 40px;
  background-color: rgba(214, 176, 215, 0.71);
  flex: 1;
  font: 400 22px Inter, sans-serif;
  outline: none;
  padding: 10px 0 10px 50px;
}

input {
  border-style: none;
  outline: none;
}

.label::placeholder {
  color: #ffffff;
}

.subscribe-form .button {
  color: #81637d;
  align-self: start;
  max-width: 300px;
  border-radius: 40px;
  background-color: rgba(217, 217, 217, 0.76);
  flex: 1;
  padding: 12px 80px 12px 34px;
  font: small-caps 400 22px Playfair Display;
  border-style: none;
  outline: none;
  margin-bottom: 10px;
}

.sale-form .button {
  border-radius: 40px;
  background: rgba(217, 217, 217, 0.75);
  font: small-caps 400 22px Playfair Display;
  grid-column-start: 2;
  grid-row-start: 3;
  display: grid;
  outline: none;
  margin-bottom: 10px;
  color: #ffffff;
  width: 350px;
  height: 38px;
  flex-shrink: 0;
}

.percent{
  grid-row-start: 2;
  justify-self: flex-start;
  grid-column-start: 3;
}

.sale-form {
  display: grid;
  justify-items: center;
  grid-template-columns: 1fr 1fr 1fr;
  grid-template-rows: 1fr 20px 1fr;
  align-items: center;
  background-image: url(../assets/sale.png);
  background-position: center;
  background-repeat: repeat;
  height: 306px;
}



.sale-title {
  text-align: center;
  color: #FFF3F3;
  grid-column-start: 2;
  font-family: Playfair Display;
  font-size: 38px;
  font-style: normal;
  font-weight: 600;
  line-height: normal;
}

@media (max-width: 1100px) {
  .percent {
    grid-column-start: 2;
    justify-self: center;
  }
  .sale-form {
    grid-template-rows: 1fr 1fr 1fr;
    gap:5px
  }
  .sale-title{
    font-size: 30px;
  }
}

</style>