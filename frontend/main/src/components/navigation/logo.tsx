import Link from 'next/link'

const Logo = () => {
  return (
    <Link href="/" title="Go to home page">
      <p className="font-Kanji text-3xl font-[600] relative bottom-[2.1px]">
        誠 /<span className="font-Jua text-2xl font-[400] "> Makoto</span>
      </p>
    </Link>
  )
}

export default Logo
