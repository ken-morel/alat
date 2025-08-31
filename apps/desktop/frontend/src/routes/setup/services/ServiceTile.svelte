<script lang="ts">
  import { nextUrl, prevUrl } from "../wizard.svelte";
  import { slide } from "svelte/transition";

  import { Switch } from "@skeletonlabs/skeleton-svelte";
  import type { Snippet } from "svelte";

  let {
    next,
    prev,
    title,
    description,
    enabled = $bindable(),
    disabled = false,
    children,
  }: {
    next: string | null;
    prev: string | null;
    title: string;
    enabled: boolean;
    disabled?: boolean;
    description: string;
    children?: Snippet;
  } = $props();

  prevUrl.set(prev);
  nextUrl.set(next);
</script>

<div class="space-y-4">
  <header>
    <div class="flex justify-between">
      <h1 class="text-2xl font-bold">{title}</h1>
      <Switch
        label="Enabled"
        onCheckedChange={(v) => (enabled = v.checked)}
        checked={enabled}
        {disabled}
      />
    </div>
    <p class="text-sm text-surface-300">
      {description}
    </p>
  </header>

  {#if enabled}
    <div transition:slide={{ duration: 100 }}>
      <div class="space-y-2 border-l-2 border-surface-200-800 pl-4">
        {#if children}
          {@render children()}
        {:else}
          <p class="text-surface-500">No configuration needed.</p>
        {/if}
      </div>
    </div>
  {/if}
</div>
