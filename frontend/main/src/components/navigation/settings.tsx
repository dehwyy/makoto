import { Dialog, DialogContent, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from '$/components/ui/dialog'
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from '$/components/ui/accordion'
import { Button } from '$/components/ui/button'
import { GearIcon, CookieIcon, CubeIcon } from '@radix-ui/react-icons'
import { Switch } from '$/components/ui/switch'
import Language from '$/components/navigation/settings/language'

const ContentClass = 'grid grid-cols-[2fr_3fr] items-center gap-y-5 py-5 font-Content'

const Settings = () => {
  return (
    <SettingsWrapper>
      <Accordion type="single" collapsible className="w-full">
        {/* 1st - General Settings */}
        <AccordionItem value="item-1">
          <Trigger text="General Settings" Icon={GearIcon} />
          <AccordionContent>
            <div className={`${ContentClass}`}>
              <Language />
            </div>
          </AccordionContent>
        </AccordionItem>

        {/* 2nd - Appearance Settings */}
        <AccordionItem value="item-2">
          <Trigger text="Appearance" Icon={CubeIcon} />
          <AccordionContent>
            <div className="py-3"></div>
          </AccordionContent>
        </AccordionItem>

        {/* 3rd - Storage Settings */}
        <AccordionItem value="item-3">
          <Trigger text="Data and Storage" Icon={CookieIcon} />
          <AccordionContent>Yes. It&apos;s animated by default, but you can disable it if you prefer.</AccordionContent>
        </AccordionItem>
      </Accordion>
    </SettingsWrapper>
  )
}

const Trigger = ({ text, Icon }: { text: string; Icon: typeof CubeIcon }) => {
  return (
    <AccordionTrigger>
      <p className="flex gap-x-2 items-center">
        <Icon className="h-6 w-6" />
        {text}
      </p>
    </AccordionTrigger>
  )
}

const SettingsWrapper = ({ children }: { children: React.ReactNode }) => {
  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="outline" size="icon">
          <GearIcon className="h-4 w-4" />
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Settings</DialogTitle>
        </DialogHeader>
        <div className="flex flex-col gap-y-4 w-full py-5">{children}</div>
        <DialogFooter>
          <Button type="submit">Save changes</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}

export default Settings
