<script lang="ts">
  import { GetPairedDeviceSysInfo } from "$lib/wailsjs/go/app/App";
  import type { pair, pbuf } from "$lib/wailsjs/go/models";
  import { onMount } from "svelte";
  import BatteryInfo from "./BatteryInfo.svelte";
  import MemInfo from "./MemInfo.svelte";

  let { enabled, device }: { enabled: boolean; device: pair.Pair } = $props();
  let info: pbuf.SysInfo | null = $state(null);
  let error: string | null = $state(null);
  async function fetchInfo() {
    GetPairedDeviceSysInfo(device).then((newInfo) => {
      info = newInfo;
    });
  }
  onMount(() => {
    fetchInfo();
    const interval = setInterval(fetchInfo, 5000);
    return () => {
      clearInterval(interval);
    };
  });
</script>

{#snippet message(color: string, msg: string, opacity: boolean = false)}
  <div class:color class:w3-opacity={opacity} class="w3-container">
    <p>{msg}</p>
  </div>
{/snippet}

{#if enabled}
  {#if error}
    {@render message("red", "Ough Scrap!, got no info:" + error)}
  {:else if info}
    <div class="w3-container">
      <span class="battery">
        {#if info.battery}
          <BatteryInfo info={info.battery} />
        {:else}
          <span>No battery available</span>
        {/if}
      </span>
      <span class="memory">
        {#if info.memory}
          <MemInfo info={info.memory} />
        {:else}
          <span>No memory available</span>
        {/if}
      </span>
    </div>
  {:else}
    {@render message("blue", "Getting system information...", true)}
  {/if}
{/if}

<style lang="sass">
@use '$lib/styles/theme'

div
  display: flex
  span.battery
    border-right: 1px theme.$border-dark solid
    padding: 10px
</style>
