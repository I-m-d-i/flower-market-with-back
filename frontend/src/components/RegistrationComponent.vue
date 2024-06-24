<script>
export default {
  name: "RegistrationComponent",
  methods: {
    registration() {
      this.$store.dispatch("registration", this.formData);
      this.$emit("close")
    },
    login() {
      console.log(this.to.name)
      this.$store.dispatch("login", {user:this.formData,to: this.to});
      this.$emit("close")
    }
  },
  props: {
    to: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      redirect: null,
      isSingIn: true,
      formData: {
        name: "",
        city: "",
        phone: "",
        password: ""
      }
    }
  }
}
</script>

<template>
    <div class="modal">
      <div class="modal_body">
        <div class="modal_h">
          <p>{{isSingIn?"Вход":"Регистрация"}}</p>
        </div>
        <form class="form_reg" @submit.prevent="isSingIn?login():registration()">
          <input v-if="!isSingIn" v-model="formData.name" type="text" name="name" placeholder="ФИО">
          <input v-if="!isSingIn" v-model="formData.city" type="text" name="city" placeholder="Город">
          <input v-model="formData.phone"  type="text" name="phone" placeholder="Номер телефона">
          <input v-model="formData.password" type="password" name="password" placeholder="Пароль">
          <a style="margin-top: -15px;color: midnightblue;font-size: 14px;" @click="isSingIn=!isSingIn">{{isSingIn?"Регистрация":"Вход"}}</a>
          <button  class="btn_reg" type="submit">Готово</button>
        </form>
      </div>
  </div>

</template>

<style scoped>
.form_reg {
  margin-top: 10px;
  display: flex;
  flex-direction: column;
  gap: 20px;
  align-items: center
}

input {
  min-height: 50px;
  max-width: 370px;
  border-radius: 40px;
  background-color: #D5CBCB;
  color: #6B5D76;
  flex: 1;
  font: 400 24px Playfair Display, sans-serif;
  outline: none;
  padding: 10px 0 10px 20px;

}

button {
  color: #4F3960;
  max-width: 264px;
  height: 65px;
  border-radius: 33px;
  background-color: #D5CBCB;
  flex: 1;
  padding: 12px 45px 12px 45px;
  margin-top: 45px;
  font: small-caps 400 28px Playfair Display;
}

.modal_h {
  text-align: center;
  font: 400 28px Playfair Display;
  padding-block: 45px;
}

.modal {
  max-width: 500px;
  max-height: 645px;
  display: block;
  border-radius: 90px;
  background-color: #EDE5E5;
}

.modal_body {
  display: flex;
  align-items: center;
  flex-direction: column;
  align-content: stretch;
  width: 100%;
  height: 100%;
  padding-block: 25px;
  padding-inline: 5px;
}

@media screen and (max-width: 400px) {
  input {
    max-width: 240px;
    font-size: 20px;
    min-height: 50px;
  }
  button {
    font-size: 20px;
    margin-top: 25px;
  }
}
</style>