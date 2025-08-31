<script lang="ts">
  import { nextUrl, prevUrl } from "../../wizard.svelte";
  import { WindowSetLightTheme } from "$lib/wails/wailsjs/runtime";
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
    return () => {
      SettingsSetFileSharingSettings(settings).catch((err: any) => {
        console.error("Failed to save file sharing settings:", err);
      });
    };
  });
</script>

<div class="space-y-4">
  <header>
    <h1 class="text-2xl font-bold">File Sharing</h1>
    <p class="text-sm text-surface-500">
      Configure how you want to share files between your devices.
    </p>
  </header>

  <div class="space-y-4">
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
        class="slider"
        min="1"
        max="10240"
        bind:value={settings.MaxFileSizeMB}
      />
    </div>
  </div>
</div>
