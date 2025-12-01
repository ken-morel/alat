<script lang="ts">
  import "../app.css";
  import Navigation from "./Navigation.svelte";
  import StatusBar from "./StatusBar.svelte";
  import PairDialog from "./PairDialog.svelte";
  import { ConfigReady } from "$lib/wails/wailsjs/go/app/App";
  import { page } from "$app/stores";
  let { children } = $props();
</script>

<div class="bg-surface-950 flex h-screen flex-col">
  <div
    class="content flex w-full flex-grow"
    style="max-height: calc(100% - 72px);"
  >
    {#key $page.url.pathname}
      {#await ConfigReady()}
        <Navigation />
      {:then ready}
        {#if ready}
          <Navigation />
        {/if}
      {/await}
    {/key}
    <div class="w-full overflow-y-auto p-8">
      {@render children?.()}
    </div>
  </div>
  <div class="">
    <StatusBar />
  </div>
  <PairDialog />
</div>
