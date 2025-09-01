<script lang="ts">
  import type { discovery } from "$lib/wails/wailsjs/go/models";
  import {
    Tv,
    Smartphone,
    Laptop,
    LucideGhost,
    Cpu,
    Network,
  } from "lucide-svelte";
  import type { DeviceType } from "$lib/device";
  import type { MouseEventHandler } from "svelte/elements";

  let {
    device: foundDevice,
    onclick,
  }: {
    device: discovery.FoundDevice;
    onclick?: MouseEventHandler<HTMLDivElement>;
  } = $props();
  const device = foundDevice.Info;

  const iconMap: Record<DeviceType, typeof Tv> = {
    tv: Tv,
    mobile: Smartphone,
    desktop: Laptop,
    arduino: Cpu,
    web: Network,
    unspecified: LucideGhost,
  };
  const DeviceIcon = iconMap[device.Type as DeviceType] || LucideGhost;
</script>

<div
  class="group card flex flex-col justify-between overflow-hidden rounded-lg
         bg-surface-100-900 ring-1 ring-surface-300/50 transition-all
         duration-300 ease-in-out hover:shadow-xl hover:-translate-y-0.5
          hover:bg-surface-200-800"
  style="--device-color: {device.Color.Hex};"
  {onclick}
  aria-label={"Select device " + device.Name}
  role="button"
  tabindex="0"
  onkeydown={() => null}
>
  <header class="h-1.5 bg-[var(--device-color)]"></header>

  <article class="flex flex-grow items-start gap-4 p-4">
    <div class="mt-1">
      <DeviceIcon color={device.Color.Hex} class="h-16 w-16 opacity-80" />
    </div>

    <div class="flex min-w-0 flex-col">
      <h4 class="h4 font-bold text-surface-700-200">
        {device.Name}
      </h4>
      <p class="truncate text-sm text-surface-500-400" title={device.ID}>
        {device.ID.slice(0, 15)}...
        <small class="text-surface-700-300"
          >{foundDevice.IP}:{foundDevice.Port}</small
        >
      </p>
    </div>
  </article>
</div>
