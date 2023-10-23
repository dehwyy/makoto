'use client'

import { useState } from 'react'
import { Select, SelectContent, SelectGroup, SelectItem, SelectTrigger } from '$/components/ui/select'
import { useChangeLocale, useCurrentLocale } from '$/locales/client'

const placeholders = {
  en: 'English / English',
  ru: 'Russian / Русский',
} as const

const Language = () => {
  const current_locale = useCurrentLocale()
  const set_locale = useChangeLocale({ preserveSearchParams: true })
  const [lang, setLang] = useState(current_locale)

  const onValueChange = (value: keyof typeof placeholders) => {
    setLang(value)
    set_locale(value)
  }

  return (
    <>
      <p className="font-Content font-[600]">Language</p>
      <Select onValueChange={onValueChange}>
        <SelectTrigger className="w-full">{placeholders[lang]}</SelectTrigger>
        <SelectContent>
          <SelectGroup defaultValue={lang} className="font-Content">
            <SelectItem value="en">English / English</SelectItem>
            <SelectItem value="ru">Russian / Русский</SelectItem>
          </SelectGroup>
        </SelectContent>
      </Select>
    </>
  )
}
export default Language
