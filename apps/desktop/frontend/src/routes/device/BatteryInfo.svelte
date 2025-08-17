<script lang="ts">
  import type { pbuf } from "$lib/wailsjs/go/models";

  let { info: batteries }: { info: pbuf.Battery[] } = $props();
</script>

<div class="battery-info-container">
  <h4>Batteries</h4>
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
</div>

<style lang="sass">
div.battery-info-container
  span.battery
    position: relative
    progress
      height: 20px
    span
      position: absolute
      right: 0
</style>
