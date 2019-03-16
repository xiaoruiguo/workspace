<template>
  <div class="sidebar-wrapper">
    <div class="sidebar-logo">可信发布系统</div>
    <div class="sidebar-list">list</div>
  </div>
</template>

<style scoped>
.sidebar-wrapper{
  border-style:solid;
  border-width: 1px;
  flex-grow: 1;

  flex-direction: column;
}
.sidebar-logo{
  border-bottom-style:solid;
  border-bottom-width: 1px;
  flex-grow: 0;
  flex-basis: 30px;
}
.sidebar-list{
  flex-grow: 1;
}
</style>
