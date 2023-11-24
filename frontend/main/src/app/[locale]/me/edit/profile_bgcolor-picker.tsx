'use client'
import { Button } from '$/components/ui/button'
import { HexColorPicker } from 'react-colorful'
import { Select, SelectContent, SelectGroup, SelectItem, SelectTrigger } from '$/components/ui/select'
import { useTheme } from 'next-themes'
import { useEffect, useState } from 'react'
import { Input } from '$/components/ui/input'

interface Props {
  default_light: string
  default_dark: string

  darkBg: string
  lightBg: string

  setDarkBg: (color: string) => void
  setLightBg: (color: string) => void
}

const ProfileColorPicker = ({ default_light, default_dark, darkBg, lightBg, setDarkBg, setLightBg }: Props) => {
  const { theme: currentTheme, setTheme } = useTheme()
  const [isColorPickerOpen, setIsColorPickerOpen] = useState(false)

  const [initialDarkBg, setInitialDarkBg] = useState(default_dark)
  const [initialLightBg, setInitialLightBg] = useState(default_light)

  const [bgTheme, setBgTheme] = useState<'light' | 'dark'>('dark')

  // I don't exactly understand how it works, but it works xd (eslint hello)
  useEffect(() => {
    isColorPickerOpen ? setTheme(bgTheme) : setBgTheme(currentTheme as 'light' | 'dark')

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [bgTheme])

  const SaveColor = () => {
    setIsColorPickerOpen(false)
    setInitialDarkBg(darkBg)
    setInitialLightBg(lightBg)
  }

  const DiscardChanges = () => {
    setIsColorPickerOpen(false)

    setDarkBg(initialDarkBg)
    setLightBg(initialLightBg)
  }

  return (
    <div className="grid grid-cols-2 gap-5">
      <p className="font-sans text-sm font-[600] mb-2 col-span-2">BACKGROUND COLOR</p>
      <div className="w-full">
        <div className="flex flex-col">
          <Button
            onClick={isColorPickerOpen ? () => SaveColor() : () => setIsColorPickerOpen(p => !p)}
            variant={isColorPickerOpen ? 'default' : 'outline'}
            className="w-full">
            {isColorPickerOpen ? 'Save' : 'Edit'}
          </Button>
          {isColorPickerOpen && (
            <Button
              onClick={DiscardChanges}
              variant="destructive"
              className={`${
                (bgTheme === 'light' && initialLightBg != lightBg) || (bgTheme === 'dark' && initialDarkBg != darkBg)
                  ? 'visible opacity-100 max-h-[41px] mt-2'
                  : 'invisible opacity-0 max-h-[0px] py-0'
              } w-full transition-all`}>
              Discard
            </Button>
          )}
        </div>
        {isColorPickerOpen && (
          <div className="pt-5 flex flex-col gap-y-3 select-none">
            <p className="text-sm font-Content dark:text-gray-300 underline">Background depenends on theme!</p>
            <p className="text-sm dark:text-gray-300 -mb-2 text-center">Select for</p>
            <Select onValueChange={v => setBgTheme(v as 'light' | 'dark')}>
              <SelectTrigger>{bgTheme[0].toUpperCase() + bgTheme.substring(1)}</SelectTrigger>
              <SelectContent>
                <SelectGroup className="font-Content">
                  <SelectItem value="light">Light</SelectItem>
                  <SelectItem value="dark">Dark</SelectItem>
                </SelectGroup>
              </SelectContent>
            </Select>
            <p className="text-sm dark:text-gray-300 text-center">Theme</p>
          </div>
        )}
      </div>
      <div
        className={
          (isColorPickerOpen ? 'visible opacity-100 max-h-[300px]' : 'invisible opacity-0 max-h-[0px]') +
          ' transition-all duraiton-300 ease-in-out w-[200px]'
        }>
        <p className="text-center text-lg underline">Pick a color</p>
        {bgTheme === 'light' ? (
          <HexColorPicker color={lightBg} onChange={setLightBg} />
        ) : bgTheme === 'dark' ? (
          <HexColorPicker color={darkBg} onChange={setDarkBg} />
        ) : (
          <></>
        )}
        <Input
          className="mt-5"
          autoComplete="disabled"
          placeholder={bgTheme === 'light' ? lightBg : darkBg}
          spellCheck={false}
          value={bgTheme === 'light' ? lightBg : darkBg}
          onChange={e => {
            if (!e.target.value.startsWith('#')) return

            if (bgTheme === 'light') setLightBg(e.target.value)
            if (bgTheme === 'dark') setDarkBg(e.target.value)
          }}
        />
      </div>
    </div>
  )
}

export default ProfileColorPicker
