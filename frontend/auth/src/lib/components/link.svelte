<script lang="ts">
	export let href = '/'
</script>

<p class="link-underline relative">
	<a {href} class="font-ContentT text-secondary"><slot /></a>
</p>

<style>
	.link-underline::before {
		content: '';
		position: absolute;
		bottom: 0;
		right: 0;
		width: 0;
		height: 2px;
		background-color: #00a2ff;
		transition: width 0.6s cubic-bezier(0.25, 1, 0.5, 1);
	}

	@media (hover: hover) and (pointer: fine) {
		.link-underline:hover::before {
			left: 0;
			right: auto;
			width: 100%;
		}
	}
</style>
