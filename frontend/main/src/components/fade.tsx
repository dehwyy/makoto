'use client'
import { cn } from '$/lib/utils'
import { useIsMounted } from '$/lib/hooks/useIsMounted'

const FadeTransition = ({ children, classes }: { children: React.ReactNode; classes?: string }) => {
  const isMounted = useIsMounted()

  return (
    <div className={`${isMounted ? 'opacity-100 visible' : 'opacity-0 invisible'} ${cn('transition-all duration-300', classes)}`}>{children}</div>
  )
}

export default FadeTransition
