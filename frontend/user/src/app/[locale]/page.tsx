import { getI18n } from "$/locales/server"

const Page = async () => {
  const t = await getI18n()
  return (
    <div>
      <div className="font-Kanji">{t("t")}</div>
    </div>
  )
}

export default Page
