<script lang="ts">
  import { goto } from "$app/navigation";
  import DeviceTile from "$lib/components/DeviceTile.svelte";
  import { selectedPairedDevice } from "$lib/state";
  import type { pair } from "$lib/wailsjs/go/models";
  import { flip } from "svelte/animate";
  let { data } = $props();
  let { pairedDevicesPromise } = data;
  function handlePairClick(pair: pair.Pair) {
    selectedPairedDevice.set(pair);
    goto("/device");
  }
</script>

<div class="dashboard-container">
  <h2 class="welcome-header">Welcome! 👋</h2>
  <section class="connected-devices">
    <div class="content-wrapper">
      <h3 class="section-title">Connected devices</h3>
      <div class="devices-list">
        {#await pairedDevicesPromise}
          <div class="loading-text">loading devices...</div>
        {:then pairedDevices}
          {#each pairedDevices as pair (pair.Token)}
            <div
              animate:flip
              onclick={() => handlePairClick(pair)}
              role="button"
              tabindex="0"
              onkeydown={null}
            >
              <DeviceTile
                deviceInfo={pair.DeviceInfo}
                services={pair.Services}
              />
            </div>
          {:else}
            <div class="no-devices-message">
              <h2>No paired device. For now</h2>
            </div>
          {/each}
        {/await}
      </div>
      <a href="/pair" class="btn btn-primary pair-button">Pair a device</a>
    </div>
  </section>
</div>

<style lang="sass">
  @use '$lib/styles/theme'

  .dashboard-container
    max-width: 800px
    margin: auto
    padding: 2rem 1.5rem

  .welcome-header
    font-size: 4rem
    font-weight: 200
    text-align: center
    padding: 3rem 0
    border-bottom: 1px solid theme.$primary-d3
    margin-bottom: 2rem

  .section-title
    font-size: 2rem
    text-align: center
    margin-bottom: 2rem

  .devices-list
    display: flex
    flex-direction: row
    flex-wrap: wrap
    justify-content: center
    gap: 1.5rem

  .loading-text, .no-devices-message
    display: flex
    align-items: center
    justify-content: center
    width: 100%
    font-size: 1.5rem
    color: theme.$text-secondary
    opacity: 0.7

  .no-devices-message h2
    font-weight: 300
    color: theme.$text-secondary

  .pair-button
    display: block
    width: 100%
    max-width: 400px
    margin: 2.5rem auto 0
    text-align: center
</style>
