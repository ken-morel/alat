<script lang="ts">
  import { nextUrl, prevUrl } from "../../wizard.svelte";
  import { SettingsSetNotificationSyncSettings } from "$lib/wails/wailsjs/go/app/App";

  let { data } = $props();

  prevUrl.set("/setup/services/clipboard");
  nextUrl.set("/setup/services/media");

  let settings = $state(data.settings);

  $effect(() => {
    return () => {
      SettingsSetNotificationSyncSettings(settings).catch((err: any) => {
        console.error("Failed to save notification settings:", err);
      });
    };
  });
</script>

<div class="space-y-4">
  <header>
    <h1 class="text-2xl font-bold">Notification Sync</h1>
    <p class="text-sm text-surface-500">
      Sync notifications between your devices.
    </p>
  </header>

  <div class="space-y-4">
    <div class="flex items-center justify-between">
      <label for="enabled" class="font-medium">Enable Notification Sync</label>
      <input id="enabled" type="checkbox" class="checkbox" bind:checked={settings.Enabled} />
    </div>

    {#if settings.Enabled}
      <div class="space-y-2 border-l-2 border-surface-200-800 pl-4">
        <div class="flex items-center justify-between">
          <label for="quick-replies" class="font-medium">Enable Quick Replies</label>
          <input
            id="quick-replies"
            type="checkbox"
            class="checkbox"
            bind:checked={settings.QuickReplies}
          />
        </div>
      </div>
    {/if}
  </div>
</div>
