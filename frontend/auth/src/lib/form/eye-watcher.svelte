<script lang="ts">
	import EyeIcon from '$lib/static/icons/eye.svg'
	import { slide, fade } from 'svelte/transition'
	import { onMount } from 'svelte'

	let x: number = 0
	let maxX: number = 0
	let y: number = 0
	let maxY: number = 0

	let isMounted = false
	let isEyeOpen = false

	const handleMouseMove = (e: MouseEvent) => {
		x = e.clientX + 0.1
		y = e.clientY + 0.1
	}

	onMount(() => {
		maxX = window.innerWidth
		maxY = window.innerHeight
		isMounted = true
	})

	$: eyeLeft = 45 + ((x || maxX / 2) / (maxX + 1)) * 10 + '%'
	$: eyeTop = 45 + ((y || maxY / 2) / (maxY + 1)) * 10 + '%'
</script>

<svelte:document on:mousemove={handleMouseMove} />
{#if isMounted}
	<div role="presentation" on:click={() => (isEyeOpen = !isEyeOpen)} class="relative">
		<img src={EyeIcon} alt="eye" class="w-[40px] h-[40px]" />
		{#if isEyeOpen}
			<div
				transition:fade={{ duration: 500 }}
				style:left={eyeLeft}
				style:top={eyeTop}
				class="w-[10px] h-[10px] rounded-full bg-primary border-[3.75px] border-white center-absolute -translate-x-1/2 -translate-y-1/2" />
		{:else}
			<div
				transition:fade={{ duration: 500 }}
				class="w-[10px] h-[10px] border-4 border-error rounded-full bg-white center-absolute -translate-x-1/2 -translate-y-1/2" />
			<div
				transition:slide={{ duration: 500 }}
				class="h-[40px] w-[3px] bg-white center-absolute -translate-x-1/2 -translate-y-1/2 -rotate-45" />
		{/if}
	</div>
{:else}
	<div class="w-[50px] h-[50px]" />
{/if}

<style lang="scss">
	.center-absolute {
		position: absolute;
		top: 50%;
		left: 50%;
	}
</style>
