<script lang="ts">
	import type { Snapshot } from './$types'

	import Input from '$lib/form/input.svelte'
	import Button from '$lib/form/button.svelte'
	import FormWindow from '$lib/form/window.svelte'
	import Icon from '$lib/form/icon.svelte'
	import LinkWrapper from '$lib/form/link-wrapper.svelte'

	import { fade } from 'svelte/transition'
	import { onMount } from 'svelte'

	let isMounted = false
	let username: string = ''
	let password: string = ''

	type SnapshotT = Record<'username', string>
	export const snapshot: Snapshot<SnapshotT> = {
		capture: () => ({
			username
		}),
		restore: value => {
			username = value.username
		}
	}

	onMount(() => {
		isMounted = true
	})

	const submit = async () => {
		const res = await fetch('/api/sign-in', {
			method: 'POST',
			body: JSON.stringify({
				username,
				password
			})
		})

		await res.json()
		const isRedirect = false

		if (isRedirect) {
			window.location.href = 'http://localhost:3000'
		}
	}
</script>

{#if isMounted}
	<main class="h-full w-full grid place-items-center overflow-hidden">
		<div
			class="w-5/6 md:w-2/3 lg:w-1/3 max-w-[400px] mx-auto"
			transition:fade={{ duration: 1500, delay: 100 }}>
			<FormWindow>
				<h2 class="text-2xl select-none">
					<span class="font-Kanji font-semibold mr-2 text-4xl text-white">誠</span><span
						class="font-Content text-white">Makoto</span>
				</h2>

				<div class="w-full flex flex-col gap-y-3">
					<Input bind:value={username} fieldName="username" />
					<Input bind:value={password} ableToChangeVisibility={true} fieldName="password" />
				</div>
				<div class="mt-10 w-full">
					<Button onClick={submit}>Log In</Button>
				</div>
				<div class="w-full flex justify-center gap-x-1 mt-2">
					<p class="font-ContentT text-[#818181]">No account yet?</p>
					<LinkWrapper
						><a href="/signup" class="font-ContentT text-secondary">Sign up</a></LinkWrapper>
				</div>
				<div class="divider font-ContentT">OR</div>
				<Icon />
			</FormWindow>
		</div>
	</main>
{/if}
