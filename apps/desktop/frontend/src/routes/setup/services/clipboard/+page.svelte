<script lang="ts">
  import { SettingsSetUniversalClipboardSettings } from "$lib/wails/wailsjs/go/app/App";
  import ServiceTile from "../ServiceTile.svelte";

  let { data } = $props();
  let settings = $state(data.settings);

  $effect(() => {
    SettingsSetUniversalClipboardSettings(settings).catch((err: any) => {
      console.error("Failed to save clipboard settings:", err);
    });
  });
</script>

<ServiceTile
  next="/setup/services/notifications"
  prev="/setup/services/fileshare"
  bind:enabled={settings.Enabled}
  title="Universal Clipboard"
  description="Sync your clipboard across all your devices."
>
  <div class="flex items-center justify-between">
    <label for="sync-text" class="font-medium">Sync Text</label>
    <input
      id="sync-text"
      type="checkbox"
      class="checkbox"
      bind:checked={settings.SyncText}
    />
  </div>
  <div class="flex items-center justify-between">
    <label for="sync-images" class="font-medium">Sync Images</label>
    <input
      id="sync-images"
      type="checkbox"
      class="checkbox"
      bind:checked={settings.SyncImages}
    />
  </div>
  <div class="flex items-center justify-between">
    <label for="ignore-password-managers" class="font-medium"
      >Ignore Password Managers</label
    >
    <input
      id="ignore-password-managers"
      type="checkbox"
      class="checkbox"
      bind:checked={settings.IgnorePasswordManagers}
    />
  </div>
</ServiceTile>
