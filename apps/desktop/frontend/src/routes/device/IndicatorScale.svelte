<script lang="ts">
  let { label, percent }: { label: string; percent: number } = $props();
  let idx = Math.random();
</script>

<div class="indicator-container">
  <progress value={percent} max="100" id="indicator-{idx}"></progress>
  <label for="indicator-{idx}">{label}</label>
</div>

<style lang="sass">
.indicator-container
</style>
