<template>
  <div data-cy="glow-block" class="flex justify-center">
    <div class="card relative">
      <div
        class="face p-5 rounded-2xl border-2 min-h-[100px] min-w-[150px] transition-all duration-500 flex justify-center items-center z-[1] relative bg-base-300">
        <h2 class="font-ContentT text-3xl font-[700] text-white">
          <slot name="title" />
        </h2>
      </div>
      <div class="pt-[1rem]">
        <div class="content border-2 rounded-2xl transition-all duration-[400ms] bg-base-300 absolute top-0 left-0 right-0">
          <div class="px-5">
            <p class="font-Content text-sm font-[600] text-white text-center">
              <slot name="content" />
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
  .card:hover {
    & .face {
      --color-one: #ff00d2;
      --color-two: #00a2ff;
      box-shadow: inset 0 0 30px white, inset 10px 0 40px var(--color-one), inset -10px 0 40px var(--color-two), inset 10px 0 150px var(--color-one),
        inset -10px 0 150px var(--color-two), 0 0 50px #fff, -10px 0 40px var(--color-one), 10px 0 40px var(--color-two);
    }

    & .content {
      top: 100%;
      padding-top: 1rem;
      padding-bottom: 1rem;
    }
  }
</style>
