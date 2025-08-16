<script lang="ts">
  import { goto } from "$app/navigation";
  import { ICONS } from "$lib";
  import Color from "$lib/color";
  import { selectedPairedDevice } from "$lib/state";
  import { GetAndSendFiles } from "$lib/wailsjs/go/app/App";
  import { onMount } from "svelte";
  import { get } from "svelte/store";

  function sendFile() {
    let device = get(selectedPairedDevice);
    if (device) GetAndSendFiles(device);
  }

  let supportsRCfile = $state(false);
  onMount(() => {
    selectedPairedDevice.subscribe((pair) => {
      if (!pair) return;
      for (let service of pair.Services) {
        switch (service.Name) {
          case "rcfile":
            supportsRCfile = service.Enabled;
            break;
        }
      }
    });
  });
</script>

{#if $selectedPairedDevice}
  <section class="hero">
    <div class="profile">
      <span
        class="logo"
        style="color: {Color.fromGO(
          $selectedPairedDevice.DeviceInfo.Color,
        ).toHexString()};">{ICONS[$selectedPairedDevice.DeviceInfo.Type]}</span
      >
      <span class="name">{$selectedPairedDevice.DeviceInfo.Name}</span>
    </div>
    <div>
      <ul>
        <li>
          <span>Address: </span>
          <span>
            <code>{$selectedPairedDevice.DeviceInfo.Address.Phrase}</code>
          </span>
        </li>
        <li>
          <span>IP Address: </span>
          <span>
            <code>{$selectedPairedDevice.DeviceInfo.Address.IP}</code>
          </span>
        </li>
      </ul>
    </div>
  </section>
  {#if supportsRCfile}
    <section class="sendfile w3-padding-32">
      <div class="w3-display-container w3-padding-16 w3-padding-right">
        <span></span>
        <h3>Send file</h3>
        <span></span>
        <button type="button" class="w3-button" onclick={sendFile}
          >Choose file</button
        >
      </div>
    </section>
  {/if}
{:else}
  {#await goto("/")}
    <p>Redirecting you back to home...</p>
  {/await}
{/if}

<style lang="sass">
@use '$lib/styles/theme'

section
  max-width: 600px
  margin: auto
section.hero
  div.profile
    font-size: xx-large
    text-align: center
    border-bottom: 1px theme.$border-dark solid
    .logo
      font-size: 1.5em
section.sendfile
  border-top: 1px theme.$border-dark solid
  div
    display: flex
    margin: 10px
    background-color: theme.$secondary-d4
    h3
      font-size: xx-large
    button
      background-color: theme.$secondary-d3
      margin-right: 20px
    span
      flex-grow: 1
</style>
