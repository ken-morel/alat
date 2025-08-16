<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import favicon from "$lib/assets/logo.svg";
  import "$lib/styles/app.sass";
  import { WasSetup } from "$lib/wailsjs/go/app/App";
  import { onMount } from "svelte";

  let { children } = $props();

  onMount(async () => {
    const isSetup = await WasSetup();
    if (!isSetup && $page.url.pathname !== "/setup") {
      await goto("/setup");
    }
  });
</script>

<svelte:head>
  <link rel="icon" href={favicon} />
</svelte:head>
<header class="w3-bar w3-top">
  <a class="w3-button" href="/">Home</a>
  <a class="w3-button" href="/setup">Setup</a>
</header>
<div class="w3-padding-16"></div>

{@render children?.()}
