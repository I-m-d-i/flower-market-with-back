<script>
export default {
  name: "Header-component",
  components: {},
  computed: {
    countBasket() {
      return this.$store.getters.countBasket;
    },
  },
  data() {
    return {
      registerWindow: false
    }
  },
  methods: {
    scrollTo(id) {
      this.$router.push({name: "home", hash: `#${id}`})
    },
    profile() {
        this.$router.push('/profile')
    },
    showMenu() {
      document.querySelector(".nav").classList.toggle("active");
    }
  },
}
</script>

<template>
  <!--Логотип выравненный по левую сторону-->
  <nav class="header">
    <div class="logo_container">
      <img class="logo" alt="logo" :src="require('@/assets/logo.png')">
    </div>
    <div class="menu">
      <div class="hamburger_container" @click="showMenu()">
        <div class="screen1_menu_hamburger"></div>
      </div>
      <ul class="nav">
        <li><a @click="scrollTo('popular')">Популярное</a></li>
        <li><a @click="scrollTo('catalog')">Каталог</a></li>
        <li><a @click="scrollTo('about')">О нас</a></li>
        <li><a @click="scrollTo('footer')">Контакты</a></li>
      </ul>
    </div>
    <div class="header_icons">
      <img style="cursor: pointer" alt="профиль" width="46" height="46" :src="require('@/assets/profile icon.png')"
           @click="profile()">
      <a href="/#/basket">
        <template v-if="countBasket===0">
          <img style="" alt="корзина" width="46" height="46" :src="require('@/assets/basket.png')">
        </template>
        <template v-else>
          <v-badge :content="countBasket">
            <img style="" alt="корзина" width="46" height="46" :src="require('@/assets/basket.png')">
          </v-badge>
        </template>
      </a>
    </div>
  </nav>
</template>

<style scoped>

.mobile_menu_icon {
  display: none;
}

.header_icons {
  padding-block: 10px;
  display: flex;
  flex-flow: row nowrap;
  gap: 10px;
}

.logo {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 200px;
  height: 250px;
  object-fit: cover;
}

.logo_container {
  position: relative;
  margin-right: 30px;
  overflow: hidden;
  width: 200px;
  height: 95px;
}

.header {
  position: relative;
  z-index: 2;
  padding: 12px 25px 30px;
  display: flex;
  justify-content: flex-start;
  align-items: flex-start;
}

.menu a {
  text-decoration: none;
  color: #4F3960;
  font-family: Inconsolata, sans-serif;
  font-size: 20px;
}

.menu {
  flex-grow: 1;
  padding-top: 20px;
}

.menu ul {
  display: flex;
  margin: 0;
  padding: 10px 0;
  list-style: none;
  text-align: left;
}

.menu li {
  padding-right: 40px;
}

@media (max-width: 960px) {
  .screen1_menu_hamburger {
    display: block !important;
  }

  .menu ul {
    display: none;
    margin-top: 20px;
    flex-direction: column;
    transition: all .5s ease-in-out;
  }

  .menu ul li {
    padding-bottom: 10px;
  }

  .header {
    backdrop-filter: blur(5px);
    box-shadow: inset 0 0 0 3000px rgba(150, 150, 150, 0.192);
    background-color: rgba(187, 143, 178, 0.32);
    padding-inline: 15px !important;
    border: 1px solid rgba(112, 112, 112, 0.56);
    overflow: hidden;
  }

  .active {
    display: block !important;
  }

  .logo_container {
    display: none;
  }
  .header_icons{
    transform: scale(0.8);
  }
}

.hamburger_container {
  height: 41px;
  width: 41px;
  padding-top: 19px;
  padding-left: 6px;
  cursor: pointer;
  border-radius: 30px;
}

.hamburger_container:hover {
  background-color: rgba(130, 100, 126, 0.65);
}

@media (hover: none) {
  .hamburger_container:hover {
    background-color: inherit;
  }
}

.screen1_menu_hamburger {
  display: none;
  width: 30px;
  height: 2px;
  background-color: #533557;
  border-radius: 5px;
  z-index: 10;
  transition: all .5s ease;
}

.screen1_menu_hamburger::before,
.screen1_menu_hamburger::after {
  content: '';
  position: absolute;
  width: 30px;
  height: 2px;
  background: #533557;
  border-radius: 5px;
  transition: all .5s ease;
}

.screen1_menu_hamburger::before {
  transform: translateY(-9px);
}

.screen1_menu_hamburger::after {
  transform: translateY(9px);
}

.screen1_menu_btn input:checked ~ .screen1_menu_hamburger::before {
  transform: rotate(45deg) translate(35px, -35px);
}

.screen1_menu_btn input:checked ~ .screen1_menu_hamburger::after {
  transform: rotate(-45deg) translate(35px, 35px);
}

.screen1_menu_btn input:checked ~ .screen1_menu_hamburger {
  transform: translateX(-50px);
  background: transparent;
}


</style>