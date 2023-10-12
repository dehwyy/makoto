import { test, expect } from '@playwright/test'

test('SignupPage', async ({ page }) => {
	// 4th line in ./index.spec describes this line
	await page.waitForTimeout(5 * 1000)
	await page.goto('/signup')

	await expect(page.getByTestId('form')).toBeVisible()

	await expect(page.getByPlaceholder('username')).toHaveValue('')
	await expect(page.getByPlaceholder(/^password$/)).toHaveValue('')
	await expect(page.getByPlaceholder('confirm password')).toHaveValue('')
	await expect(page.getByPlaceholder('control question')).toHaveValue('')
	await expect(page.getByPlaceholder('answer on question')).toHaveValue('')
	await expect(page.getByRole('button')).toBeVisible()

	await expect(page.getByText(/sign in/i)).toBeVisible()
})
