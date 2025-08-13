<script lang="ts">
  let { data } = $props();
  let { pairedDevicesPromise } = data;
</script>

<h2 class="w3-jumbo w3-padding-64 w3-center w3-border-bottom">Welcome! 👋</h2>
<section class="connected-devices w3-container">
  <h3 class="w3-large w3-center w3-xxlarge">Connected devices</h3>
  <div class="devices-list w3-center w3-container">
    {#await pairedDevicesPromise}
      <div class="loading-devices w3-xlarge w3-opacity">loading devices...</div>
    {:then pairedDevices}
      {#each pairedDevices as device}
        <div class="device-tile w3-card">
          <span class="logo">
            <span></span>
          </span>
          <span class="info w3-panel">
            <span class="name">{device.Name}</span>
            <span class="address">
              <pre>{device.Address.IP}@{device.Address.Port}</pre>
            </span>
          </span>
        </div>
      {/each}
    {/await}
    <button class="w3-button">Pair a device</button>
  </div>
</section>

<style lang="sass">
@use '$lib/styles/theme'

h2
  background-color: theme.$secondary-d4
  margin: 0
  border-color: theme.$border-light !important
section.connected-devices
  background-color: theme.$secondary-d4
  div.devices-list
    max-width: 800px
    margin: auto
    display: flex
    div.device-tile
      margin: 20px
      display: flex
      max-width: 350px
      background-color: theme.$secondary-d3
      span.logo
        height: 100%
        width: 100px
        padding: 10px
        span
          display: inline-block
          width: 100px
          height: 100px
          border-radius: 50%
          background-color: theme.$tertiary
      span.info
        flex-grow: 1
        margin-left: 50px
        text-align: left
        span.name
          display: block
          font-size: xx-large
        span.address
          display: block
</style>
