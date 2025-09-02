<script lang="ts">
  import "../app.css";
  import "../color-fixer.js";
  import Navigation from "./Navigation.svelte";
  import StatusBar from "./StatusBar.svelte";
  import PairDialog from "./PairDialog.svelte";
  let { children, data } = $props();
</script>

<div class="flex h-screen flex-col">
  <div class="content flex w-full flex-grow overflow-y-auto">
    {#if data.ready}
      <Navigation />
    {/if}
    {@render children?.()}
  </div>
  <div class="h-16 w-full">
    <StatusBar />
  </div>
  <PairDialog />
</div>
