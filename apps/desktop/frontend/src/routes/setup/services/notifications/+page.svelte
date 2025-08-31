<script lang="ts">
  import { SettingsSetNotificationSyncSettings } from "$lib/wails/wailsjs/go/app/App";
  import ServiceTile from "../ServiceTile.svelte";

  let { data } = $props();

  let settings = $state(data.settings);

  $effect(() => {
    SettingsSetNotificationSyncSettings(settings).catch((err: any) => {
      console.error("Failed to save notification settings:", err);
    });
  });
</script>

<ServiceTile
  title="Notification Sync"
  description="Sync notifications between your devices."
  bind:enabled={settings.Enabled}
  next="/setup/services/media"
  prev="/setup/services/clipboard"
>
  <div class="flex items-center justify-between">
    <label for="quick-replies" class="font-medium">Enable Quick Replies</label>
    <input
      id="quick-replies"
      type="checkbox"
      class="checkbox"
      bind:checked={settings.QuickReplies}
    />
  </div>
</ServiceTile>
