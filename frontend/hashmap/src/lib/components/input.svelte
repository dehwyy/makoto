<script lang="ts">
	// Just a wrapper for MakotoUI/Input
	import { Input } from 'makoto-ui-svelte'

	export let value = ''
	export let placeholder = ''
	export let ableToChangeVisibility = false
	export let isPasswordType = false
</script>

<Input bind:value {placeholder} {ableToChangeVisibility} {isPasswordType}>
	<span class="text-md font-ContentT">{placeholder}</span>
</Input>
