<script lang="ts">
  import { SettingsSetFolderSyncSettings } from "$lib/wails/wailsjs/go/app/App";
  import ServiceTile from "../ServiceTile.svelte";

  let { data } = $props();

  let settings = $state(data.settings);

  $effect(() => {
    SettingsSetFolderSyncSettings(settings).catch((err: any) => {
      console.error("Failed to save folder sync settings:", err);
    });
  });
</script>

<ServiceTile
  title="Folder Sync"
  description="Sync folders between your devices."
  bind:enabled={settings.Enabled}
  next="/setup/done"
  prev="/setup/services/remoteinput"
></ServiceTile>
