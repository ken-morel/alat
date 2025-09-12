<script lang="ts">
  import { connectedDevice } from "$lib/store";
  import UploadIcon from "@lucide/svelte/icons/upload";
  const dev = connectedDevice;
</script>

<div class="h-full w-full grid place-items-center">
  {#if $dev}
    <div
      class="card preset-filled-surface-100-900 border-[1px] border-surface-200-800 w-full max-w-lg"
    >
      <header class="flex border-b border-surface-200-800 p-8">
        <h3 class="ml-8 h4">Send files to {$dev.Info.Name}</h3>
      </header>
      <article class="place-items-center w-full h-full grid p-10">
        <div
          class="card preset-filled-surface-200-800 p-8 place-items-center grid text-surface-600-400"
        >
          <UploadIcon size={70} color="var(--color-surface-400-600)" />
          Select files
        </div>
      </article>
    </div>
  {:else}
    <div class="card preset-filled-error-400-600 p-8">
      <p class="h2 text-error-700-300">No device selected.</p>
    </div>
  {/if}
</div>
