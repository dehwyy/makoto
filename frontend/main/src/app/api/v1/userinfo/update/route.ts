import { TypedFetch as tp } from '@makoto/lib/typed-fetch'
import { SafeTwirpClient } from '@makoto/grpc/clients'
import { RpcInterceptors } from '@makoto/grpc'
import { UpdateUserInfo } from '$/lib/fetches'
import { cookies, RouteResponse } from '$/lib/route'

export async function POST(req: Request) {
  const payload = await tp.Get(req, UpdateUserInfo)
  console.log(payload, 'PYALOAD')
  const { response, status } = await SafeTwirpClient(cookies).UserInfo.updateUser(
    {
      ...payload,
    },
    {
      interceptors: [RpcInterceptors.AddAuthorizationHeader(cookies.get('token'))],
    },
  )

  const success = response.isSuccess

  return success ? RouteResponse.success({}) : RouteResponse.error(500, status.detail)
}
