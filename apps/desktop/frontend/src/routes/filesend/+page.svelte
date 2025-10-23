<script lang="ts">
  import { Accordion, Tabs, type AccordionItem } from "melt/builders";
  import { slide } from "svelte/transition";
  import { connectedDevices } from "$lib/connected";
  import guessIcon from "$lib/icons";
  import { connected } from "$lib/wails/wailsjs/go/models";

  let choosedDevices: string[] = $state([]);
  let alatShare: boolean = $state(true);

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
    WebShareStart,
    WebShareStop,
    WebShareGetStatus,
    WebShareAddSharedFiles,
    WebShareRemoveSharedFile,
    WebShareClearSharedFiles,
    WebShareSetPasscode,
    WebShareGetPasscode,
  } from "$lib/wails/wailsjs/go/app/App";
  import type { webshare } from "$lib/wails/wailsjs/go/models";
  import fsizeText from "$lib/fsize";
  import { get } from "svelte/store";
  import type { app } from "$lib/wails/wailsjs/go/models";
  import { onMount } from "svelte";

  let webShareStatus: webshare.Status | null = $state(null);
  let newPasscode: string = $state("");
  let filesToShare: FileList | null = $state(null);

  onMount(async () => {
    webShareStatus = await WebShareGetStatus();
  });

  const tabNames = ["filesend", "browsershare"];
  const tabs = new Tabs({
    value: tabNames[0],
  });

  function selectFiles() {
    AskFilesSend().then((selectedFiles) => {
      sendingFiles.update((files) => {
        return selectedFiles.concat(files);
      });
    });
  }

  async function startWebShare() {
    try {
      const port = await WebShareStart();
      console.log(`WebShare started on port ${port}`);
      await refreshWebShareStatus();
    } catch (error) {
      console.error("Failed to start WebShare:", error);
    }
  }

  async function stopWebShare() {
    try {
      await WebShareStop();
      console.log("WebShare stopped");
      await refreshWebShareStatus();
    } catch (error) {
      console.error("Failed to stop WebShare:", error);
    }
  }

  async function setWebSharePasscode() {
    if (newPasscode) {
      try {
        await WebShareSetPasscode(newPasscode);
        console.log("Passcode updated");
        await refreshWebShareStatus();
        newPasscode = ""; // Clear input after setting
      } catch (error) {
        console.error("Failed to set passcode:", error);
      }
    }
  }

  async function addFilesToWebShare() {
    if (filesToShare && filesToShare.length > 0) {
      const paths = Array.from(filesToShare).map((file) => file.path);
      try {
        await WebShareAddSharedFiles(paths);
        console.log("Files added to web share");
        await refreshWebShareStatus();
        filesToShare = null; // Clear input
      } catch (error) {
        console.error("Failed to add files to web share:", error);
      }
    }
  }

  async function removeSharedFile(uuid: string) {
    try {
      await WebShareRemoveSharedFile(uuid);
      console.log(`File with UUID ${uuid} removed`);
      await refreshWebShareStatus();
    } catch (error) {
      console.error("Failed to remove file:", error);
    }
  }

  async function clearAllSharedFiles() {
    try {
      await WebShareClearSharedFiles();
      console.log("All shared files cleared");
      await refreshWebShareStatus();
    } catch (error) {
      console.error("Failed to clear all shared files:", error);
    }
  }

  async function refreshWebShareStatus() {
    webShareStatus = await WebShareGetStatus();
  }
</script>

