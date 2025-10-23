<script lang="ts">
  import { slide } from "svelte/transition";
  import { page } from "$app/stores";
  import { nextUrl, prevUrl } from "./wizard.svelte";

  let { children } = $props();

  nextUrl.set(null);
  prevUrl.set(null);
</script>

<div class="h-full w-full grid place-items-center">
  <div
    class="card preset-filled-surface-100-900 border-[1px] border-surface-200-800 w-full max-w-lg"
  >
    <div class="overflow-hidden">
      {#key $page.url.pathname}
        <div
          class="p-8"
          in:slide|local={{ duration: 200, delay: 200 }}
          out:slide|local={{ duration: 200 }}
        >
          {@render children?.()}
        </div>
      {/key}
    </div>

    <footer
      class="flex items-center border-t border-surface-200-800"
      class:justify-between={$prevUrl}
      class:justify-end={!$prevUrl}
    >
      {#if $prevUrl}
        <a href={$prevUrl} class="btn">Back</a>
      {/if}

      {#if $nextUrl}
        <a href={$nextUrl} class="btn variant-filled"
          >{$nextUrl == "/" ? "Complete" : "Next"}</a
        >
      {/if}
    </footer>
  </div>
</div>
