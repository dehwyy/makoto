<script lang="ts" setup>
  import { StateKeys } from '~/constant/state-keys'
  import { servicesUrl as s } from '../../../../.config/services-url'

  const isOpen = useState(StateKeys.menu.isOpen, () => false)
  const isAuth = ref(false)

  const { t } = useI18n({
    useScope: 'local',
  })
  const Hexagons = computed(() =>
    [
      { href: s.home, label: 'home' },
      { href: isAuth.value ? s.user : s.auth, label: isAuth.value ? 'user' : 'login' },
      { href: s.chat, label: 'chat' },
      { href: s.logout, label: 'logout' },
      { href: s.games, label: 'games' },
      { href: s.codespace, label: 'codespace' },
    ].filter(x => {
      // if label is not "user" or "logout" return Val
      if (!['user', 'logout'].includes(x.label)) return x
      // if not Authed and lable is logout
      else if (isAuth.value && x.label === 'logout') return x
    })
  )

  // for CSS binding
  const clipPath = computed(() => (isOpen.value ? 'circle(75%)' : 'circle(25px at calc(100% - 45px) 45px)'))
</script>
<template>
  <div data-cy="menu" class="opacity-100 transition-all delay-300 duration-500 absolute">
    <header-icon-menu-close />
    <div :class="[isOpen ? 'bg-base-300' : '', ' wrapper']">
      <div class="logo flex flex-col gap-y-5">
        <nuxt-link href="/">
          <header-logo />
        </nuxt-link>
        <header-locale-changer />
      </div>
      <div class="border-t-4 min-h-0 border-primary h-full flex items-center justify-center overflow-y-scroll overflow-x-hidden content-wrapper">
        <div class="pt-10 ml-10 lg:ml-0 lg:w-min lg:h-min grid lg:grid-flow-col lg:grid-rows-2 lg:grid-cols-none lg:gap-0 gap-y-6 grid-cols-2 low">
          <div v-for="hex in Hexagons" class="hexagon-item">
            <div v-for="_ in 2" class="hex-item">
              <div v-for="__ in 3"></div>
            </div>
            <a :href="hex.href">
              <p class="hex-content pb-10">
                <span class="flex items-center h-full justify-center svg-wrapper">
                  <svgo-welcome v-if="hex.label === 'home'" />
                  <svgo-chat v-else-if="hex.label === 'chat'" />
                  <svgo-games v-else-if="hex.label === 'games'" />
                  <svgo-codespace v-else-if="hex.label === 'codespace'" />
                  <svgo-logout v-else-if="hex.label === 'logout' && !isAuth" />
                  <svgo-user v-else />
                </span>

                <span class="hex-content-inner font-ContentT font-[600] text-xl title pt-14">{{ t(hex.label) }}</span>
                <svg viewBox="0 0 174 200" height="200" width="174" version="1.1" xmlns="http://www.w3.org/2000/svg">
                  <path d="M87 0L174 50L174 150L87 200L0 150L0 50Z" fill="#1d232a"></path>
                </svg>
              </p>
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<i18n lang="json" scoped>
{
  "en": {
    "home": "Home",
    "user": "Profile",
    "chat": "Chat",
    "login": "Login",
    "logout": "Logout",
    "games": "Games",
    "codespace": "CodeSpace"
  },
  "ru": {
    "home": "Главная",
    "user": "Профиль",
    "chat": "Чат",
    "login": "Войти",
    "logout": "Выйти",
    "games": "Игры",
    "codespace": "CodeSpace"
  }
}
</i18n>

