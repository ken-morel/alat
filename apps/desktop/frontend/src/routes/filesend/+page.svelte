<script lang="ts">
  import { Accordion, type AccordionItem } from "melt/builders";
  import { slide } from "svelte/transition";
  import { connectedDevices } from "$lib/connected";
  import guessIcon from "$lib/icons";
  import { connected } from "$lib/wails/wailsjs/go/models";

  let choosedDevices = $state<string[]>([]);

  const accordion = new Accordion();
  type Item = AccordionItem<{ id: string }>;
  const devices: Item = accordion.getItem({
    id: "devices",
  });
  const files: Item = accordion.getItem({
    id: "files",
  });

  import { connectedDevice, sendingFiles } from "$lib/store";
  import UploadIcon from "@lucide/svelte/icons/upload";
  import CloseIcon from "@lucide/svelte/icons/x";
  import {
    AskFilesSend,
    ServiceStartSendFilesToDevice,
  } from "$lib/wails/wailsjs/go/app/App";
  import fsizeText from "$lib/fsize";
  import { get } from "svelte/store";
  import type { app } from "$lib/wails/wailsjs/go/models";

  function selectFiles() {
    AskFilesSend().then((selectedFiles) => {
      sendingFiles.update((files) => {
        return selectedFiles.concat(files);
      });
    });
  }
</script>

<div class="w-full h-full grid place-items-center transition-all">
  <div
    class="card preset-filled-surface-100-900 border-[1px] border-surface-200-800 w-full max-w-lg"
  >
    <header class="p-6 pl-8 border-b-2 border-surface-200-800">
      <h3 class="h3">Send files</h3>
    </header>
    <div {...accordion.root}>
      <div class="border-b-[1px] border-surface-200-800 w-full">
        <h5 {...devices.heading} class="h5 w-full bg-surface-300-700">
          <button {...devices.trigger} class="w-full"> Select devices </button>
        </h5>
        {#if devices.isExpanded}
          <div {...devices.content} transition:slide>
            {#each $connectedDevices as dev}
              {@const Icon = guessIcon(dev.info.type)}
              <div
                class:bg-surface-200-800={choosedDevices.includes(dev.info.id)}
                class="group card flex flex-col justify-between overflow-hidden
         bg-surface-100-900 ring-1 ring-surface-300/50 transition-all
         duration-300 ease-in-out hover:shadow-xl hover:-translate-y-0.5
          hover:border-primary-300-700 m-2"
                style="--device-color: {dev.info.color.hex};"
                aria-label={"Select device " + dev.info.name}
                role="button"
                tabindex="0"
                onkeydown={() => null}
                onclick={() => {
                  if (choosedDevices.includes(dev.info.id)) {
                    choosedDevices = choosedDevices.filter(
                      (d) => d != dev.info.id,
                    );
                  } else {
                    choosedDevices.push(dev.info.id);
                    choosedDevices = choosedDevices;
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
        {/if}
      </div>

      <div class="border-b-[1px] border-surface-200-800 w-full">
        <h5 {...files.heading} class="h5 w-full bg-surface-300-700">
          <button {...files.trigger} class="w-full"> Select files </button>
        </h5>
        {#if files.isExpanded}
          <div {...files.content} transition:slide>
            <article class="place-items-center w-full h-full grid p-10">
              <button
                class="w-full card preset-filled-surface-200-800 hover:bg-surface-300-700 p-8 place-items-center grid text-surface-600-400"
                onclick={selectFiles}
              >
                <UploadIcon size={70} color="var(--color-surface-400-600)" />
                Select files
              </button>
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
      <div>
        <button class="btn preset-filled-surface-300-700 p-3 w-full"
          >Send {$sendingFiles.length} files to {choosedDevices.length} devices</button
        >
      </div>
    </div>
  </div>
</div>

<style lang="sass">

</style>
