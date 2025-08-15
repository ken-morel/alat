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
  // device name, device color,language, autostart, theme
</script>

<h1 class="w3-center w3-xxlarge w3-padding-32">Device Setup</h1>

<form
  onsubmit={(e) => {
    e.preventDefault();
    save();
  }}
>
  {#if pstate == "ready" && cfg}
    <section class="app-config">
      <article class="dname">
        <div class="name">
          <label for="dname">Device name</label>
          <input
            id="dname"
            name="dname"
            bind:value={cfg.DeviceName}
            class="w3-input"
            style="border-color: {deviceColorHex}"
          />
          <p class="hint">
            Right click or press <kbd class="">Win + Period</kbd> to enter emoji
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
      </article>
      <article class="lang">
        <label for="language"> Language </label>
        <div class="options w3-container w3-row">
          <span class="selected {cfg.Language}"></span>
          <button
            type="button"
            class="option w3-input w3-button w3-large w3-col w3-half"
            onclick={() => {
              if (cfg) cfg.Language = "en-cm";
            }}
            style="border-color: {deviceColorHex}"
          >
            English
          </button>
          <button
            type="button"
            class="option w3-input w3-button w3-large w3-col w3-half"
            onclick={() => {
              if (cfg) cfg.Language = "fr-cm";
            }}
            style="border-color: {deviceColorHex}"
          >
            French
          </button>
        </div>
      </article>
    </section>
    <section>
      <h2 class="w3-center">Services</h2>
      <p>
        Services are the features of your device other devices can have access.
        It could be file receiving, notification sending and more. It is
        important for them to be configured at setup, but every time a device
        will be connected you will need to select services you want it to
        access. For more information, refer to
        <a href="/setup"> Services documentation </a>
      </p>
    </section>
    <section class="rcfile">
      <RcFile cfg={rfconf} />
    </section>
  {:else}
    <p class="loading w3-large w3-center w3-opacity">Loading...</p>
  {/if}

  <section><button class="save w3-button">Save</button></section>
</form>

<style lang="sass">
  @use '$lib/styles/theme'
  form
    button.save
      margin-top: 10px
      width: 100%
      font-size: x-large
      color: theme.$text-primary !important
      background-color: theme.$primary-d3 !important
  
      &:hover
        background-color: theme.$primary-d2 !important
    section
      background-color: theme.$background
      margin: auto
      max-width: 700px
      padding: 30px
      transition: 1s
      margin-top: 2px
    section.app-config
      article
        display: block
        border-bottom: 2px theme.$tertiary-d3 solid
        padding: 10px
        label
          font-size: xx-large
          color: theme.$text-primary
        input
          background-color: theme.$primary-d3
          color: theme.$text-primary
          font-size: x-large
          margin-bottom: 0!important  
        p.hint
          margin-top: 0!important
          text-align: right
          color: theme.$text-primary
      article.dname
        display: flex
        div.name
          flex-grow: 1
        div.dcolor
          text-align: left
          margin-left: 20px
          input
            height: 53px
            width: 53px
            padding: 0px !important
            margin: 0px !important
            display: block
      article.lang
        div.options
          position: relative
          span.selected
            position: absolute
            top: 0
            border: 3px theme.$primary solid
            height: 100%
            width: 50%
            transition: 0.2s
            &.en-cm
              left: 0
            &.fr-cm
              left: 50%
          button
            color: theme.$text-primary !important
            &:hover
              background-color: theme.$primary-d4 !important
</style>
