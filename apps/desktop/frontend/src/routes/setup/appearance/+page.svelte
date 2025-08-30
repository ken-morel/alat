<script lang="ts">
	import { wizardState } from '../wizard-state';

	// Set the navigation state for this step
	wizardState.prevUrl = '/setup/welcome';
	wizardState.nextUrl = '/setup/services';

	// A simple array of colors for the user to choose from
	const colorSwatches = [
		'#3498db', // Blue
		'#2ecc71', // Green
		'#e74c3c', // Red
		'#f1c40f', // Yellow
		'#9b59b6', // Purple
		'#e67e22', // Orange
	];

	let selectedColor = $state(colorSwatches[0]);
</script>

<div class="space-y-2">
	<h3 class="h3">Choose a Custom Color</h3>
	<p class="text-surface-500">This color will help you identify this device.</p>
</div>

<!-- A grid of color swatches -->
<div class="grid grid-cols-6 gap-4 py-4">
	{#each colorSwatches as color}
		<button
			type="button"
			class="w-full aspect-square rounded-full transition-all"
			class:ring-2={selectedColor === color}
			class:ring-primary-500={selected-Color === color}
			class:ring-offset-2={selectedColor === color}
			style="background-color: {color};"
			on:click={() => (selectedColor = color)}
			aria-label="Select color {color}"
		/>
	{/each}
</div>
