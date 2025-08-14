<script lang="ts">
	import { onMount } from 'svelte';
	import { GetConfig, SaveConfig } from '../../../wailsjs/go/config/Config';
	import { options } from 'wailsjs/go/models';

	let config: options.RGBA | null = null;

	onMount(async () => {
		const result = await GetConfig();
		config = result;
	});

	async function save() {
		if (config) {
			await SaveConfig(config);
		}
	}
</script>

<div class="container">
	<h1>Setup</h1>

	{#if config}
		<form on:submit|preventDefault={save}>
			<div class="form-group">
				<label for="deviceName">Device Name</label>
				<input id="deviceName" type="text" bind:value={config.deviceName} />
			</div>

			<div class="form-group">
				<label for="deviceColor">Device Color</label>
				<input id="deviceColor" type="color" bind:value={config.deviceColor} />
			</div>

			<div class="form-group">
				<label for="language">Language</label>
				<select id="language" bind:value={config.language}>
					<option value="en-cm">English</option>
					<option value="fr-cm">French</option>
				</select>
			</div>

			<div class="form-group">
				<label for="autoStart">
					<input id="autoStart" type="checkbox" bind:checked={config.autoStart} />
					Auto Start
				</label>
			</div>

			<div class="form-group">
				<label for="theme">Theme</label>
				<select id="theme" bind:value={config.theme}>
					<option value="light">Light</option>
					<option value="dark">Dark</option>
				</select>
			</div>

			<button type="submit">Save</button>
		</form>
	{:else}
		<p>Loading...</p>
	{/if}
</div>
