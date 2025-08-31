<script lang="ts">
  import { nextUrl, prevUrl } from "../../wizard.svelte";
  import { SettingsSetRemoteInputSettings } from "$lib/wails/wailsjs/go/app/App";

  let { data } = $props();

  prevUrl.set("/setup/services/media");
  nextUrl.set("/setup/services/foldersync");

  let settings = $state(data.settings);

  $effect(() => {
    return () => {
      SettingsSetRemoteInputSettings(settings).catch((err: any) => {
        console.error("Failed to save remote input settings:", err);
      });
    };
  });
</script>

<div class="space-y-4">
  <header>
    <h1 class="text-2xl font-bold">Remote Input</h1>
    <p class="text-sm text-surface-500">
      Use your phone as a remote control for your computer.
    </p>
  </header>

  <div class="space-y-4">
    <div class="flex items-center justify-between">
      <label for="enabled" class="font-medium">Enable Remote Input</label>
      <input id="enabled" type="checkbox" class="checkbox" bind:checked={settings.Enabled} />
    </div>

    {#if settings.Enabled}
      <div class="space-y-2 border-l-2 border-surface-200-800 pl-4">
        <div class="flex items-center justify-between">
          <label for="natural-scrolling" class="font-medium"
            >Natural Scrolling</label
          >
          <input
            id="natural-scrolling"
            type="checkbox"
            class="checkbox"
            bind:checked={settings.NaturalScrolling}
          />
        </div>
        <div class="space-y-2">
          <label for="mouse-sensitivity" class="font-medium"
            >Mouse Sensitivity ({settings.MouseSensitivity})</label
          >
          <input
            id="mouse-sensitivity"
            type="range"
            class="slider"
            min="0.1"
            max="5"
            step="0.1"
            bind:value={settings.MouseSensitivity}
          />
        </div>
      </div>
    {/if}
  </div>
</div>
