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

        <nav class="mt-5 w-[75%]">
          <nuxt-link href="/map">
            <x-translate class="font-[600] font-ContentT text-2xl" direction="right" :delay="750">
              <button-makoto :on-click="() => {}">
                <span>{{ t('button') }}</span>
              </button-makoto>
            </x-translate>
          </nuxt-link>
        </nav>
      </article>

      <article>
        <x-translate direction="left" :delay="1250">
          <h2 data-cy="logo" class="font-Kanji text-[15rem] font-[600] select-none">誠</h2>
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

  .cut-top {
    position: absolute;
    top: 12px;
    display: inline-block;
    clip-path: polygon(0% 66%, 0% 0%, 100% 0%, 100% 40%);
    transition: transform 0.6s cubic-bezier(0.16, 1, 0.3, 1) 0.25s, color 0.5s ease 0.22s;
  }

  .cut-top:after {
    content: '';
    position: absolute;

    top: 40%;
    left: 0;
    width: 100%;
    height: 4px;
    background: #ff00d2;
    transform: rotateZ(-3deg) scaleX(0%);
    transform-origin: right top;
    transition: transform 0.2s ease 0.22s;
  }

  .cut-bottom {
    display: inline-block;
    clip-path: polygon(0% 65%, 100% 40%, 100% 110%, 0% 110%);
    text-decoration: underline;
    transition: background-position 0.2s ease 0.22s;
    text-decoration: none;
    background-size: 200% 8%;
    background-position: left bottom;
    background-repeat: no-repeat;
    background-image: linear-gradient(to right, white 50%, transparent 50%);
  }

  .cut-bottom:after {
    content: '';
    position: absolute;
    background: #ff00d2;
    top: 10%;
    left: 0;
    width: 100%;
    height: 4px;
    transform: rotateZ(-2deg) scaleX(0%);
    transform-origin: right top;
    transition: transform 0.2s ease 0.22s;
  }

  .cut:hover {
    & .cut-top {
      transform: translateY(-0.5em) rotateZ(-3deg);
      transition: transform 0.5s cubic-bezier(0.12, 0.8, 0.57, 1) 0.42s, color 0.5s ease 0.22s;
    }

    & .cut-bottom {
      background-position: 100% bottom;
    }
    & .cut-top:after {
      top: 60%;
      transform-origin: left top;
      transform: rotateZ(-3.1deg) scaleX(100%);
    }

    & .cut-bottom:after {
      top: 61%;
      transform-origin: left top;
      transform: rotateZ(-3.1deg) scaleX(100%);
    }
  }
</style>
