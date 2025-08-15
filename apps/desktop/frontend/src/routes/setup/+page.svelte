<script lang="ts">
  import { onMount } from "svelte";
  import { GetConfig, SaveConfig } from "$lib/wailsjs/go/app/App";
  import { config, options, rcfile } from "$lib/wailsjs/go/models";
  import Color from "$lib/color";
  import { goto } from "$app/navigation";
  import RcFile from "./RCFile.svelte";
  import { writable } from "svelte/store";

  let cfg: config.Config | null = $state(null);
  let deviceColorHex = $state("");
  let pstate: "loading" | "ready" = $state("loading");
  let rfconf = writable<rcfile.ServiceConfig>(new rcfile.ServiceConfig());
  rfconf.subscribe((val: rcfile.ServiceConfig) => {
    if (cfg) cfg.Services.RCFile = val;
  });

  onMount(async () => {
    await new Promise((r) => setTimeout(r, 200));
    cfg = await GetConfig();
    rfconf.set(cfg.Services.RCFile);
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
    if (pstate === "ready" && cfg) {
      const col = new Color(deviceColorHex);
      const { red, green, blue, opacity } = col;
      cfg.DeviceColor = new options.RGBA({
        r: red,
        g: green,
        b: blue,
        a: opacity * 255,
      });
      SaveConfig(cfg).then(() => {
        goto("/dashboard");
      });
    }
  }
</script>

<div class="setup-container">
  <h1 class="page-title">Device Setup</h1>

  <form
    class="setup-form"
    onsubmit={(e) => {
      e.preventDefault();
      save();
    }}
  >
    {#if pstate == "ready" && cfg}
      <section class="form-section">
        <article class="form-article device-name-color">
          <div class="form-group device-name">
            <label for="dname">Device name</label>
            <input
              id="dname"
              name="dname"
              bind:value={cfg.DeviceName}
              style="border-color: {deviceColorHex}"
            />
            <p class="hint">
              Right click or press <kbd>Win + Period</kbd> to enter emoji
            </p>
          </div>
          <div class="form-group device-color">
            <label for="dcolor">Color</label>
            <input
              type="color"
              id="dcolor"
              name="dcolor"
              bind:value={deviceColorHex}
            />
            <p class="hint">Choose a color</p>
          </div>
        </article>
        <article class="form-article">
          <label for="language">Language</label>
          <div class="language-options">
            <span class="selected-lang {cfg.Language}" style="border-color: {deviceColorHex}"></span>
            <button
              type="button"
              class="lang-button"
              onclick={() => {
                if (cfg) cfg.Language = "en-cm";
              }}
            >
              English
            </button>
            <button
              type="button"
              class="lang-button"
              onclick={() => {
                if (cfg) cfg.Language = "fr-cm";
              }}
            >
              French
            </button>
          </div>
        </article>
      </section>

      <section class="form-section">
        <h2 class="section-title">Services</h2>
        <p class="section-intro">
          Services are the features of your device other devices can have access.
          It could be file receiving, notification sending and more. It is
          important for them to be configured at setup, but every time a device
          will be connected you will need to select services you want it to
          access. For more information, refer to
          <a href="/setup">Services documentation</a>.
        </p>
      </section>

      <section class="form-section">
        <RcFile cfg={rfconf} />
      </section>
    {:else}
      <p class="loading-text">Loading...</p>
    {/if}

    <section class="form-section">
      <button class="save-button btn btn-primary">Save</button>
    </section>
  </form>
</div>

<style lang="sass">
  @use '$lib/styles/theme'

  .setup-container
    max-width: 700px
    margin: auto
    padding: 2rem 1.5rem

  .page-title
    text-align: center
    font-size: 2.5rem
    margin-bottom: 2rem

  .form-section
    margin-bottom: 2.5rem

  .form-article
    padding-bottom: 1.5rem
    border-bottom: 1px solid theme.$primary-d3
    margin-bottom: 1.5rem

  .device-name-color
    display: flex
    gap: 1.5rem

  .form-group
    label
      display: block
      font-size: 1.5rem
      margin-bottom: 0.5rem
      color: theme.$text-secondary
    
    input
      font-size: 1.25rem
      border-width: 2px

    .hint
      margin-top: 0.25rem
      font-size: 0.875rem
      text-align: right
      color: theme.$text-secondary

  .device-name
    flex-grow: 1

  .device-color
    input[type="color"]
      height: 53px
      width: 53px
      padding: 0
      margin: 0
      border-width: 2px

  .language-options
    position: relative
    display: flex

    .selected-lang
      position: absolute
      top: 0
      height: 100%
      width: 50%
      border: 3px solid theme.$primary
      transition: left 0.2s ease-in-out
      
      &.en-cm
        left: 0
      &.fr-cm
        left: 50%

    .lang-button
      width: 50%
      padding: 1rem
      font-size: 1.25rem
      background-color: theme.$primary-d3
      color: theme.$text-primary
      border: none
      cursor: pointer
      
      &:hover
        background-color: theme.$primary-d2

  .section-title
    text-align: center
    font-size: 2rem
    margin-bottom: 1rem

  .section-intro
    color: theme.$text-secondary
    text-align: center

  .loading-text
    font-size: 1.5rem
    text-align: center
    opacity: 0.7

  .save-button
    width: 100%
    font-size: 1.5rem
    padding: 1rem
</style>
