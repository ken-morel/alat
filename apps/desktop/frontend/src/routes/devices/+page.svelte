<script lang="ts">
  import {
    GetFoundDevices,
    RequestPairingFoundDevice,
    GetConnectedDevices,
  } from "$lib/wails/wailsjs/go/app/App";
  import { onMount } from "svelte";
  import { ProgressRing } from "@skeletonlabs/skeleton-svelte";
  import { discovery, type device } from "$lib/wails/wailsjs/go/models";
  import FoundDeviceTile from "$lib/components/tiles/FoundDeviceTile.svelte";
  import { pairDialogOptions } from "../PairDialog.svelte";
  let { data } = $props();
  let devices: discovery.FoundDevice[] = $state(data.found);

  onMount(() => {
    const searchInterval = setInterval(async () => {
      const foundDevices = (await GetFoundDevices()) || [];
      const connectedDevices = await GetConnectedDevices();
      devices = [];
      for (const foundDev of foundDevices)
        for (const connDev of connectedDevices)
          if (connDev.Info.ID !== foundDev.Info.ID) devices.push(foundDev);
    }, 500);
    return () => {
      clearInterval(searchInterval);
    };
  });
</script>

<div class="w-full h-full grid place-items-center transition-all">
  <div
    class="card preset-filled-surface-100-900 border-[1px] border-surface-200-800 w-full max-w-lg p-8"
  >
    <header><h3 class="h3">Found devices</h3></header>
    <div class="pt-4">
      {#each devices as device}
        <FoundDeviceTile
          {device}
          onclick={() => {
            pairDialogOptions.set({
              info: device.Info,
              accept: () => {
                RequestPairingFoundDevice(device.Info.ID);
              },
              decline: () => {},
            });
          }}
        />
      {:else}
        <p class="text-surface-300-700">No device found</p>
      {/each}
    </div>
    <div class=" flex mt-8">
      <ProgressRing
        value={null}
        size="size-10"
        meterStroke="stroke-tertiary-600-400"
      />
      <span class="m-3 text-surface-600-400">Searching devices</span>
    </div>
  </div>
</div>
