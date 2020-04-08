<template>
    <div>
      <input v-model="inputValue"/>
      <button @click="clickHandler"></button>
    </div>
</template>

<script>
  export default {
    name: "TodoList",
    data: function() {    /* 或者： data(){*/
      return {
        inputValue: "inputValue",
        list: [],
      }
    },
    methods:{
      clickHandler() {  /* 或者 clickHandler: function(){*/
        this.list.push((this.inputValue))
      }
    }
  }
</script>

<style scoped>

</style>
