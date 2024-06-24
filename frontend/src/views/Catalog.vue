<script>
import CardProduct from "@/components/CardProduct.vue";

export default {
  name: "Catalog",
  components: {CardProduct},
  data() {
    return {
      page: 0,
      // itemsPerPage: 6,
    }
  },
  methods: {
    showMenu() {
      document.getElementById("catalog-menu").classList.toggle("active");
    }
  },
  computed: {
    itemsPerPage() {
      if (this.isMobileCarousel) {
        return 4
      } else {
        return 6
      }
    },
    displayData() {
      //разбитый по страницам массив
      let listingProducts = []

      // Рассчитываем начальный и конечный индексы элементов для текущей страницы
      for (let i = 1; i <= Math.ceil(this.$store.getters.getProducts.length / this.itemsPerPage); i++) {
        let startIndex = (i - 1) * this.itemsPerPage;
        let endIndex = startIndex + this.itemsPerPage;
        listingProducts.push(this.$store.getters.getProducts.slice(startIndex, endIndex))
      }
      return listingProducts
    },
    isMobileMenuCatalog() {
      return window.innerWidth <= 1000
    },
    isMobileCarousel() {
      return window.innerWidth <= 660
    }
  },
}
</script>

<template>
  <div id="catalog" style="">
    <h1>Каталог</h1>
  </div>
  <div class="catalog-container">
    <div class="catalog-menu-container">
      <v-icon @click="showMenu()" v-if="isMobileMenuCatalog" class="mobile_menu_icon">mdi-filter-menu</v-icon>
      <div id="catalog-menu" class="catalog-menu" >
        <h2 class="title">Каталог</h2>
        <ul class="navigation">
          <li><a href="#" class="link">Весенние букеты</a></li>
          <li><a href="#" class="link">Свадебные букеты</a></li>
          <li><a href="#" class="link">Дизайнерские букеты</a></li>
          <li><a href="#" class="link">Монобукеты</a></li>
          <li><a href="#" class="link">Цветы в коробке</a></li>
          <li><a href="#" class="link">Корзина с цветами</a></li>
        </ul>
        <div class="filter">
          <h3 class="filter-title">Фильтр</h3>
          <div class="filter-item">
            <h4 class="filter-subtitle">Размер</h4>
            <ul>
              <li><label class="filter-label">
                <input type="radio" name="size" class="filter-input"/>Большой
              </label></li>
              <li><label class="filter-label">
                <input type="radio" name="size" class="filter-input"/>Средний
              </label></li>
              <li><label class="filter-label">
                <input type="radio" name="size" class="filter-input"/>Маленький
              </label></li>
            </ul>
          </div>
          <div class="filter-item">
            <h4 class="filter-subtitle">Цвет</h4>
            <div class="color-picker">
              <div class="color-option">
                <svg xmlns="http://www.w3.org/2000/svg" width="14" height="15" viewBox="0 0 14 15" fill="none">
                  <ellipse cx="7" cy="7.5" rx="7" ry="7.5" fill="#D9D9D9"/>
                </svg>
              </div>
              <div class="color-option">
                <svg xmlns="http://www.w3.org/2000/svg" width="14" height="15" viewBox="0 0 14 15" fill="none">
                  <ellipse cx="7" cy="7.5" rx="7" ry="7.5" fill="#785C75"/>
                </svg>
              </div>
              <div class="color-option">
                <svg xmlns="http://www.w3.org/2000/svg" width="14" height="15" viewBox="0 0 14 15" fill="none">
                  <ellipse cx="7" cy="7.5" rx="7" ry="7.5" fill="#F5A3C5"/>
                </svg>
              </div>
              <div class="color-option">
                <svg xmlns="http://www.w3.org/2000/svg" width="14" height="15" viewBox="0 0 14 15" fill="none">
                  <ellipse cx="7" cy="7.5" rx="7" ry="7.5" fill="#ECA082"/>
                </svg>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="catalog-content">
      <div class="catalog-sort">
        <div class="catalog-sort-content">
          <span class="black-text">Сортировать по: </span>
          <span class="gray-text">лучшей цене</span>
        </div>
      </div>
      <div class="catalog-products">
        <v-carousel  style="height: 100%" v-model="page" hide-delimiter-background
                    :show-arrows="false" color="#ab91a8">
          <v-carousel-item style="width: 100%"
            v-for="(products,i) in displayData"
            :key="i"
          >
            <div class="catalog-products-list">
              <card-product class="card-product" v-for="(product,i) in products"
                            :key="i"
                            :title="product.product_name"
                            :price="product.price"
                            :description="product.description"
                            :image_src="product.img_src"
                            :id="product.product_id"
              />
            </div>

          </v-carousel-item>
        </v-carousel>
      </div>
    </div>
  </div>
