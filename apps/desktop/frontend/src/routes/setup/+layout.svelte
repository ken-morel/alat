<script lang="ts">
  import { slide } from "svelte/transition";
  import { page } from "$app/stores";
  import { wizardState } from "./wizard-state";

  let { children } = $props();

  $effect(() => {
    return () => {
      wizardState.nextUrl = null;
      wizardState.prevUrl = null;
    };
  });
</script>

<div class="h-screen w-screen grid place-items-center">
  <div
    class={[
      "card preset-filled-surface-100-900 border-[1px]",
      "border-surface-200-800 w-full max-w-lg p-8 space-y-6",
    ]}
  >
    <div class="overflow-hidden">
      {#key $page.url.pathname}
        <div
          in:slide|local={{ duration: 200, delay: 200 }}
          out:slide|local={{ duration: 200 }}
        >
          {@render children?.()}
        </div>
      {/key}
    </div>

    <footer
      class="flex items-center"
      class:justify-between={wizardState.prevUrl}
      class:justify-end={!wizardState.prevUrl}
    >
      {#if wizardState.prevUrl}
        <a href={wizardState.prevUrl} class="btn">Back</a>
      {/if}

      {#if wizardState.nextUrl}
        <a href={wizardState.nextUrl} class="btn variant-filled">Next</a>
      {/if}
    </footer>
  </div>
</div>
