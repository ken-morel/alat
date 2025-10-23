<script lang="ts">
  import { Accordion, type AccordionItem } from "melt/builders";
  import { slide } from "svelte/transition";
  import { onMount } from "svelte";
  import {
    AskFilesSend,
    WebShareStart,
    WebShareStop,
    WebShareGetStatus,
    WebShareAddSharedFiles,
    WebShareRemoveSharedFile,
    WebShareClearSharedFiles,
    WebShareSetPasscode,
  } from "$lib/wails/wailsjs/go/app/App";
  import type { webshare } from "$lib/wails/wailsjs/go/models";
  import fsizeText from "$lib/fsize";
  import UploadIcon from "@lucide/svelte/icons/upload";
  import CloseIcon from "@lucide/svelte/icons/x";

  let webShareStatus: webshare.Status | null = $state(null);
  let newPasscode: string = $state("");

  const accordion = new Accordion();
  const serverItem: AccordionItem<{ id: string }> = accordion.getItem({
    id: "server",
  });
  const filesItem: AccordionItem<{ id: string }> = accordion.getItem({
    id: "files",
  });

  onMount(async () => {
    webShareStatus = await WebShareGetStatus();
  });

  async function startWebShare() {
    await WebShareStart();
    await refreshWebShareStatus();
  }

  async function stopWebShare() {
    await WebShareStop();
    await refreshWebShareStatus();
  }

  async function setWebSharePasscode() {
    if (newPasscode) {
      await WebShareSetPasscode(newPasscode);
      await refreshWebShareStatus();
      newPasscode = "";
    }
  }

  async function addFilesToWebShare() {
    const selectedFiles = await AskFilesSend();
    if (selectedFiles && selectedFiles.length > 0) {
      const paths = selectedFiles.map((file) => file.Path);
      await WebShareAddSharedFiles(paths);
      await refreshWebShareStatus();
    }
  }

  async function removeSharedFile(uuid: string) {
    await WebShareRemoveSharedFile(uuid);
    await refreshWebShareStatus();
  }

  async function clearAllSharedFiles() {
    await WebShareClearSharedFiles();
    await refreshWebShareStatus();
  }

  async function refreshWebShareStatus() {
    webShareStatus = await WebShareGetStatus();
  }
</script>

<div {...accordion.root} class="container p-4">
  <!-- Server Status and Controls -->
  <div class="border-b-[1px] border-surface-200-800 w-full">
    <h5 {...serverItem.heading} class="h5 w-full bg-secondary-600-400">
      <button {...serverItem.trigger} class="w-full p-2">
        Server Controls
      </button>
    </h5>
    {#if serverItem.isExpanded}
      <div {...serverItem.content} transition:slide class="p-4 space-y-4">
        {#if webShareStatus}
          <div class="flex justify-between items-center">
            <span
              >Status: <span class="font-bold"
                >{webShareStatus.isRunning ? "Running" : "Stopped"}</span
              ></span
            >
            {#if webShareStatus.isRunning}
              <button
                class="btn preset-filled-error-500-400 p-2"
                onclick={stopWebShare}
              >
                Stop Server
              </button>
            {:else}
              <button
                class="btn preset-filled-success-500-400 p-2"
                onclick={startWebShare}
              >
                Start Server
              </button>
            {/if}
          </div>
          {#if webShareStatus.isRunning}
            <p>
              Share URL:
              <a
                href={webShareStatus.shareURL}
                target="_blank"
                rel="noopener noreferrer"
                class="text-blue-500 hover:underline"
                >{webShareStatus.shareURL}</a
              >
            </p>
            <p>
              Passcode: <span class="font-bold">{webShareStatus.passcode}</span>
            </p>
          {/if}
          <div class="flex gap-2">
            <input
              type="text"
              class="input flex-grow"
              placeholder="Set a new passcode..."
              bind:value={newPasscode}
            />
            <button
              class="btn preset-filled-primary-500-400 p-2"
              onclick={setWebSharePasscode}
              disabled={!newPasscode}>Set</button
            >
          </div>
        {:else}
          <p>Loading status...</p>
        {/if}
      </div>
    {/if}
  </div>

  <!-- List Shared Files -->
  <div class="border-b-[1px] border-surface-200-800 w-full">
    <h5 {...filesItem.heading} class="h5 w-full bg-secondary-600-400">
      <button {...filesItem.trigger} class="w-full p-2"> Shared Files </button>
    </h5>
    {#if filesItem.isExpanded}
      <div {...filesItem.content} transition:slide class="p-4">
        <article class="place-items-center w-full h-full grid">
          <span class="w-full relative mb-4">
            <button
              class="w-full card preset-filled-surface-200-800 hover:bg-surface-300-700 p-8 place-items-center grid text-surface-600-400"
              onclick={addFilesToWebShare}
            >
              <UploadIcon size={70} color="var(--color-surface-400-600)" />
              Select files to share
            </button>
            <button
              class="absolute bottom-0 right-0 text-error-700-300 p-2"
              onclick={clearAllSharedFiles}>Clear All</button
            >
          </span>
          <div class="bg-surface-50-950 w-full p-2">
            {#if webShareStatus && webShareStatus.sharedFiles.length > 0}
              {#each webShareStatus.sharedFiles as file (file.uuid)}
                <div
                  transition:slide={{ duration: 100 }}
                  class="flex flex-2 text-[12px] text-surface-800-200 items-center"
                >
                  <div class="flex-grow flex flex-col p-1.5">
                    <span class="font-medium">{file.name}</span>
                    <span class="text-sm text-surface-600-400"
                      >{fsizeText(file.size)}</span
                    >
                  </div>
                  <button
                    class="btn m-0 p-1.5"
                    onclick={() => removeSharedFile(file.uuid)}
                  >
                    <CloseIcon size={20} color="var(--color-error-600-400)" />
                  </button>
                </div>
              {/each}
            {:else}
              <div
                transition:slide
                class="p-3 text-surface-300-700 text-center"
              >
                No files are currently shared.
              </div>
            {/if}
          </div>
        </article>
      </div>
    {/if}
  </div>
</div>
