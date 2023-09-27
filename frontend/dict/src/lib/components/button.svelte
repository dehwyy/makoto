<script lang="ts">
	export let onClick = () => {}
</script>

<button on:click={onClick} class="bg-base-300 fancy px-3 py-1 outline-none select-none w-full">
	<span class="top-key" />
	<p class="font-ContentT font-semibold text-white text-center text-xl">
		<slot />
	</p>
	<span class="bottom-key-1" />
	<span class="bottom-key-2" />
</button>

<style lang="scss">
	$primary: white;
	$secondary: black;

	.fancy {
		border: 2px solid $primary;
		float: right;
		position: relative;
		text-align: center;
		text-decoration: none;
		text-transform: none;
		transition: all 0.3s ease-in-out;

		&::before {
			content: ' ';
			width: 1.5625rem;
			height: 2px;
			background: $secondary;
			top: 50%;
			left: 1.5em;
			position: absolute;
			transform: translateY(-50%);
			transform-origin: center;
			transition: background 0.3s linear, width 0.3s linear;
		}
		p {
			display: block;
			transition: all 0.3s ease-in-out;
		}

		.top-key {
			height: 2px;
			width: 1.5625rem;
			top: -2px;
			left: 0.625rem;
			position: absolute;
			background: $secondary;
			transition: width 0.5s ease-out, left 0.3s ease-out;
		}

		.bottom-key-1 {
			height: 2px;
			width: 1.5625rem;
			right: 1.875rem;
			bottom: -2px;
			position: absolute;
			background: $secondary;
			transition: width 0.5s ease-out, right 0.3s ease-out;
		}
		.bottom-key-2 {
			height: 2px;
			width: 0.625rem;
			right: 0.625rem;
			bottom: -2px;
			position: absolute;
			background: $secondary;
			transition: width 0.5s ease-out, right 0.3s ease-out;
		}

		&:hover {
			color: $secondary;
			background: $primary;

			&::before {
				width: 0.9375rem;
				background: $secondary;
			}

			p {
				color: $secondary;
				padding-right: 0.5em;
			}

			.top-key {
				left: -2px;
				width: 0px;
			}
			.bottom-key-1,
			.bottom-key-2 {
				right: 0;
				width: 0;
			}
		}
	}
</style>
