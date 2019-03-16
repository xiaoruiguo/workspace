<template>
<scroller>
  <div class="group" v-for="(value,key) in object">
    <text class="text key"> {{key}} </text>
    <text class="text type"> {{typeof value}} </text>
    <text class="text value"> {{value}} </text>
  </div>
</scroller>
</template>

<script>
export default {
  data () {
    return {
      object: weex.config.env
    }
  }
}
</script>

<style scoped>
  .group{
    flex-direction: row;
    justify-content:space-around;
  }
  .text {
    font-size: 25px;
    flex-grow: 0;
  }
  .key {
    width: 200px;
    text-align: right;
    color: #00B4FF;
  }
  .type{
    width: 100px;
    text-align:center;
    color: #A0A0A0;
  }
  .value{
    width: 200px;
    color: #525252;
  }
</style>
