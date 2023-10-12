import { dataCY } from '../utils/get-data-attr-cy'
import { waitFor } from '../utils/wait-for'

const BASE_URL = Cypress.config().baseUrl
const TRANSITION_MS = 2250

const LOCALE_COOKIE_NAME = 'makoto_locale'

const LOCALE_EN = 'en'
const LOCALE_RU = 'ru'

type LocaleType = typeof LOCALE_EN | typeof LOCALE_RU
type KeysType = 'h1' | 'h2' | 'logo' | 'button'

const LOCALES: Record<LocaleType, Record<KeysType, string | RegExp>> = {
  en: {
    h1: 'Makoto',
    h2: 'All-in-one',
    logo: '誠',
    button: 'Investigate',
  },
  ru: {
    h1: 'Makoto',
    h2: /Вс[ёе]-в-одном/, // doesn't make any sense whether letter is 'e' or 'ё'
    logo: '誠',
    button: 'Исследовать',
  },
}

const TranslationT = (locale: LocaleType) => {
  const currentT = LOCALES[locale]

  // main page content
  cy.get('h1').contains(currentT.h1)
  cy.get(dataCY('subheading')).contains(currentT.h2)
  cy.get(dataCY('button')).contains(currentT.button)
  cy.get(dataCY('logo')).contains(LOCALES[locale].logo)

  // heading & heading's components
  cy.get('header').should('exist')
  cy.get(dataCY('menu')).should('exist')
}

describe('IndexPage:E2E', () => {
  beforeEach(() => {
    // visit IndexPage
    cy.visit('/')
  })
  afterEach(() => {
    cy.clearCookie(LOCALE_COOKIE_NAME)
  })

  // default stays for DEFAULT_LOCALE ( en )
  it('IndexPage:DEFAULT', async () => {
    cy.getCookie(LOCALE_COOKIE_NAME).should('exist').should('have.property', 'value', LOCALE_EN)

    cy.url().should('eq', BASE_URL + '/')

    await waitFor(TRANSITION_MS)
    TranslationT(LOCALE_EN)
  })

  // for locale: RU
  it('IndexPage:RU', async () => {
    cy.setCookie(LOCALE_COOKIE_NAME, LOCALE_RU)
    cy.getCookie(LOCALE_COOKIE_NAME).should('exist').should('have.property', 'value', LOCALE_RU)

    cy.url().should('eq', `${BASE_URL}/${LOCALE_RU}`)

    await waitFor(TRANSITION_MS)
    TranslationT(LOCALE_RU)
  })
})
