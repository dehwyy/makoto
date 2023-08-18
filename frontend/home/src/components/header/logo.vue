<template>
  <p class="text-fade z-50 shadow-black drop-shadow-lg font-Kanji font-[600] text-6xl bg-gradient-to-r from-error via-info to-error">誠</p>
</template>
<style scoped>
  .text-fade {
    color: transparent;
    background-clip: text;
    background-size: 200% auto;
    animation-name: bg-pseudo-translate;
    animation-duration: 10s;
    animation-timing-function: linear;
    animation-iteration-count: infinite;
  }

  @keyframes bg-pseudo-translate {
    from {
      background-position: -100% center;
    }
    to {
      background-position: 100% center;
    }
  }
</style>
