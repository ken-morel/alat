<script lang="ts">
  import { nextUrl, prevUrl } from "../../wizard.svelte";
  import { WindowSetLightTheme } from "$lib/wails/wailsjs/runtime";
  import ServiceTile from "../ServiceTile.svelte";
  import {
    AskFileSharingDestDirectory,
    SettingsSetFileSharingSettings,
  } from "$lib/wails/wailsjs/go/app/App";

  WindowSetLightTheme();

  let { data } = $props();

  nextUrl.set("/setup/services/clipboard");
  prevUrl.set("/setup/services");

  let settings = $state(data.settings);

  async function selectDirectory() {
    const dest: string = await AskFileSharingDestDirectory();
    if (dest && dest.length > 0) {
      settings.DefaultDownloadLocation = dest;
    }
  }

  $effect(() => {
    SettingsSetFileSharingSettings(settings).catch((err: any) => {
      console.error("Failed to save file sharing settings:", err);
    });
  });
</script>

<ServiceTile
  title="File Sharing"
  description="Easily share files between your devices."
  enabled={true}
  disabled={true}
  next="/setup/services/clipboard"
  prev="/setup/services"
>
  <div class="space-y-2">
    <label for="download-location" class="font-medium"
      >Default Download Location</label
    >
    <div class="flex items-center gap-2">
      <input
        id="download-location"
        type="text"
        class="input w-full"
        bind:value={settings.DefaultDownloadLocation}
        readonly
      />
      <button class="btn" onclick={selectDirectory}>Browse</button>
    </div>
  </div>
  <div class="flex items-center gap-2">
    <input
      id="ask-before-receiving"
      type="checkbox"
      class="checkbox"
      bind:checked={settings.AskBeforeReceiving}
    />
    <label for="ask-before-receiving" class="font-medium"
      >Ask before receiving files</label
    >
  </div>

  <div class="space-y-2">
    <label for="max-file-size" class="font-medium"
      >Max File Size ({settings.MaxFileSizeMB} MB)</label
    >
    <input
      id="max-file-size"
      type="range"
      class="slider block"
      min="1"
      max="10240"
      bind:value={settings.MaxFileSizeMB}
    />
  </div>
</ServiceTile>
