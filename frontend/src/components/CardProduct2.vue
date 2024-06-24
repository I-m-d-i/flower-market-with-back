<script>
export default {
  name: "CardProduct2",
  props: {
    id:{
      type: String,
      required: true,
    },
  },
  computed: {
    product() {
      return this.$store.getters.getProductById(this.id);
    }
  }
}
</script>

<template>
<div class="card-product2">
  <a class="img" :href="'/#/product/'+id">
    <img :src="product.img_src" class="img" alt="img"/>
  </a>
  <div class="div-3">
    <div class="div-4">{{ product.product_name }}</div>
    <button class="add-to-basket">{{product.price}}</button>
  </div>
</div>
</template>

<style scoped>
.div-3{
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 17px;
  gap: 10px;
}
.div-4 {
  color: #000;
  font-family: Playfair Display;
  font-size: 16px;
  line-height: normal;
  font-variant: small-caps;
  margin-left: 10px;
}
.card-product2 {
  max-width: 350px;
  padding-bottom: 50px;
}
.img {
  width: 100%;
  aspect-ratio: 0.9;
  min-width: 200px;
  border-radius: 24px;
}
.add-to-basket {
  max-width: 110px;
  color: #fffbfb;
  align-self: stretch;
  border-radius: 40px;
  background-color: #81637d;
  flex-grow: 1;
  padding: 12px 25px 12px 25px;
  font: small-caps 400 20px Playfair Display, -apple-system, Roboto, Helvetica, sans-serif;
}
@media (max-width: 500px) {
  .add-to-basket{
    font-size: 16px;
    padding: 7px 25px 7px 25px;
  }
}
</style>