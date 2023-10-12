<script lang="ts" setup>
  import { StateKeys } from '~/constant/state-keys'

  const isOpen = useState(StateKeys.menu.isOpen, () => false)
</script>
<template>
  <div class="relative z-50 left-6 top-4">
    <input @click="isOpen = !isOpen" class="input" type="checkbox" />
    <div v-for="_ in 2" class="wrap-icon">
      <svg v-for="__ in 3" class="svg" viewBox="0 0 820 864" x="0px" xmlns="http://www.w3.org/2000/svg" y="0px">
        <g>
          <g>
            <path
              class="path stroke-white"
              d="M341.8,326.1h213c89.9-9.3,148.7-84.1,140-152c-7.6-58.9-67.1-121.3-133-109c-65.2,12.2-103.6,91.5-87,133
			c15.6,39.1,72.5,25.7,104,73c34.8,52.2,4.4,126.9,1,135c-38.6,91.3-166.6,159.9-283,120c-83.1-28.5-148.1-108-153-203
			c-5.2-102,61.1-181.2,130-217c105.2-54.7,222.8-11.2,288,42c134.2,109.5,141.4,337.4,56,424c-58.7,59.5-164.8,55.8-217,147
			c-6.4,11.2-13.2,25.9-9,41c12.1,43.6,109.8,70.9,180,42c82.7-34.1,134.7-149.4,97-207c-35.2-53.7-145.1-52.1-210-14
			c-39,22.9-40.8,46.7-85,69c-55.6,28.1-132.6,30.5-146,7c-14.9-26.2,38-103.1,320-320"></path>
          </g>
        </g>
      </svg>
    </div>
  </div>
</template>
<style lang="scss" scoped>
  .wrap-icon {
    top: 0;
    left: 0;
    position: absolute;
    display: flex;
    justify-content: center;
    align-items: center;
    transition: 0.2s cubic-bezier(0.175, 0.885, 0.32, 1.575);
    transition-delay: 1.85s;
    pointer-events: none;
    &:nth-of-type(2) {
      z-index: -1;
      &:before {
        content: '';
        text-transform: uppercase;
        position: absolute;
        color: rgba(255, 255, 255, 0.75);
        z-index: 999;
        font-family: Futura;
        letter-spacing: 2px;
        font-size: 8px;
        left: 5px;
        top: calc(50% + 40px);
        text-align: center;
        transition: transform 0.4s ease-in-out, opacity 0.4s ease-in-out, letter-spacing 0.3s ease-in-out;
        transition-delay: 2.5s, 2.5s, 0s;
        opacity: 1;
      }
      .svg {
        path {
          transition-delay: 0s, 0s;
          stroke: #fff;
        }
        &:nth-of-type(2) {
          path {
            transition-delay: 0.75s, 0.75s;
          }
        }
        &:nth-of-type(3) {
          path {
            transition-delay: 1.5s, 1.5s;
          }
        }
      }
    }
  }
  .input {
    position: absolute;
    width: 30px;
    height: 30px;
    z-index: 999;
    opacity: 0;
    cursor: pointer;
    left: 50%;
    top: 0%;
    transform: translateX(-50%) translateY(-50%);
    &:hover ~ .wrap-icon {
      &:before {
        transition-delay: 2.5s, 2.5s, 0s;
        letter-spacing: 5px;
      }
    }
    &:checked ~ .wrap-icon {
      transform: scaleX(0.9);
      &:before {
        opacity: 0;
        transform: translateY(25vh);
        transition-delay: 0s, 0s, 0s;
      }
      &:nth-of-type(2) {
        .svg {
          .path {
            stroke: #fff;
          }
        }
      }
      .svg {
        transition-delay: 0s;
        &:first-of-type {
          transform: translateY(0px);
          .path {
            transition-delay: 0.5s, 0.5s;
          }
        }
        &:nth-of-type(2) {
          transform: translateY(0px) scaleY(-1);
          .path {
            transition-delay: 0.5s, 0.5s;
          }
        }
        &:nth-of-type(3) {
          transform: translateY(0);
          .path {
            stroke-dasharray: 10000px;
            stroke-dashoffset: 10000px;
            transition-delay: 1s, 1s;
          }
        }
        .path {
          stroke-dasharray: 5000px;
          stroke-dashoffset: -4150px;
          stroke: #f87272;
        }
      }
    }
  }

  .svg {
    position: absolute;
    width: 90px;
    transition: 0.2s cubic-bezier(0.175, 0.885, 0.32, 1.575);
    transition-delay: 2s;
    margin-left: -9px;
    &:first-of-type {
      transform: translateY(0);
    }
    &:nth-of-type(2) {
      transform: translateY(-13px) scaleY(-1);
      .path {
        transition: stroke-dasharray 1.2s ease-in-out, stroke-dashoffset 1.2s ease-in-out, stroke 1.2s ease-in-out;
        transition-delay: 0.6s, 0.6s, 0.6s;
      }
    }
    &:nth-of-type(3) {
      transform: translateY(20px);
      .path {
        transition: stroke-dasharray 0.25s ease-in-out, stroke-dashoffset 0.25s ease-in-out, stroke 0.25s ease-in-out;
        transition-delay: 1.8s, 1.8s, 1.8s;
      }
    }
    .path {
      stroke-width: 40px;
      fill: none;
      stroke-linecap: round;
      stroke-dasharray: 4400px;
      stroke-dashoffset: 4200px;
      transition: stroke-dasharray 1.5s ease-in-out, stroke-dashoffset 1.5s ease-in-out, stroke 1.5s ease-in-out;
    }
  }
</style>
