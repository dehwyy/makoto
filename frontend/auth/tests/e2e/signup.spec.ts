import { test, expect } from '@playwright/test'

test('SignupPage', async ({ page }) => {
	await page.goto('/signup')

	await expect(page.getByTestId('hoshi-boshi-bg')).toBeVisible()
	await expect(page.getByTestId('form')).toBeVisible()

	await expect(page.getByPlaceholder('username')).toHaveValue('')
	await expect(page.getByPlaceholder(/^password$/)).toHaveValue('')
	await expect(page.getByPlaceholder('confirm password')).toHaveValue('')
	await expect(page.getByRole('button')).toBeVisible()
	await expect(page.getByTestId('google')).toBeVisible()

	await expect(page.getByText(/sign in/i)).toBeVisible()
})
