<script lang="ts">
  import { GetPairedDeviceSysInfo } from "$lib/wailsjs/go/app/App";
  import type { pair, pbuf } from "$lib/wailsjs/go/models";

  let { device }: { device: pair.Pair } = $props();
  let batteries: pbuf.Battery[] | null = $state(null);
  GetPairedDeviceSysInfo(device).then((info) => {
    batteries = info.battery ?? null;
  });
</script>

<div class="battery-info-container">
  <h4>Batteries</h4>
  {#if batteries}
    {#each batteries as battery}
      {#if battery && battery.current_capacity && battery.full_charged_capacity}
        <span class="battery">
          <progress
            value={(battery.current_capacity / battery.full_charged_capacity) *
              100}
            max="100"
            style:background-color="green"
          >
          </progress>
          <span>
            {battery.state}
          </span>
        </span>
      {/if}
    {/each}
  {:else}
    <span class="battery">
      <progress style:background-color="green"> </progress>
      <span> Fetching ... </span>
    </span>
  {/if}
</div>

<style lang="sass">
div.battery-info-container
  font-size: medium
  span.battery
    margin: 10px
    position: relative
    progress
      height: 20px
      width: 100px
    span
      position: absolute
      right: 0
</style>
