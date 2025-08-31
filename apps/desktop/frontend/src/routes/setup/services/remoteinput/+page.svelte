<script lang="ts">
  import { SettingsSetRemoteInputSettings } from "$lib/wails/wailsjs/go/app/App";
  import ServiceTile from "../ServiceTile.svelte";

  let { data } = $props();
  let settings = $state(data.settings);

  $effect(() => {
    SettingsSetRemoteInputSettings(settings).catch((err: any) => {
      console.error("Failed to save remote input settings:", err);
    });
  });
</script>

<ServiceTile
  title="Remote Input"
  description="Use your phone as a remote control for your computer."
  bind:enabled={settings.Enabled}
  next="/setup/services/foldersync"
  prev="/setup/services/media"
>
  <div class="flex items-center justify-between">
    <label for="natural-scrolling" class="font-medium">Natural Scrolling</label>
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
      class="slider block"
      min="0.1"
      max="5"
      step="0.1"
      bind:value={settings.MouseSensitivity}
    />
  </div>
</ServiceTile>
