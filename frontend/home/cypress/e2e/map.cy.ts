import { dataCY } from '../utils/get-data-attr-cy'

describe('MapPage:E2E', () => {
  it('MapPage', () => {
    cy.visit('/map')

    // header
    cy.get('header').should('exist')
    cy.get(dataCY('menu')).should('exist')

    // elements
    let lineCounter = 0

    cy.get(dataCY('line'))
      .each(_ => lineCounter++)
      .then(() => {
        cy.get(dataCY('moveable-block')).should('have.length', lineCounter)
        cy.get(dataCY('glow-block')).should('have.length', lineCounter)
        expect(lineCounter).to.be.greaterThan(0)
      })

    cy.get(dataCY('makoto-text')).should('exist')
  })
})
