<script lang="ts" module>
  let isOpen: boolean = $state(false);

  export function openPane() {
    isOpen = true;
  }
  export function closePane() {
    isOpen = false;
  }
  export function togglePane() {
    isOpen = !isOpen;
  }

  const popover = new Popover({
    open: () => isOpen,
    onOpenChange(value) {
      isOpen = value;
    },
  });

  const BTNS: ("receiving" | "sending")[] = ["receiving", "sending"];
  const BTNSL = ["Receiving", "Sending"];
</script>

<script lang="ts">
  import { Popover, Tabs } from "melt/builders";
  import { fly } from "svelte/transition";
  import { ProgressRing } from "@skeletonlabs/skeleton-svelte";
  import IconFile from "@lucide/svelte/icons/file";
  import IconDownload from "@lucide/svelte/icons/download";
  import IconUpload from "@lucide/svelte/icons/upload";
  import { ServiceGetFileSendStatus } from "$lib/wails/wailsjs/go/app/App";
  import type { filesend } from "$lib/wails/wailsjs/go/models";
  import { Progress } from "@skeletonlabs/skeleton-svelte";
  import guessIcon from "$lib/icons";
  import { onMount } from "svelte";

  const tabs = new Tabs<"sending" | "receiving">({
    value: "sending",
  });

  let status: filesend.FileTransfersStatus | null = $state(null);
  let sendingPercent = $state(100);
  let receivingPercent = $state(100);

  let operationsPercent = $state(100);

  $effect(() => {
    operationsPercent = (sendingPercent + receivingPercent) / 2;
  });

  $effect(() => {
    if (!status) {
      sendingPercent = 100;
      receivingPercent = 100;
    } else {
      sendingPercent =
        status.percentSending === 0 ? 100 : status.percentSending;
      receivingPercent =
        status.percentReceiving === 0 ? 100 : status.percentReceiving;
    }
  });
  onMount(() => {
    const interval = setInterval(async () => {
      status = await ServiceGetFileSendStatus();
    }, 200);
    return () => {
      clearInterval(interval);
    };
  });
</script>

