import type { Metadata } from 'next'
import './globals.css'

import '@fontsource-variable/comfortaa'
import '@fontsource/klee-one'
import '@fontsource/jua'
import '@fontsource-variable/shantell-sans'
import { ThemeProvider } from '$/components/theme-provider'
import Navigation from '$/components/navigation'

// ! Can't use next/font/goole as I'm getting error (user aborted request)
// import { Comfortaa, Klee_One, Jua, Shantell_Sans } from 'next/font/google'

// const comfortaa = Comfortaa({
//   subsets: ['latin'],
//   variable: '--font-comfortaa',
//   weight: ['400', '500', '600', '700'],
// })
// const kanji = Klee_One({
//   weight: ['400', '600'],
//   variable: '--font-kanji',
//   subsets: ['latin-ext', 'greek-ext'],
//   preload: true,
// })
// const jua = Jua({
//   subsets: ['latin'],
//   variable: '--font-jua',
//   weight: ['400'],
//   preload: true,
// })
// const shantell_sans = Shantell_Sans({
//   subsets: ['latin'],
//   variable: '--font-shantell-sans',
//   weight: ['400', '500', '600', '700'],
//   preload: true,
// })

export const metadata: Metadata = {
  title: 'Makoto',
  description: 'Makoto user microservice',
  icons: {
    icon: ['/meta/favicon.ico?v=4'],
    apple: ['/meta/apple-touch-icon.png?v=4'],
    shortcut: ['/meta/apple-touch-icon.png'],
  },
  manifest: '/meta/site.webmanifest',
}

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body>
        <ThemeProvider attribute="class" defaultTheme="system" enableSystem disableTransitionOnChange>
          <Navigation />
          {children}
        </ThemeProvider>
      </body>
    </html>
  )
}
