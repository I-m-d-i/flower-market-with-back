<script>
import axios from "axios";

export default {
  name: "PaymentPage",
  data() {
    return {
      isOnline: true,
      isOnline2: false,
      rating: 4.95,
    }
  },
  computed: {
    countProduct() {
      return this.$store.getters.countBasket
    },
    totalCost() {
      return this.$store.getters.totalCost
    }
  },
  methods: {
    toPay(){
      const basket = this.$store.getters.getBasket
      axios.post('/api/purchases/add',{store_id:1,products:basket},{withCredentials: true}).then((r)=>{
        console.log(r)
        this.$router.push({path: '/potsPayment/'+r.data.purchase_id})
      }).catch((e)=>{
        console.log(e)
      })
    }
  }

}
</script>

<template>
  <div class="payment-page">
    <v-btn variant="text" @click="$router.go(-1)" style="position:absolute; top:10px;left:10px">назад</v-btn>
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
    <div class="payment-method">
      <p style="padding-bottom: 30px">СПОСОБ ОПЛАТЫ</p>
      <div style="display: grid; gap: 21px">
        <div class="payment-method-card" @click="isOnline2 = true">
          <span>Карта: ·· 7636 </span>
          <div class="payment-method-icon"
               v-bind:style="isOnline2?{backgroundColor: '#D5CBCB',border: 'solid 3px #FFFFFF'}:''"></div>
        </div>
        <div class="payment-method-card" @click="isOnline2 = false">
          <span>Наличными при получении</span>
          <div class="payment-method-icon"
               v-bind:style="!isOnline2?{backgroundColor: '#D5CBCB', border: 'solid 3px #FFFFFF'}:''"></div>
        </div>
      </div>
    </div>
    <div v-if="isOnline2" class="payment-time">
      <p style="align-self: start">Когда списать оплату? </p>
      <div class="payment-time-container">
        <button v-bind:style="isOnline?{backgroundColor: '#DDB9B9'}:''" @click=" isOnline = true">СРАЗУ</button>
        <button v-bind:style="!isOnline?{backgroundColor: '#DDB9B9'}:''" @click=" isOnline = false">ПРИ ПОЛУЧЕНИИ
        </button>
      </div>
    </div>
    <div class="total">
      <div class="total-title">
        <p>К ОПЛАТЕ</p>
        <p style="font-size: 24px">ДОСТАВКА</p>
      </div>
      <div class="total-cost-container">
        <p class="totalCost" >{{totalCost}} ₽</p>
        <p style="color: #407D4D;font-size: 24px; margin-top: -6px;">БЕСПЛАТНО</p>
      </div>

    </div>
    <button class="btn_to_pay" @click="toPay()">Заказать</button>
  </div>
</template>

<style scoped>
span, p, button {
  font-family: 'Playfair Display', serif;
  font-size: 20px;
  font-style: normal;
  line-height: normal;
}
.total p{
  padding-bottom: 24px;
}

.total {
  width: 100%;
  max-width: 1376px;
  display: flex;
  flex-direction: row;
  justify-content: space-evenly;
  margin-top: 58px;
}

.totalCost{
  font-family: Source Serif Pro;
  font-size: 32px;
  font-style: normal;
  line-height: normal;
  font-variant: all-small-caps;
  margin-top: -10px;
}

.payment-page p {
  align-self: start
}

.rating {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: left;
  margin-bottom: 15px;
}

.rating span {
  line-height: 20px;
}

.payment-method-card {
  display: flex;
  justify-content: space-between;
  border-radius: 14px;
  background: #DDB9B9;
  height: 59px;
  align-items: center;
  cursor: pointer;
  padding-inline: 20px;
}

.payment-page {
  padding-block: 30px;
  padding-inline: 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.pick-up-point {
  border-radius: 51px;
  background: rgba(238, 207, 207, 0.42);
  width: 100%;
  max-width: 1376px;
  height: 283px;
  flex-shrink: 0;
  padding: 15px 60px 30px 40px;
}

.payment-method {
  width: 100%;
  border-radius: 51px;
  background: rgba(238, 207, 207, 0.42);
  max-width: 1376px;
  max-height: 283px;
  margin-top: 44px;
  padding: 25px 25px 80px;
}

.payment-time {
  width: 100%;
  justify-content: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  max-width: 1376px;
}

.payment-time p {
  align-self: start;
  padding-block: 21px;
  padding-left: 49px;
}

.payment-time-container {
  width: 100%;
  margin-top: 22px;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  padding-inline: 25px;
  gap:15px;
  max-width: 901px;
  height: 131px;
  border-radius: 51px;
  background: rgba(238, 207, 207, 0.42);
}

.payment-time button {
  width: 100%;
  max-width: 348px;
  border-radius: 14px;
  background: #EAD4D4;
  height: 59px;
}

@media screen and (max-width: 500px) {
  .payment-time button {
    font-size: 14px;
  }
}

.payment-method-icon {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  background: #FFFFFF;
}
.btn_to_pay{
  width: 100%;
  max-width: 350px;
  border-radius: 14px;
  height: 40px;
  background-color: #D9D9D9;
}
</style>