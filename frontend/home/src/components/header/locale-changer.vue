<script lang="ts" setup>
  const LOCALES = ['ru', 'en']

  const cookieLocale = useCookie('makoto_locale')

  const isExpanded = ref(false)
  const currentLocale = ref(cookieLocale.value)

  const clickHandler = (l: string) => {
    if (l !== currentLocale.value) {
      currentLocale.value = l
      cookieLocale.value = l
      setTimeout(() => {
        location.reload()
      }, 150)
    }
    isExpanded.value = !isExpanded.value
  }
</script>
<!-- for 2nd div @line:24 -->
<!-- if current locales is selected: TOP -->
<!--     if expanded => show other icons -->
<!--     else => only current -->
<!-- default styles -->
<template>
  <div class="locale-changer-wrapper">
    <div
      v-for="l in LOCALES"
      @click="clickHandler(l)"
      :class="[
        currentLocale === l ? 'top-0 z-20' : isExpanded ? 'top-[55px] opacity-100 visible' : 'top-0 opacity-0 invisible',
        'transition-all cursor-pointer absolute left-0 right-0',
      ]">
      <svgo-flags-ru v-if="l === 'ru'" />
      <svgo-flags-uk v-else />
    </div>
  </div>
</template>

<style>
  .locale-changer-wrapper {
    display: flex;
    flex-direction: column;
    justify-items: center;
    position: relative;
    height: 50px;
    width: 50px;
    margin: 0 auto;
  }

  .nuxt-icon {
    width: auto;
    height: auto;
    margin-bottom: 0;
  }
</style>
