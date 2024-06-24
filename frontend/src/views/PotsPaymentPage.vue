<script>
import axios from "axios";

export default {
  name: "PotsPaymentPage",
  data() {
    return {
      rating: 4.95,
      countProduct:0,
      user: null,
      productsNameList:[],
    }
  },
  props: {
    id: String,
  },
  mounted() {
    console.log(this.id);
    axios.postForm("/api/purchase/detail", {purchase_id: this.id},{withCredentials:true}).then((r) => {
      console.log(r.data)
      for (let i = 0; i < r.data.items.length; i++) {
        this.countProduct += r.data.items[i].product_count
        this.productsNameList.push(r.data.items[i].product_name)
      }
      console.log(this.productsNameList)
    }).catch((e)=>{
      console.log(e)
    })
    axios.get("/api/customer/profile",{withCredentials:true}).then((r)=>{
      this.user = r.data
    })
  },
}
</script>

<template>
  <div>
    <v-btn variant="text" @click="$router.push({name:'home'})" style="position:absolute; top:10px;left:10px">назад</v-btn>
    <div class="pick-up-point">
      <p style="padding-bottom: 30px">Пункт выдачи</p>
      <div style=" display: inline-flex;width: 100%;   align-items: center;    justify-content: space-between;">
        <div/>
        <div>
          <p>Красноярск, Красноярский рабочий 156</p>
          <div class="rating">
            <span>{{ rating }}</span>
            <v-rating
                style="padding-top: 5px;padding-left: 5px"
                v-model="rating"
                active-color="#DBCB3D"
                color="#DBCB3D"
                readonly
                :size="25"
                half-increments
            ></v-rating>
          </div>
          <span>11-13 декабря</span>
        </div>
        <div class="count-product">
          <span>{{countProduct}} товар</span>
        </div>
      </div>
    </div>
    <div style="margin-left: 60px; margin-top:35px">
      <p>СПАСИБО ЗА ЗАКАЗ!</p>
      <br/>
      <p>
        АДРЕС ДОСТАВКИ:<br/>
        Красноярск, Красноярский рабочий 156
      </p>
      <br/>
      <p style="text-transform: uppercase">
        СПОСОБ ОПЛАТЫ: карта
      </p>
      <br/>
      <p>ИНФОРМАЦИЯ О ПОКУПКЕ:</p>
        <span>ФИО: {{user.customer_name}}</span><br/>
        <span>ГОРОД: {{user.customer_city}}</span><br/>
        <span>ТОВАРЫ: {{productsNameList.join('\n')}}</span><br/>
    </div>
  </div>
</template>

<style scoped>
.pick-up-point {
  margin-top: 50px;
  margin-left:51px;
  border-radius: 51px;
  background: rgba(238, 207, 207, 0.42);
  width: 100%;
  max-width: 1376px;
  height: 283px;
  flex-shrink: 0;
  padding: 15px 60px 30px 40px;
}
p {
  font-family: 'Playfair Display', serif;
 font-size: 24px;
}
span {
  font-family: 'Playfair Display', serif;
  font-size: 24px;
}
</style>