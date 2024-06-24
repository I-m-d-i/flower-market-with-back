<script>
import axios from "axios";

export default {
  name: "AdminAddProductPage",
  data() {
    return {
      image_src: null,
      title: null,
      description: null,
      price: null,
    }
  },
  methods:{
    uploadImage(){
      const file = this.$refs.fileInput.files[0];

      if (!file.type.startsWith("image/")) {
        alert("Please upload an image file");
        return;
      }
      const formData = new FormData();
      formData.append('file', file);
      axios.post( '/api/admin/image/upload',
          formData,
          {
            headers: {
              'Content-Type': 'multipart/form-data'
            }
          }
      ).then((r)=>{
        this.image_src = r.data
      })
          .catch(function(){
            console.log('FAILURE!!');
          });
    },
    save(){
      axios.post("api/admin/product/add",{
        product_name: this.title,
        description: this.description,
        price: this.price,
        img_src: this.image_src
      },{withCredentials: true}).then((r)=>{
        console.log(r)
      }).catch((e)=>{
        console.log(e)
      })
    }

  }
}
</script>

<template>
  <h2>Товары</h2>
  <div class="div">
    <a class="img_container">
      <img :src="image_src || require('@/assets/imgProductSkeleton.png')" class="img" :alt="title">
      <label v-if="image_src==null" class="input-file">
        <input  type="file" id="img" ref="fileInput" name="file" @change="uploadImage"  placeholder="ФОТО" >
        <span>фото</span>
      </label>
    </a>
    <div class="div-2">
      <div class="div-3">
        <div class="div-4"><v-textarea hide-details rows="2" variant="plain" v-model="title" label="Название"></v-textarea></div>
        <div class="div-5"><v-textarea hide-details rows="2" variant="plain" v-model="description" label="Описание"></v-textarea></div>
      </div>
      <div class="div-6">
        <div class="div-7">
          <v-textarea hide-details rows="2" variant="plain" v-model.number="price" label="Цена"/>
        </div>
        <button class="div-8" @click="save()">Сохранить</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
h2 {
  font-size: 24px;
  margin-left: 30px;
  font-weight: normal;
}
.input-file {
  position: absolute;
  left: 32%;
  top: 50%;
  display: inline-block;
}
.input-file span {
  position: relative;
  display: inline-block;
  cursor: pointer;
  outline: none;
  text-decoration: none;
  font-size: 36px;
  vertical-align: middle;
  color: rgb(0, 0, 0);
  text-align: center;
  border-radius: 4px;
  background-color: rgba(65, 145, 82, 0);
  line-height: 22px;
  height: 40px;
  padding: 10px 20px;
  box-sizing: border-box;
  border: none;
  margin: 0;
  transition: background-color 0.2s;
}
.input-file input[type=file] {
  position: absolute;
  z-index: -1;
  opacity: 0;
  display: block;
  width: 0;
  height: 0;
}

/* Focus */
.input-file input[type=file]:focus + span {
  box-shadow: 0 0 0 0.2rem rgba(0,123,255,.25);
}

/* Hover/active */
.input-file:hover span {
  background-color: #dcdcdc;
}
.input-file:active span {
  background-color: #cdcdcd;
}

/* Disabled */
.input-file input[type=file]:disabled + span {
  background-color: #eee;
}

.div {
  display: flex;
  flex-direction: column;
  padding: 5px;
  margin-left: 50px;
  max-width: 300px;
  align-items: center;
}

button {
  border: none;
}

.img_container {
  position: relative;
}

.img {
  border-radius: 368px 368px 0px 0px;
  width: 100%;
  aspect-ratio: 0.84;
}

.div-2 {
  align-self: start;
  display: flex;
  width: 100%;
  align-items: flex-start;
  justify-content: space-between;
  margin: 20px 0 0 0px;
  gap: 5px;
}

.div-3 {
  align-self: start;
  flex-grow: 1;
  width: 100px;
  align-items: flex-start;
  display: flex;
  flex-direction: column;
}

.div-4 {
  color: #000;
  font-family: Playfair Display;
  font-size: 24px;
  line-height: normal;
  font-variant: small-caps;
}

.div-5 {
  color: #000;
  margin-top: 18px;
  font-family: Playfair Display;
  font-size: 24px;
  max-height: 100px;
  line-height: normal;
  font-variant: small-caps;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.div-6 {
  flex: 1;
  padding-right: 10px;
}

.div-7 {
  color: #000;
  align-self: stretch;
  font-family: Playfair Display;
  font-size: 24px;
  line-height: normal;
  font-variant: small-caps;
  text-align: center;
}

.div-8 {
  color: #fffbfb;
  align-self: stretch;
  border-radius: 40px;
  background-color: #81637d;
  margin-top: 26px;
  width: 100%;
  flex-grow: 1;
  min-width: 120px;
  min-height: 30px;
  font: small-caps 400 18px Playfair Display, -apple-system, Roboto, Helvetica, sans-serif;
}


@media (max-width: 1100px) {
  .div-8 {
    padding: 0 -5px 0 2px;
  }

  .div-4 {
    font-size: 16px;
  }

  .div-5 {
    font-size: 14px;
    margin-top: 10px;
  }

  .div-7 {
    font-size: 24px;
  }

  .div-8 {
    font-size: 10px;;
    margin-top: 10px;
    min-width: 75px;
  }

  .div-6 {
    padding-right: 0px;
  }

  .div-2 {
    margin-top: 5px;
  }
}

@media (max-width: 440px) {
  .div-4 {
    font-size: 14px;
  }

  .div-5 {
    font-size: 10px;
  }
}

@media (max-width: 991px) {
  .div-8 {
    padding: 0 10px;
  }
}
</style>