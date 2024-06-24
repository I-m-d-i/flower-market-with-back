<script>
export default {
  name: "Header2"
}
</script>

<template>
  <div class="new_header">
    <img class="arrow" :src="require('@/assets/arrow_white.png')" alt="Назад" @click="$router.go(-1)" />
    <h1>{{ $route.name}}</h1>
    <div style="justify-self: end;    margin-top: -15px;">
      <slot></slot>
    </div>
  </div>
</template>

<style scoped>
.new_header{
  width: 100%;
  height: 187px;
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  align-items: center;
  background-color: #AF9595;
  justify-items: stretch
}
.arrow {
  padding-left: 82px;
  width: 130px;
  cursor: pointer;
  height: 75px;
}
.new_header h1 {
  margin-top: -15px;
  font-family: Playfair Display;
  margin-left: -50px;
  text-align: center;
  font-size: 40px;
  font-weight: 700;
  line-height: 53.32px;
  color: #FFFFFF;
}

@media (max-width: 700px) {
  .new_header h1 {
    font-size: 23px;
    padding-left: 10px;
  }
  .arrow{
    padding-left: 20px;
    width: 55px;
    height: 50px;
  }
}
@media (max-width: 375px) {
  .new_header h1 {
    font-size: 19px;
    padding-left: 30px;
  }
  .arrow{
    padding-left: 20px;
    height: 35px;
  }
  .new_header{
    height: 100px;
  }
}
@media (max-width: 280px) {
  .new_header h1 {
    font-size: 18px;
    padding-left: 32px;
  }
  .arrow{
    padding-left: 20px;
    height: 35px;
  }
}
</style>