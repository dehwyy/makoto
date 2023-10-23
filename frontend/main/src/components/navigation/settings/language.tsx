import { Select, SelectContent, SelectGroup, SelectItem, SelectTrigger } from '$/components/ui/select'

const placeholders = {
  ru: 'English / English',
  en: 'Russian / Русский',
} as const

const Language = () => {
  const language_value_from_store = 'ru' as keyof typeof placeholders

  return (
    <>
      <p className="font-Content font-[600]">Язык</p>
      <Select>
        <SelectTrigger className="w-full">{placeholders[language_value_from_store]}</SelectTrigger>
        <SelectContent>
          <SelectGroup defaultValue={language_value_from_store} className="font-Content">
            <SelectItem value="english">English / English</SelectItem>
            <SelectItem value="russian">Russian / Русский</SelectItem>
          </SelectGroup>
        </SelectContent>
      </Select>
    </>
  )
}
export default Language
