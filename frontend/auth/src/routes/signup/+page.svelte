<script lang="ts">
	import Input from '$lib/form/input.svelte'
	import Button from '$lib/form/button.svelte'
	import FormWindow from '$lib/form/window.svelte'
	import Icon from '$lib/form/icon.svelte'
	import LinkWrapper from '$lib/form/link-wrapper.svelte'

	import { fade } from 'svelte/transition'
	import { onMount } from 'svelte'
	import type { Snapshot } from './$types'

	let isMounted = false

	let username: string = ''
	let password: string = ''
	let repeatedPassword: string = ''
	let question: string = ''
	let answer: string = ''

	onMount(() => {
		isMounted = true
	})

	type SnapshotT = Record<'username', string>
	export const snapshot: Snapshot<SnapshotT> = {
		capture: () => ({
			username
		}),
		restore: value => {
			username = value.username
		}
	}

	const submit = async () => {
		console.log('HELLO!')

		const res = await fetch('/api/sign-up', {
			method: 'POST',
			body: JSON.stringify({
				username,
				password,
				question,
				answer
			})
		})

		const userId = await res.json()
		console.log(userId)
		if (false) {
			window.location.href = 'http://localhost:3000'
		}
	}
</script>

{#if isMounted}
	<main class="h-full w-full grid py-5 place-items-center overflow-x-hidden">
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
					<Input
						bind:value={repeatedPassword}
						isPasswordConstType={true}
						fieldName="confirm password" />
					<Input bind:value={question} fieldName="control question" />
					<Input bind:value={answer} fieldName="answer on question" />
				</div>
				<div class="mt-10 w-full">
					<Button onClick={submit}>Sign Up</Button>
				</div>
				<div class="w-full flex justify-center gap-x-1 mt-2">
					<p class="text-Content text-[#818181]">Already has account?</p>
					<LinkWrapper><a href="/" class="font-ContentT text-secondary">Sign in</a></LinkWrapper>
				</div>
				<!-- hidden, while Google Auth is not providede -->
				<div class="hidden">
					<div class="divider font-ContentT">OR</div>
					<Icon />
				</div>
			</FormWindow>
		</div>
	</main>
{/if}
