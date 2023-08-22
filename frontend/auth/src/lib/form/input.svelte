<script lang="ts">
	import Eye from '$lib/form/eye-watcher.svelte'

	// const
	export let fieldName: string = ''
	export let ableToChangeVisibility: boolean = false
	export let isPasswordConstType: boolean = false
	// not const
	export let value: string = ''

	// local state
	let isHidePassword: boolean = true

	// handlers
	function handleInput(e: Event) {
		value = (e.target as HTMLInputElement).value
	}

	//
	$: inputType = ableToChangeVisibility && isHidePassword ? 'password' : 'text'
</script>

<!-- not using "value:bind" cuz input has "dynimic attribute-type", so Svelte complains -->
<div class="form__group field w-full relative select-none">
	<input
		{value}
		on:input={handleInput}
		class="form__field w-full font-Content font-[700]"
		autocomplete="off"
		spellcheck="false"
		id={fieldName}
		placeholder={fieldName}
		type={isPasswordConstType ? 'password' : inputType} />
	{#if ableToChangeVisibility}
		<div
			role="presentation"
			on:click={() => (isHidePassword = !isHidePassword)}
			class="absolute top-3 right-0 cursor-pointer">
			<Eye />
		</div>
	{/if}
	<label for={fieldName} class="form__label font-ContentT font-semibold select-none"
		>{fieldName}</label>
</div>

<style lang="scss">
	$primary: hsl(314, 85%, 66%);
	$secondary: hsl(224, 85%, 66%);
	$gray: #9b9b9b;
	.form__group {
		padding: 15px 0 0;
		margin-top: 10px;
	}

	.form__field {
		border: 0;
		border-bottom: 2px solid $gray;
		outline: 0;
		padding: 7px 0;
		background: transparent;
		transition: border-color 0.2s;

		&::placeholder {
			color: transparent;
		}

		&:placeholder-shown ~ .form__label {
			cursor: text;
			top: 20px;
		}
	}

	.form__label {
		position: absolute;
		top: 0;
		display: block;
		transition: 0.2s;
		color: $gray;
	}

	.form__field:focus {
		~ .form__label {
			position: absolute;
			top: -5px;
			display: block;
			transition: 0.2s;
		}
		padding-bottom: 6px;
		border-width: 3px;
		border-image: linear-gradient(60deg, hsl(224, 85%, 66%), hsl(314, 85%, 66%));
		border-image-slice: 1;
	}
	/* reset input */
	.form__field {
		&:required,
		&:invalid {
			box-shadow: none;
		}
	}
</style>
