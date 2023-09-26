'use client'
import { Parallax, ParallaxLayer } from '@react-spring/parallax'
import SkewedText from '@/components/text-effects/SkewedRainboxText'
import Link from 'next/link'

const page = () => {
  return (
    <Parallax pages={3}>
      {/* @1 */}
      <ParallaxLayer factor={0.6} speed={2}>
        <section className="text-7xl font-bold font-ContentT grid place-items-center text-center h-full">
          <div className="flex flex-col gap-y-10">
            <h2>
              <span className="text-secondary">Makoto</span>&nbsp;<span className="underline text-primary">team</span>
            </h2>
            <p>Introduces</p>
          </div>
        </section>
      </ParallaxLayer>

      {/* @2 */}
      <ParallaxLayer offset={0.25} speed={0.5} factor={1}>
        <section className="font-ContentT font-bold text-9xl h-full grid place-items-center pb-20 select-none">
          <SkewedText text="Makoto" subtext="Games" />
        </section>
      </ParallaxLayer>

      {/* @3 */}
      <ParallaxLayer offset={1} factor={1} speed={0.1}>
        <section className="text-center flex flex-col gap-y-24 items-center">
          <h2 className="text-9xl font-Pixel font-bold">tetris</h2>
          <Link href="/tetris">
            <button className="text-primary font-ContentT font-bold underline py-2 px-5 text-8xl">PLAY</button>
          </Link>
        </section>
      </ParallaxLayer>
    </Parallax>
  )
}

export default page
