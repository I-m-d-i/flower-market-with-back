<script>
import CardProduct2 from "@/components/CardProduct2.vue";

export default {
  name: "CardPage",
  components: {CardProduct2},
  data() {
    return {
      id: this.$route.params.id
    }
  },
  computed: {
    product() {
      return this.$store.getters.getProductById(this.id);
    }
  },
  methods: {
    addToBasket(id) {
      this.$store.commit("addToBasket", id)
    },
  },
  created() {
    window.scrollTo(0, 0)
  }
}
</script>

<template>
  <div>
    <div>
      <div class="icon-back" @click="this.$router.go(-1)">
        <svg xmlns="http://www.w3.org/2000/svg" width="70" height="24" viewBox="0 0 73 24" fill="none">
          <path
            d="M0.939339 10.9393C0.353554 11.5251 0.353554 12.4749 0.939339 13.0607L10.4853 22.6066C11.0711 23.1924 12.0208 23.1924 12.6066 22.6066C13.1924 22.0208 13.1924 21.0711 12.6066 20.4853L4.12132 12L12.6066 3.51472C13.1924 2.92893 13.1924 1.97919 12.6066 1.3934C12.0208 0.807611 11.0711 0.807611 10.4853 1.3934L0.939339 10.9393ZM73 10.5L2 10.5V13.5L73 13.5V10.5Z"
            fill="#81637D"/>
        </svg>
      </div>
      <div class="card-container">
        <div>
          <h1 class="header header-left">КАРТОЧКА БУКЕТА</h1>
          <img :src="product.img_src" alt="">
        </div>
        <div>
          <div class="header header-right"/>
          <div class="card-info">
            <h2> {{ product.title }}</h2>
            <p>{{ product.description }}</p>
            <div class="product-parameters">
              <p>Диаметр: {{ product.diameter }}</p>
              <p>Высота: {{ product.height }}</p>
            </div>
            <p class="price">{{ product.price }}<span> руб</span></p>
            <div>
              <input type="checkbox" id="cb1"><label for="cb1">Добавить открытку</label>
            </div>
            <div class="text">
              <textarea id="text" placeholder="напишите текст"/>
            </div>
            <div style="text-align: end">
              <button class="div-8" @click="addToBasket(product.id)">Купить</button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="other-products">
      <h2>вам так же может понравится:</h2>
      <div class="other-products-container">
        <card-product2 id="30"></card-product2>
        <card-product2 id="31"></card-product2>
        <card-product2 id="34"></card-product2>
      </div>
    </div>
  </div>
</template>

<style scoped>

.card-info{
  padding-inline: 15px;
}
.other-products-container{
  margin-inline: 100px;
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  flex-direction: row;
  justify-content: center;
  justify-items: center;
  gap: 10%;
}

.other-products h2{
  color: #000;
  font-family: Source Serif Pro;
  margin-left: 110px;
  font-size: 24px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  font-variant: small-caps;
}
.icon-back {
  position: absolute;
  top: 89px;
  left: 2%;
  cursor: pointer;
  overflow: hidden;
}

.text {
  margin-top: 40px;
}

