<script lang="ts">
  import { Popover, Tabs } from "melt/builders";
  import { scale } from "svelte/transition";
  import { ProgressRing } from "@skeletonlabs/skeleton-svelte";
  import IconFile from "@lucide/svelte/icons/file";
  import { onMount } from "svelte";
  import { ServiceGetFileSendStatus } from "$lib/wails/wailsjs/go/app/App";
  import type { filesend } from "$lib/wails/wailsjs/go/models";
  let isOpen: boolean = $state(true);
  let status: filesend.FileTransfersStatus | null = $state(null);
  const popover = new Popover({
    onOpenChange: (value) => (isOpen = value),
  });
  let tabIds = ["Hello", "World"];
  const tabs = new Tabs({
    value: tabIds[0],
  });
  let sendingPercent = $state(100);
  onMount(() => {
    const interval = setInterval(async () => {
      status = (await ServiceGetFileSendStatus()) || [];
    }, 100);
    return () => {
      clearInterval(interval);
    };
  });
</script>

<div {...popover.trigger}>
  <div class="flex items-center space-x-0 mr-0">
    <ProgressRing
      value={sendingPercent}
      size="size-8"
      meterStroke="stroke-success-600-400"
      trackStroke="stroke-warning-50-950"
    >
      <IconFile size={20} />
    </ProgressRing>
  </div>
</div>

{#key isOpen}
  <div
    {...popover.content}
    in:scale={{ opacity: 0.5, duration: 200, start: 0.9 }}
    class="preset-filled-surface-300-700 card"
  >
    {#if status}
      <div class="tabs">
        <div {...tabs.getContent("sending")}>Sending</div>
        <div {...tabs.getContent("receiving")}>Receiving</div>
      </div>
      <div {...tabs.triggerList}>
        <button class="btn p-1" {...tabs.getTrigger("sending")}>Sending</button>
        <button class="btn p-1" {...tabs.getTrigger("receiving")}
          >Receiving</button
        >
      </div>
    {/if}
  </div>
{/key}
