<script>
export default {
  name: "About"
}
</script>

<template>
  <div id="about">
    <h1 style="margin-left: 306px">О нас</h1>
    <div style="
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
padding: 20px;
gap:30px">
            <span>
      ПРИВЕТ, МЫ СТУДИЯ ФЛОРИСТИКИ И ДЕКОРА.<br/>
      НАША КОМАНДА СОСТОИТ ИЗ СПЕЦИАЛИСТОВ В<br/>
      ОБЛАСТИ ФЛОРИСТИКИ И ДИЗАЙНА,<br/>
      ПРОДАЖ.ЛОГИСТИКИ И МАРКЕТИНГА.<br/>
      У НАС ЦАРИТ ОСОБАЯ АТМОСФЕРА, КОТОРАЯ<br/>
      ПОМОГАЕТ ПОЛУЧАТЬ ВДОХНОВЕНИЕ ДЛЯ СОЗДАНИЯ<br/>
      НОВЫХ ШЕДЕВРОВ ФЛОРИСТИКИ, ДИЗАЙНА И ДЕКОРА. МЫ ЛЮБИМ ТО,<br/>
      ЧТО ДЕЛАЕМ, И ЭТО КЛЮЧЕВОЙ ПОКАЗАТЕЛЬ<br/>
      НАШЕГО УСПЕХА.<br/>
      ЛЮБОВЬ К ЦВЕТАМ И ЛЮДЯМ, ИСКРЕННОСТЬ<br/>
      ЭМОЦИЙ И ВПЕЧАТЛЕНИЙ, ГОРЯЩИЕ ГЛАЗА - ВСЕ ТО,<br/>
      ЧТО ДЕЛАЕТ НАС СОВЕРШЕННЕЕ ДЕНЬ ОТО ДНЯ.<br/>
    </span>
      <img :src="require('@/assets/About1.png')" alt="О нас"/>
    </div>
    <div style="
  display: flex;
   flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 20px;
    margin-top: 30px;
gap:30px"
    >
      <img :src="require('@/assets/About2.png')" alt="О нас"/>
      <span>
      У НАС ВСЕГДА ПРЕДСТАВЛЕН<br/>
      ШИРОКИЙАССОРТИМЕНТ ОТЛИЧНЫХ ЦВЕТОВ.<br/>
      НАШ КАТАЛОГ НАСЧИТЫВАЕТ МНОГО<br/>
      ДЕСЯТКОВ НАИМЕНОВАНИЙ БУКЕТОВ, ОЧЕНЬ<br/>
      КРАСИВЫХ И ГАРМОНИЧНЫХ, КОТОРЫЕ БУДУТ<br/>
      ПРЕКРАСНО СМОТРЕТЬСЯВ ЛЮБОМ ЖИЛИЩЕ<br/>
      ИЛИ ОФИСЕ. КРОМЕ БУКЕТОВ, НАШИ КЛИЕНТЫ<br/>
      МОГУТ ПРИОБРЕСТИ<br/>
      ЦВЕТОЧНЫЕ КОМПОЗИЦИИ В ВИДЕ<br/>
      ДЕКОРАТИВНЫХ КОРЗИНОК, ШЛЯПНЫХ<br/>
      КОРОБОК.ЕСТЬ ТАКЖЕ «СОПУТСТВУЮЩИЕ»<br/>
      ТОВАРЫ: МЯГКИЕ ИГРУШКИ, ШАРЫ С ГЕЛИЕМ,<br/>
      ШОКОЛАД.<br/>
    </span>
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
  font-variant: small-caps;
}

span {
  color: #000;
  font-family: Source Serif Pro;
  font-size: 32px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  font-variant: all-small-caps;
}

@media (max-width: 1050px) {
  span {
    font-size: 24px;
  }
  img {
    width: 350px;
  }
}
@media (max-width: 700px) {
  span {
    font-size: 16px;
  }
}
</style>