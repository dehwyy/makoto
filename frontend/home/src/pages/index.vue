<script setup lang="ts">
  import { Button as ButtonMakoto } from 'makoto-ui-vue'

  const { t } = useI18n({
    useScope: 'local',
  })
</script>
<template>
  <header-makoto />
  <main class="main">
    <section class="flex items-center flex-col-reverse xl:flex-row justify-center xl:justify-between w-full md:w-[60%] mx-auto text-white">
      <article class="pb-1 flex flex-col items-center xl:items-start">
        <x-translate direction="left">
          <h1 class="font-Jua text-7xl md:text-8xl lg:text-9xl">Makoto</h1>
        </x-translate>

        <x-translate direction="right" :delay="250">
          <h2 data-cy="subheading" class="font-[600] text-4xl md:text-5xl lg:text-[3.2rem] font-ContentT">{{ t('subheading') }}</h2>
        </x-translate>

        <nav class="mt-5 w-full lg:w-[75%]">
          <nuxt-link href="/map">
            <x-translate class="font-[600] font-ContentT text-lg lg:text-2xl" direction="right" :delay="750">
              <button-makoto :on-click="() => {}">
                <span>{{ t('button') }}</span>
              </button-makoto>
            </x-translate>
          </nuxt-link>
        </nav>
      </article>

      <article>
        <x-translate direction="left" :delay="1250">
          <h2 data-cy="logo" class="font-Kanji text-[12rem] lg:text-[15rem] font-[600] select-none">誠</h2>
        </x-translate>
      </article>
    </section>
  </main>
</template>

<i18n lang="json" scoped>
{
  "en": {
    "subheading": "All-in-one",
    "button": "Investigate"
  },
  "ru": {
    "subheading": "Всё-в-одном",
    "button": "Исследовать"
  }
}
</i18n>

<style lang="scss">
  .main {
    background: url('/static/images/gradient-black.svg') no-repeat;
    background-position: 100% 0%;
    display: flex;
    flex: 1;
    height: 100vh;
    overflow-y: scroll;
    overflow-x: hidden;

    @media (min-height: 800px) {
      background-size: cover;
    }
  }
</style>
