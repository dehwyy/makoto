<script lang="ts">
	import type { Snapshot } from './$types'

	import { Button, RainbowBlock } from 'makoto-ui-svelte'
	import GoogleIcon from '$lib/components/icons/google.svelte'
	import Underscore from '$lib/components/link.svelte'
	import Input from '$lib/components/input.svelte'
	import { SignInFetch } from '$lib/api/fetches'

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

	const SignIn = async () => {
		const response = await SignInFetch({ password, username })
		console.log(response.status)
	}
</script>

<RainbowBlock>
	<div class="flex flex-col gap-y-3 p-5 sm:p-10">
		<h2 class="text-2xl select-none text-center">
			<span class="font-Kanji font-semibold mr-2 text-4xl text-white">誠</span><span
				class="font-Content text-white">Makoto</span>
		</h2>

		<div class="w-full flex flex-col gap-y-3 font-Content">
			<Input bind:value={username} placeholder="username" />
			<Input bind:value={password} placeholder="password" />
		</div>
		<div class="mt-10 w-full font-ContentT font-[600] text-lg">
			<Button onClick={SignIn}>Log In</Button>
		</div>
		<div class="w-full flex justify-center gap-x-1 mt-2">
			<p class="font-ContentT text-[#818181]">No account yet?</p>
			<Underscore><a href="/signup" class="font-ContentT text-secondary">Sign up</a></Underscore>
		</div>
		<div class="divider font-ContentT">OR</div>
		<a href="/api/v1/auth/google">
			<GoogleIcon />
		</a>
	</div>
</RainbowBlock>
