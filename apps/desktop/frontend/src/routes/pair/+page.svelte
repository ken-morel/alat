<script lang="ts">
  import {
    GetFoundDevices,
    IsSearchingDevices as IsSearching,
    SearchDevices,
  } from "$lib/wails/wailsjs/go/app/App";
  import { onMount } from "svelte";
  import { ProgressRing } from "@skeletonlabs/skeleton-svelte";
  import type { device } from "$lib/wails/wailsjs/go/models";
  import FoundDeviceTile from "$lib/components/tiles/FoundDeviceTile.svelte";
  import { pairDialogOptions } from "../PairDialog.svelte";
  let devices: device.Info[] = $state([]);
  let isSearching: boolean = $state(false);
  async function startSearch() {
    await SearchDevices();
  }
  onMount(() => {
    const searchInterval = setInterval(() => {
      GetFoundDevices().then((found) => {
        if (found) {
          devices = found;
        }
      });
      IsSearching().then((searching) => {
        if (searching !== isSearching) {
          isSearching = searching;
        }
      });
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
    <h4 class="h4 text-surface-400-600">Searching devices</h4>
    <div>
      {#each devices as device}
        <FoundDeviceTile
          {device}
          onclick={() => {
            pairDialogOptions.set({
              info: device,
              accept: () => {
                alert("Accepted" + device.Name);
              },
              decline: () => {
                alert("Declined" + device.Name);
              },
            });
          }}
        />
      {:else}
        <p class="text-surface-300-700">No device found</p>
      {/each}
    </div>
    <div class="place-items-center grid mt-8">
      {#if isSearching}
        <ProgressRing
          value={null}
          size="size-10"
          meterStroke="stroke-tertiary-600-400"
        />
      {:else}
        <button
          class="btn preset-filled-secondary-600-400 w-full h-10"
          onclick={startSearch}>Search</button
        >
      {/if}
    </div>
  </div>
</div>
