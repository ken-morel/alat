<script lang="ts">
  import { nextUrl, prevUrl } from "../../wizard.svelte";
  import { SettingsSetUniversalClipboardSettings } from "$lib/wails/wailsjs/go/app/App";

  let { data } = $props();

  prevUrl.set("/setup/services/fileshare");
  nextUrl.set("/setup/services/notifications");

  let settings = $state(data.settings);

  $effect(() => {
    return () => {
      SettingsSetUniversalClipboardSettings(settings).catch((err: any) => {
        console.error("Failed to save clipboard settings:", err);
      });
    };
  });
</script>

<div class="space-y-4">
  <header>
    <h1 class="text-2xl font-bold">Universal Clipboard</h1>
    <p class="text-sm text-surface-500">
      Sync your clipboard across all your devices.
    </p>
  </header>

  <div class="space-y-4">
    <div class="flex items-center justify-between">
      <label for="enabled" class="font-medium">Enable Universal Clipboard</label
      >
      <input id="enabled" type="checkbox" class="checkbox" bind:checked={settings.Enabled} />
    </div>

    {#if settings.Enabled}
      <div class="space-y-2 border-l-2 border-surface-200-800 pl-4">
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
      </div>
    {/if}
  </div>
</div>
