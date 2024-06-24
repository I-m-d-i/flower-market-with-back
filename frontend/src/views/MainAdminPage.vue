<script>
import GeneralPage from "@/views/AdminGeneralPage.vue";
import AdminProductsPage from "@/views/AdminProductsPage.vue";
import AdminAddProductPage from "@/views/AdminAddProductPage.vue";
import AdminPurchasesPage from "@/views/AdminPurchasesPage.vue";
export default {
  name: "MainAdminPage",
  components: {
    AdminPurchasesPage,
    AdminAddProductPage,
    AdminProductsPage,
    GeneralPage
  },
  data() {
    return {
      tab: "1"
    }
  },
  methods:{
    newTab(tab){
      console.log(tab)
      this.tab = tab
    }
  }
}
</script>

<template>
<div class="d-flex flex-row" style="height: 100vmax;width: 100%">
  <v-tabs class="sel-item" bg-color="#CDC3CE" style="width: 330px;height: 100%;padding-left:30px" v-model="tab" direction="vertical">
    <v-tab style="font-weight: 400;text-transform: none;" text="Консоль" value="0"/>
    <v-tab style="font-weight: 400;text-transform: none;" text="Общие" value="1"/>
    <v-tab style="font-weight: 400;text-transform: none;" text="Внешний вид" value="2"/>
    <v-tab style="font-weight: 400;text-transform: none;" text="Заказы" value="3"/>
    <v-tab style="font-weight: 400;text-transform: none;" text="Отчеты" value="4"/>
    <v-tab style="font-weight: 400;text-transform: none;display: none" text="Отчеты" value="5"/>
    <v-tab style="font-weight: 400;text-transform: none;display: none" text="Отчеты" value="6"/>
  </v-tabs>
  <v-tabs-window style="flex-grow: 1;" v-model="tab">
    <v-tabs-window-item  value="0"></v-tabs-window-item>
    <v-tabs-window-item  value="1"><GeneralPage @new-tab="newTab($event)"/></v-tabs-window-item>
    <v-tabs-window-item  value="2"></v-tabs-window-item>
    <v-tabs-window-item  value="3"><admin-purchases-page/></v-tabs-window-item>
    <v-tabs-window-item  value="4"></v-tabs-window-item>
    <v-tabs-window-item  value="5"><admin-products-page @new-tab="newTab($event)"/></v-tabs-window-item>
    <v-tabs-window-item  value="6"><admin-add-product-page /></v-tabs-window-item>

  </v-tabs-window>
</div>
</template>

<style scoped>
.sel-item >>> .v-slide-group--vertical{

}
.sel-item >>> .v-tab__slider{
  content: '' !important;
  position: absolute !important; /* Абсолютное позиционирование */
  right: 0px !important;
  left: auto !important;
  bottom: -20px !important; /* Положение треугольника */
  border: 23px solid transparent !important;
  border-right: 20px solid white !important; /* Добавляем треугольник */
  background: transparent !important;
}

</style>