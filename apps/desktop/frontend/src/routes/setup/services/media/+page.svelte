<script lang="ts">
  import { SettingsSetMediaControlSettings } from "$lib/wails/wailsjs/go/app/App";
  import ServiceTile from "../ServiceTile.svelte";

  let { data } = $props();

  let settings = $state(data.settings);

  $effect(() => {
    SettingsSetMediaControlSettings(settings).catch((err: any) => {
      console.error("Failed to save media control settings:", err);
    });
  });
</script>

<ServiceTile
  title="Media Control"
  description="Control media playback on other devices."
  bind:enabled={settings.Enabled}
  next="/setup/services/remoteinput"
  prev="/setup/services/notifications"
></ServiceTile>