<!-- Enhanced Trigger Button -->
<div {...popover.trigger} class="relative">
  <div class="flex items-center justify-center group cursor-pointer">
    <div class="relative">
      <button>
        <ProgressRing
          value={operationsPercent}
          size="size-10"
          meterStroke="stroke-success-400-600"
          trackStroke="stroke-warning-200-800"
          classes="transition-all duration-300 group-hover:scale-110 drop-shadow-lg"
        >
          <IconFile
            size={24}
            class="text-primary-500 group-hover:text-primary-400 transition-colors"
          />
        </ProgressRing>
      </button>

      <!-- Pulse indicator for active transfers -->
      {#if sendingPercent < 100}
        <div
          class="absolute -top-1 -right-1 w-5 h-5 p-2 bg-success-500 rounded-full animate-pulse shadow-lg"
        ></div>
      {/if}
    </div>
  </div>
</div>

<!-- Enhanced Popover Content -->
<div
  class="w-[400px] h-[500px] card preset-filled-surface-200-800 relative"
  {...popover.content}
>
  <div class="flex flex-col grow">
    {#if status}
      <!-- Enhanced Tab Navigation -->
      <div class="border-b border-surface-300-700" {...tabs.triggerList}>
        <div class="flex bg-surface-200-800 relative">
          {#each BTNS as btn, idx}
            <button
              class="flex-1 flex items-center justify-center gap-2 px-4 py-2 text-sm font-medium transition-all duration-200 relative z-10"
              class:bg-primary-400={tabs.value === btn}
              class:text-white={tabs.value === btn}
              class:text-surface-700-300={tabs.value !== btn}
              class:hover:bg-surface-300-700={tabs.value !== btn}
              {...tabs.getTrigger(btn)}
            >
              <IconUpload size={16} />
              {BTNSL[idx]}
            </button>
          {/each}
        </div>
      </div>

      <!-- Tab Content Area -->
      <div class="flex-1 h-[350px] overflow-hidden">
        <div class="h-full overflow-y-auto" {...tabs.getContent("sending")}>
          {#if !status.sending || status.sending.length === 0}
            <div class="w-full h-full place-items-center grid">
              <div
                class="card flex flex-col items-center justify-center text-surface-500"
              >
                <IconUpload size={48} class="mb-4 opacity-50" />
                <p class="text-lg font-medium">No active uploads</p>
                <p class="text-sm">Files you send will appear here</p>
              </div>
            </div>
          {:else}
            {@render transferTab(status.sending || [])}
          {/if}
        </div>

        <div class="h-full overflow-y-auto" {...tabs.getContent("receiving")}>
          {#if !status.receiving || status.receiving.length === 0}
            <div
              class="flex flex-col items-center justify-center h-full text-surface-500"
            >
              <IconDownload size={48} class="mb-4 opacity-50" />
              <p class="text-lg font-medium">No active downloads</p>
              <p class="text-sm">Files you receive will appear here</p>
            </div>
          {:else}
            {@render transferTab(status.receiving || [])}
          {/if}
        </div>
      </div>
    {:else}
      <!-- Loading State -->
      <div class="flex-1 flex items-center justify-center">
        <div class="text-center">
          <div
            class="animate-spin w-8 h-8 border-2 border-primary-500 border-t-transparent rounded-full mx-auto mb-4"
          ></div>
          <p class="text-surface-500">Loading transfers...</p>
        </div>
      </div>
    {/if}
  </div>
</div>

{#snippet transferTab(devices: filesend.FileTransfersStatusDevice[])}
  {#each devices.toSorted( (a, b) => a.device.id.localeCompare(b.device.id), ) as device, i (device.device.id)}
    {@const Icon = guessIcon(device.device.type)}
    <div
      class="bg-surface-100-900 border border-surface-300-700 overflow-hidden hover:shadow-lg transition-all duration-200 hover:border-primary-500/30"
      in:fly={{ y: 20, duration: 300, delay: i * 100 }}
    >
      <!-- Device Header -->
      <div
        class=" from-surface-200-800 to-surface-100-900 px-4 py-3 border-b border-surface-300-700"
      >
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            <div class="p-2 bg-primary-500/10">
              <Icon
                class="text-primary-500"
                size={18}
                color={device.device.color.hex}
              />
            </div>
            <div>
              <h5 class="font-semibold text-surface-900-50">
                {device.device.name}
              </h5>
              <p class="text-xs text-surface-600-400">
                {device.transfers.length} file{device.transfers.length !== 1
                  ? "s"
                  : ""}
              </p>
            </div>
          </div>
          <div class="text-right">
            <p class="text-sm font-medium text-surface-900-50">
              {Math.round(device.percent)}%
            </p>
          </div>
        </div>

        <!-- Overall Device Progress -->
        <div class="mt-3">
          <Progress
            value={device.percent}
            max={100}
            height="5px"
            classes="h-2 bg-surface-300-700  [&>*]:from-primary-500 [&>*]:to-secondary-500"
          />
        </div>
      </div>

      <!-- Files List -->
      <div class="px-1 max-h-80 overflow-y-auto">
        {#each device.transfers.toSorted( (a, b) => a.fileName.localeCompare(b.fileName), ) as transfer, j (transfer.fileName)}
          <div
            class="flex items-center justify-between p-1 bg-surface-50-950 border border-surface-200-800 hover:border-primary-500/20 transition-colors"
            in:fly={{ x: -20, duration: 200, delay: j * 50 }}
          >
            <div class="flex items-center gap-2 flex-1 min-w-0">
              <div class="p-2 bg-surface-200-800 rounded-md flex-shrink-0">
                <IconFile size={14} class="text-surface-600-400" />
              </div>
              <div class="flex-1 min-w-0">
                <p
                  class="font-medium text-sm text-surface-900-50 truncate"
                  title={transfer.fileName}
                >
                  {transfer.fileName.replaceAll("\\", "/").split("/").at(-1)}
                </p>
                <div class="flex items-center justify-between">
                  <Progress
                    value={transfer.percent}
                    max={100}
                    classes="h-1.5 bg-surface-300-700 flex-1 mr-3  [&>*]:to-primary-500 [&>*]:rounded-full"
                  />
                  <span
                    class="text-xs text-surface-600-400 font-medium tabular-nums"
                  >
                    {Math.round(transfer.percent)}%
                  </span>
                </div>
              </div>
            </div>

            <!-- Transfer Status Indicator -->
            <div class="flex-shrink-0 ml-3">
              {#if transfer.percent === 100}
                <div class="w-5 h-5 bg-success-500 rounded-full"></div>
              {:else if transfer.percent > 0}
                <div
                  class="w-5 h-5 bg-primary-500 rounded-full animate-pulse"
                ></div>
              {:else}
                <div class="w-5 h-5 bg-warning-400-600 rounded-full"></div>
              {/if}
            </div>
          </div>
        {/each}
      </div>
    </div>
  {/each}
{/snippet}
