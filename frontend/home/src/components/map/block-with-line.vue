<script lang="ts" setup>
  import { StateKeys } from '~/constant/state-keys'

  const rootX = useState(StateKeys.map.makoto.x, () => 0)
  const rootY = useState(StateKeys.map.makoto.y, () => 0)

  const isClick = ref(false)
  const isMounted = ref(false)
  const el = ref<HTMLDivElement>()

  const pos = reactive({
    x: 0,
    y: 0,
  })

  const elCenter = reactive({
    x: 0,
    y: 0,
  })

  const onMouseMove = (e: MouseEvent) => {
    if (!isClick.value || !el.value) return

    el.value.style.left = el.value.offsetLeft - pos.x + e.clientX + 'px'
    el.value.style.top = el.value.offsetTop - pos.y + e.clientY + 'px'

    pos.x = e.clientX
    pos.y = e.clientY

    const { width, height } = el.value.getBoundingClientRect()

    elCenter.x = el.value.offsetLeft + width / 2
    elCenter.y = el.value.offsetTop + height / 2
  }

  const onClick = (e: MouseEvent) => {
    isClick.value = true
    el.value!.style.zIndex = '31'
    pos.x = e.clientX
    pos.y = e.clientY
  }

  const onMouseEnd = () => {
    isClick.value = false
    el.value!.style.zIndex = '30'
  }

  const init = () => {
    const e = el.value!

    const { width, height } = e.getBoundingClientRect()

    const isTop = Math.random() > 0.5

    el.value!.style.top = Number(isTop) * 55 + Math.random() * 35 + '%'
    el.value!.style.left = Math.random() * 80 + 10 + '%'

    elCenter.x = e.offsetLeft + width / 2
    elCenter.y = e.offsetTop + height / 2
  }

  onMounted(() => {
    init()
    isMounted.value = true
    window.addEventListener('resize', init)
  })

  onUnmounted(() => {
    window.removeEventListener('resize', init)
  })
</script>
<template>
  <makoto-line v-if="elCenter.x && elCenter.y && isMounted" :x1="elCenter.x" :y1="elCenter.y" :x2="rootX" :y2="rootY" />
  <div
    ref="el"
    @dragstart.prevent=""
    @drag.prevent=""
    @mouseup="onMouseEnd"
    @mousedown="e => onClick(e)"
    @mousemove="e => onMouseMove(e)"
    @mouseleave="onMouseEnd"
    class="absolute cursor-grab z-30 select-none">
    <slot />
  </div>
</template>
