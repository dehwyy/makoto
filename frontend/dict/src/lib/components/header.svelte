<script lang="ts">
	import { fade } from 'svelte/transition'

	// header's icons
	import user from '$lib/assets/user.svg'
	import { onMount } from 'svelte'

	let isMounted = false

	const AUTH = 'http://localhost:3001'

	onMount(() => {
		isMounted = true
	})
</script>

{#if isMounted}
	<header transition:fade={{ duration: 300, delay: 500 }} class="fixed top-0 left-0 right-0">
		<a href={AUTH}>
			<div class="cursor-pointer py-5 px-10 h-[40px] w-full mx-auto flex justify-end">
				<img class="h-[40px] w-[40px]" src={user} alt="logo" />
			</div>
		</a>
	</header>
{/if}