textarea {
  width: 100%;
  border-radius: 26px;
  border: 2px solid #846680;
  background: linear-gradient(0deg, #FFF 0%, #FFF 100%), linear-gradient(0deg, #D5CBCB 0%, #D5CBCB 100%), linear-gradient(0deg, #D5CBCB 0%, #D5CBCB 100%), #D5CBCB;
  resize: none;
  height: 149px;
  padding: 10px;
}

textarea::placeholder {
  padding-left: 10px;
  color: #948C8C;
  font-family: Playfair Display;
  font-size: 16px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  font-variant: small-caps;
}

textarea:focus {
  border: 2px solid #846680;
  opacity: 0.6;
  outline: none;
}

textarea:focus::placeholder {
  color: #FFFFFF;
}

p {
  color: #000;
  font-family: Playfair Display;
  font-size: 16px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  font-variant: small-caps;
}

.price {
  color: #000;
  font-family: Crimson Text;
  font-size: 24px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  font-variant: small-caps;
  margin-bottom: 30px;
}

.price span {
  color: #000;
  font-family: Playfair Display;
  font-size: 24px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  font-variant: small-caps;
}

h2 {
  color: #000;
  font-family: Playfair Display;
  font-size: 32px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  font-variant: small-caps;
  margin-bottom: 20px;
}

.card-container {
  padding-inline: 20px;
  padding-bottom: 20px;
  display: grid;
  grid-template-columns: 1fr 1fr;
  justify-items: stretch;
  max-width: 80%;
  margin-inline: auto;
}

.product-parameters {
  margin-block: 42px;
}

img {
  border-radius: 368px 368px 0px 0px;
  aspect-ratio: 0.84;
  object-position: center;
  position: relative;
  height: 40%;
  overflow: hidden;
  align-self: start;
  background: lightgray 50% / cover no-repeat;
}

label {
  color: #000;
  font-family: Playfair Display;
  font-size: 20px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  font-variant: small-caps;
}

.div-8 {
  color: #fffbfb;
  align-self: stretch;
  border-radius: 40px;
  background-color: #81637d;
  margin-top: 26px;
  flex-grow: 1;
  padding: 12px 25px 12px 25px;
  font: small-caps 400 18px Playfair Display, -apple-system, Roboto, Helvetica, sans-serif;
}

.header {
  display: flex;
  flex-direction: row;
  margin-top: 50px;
  margin-bottom: 30px;
  height: 92px;
  background-color: #e7e0e7;
}

.header-left   {
  padding-block: 25px;
  padding-left: 35px;
  color: #000;
  font-family: Source Serif Pro;
  font-size: 36px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  font-variant: small-caps;
}
@media (max-width: 1000px) {
  .header-left{
    font-size: 28px;
  }
}
@media (max-width: 800px) {
  .other-products-container{
    grid-template-columns: repeat(1, 1fr);
    gap: 10px;
    margin-inline:50px;
  }
  .icon-back {
    width: 60px;
  }
}

@media (max-width: 650px) {
  .header-title h1 {
    font-size: 24px;
  }
  .card-container{
    grid-template-columns: 1fr;
  }

  .header-right {
    display: none;
  }
  .icon-back {
    width: 50px;
     }

  img {
    width: 100%;
    height: auto;
  }

  .product-parameters{
    margin-block: 15px;
  }
  .other-products h2{
    margin-left: 25px;
  }
}


input[type="checkbox"]:checked,
input[type="checkbox"]:not(:checked) {
  position: absolute;
  left: -9999px;
}

input[type="checkbox"]:checked + label,
input[type="checkbox"]:not(:checked) + label {
  display: inline-block;
  position: relative;
  padding-left: 28px;
  line-height: 20px;
  cursor: pointer;
  user-select: none;
}

input[type="checkbox"]:checked + label:before,
input[type="checkbox"]:not(:checked) + label:before {
  content: "";
  position: absolute;
  left: 0px;
  top: 0px;
  width: 18px;
  height: 18px;
  background-color: #A795A5;
}

input[type="checkbox"]:checked + label:before,
input[type="checkbox"]:not(:checked) + label:before {
  border-radius: 100%;
}

input[type="checkbox"]:checked + label:after,
input[type="checkbox"]:not(:checked) + label:after {
  content: "";
  position: absolute;
}

input[type="checkbox"]:checked + label:after,
input[type="checkbox"]:not(:checked) + label:after {
  left: 0px;
  top: 0px;
  width: 18px;
  height: 18px;
  border-radius: 100%;
  background-color: #81637D;
}

input[type="checkbox"]:not(:checked) + label:after {
  opacity: 0;
}

input[type="checkbox"]:checked + label:after {
  opacity: 1;
}

</style>