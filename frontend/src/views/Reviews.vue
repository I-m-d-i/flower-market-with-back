<script>
import Review from "@/components/review.vue";

export default {
  name: "Reviews",
  components: {Review},
  data() {
    return {
      reviews: [
        {
          id: 1,
          image_src: require('@/assets/Tom1.png'),
          author: "том Холланд",
          rating: 5,
          text: "Спасибо за такой удобный формат магазина и быструю доставку!"
        },
        {
          id: 2,
          image_src: require('@/assets/Tom2.png'),
          author: "том Эднотт",
          rating: 5,
          text: "отличный и удобный сайт!"
        },
        {
          id: 3,
          image_src: require('@/assets/Elli.png'),
          author: "Элли Хьюбберт",
          rating: 4,
          text: "немного заезженный сайт, но в целом нормально"
        },

      ]
    }
  }
}
</script>

<template>
  <div id="reviews" style="margin-top: 100px">
    <h1>Отзывы</h1>
    <div style="display: flex; flex-wrap: wrap  ;justify-content: space-evenly;margin-bottom: 200px;">
      <review v-for="(review,i) in reviews" :key="i" :id="review.id" :rating="review.rating" :text="review.text" :author="review.author" :image_src="review.image_src"></review>
    </div>
  </div>

</template>

<style scoped>
h1 {
  color: #000;
  font-family: Source Serif Pro;
  font-size: 48px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  font-variant: all-small-caps;
  margin-left: 30%;
  padding-top: 40px;
  padding-bottom: 50px;
}
</style>