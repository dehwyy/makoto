'use client'
import { I18nProviderClient } from '@/locales/client'

const Providers = ({children}: {children: React.ReactNode}) => {
  return <>
  <I18nProviderClient>
      {children}
    </I18nProviderClient>
  </>
}

export default Providers
