<script lang="ts">
  import { nextUrl, prevUrl } from "../../wizard.svelte";
  import { SettingsSetMediaControlSettings } from "$lib/wails/wailsjs/go/app/App";

  let { data } = $props();

  prevUrl.set("/setup/services/notifications");
  nextUrl.set("/setup/services/remoteinput");

  let settings = $state(data.settings);

  $effect(() => {
    return () => {
      SettingsSetMediaControlSettings(settings).catch((err: any) => {
        console.error("Failed to save media control settings:", err);
      });
    };
  });
</script>

<div class="space-y-4">
  <header>
    <h1 class="text-2xl font-bold">Media Control</h1>
    <p class="text-sm text-surface-500">
      Control media playback on other devices.
    </p>
  </header>

  <div class="space-y-4">
    <div class="flex items-center justify-between">
      <label for="enabled" class="font-medium">Enable Media Control</label>
      <input id="enabled" type="checkbox" class="checkbox" bind:checked={settings.Enabled} />
    </div>
  </div>
</div>
