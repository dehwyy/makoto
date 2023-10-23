import { createI18nMiddleware } from 'next-international/middleware'
import { NextRequest } from 'next/server'

const default_locale = 'en'

const I18nMiddleware = createI18nMiddleware({
  locales: ['en', 'ru'],
  defaultLocale: default_locale,
  urlMappingStrategy: 'rewrite',
  // @ts-ignore
  resolveLocaleFromRequest: ({ cookies }) => {
    const makoto_locale = cookies.get('makoto_locale')
    if (!makoto_locale) {
      cookies.set('makoto_locale', default_locale)
      return default_locale
    }

    return makoto_locale.value
  },
})

export function middleware(request: NextRequest) {
  return I18nMiddleware(request)
}

export const config = {
  matcher: ['/((?!api|static|.*\\..*|_next|favicon.ico|robots.txt).*)'],
}
