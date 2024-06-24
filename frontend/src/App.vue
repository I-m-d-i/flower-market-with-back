<template >
  <div>
    <router-view />
    <v-dialog v-model="registerWindow" max-width="400">
      <registration-component :to="to" @close="registerWindow = false"/>
    </v-dialog>
  </div>
</template>
<script>


import RegistrationComponent from "@/components/RegistrationComponent.vue";
import bus from './plugins/eventBus'

export default {
  name: "App",
  data() {
    return{
      registerWindow: false,
      to: null,
    }
  },
  components: {
    RegistrationComponent
  },
  mounted() {
    this.$store.dispatch("loadAllProducts");
    bus.$on("openRegistration", (to) => {
      console.log(to)
      this.registerWindow = true
      this.to = to
    })
  }

}
</script>


<style>
@import url('https://fonts.cdnfonts.com/css/source-serif-pro');
@import url('https://fonts.cdnfonts.com/css/crimson-text');
@import url('https://fonts.cdnfonts.com/css/playfair-display');
@import url('https://fonts.cdnfonts.com/css/anonymous-pro');
@import url('https://fonts.cdnfonts.com/css/inconsolata-2');
@import url('https://fonts.cdnfonts.com/css/podkova');
a {
  cursor: pointer;
}
</style>
