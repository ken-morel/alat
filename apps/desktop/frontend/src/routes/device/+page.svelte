<script lang="ts">
  import { goto } from "$app/navigation";
  import { ICONS } from "$lib";
  import Color from "$lib/color";
  import { selectedPairedDevice } from "$lib/state";
  import { onMount } from "svelte";
  import RcFile from "./RcFile.svelte";
  import SysInfo from "./SysInfo.svelte";
  import type { pair } from "$lib/wailsjs/go/models";

  function supports(name: string, device: pair.Pair | null): boolean {
    if (!device) return false;
    for (let service of device.Services) {
      if (service.Name == name) return service.Enabled;
    }
    return false;
  }

  let supportsRCfile: boolean = $derived(
    supports("rcfile", $selectedPairedDevice),
  );
  let supportsSysInfo: boolean = $derived(
    supports("sysinfo", $selectedPairedDevice),
  );
  let device: pair.Pair | null = $derived($selectedPairedDevice);
  onMount(() => {});
</script>

{#if device}
  <section class="hero">
    <div class="profile">
      <span
        class="logo"
        style="color: {Color.fromGO(device.DeviceInfo.Color).toHexString()};"
        >{ICONS[device.DeviceInfo.Type]}</span
      >
      <span class="name">{device.DeviceInfo.Name}</span>
    </div>
    <div>
      <ul>
        <li>
          <span>Address: </span>
          <span>
            <code>{device.DeviceInfo.Address.Phrase}</code>
          </span>
        </li>
        <li>
          <span>IP Address: </span>
          <span>
            <code>{device.DeviceInfo.Address.IP}</code>
          </span>
        </li>
      </ul>
    </div>
  </section>
  <section class="service-container">
    <RcFile enabled={supportsRCfile} {device} />
  </section>
  <section class="service-container">
    <SysInfo enabled={supportsSysInfo} {device} />
  </section>
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
section.service-container
  border-top: 1px theme.$border-dark solid
  padding: 32px

</style>
