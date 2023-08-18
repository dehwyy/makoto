<script lang="ts" setup>
  import { useState } from 'nuxt/app'
  import { StateKeys } from '~/constant/state-keys'

  const isOpen = useState(StateKeys.menu.isOpen, () => false)
  const timeout = ref<NodeJS.Timeout | null>(null)
  const isOpenWithDelay = ref(false)

  watch(isOpen, v => {
    timeout.value = null
    // if on open: immediately hide all underlying stuff
    if (v) {
      isOpenWithDelay.value = v
    }
    // else: only after animation show that stuff
    else {
      timeout.value = setTimeout(() => {
        timeout.value && (isOpenWithDelay.value = v)
      }, 1500)
    }
  })
</script>
<template>
  <header :class="[isOpenWithDelay ? 'z-40' : 'z-30']" class="fixed top-0 left-0 right-0 p-3 h-[95px]">
    <div class="w-full lg:w-[80%] h-full mx-auto flex justify-between items-center select-none">
      <nav>
        <nuxt-link href="/">
          <!-- <makoto-logo /> -->
        </nuxt-link>
      </nav>
      <section class="flex gap-x-1 lg:gap-x-5 pr-16 lg:pr-0 min-h-[30px]">
        <header-menu />
      </section>
    </div>
  </header>
</template>
