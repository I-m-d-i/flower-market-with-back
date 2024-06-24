<script>
import Header2 from "@/components/Header2.vue";

export default {
  name: "NotificationsPage",
  components: {Header2},
  data() {
    return {
      listNotifications: [
        {
          id: 1,
          name: "SMS-сообщения",
          active: false,
        },
      ]
    }
  }
}
</script>

<template>
  <header2/>
  <div class="container">
    <h2>SMS Рассылки</h2>
    <ul>
      <li v-for="notification in listNotifications" :key="notification.id">
        <div class="btn-menu">
          <span style="padding-left: 30px">{{ notification.name }}</span>
          <v-spacer/>
          <v-switch hide-details
                    inset
                    color="#81637D"
                    v-model="notification.active"/>
        </div>
      </li>
    </ul>
  </div>
</template>

<style scoped>
.v-input {
  padding-top: 15px;
  justify-items: end;
}
.btn-menu >>> .v-switch--inset .v-switch__track {
  height: 58px !important;
  width: 105px !important;
  background-color: #81637D !important;
}

.btn-menu >>> .v-selection-control--dirty .v-selection-control__input {
  transform: translateX(40px);
}

.btn-menu >>> .v-switch .v-selection-control__input {
  margin-left: -30px;
}
.btn-menu >>> .v-switch__thumb {
  transform: scale(1.2) !important;
  height: 40px;
  width: 40px;
}


h2 {
  font-family: Playfair Display;
  font-size: 40px;
  font-weight: 700;
  line-height: 53px;
  text-align: left;
  color: #6B5D76;
  padding-left: 30px;
  padding-bottom: 5px;
}
.btn-menu {
  background-color: #D5CBCB;
  padding: 10px 30px 20px 0;
  height: 110px;
  width: 100%;
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: space-between;
}


span {
  font-family: Playfair Display;
  font-size: 40px;
  font-weight: 700;
  line-height: 53.32px;
  color: #6B5D76;
  text-align: left;
}
ul {
  list-style-type: none;
}
li {
  margin-bottom: 30px;
}

.container {
  margin-top: 100px;
}

@media screen and (max-width: 600px) {
  span {
    font-size: 30px;
  }
  h2 {
    font-size: 30px;
  }
  .container {
    margin-top: 50px;
  }
}
</style>