import { ReactElement } from 'react'
import { I18nProviderClient } from '$/locales/client'

export default function SubLayout({ params: { locale }, children }: { params: { locale: string }; children: ReactElement }) {
  return (
    <I18nProviderClient locale={locale} fallback="en">
      {children}
    </I18nProviderClient>
  )
}
