import { Button } from '$/components/ui/button'
import { ClearCookies } from '$/app/actions/cookies'

const DataAndStorage = () => {
  const cookies_usage = 2.37
  const LocalStorage_usage = 0

  const data = [
    { name: 'Cookies', value: cookies_usage, action: () => ClearCookies() },
    { name: 'LocalStorage', value: LocalStorage_usage, action: () => {} },
  ]
  return (
    <>
      {data.map(item => (
        <>
          <p>{item.name}</p>
          <p className="text-center select-none">{item.value.toFixed(2)} B</p>
          {item.value === 0 ? (
            <Button variant="secondary" className="select-none" disabled>
              Clear
            </Button>
          ) : (
            <Button variant="secondary" onClick={item.action}>
              Clear
            </Button>
          )}
        </>
      ))}
    </>
  )
}

export default DataAndStorage
