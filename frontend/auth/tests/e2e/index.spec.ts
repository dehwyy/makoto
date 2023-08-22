import { expect, test } from '@playwright/test'

test('index page', async ({ page }) => {
	await page.goto('/')

	await expect(page.getByTestId('hoshi-boshi-bg')).toBeVisible()
	await expect(page.getByTestId('form')).toBeVisible()

	// heading
	await expect(page.getByText(/Makoto/i)).toBeVisible()
	await expect(page.getByText('誠')).toBeVisible()

	//? inputs
	// username
	await expect(page.getByPlaceholder('username')).toHaveValue('')

	page.getByPlaceholder('username').type('SomeUsername')
	await expect(page.getByPlaceholder('username')).toHaveValue('SomeUsername')
	// password
	await expect(page.getByPlaceholder('password')).toHaveValue('')

	page.getByPlaceholder('password').type('SomePassword')
	await expect(page.getByPlaceholder('password')).toHaveValue('SomePassword')

	// button
	await expect(page.getByRole('button')).toBeVisible()

	// google
	await expect(page.getByTestId('google')).toBeVisible()
})
