<script lang="ts">
  import { nextUrl, prevUrl } from "../wizard.svelte";
  import {
    SettingsSetDeviceColorName,
    SettingsSetDeviceName,
  } from "$lib/wails/wailsjs/go/app/App";
  import type { PageData } from "./$types";

  const { data }: { data: PageData } = $props();

  let deviceName: string = $state(data.deviceName);
  let deviceColorName: string = $state(data.deviceColorName);
  let error: string | null = $state(null);

  prevUrl.set("/setup");
  nextUrl.set("/setup/services");

  $effect(() => {
    if (deviceName.length === 0) {
      error = "Device name cannot be empty.";
    } else {
      SettingsSetDeviceName(deviceName.trim()).catch((err) => {
        console.error("Failed to set device name:", err);
      });
    }
    SettingsSetDeviceColorName(deviceColorName).then((e) =>
      console.error("Failed setting device color: ", e),
    );
  });
</script>

<div class="">
  <h3 class="h3">Your device</h3>
  <p class="text-surface-500">Let's get your device set up.</p>
</div>
<div style:padding-top="20px"></div>

<label class="label py-4">
  <span>Device Name</span>
  <input
    class="input"
    type="text"
    bind:value={deviceName}
    placeholder="My Awesome Laptop"
  />
  {#if error}
    <span class="text-error-500">{error}</span>
  {/if}
</label>
<label class="label py-4">
  <span>Device Color</span>

  <div class="grid grid-cols-10 gap-4">
    {#each data.alatColors as color}
      <button
        type="button"
        class="w-full aspect-square rounded-full transition-all"
        class:ring-2={deviceColorName === color.name}
        class:ring-primary-500={deviceColorName === color.name}
        class:ring-offset-2={deviceColorName === color.name}
        style="background-color: {color.hex};"
        onclick={() => (deviceColorName = color.name)}
        aria-label="Select color {color.name}"
      ></button>
    {/each}
  </div>
</label>
