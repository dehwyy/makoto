import { Switch } from '$/components/ui/switch'

const PrivacyTab = () => {
  return (
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

export default PrivacyTab
