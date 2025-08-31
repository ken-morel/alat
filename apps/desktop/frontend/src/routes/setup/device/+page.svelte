<script lang="ts">
  import { nextUrl, prevUrl } from "../wizard.svelte";
  import {
    SettingsSetDeviceColor,
    SettingsSetDeviceName,
  } from "$lib/wails/wailsjs/go/app/App";
  import type { PageData } from "./$types";

  const { data }: { data: PageData } = $props();

  let deviceName: string = $state(data.deviceName);
  let deviceColor: string = $state(data.deviceColor);
  let error: string | null = $state(null);

  prevUrl.set("/setup");
  nextUrl.set("/setup/services");

  $effect(() => {
    deviceName = deviceName.trim();
    if (deviceName.length === 0) {
      error = "Device name cannot be empty.";
    } else {
      SettingsSetDeviceName(deviceName).catch((err) => {
        console.error("Failed to set device name:", err);
      });
    }
    SettingsSetDeviceColor(deviceColor).then((e) =>
      console.error("Failed setting device color: ", e),
    );
  });
</script>

<div class="">
  <h3 class="h3">Welcome to Alat</h3>
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
        class:ring-2={deviceColor === color.hex}
        class:ring-primary-500={deviceColor === color.hex}
        class:ring-offset-2={deviceColor === color.hex}
        style="background-color: {color.hex};"
        onclick={() => (deviceColor = color.hex)}
        aria-label="Select color {color.name}"
      ></button>
    {/each}
  </div>
</label>
