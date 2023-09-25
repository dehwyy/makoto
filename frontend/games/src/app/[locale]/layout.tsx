import '@/styles/global.css'
import Providers from '@/components/Providers'

// fonts
import fonts from '@/lib/fonts'
import '@fontsource-variable/shantell-sans'

const Layout = ({ children }: { children: React.ReactNode }) => {
  return (
    <html lang="en" className={`text-white font-Content`} suppressHydrationWarning>
      <body className={`${fonts} bg-base-300`}>
        <Providers>
          <main className="w-full max-h-screen min-h-screen h-full">{children}</main>
        </Providers>
      </body>
    </html>
  )
}

export default Layout