</template>

<style scoped>

.catalog-products >>> .v-btn--active {
  color: #785C75 !important;
}

.catalog-products >>> .v-btn--icon .v-icon {
  --v-icon-size-multiplier: 2;
}

.catalog-products >>> .v-btn--active > .v-btn__overlay, .v-btn[aria-haspopup=menu][aria-expanded=true] > .v-btn__overlay {
  opacity: 0;
}

.catalog-products >>> .v-carousel__controls {
  gap: 31px
}

.card-product {
  max-width: 315px;
}

h1 {
  color: #000;
  font-family: Source Serif Pro;
  font-size: 48px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  font-variant: small-caps;
  text-align: left; margin-left: 20%; margin-block: 100px;
}

.catalog-container {
  display: grid;
  grid-template-columns: 2fr 5fr;
  padding-inline: 30px;
  gap: 20px;
}
.catalog-menu-container {
  justify-self: end;
}
.catalog-menu {
  background-color: rgba(255, 253, 253, 082);
  display: flex;
  flex-direction: column;
  padding: 32px 80px 77px 11px;
  height: 900px;
  border: #D9D9D947 solid 6px;
  max-width: 350px;
}

.catalog-products-list {
  width: 100%;
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
  padding-bottom: 75px;
  padding-top: 20px;
  justify-items: center;
}

.catalog-products {
  width: 100%;
}

.catalog-sort-content {
  background-color: #e7e0e7;
  padding-block: 20px;
  height: 100%;
  padding-left: 5px;
}

.catalog-content {
  display: grid;
  grid-template-rows: 90px 1fr;
}

@media (max-width: 1000px) {
  .catalog-menu {
    position: absolute;
    left: 0;
    z-index: 2;
    transform: translateX(-600px);
  }
  .active {
    transform: translateX(0px);
    transition: 0.5s;
  }
  h1 {
    font-size: 36px;
    margin-block: 10px;
  }
  .catalog-container{
    grid-template-columns: 1fr;
    margin-inline: 0px;
  }
  .catalog-sort-content {
    padding-block: 5px;
  }
  .catalog-sort-content span{
    font-size: 20px;
  }
  .catalog-content {
    grid-template-rows: 60px 1fr;
  }
  .catalog-menu-container {
    justify-self: start;
  }
}

@media (max-width: 991px) {
  .catalog-menu-container {
    padding: 0 20px;
  }
}

@media (max-width: 660px) {
  .catalog-container{
    grid-template-columns: 1fr;
    width: 100%;
    padding-inline: 2px;
    justify-items: center;
  }

  .catalog-products-list {
    grid-template-columns: repeat(2, 1fr);
    gap: 0px !important;
  }
}

.catalog-menu-container ul {
  padding-left: 30px;
}


.title {
  font-family: Source Serif Pro;
  font-size: 32px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  font-variant: small-caps;
  margin: 0 0 0 80px;
}

@media (max-width: 991px) {
  .title {
    margin-left: 10px;
    white-space: initial;
  }
}

.link {
  color: #000;
  text-decoration-line: underline;
  text-align: start;
  margin-top: 39px;
  width: 100%;
  font: 400 20px Inter, sans-serif;
}

.filter {
  color: #000;
}

.filter-title {
  color: #000;
  font-family: Source Serif Pro;
  font-size: 32px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  margin: 30px 0px 30px 80px;
}

.filter-subtitle {
  margin-left: 50px;
  text-align: left;
  font-family: Playfair Display;
  font-size: 20px;
  font-style: normal;
  font-weight: 600;
  line-height: normal;
}

.filter-label {
  color: #000;
  text-decoration-line: underline;
  align-self: start;
  width: 100%;
  font: 400 20px Inter, sans-serif;
}

.color-picker {
  display: flex;
  margin-top: 32px;
  margin-left: 40px;
  width: 100%;
  justify-content: flex-start;
  gap: 20px;
}

@media (max-width: 991px) {
  .filter {
    margin-top: 40px;
  }

  .filter-title {
    margin: 40px 0 0 10px;
  }

  .color-picker {
    justify-content: center;
    margin-top: 40px;
  }
}

li {
  line-height: 3;
  list-style: none;
  text-align: left;
  padding: 0;
}

.filter-input {
  appearance: none;
}

.catalog-sort {
  justify-content: center;
  color: rgba(0, 0, 0, 0.6);
  padding-inline: 10px;
  width: 100%;
  text-align: start;
  font: small-caps 400 36px Source Serif Pro, -apple-system, Roboto, Helvetica, sans-serif;
}

.black-text {
  color: rgba(0, 0, 0, 1);
}

.gray-text {
  color: rgba(0, 0, 0, 0.6);
  text-decoration-line: underline;
}
</style>