<script lang="ts">
  import { Popover, Tabs } from "melt/builders";
  import { scale } from "svelte/transition";
  import { ProgressRing } from "@skeletonlabs/skeleton-svelte";
  import IconFile from "@lucide/svelte/icons/file";
  import { onMount } from "svelte";
  import { ServiceGetFileSendStatus } from "$lib/wails/wailsjs/go/app/App";
  import type { filesend } from "$lib/wails/wailsjs/go/models";
  import { Progress } from "@skeletonlabs/skeleton-svelte";
  let isOpen: boolean = $state(true);
  let status: filesend.FileTransfersStatus | null = $state(null);
  const popover = new Popover({
    onOpenChange: (value) => (isOpen = value),
  });
  const tabs = new Tabs<"sending" | "receiving">({
    value: "sending",
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
{#snippet transferTab(devices: filesend.FileTransfersStatusDevice[])}
  {#each devices as device (device.Device.ID)}
    <div class="w-full card preset-outlined-surface-300-700">
      <header>
        <h5 class="h5">{device.Device.Name}</h5>
        <Progress value={device.Percent} max={100} classes="bg-surface-300-700"
        ></Progress>
      </header>
      <article>
        {#each device.Transfers as transfer (transfer.FileName)}
          <div>
            {transfer.FileName.split("/").at(-1)}
            <Progress value={transfer.Percent} max={100} />
          </div>
        {/each}
      </article>
    </div>
  {/each}
{/snippet}

{#key isOpen}
  <div
    {...popover.content}
    in:scale={{ opacity: 0.5, duration: 200, start: 0.9 }}
    class="preset-filled-surface-300-700 card w-[500px] h-[400px]"
  >
    <div class="p-2">
      {#if status}
        <div
          class="tab-buttons-container grid place-items-center w-full"
          {...tabs.triggerList}
        >
          <div class="flex tab-buttons">
            <button class="btn p-1" {...tabs.getTrigger("sending")}
              >Sending</button
            >
            <button class="btn p-1" {...tabs.getTrigger("receiving")}>
              Receiving
            </button>
          </div>
          <hr class="hr" />
        </div>

        <div class="tabs grow">
          <div {...tabs.getContent("sending")}>
            {@render transferTab(status.Sending)}
          </div>
          <div {...tabs.getContent("receiving")}>
            {@render transferTab(status.Receiving)}
          </div>
        </div>
      {/if}
    </div>
  </div>
{/key}

<style lang="sass">
.tab-buttons
  button[data-active]
    background: var(--color-secondary-600)
</style>
