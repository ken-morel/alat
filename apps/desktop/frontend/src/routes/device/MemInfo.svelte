<script lang="ts">
  import type { pbuf } from "$lib/wailsjs/go/models";

  let { info: memory }: { info: pbuf.VirtualMemoryStat } = $props();
</script>

<div class="memory-info-container">
  <h4>Memory</h4>
  <span class="battery">
    <progress
      value={memory.used_percent}
      max="100"
      style:background-color="green"
    >
    </progress>
  </span>
</div>

<style lang="sass">
div.memory-info-container
  span.battery
    position: relative
    progress
      height: 20px
</style>
