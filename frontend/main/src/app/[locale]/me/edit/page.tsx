'use client'
import { Dialog, DialogTrigger, DialogContent } from '$/components/ui/dialog'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '$/components/ui/tabs'
import { Input } from '$/components/ui/input'
import { Label } from '$/components/ui/label'
import { useState } from 'react'
import { Separator } from '$/components/ui/separator'
import { Textarea } from '$/components/ui/textarea'
import dynamic from 'next/dynamic'
import { Button } from '$/components/ui/button'
import Image from 'next/image'
import { cn } from '$/lib/utils'
import Link from 'next/link'

const PictureEditor = dynamic(() => import('$/components/picture-editor'), { ssr: false })

const Page = () => {
  const initialName = 'dehwyy!'
  const customId = 'a0131d5d-6b2b-4b2a-8b5b-4b2a8b5b4b2a'
  const initialDescription = "Hello, I'm dehwyy and using Makoto but way longer sentence to test if it works."
  const dark_background = 'dark:bg-[#171717]'
  const background = 'bg-[#00a2ff]'

  const [name, setName] = useState(initialName)
  const [id, setId] = useState(customId)
  const [description, setDescription] = useState(initialDescription)

  const [image, setImage] = useState('')

  return (
    <div className="min-h-screen mx-auto pt-28 w-[90%] md:w-2/3">
      <Tabs defaultValue="profile" className="w-full">
        <TabsList className="grid grid-cols-2">
          <TabsTrigger value="profile">User Profile</TabsTrigger>
          <TabsTrigger value="privacy">Privacy</TabsTrigger>
        </TabsList>

        {/* User profile  */}
        <TabsContent value="profile" className="p-5 dark:bg-black rounded-xl overflow-hidden border-2 border-secondary grid lg:grid-cols-2 gap-x-8">
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
            <Separator />
            <WithLabel id="custom_id" text="custom id">
              <Input
                className="w-full mt-2"
                type="text"
                id="custom_id"
                placeholder={id}
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
              <div className="h-[20px] dark:bg-black flex items-center mb-0.5">
                <div className="bg-secondary rounded-full h-[13px] w-[93%] mx-auto px-2">
                  <span className="font-Content text-[8px] relative bottom-[9px] underline select-none">https://makoto/me/{id}</span>
                </div>
              </div>
              <div className={cn(dark_background, background, 'h-[60px]')} />
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
        </TabsContent>

        {/* User privacy  */}
        <TabsContent value="privacy">Change your password here.</TabsContent>
      </Tabs>
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
