import { Jua, Yuji_Boku, Comfortaa, VT323 } from 'next/font/google'

const jua = Jua({
  weight: '400',
  subsets: ['latin'],
  fallback: ['sans-serif'],
  preload: true,
  variable: '--font-jua',
})

const contentFont = Comfortaa({
  weight: ['400', '700', '600', '500', '300'],
  subsets: ['cyrillic', 'cyrillic-ext', 'latin', 'latin-ext'],
  preload: true,
  variable: '--font-content',
})

const kanji = Yuji_Boku({
  weight: ['400'],
  subsets: ['cyrillic', 'latin', 'latin-ext'],
  variable: '--font-kanji',
})

const pixel = VT323({
  weight: ['400'],
  subsets: ['latin', 'latin-ext'],
  variable: '--font-pixel',
})

export default [contentFont.variable, kanji.variable, jua.variable, pixel.variable].join(' ')
