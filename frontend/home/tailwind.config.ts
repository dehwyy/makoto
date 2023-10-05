import type { Config } from 'tailwindcss'
import defaultTheme from 'tailwindcss/defaultTheme'

export default <Partial<Config>>{
  darkMode: 'class',
  content: ['./src/**/*.vue'],
  theme: {
    extend: {
      colors: {
        'white-darker': '#f7f7f7',
      },
      fontFamily: {
        Jua: ['Jua', ...defaultTheme.fontFamily.sans],
        Content: ['Confortaa', ...defaultTheme.fontFamily.sans],
        CherryBomb: ['var(--font-cherry)'],
        Kanji: ['Klee One', ...defaultTheme.fontFamily.sans],
        ContentT: ['Shantell Sans', ...defaultTheme.fontFamily.sans],
      },
      backgroundImage: {
        sun: 'url(/static/png/switch/sun.png)',
        sunBg: 'url(/static/png/switch/sun-background.png)',
        moon: 'url(/static/png/switch/moon.png)',
        moonBg: 'url(/static/png/switch/moon-background.png)',
      },
    },
  },
  daisyui: {
    themes: [
      {
        mytheme: {
          primary: '#4629f2',
          secondary: '#1d1748',
          accent: '#1D082C',
          neutral: '#3e3f40',
          'base-100': '#1d232a',
          'base-300': '#000000',
          info: '#8e97fd',
          success: '#36d399',
          warning: '#fbbd23',
          error: '#f87272',
        },
      },
    ],
  },
  plugins: [require('daisyui')],
}
