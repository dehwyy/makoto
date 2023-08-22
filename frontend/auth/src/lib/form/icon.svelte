<script lang="ts">
	import googleIconWhite from '$lib/static/icons/google-white.svg'
	import googleIconGray from '$lib/static/icons/google-gray.svg'

	export let href: string = '#'

	let isHover: boolean = false

	$: icon = isHover ? googleIconWhite : googleIconGray
</script>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<div
	on:mouseenter={() => (isHover = true)}
	on:mouseleave={() => (isHover = false)}
	class="google flex gap-x-5 items-center justify-center w-full cursor-pointer">
	<a class="block relative w-[40px] h-[40px] transition-all duration-500" {href}>
		<span />
		<span />
		<span />
		<span />
		<img src={icon} class="p-2 min-w-[40px] min-h-[40px]" alt="Google" />
	</a>
	<p class="font-Content text-white">Log in with Google</p>
</div>

<style lang="scss">
	.google span {
		position: absolute;
		transition: transform 0.5s;
	}
	.google span:nth-child(1),
	.google span:nth-child(3) {
		width: 100%;
		height: 4px;
		background: #404040;
	}
	.google span:nth-child(1) {
		top: 0;
		left: 0;
		transform-origin: right;
	}
	.google:hover span:nth-child(1) {
		transform: scaleX(0);
		transform-origin: left;
		transition: transform 0.5s;
	}

	.google span:nth-child(3) {
		bottom: 0;
		left: 0;
		transform-origin: left;
	}
	.google:hover span:nth-child(3) {
		transform: scaleX(0);
		transform-origin: right;
		transition: transform 0.5s;
	}

	.google span:nth-child(2),
	.google span:nth-child(4) {
		width: 4px;
		height: 100%;
		background: #404040;
	}
	.google span:nth-child(2) {
		top: 0;
		left: 0;
		transform: scale(0);
		transform-origin: bottom;
	}
	.google:hover span:nth-child(2) {
		transform: scale(1);
		transform-origin: top;
		transition: transform 0.5s;
	}
	.google span:nth-child(4) {
		top: 0;
		right: 0;
		transform: scale(0);
		transform-origin: top;
	}
	.google:hover span:nth-child(4) {
		transform: scale(1);
		transform-origin: bottom;
		transition: transform 0.5s;
	}

	.google:hover span {
		background: white;
	}
</style>
