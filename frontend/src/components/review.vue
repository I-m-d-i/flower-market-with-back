<script>
export default {
  name: "review",
  props: {
    image_src: {
      type: String,
      required: true,
    },
    author: {
      type: String,
      required: true,
    },
    text: {
      type: String,
      required: true,
    },
    rating: {
      type: Number,
      required: true,
    },
    id:{
      type: Number,
      required: true,
    }
  },
}
</script>

<template>
<div  class="review_card">
  <div class="review_container">
    <div class="review_img">
      <img :src="image_src" alt=""/>
    </div>
    <div class="review_author">
      <span>{{author}}</span>
    </div>
    <div class="review_rating">
      <v-rating
          hover
          readonly
          :length="5"
          :size="32"
          :model-value="rating"
          color="#FFFFFF"
          active-color="#A795A5"
      />
    </div>
    <div class="review_text">
      <span>{{text}}</span>
    </div>
  </div>
</div>
</template>

<style scoped>
.review_card{
  padding: 10px;
}

.review_container{
  display: flex;
  flex-direction: column;
  max-width: 350px;
}
.review_img{
  align-self: center;
  margin-bottom: 10px;
}
span {
  color: #000;
  font-family: Anonymous Pro;
  font-size: 24px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
}
</style>