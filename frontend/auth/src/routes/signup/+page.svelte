<script lang="ts">
	import { Button, RainbowBlock } from 'makoto-ui-svelte'
	import Link from '$lib/components/link.svelte'
	import GoogleIcon from '$lib/components/icons/google.svelte'
	import type { Snapshot } from './$types'
	import Input from '$lib/components/input.svelte'

	let username: string = ''
	let password: string = ''
	let repeatedPassword: string = ''
	let question: string = ''
	let answer: string = ''

	type SnapshotT = Record<'username', string>
	export const snapshot: Snapshot<SnapshotT> = {
		capture: () => ({
			username
		}),
		restore: value => {
			username = value.username
		}
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
			<Input ableToChangeVisibility={true} bind:value={password} placeholder="password" />
			<Input isPasswordType={true} bind:value={repeatedPassword} placeholder="confirm password" />
			<Input bind:value={question} placeholder="control question" />
			<Input bind:value={answer} placeholder="answer on question" />
		</div>
		<div class="mt-10 w-full font-ContentT font-[600] text-lg">
			<Button onClick={() => {}}>Sign Up</Button>
		</div>
		<div class="w-full flex justify-center gap-x-1 mt-2">
			<p class="text-Content text-[#818181]">Already has account?</p>
			<Link href="/">Sign in</Link>
		</div>
		<!-- hidden, while Google Auth is not providede -->
		<div class="hidden">
			<div class="divider font-ContentT">OR</div>
			<GoogleIcon />
		</div>
	</div>
</RainbowBlock>
