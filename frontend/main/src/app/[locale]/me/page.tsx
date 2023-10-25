import { Button } from '$/components/ui/button'
import Image from 'next/image'
import { DiscordLogoIcon, GitHubLogoIcon, InstagramLogoIcon } from '@radix-ui/react-icons'
import { cn } from '$/lib/utils'
import Link from 'next/link'
import { Routes } from '$/lib/constants'

interface PageDefaultProps {
  searchParams: {
    id?: string
  }
  params: {}
}

const Page = async ({ searchParams }: PageDefaultProps) => {
  console.log(searchParams.id)

  const dark_background = 'dark:bg-[#171717]'
  const background = 'bg-[#00a2ff]'
  const nickname = 'dehwyy'
  const description = "Hello, I'm dehwyy and using Makoto but way longer sentence to test if it works."
  const memberSince = '21.10.2023'

  return (
    <div className="w-[85%] mx-auto min-h-screen grid grid-cols-[4fr_5fr] pt-20 gap-x-10 ">
      <section className="p-5">
        <div className="dark:bg-black rounded-xl overflow-hidden border-2 border-secondary">
          <div className={cn(dark_background, background, 'h-[100px]')} />
          <div className="px-7 flex items-center justify-between gap-x-7">
            <div className="dark:border-transparent border-secondary border-2 max-h-[110px] max-w-[110px] min-h-[110px] min-w-[110px] overflow-hidden rounded-full object-cover select-none relative bottom-4">
              <Image priority={true} src="/images/kawaii.png" width={350} height={350} alt="avatar" />
            </div>
            <div className="flex-auto">
              <h3 className="font-Jua text-2xl dark:font-[800] font-[500] tracking-wider ">{nickname}</h3>
            </div>
            <Link href={Routes.MeEdit}>
              <Button className="px-10 font-ContentT">Edit</Button>
            </Link>
          </div>
          <div className="px-7 pb-5">
            {/* integrations */}
            <div className="flex justify-between w-[110px] px-5">
              <DiscordLogoIcon />
              <GitHubLogoIcon />
              <InstagramLogoIcon />
            </div>

            {/* information */}
            <div className="pt-5 pl-3 flex flex-col gap-y-3">
              <p className="font-ContentT font-[500]">{description}</p>
              <p className="text-end font-ContentT">
                Member since <span className="underline">{memberSince}</span>
              </p>
            </div>
          </div>
        </div>
      </section>
      <section className="bg-[#222222] flex-auto">
        <div className="text-center font-Jua py-10 text-2xl">Future content here</div>
      </section>
    </div>
  )
}

export default Page
