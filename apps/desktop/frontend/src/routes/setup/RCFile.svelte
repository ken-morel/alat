<script lang="ts">
  import { rcfile } from "$lib/wailsjs/go/models";
  import { get, type Writable } from "svelte/store";
  import { AskDirectory } from "$lib/wailsjs/go/app/App";
  import { slide } from "svelte/transition";
  const SIZES = {
    KB: 1024,
    MB: 1024 * 1024,
    GB: 1024 * 1024 * 1024,
  };

  let { cfg }: { cfg: Writable<rcfile.ServiceConfig> } = $props();
  let start = get(cfg);

  let unit: "MB" | "KB" | "GB" = $state("KB");
  let enabled: boolean = $state(start.Enabled);
  let maxSize: number = $state(start.FileMaxSize / 1024);
  let destination: string = $state(start.Destination);
  function choosedir() {
    AskDirectory().then((dirname) => (destination = dirname));
  }
  $effect(() => {
    cfg.set(
      new rcfile.ServiceConfig({
        Enabled: enabled,
        FileMaxSize: SIZES[unit] * maxSize,
        Destination: destination,
      }),
    );
    console.log("updating conf", get(cfg));
  });
  // enabled, dest, file max size
</script>

<div class="rcfile-container">
  <h3>File receive</h3>
  <div class="enabling">
    <div class="check">
      <button
        class={["w3-button", enabled && "enabled"]}
        aria-label="Status"
        type="button"
        onclick={() => (enabled = !enabled)}
      ></button>
      {#if enabled}
        Enabled
      {:else}
        Disabled
      {/if}
    </div>
    <div class="doc">
      <p>
        File receive allows connected devices to send files to this device,
        disabling this will disable the feature globally. Currently
        <code style:color|important={enabled ? "#338844" : "#aa3333"}>
          {#if enabled}
            enabled
          {:else}
            disabled
          {/if}
        </code>
      </p>
    </div>
  </div>
  {#if enabled}
    <div transition:slide class="contents show">
      <div class="dest">
        <label for="rcfile-dest">Destination</label>
        <div>
          <input
            type="text"
            class="w3-input"
            bind:value={destination}
            id="rcfile-dest"
          />
          <button type="button" class="w3-button" onclick={choosedir}
            >Select</button
          >
        </div>
        <p>The folder inwhich received files will be saved</p>
      </div>
      <summary class="maxsize">
        <label for="maxsize">File max size</label>
        <div>
          <input
            type="number"
            class="w3-input"
            accept="[0-9]"
            id="rcfile-dest"
            bind:value={maxSize}
          />
          <div>
            {#each ["KB", "MB", "GB"] as thisUnit (thisUnit)}
              <button
                class="w3-button w3-bar-item {unit == thisUnit && 'selected'}"
                type="button"
                onclick={() => {
                  unit = thisUnit as "MB" | "KB" | "GB";
                }}
              >
                {thisUnit}
              </button>
            {/each}
          </div>
        </div>
        <p>Choose a maximum size for received files, leave empty to disable.</p>
      </summary>
    </div>
  {/if}
</div>

<style lang="sass">
@use "$lib/styles/theme"
div.rcfile-container
  transition: 0.4s
  h3
    text-align: center
  div.enabling
    display: flex
    div.check
      flex-grow: 0
      button
        width: 50px
        height: 50px
        margin: 5px
        border: 5px theme.$primary-d2 solid
        background-color: transparent  !important
        &.enabled
          background-color: theme.$primary-d1 !important
    div.doc
      flex-grow: 1
      border-left: 1px theme.$primary-d3 solid
      padding-left: 10px
      margin-bottom: 20px
  div.contents
    border-top: 1px theme.$primary-d3 solid
    padding-top: 20px
    height: 0
    display: none
    overflow: hidden
    transition: 0.4s
    &.show
      display: block
      height: 100% // fit-content
    div.dest,
    summary.maxsize
      div
        display: flex
        input
          background-color: theme.$primary-d3
          color: theme.$text-primary
          transition: border-color 0.1s
          &:invalid
            border-color: red
      p
        font-size: small
        text-align: right
    summary.maxsize
      div
        div
          button.selected
            background-color: theme.$primary-d2
</style>
