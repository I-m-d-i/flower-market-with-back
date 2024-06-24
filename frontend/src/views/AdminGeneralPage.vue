<script>
export default {
  name: "Ganeral"
}
</script>

<template>
  <h2>Общие</h2>
  <div class="general">
    <v-row>
        <v-card
          class="pa-2 card"
          variant="flat"
          @click="$emit('new-tab',5)"
        >
          Товары
          <div class="counter">
            100
          </div>
        </v-card>
        <v-card
          class="pa-2 card"
          variant="flat"
          @click="$emit('new-tab',6)"
        >
          Скидки
          <div class="counter">
            100
          </div>
        </v-card>
    </v-row>
    <v-row>
      <v-card
          class="pa-2 card"
          variant="flat"
          @click="$emit('new-tab',7)"
      >
        Партнеры
        <div class="counter">
          100
        </div>
      </v-card>
      <v-card
          class="pa-2 card"
          variant="flat"
          @click="$emit('new-tab',8)"
      >
        Общие
        <div class="counter">
          100
        </div>
      </v-card>
    </v-row>
  </div>
</template>

<style scoped>
.general {
  display:flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  gap:50px
}
.card {
  width: 300px;
  height: 175px;
  background-color: #FFFFFF;
  margin-left: 100px;
}
.counter {
  display: flex;
  height: 100%;
  flex-direction: row;
  align-items: flex-end;
  justify-content: flex-end;
  padding: 20px;
}
</style>