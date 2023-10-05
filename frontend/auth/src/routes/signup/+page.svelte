<script lang="ts">
	import { Input, Button, RainbowBlock } from 'makoto-ui-svelte'
	import Icon from '$lib/form/google.svelte'
	import LinkWrapper from '$lib/form/link-wrapper.svelte'

	import { fade } from 'svelte/transition'
	import { onMount } from 'svelte'
	import type { Snapshot } from './$types'
	import Placeholder from '$lib/form/placeholder.svelte'

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
		const res = await fetch('/api/sign-up', {
			method: 'POST',
			body: JSON.stringify({
				username,
				password,
				question,
				answer
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
	<main class="h-full w-full grid py-5 place-items-center overflow-x-hidden">
		<div
			data-testid="form"
			class="w-5/6 md:w-2/3 lg:w-1/3 max-w-[400px] mx-auto"
			transition:fade={{ duration: 1500, delay: 100 }}>
			<RainbowBlock>
				<div class="flex flex-col gap-y-3 p-5 sm:p-10">
					<h2 class="text-2xl select-none text-center">
						<span class="font-Kanji font-semibold mr-2 text-4xl text-white">誠</span><span
							class="font-Content text-white">Makoto</span>
					</h2>

					<div class="w-full flex flex-col gap-y-3 font-Content">
						<Input bind:value={username} placeholder="username">
							<Placeholder placeholder="username" />
						</Input>
						<Input ableToChangeVisibility={true} bind:value={password} placeholder="password">
							<Placeholder placeholder="password" />
						</Input>
						<Input
							isPasswordType={true}
							bind:value={repeatedPassword}
							placeholder="confirm password">
							<Placeholder placeholder="confirm password" />
						</Input>
						<Input bind:value={question} placeholder="control question">
							<Placeholder placeholder="control question" />
						</Input>
						<Input bind:value={answer} placeholder="answer on question">
							<Placeholder placeholder="answer on question" />
						</Input>
					</div>
					<div class="mt-10 w-full font-ContentT font-[600] text-lg">
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
				</div>
			</RainbowBlock>
		</div>
	</main>
{/if}
