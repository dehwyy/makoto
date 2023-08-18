<template>
  <div class="line" :style="lineStyle">
    <div class="pulse right-0 bottom-0 h-[2px] w-[50px]"></div>
  </div>
</template>

<script lang="ts" setup>
  import { CSSProperties } from 'nuxt/dist/app/compat/capi'

  const props = defineProps<{
    x1: number
    y1: number
    x2: number
    y2: number
  }>()

  const overallLineLength = ref(0)
  const randomPulseInterval = ref(Math.random() * 10000 + 2000)

  const computedLineLength = computed(() => overallLineLength.value + 'px')
  const computedPulseInterval = computed(() => randomPulseInterval.value + 'ms')

  const lineStyle = computed((): CSSProperties => {
    // chatGPT wrote this file (bruh)
    const deltaX = props.x2 - props.x1
    const deltaY = props.y2 - props.y1
    const length = Math.sqrt(deltaX * deltaX + deltaY * deltaY)
    const angle = Math.atan2(deltaY, deltaX) * (180 / Math.PI)

    overallLineLength.value = length - 150

    return {
      width: `${length - 150}px`,
      height: '2px',
      top: `${props.y1}px`,
      left: `${props.x1}px`,
      transform: `rotate(${angle}deg)`,
    } as const
  })
</script>

<style>
  .line {
    position: absolute;
    background-color: #151617;
    transform-origin: 0% 50%;
  }

  .pulse {
    --color-one: #ff00d2;
    --color-two: #00a2ff;
    position: absolute;
    bottom: 0;
    box-shadow: inset 0 0 30px white, inset 10px 0 40px var(--color-one), inset -10px 0 40px var(--color-two), inset 10px 0 150px var(--color-one),
      inset -10px 0 150px var(--color-two), 0 0 50px #fff, -10px 0 40px var(--color-one), 10px 0 40px var(--color-two);
    animation: pulse-animation infinite;
    animation-duration: v-bind(computedPulseInterval);
    transform-origin: 0% 50%;
    opacity: 0;
  }

  @keyframes pulse-animation {
    0% {
      left: v-bind(computedLineLength);
    }
    1% {
      opacity: 100;
    }
    50% {
      left: 0;
    }
    51% {
      opacity: 0;
    }
    100% {
      opacity: 0;
    }
  }
</style>
