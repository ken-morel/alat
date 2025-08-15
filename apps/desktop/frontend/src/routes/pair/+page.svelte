<script lang="ts">
  import DeviceTile from "$lib/components/DeviceTile.svelte";
  import { SearchDevices } from "$lib/wailsjs/go/app/App";
  import { device } from "$lib/wailsjs/go/models";
  import { onMount } from "svelte";

  let {} = $props();
  let tdots: string = $state(".");
  let deviceInfos: device.DeviceInfo[] = $state([]);
  let error: string | null = $state(null);

  onMount(() => {
    const tdotsInterval = setInterval(() => {
      tdots = tdots.length < 3 ? tdots + "." : ".";
    }, 500);

    SearchDevices()
      .then((result: device.DeviceInfo[]) => {
        if (result) deviceInfos = result;
      })
      .catch((err: Error) => {
        error =
          "Could not find any local addresses. Make sure you are connected to a network.";
        console.error(err);
      });

    return () => {
      clearInterval(tdotsInterval);
    };
  });
</script>

<h2 class="w3-center w3-xxlarge w3-padding-32">Pair a device</h2>

<section class="w3-container devices-list-container">
  <div>
    {#if error}
      <div class="w3-panel w3-red w3-center">
        <p>{error}</p>
      </div>
    {:else if deviceInfos.length === 0}
      <div class="title">
        <h4>Searching for devices {tdots}</h4>
      </div>
    {:else}
      <div class="title">
        <h4>Select a device to pair with:</h4>
      </div>
      <div class="device-list">
        {#each deviceInfos as info}
          <DeviceTile deviceInfo={info} />
        {/each}
      </div>
    {/if}
    <div class="devices-list"></div>
  </div>
</section>

<style lang="sass">
@use "$lib/styles/theme"
h2
  background-color: theme.$secondary-d4
  margin: 0
section
  background-color: theme.$secondary-d4
section.devices-list-container
  div
    max-width: 800px
    margin: auto
    div.title
      text-align: center
    div.devices-list
      display: flex

</style>
