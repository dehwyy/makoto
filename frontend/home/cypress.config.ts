import { defineConfig } from 'cypress'
import viteConfig from './vite.config.cypress.component'

export default defineConfig({
  video: true,
  videoCompression: 1,
  waitForAnimations: true,
  viewportWidth: 1280,
  viewportHeight: 720,
  videosFolder: 'videos',
  screenshotsFolder: 'screenshots',

  component: {
    devServer: {
      framework: 'vue',
      bundler: 'vite',
      viteConfig,
    },
  },
  e2e: {
    baseUrl: 'http://localhost:3000',
    setupNodeEvents(on, config) {
      // implement node event listeners here
    },
  },
})
