import st from './SkewedRainboxText.module.scss'

interface Props {
  text: string
  subtext: string
}

const SkewedRainboxText = ({ text, subtext }: Props) => {
  return (
    <div className={st['wrapper']}>
      <p className={st['text']}>
        <span data-text={text}>{text}</span>
        <span data-text={subtext}>{subtext}</span>
      </p>
    </div>
  )
}

export default SkewedRainboxText
