export default defineNuxtRouteMiddleware((to, from) => {
  const cookie = useCookie('makoto_locale')

  if (cookie.value === undefined) {
    cookie.value = 'en'
  }

  if (cookie.value !== 'en' && !to.path.startsWith('/ru')) {
    return navigateTo({
      path: `/ru${to.path.length > 1 ? to.path : ''}`,
    })
  }

  if (cookie.value !== 'ru' && to.path.startsWith('/ru')) {
    return navigateTo({
      path: `/${to.path.slice(3)}`,
    })
  }
})
