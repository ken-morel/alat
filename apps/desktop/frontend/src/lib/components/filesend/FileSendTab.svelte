<script lang="ts">
  import { Accordion, type AccordionItem } from "melt/builders";
  import { slide } from "svelte/transition";
  import { connectedDevices } from "$lib/connected";
  import guessIcon from "$lib/icons";
  import { sendingFiles, sendingDevices, isSendingTo } from "$lib/filesend";
  import UploadIcon from "@lucide/svelte/icons/upload";
  import CloseIcon from "@lucide/svelte/icons/x";
  import {
    AskFilesSend,
    ServiceStartSendFilesToDevice,
  } from "$lib/wails/wailsjs/go/app/App";
  import fsizeText from "$lib/fsize";
  import { get } from "svelte/store";

  const accordion = new Accordion();
  const devices: AccordionItem<{ id: string }> = accordion.getItem({
    id: "devices",
  });
  const files: AccordionItem<{ id: string }> = accordion.getItem({
    id: "files",
  });

  function selectFiles() {
    AskFilesSend().then((selectedFiles) => {
      if (selectedFiles) {
        sendingFiles.update((files) => [...selectedFiles, ...files]);
      }
    });
  }

  function sendFiles() {
    const filesToSend = $sendingFiles.map((f) => f.Path);
    const devicesToSend = get(sendingDevices);
    if (filesToSend.length === 0 || devicesToSend.length === 0) {
      return;
    }

    for (const peer of devicesToSend) {
      ServiceStartSendFilesToDevice(peer, filesToSend);
    }

    // Clear selections after sending
    $sendingFiles = [];
    $sendingDevices = [];
  }
</script>

<div {...accordion.root} class="container p-4">
  <div class="border-b-[1px] border-surface-200-800 w-full">
    <h5 {...files.heading} class="h5 w-full bg-secondary-600-400">
      <button {...files.trigger} class="w-full p-2"> Select files </button>
    </h5>
    {#if files.isExpanded}
      <div {...files.content} transition:slide>
        <article class="place-items-center w-full h-full grid p-10">
          <span class="w-full relative">
            <button
              class="w-full card preset-filled-surface-200-800 hover:bg-surface-300-700 p-8 place-items-center grid text-surface-600-400"
              onclick={selectFiles}
            >
              <UploadIcon size={70} color="var(--color-surface-400-600)" />
              Select files
            </button>
            <button
              class="absolute bottom-0 right-0 text-error-700-300 p-2"
              onclick={() => ($sendingFiles = [])}>Clear</button
            >
          </span>
          <div class="bg-surface-50-950 w-full p-2">
            {#each $sendingFiles as file, idx (file.Path)}
              <div
                transition:slide={{ duration: 100 }}
                class="flex flex-2 text-[12px] text-surface-800-200"
              >
                <div class="flex flex-3 p-1.5 pr-0">
                  <span class="">{file.Path}</span>
                  <div class="grow"></div>
                  <span class="">{fsizeText(file.Size)}</span>
                </div>
                <button
                  class="btn m-0 p-1.5"
                  onclick={() => {
                    sendingFiles.update((files) => {
                      return files.filter((f) => f.Path != file.Path);
                    });
                  }}
                >
                  <CloseIcon size={20} color="var(--color-error-600-400)" />
                </button>
              </div>
              {#if idx < $sendingFiles.length - 1}
                <hr class="border-b-[1px] border-surface-300-700 w-full" />
              {/if}
            {:else}
              <div transition:slide class="p-3 text-surface-300-700">
                No file selected
              </div>
            {/each}
          </div>
        </article>
      </div>
    {/if}
  </div>

  <div class="border-b-[1px] border-surface-200-800 w-full">
    <h5 {...devices.heading} class="h5 w-full bg-secondary-600-400">
      <button {...devices.trigger} class="w-full p-2"> Select devices </button>
    </h5>
    {#if devices.isExpanded}
      <div transition:slide class="p-7">
        <div {...devices.content}>
          {#each $connectedDevices as dev}
            {@const Icon = guessIcon(dev.info.type)}
            <div
              class:bg-surface-200-800={isSendingTo(dev)}
              class="group card flex flex-col justify-between overflow-hidden
         bg-surface-100-900 ring-1 ring-surface-300/50 transition-all
         duration-300 ease-in-out hover:shadow-xl
          hover:border-primary-300-700"
              style="--device-color: {dev.info.color.hex};"
              aria-label={"Select device " + dev.info.name}
              role="button"
              tabindex="0"
              onkeydown={() => null}
              onclick={() => {
                if (isSendingTo(dev)) {
                  sendingDevices.update((devices) =>
                    devices.filter((d) => d.info.id != dev.info.id),
                  );
                } else {
                  sendingDevices.update((devices) => devices.concat([dev]));
                }
              }}
            >
              <article class="flex flex-grow items-start gap-4 p-1">
                <div class="mt-1">
                  <Icon
                    color={dev.info.color.hex}
                    class="h-10 w-10 opacity-80"
                  />
                </div>

                <div class="flex min-w-0 flex-col">
                  <h6 class="h6 font-bold text-surface-700-200">
                    {dev.info.name}
                  </h6>
                </div>
              </article>
            </div>
          {:else}
            <div class="p-4">
              <p class="text-surface-600-400 p-4">No connected device</p>
            </div>
          {/each}
        </div>
      </div>
    {/if}
  </div>
  <div>
    <button
      class="btn preset-filled-surface-300-700 p-3 w-full"
      onclick={sendFiles}
      disabled={$sendingFiles.length === 0 || $sendingDevices.length === 0}
    >
      Send {$sendingFiles.length} files to {$sendingDevices.length} devices
    </button>
  </div>
</div>
