import { Button } from '$/components/ui/button'
import { Dialog, DialogContent, DialogTrigger } from '$/components/ui/dialog'
import { Input } from '$/components/ui/input'
import { Label } from '$/components/ui/label'
import { Separator } from '$/components/ui/separator'
import { Textarea } from '$/components/ui/textarea'
import dynamic from 'next/dynamic'
import { useState } from 'react'
import ProfileColorPicker from './profile_bgcolor-picker'
import Preview from './profile_preview'
import { UpdateUserInfo } from '$/lib/fetches'

interface Props {
  name: string
  email: string
  customId: string
  description: string
  dark_background: string
  light_background: string
  image: string

  isVerifiedEmail: boolean
}

const PictureEditor = dynamic(() => import('$/components/picture-editor'), { ssr: false })

const Profile = (props: Props) => {
  const [name, setName] = useState(props.name)
  const [email, setEmail] = useState(props.email)
  const [id, setId] = useState(props.customId)
  const [description, setDescription] = useState(props.description)

  const [darkBg, setDarkBg] = useState(props.dark_background)
  const [lightBg, setLightBg] = useState(props.light_background)

  const [image, setImage] = useState(props.image)

  const UpdateUser = async () => {
    const res = await UpdateUserInfo({
      darkBg,
      lightBg,
      picture: image,
      userId: id,
      languages: ['russian', 'english'],
      description,
    })
    console.log(res.data)
  }

  return (
    <div className="p-5 dark:bg-black rounded-xl border-2 border-secondary grid lg:grid-cols-2 gap-x-8">
      <section className="font-Content px-3 flex flex-col gap-y-7">
        <WithLabel id="display_name" text="display name">
          <Input
            className="w-full mt-2"
            type="text"
            id="display_name"
            maxLength={20}
            placeholder={props.name}
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
            placeholder={props.email}
            spellCheck={false}
            value={email}
            onChange={e => setEmail(e.target.value)}
          />
          {!props.isVerifiedEmail && (
            <div className="flex flex-col sm:flex-row items-center sm:items-start  gap-x-3 gap-y-2 pl-4 pt-3 -mb-5 sm:-mb-2 ">
              <p className="text-sm text-gray-300">Email not verified ❌</p>
              <Button variant="link" className="-mt-2.5">
                Verify
              </Button>
            </div>
          )}
        </WithLabel>
        <Separator />
        <ProfileColorPicker
          default_dark={props.dark_background}
          default_light={props.light_background}
          darkBg={darkBg}
          lightBg={lightBg}
          setDarkBg={setDarkBg}
          setLightBg={setLightBg}
        />
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
            placeholder={props.description}
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
        <Separator />
        <div className="grid sm:grid-cols-2 grid-cols-1 gap-x-5 gap-y-3">
          <Button variant="default" onClick={UpdateUser}>
            Save
          </Button>
          <Button variant="destructive">Discard</Button>
        </div>
      </section>
      <Preview name={name} id={id} description={description} dark_background={darkBg} light_background={lightBg} image={image} />
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
export default Profile
