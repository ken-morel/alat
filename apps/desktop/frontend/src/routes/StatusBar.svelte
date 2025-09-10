<script lang="ts">
  import {
    GetNodeStatus,
    GetPairedDevices,
  } from "$lib/wails/wailsjs/go/app/App";
  import type { node } from "$lib/wails/wailsjs/go/models";
  import { AppBar, ProgressRing } from "@skeletonlabs/skeleton-svelte";
  import { onMount } from "svelte";

  // Icons from Skeleton's built-in icon library (or you can use your own)
  import IconBolt from "@lucide/svelte/icons/bolt"; // For server (running/not running)
  import IconMagnifyingGlass from "@lucide/svelte/icons/zoom-in"; // For discovery (running/not running)
  import IconNode from "@lucide/svelte/icons/network";

  let status: node.Status | null = $state(null);
  let numberOfPairedDevices: number = $state(0);

  onMount(() => {
    const interval = setInterval(() => {
      GetNodeStatus().then((stat) => {
        status = stat;
      });
      GetPairedDevices().then((dev) => {
        if (dev) {
          numberOfPairedDevices = dev.length;
        } else {
          numberOfPairedDevices = 0;
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
        <div class="chip flex items-center space-x-1">
          <IconBolt
            class="h-5 w-5 {status.ServerRunning
              ? 'text-success-500'
              : 'text-warning-500'}"
          />
          <span
            class="text-sm {status.ServerRunning
              ? 'text-success-500'
              : 'text-warning-500'}"
          >
            Server: {status.ServerRunning ? "Running" : "Stalked"}
          </span>
        </div>

        <div class="chip flex items-center space-x-1">
          <IconMagnifyingGlass
            class="h-5 w-5 {status.DiscoveryRunning
              ? 'text-success-500'
              : 'text-warning-500'}"
          />
          <span
            class="text-sm {status.DiscoveryRunning
              ? 'text-success-500'
              : 'text-warning-500'}"
          >
            {status.DiscoveryRunning ? "Device discoverable" : "Device hidden"}
          </span>
        </div>

        <div class="flex items-center space-x-1">
          <span class="h-5 w-5">{numberOfPairedDevices}</span>
          <span class="text-sm"
            >Paired device{numberOfPairedDevices > 1 ? "s" : ""}</span
          >
        </div>
        <div class="chip flex items-center space-x-1">
          <IconNode
            class="h-5 w-5 {status.WorkerRunning
              ? 'text-success-500'
              : 'text-warning-500'}"
          />
          <span
            class="text-sm {status.WorkerRunning
              ? 'text-success-500'
              : 'text-warning-500'}"
          >
            {status.WorkerRunning ? "Node running" : "Node stopped"}
          </span>
        </div>
      </div>
    {/if}
  {/snippet}

  {#snippet trail()}
    <!-- You can add more elements here if needed, e.g., current time, user info -->
    <!-- For now, we'll keep it simple as per the request focusing on server statuses -->
  {/snippet}
</AppBar>
