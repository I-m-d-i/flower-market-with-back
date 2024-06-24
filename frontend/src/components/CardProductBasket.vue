<script>
export default {
  name: "CardProductBasket",
  props: {
    id: {
      type: String,
      required: true,
    },
    count: {
      type: Number,
      required: true,
    }
  },
  computed: {
    product() {
      return this.$store.getters.getProductById(this.id);
    }
  },
  methods: {
    addToBasket(id) {
      this.$store.commit("addToBasket", id);
    },
    deleteFromBasket(id) {
      this.$store.commit("deleteFromBasket", id);
    }
  }
}
</script>

<template>
  <div class="card-product">
    <a class="img_container" :href="'/#/product/'+id">
      <img :src="product.img_src" class="img" :alt="product.title"/>
    </a>
    <div class="div-3">
      <div class="div-4">{{ product.title }}</div>
      <div class="div-5">
        <p>Цвет: <span> белоснежный</span></p>
        <p>Склад:<span> красноярск</span></p>
      </div>
      <div class="actions_container">
        <div class="count-product">
          <button class="decrement" @click="deleteFromBasket(id)">—</button>
          <span>{{ count }}</span>
          <button class="increment" @click="addToBasket(id)">+</button>
        </div>
        <button class="add-to-cart" @click="$router.push({ name: 'PaymentPage'})">Купить</button>
      </div>

    </div>
  </div>
</template>


<style scoped>


img {
  min-width: 210px;
  max-width: 400px;
  width: 100%;
  height: auto;
  border-radius: 25px;
}

.div-3 {
  width: 100%;
  padding-top: 54px;
}

.count-product button {
  color: white;
  border-radius: 3px;
  width: 40px;
  height: 21px;
  user-select: none;
}

.count-product .decrement {
  background: #D9D9D9;
}

.count-product .increment {
  background: #81637D;
}

.count-product {
  display: inline-grid;
  grid-template-columns: 1fr 1fr 1fr;
  max-width: 120px;
  gap: 5px;
  align-items: center;
  justify-content: center;
}

.count-product span {
  text-align: center;
  color: #000;
  font-family: Playfair Display;
  padding-bottom: 8px;
  font-size: 24px;
  line-height: normal;
  font-variant: small-caps;
}

.card-product {
  display: grid;
  grid-template-columns: 1fr 1fr;
  width: 100%;
  flex-direction: row;
  align-items: flex-start;
  gap: 20px;
  max-width: 1000px;
  padding: 20px;
}

.div-4 {
  padding-bottom: 25px;
  color: #000;
  font-family: Playfair Display;
  font-size: 32px;
  line-height: normal;
  font-variant: small-caps;
}

.div-5 {
  color: #757171;
  font-family: Playfair Display;
  padding-bottom: 31px;
  font-size: 20px;
  line-height: normal;
  font-variant: small-caps;
}

.div-5 span {
  color: #000;
  font-family: Playfair Display;
  font-size: 20px;
  line-height: normal;
  font-variant: small-caps;
}

.add-to-cart {
  color: #fffbfb;
  align-self: stretch;
  border-radius: 40px;
  background-color: #81637d;
  max-width: 140px;
  min-width: 90px;
  flex-grow: 1;
  padding: 12px 25px 12px 25px;
  font: small-caps 18px Playfair Display, -apple-system, Roboto, Helvetica, sans-serif;
}

.actions_container {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: space-between;
  padding-inline: 20px;
  gap: 10px
}

@media screen and (max-width: 760px) {
  .card-product{
    gap: 15px;
  }
  .div-3 {
    padding-top: 10px;
  }
  .div-4{
    font-size: 22px;
    padding-bottom: 10px;
  }
  .div-5 {
    font-size: 16px;
    padding-bottom: 10px !important;
  }

  .div-5 span {
    font-size: 16px;
  }
  img {
    min-width: 150px;
    max-width: 300px;
  }
  .add-to-cart{
    font-size: 10px;
    padding: 7px 25px 7px 25px;
  }
  .count-product button{
    width: 35px;
  }
}

@media screen and (max-width: 500px) {
  .decrement {
    height: 20px !important;
  }
  .increment {
    height: 20px !important;
  }
  .actions_container {
    gap: 5px;
    padding-inline: 10px;
  }
  .div-5, .div-5 span {
    padding-bottom: 5px !important;
    font-size: 16px;
  }
}
</style>