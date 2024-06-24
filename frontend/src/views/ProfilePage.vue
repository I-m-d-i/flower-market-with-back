<script>
import MessageLogout from "@/components/messageLogout.vue";
import axios from "axios"

export default {
  name: "ProfilePage",
  components: {MessageLogout},
  data() {
    return {
      showWindowLogout: false,
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
    methods: {
      logout() {
        this.showWindowLogout = true;
      }
    }
}
</script>

<template>
  <div class="profile">
    <div class="profile_body">
      <div class="profile_header">
        <img class="profile_img" :src="require('@/assets/user icon.png')"/>
        <span class="profile_name">{{ user?.name }}</span>

      </div>
      <div class="profile_menu">
        <ul>
          <li>
            <button class="btn-menu" @click="$router.push('/profile/editName')">
              <span>Редактировать имя</span>
              <img style="width: 27px; height: 43px; align-self: center;" :src="require('@/assets/Group 1.png')">
            </button>
          </li>
          <li>
            <button class="btn-menu" @click="$router.push('/profile/notifications')">
              <span>Настройка уведомлений</span>
              <img style="width: 27px; height: 43px; align-self: center;" :src="require('@/assets/Group 1.png')">
            </button>
          </li>
          <li>
            <button class="btn-menu">
              <span>Мои реквизиты</span>
              <img style="width: 27px; height: 43px; align-self: center;" :src="require('@/assets/Group 1.png')">
            </button>
          </li>
          <li>
            <button class="btn-menu" @click="logout()">
              <span>Выйти из аккаунта</span>
              <img style="width: 27px; height: 43px; align-self: center;" :src="require('@/assets/Group 1.png')">
            </button>
          </li>
        </ul>
      </div>
    </div>
  </div>
<message-logout v-if="showWindowLogout" @close="showWindowLogout = false"/>
</template>

<style scoped>

.profile_header {
  display: grid;
  grid-template-columns: 300px 1fr;
  gap: 40px;
  align-items: center;
}

.profile_name {
  font: 400 40px Playfair Display;
  padding-bottom: 20px;
}

.btn-menu {
  background-color: #D5CBCB;
  padding: 10px 30px 20px 0;
  height: 110px;
  width: 100%;
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  justify-content: space-between;
}

li {
  padding: 15px 0 15px 0;
}

ul {
  list-style-type: none;
  width: 100%;
}
.profile_img {
  width: 330px;
}

ul span {
  padding-left: 20px;
  text-decoration: none;
  color: #6B5D76;
  font: 400 40px Playfair Display;
}
@media screen and (max-width: 768px) {
  .profile_header {
    grid-template-columns: 1fr;
    gap: 0px;
    justify-items: center;
  }
  .btn-menu span{
    font: 400 35px Playfair Display;
  }
}
@media  screen and (max-width: 290px) {
  .btn-menu span{
    font: 400 30px Playfair Display;
  }
  .profile_img {
    width: 200px;
  }
}


</style>