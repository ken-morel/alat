<script lang="ts">
  import {
    GetConnectedDevices,
    GetNodeStatus,
    GetPairedDevices,
  } from "$lib/wails/wailsjs/go/app/App";
  import type { node } from "$lib/wails/wailsjs/go/models";
  import { AppBar, ProgressRing } from "@skeletonlabs/skeleton-svelte";
  import { onMount } from "svelte";
  import Tooltip from "$lib/widgets/Tooltip.svelte";

  // Icons from Skeleton's built-in icon library (or you can use your own)
  import IconBolt from "@lucide/svelte/icons/bolt"; // For server (running/not running)
  import IconMagnifyingGlass from "@lucide/svelte/icons/zoom-in"; // For discovery (running/not running)
  import IconNode from "@lucide/svelte/icons/network";
  import IconConnectedDevices from "@lucide/svelte/icons/laptop-minimal";

  let status: node.Status | null = $state(null);
  let numberOfConnectedDevices: number = $state(0);

  onMount(() => {
    const interval = setInterval(() => {
      GetNodeStatus().then((stat) => {
        status = stat;
      });
      GetPairedDevices().then((dev) => {
        if (dev) {
          numberOfConnectedDevices = dev.length;
        } else {
          numberOfConnectedDevices = 0;
        }
      });
    }, 500);
    return () => clearInterval(interval);
  });
</script>

<AppBar base="h-16 w-full border-t border-surface-200-800 px-4">
  {#snippet lead()}
    {#if !status}
      <div class="flex items-center space-x-2">
        <ProgressRing
          value={null}
          size="size-8"
          meterStroke="stroke-tertiary-600-400"
          trackStroke="stroke-tertiary-50-950"
        />
        <span class="text-surface-500-400">Loading Node Status...</span>
      </div>
    {:else}
      <div class="flex items-center space-x-4">
        <Tooltip>
          {#snippet tooltip()}
            {#if status}
              <div class="p-4">
                <div class=" flex space-x-1">
                  <IconBolt
                    class="h-5 w-5 {status.ServerRunning
                      ? 'text-success-500'
                      : 'text-warning-500'}"
                  />
                  <span class="text-sm">
                    Services are {status.ServerRunning ? "running" : "stalked"}
                  </span>
                </div>
                <div class=" flex space-x-1">
                  <IconMagnifyingGlass
                    class="h-5 w-5 {status.DiscoveryRunning
                      ? 'text-success-500'
                      : 'text-warning-500'}"
                  />
                  <span class="text-sm">
                    {status.DiscoveryRunning
                      ? "Device visible"
                      : "Device hidden"}
                  </span>
                </div>
                <div class="flex space-x-1">
                  <IconNode
                    class="h-5 w-5 {status.WorkerRunning
                      ? 'text-success-500'
                      : 'text-warning-500'}"
                  />
                  <span class="text-sm">
                    {status.WorkerRunning ? "Alat is running" : "Alat stalked"}
                  </span>
                </div>
              </div>
            {/if}
          {/snippet}

          <div class="chip flex items-center space-x-1">
            <IconBolt
              class="h-5 w-5 {status.ServerRunning
                ? 'text-success-500'
                : 'text-warning-500'}"
            />
            <IconMagnifyingGlass
              class="h-5 w-5 {status.DiscoveryRunning
                ? 'text-success-500'
                : 'text-warning-500'}"
            />
            <IconNode
              class="h-5 w-5 {status.WorkerRunning
                ? 'text-success-500'
                : 'text-warning-500'}"
            />
          </div>
        </Tooltip>
      </div>
    {/if}
  {/snippet}

  {#snippet trail()}
    <!-- You can add more elements here if needed, e.g., current time, user info -->
    <!-- For now, we'll keep it simple as per the request focusing on server statuses -->
    <div class="flex items-center space-x-1">
      <span class="snippen text-surface-800-200">
        <span class="icon">
          <IconConnectedDevices size="23" />
        </span>
        <span class="count">{numberOfConnectedDevices}</span>
        <span class="label text-sm">
          Active device{numberOfConnectedDevices > 1 ? "s" : ""} amongs
        </span>
      </span>
    </div>
  {/snippet}
</AppBar>

<style lang="sass">
span.snippen
  display: flex
  background-color: var(--color-tertiary-500)
  border-radius: 20px
  height: 33px
  span.icon
    background-color: var(--color-secondary-500)
    padding: 5px
    border-radius: 20px
  span.count
    padding: 5px
    padding-right: 10px
  span.label
    padding: 2px
    width: 0px
    transition: width 0.4s
    overflow-x: hidden
  &:hover span.label
    width: 120px
    display: inline
</style>
