<script>
import Header2 from "@/components/Header2.vue";
import axios from "axios";

export default {
  name: "EditNamePage",
  components: {Header2},
  data() {
    return {
      user: null
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
    accept() {
      axios.post()
      this.$router.push("/profile");
    }
  }
}
</script>

<template>
  <header2><button class="btn_ready" @click="accept()">Готово</button></header2>
  <div class="row_edit_name">
    <span style="padding-left: 30px">Имя</span>
    <v-spacer/>
    <input class="input_name" v-model="user.customer_name" type="text">
  </div>

</template>

<style scoped>
span {
  font-family: Playfair Display;
  font-size: 40px;
  font-weight: 700;
  line-height: 53.32px;
  text-align: left;
  color: #6B5D76;
}
.input_name{
  font-family: Playfair Display;
  font-size: 40px;
  font-weight: 700;
  line-height: 53.32px;
  text-align: center;
  color: #514F4F;
}
.btn_ready{
  font-family: Playfair Display,serif;
  font-size: 40px;
  font-weight: 700;
  line-height: 53.32px;
  text-align: left;
  margin-right: 100px;
  color: #FFFFFF
}
.row_edit_name {
  margin-top: 120px;
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
@media screen and (max-width: 760px) {
  .btn_ready {
    font-size: 30px;
    margin-right: 50px;
  }
}
</style>