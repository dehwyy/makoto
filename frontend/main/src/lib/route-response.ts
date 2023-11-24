class RouteResponseWrapper {
  success(payload: any) {
    return new Response(
      JSON.stringify({
        ...payload,
      }),
      {
        status: 200,
        statusText: 'OK',
      },
    )
  }

  error(code: number, reason?: string) {
    return new Response(
      JSON.stringify({
        message: reason,
      }),
      {
        status: 500,
        statusText: 'Error',
      },
    )
  }
}

export const RouteResponse = new RouteResponseWrapper()
