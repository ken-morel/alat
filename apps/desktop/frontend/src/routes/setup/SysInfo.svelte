<script lang="ts">
  import { sysinfo } from "$lib/wailsjs/go/models";
  import { get, type Writable } from "svelte/store";

  let { cfg }: { cfg: Writable<sysinfo.ServiceConfig> } = $props();
  let start = get(cfg);

  let enabled: boolean = $state(start.Enabled);
  $effect(() => {
    cfg.set(
      new sysinfo.ServiceConfig({
        Enabled: enabled,
      }),
    );
    console.log("updating conf", get(cfg));
  });
  // enabled, dest, file max size
</script>

<div class="sysinfo-container">
  <h3>Share system stats</h3>
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
        When enabled, connected devices can read this device's battery, memory
        usage and more. Currently
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
</div>

<style lang="sass">
@use "$lib/styles/theme"
div.sysinfo-container
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
</style>
