<script lang="ts">
  import DeviceTile from "$lib/components/DeviceTile.svelte";
  import ServiceConfigModal from "$lib/components/ServiceConfigModal.svelte";
  import { selectedDeviceForPairing } from "$lib/state";
  import { SearchDevices } from "$lib/wailsjs/go/app/App";
  import { device } from "$lib/wailsjs/go/models";
  import { onMount } from "svelte";

  let {} = $props();
  let tdots: string = $state(".");
  let deviceInfos: device.DeviceInfo[] = $state([]);
  let error: string | null = $state(null);

  function handleDeviceClick(info: device.DeviceInfo) {
    selectedDeviceForPairing.set(info);
  }

  onMount(() => {
    const tdotsInterval = setInterval(() => {
      tdots = tdots.length < 3 ? tdots + "." : ".";
    }, 500);

    SearchDevices()
      .then((result: device.DeviceInfo[]) => {
        setTimeout(() => {
          if (result) deviceInfos = result;
        }, 1500);
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

<ServiceConfigModal />

<h2 class="w3-center w3-xxlarge w3-padding-32">Connect a device</h2>

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
        <h4>Select a device to connect to:</h4>
      </div>
      <div class="device-list">
        {#each deviceInfos as info}
          <div
            onclick={() => handleDeviceClick(info)}
            role="button"
            tabindex="0"
            onkeydown={null}
          >
            <DeviceTile deviceInfo={info} />
          </div>
        {/each}
      </div>
    {/if}
    <div class="devices-list"></div>
  </div>
</section>

<style lang="sass">
@use "$lib/styles/theme"
h2
  background-color: theme.$background
  margin: 0
section
  background-color:  theme.$background
section.devices-list-container
  border-top: 2px theme.$primary-d3 solid
  border-bottom: 2px theme.$primary-d3 solid
  padding: 30px
  div
    max-width: 800px
    margin: auto
    div.title
      text-align: center
    div.device-list
      display: flex
      flex-direction: row
      flex-wrap: wrap
      align-content: flex-start
      justify-content: space-evenly
      align-items: flex-start

</style>
