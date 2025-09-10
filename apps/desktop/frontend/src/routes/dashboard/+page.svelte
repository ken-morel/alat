<script lang="ts">
  import { device, type connected } from "$lib/wails/wailsjs/go/models";
  import type { PageData } from "./$types";
  import { GetConnectedDevices } from "$lib/wails/wailsjs/go/app/App";
  import {
    Tv,
    Smartphone,
    Laptop,
    LucideGhost,
    Cpu,
    Network,
  } from "lucide-svelte";
  import type { DeviceType } from "$lib/device";

  import { onMount } from "svelte";

  let { data }: { data: PageData } = $props();
  let connectedDevices: connected.Connected[] = $state(data.connectedDevices);
  onMount(() => {
    const interval = setInterval(async () => {
      connectedDevices = (await GetConnectedDevices()) ?? [];
    }, 1000);
    return () => {
      clearInterval(interval);
    };
  });

  const iconMap: Record<DeviceType, typeof Tv> = {
    tv: Tv,
    mobile: Smartphone,
    desktop: Laptop,
    arduino: Cpu,
    web: Network,
    unspecified: LucideGhost,
  };
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
          <div
            class="group card flex flex-col justify-between overflow-hidden rounded-lg
         bg-surface-100-900 ring-1 ring-surface-300/50 transition-all
         duration-300 ease-in-out hover:shadow-xl hover:-translate-y-0.5
          hover:bg-surface-200-800"
            style="--device-color: {dev.Info.Color.Hex};"
            aria-label={"Select device " + dev.Info.Name}
            role="button"
            tabindex="0"
            onkeydown={() => null}
          >
            <header class="h-1.5 bg-[var(--device-color)]"></header>

            <article class="flex flex-grow items-start gap-4 p-4">
              <div class="mt-1">
                <!-- svelte-ignore svelte_component_deprecated -->
                <svelte:component
                  this={iconMap[dev.Info.Type as DeviceType] || LucideGhost}
                  color={dev.Info.Color.Hex}
                  class="h-16 w-16 opacity-80"
                />
              </div>

              <div class="flex min-w-0 flex-col">
                <h4 class="h4 font-bold text-surface-700-200">
                  {dev.Info.Name}
                </h4>
                <p
                  class="truncate text-sm text-surface-500-400"
                  title={dev.Info.ID}
                >
                  <small class="text-surface-700-300">{dev.IP}:{dev.Port}</small
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
