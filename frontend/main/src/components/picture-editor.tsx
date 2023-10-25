'use client'
import React, { useEffect, useRef, useState } from 'react'
import Cropper, { ReactCropperElement } from 'react-cropper'
import 'cropperjs/dist/cropper.css'
import './picture-editor.css'
import { useWindowSize } from 'usehooks-ts'
import { Input } from './ui/input'
import { readFileAsBase64 } from '$/lib/utils'
import { DialogClose } from '@radix-ui/react-dialog'
import { Button } from './ui/button'

interface Props {
  onSave: (imageData64: string) => void
}

const PictureEditor = ({ onSave: onSaveFromParent }: Props) => {
  const [file, setFile] = useState<File | null>()
  const [image, setImage] = useState('')

  useEffect(() => {
    if (!file) return
    readFileAsBase64(file).then(imageBase64 => {
      setImage(imageBase64)
    })
  }, [file])

  const cropperRef = useRef<ReactCropperElement>(null)
  const { height: windowHeight } = useWindowSize()
  const onCrop = () => {
    const cropper = cropperRef.current?.cropper!
  }

  const onSave = () => {
    onSaveFromParent(cropperRef.current!.cropper.getCroppedCanvas().toDataURL())
  }

  return (
    <>
      {image && (
        <Cropper
          src={image}
          className="mt-3"
          style={{ height: Math.min(windowHeight, 500) }}
          initialAspectRatio={1 / 1}
          aspectRatio={1 / 1}
          guides={false}
          viewMode={2}
          crop={onCrop}
          ref={cropperRef}
        />
      )}

      <Input type="file" className="mt-3" onChange={e => setFile(e.target.files?.item(0))} />
      <DialogClose asChild>
        <Button onClick={onSave}>Save</Button>
      </DialogClose>
    </>
  )
}

export default PictureEditor