<style lang="scss" scoped>
  .wrapper {
    position: fixed;
    top: 0;
    left: 0;
    height: 100%;
    width: 100%;
    clip-path: v-bind(clipPath);
    transition: all 0.3s ease-in-out;
    transition-delay: 0.2s;
  }

  .logo {
    z-index: 50;
    position: fixed;
    top: 20px;
    left: 20px;
  }

  .menu-btn {
    position: absolute;
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 2;
    right: 20px;
    top: 20px;
    height: 50px;
    width: 50px;
    color: #fff;
    cursor: pointer;
    transition: all 0.3s ease-in-out;
  }

  @keyframes rotate {
    0% {
      filter: hue-rotate(0deg);
    }
    100% {
      filter: hue-rotate(360deg);
    }
  }

  .hexagon-item {
    cursor: pointer;
    width: 200px;
    height: 174px;
    float: left;
    margin-left: -29px;
    z-index: 0;
    position: relative;
    transform: rotate(30deg);
  }
  .hexagon-item:hover {
    z-index: 1;
  }
  .hexagon-item:hover .hex-item:last-child {
    opacity: 1;
    transform: scale(1.3);
  }
  .hexagon-item:hover .hex-item:first-child {
    opacity: 1;
    transform: scale(1.2);
  }
  .hexagon-item:hover .hex-item:first-child div:before,
  .hexagon-item:hover .hex-item:first-child div:after {
    height: 5px;
  }
  .hexagon-item:hover .hex-item div::before,
  .hexagon-item:hover .hex-item div::after {
    background-color: #4629f2;
  }

  .hex-item {
    position: absolute;
    top: 0;
    left: 50px;
    width: 100px;
    height: 170px;
  }
  .hex-item:first-child {
    z-index: 0;
    transform: scale(0.89);
    transition: all 0.3s cubic-bezier(0.165, 0.84, 0.44, 1);
  }
  .hex-item:last-child {
    transition: all 0.3s cubic-bezier(0.19, 1, 0.22, 1);
    z-index: 1;
  }
  .hex-item div {
    box-sizing: border-box;
    position: absolute;
    top: 0;
    width: 100px;
    height: 174px;
    transform-origin: center center;
  }
  .hex-item div::before,
  .hex-item div::after {
    background-color: #1e2530;
    content: '';
    position: absolute;
    width: 100%;
    height: 3px;
    transition: all 0.3s cubic-bezier(0.165, 0.84, 0.44, 1) 0s;
  }
  .hex-item div:before {
    top: 0;
  }
  .hex-item div:after {
    bottom: 0;
  }
  .hex-item div:nth-child(1) {
    transform: rotate(0deg);
  }
  .hex-item div:nth-child(2) {
    transform: rotate(60deg);
  }
  .hex-item div:nth-child(3) {
    transform: rotate(120deg);
  }

  .hex-content {
    color: #fff;
    display: block;
    height: 180px;
    margin: 0 auto;
    position: relative;
    text-align: center;
    transform: rotate(-30deg);
    width: 156px;
  }
  .hex-content .hex-content-inner {
    left: 50%;
    margin: -3px 0 0 2px;
    position: absolute;
    top: 50%;
    transform: translate(-50%, -50%);
  }

  .hexagon-item:hover .svg-wrapper > svg {
    fill: #4629f2;
    stroke: #4629f2;
    transition: 0.6s;
  }
  .svg-wrapper > svg {
    stroke-width: 0;
    width: 60px;
    height: 60px;
    fill: white;
  }

  .hex-content > svg {
    left: -7px;
    position: absolute;
    top: -13px;
    transform: scale(0.87);
    z-index: -1;
    transition: all 0.3s cubic-bezier(0.165, 0.84, 0.44, 1) 0s;
  }

  .hexagon-item:nth-child(even) {
    transform: rotate(30deg) translate(62px, -63px);
  }

  .hexagon-item:hover .title {
    animation: focus-in-contract 0.5s cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
  }

  @media (max-width: 350px) {
    .low {
      grid-gap: 0px;
      grid-template-columns: 1fr;
      width: 100%;
      margin-left: 20px;
    }
    .hexagon-item:nth-child(even) {
      transform: rotate(30deg) translate(63px, -36px) !important;
      margin-left: 0px;
    }
    .content-wrapper {
      align-items: start;
    }
  }

  @media (max-width: 1023px) and (min-width: 350px) {
    .hexagon-item:nth-child(even) {
      transform: rotate(30deg) translate(50px, 86px);
    }
  }

  @media (min-width: 1023px) and (max-height: 350px) {
    .content-wrapper {
      align-items: start;
    }
  }

  @media (max-width: 1023px) and (min-width: 350px) and (max-height: 800px) {
    .content-wrapper {
      align-items: start;
    }
  }

  @keyframes focus-in-contract {
    0% {
      letter-spacing: 0.7em;
      filter: blur(10px);
      opacity: 0;
    }
    100% {
      filter: blur(0px);
      opacity: 1;
    }
  }
</style>
