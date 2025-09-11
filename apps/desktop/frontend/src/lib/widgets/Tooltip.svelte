<script lang="ts">
  import { Tooltip } from "melt/builders";
  import type { Snippet } from "svelte";
  import { scale } from "svelte/transition";
  let isOpen: boolean = $state(true);
  const builder = new Tooltip({
    closeDelay: 400,
    closeOnPointerDown: true,
    openDelay: 400,
    onOpenChange: (value) => (isOpen = value),
  });
  let { children, tooltip }: { children: Snippet; tooltip: Snippet } = $props();
</script>

<div {...builder.trigger}>
  {@render children?.()}
</div>
{#key isOpen}
  <div
    {...builder.content}
    in:scale={{ opacity: 0.5, duration: 200, start: 0.9 }}
    class="card preset-filled-surface-50-950"
  >
    <div {...builder.arrow}></div>
    {@render tooltip?.()}
  </div>
{/key}
