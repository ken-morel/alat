<script lang="ts">
  import ServiceTile from "../ServiceTile.svelte";
  import {
    SettingsSetFileSend,
    AskFileSharingDestDirectory,
  } from "$lib/wails/wailsjs/go/app/App";
  import { Select } from "melt/components";

  let { data } = $props();

  let settings = $state(data.settings);

  let unit: string = $state("MB");
  let maxSize: number = $state(settings.MaxSize / (1024 * 1024));
  $effect(() => {
    // @ts-ignore
    settings.MaxSize = SIZES[unit] * maxSize;
  });

  $effect(() => {
    SettingsSetFileSend(settings).catch((err: any) => {
      console.error("Failed to save notification settings:", err);
    });
  });
  const UNITS = ["KB", "MB", "GB", "TB"];
  const SIZES = {
    KB: 1024,
    MB: 1024 * 1024,
    GB: 1024 * 1024 * 1024,
    TB: 1024 * 1024 * 1024 * 1024,
  };
  function changeSaveFolder() {
    AskFileSharingDestDirectory().then((path) => {
      if (path) settings.SaveFolder = path;
    });
  }
</script>

<ServiceTile
  title="File receive"
  description="Allow other devices send you files."
  bind:enabled={settings.Enabled}
  prev="/setup/services/sysinfo"
  next="/setup/done"
>
  <div class="flex items-center justify-between">
    <label for="dest" class="font-medium text-surface-600-400">
      Destination</label
    >
    <input
      type="text"
      class="input text-surface-700-300 ml-2"
      bind:value={settings.SaveFolder}
      disabled
    />
    <button class="btn preset-filled-surface-200-800" onclick={changeSaveFolder}
      >Change</button
    >
  </div>
  <div class="flex flex-2 items-center justify-between">
    <label for="dest" class="grow font-medium text-surface-600-400">
      Maxsize
    </label>
    <input
      type="number"
      pattern="\d+"
      class="input text-surface-700-300 ml-2"
      bind:value={maxSize}
    />
    <Select bind:value={unit}>
      {#snippet children(select)}
        <button class="btn preset-filled-surface-200-800" {...select.trigger}>
          {select.value}
        </button>

        <div {...select.content} class="p-1 bg-surface-200-800">
          {#each UNITS as unit}
            <div
              class="p-1 btn preset-filled-surface-200-800 w-full"
              {...select.getOption(unit)}
            >
              {unit}
            </div>
          {/each}
        </div>
      {/snippet}
    </Select>
  </div>
</ServiceTile>
