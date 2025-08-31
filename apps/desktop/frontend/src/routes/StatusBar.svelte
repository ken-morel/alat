<script lang="ts">
  import { GetNodeStatus } from "$lib/wails/wailsjs/go/app/App";
  import type { node } from "$lib/wails/wailsjs/go/models";
  import { AppBar, ProgressRing } from "@skeletonlabs/skeleton-svelte";
  import { onMount } from "svelte";

  // Icons from Skeleton's built-in icon library (or you can use your own)
  import IconBolt from "@lucide/svelte/icons/bolt"; // For server (running/not running)
  import IconMagnifyingGlass from "@lucide/svelte/icons/zoom-in"; // For discovery (running/not running)
  import IconCheckCircle from "@lucide/svelte/icons/check"; // For active/running
  import IconXCircle from "@lucide/svelte/icons/x"; // For inactive/stopped

  let status: node.Status | null = $state(null);
  let serverStatusIcon: typeof IconCheckCircle | null = $state(null);
  let discoveryStatusIcon: typeof IconCheckCircle | null = $state(null);
  let serverStatusText: string = $state("Server Status: Unknown");
  let discoveryStatusText: string = $state("Discovery Status: Unknown");
  let serverStatusColor: string = $state("text-surface-400-500"); // Default gray
  let discoveryStatusColor: string = $state("text-surface-400-500"); // Default gray

  onMount(() => {
    const interval = setInterval(() => {
      GetNodeStatus().then((stat) => {
        status = stat;
        updateStatusDisplay(stat);
      });
    }, 1000);
    return () => clearInterval(interval);
  });

  function updateStatusDisplay(currentStatus: node.Status) {
    // Update Server Status
    if (currentStatus.ServerRunning) {
      serverStatusText = "Server: Active";
      serverStatusIcon = IconCheckCircle;
      serverStatusColor = "text-success-500"; // Green for active
    } else {
      serverStatusText = "Server: Inactive";
      serverStatusIcon = IconXCircle;
      serverStatusColor = "text-error-500"; // Red for inactive
    }

    // Update Discovery Status
    if (currentStatus.DiscoveryRunning) {
      discoveryStatusText = "Discovery: Running";
      discoveryStatusIcon = IconCheckCircle;
      discoveryStatusColor = "text-success-500"; // Blue/Primary for searching
    } else {
      discoveryStatusText = "Discovery: Idle";
      discoveryStatusIcon = IconXCircle;
      discoveryStatusColor = "text-warning-500"; // Yellow/Warning for idle (not necessarily an error)
    }
  }
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
          <IconBolt class="h-5 w-5 {serverStatusColor}" />
          <span class="text-sm {serverStatusColor}">{serverStatusText}</span>
        </div>

        <div class="flex items-center space-x-1">
          <IconMagnifyingGlass class="h-5 w-5 {discoveryStatusColor}" />
          <span class="text-sm {discoveryStatusColor}"
            >{discoveryStatusText}</span
          >
        </div>
      </div>
    {/if}
  {/snippet}

  {#snippet trail()}
    <!-- You can add more elements here if needed, e.g., current time, user info -->
    <!-- For now, we'll keep it simple as per the request focusing on server statuses -->
  {/snippet}
</AppBar>
