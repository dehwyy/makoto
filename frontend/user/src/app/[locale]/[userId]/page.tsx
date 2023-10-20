import { Button } from '$/components/ui/button'
import Image from 'next/image'
import { DiscordLogoIcon, GitHubLogoIcon, InstagramLogoIcon } from '@radix-ui/react-icons'
import { cn } from '$/lib/utils'

const Page = () => {
  const dark_background = 'dark:bg-[#171717]'
  const background = 'bg-[#00a2ff]'
  const nickname = 'dehwyy'

  return (
    <main className="w-[85%] mx-auto min-h-screen grid grid-cols-[4fr_5fr] pt-20 gap-x-10">
      <section className="p-5">
        <div className="dark:bg-black rounded-xl overflow-hidden">
          <div className={cn(dark_background, background, 'h-[100px]')} />
          <div className="px-7 flex items-center justify-between gap-x-7">
            <div className="dark:border-transparent border-secondary border-2 max-h-[110px] max-w-[110px] min-h-[110px] min-w-[110px] overflow-hidden rounded-full object-cover select-none relative bottom-4">
              <Image priority={true} src="/images/kawaii.png" width={350} height={350} alt="avatar" />
            </div>
            <div className="flex-auto">
              <h3 className="font-Jua text-2xl font-[800] tracking-wider ">{nickname}</h3>
            </div>
            <div>
              <Button>Edit profile</Button>
            </div>
          </div>
          <div className="px-7 pb-5">
            {/* integrations */}
            <div className="flex justify-between w-[110px] px-5">
              <DiscordLogoIcon />
              <GitHubLogoIcon />
              <InstagramLogoIcon />
            </div>
          </div>
        </div>
      </section>
      <section className="bg-[#222222] flex-auto">
        <div className="text-center font-Jua py-10 text-2xl">Future content here</div>
      </section>
    </main>
  )
}

export default Page
