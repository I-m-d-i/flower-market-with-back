<script>
import axios from "axios";

export default {
  name: "AdminPurchasesPage",
  data(){
    return{
      purchases: null,
    }
  },
  methods:{

  },
  mounted() {
    axios.get("api/admin/purchases",{withCredentials:true}).then((r)=>{
      this.purchases = r.data
    }).catch((e)=>{
      console.log(e)
    })
  }
}
</script>

<template>
<div style="">

  <h2>Заказы</h2>
  <div style="display: flex;justify-content: space-around;flex-direction: row; margin-top: 30px">
    <div class="card">
       <span>Номер заказа</span>
    </div>
    <div class="card">
      <span>Номер телефона</span>
    </div>
  </div>
  <div style="display: flex;justify-content: space-around;flex-direction: row; margin-top: 30px">
    <div class="card">
      <span>Статус заказа</span>
    </div>
    <div class="card">
      <span>Покупатель</span>
    </div>
  </div>
  <div style="padding-block: 5px;width: 100%; background-color: #D9D9D9; margin-top: 30px">
    <img :src="require('@/assets/iconPurchase.png')" style="width: 30px;height: 15px" alt="">
    <span>Номер заказа</span>
  </div>

  <table style="margin-top: 30px;padding-inline:20px">
    <thead>
    <tr>
      <th>Заказ №</th>
      <th>Покупатель</th>
      <th>Телефон</th>
      <th>Статус</th>
      <th>Всего</th>
      <th>Создан</th>
    </tr>
    </thead>
    <tbody>

    <tr v-for="purchase in purchases" :key="purchase.purchase_id">
      <td>{{purchase.purchase_id}}</td>
      <td>{{purchase.customer_name}}</td>
      <td>{{purchase.customer_phone}}</td>
      <td>В пути</td>
      <td>{{purchase.purchase_price}}</td>
      <td>{{purchase.purchase_date}}</td>
    </tr>
    </tbody>
  </table>
</div>
</template>

<style scoped>
table, th, td {
  border:1px solid black;
  border-spacing: 0px;
  border-collapse: collapse;
  text-align: center;

}
.card {
  border:1px solid black;
  width: 300px;
  height: 35px;
}
span {
  font-family: 'Playfair Display';
  font-style: normal;
  font-weight: 400;
  font-size: 16px;
  line-height: 21px;

  color: #000000;
}
</style>