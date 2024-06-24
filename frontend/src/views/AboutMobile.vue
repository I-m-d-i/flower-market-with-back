<script>
export default {
  name: "AboutMobile"
}
</script>

<template>
  <div style="    display: flex;
    flex-direction: column;
    align-items: center;
    padding-inline: 10%;
margin-top: 20px">
    <h1 style="align-self: flex-start;">О нас</h1>
    <div class="header-about" style="">
      <img :src="require('@/assets/About2.png')" alt="О нас"/>
      <div class="ellipse">
        МЫ РАБОТАЕМ<br/> ДЛЯ ВАШЕЙ<br/> УЛЫБКИ!
      </div>
    </div>
    <span style="margin-top: 30px;align-self: start; font: 11px Source Serif Pro,SemiBold ; font-weight: 600">
      Студия флористики и декора
    </span>
    <span style="margin-top: 10px; align-self: start; color: #71314C;font: 11px Source Serif Pro,SemiBold; font-weight: 600">
      CARDENIA
    </span>
    <span style="margin-top: 10px">
      ПРИВЕТ, МЫ СТУДИЯ ФЛОРИСТИКИ И ДЕКОРА.<br/>
      НАША КОМАНДА СОСТОИТ ИЗ СПЕЦИАЛИСТОВ В
      ОБЛАСТИ ФЛОРИСТИКИ И ДИЗАЙНА,
      ПРОДАЖ.ЛОГИСТИКИ И МАРКЕТИНГА.
      У НАС ЦАРИТ ОСОБАЯ АТМОСФЕРА, КОТОРАЯ
      ПОМОГАЕТ ПОЛУЧАТЬ ВДОХНОВЕНИЕ ДЛЯ СОЗДАНИЯ
      НОВЫХ ШЕДЕВРОВ ФЛОРИСТИКИ, ДИЗАЙНА И ДЕКОРА. МЫ ЛЮБИМ ТО,
      ЧТО ДЕЛАЕМ, И ЭТО КЛЮЧЕВОЙ ПОКАЗАТЕЛЬ
      НАШЕГО УСПЕХА.
      ЛЮБОВЬ К ЦВЕТАМ И ЛЮДЯМ, ИСКРЕННОСТЬ
      ЭМОЦИЙ И ВПЕЧАТЛЕНИЙ, ГОРЯЩИЕ ГЛАЗА - ВСЕ ТО,
      ЧТО ДЕЛАЕТ НАС СОВЕРШЕННЕЕ ДЕНЬ ОТО ДНЯ.
    </span>
    <span style="margin-top: 40px">
      У НАС ВСЕГДА ПРЕДСТАВЛЕН
      ШИРОКИЙАССОРТИМЕНТ ОТЛИЧНЫХ ЦВЕТОВ.<br/>
      НАШ КАТАЛОГ НАСЧИТЫВАЕТ МНОГО
      ДЕСЯТКОВ НАИМЕНОВАНИЙ БУКЕТОВ, ОЧЕНЬ
      КРАСИВЫХ И ГАРМОНИЧНЫХ, КОТОРЫЕ БУДУТ
      ПРЕКРАСНО СМОТРЕТЬСЯВ ЛЮБОМ ЖИЛИЩЕ
      ИЛИ ОФИСЕ. КРОМЕ БУКЕТОВ, НАШИ КЛИЕНТЫ
      МОГУТ ПРИОБРЕСТИ
      ЦВЕТОЧНЫЕ КОМПОЗИЦИИ В ВИДЕ
      ДЕКОРАТИВНЫХ КОРЗИНОК, ШЛЯПНЫХ
      КОРОБОК.ЕСТЬ ТАКЖЕ «СОПУТСТВУЮЩИЕ»
      ТОВАРЫ: МЯГКИЕ ИГРУШКИ, ШАРЫ С ГЕЛИЕМ,
      ШОКОЛАД.
    </span>
  </div>
</template>

<style scoped>

.header-about {
  display: inline-flex;
  align-items: center;
  justify-content: flex-end;
  padding-inline: 20%
}

img {
  border-radius: 282px;
  width: 100%;
}

span {
  font: 470 10px Source Serif Pro, regular;
}

.ellipse {
  background-color: #FFFFFF;
  border-radius: 81px;
  position: absolute;
  width: 162px;
  height: 162px;
  padding-top: 50px;
  text-align: center;
  margin-right: -81px;
  margin-top: 162px;
  font-family: Source Serid Pro, SemiBold;
  font-size: 14px;
  font-weight: 550;
}

@media (max-width: 700px) {
  .ellipse {
    width: 102px;
    height: 102px;
    margin-right: -51px;
    font-size: 9px;
    margin-top: 102px;
    padding-top: 30px;
  }
  .header-about{
    padding-inline: 10%
  }


}
</style>