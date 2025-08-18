<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import favicon from "$lib/assets/logo.svg";
  import "$lib/styles/app.sass";
  import { WasSetup } from "$lib/wailsjs/go/app/App";
  import { onMount } from "svelte";
  import { writable, type Writable } from "svelte/store";

  let { children } = $props();

  onMount(async () => {
    const isSetup = await WasSetup();
    if (!isSetup && $page.url.pathname !== "/setup") {
      await goto("/setup");
    }
  });
  type Dload = {
    file: string;
    percent: number;
  };
  let dloads: Dload[] = [];
  let downloadPercent: Writable<number | null> = writable(null);

  // @ts-ignore
  window.downloadCallback = async (file: string, percent: number) => {
    console.log(percent);
    let found: boolean = false;
    for (const dload of dloads) {
      if (dload.file == file) {
        dload.percent = percent;
        found = true;
        break;
      }
    }
    if (!found) {
      dloads.push({ file, percent });
    }
    updatePercent();
  };
  function updatePercent() {
    dloads = dloads.filter((d) => d.percent != 100);
    if (dloads.length == 0) {
      downloadPercent.set(null);
      return;
    }
    let sum: number = 0;
    for (let d of dloads) sum += d.percent;
    downloadPercent.set(sum / dloads.length);
  }
</script>

<svelte:head>
  <link rel="icon" href={favicon} />
</svelte:head>
<header class="w3-bar w3-top">
  <a class="w3-button w3-bar-item" href="/">Home</a>
  <a class="w3-button w3-bar-item" href="/setup">Setup</a>

  {#if $downloadPercent}
    <span class="w3-bar-item w3-right w3-small w3-opacity">
      Receiving files...
    </span>

    <span class="progress" style="width: {$downloadPercent}%;"></span>
  {/if}
</header>
<div class="w3-padding-32"></div>

{@render children?.()}

<style lang="sass">
header
  span.progress
    position: absolute
    bottom: 0
    left: 0
    height: 5px
    background-color: blue;
    transition: 0.1s
</style>
