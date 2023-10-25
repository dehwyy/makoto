import { type ClassValue, clsx } from 'clsx'
import { twMerge } from 'tailwind-merge'

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export const readFileAsBase64 = async (file: File) => {
  return new Promise<string>(resolve => {
    const reader = new FileReader()
    reader.readAsDataURL(file)
    reader.onload = async () => {
      const image = reader.result
      return resolve(String(image))
    }
  })
}
