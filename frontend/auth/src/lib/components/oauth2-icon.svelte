<script lang="ts">
	export let icon: string
	export let provider: string
</script>

<button
	class="border-solid border border-gray-400 w-full rounded-lg pl-5 pr-10 py-1 flex items-center justify-center gap-x-3 hover:bg-base-200 hover:border-gray-300 transition-all">
	<img src={icon} class="p-2 min-w-[40px] min-h-[40px] max-w-[40px] max-h-[40px]" alt="Google" />
	<span class="font-Jua text-lg relative top-[1px]">Continue with {provider}</span>
</button>
