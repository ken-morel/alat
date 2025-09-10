<script lang="ts">
  import { connectedDevice } from "$lib/store";
  let dev = $derived($connectedDevice);
</script>

<div class="h-full w-full grid place-items-center">
  {#if dev}
    <div
      class="card preset-filled-surface-100-900 border-[1px] border-surface-200-800 w-full max-w-lg"
    >
      <header class="border-b border-surface-200-800 p-8">
        <h3 class="h3">{dev.Info.Name}</h3>
      </header>
      <article class="p-8"></article>
    </div>
  {:else}
    <div class="card preset-filled-error-400-600 p-8">
      <p class="h2 text-error-700-300">No device selected.</p>
    </div>
  {/if}
</div>
