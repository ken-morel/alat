<script lang="ts">
  import { nextUrl, prevUrl } from "../../wizard.svelte";
  import { SettingsSetFolderSyncSettings } from "$lib/wails/wailsjs/go/app/App";
  import { createSwitch } from "@melt-ui/svelte";

  let { data } = $props();

  prevUrl.set("/setup/services/remoteinput");
  nextUrl.set("/setup/done");

  let settings = $state(data.settings);

  const {
    elements: { root: enabledSwitch, thumb: enabledThumb },
    states: { checked: enabledChecked },
  } = createSwitch({
    defaultChecked: settings.Enabled,
  });

  $effect(() => {
    settings.Enabled = $enabledChecked;
    return () => {
      SettingsSetFolderSyncSettings(settings).catch((err: any) => {
        console.error("Failed to save folder sync settings:", err);
      });
    };
  });
</script>

<div class="space-y-4">
  <header>
    <h1 class="text-2xl font-bold">Folder Sync</h1>
    <p class="text-sm text-surface-500">
      Sync folders between your devices.
    </p>
  </header>

  <div class="space-y-4">
    <div class="flex items-center justify-between">
      <label for="enabled" class="font-medium">Enable Folder Sync</label>
      <button use:enabledSwitch class="switch">
        <span use:enabledThumb />
      </button>
    </div>
  </div>
</div>
