<script>
import CardProductBasket from "@/components/CardProductBasket.vue";
import HeaderComponent from "@/components/Header.vue";

export default {
  name: "Basket",
  components: {HeaderComponent, CardProductBasket},
  computed: {
    productsInBasket() {
      return this.$store.getters.getBasket;
    },
    collectionProducts() {
      return [...new Set(this.productsInBasket)];
    }
  },
  methods: {
    countProductInBasket(id) {
      return this.productsInBasket.filter(item => item === id).length
    }
  }
}
</script>

<template>
  <header-component/>
  <div style="margin-top: 73px">
    <h1>Корзина</h1>
    <div class="basket">
      <template v-if="collectionProducts.length===0">
        корзина пуста
      </template>
      <template v-else>
        <card-product-basket v-for="id in collectionProducts" :key="id" :id="id"
                             :count="countProductInBasket(id)"></card-product-basket>
      </template>
    </div>
  </div>
</template>

<style scoped>
h1 {
  color: #000;
  font-family: Source Serif Pro;
  font-size: 48px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  font-variant: small-caps;
  margin-left: 74px;
  margin-bottom: 60px;
}

.basket {
  display: grid;
  justify-items: center;
  gap: 50px;
}

@media (max-width: 800px) {
  h1 {
    font-size: 38px;
  }
}
</style>