'use client'
import { Parallax, ParallaxLayer } from '@react-spring/parallax'
import SkewedText from '@/components/text-effects/SkewedRainboxText'

const page = () => {
  return (
    <Parallax pages={4}>
      <ParallaxLayer factor={1} speed={1.5}>
        <section className="text-9xl font-bold font-ContentT leading-[60px] grid place-items-center text-center h-full">
          <div>
            <h2>
              <span className="text-secondary">Makoto</span>&nbsp;<span className="underline text-primary">team</span>
            </h2>
            <br />
            <p>Introduces</p>
          </div>
        </section>
      </ParallaxLayer>
      <ParallaxLayer
        offset={1}
        speed={0.1}
        factor={1}
        style={{
          backgroundImage: 'url(/ab.jpg)',
          backgroundSize: 'cover',
        }}>
        <div style={{ backgroundPosition: '0 100px' }} className="mt-[-50px] blur-md h-[70px] w-[110%] ml-[-50px] bg-[url('/ab.jpg')]"></div>
        <section className="font-ContentT font-bold text-9xl h-full grid place-items-center pb-20 select-none">
          <SkewedText text="Makoto" subtext="Games" />
        </section>
        <div style={{ backgroundPosition: '0 50px' }} className="mt-[-50px] blur-md h-[70px] w-[110%] ml-[-50px] bg-[url('/ab.jpg')]"></div>
      </ParallaxLayer>

      <ParallaxLayer
        offset={2}
        factor={2}
        speed={0.7}
        style={{
          backdropFilter: 'blur(10px)',
          backgroundImage: 'url(/tetris.jpg)',
          backgroundSize: 'cover',
        }}>
        <section className="">
          <h2>HELLO</h2>
        </section>
      </ParallaxLayer>
    </Parallax>
  )
}

export default page
