import type { Config } from 'tailwindcss'
import defaultTheme from 'tailwindcss/defaultTheme'

const config: Config = {
  content: ['./src/pages/**/*.{js,ts,jsx,tsx,mdx}', './src/components/**/*.{js,ts,jsx,tsx,mdx}', './src/app/**/*.{js,ts,jsx,tsx,mdx}'],
  theme: {
    extend: {
      fontFamily: {
        ContentT: ['Shantell Sans Variable', ...defaultTheme.fontFamily.sans],
        Jua: ['var(--font-jua)'],
        Content: ['var(--font-content)'],
        Kanji: ['var(--font-kanji)'],
      },
    },
  },
  daisyui: {
    themes: [
      {
        mytheme: {
          primary: '#ff00d2',
          secondary: '#00a2ff',
          accent: '#4629f2',
          neutral: '#2a323c',
          'base-100': '#1d232a',
          info: '#7583ca',
          success: '#36d399',
          warning: '#fbbd23',
          error: '#f87272',
        },
      },
    ],
  },
  plugins: [require('daisyui')],
}
export default config
