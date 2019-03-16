<template>
  <div class="window-wrapper">
    <div class="window-status">status</div>
    <div class="window-content">content</div>
    <div class="window-input">input</div>
  </div>
</template>

<style scoped>
.window-wrapper{
  border-style:solid;
  border-width: 1px;
  flex-grow: 1;

  flex-direction: column;
}

.window-status{
  border-bottom-style:solid;
  border-bottom-width: 1px;
  flex: 0 1 20px;
}

.window-content{
  border-bottom-style:solid;
  border-bottom-width: 1px;
  flex: 1 1;
}

.window-input{
  flex: 0 1 100px;
}
</style>
