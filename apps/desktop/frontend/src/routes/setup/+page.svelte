<script lang="ts">
  import { onMount } from "svelte";
  import { GetConfig, SaveConfig } from "$lib/wailsjs/go/app/App";
  import { config } from "$lib/wailsjs/go/models";

  let cfg: config.Config | null = null;
  let deviceColorHex = "";

  onMount(async () => {
    const result = await GetConfig();
    cfg = result;
    if (cfg) {
      deviceColorHex = rgbaToHex(cfg.deviceColor);
    }
  });

  async function save() {
    if (cfg) {
      cfg.deviceColor = hexToRgba(deviceColorHex);
      await SaveConfig(cfg);
    }
  }

  function rgbaToHex(rgba: {
    r: number;
    g: number;
    b: number;
    a: number;
  }): string {
    const r = rgba.r.toString(16).padStart(2, "0");
    const g = rgba.g.toString(16).padStart(2, "0");
    const b = rgba.b.toString(16).padStart(2, "0");
    return `#${r}${g}${b}`;
  }

  function hexToRgba(hex: string): {
    r: number;
    g: number;
    b: number;
    a: number;
  } {
    const result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex);
    return result
      ? {
          r: parseInt(result[1], 16),
          g: parseInt(result[2], 16),
          b: parseInt(result[3], 16),
          a: 255,
        }
      : { r: 0, g: 0, b: 0, a: 255 };
  }
</script>

<div class="container">
  <h1>Setup</h1>

  {#if cfg}
    <form on:submit|preventDefault={save}>
      <div class="form-group">
        <label for="deviceName">Device Name</label>
        <input id="deviceName" type="text" bind:value={cfg.deviceName} />
      </div>

      <div class="form-group">
        <label for="deviceColor">Device Color</label>
        <input id="deviceColor" type="color" bind:value={deviceColorHex} />
      </div>

      <div class="form-group">
        <label for="language">Language</label>
        <select id="language" bind:value={cfg.language}>
          <option value="en-cm">English</option>
          <option value="fr-cm">French</option>
        </select>
      </div>

      <div class="form-group">
        <label for="autoStart">
          <input id="autoStart" type="checkbox" bind:checked={cfg.autoStart} />
          Auto Start
        </label>
      </div>

      <div class="form-group">
        <label for="theme">Theme</label>
        <select id="theme" bind:value={cfg.theme}>
          <option value="light">Light</option>
          <option value="dark">Dark</option>
        </select>
      </div>

      <button type="submit">Save</button>
    </form>
  {:else}
    <p>Loading...</p>
  {/if}
</div>
