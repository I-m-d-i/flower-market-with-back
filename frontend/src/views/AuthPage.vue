<script>
import axios from "axios";

export default {
  name: "AuthPage",
  data() {
    return {
      user: null,
    }
  },
  mounted() {
    axios.get("/api/customer/profile",{withCredentials:true}).then((r)=>{
      this.user = r.data
    }).catch((e)=>{
      console.log(e)
    })
  },
  computed: {
    phone() {
      return this.user.customer_phone
    }
  }
}
</script>

<template>
  <div class="container">
    <img class="arrow" :src="require('@/assets/arrow_white.png')" alt="Назад" @click="$router.go(-1)" />
    <div class="img_container">
      <img :src="require('@/assets/phones.png')"/>
    </div>
    <span>Вы хотите войти под номером</span>
    <span class="phone"> {{phone}}?</span>
    <span >После входа вам будут доступны товары с персональными скидками</span>
    <button class="btn btn_login" @click="$store.dispatch('auth'); $router.push({name: 'home'})">Да, войти</button>
    <button class="btn btn_register" @click="$store.dispatch('logout'); $router.push({name: 'home'})">Войти под другим номером</button>
  </div>
</template>

<style scoped>
.arrow {
  position: absolute;
  top: 70px;
  left: 45px;
  cursor: pointer;
}
span {
  font-family: Playfair Display;
  font-size: 40px;
  font-weight: 400;
  line-height: 53.32px;
  text-align: center;
  max-width: 900px;
}
.container {
  background-color: #EEE1E1;
  display: grid;
  gap: 20px;
  align-items: center;
  justify-items: center;
  align-content: center;
  padding-block: 20px;
  justify-content: center
}
.phone {
  font-family: Podkova, serif;
  font-size: 40px;
  font-weight: 400;
  line-height: 44.32px;
  text-align: left;
  margin-block: 10px;
}
.img_container {
  position: relative;
  overflow:hidden;
  height:400px;
  width: 100%;
}
.img_container img {
  position: absolute;
  object-fit:cover;
  transform:translate(-50%,-50%);
  width: 100%;
  top:50%;
  left:50%;
}
.btn {
  font-family: Playfair Display;
  font-size: 28px;
  font-weight: 400;
  line-height: 37.32px;
  text-align: center;
  color: #FFFFFF;
  border-radius: 21px;
  height: 92px;
  min-width:  528px;
}
.btn_login {
  margin-top: 50px;
  background: #A590A1;
}
.btn_register{
  background: #A590A1AD;
}


@media screen and (max-width: 760px) {
  .btn {
    min-width:  300px;
  }
}

@media screen and (max-width: 620px) {
  span {
    font-size: 30px;
  }
}
@media screen and (max-width: 500px) {
  span {
    font-size: 20px;
    max-width: 400px;
  }
  .arrow{
    top: 35px;
    transform: scale(0.7);
  }
}
@media screen and (max-width: 400px) {
  span {
    font-size: 15px;
    max-width: 300px;
  }
  .phone {
    font-size: 30px;
  }
  .container {
    gap: 15px;
  }
  .btn {
    min-width:  200px;
    font-size: 20px;
  }
  .btn_login {
    margin-top: 10px;
  }
}
</style>