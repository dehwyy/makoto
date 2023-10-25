import { PORTS } from '@makoto/config'

export const enum Services {
  Auth = `http://localhost:${PORTS.AUTH}`,
  Hashmap = `http://localhost:${PORTS.HASHMAP}`,
}

// this service
export const enum Routes {
  Index = '/',
  Me = '/me',
  MeEdit = '/me/edit',
  Docs = '/docs',
}
