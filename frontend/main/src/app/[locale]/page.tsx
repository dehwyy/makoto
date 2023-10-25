import { Button } from '$/components/ui/button'
import FadeTransition from '$/components/fade'
import { getScopedI18n } from '$/locales/server'
import Link from 'next/link'
import { Routes, Services } from '$/lib/constants'

const Page = async () => {
  const t = await getScopedI18n('index')
  return (
    <main className="bg-[url('/images/gradient-black.svg')] bg-right-top flex flex-1 h-screen overflow-x-hidden overlfow-y-scroll bg-no-repeat bg-cover">
      <section className="w-full grid place-items-center">
        <FadeTransition>
          <div className="font-[600] flex flex-col gap-y-5 text-center select-none text-white">
            <h2 className="font-Kanji text-9xl">誠</h2>
            <h1 className="font-Jua text-8xl">Makoto</h1>
            <div className="grid grid-cols-2 gap-x-5 font-ContentT ">
              <a href={Services.Auth}>
                <Button className="text-lg w-full">{t('sign-in')}</Button>
              </a>
              <Link href={Routes.Docs || '/'}>
                {' '}
                <Button variant="secondary" className="text-lg w-full">
                  {t('docs')}
                </Button>{' '}
              </Link>
            </div>
          </div>
        </FadeTransition>
      </section>
    </main>
  )
}

export default Page
