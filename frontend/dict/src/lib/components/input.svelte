<script lang="ts">
	// const information
	export let placeholder = ''

	// value
	export let value = ''

	// handle input

	function handleInput(e: Event) {
		value = (e.target as HTMLInputElement).value
	}
</script>

<div class="form__group field w-full relative select-none">
	<input
		{value}
		on:input={handleInput}
		class="form__field w-full font-Content font-[700]"
		autocomplete="off"
		spellcheck="false"
		id={placeholder}
		{placeholder} />
	<label for={placeholder} class="form__label font-ContentT font-semibold select-none"
		>{placeholder}</label>
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
