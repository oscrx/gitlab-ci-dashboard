import { Text } from '@mantine/core'
import { useEffect, useState } from 'react'

export default function Version() {
  const [version, setVersion] = useState('')

  useEffect(() => {
    window
      .fetch('/api/version')
      .then((r) => r.text())
      .then((v) => setVersion(v))
  }, [])

  return (
    <Text className="text-white hidden sm:block" size="xs">
      {version}
    </Text>
  )
}
