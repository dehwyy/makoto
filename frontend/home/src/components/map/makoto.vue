<script lang="ts" setup>
  import { StateKeys } from '~/constant/state-keys'

  const rootPositionX = useState(StateKeys.map.makoto.x, () => 0)
  const rootPositionY = useState(StateKeys.map.makoto.y, () => 0)

  const makoto = ref<HTMLDivElement>()

  const setMakotoPosition = () => {
    if (!makoto.value) return

    const { width, height, x, y } = makoto.value.getBoundingClientRect()
    rootPositionX.value = x + width / 2
    rootPositionY.value = y + height / 2
  }

  onMounted(() => {
    setMakotoPosition()

    window.addEventListener('resize', setMakotoPosition)
  })

  onUnmounted(() => {
    window.removeEventListener('resize', setMakotoPosition)
  })
</script>
<template>
  <div class="min-h-screen h-full flex justify-center items-center select-none relative">
    <div ref="makoto" class="title">
      <h1 class="font-Jua text-7xl z-20">
        Makoto
        <div class="text-highlight"></div>
      </h1>
      <div class="text-underline"></div>
      <div aria-hidden class="text-filled">
        <h2 class="font-Jua text-7xl z-20">Makoto</h2>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
  .title {
    color: white;
    position: relative;

    h1 {
      position: relative;
      margin: 0;
      padding: 0;
      overflow: hidden;
      color: #151617;
      -webkit-text-stroke: 2px transparent;
      background-clip: text;
      background-image: linear-gradient(to right, #00a2ff, #ff00d2);
    }

    & .text-underline {
      position: absolute;
      width: 100%;
      left: 0;
      bottom: -3px;
      height: 3px;
      background-image: linear-gradient(to right, #00a2ff, #ff00d2);
      z-index: 5;
      transition: transform 580ms cubic-bezier(0.2, 0.1, 0.15, 1.32);
    }

    & .text-highlight {
      position: absolute;
      width: 100%;
      bottom: -55px;
      left: 0;
      height: 55px;
      background-color: #00a2ff;
      z-index: -1;
      transition: transform 400ms ease;
    }

    & .text-filled {
      position: absolute;
      margin: 0;
      padding: 0;
      top: 0;
      background-clip: text;
      background-image: linear-gradient(to right, #00a2ff, #ff00d2);
      color: transparent;
      z-index: 1000;
      clip-path: polygon(0% 100%, 0% 99%, 100% 99%, 100% 100%);
      transition: clip-path 600ms cubic-bezier(0.2, 0.1, 0.15, 1.32);
    }

    &:hover {
      .text-filled {
        clip-path: polygon(0% 100%, 0% 0%, 100% 0%, 100% 100%);
      }
      .text-underline {
        transform: translateY(-75px);
      }
      .text-highlight {
        transform: translateY(-2.2em);
      }

      h1 {
        -webkit-text-stroke: 0px transparent;
      }
    }
  }
</style>
