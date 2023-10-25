import Image from 'next/image'
import { forwardRef } from 'react'

interface IProps extends React.ImgHTMLAttributes<unknown> {
  image: string
  alt?: string
  w: number
  h: number
  cls?: string
}

/**
 *
 * @description Wrapper on next/image Image Component which has PRELOAD and QUALITY=100 by default
 */

const SSRImage = forwardRef(({ image, h, w, cls, alt, placeholder: _, ...attrs }: IProps, ref: React.Ref<HTMLImageElement>) => {
  return (
    <Image
      {...attrs}
      ref={ref}
      priority={true}
      quality={100}
      src={image}
      alt={alt || ''}
      width={w}
      height={h}
      className={`${cls} ${attrs.className}`}
    />
  )
})

SSRImage.displayName = 'SSRImage'

export default SSRImage
