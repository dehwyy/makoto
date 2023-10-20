'use client'
import { PORTS } from '@makoto/config'
import Link from 'next/link'
import { forwardRef } from 'react'
import { cn } from '$/lib/utils'
import {
  NavigationMenu,
  NavigationMenuContent,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  NavigationMenuTrigger,
  navigationMenuTriggerStyle,
} from '$/components/ui/navigation-menu'
import ThemeToggler from '$/components/theme-toggler'

const services: { title: string; href: string; description: string }[] = [
  {
    title: 'Hashmap',
    href: `http://localhost:${PORTS.HASHMAP}`,
    description: "A service to store key-value pairs which support features like searching, sorting, editing and viewing other users' data",
  },
  {
    title: 'Intergrations/Discord',
    href: '/docs/primitives/hover-card',
    description: 'Discord integration with Makoto which allows users to use Makoto as a bot for server.',
  },
]

export default function Navigation() {
  return (
    <div className="fixed right-0 left-0 top-0 flex justify-between px-5 py-3 gap-x-5 dark:bg-black bg-white">
      <NavigationMenu>
        <NavigationMenuList className="flex gap-x-3">
          <NavigationMenuItem>
            <NavigationMenuTrigger>Getting started</NavigationMenuTrigger>
            <NavigationMenuContent>
              <ul className="grid gap-3 p-6 md:w-[400px] lg:w-[500px] lg:grid-cols-[.75fr_1fr]">
                <li className="row-span-3">
                  <NavigationMenuLink asChild>
                    <a
                      className="flex h-full w-full select-none flex-col justify-end rounded-md bg-gradient-to-b from-muted/50 to-muted p-6 no-underline outline-none focus:shadow-md"
                      href="/">
                      <div className="font-Kanji font-[600] text-8xl h-full w-full flex flex-col justify-center">誠</div>
                      <div className="mb-2 font-[600] font-ContentT text-2xl">Makoto</div>
                      <p className="leading-tight font-Content">Multitool for web</p>
                    </a>
                  </NavigationMenuLink>
                </li>
                <ListItem href={`http://localhost:${PORTS.AUTH}`} title="Authorize">
                  Login using Google, Discord or create a new Makoto account.
                </ListItem>
                <ListItem href="/" title="Customize">
                  Make your new Makoto profile look unique.
                </ListItem>
                <ListItem href="/" title="Use">
                  Makoto provides a lot of useful services which you can enjoy.
                </ListItem>
              </ul>
            </NavigationMenuContent>
          </NavigationMenuItem>
          <NavigationMenuItem>
            <NavigationMenuTrigger>Services</NavigationMenuTrigger>
            <NavigationMenuContent>
              <ul className="grid w-[400px] gap-3 p-4 md:w-[500px] md:grid-cols-2 lg:w-[600px] ">
                {services.map(component => (
                  <ListItem key={component.title} title={component.title} href={component.href}>
                    {component.description}
                  </ListItem>
                ))}
              </ul>
            </NavigationMenuContent>
          </NavigationMenuItem>
          <NavigationMenuItem>
            <Link href="/integrations" legacyBehavior passHref>
              <NavigationMenuLink className={navigationMenuTriggerStyle()}>Integrations</NavigationMenuLink>
            </Link>
          </NavigationMenuItem>
        </NavigationMenuList>
      </NavigationMenu>
      <ThemeToggler />
    </div>
  )
}

const ListItem = forwardRef<React.ElementRef<'a'>, React.ComponentPropsWithoutRef<'a'>>(({ className, title, children, ...props }, ref) => {
  return (
    <li>
      <NavigationMenuLink asChild>
        <a
          ref={ref}
          className={cn(
            'block select-none space-y-1 rounded-md p-3 leading-none no-underline outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground',
            className,
          )}
          {...props}>
          <div className="text-sm font-medium leading-none">{title}</div>
          <p className="line-clamp-2 text-sm leading-snug text-muted-foreground">{children}</p>
        </a>
      </NavigationMenuLink>
    </li>
  )
})
ListItem.displayName = 'ListItem'
