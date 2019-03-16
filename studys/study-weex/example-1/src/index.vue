<template>
  <div class="wrapper">
    <div class="sidebar">sidebar</div>
    <div class="window"><Info/></div>
  </div>
</template>

<script>
import SideBar from './components/SideBar.vue'
import Window from './components/Window.vue'
import Info from './components/Info.vue'
export default {
  name: 'App',
  components: {
    SideBar,
    Window,
    Info
  },
  data () {
    return {
      logo: 'https://gw.alicdn.com/tfs/TB1yopEdgoQMeJjy1XaXXcSsFXa-640-302.png',
      object: weex.config.env
    }
  }
}
</script>

<style scoped>
  .wrapper {
    flex-direction: row;
    justify-content:flex-start;
    align-items: stretch;
    flex-grow:1;
  }
  .sidebar{
    border-style:solid;
    border-width: 1px;
    flex-grow:1;
    flex-basis: 50px;
    width: 150px;
  }
  .window{
    border-style:solid;
    border-width: 1px;
    flex-grow:4;
    flex-basis: 100px;
    width: 600px;
  }
</style>
