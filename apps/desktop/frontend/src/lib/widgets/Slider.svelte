<script lang="ts">
  import type { Snippet } from "svelte";
  import { Slider } from "melt/components";
  import Tooltip from "./Tooltip.svelte";

  let {
    tooltip: tooltipSnippet,
    subtext,
    value = $bindable(),
    max,
    min,
    step,
    width = "300px",
  }: {
    tooltip?: Snippet;
    subtext?: Snippet;
    value: number;
    max?: number;
    min?: number;
    step?: number;
    width?: string;
  } = $props();
</script>

{#snippet slider()}
  <Slider bind:value {max} {min} {step}>
    {#snippet children(slider)}
      <div class="slider" {...slider.root}>
        <div class="track">
          <div class="range"></div>
          <div {...slider.thumb}></div>
        </div>
        {@render subtext?.()}
      </div>
    {/snippet}
  </Slider>
{/snippet}

<div style="--width: {width};">
  {#if tooltipSnippet}
    <Tooltip tooltip={tooltipSnippet}>
      {@render slider()}
    </Tooltip>
  {:else}
    {@render slider()}
  {/if}
</div>

<style lang="sass">
      .slider 
        width: var(--width)
        height: 40px
        margin: 0 auto
        padding-block: 16px // padding to increase touch area
        *
          transition: 0.1s
        

        .track 
          background: var(--color-surface-300-700)
          height: 100%
          position: relative;
          border-radius: 5px

        .range
          position: absolute
          background: var(--color-secondary-200)
          inset: 0
          right: var(--percentage-inv)
          border-radius: 5px

        [data-melt-slider-thumb]
          position: absolute
          background: var(--color-tertiary-300)
          left: var(--percentage)
          top: 50%
          transform: translate(-50%, -50%)
          width: 20px
          height: 20px
          border-radius: 5px
</style>
