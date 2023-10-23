'use client'

import ThemeToggler from '$/components/navigation/theme-toggler'
import NavigationItems from '$/components/navigation/navigation-items'
import Logo from '$/components/navigation/logo'
import Settings from '$/components/navigation/settings'

export default function Navigation() {
  return (
    <div className="fixed right-0 left-0 top-0 flex px-5 py-3 gap-x-14 dark:bg-black bg-white z-10">
      <Logo />
      <NavigationItems />
      <div className="ml-auto flex gap-x-3">
        <Settings />
        <ThemeToggler />
      </div>
    </div>
  )
}
