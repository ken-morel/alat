<script lang="ts" module>
  let status: filesend.FileTransfersStatus | null = $state(null);

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
  setInterval(async () => {
    status = (await ServiceGetFileSendStatus()) || [];
  }, 100);
</script>

<script lang="ts">
  import { Popover, Tabs } from "melt/builders";
  import { scale, fly } from "svelte/transition";
  import { ProgressRing } from "@skeletonlabs/skeleton-svelte";
  import IconFile from "@lucide/svelte/icons/file";
  import IconDownload from "@lucide/svelte/icons/download";
  import IconUpload from "@lucide/svelte/icons/upload";
  import IconDevice from "@lucide/svelte/icons/smartphone";
  import { ServiceGetFileSendStatus } from "$lib/wails/wailsjs/go/app/App";
  import type { filesend } from "$lib/wails/wailsjs/go/models";
  import { Progress } from "@skeletonlabs/skeleton-svelte";
  import guessIcon from "$lib/icons";

  const tabs = new Tabs<"sending" | "receiving">({
    value: "sending",
  });

  let sendingPercent = $state(100);
</script>

<!-- Enhanced Trigger Button -->
<div {...popover.trigger} class="relative">
  <div class="flex items-center justify-center group cursor-pointer">
    <div class="relative">
      <button>
        <ProgressRing
          value={sendingPercent}
          size="size-10"
          meterStroke="stroke-primary-500"
          trackStroke="stroke-surface-300-700"
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
  <div
    class="bg-gradient-to-r from-primary-500/10 to-secondary-500/10 border-b border-surface-300-700 px-6 py-4"
  >
    <h3
      class="text-lg font-semibold text-surface-900-50 flex items-center gap-3"
    >
      <IconFile class="text-primary-500" size={20} />
      File Transfers
    </h3>
  </div>

  <div class="flex flex-col grow">
    {#if status}
      <!-- Enhanced Tab Navigation -->
      <div
        class="px-6 py-4 border-b border-surface-300-700"
        {...tabs.triggerList}
      >
        <div class="flex bg-surface-200-800 p-1 relative">
          <button
            class="flex-1 flex items-center justify-center gap-2 px-4 py-2 text-sm font-medium transition-all duration-200 relative z-10"
            class:bg-primary-500={tabs.value === "sending"}
            class:text-white={tabs.value === "sending"}
            class:text-surface-700-300={tabs.value !== "sending"}
            class:hover:bg-surface-300-700={tabs.value !== "sending"}
            {...tabs.getTrigger("sending")}
          >
            <IconUpload size={16} />
            Sending
          </button>
          <button
            class="flex-1 flex items-center justify-center gap-2 px-4 py-2 text-sm font-medium transition-all duration-200 relative z-10"
            class:bg-primary-500={tabs.value === "receiving"}
            class:text-white={tabs.value === "receiving"}
            class:text-surface-700-300={tabs.value !== "receiving"}
            class:hover:bg-surface-300-700={tabs.value !== "receiving"}
            {...tabs.getTrigger("receiving")}
          >
            <IconDownload size={16} />
            Receiving
          </button>
        </div>
      </div>

      <!-- Tab Content Area -->
      <div class="flex-1 overflow-hidden">
        <div
          class="h-full overflow-y-auto px-6 py-4 space-y-4"
          {...tabs.getContent("sending")}
        >
          {#if !status.Sending || status.Sending.length === 0}
            <div
              class="flex flex-col items-center justify-center h-full text-surface-500"
            >
              <IconUpload size={48} class="mb-4 opacity-50" />
              <p class="text-lg font-medium">No active uploads</p>
              <p class="text-sm">Files you send will appear here</p>
            </div>
          {:else}
            {@render transferTab(status.Sending || [])}
          {/if}
        </div>

        <div
          class="h-full overflow-y-auto px-6 py-4 space-y-4"
          {...tabs.getContent("receiving")}
        >
          {#if !status.Receiving || status.Receiving.length === 0}
            <div
              class="flex flex-col items-center justify-center h-full text-surface-500"
            >
              <IconDownload size={48} class="mb-4 opacity-50" />
              <p class="text-lg font-medium">No active downloads</p>
              <p class="text-sm">Files you receive will appear here</p>
            </div>
          {:else}
            {@render transferTab(status.Receiving || [])}
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
  {#each devices as device, i (device.Device.ID)}
    {@const Icon = guessIcon(device.Device.Type)}
    <div
      class="bg-surface-100-900 border border-surface-300-700 overflow-hidden hover:shadow-lg transition-all duration-200 hover:border-primary-500/30"
      in:fly={{ y: 20, duration: 300, delay: i * 100 }}
    >
      <!-- Device Header -->
      <div
        class="bg-gradient-to-r from-surface-200-800 to-surface-100-900 px-4 py-3 border-b border-surface-300-700"
      >
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            <div class="p-2 bg-primary-500/10">
              <Icon class="text-primary-500" size={18} color={device.Device.Color.Hex} />
            </div>
            <div>
              <h5 class="font-semibold text-surface-900-50">
                {device.Device.Name}
              </h5>
              <p class="text-xs text-surface-600-400">
                {device.Transfers.length} file{device.Transfers.length !== 1
                  ? "s"
                  : ""}
              </p>
            </div>
          </div>
          <div class="text-right">
            <p class="text-sm font-medium text-surface-900-50">
              {Math.round(device.Percent)}%
            </p>
          </div>
        </div>

        <!-- Overall Device Progress -->
        <div class="mt-3">
          <div class="flex justify-between text-xs text-surface-600-400 mb-1">
            <span>Overall Progress</span>
            <span>{Math.round(device.Percent)}%</span>
          </div>
          <Progress
            value={device.Percent}
            max={100}
            classes="h-2 bg-surface-300-700 [&>*]:bg-gradient-to-r [&>*]:from-primary-500 [&>*]:to-secondary-500 [&>*]:rounded-full"
          />
        </div>
      </div>

      <!-- Files List -->
      <div class="p-4 space-y-3 max-h-48 overflow-y-auto">
        {#each device.Transfers as transfer, j (transfer.FileName)}
          <div
            class="flex items-center justify-between p-3 bg-surface-50-950 rounded-lg border border-surface-200-800 hover:border-primary-500/20 transition-colors"
            in:fly={{ x: -20, duration: 200, delay: j * 50 }}
          >
            <div class="flex items-center gap-3 flex-1 min-w-0">
              <div class="p-2 bg-surface-200-800 rounded-md flex-shrink-0">
                <IconFile size={14} class="text-surface-600-400" />
              </div>
              <div class="flex-1 min-w-0">
                <p
                  class="font-medium text-sm text-surface-900-50 truncate"
                  title={transfer.FileName}
                >
                  {transfer.FileName.split("/").at(-1)}
                </p>
                <div class="flex items-center justify-between mt-2">
                  <Progress
                    value={transfer.Percent}
                    max={100}
                    classes="h-1.5 bg-surface-300-700 flex-1 mr-3 [&>*]:bg-gradient-to-r [&>*]:from-success-500 [&>*]:to-primary-500 [&>*]:rounded-full"
                  />
                  <span
                    class="text-xs text-surface-600-400 font-medium tabular-nums"
                  >
                    {Math.round(transfer.Percent)}%
                  </span>
                </div>
              </div>
            </div>

            <!-- Transfer Status Indicator -->
            <div class="flex-shrink-0 ml-3">
              {#if transfer.Percent === 100}
                <div class="w-2 h-2 bg-success-500 rounded-full"></div>
              {:else if transfer.Percent > 0}
                <div
                  class="w-5 h-5 bg-primary-500 rounded-full animate-pulse"
                ></div>
              {:else}
                <div class="w-2 h-2 bg-surface-400-600 rounded-full"></div>
              {/if}
            </div>
          </div>
        {/each}
      </div>
    </div>
  {/each}
{/snippet}

<!-- <style lang="sass"> -->
<!-- /* Custom scrollbar for webkit browsers */ -->
<!-- :global(.overflow-y-auto::-webkit-scrollbar) -->
<!--   width: 6px -->
<!---->
<!-- :global(.overflow-y-auto::-webkit-scrollbar-track) -->
<!--   background: transparent -->
<!---->
<!-- :global(.overflow-y-auto::-webkit-scrollbar-thumb) -->
<!--   background: rgb(var(--color-surface-400) / 0.5) -->
<!--   border-radius: 3px -->
<!---->
<!-- :global(.overflow-y-auto::-webkit-scrollbar-thumb:hover) -->
<!--   background: rgb(var(--color-surface-500) / 0.7) -->
<!-- </style> -->
