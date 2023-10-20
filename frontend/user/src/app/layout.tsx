import type { Metadata } from 'next'
import './globals.css'

import '@fontsource-variable/comfortaa'
import '@fontsource/klee-one'
import '@fontsource/jua'
import '@fontsource-variable/shantell-sans'
import { ThemeProvider } from '$/components/theme-provider'
import Navigation from '$/components/navigation'

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
    <html lang="en">
      <body>
        <ThemeProvider attribute="class" defaultTheme="system" enableSystem disableTransitionOnChange>
          <Navigation />
          {children}
        </ThemeProvider>
      </body>
    </html>
  )
}
