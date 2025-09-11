<script lang="ts">
  import ServiceTile from "../ServiceTile.svelte";
  import { SettingsSetSysInfo } from "$lib/wails/wailsjs/go/app/App";
  import Slider from "$lib/widgets/Slider.svelte";
  import { fade } from "svelte/transition";

  let { data } = $props();

  let settings = $state(data.settings);

  $effect(() => {
    SettingsSetSysInfo(settings).catch((err: any) => {
      console.error("Failed to save notification settings:", err);
    });
  });
</script>

<ServiceTile
  title="System Info"
  description="Allow other devices to get system information like disk usage and battery status."
  bind:enabled={settings.Enabled}
  prev="/setup/services"
  next="/setup/done"
>
  <div class="flex items-center justify-between">
    <label for="cache-time" class="font-medium text-surface-600-400"
      >Cache time</label
    >
    <Slider bind:value={settings.CacheSeconds} max={30} min={1} step={1}>
      {#snippet subtext()}
        <span class="text-xs">
          {settings.CacheSeconds} second{settings.CacheSeconds > 1 ? "s" : ""}
        </span>
      {/snippet}
      {#snippet tooltip()}
        <p
          class="preset-filled-surface-300-700 text-surface-900-100 p-2 max-w-[200px]"
        >
          The minimum interval for alat to fetch system information from your
          os.
        </p>
      {/snippet}
    </Slider>
  </div>
</ServiceTile>
