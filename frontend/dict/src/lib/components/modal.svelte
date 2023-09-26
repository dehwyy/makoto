<script lang="ts">
	export let isOpen: boolean
	export let close = () => {}
</script>

{#if isOpen}
	<section
		on:click={close}
		aria-hidden="true"
		class="bg-[rgba(0,0,0,0.7)] fixed grid place-items-center top-0 left-0 right-0 bottom-0">
		<div
			aria-hidden="true"
			on:click={e => e.stopPropagation()}
			class="w-[400px] bg-base-300 border-4 border-base-100 rounded-2xl relative">
			<slot />
			<div>
				<span
					on:click={close}
					aria-hidden="true"
					class="text-2xl absolute right-2 top-2 cursor-pointer font-Jua hover:text-red-500 transition-all"
					>X</span>
			</div>
		</div>
	</section>
{/if}
