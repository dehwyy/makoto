import type { Metadata } from 'next'
import './globals.css'

import '@fontsource-variable/comfortaa'
import '@fontsource/klee-one'
import '@fontsource/jua'
import '@fontsource-variable/shantell-sans'
import { ThemeProvider } from '$/components/theme-provider'
import Navigation from '$/components/navigation'

export const metadata: Metadata = {
  title: 'MakotoUser',
  description: 'Makoto user microservice',
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
