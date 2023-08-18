// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ['@nuxtjs/tailwindcss', '@nuxtjs/i18n', 'nuxt-svgo', '@nuxtjs/google-fonts'],
  components: {
    dirs: ['~/components', '~/components/map', '~/components/header', '~/components/index'],
  },
  svgo: {},
  devServer: {
    port: 3000,
  },
  googleFonts: {
    preload: true,
    subsets: ['cyrillic', 'latin'],
    families: {
      Jua: [400],
      Comfortaa: [400, 600, 700],
      'Klee One': [600],
      'Shantell Sans': [400, 500, 600, 700],
    },
  },
  i18n: {
    locales: ['ru', 'en'],
    defaultLocale: 'en',
  },
  srcDir: 'src',
  css: ['~/assets/global.scss'],
})
