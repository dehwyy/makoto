'use client'
import { Dialog, DialogTrigger, DialogContent } from '$/components/ui/dialog'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '$/components/ui/tabs'
import { Input } from '$/components/ui/input'
import { Label } from '$/components/ui/label'
import { useEffect, useState } from 'react'
import { Separator } from '$/components/ui/separator'
import { Textarea } from '$/components/ui/textarea'
import dynamic from 'next/dynamic'
import { Button } from '$/components/ui/button'
import Image from 'next/image'
import { Switch } from '$/components/ui/switch'
import { HexColorPicker } from 'react-colorful'
import { Select, SelectItem, SelectContent, SelectGroup, SelectTrigger } from '$/components/ui/select'
import { useTheme } from 'next-themes'

const PictureEditor = dynamic(() => import('$/components/picture-editor'), { ssr: false })

const Page = () => {
  const initialName = 'dehwyy!'
  const initialEmail = 'dehwyy@google.com'
  const customId = 'a0131d5d-6b2b-4b2a-8b5b-4b2a8b5b4b2a'
  const initialDescription = "Hello, I'm dehwyy and using Makoto but way longer sentence to test if it works."
  const dark_background = '#171717'
  const background = '#2ccce7'

  const { theme: currentTheme, setTheme } = useTheme()

  const [name, setName] = useState(initialName)
  const [email, setEmail] = useState(initialEmail)
  const [id, setId] = useState(customId)
  const [description, setDescription] = useState(initialDescription)

  const [isColorPickerOpen, setIsColorPickerOpen] = useState(false)
  const [bgTheme, setBgTheme] = useState<'light' | 'dark'>('dark')
  const [initialDarkBg, setInitialDarkBg] = useState(dark_background)
  const [darkBg, setDarkBg] = useState(dark_background)
  const [initialLightBg, setInitialLightBg] = useState(background)
  const [lightBg, setLightBg] = useState(background)

  useEffect(() => {
    isColorPickerOpen ? setTheme(bgTheme) : setBgTheme(currentTheme as 'light' | 'dark')
  }, [bgTheme])

  const [image, setImage] = useState('')

  return (
    <div className="min-h-screen mx-auto pt-28 w-[90%] md:w-2/3">
      <Tabs defaultValue="profile" className="w-full">
        <TabsList className="grid grid-cols-2">
          <TabsTrigger value="profile">User Profile</TabsTrigger>
          <TabsTrigger value="privacy">Privacy</TabsTrigger>
        </TabsList>

        {/* User profile  */}
        <TabsContent value="profile">
          <div className="p-5 dark:bg-black rounded-xl overflow-hidden border-2 border-secondary grid lg:grid-cols-2 gap-x-8">
            <section className="font-Content px-3 flex flex-col gap-y-7">
              <WithLabel id="display_name" text="display name">
                <Input
                  className="w-full mt-2"
                  type="text"
                  id="display_name"
                  maxLength={20}
                  placeholder={initialName}
                  spellCheck={false}
                  value={name}
                  onChange={e => setName(e.target.value)}
                />
              </WithLabel>
              <WithLabel id="email" text="email">
                <Input
                  className="w-full mt-2"
                  type="email"
                  id="email"
                  autoComplete="email"
                  placeholder={initialEmail}
                  spellCheck={false}
                  value={email}
                  onChange={e => setEmail(e.target.value)}
                />
              </WithLabel>
              <Separator />
              <div className="grid grid-cols-2 gap-5">
                <p className="font-sans text-sm font-[600] mb-2 col-span-2">BACKGROUND COLOR</p>
                <div className="w-full">
                  <div className="flex flex-col">
                    <Button
                      onClick={() => {
                        setIsColorPickerOpen(p => !p)
                        setInitialDarkBg(darkBg)
                        setInitialLightBg(lightBg)
                      }}
                      variant={isColorPickerOpen ? 'default' : 'outline'}
                      className="w-full">
                      {isColorPickerOpen ? 'Save' : 'Edit'}
                    </Button>
                    {isColorPickerOpen && (
                      <Button
                        onClick={() => {
                          setIsColorPickerOpen(false)
                          setLightBg(initialLightBg)
                          setDarkBg(initialDarkBg)
                        }}
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
              <Separator />
              <WithLabel id="custom_id" text="custom id">
                <Input
                  className="w-full mt-2"
                  type="text"
                  id="custom_id"
                  placeholder={id}
                  autoComplete="disabled"
                  spellCheck={false}
                  value={id}
                  onChange={e => setId(e.target.value)}
                />
              </WithLabel>
              <Separator />
              <WithLabel id="description" text="description">
                <Textarea
                  className="w-full mt-2"
                  id="description"
                  placeholder={initialDescription}
                  spellCheck={false}
                  value={description}
                  onChange={e => setDescription(e.target.value)}
                />
              </WithLabel>
              <Separator />

              <div>
                <p className="font-sans text-sm font-[600] mb-2">AVATAR</p>
                <Dialog>
                  <DialogTrigger className="block" asChild>
                    <Button className="min-w-[50%]">Edit</Button>
                  </DialogTrigger>
                  <DialogContent>
                    <PictureEditor onSave={image => setImage(image)} />
                  </DialogContent>
                </Dialog>
              </div>
            </section>
            <section>
              <p className="font-sans text-sm font-[600] mb-2">PREVIEW</p>
              <div className="dark:bg-black rounded-xl overflow-hidden border-2 border-secondary">
                <div className="h-[20px] dark:bg-black flex items-center my-0.5">
                  <div className="bg-secondary rounded-md h-[13px] w-[93%] mx-auto px-2">
                    <span className="font-Content text-[10px] relative bottom-[9px] underline select-none">https://makoto/me/{id}</span>
                  </div>
                </div>
                <div
                  style={{
                    backgroundColor: currentTheme === 'light' ? lightBg : darkBg,
                  }}
                  className="h-[60px]"
                />
                <div className="px-7 flex justify-between gap-x-7 pt-1">
                  <div className="dark:border-transparent border-secondary border-2 max-h-[110px] max-w-[110px] min-h-[110px] min-w-[110px] overflow-hidden rounded-full object-cover select-none relative bottom-4">
                    <Image priority={true} src={image || '/images/kawaii.png'} width={350} height={350} alt="avatar" />
                  </div>
                  <div className="flex-auto self-center">
                    <h3 className="font-Jua text-2xl dark:font-[800] font-[500] tracking-wider break-all">{name || initialName}</h3>
                  </div>
                </div>
                <div className="px-7 pb-5">
                  {/* integrations */}
                  <div className="flex justify-between w-[110px] px-5"></div>

                  {/* information */}
                  <div className="pt-5 pl-3 flex flex-col gap-y-3">
                    <p className="font-ContentT font-[500]">{description}</p>
                  </div>
                </div>
              </div>
            </section>
          </div>
        </TabsContent>

        {/* User privacy  */}
        <TabsContent value="privacy">
          <div className="p-5 dark:bg-black rounded-xl overflow-hidden border-2 border-secondary grid justify-items-center lg:grid-cols-2 gap-x-8">
            <section>
              <h2 className="font-ContentT text-lg p-3 font-[600]">Profile privacy:</h2>
              <div className="flex flex-col gap-y-6 mt-7 mb-3 font-ContentT">
                <Privacy text="Who can see my profile?">
                  <p>Nobody</p>
                  <Switch />
                  <p>Everyone</p>
                </Privacy>
                <Privacy text="Who can follow me?">
                  <p>Nobody</p>
                  <Switch />
                  <p>Everyone</p>
                </Privacy>
              </div>
            </section>
            <section>
              <h2 className="font-ContentT text-lg p-3 font-[600]">Services privacy:</h2>
              <div className="flex flex-col gap-y-6 mt-7 mb-3 font-ContentT">
                <Privacy text="Who can see my services?">
                  <p>Nobody</p>
                  <Switch />
                  <p>Everyone</p>
                </Privacy>
              </div>
            </section>
          </div>
        </TabsContent>
      </Tabs>
    </div>
  )
}

const Privacy = (props: { children: React.ReactNode; text: string }) => {
  return (
    <div className="flex flex-col gap-y-3">
      <p className="underline">{props.text}</p>
      <div className="flex gap-x-3">{props.children}</div>
    </div>
  )
}

const WithLabel = (props: { children: React.ReactNode; text: string; id: string }) => {
  return (
    <div>
      <Label htmlFor={props.id} className="font-sans text-sm font-[600]">
        {props.text.toUpperCase()}
      </Label>
      {props.children}
    </div>
  )
}

export default Page
