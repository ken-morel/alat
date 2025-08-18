<script lang="ts">
  import { GetAndSendFiles } from "$lib/wailsjs/go/app/App";
  import type { pair } from "$lib/wailsjs/go/models";

  let { enabled, device }: { enabled: boolean; device: pair.Pair } = $props();

  function sendFile() {
    GetAndSendFiles(device);
  }
</script>

{#if enabled}
  <div class="w3-display-container w3-padding-16 w3-padding-right">
    <span></span>
    <h3>Send file</h3>
    <span></span>
    <button type="button" class="w3-button" onclick={sendFile}>
      Choose file
    </button>
  </div>
{:else}
  <p class="w3-opacity">File send disabled...</p>
{/if}

<style lang="sass">
@use '$lib/styles/theme'
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
