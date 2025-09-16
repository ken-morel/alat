<script lang="ts">
  import type { connected } from "$lib/wails/wailsjs/go/models";
  import type { PageData } from "./$types";
  import { GetConnectedDevices } from "$lib/wails/wailsjs/go/app/App";
  import { connectedDevice } from "$lib/store";

  import guessIcon from "$lib/icons";

  import { onMount } from "svelte";
  import { goto } from "$app/navigation";

  let { data }: { data: PageData } = $props();
  let connectedDevices: connected.Connected[] = $state(data.connectedDevices);

  function viewConnectedDevice(dev: connected.Connected) {
    connectedDevice.set(dev);
    goto("/dashboard/device");
  }
  onMount(() => {
    const interval = setInterval(async () => {
      connectedDevices = (await GetConnectedDevices()) || [];
    }, 500);
    return () => {
      clearInterval(interval);
    };
  });
</script>

<div class="h-full w-full grid place-items-center">
  <div
    class="card preset-filled-surface-100-900 border-[1px] border-surface-200-800 w-full max-w-lg"
  >
    <header class="border-b border-surface-200-800 p-8">
      <h3 class="h3">Active devices</h3>
    </header>
    <article class="p-8">
      <div class="flex p-2">
        {#each connectedDevices as dev, i (i)}
          {@const Icon = guessIcon(dev.info.type)}
          <div
            class="group card flex flex-col justify-between overflow-hidden rounded-lg
         bg-surface-100-900 ring-1 ring-surface-300/50 transition-all
         duration-300 ease-in-out hover:shadow-xl hover:-translate-y-0.5
          hover:bg-surface-200-800"
            style="--device-color: {dev.info.color.hex};"
            aria-label={"Select device " + dev.info.name}
            role="button"
            tabindex="0"
            onkeydown={() => null}
            onclick={() => viewConnectedDevice(dev)}
          >
            <header class="h-1.5 bg-[var(--device-color)]"></header>

            <article class="flex flex-grow items-start gap-4 p-4">
              <div class="mt-1">
                <!-- svelte-ignore svelte_component_deprecated -->
                <Icon color={dev.info.color.hex} class="h-16 w-16 opacity-80" />
              </div>

              <div class="flex min-w-0 flex-col">
                <h4 class="h4 font-bold text-surface-700-200">
                  {dev.info.name}
                </h4>
                <p
                  class="truncate text-sm text-surface-500-400"
                  title={dev.info.id}
                >
                  <small class="text-surface-700-300">{dev.ip}:{dev.port}</small
                  >
                </p>
              </div>
            </article>
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
