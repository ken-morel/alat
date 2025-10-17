<script lang="ts">
  import "../app.css";
  // import "../color-fixer.js";
  import Navigation from "./Navigation.svelte";
  import StatusBar from "./StatusBar.svelte";
  import PairDialog from "./PairDialog.svelte";
  import { ConfigReady } from "$lib/wails/wailsjs/go/app/App";
  import { ProgressRing } from "@skeletonlabs/skeleton-svelte";
  import { page } from "$app/stores";
  let { children } = $props();
</script>

<div class="flex h-screen flex-col">
  <div class="content flex w-full flex-grow overflow-y-auto">
    {#key $page.url.pathname}
      {#await ConfigReady()}
        <Navigation />
      {:then ready}
        {#if ready}
          <Navigation />
        {/if}
      {/await}
    {/key}
    {@render children?.()}
  </div>
  <div class="h-18 w-full">
    <StatusBar />
  </div>
  <PairDialog />
</div>
