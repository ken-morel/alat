<script lang="ts">
  import { connectedDevice } from "$lib/store";
  import guessIcon from "$lib/icons";
  import UploadIcon from "@lucide/svelte/icons/upload";
  const dev = connectedDevice;
</script>

<div class="h-full w-full grid place-items-center">
  {#if $dev}
    {@const Icon = guessIcon($dev.Info.Type)}
    <div
      class="card preset-filled-surface-100-900 border-[1px] border-surface-200-800 w-full max-w-lg"
    >
      <header class="flex border-b border-surface-200-800 p-8">
        <Icon size={50} color={$dev.Info.Color.Hex} />
        <h3 class="ml-8 h3">{$dev.Info.Name}</h3>
      </header>
      <article class="p-4">
        <div class="card preset-filled-surface-200-800 border-2">
          <UploadIcon />
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
