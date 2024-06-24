<script>
import CardProduct from "@/components/CardProduct.vue";

export default {
  name: "popular",
  components: {CardProduct},
  computed: {
    products() {
      return this.$store.getters.getProducts.slice(0,3)
    }
  }
}
</script>

<template>
  <div id="popular" style="text-align: left">
    <span class="popular_title">Популярное</span>
  </div>
  <div class="popular">
    <card-product class="card-product" v-for="product in products" :key="product.product_id" :id="product.product_id" :title="product.product_name" :price="product.price" :description="product.description" :image_src="product.img_src"/>
  </div>
</template>

<style scoped>
.popular {
  display: flex;
  flex-wrap: wrap;
  margin-top: 80px;
  justify-content: space-around;
  margin-inline: 20px;
  gap: 5%;
  justify-items: center
}
@media (max-width: 1100px) {
  .card-product{
    margin-bottom: 20px;
  }
}

.popular_title{
  font-family: Source Serif Pro;
  font-size: 40px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  font-variant: small-caps;
  margin-left: 33%
}

@media (max-width: 680px) {
  .popular_title{
    margin-left: 20px
  }
}
</style>