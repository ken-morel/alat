<script lang="ts">
  import { Tooltip } from "melt/builders";
  import type { Snippet } from "svelte";
  import { scale } from "svelte/transition";
  let isOpen: boolean = $state(true);
  const builder = new Tooltip({
    closeDelay: 200,
    closeOnPointerDown: true,
    openDelay: 400,
    onOpenChange: (value) => (isOpen = value),
  });
  let {
    children,
    tooltip,
    classes = "rounded-2xl",
  }: { children: Snippet; tooltip: Snippet; classes: string } = $props();
</script>

<div {...builder.trigger}>
  {@render children?.()}
</div>
{#key isOpen}
  <div
    {...builder.content}
    in:scale={{ opacity: 0.5, duration: 200, start: 0.9 }}
    class={classes}
  >
    <div {...builder.arrow} class="h-5 w-5 bg-green"></div>
    {@render tooltip?.()}
  </div>
{/key}
