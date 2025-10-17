<script lang="ts">
  import { QueryDeviceSysInfo } from "$lib/wails/wailsjs/go/app/App";
  import type { connected } from "$lib/wails/wailsjs/go/models";
  import { ProgressRing } from "@skeletonlabs/skeleton-svelte";
  import { onMount } from "svelte";
  import { Spring } from "svelte/motion";
  import IconBattery from "@lucide/svelte/icons/battery";
  import IconBatteryCharging from "@lucide/svelte/icons/battery-charging";
  import IconError from "@lucide/svelte/icons/battery-warning";
  import Tooltip from "$lib/widgets/Tooltip.svelte";

  let { dev, size = 20 }: { dev: connected.Connected; size?: number } =
    $props();

  let iconSize = size * 2;

  let percent = new Spring<number>(0);
  let charging: boolean = $state(false);
  let loaded: boolean = $state(false);
  let error: string | null = $state(null);

  let stroke: string = $derived(
    error ? "error" : loaded ? (charging ? "success" : "tertiary") : "warning",
  );
  let meterStroke = $derived("stroke-" + stroke + "-700-300");
  onMount(() => {
    const interval = setInterval(() => {
      QueryDeviceSysInfo(dev)
        .then((info) => {
          error = null;
          loaded = true;
          percent.set(info.batteryPercent || 0);
          charging = info.batteryCharging || info.batteryPercent == 100;
        })
        .catch((err: any) => {
          error = err.toString();
        });
    }, 5000);
    return () => {
      clearInterval(interval);
    };
  });
</script>

<Tooltip classes="p-4 rounded-xl bg-{error ? 'error' : 'surface'}-300-700">
  <ProgressRing
    value={loaded && !error ? percent.current : null}
    size="size-20"
    {meterStroke}
  >
    {#if error}
      <IconError size={iconSize} />
    {:else if loaded}
      {#if charging}
        <IconBatteryCharging size={iconSize} />
      {:else}
        <IconBattery size={iconSize} />
      {/if}
    {:else}
      ...
    {/if}
  </ProgressRing>
  {#snippet tooltip()}
    <div>
      {#if error}
        <p>{error}</p>
      {:else if loaded}
        Battery at {percent.target.toFixed(1)}%
        <br />
        <div class="flex">
          <IconBatteryCharging class="mr-3.5" />{charging
            ? "Charging"
            : "Not charging"}
        </div>
      {:else}
        <i class="italic">Loading ...</i>
      {/if}
    </div>
  {/snippet}
</Tooltip>
