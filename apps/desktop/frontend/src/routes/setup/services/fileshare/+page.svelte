<script lang="ts">
  import { nextUrl, prevUrl } from "../../wizard.svelte";
  import {
    AskFileSharingDestDirectory,
    SettingsSetFileSharingSettings,
  } from "$lib/wails/wailsjs/go/app/App";
  import { createButton, createCheckbox, createSlider } from "@melt-ui/svelte";

  const {
    elements: { root: browseButton },
  } = createButton();
  const {
    elements: { root: checkbox, input },
    states: { checked },
  } = createCheckbox({
    defaultChecked: data.settings.AskBeforeReceiving,
  });
  const {
    elements: { root: slider, range, thumb },
    states: { value },
  } = createSlider({
    defaultValue: [data.settings.MaxFileSizeMB],
    max: 10240,
    min: 1,
  });

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
    settings.AskBeforeReceiving = $checked;
    settings.MaxFileSizeMB = $value[0];
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
        <button class="btn" use:browseButton on:click={selectDirectory}
          >Browse</button
        >
      </div>
    </div>

    <div class="flex items-center gap-2">
      <button class="checkbox" use:checkbox>
        <input use:input />
      </button>
      <label for="ask-before-receiving" class="font-medium"
        >Ask before receiving files</label
      >
    </div>

    <div class="space-y-2">
      <label for="max-file-size" class="font-medium"
        >Max File Size ({settings.MaxFileSizeMB} MB)</label
      >
      <span use:slider class="slider">
        <span use:range />
        <span use:thumb />
      </span>
    </div>
  </div>
</div>
