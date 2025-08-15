<script lang="ts">
  import { onMount } from "svelte";
  import { GetConfig, SaveConfig } from "$lib/wailsjs/go/app/App";
  import { config, options } from "$lib/wailsjs/go/models";
  import Color from "$lib/color";

  let cfg: config.Config | null = $state(null);
  let deviceColorHex = $state("");
  let pstate: "loading" | "ready" = $state("loading");

  onMount(async () => {
    await new Promise((r) => setTimeout(r, 2000));
    cfg = await GetConfig();
    if (cfg) {
      const { r, g, b, a } = cfg.DeviceColor;
      if (r + g + b > 10)
        deviceColorHex = Color.rgba(r, g, b, a / 255).toHexString();
    }

    if (deviceColorHex == "") {
      const colors = Object.values(Color.COLORS);
      const col = colors.at(Math.round(Math.random() * colors.length));
      if (col) {
        deviceColorHex = "#" + col;
      }
    }
    pstate = "ready";
  });

  async function save() {
    if (cfg) {
      const col = new Color(deviceColorHex);
      const { red, green, blue, opacity } = col;
      cfg.DeviceColor = new options.RGBA({
        r: red,
        g: green,
        b: blue,
        a: opacity * 255,
      });
      await SaveConfig(cfg);
    }
  }

  // device name, device color,language, autostart, theme
</script>

<h1 class="w3-center w3-xxlarge w3-padding-32">Device Setup</h1>

{#if pstate == "ready" && cfg}
  <form
    onsubmit={(e) => {
      e.preventDefault();
      save();
    }}
  >
    <section class="dname">
      <div class="name">
        <label for="dname">Device name</label>
        <input
          id="dname"
          name="dname"
          bind:value={cfg.DeviceName}
          class="w3-input"
        />
        <p class="hint">
          Right click or press <kbd>Win + Period</kbd> to enter emoji
        </p>
      </div>
      <div class="dcolor">
        <label for="dcolor">Color</label>
        <input
          type="color"
          id="dcolor"
          name="dcolor"
          class="w3-input"
          bind:value={deviceColorHex}
        />
        <p class="hint">Choose a color</p>
      </div>
    </section>
  </form>
{:else}
  <p>Loading...</p>
{/if}

<style lang="sass">
  @use '$lib/styles/theme'
  form
    background-color: theme.$background-dark
    margin: auto
    max-width: 700px
    padding: 30px
    section
      display: block
      border-bottom: 2px theme.$tertiary-d1 solid
      padding: 10px
      label
        font-size: xx-large
        color: theme.$text-primary
      input
        background-color: theme.$secondary-d3
        border-color: theme.$primary-d2
        color: theme.$text-secondary
        font-size: x-large
        margin-bottom: 0!important  
      p.hint
        margin-top: 0!important
        text-align: right
        color: theme.$text-secondary
    section.dname
      display: flex
      div.name
        flex-grow: 1
      div.dcolor
        text-align: left
        margin-left: 20px
        input
          height: 52px
          width: 52px
          padding: 0px
          margin: 0px
          display: block
</style>
