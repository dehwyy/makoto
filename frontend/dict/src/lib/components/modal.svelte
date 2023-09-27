<script lang="ts">
	export let isOpen: boolean
	export let close = () => {}
</script>

<svelte:body on:keydown={e => e.key === 'Escape' && isOpen && close()} />

{#if isOpen}
	<section
		on:mousedown={close}
		aria-hidden="true"
		class="bg-[rgba(0,0,0,0.7)] z-10 fixed grid place-items-center top-0 left-0 right-0 bottom-0">
		<div
			aria-hidden="true"
			on:mousedown={e => e.stopPropagation()}
			class="w-[400px] bg-base-300 gradient-border border-base-100 rounded-2xl relative">
			<slot />
			<div>
				<span
					on:click={close}
					aria-hidden="true"
					class="text-2xl absolute right-4 top-3 cursor-pointer font-Jua hover:text-red-500 transition-all"
					>X</span>
			</div>
		</div>
	</section>
{/if}

<style lang="scss">
	.gradient-border {
		--border-width: 3px;
		border-radius: var(--border-width);

		&::after {
			position: absolute;
			content: '';
			top: calc(-1 * var(--border-width));
			left: calc(-1 * var(--border-width));
			z-index: -1;
			width: calc(100% + var(--border-width) * 2);
			height: calc(100% + var(--border-width) * 2);
			background: linear-gradient(
				60deg,
				hsl(224, 85%, 66%),
				hsl(269, 85%, 66%),
				hsl(314, 85%, 66%),
				hsl(359, 85%, 66%),
				hsl(44, 85%, 66%),
				hsl(89, 85%, 66%),
				hsl(134, 85%, 66%),
				hsl(179, 85%, 66%)
			);
			background-size: 300% 300%;
			background-position: 0 50%;
			border-radius: calc(2 * var(--border-width));
			animation: moveGradient 4s alternate infinite;
		}
	}

	@keyframes moveGradient {
		50% {
			background-position: 100% 50%;
		}
	}
</style>
