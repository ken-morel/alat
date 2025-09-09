<script lang="ts">
  import type { connected } from "$lib/wails/wailsjs/go/models";
  import type { PageData } from "./$types";
  import { GetConnectedDevices } from "$lib/wails/wailsjs/go/app/App";

  import { onMount } from "svelte";

  let { data }: { data: PageData } = $props();
  let connectedDevices: connected.Connected[] = $state(data.connectedDevices);
  onMount(() => {
    const interval = setInterval(async () => {
      connectedDevices ??= await GetConnectedDevices();
    }, 100);
    return () => {
      clearInterval(interval);
    };
  });
</script>

<div class="h-full w-full grid place-items-center">
  <div
    class="card preset-filled-surface-100-900 border-[1px] border-surface-200-800 w-full max-w-lg"
  >
    <header class="border-b border-surface-200-800 p-8"></header>
    <article class="p-8">
      <h3 class="h3">Active devices</h3>
      <div class="flex p-2">
        {#each connectedDevices as dev, i (i)}
          <div class="card preset-outlined-primary-200-800 p-2">
            <div>{dev.Info.Name}</div>
          </div>
        {:else}
          <div class="p-2">
            <p class="text-surface-600-400">No connected active device found</p>
          </div>
        {/each}
      </div>
    </article>
  </div>
</div>
tton