<div class="w-full h-full grid place-items-center transition-all">
  <div
    class="card preset-filled-surface-100-900 border-[1px] border-surface-200-800 w-full max-w-lg"
  >
    <header
      {...tabs.triggerList}
      class="border-b-2 border-surface-200-800 flex flex-2 place-content-evenly"
    >
      <span class="p-2 pl-2" {...tabs.getTrigger(tabNames[0])}>
        <h4 class="h4">Send files</h4>
      </span>
      <span class="p-2 pl-2" {...tabs.getTrigger(tabNames[1])}>
        <h4 class="h4">Browser share</h4>
      </span>
    </header>
    <div id="tabs-content" style="transition: all 0.4s;">
      <div {...tabs.getContent(tabNames[0])}>
        <div {...accordion.root}>
          <div
            class="border-b-[1px] border-surface-200-800 w-full"
            transition:slide
          >
            <h5 {...files.heading} class="h5 w-full bg-secondary-600-400">
              <button {...files.trigger} class="w-full"> Select files </button>
            </h5>
            {#if files.isExpanded}
              <div {...files.content} transition:slide>
                <article class="place-items-center w-full h-full grid p-10">
                  <span class="w-full relative"
                    ><button
                      class="w-full card preset-filled-surface-200-800 hover:bg-surface-300-700 p-8 place-items-center grid text-surface-600-400"
                      onclick={selectFiles}
                    >
                      <UploadIcon
                        size={70}
                        color="var(--color-surface-400-600)"
                      />
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
                          <CloseIcon
                            size={20}
                            color="var(--color-error-600-400)"
                          />
                        </button>
                      </div>
                      {#if idx < $sendingFiles.length - 1}
                        <hr
                          class="border-b-[1px] border-surface-300-700 w-full"
                        />
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
              <button {...devices.trigger} class="w-full">
                Select devices
              </button>
            </h5>
            {#if devices.isExpanded}
              <div transition:slide class="p-7">
                <div {...devices.content}>
                  {#each $connectedDevices as dev}
                    {@const Icon = guessIcon(dev.info.type)}
                    <div
                      class:bg-surface-200-800={choosedDevices.includes(
                        dev.info.id,
                      )}
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
                      <p class="text-surface-600-400 p-4">
                        No connected device
                      </p>
                    </div>
                  {/each}
                </div>
              </div>
            {/if}
          </div>
          <div>
            <button class="btn preset-filled-surface-300-700 p-3 w-full">
              Send {$sendingFiles.length} files to {choosedDevices.length} devices
            </button>
          </div>
        </div>
      </div>
      <div {...tabs.getContent(tabNames[1])}>
        <div class="container p-4">
          <h4 class="h4 mb-4">Browser Share</h4>

          <!-- Server Status and Controls -->
          <div class="card mb-4 p-4">
            <h2 class="h5 mb-3">Server Status</h2>
            {#if webShareStatus}
              <p class="mb-2">
                Status: <span class="font-bold"
                  >{webShareStatus.isRunning ? "Running" : "Stopped"}</span
                >
              </p>
              {#if webShareStatus.isRunning}
                <p class="mb-2">
                  Port: <span class="font-bold">{webShareStatus.port}</span>
                </p>
                <p class="mb-2">
                  Share URL: <a
                    href={webShareStatus.shareURL}
                    target="_blank"
                    rel="noopener noreferrer"
                    class="text-blue-500 hover:underline"
                    >{webShareStatus.shareURL}</a
                  >
                </p>
                <p class="mb-4">
                  Passcode: <span class="font-bold"
                    >{webShareStatus.passcode}</span
                  >
                </p>
                <button
                  class="btn preset-filled-error-500-400 p-2 w-full"
                  onclick={stopWebShare}
                >
                  Stop Server
                </button>
              {:else}
                <button
                  class="btn preset-filled-success-500-400 p-2 w-full"
                  onclick={startWebShare}
                >
                  Start Server
                </button>
              {/if}
            {:else}
              <p>Loading status...</p>
            {/if}
          </div>

          <!-- Set Passcode -->
          <div class="card mb-4 p-4">
            <h2 class="h5 mb-3">Set Custom Passcode</h2>
            <div class="flex gap-2">
              <input
                type="text"
                class="input flex-grow"
                placeholder="New Passcode (e.g., 6-digit code)"
                bind:value={newPasscode}
              />
              <button
                class="btn preset-filled-primary-500-400 p-2"
                onclick={setWebSharePasscode}
                disabled={!newPasscode}
              >
                Set Passcode
              </button>
            </div>
          </div>

          <!-- Upload Files for Sharing -->
          <div class="card mb-4 p-4">
            <h2 class="h5 mb-3">Add Files to Share</h2>
            <div class="flex gap-2">
              <input
                type="file"
                class="input flex-grow"
                onchange={(e) => (filesToShare = e.target?.files)}
                multiple
              />
              <button
                class="btn preset-filled-primary-500-400 p-2"
                onclick={addFilesToWebShare}
                disabled={!filesToShare || filesToShare.length === 0}
              >
                Add Files
              </button>
            </div>
          </div>

          <!-- List Shared Files -->
          <div class="card p-4">
            <h2 class="h5 mb-3">Currently Shared Files</h2>
            {#if webShareStatus && webShareStatus.sharedFiles.length > 0}
              <ul class="list-none p-0 m-0">
                {#each webShareStatus.sharedFiles as file (file.uuid)}
                  <li
                    class="flex justify-between items-center py-2 border-b border-surface-200-800 last:border-b-0"
                  >
                    <div>
                      <p class="font-medium">{file.name}</p>
                      <p class="text-sm text-surface-600-400">
                        {fsizeText(file.size)}
                      </p>
                    </div>
                    <button
                      class="btn preset-filled-error-500-400 p-1.5 text-sm"
                      onclick={() => removeSharedFile(file.uuid)}
                    >
                      Remove
                    </button>
                  </li>
                {/each}
              </ul>
              <button
                class="btn preset-filled-error-500-400 p-2 w-full mt-4"
                onclick={clearAllSharedFiles}
              >
                Clear All Shared Files
              </button>
            {:else}
              <p class="text-center text-surface-600-400">
                No files are currently shared via browser.
              </p>
            {/if}
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

<style lang="sass">

</style>
