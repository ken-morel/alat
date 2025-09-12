<script lang="ts">
  import { connectedDevice, sendingFiles } from "$lib/store";
  import UploadIcon from "@lucide/svelte/icons/upload";
  import CloseIcon from "@lucide/svelte/icons/x";
  import { AskFilesSend } from "$lib/wails/wailsjs/go/app/App";
  import fsizeText from "$lib/fsize";
  import { slide } from "svelte/transition";

  const dev = connectedDevice;

  function selectFiles() {
    const device = $dev;
    if (!device) return;
    AskFilesSend(device.Info.Name).then((selectedFiles) => {
      sendingFiles.update((files) => {
        return selectedFiles.concat(files);
      });
    });
  }
</script>

<div class="h-full w-full grid place-items-center">
  {#if $dev}
    <div
      class="card preset-filled-surface-100-900 border-[1px] border-surface-200-800 w-full max-w-lg"
    >
      <header class="flex border-b border-surface-200-800 p-8">
        <h3 class="ml-8 h4">Send files to {$dev.Info.Name}</h3>
      </header>
      <article class="place-items-center w-full h-full grid p-10">
        <button
          class="w-full card preset-filled-surface-200-800 hover:bg-surface-300-700 p-8 place-items-center grid text-surface-600-400"
          onclick={selectFiles}
        >
          <UploadIcon size={70} color="var(--color-surface-400-600)" />
          Select files
        </button>
        <div class="bg-surface-50-950 w-full p-2">
          {#each $sendingFiles as file (file.Path)}
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
            <hr class="border-b-[1px] border-surface-300-700 w-full" />
          {:else}
            <div class="p-3 text-surface-300-700">No file selected</div>
          {/each}
        </div>
      </article>
      <footer class="px-10 pb-9">
        <button class="btn w-full preset-filled-primary-600-400"
          >Send files</button
        >
      </footer>
    </div>
  {:else}
    <div
      transition:slide={{ duration: 100 }}
      class="card preset-filled-error-400-600 p-8"
    >
      <p class="h2 text-error-700-300">No device selected.</p>
    </div>
  {/if}
</div>
