<script>
export default {
  name: "messageLogout",
  data() {
  },
  methods: {
    logout() {
      this.$store.dispatch('logout');
      this.$router.push({name: "home"});
      this.$emit("close")
    }
  }
}
</script>

<template>
  <div class="message_container">
    <div class="message">
      <div class="message_body">
        <div>
          <div class="message_h">
            <p>обратите внимание</p>
          </div>
          <div class="message_text">
            <p>Вы уверены, что хотите выйти?</p>
          </div>
        </div>
        <div class="message_action">
          <v-btn flat class="btn button-logout" @click="logout()"><span>Выйти</span></v-btn>
          <v-spacer/>
          <v-btn flat class="btn button-cancel" @click="$emit('close')">Отмена</v-btn>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.message_action{
  display: flex;
  gap: 30px;
  padding-bottom: 20px;
}
.btn {
  width: 248px;
  height: 67px;
  align-self: stretch;
  border-radius: 21px;
  background-color: #D9D9D9;
  flex-grow: 1;
  padding: 12px 25px 12px 25px;
  font: small-caps 400 28px Playfair Display, -apple-system, Roboto, Helvetica, sans-serif;
}
.button-logout {
  color: #C20000;
}
.button-cancel {
  color: #870B7A;
}
.message_text{
  font: 400 28px Playfair Display;
}
.message_h {
  text-align: center;
  font: 400 36px Playfair Display;
}
.message_body {
  display: flex;
  flex-direction: column;
  height: 100%;
  align-items: center;
  padding-top: 10px;
  padding-bottom: 15px ;
  justify-content: space-between;
}
.message_container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  position: absolute;
  top: 0;
}
.message {
  display: block;
  width: 635px;
  height: 273px;
  border-radius: 51px;
  background-color: #E7E0E7;
}
.v-btn {
  text-transform: none;
}
</style>