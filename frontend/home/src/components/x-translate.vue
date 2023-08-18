<script lang="ts" setup>
  // Props

  interface Props {
    direction: 'left' | 'right'
    delay?: number
  }
  const props = withDefaults(defineProps<Props>(), {
    direction: 'left',
    delay: 0,
  })

  // State
  const isMounted = ref(false)

  const transitionDelay = computed(() => {
    return (props.delay ?? 0) + 'ms'
  })
  const transformTranslation = computed(() => {
    return `translateX(${props.direction === 'left' ? '-100%' : '100%'})`
  })

  // update on mount
  onMounted(() => {
    isMounted.value = true
  })
</script>

<template>
  <div :class="[isMounted ? 'mounted' : 'before-mount', 'default']">
    <slot />
  </div>
</template>

<style lang="scss" scoped>
  .mounted {
    transform: translateX(0);
    opacity: 100;
    visibility: visible;
  }

  .before-mount {
    opacity: 0;
    visibility: hidden;
    transform: v-bind(transformTranslation);
  }

  .default {
    transition-delay: v-bind(transitionDelay);
    transition-duration: 1000ms;
    transition-property: all;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  }
</style>
