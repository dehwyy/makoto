<script lang="ts">
	import { fade } from 'svelte/transition'

	// header's icons
	import user from '$lib/assets/user.svg'
	//
	import { onMount } from 'svelte'
	import { Username } from '$lib/stores/user-store'
	import { SERVICES } from '$lib/const'

	let isMounted = false
	onMount(() => {
		isMounted = true
	})
</script>

{#if isMounted}
	<header transition:fade={{ duration: 300, delay: 500 }} class="header">
		{#if Username}
			<div class="user_wrapper">
				<p class="user_item">{$Username}</p>
			</div>
		{:else}<a href={SERVICES.AUTH}>
				<div class="icon_wrapper">
					<img class="icon_item" src={user} alt="logo" />
				</div>
			</a>
		{/if}
	</header>
{/if}

<style lang="scss">
	.header {
		@apply fixed top-0 left-0 right-0;
	}
	.user {
		&_wrapper {
			@apply py-5 px-10 mx-auto w-full flex justify-end;
		}
		&_item {
			@apply font-Jua text-2xl select-none cursor-pointer font-[700];
		}
	}
	.icon {
		&_wrapper {
			@apply cursor-pointer py-5 px-10 h-[40px] w-full mx-auto flex justify-end;
		}
		&_item {
			@apply h-[40px] w-[40px];
		}
	}
</style>
