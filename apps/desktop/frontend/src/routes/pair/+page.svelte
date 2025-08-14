<script lang="ts">
  import { GetAvailableAddresses } from "$lib/wailsjs/go/app/App";
  import type { core } from "$lib/wailsjs/go/models";
  import { onMount } from "svelte";

  let {} = $props();
  let tdots: string = $state(".");
  let addresses: core.DeviceAddress[] = $state([]);
  let error: string | null = $state(null);

  onMount(() => {
    const tdotsInterval = setInterval(() => {
      tdots = tdots.length < 3 ? tdots + "." : ".";
    }, 500);

    GetAvailableAddresses()
      .then((result: any) => {
        addresses = result;
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
    {:else if addresses.length === 0}
      <div class="title">
        <h4>Searching for devices {tdots}</h4>
      </div>
    {:else}
      <div class="title">
        <h4>Select a device to pair with:</h4>
      </div>
      <ul class="w3-ul w3-card-4">
        {#each addresses as addr}
          <li class="w3-bar">
            <div class="w3-bar-item">
              <span class="w3-large">{addr.Phrase}</span><br />
              <span>{addr.IP}</span>
            </div>
          </li>
        {/each}
      </ul>
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
    ul
        margin-top: 16px
        li
            cursor: pointer
            &:hover
                background-color: theme.$secondary-d1

</style>
