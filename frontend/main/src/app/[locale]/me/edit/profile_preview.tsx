import { useTheme } from 'next-themes'
import Image from 'next/image'

interface Props {
  name: string
  id: string
  description: string
  dark_background: string
  light_background: string
  image: string
}

const Preview = ({ name, id, description, dark_background, light_background, image }: Props) => {
  const { theme } = useTheme()
  return (
    <section className="h-min sticky top-16">
      <p className="font-sans text-sm font-[600] mb-2">PREVIEW</p>
      <div className=" dark:bg-black rounded-xl overflow-hidden border-2 border-secondary">
        <div className="h-[20px] dark:bg-black flex items-center my-0.5">
          <div className="bg-secondary rounded-md h-[13px] w-[93%] mx-auto px-2">
            <span className="font-Content text-[10px] relative bottom-[9px] underline select-none">https://makoto/me/{id}</span>
          </div>
        </div>
        <div
          style={{
            backgroundColor: theme === 'light' ? light_background : dark_background,
          }}
          className="h-[60px]"
        />
        <div className="px-7 flex justify-between gap-x-7 pt-1">
          <div className="dark:border-transparent border-secondary border-2 max-h-[110px] max-w-[110px] min-h-[110px] min-w-[110px] overflow-hidden rounded-full object-cover select-none relative bottom-4">
            <Image priority={true} src={image || '/images/kawaii.png'} width={350} height={350} alt="avatar" />
          </div>
          <div className="flex-auto self-center">
            <h3 className="font-Jua text-2xl dark:font-[800] font-[500] tracking-wider break-all">{name}</h3>
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
  )
}

export default Preview
