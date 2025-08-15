<script lang="ts">
  import DeviceTile from "$lib/components/DeviceTile.svelte";

  let { data } = $props();
  let { pairedDevicesPromise } = data;
</script>

<h2 class="w3-jumbo w3-padding-64 w3-center w3-border-bottom">Welcome! 👋</h2>
<section class="connected-devices w3-container">
  <div>
    <h3 class="w3-large w3-center w3-xxlarge">Connected devices</h3>
    <div class="devices-list w3-center w3-container">
      {#await pairedDevicesPromise}
        <div class="loading-devices w3-xlarge w3-opacity">
          loading devices...
        </div>
      {:then pairedDevices}
        {#each pairedDevices as device}
          <DeviceTile deviceInfo={device} />
        {/each}
      {/await}
    </div>
    <a href="/pair" class="w3-button w3-block w3-margin-top">Pair a device</a>
  </div>
</section>

<style lang="sass">
@use '$lib/styles/theme'

h2
  background-color: theme.$background
  margin: 0
  border-color: theme.$border-dark !important
section.connected-devices
  background-color: theme.$background
  div
    max-width: 800px
    margin: auto
    a
      max-width: 90%
      background-color: theme.$primary-d3
      margin: auto
    div.devices-list
      display: flex
      flex-direction: row
      flex-wrap: wrap
      align-content: flex-start
      justify-content: space-evenly
      align-items: flex-start
</style>
