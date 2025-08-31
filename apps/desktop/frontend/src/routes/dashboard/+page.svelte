<script lang="ts">
  let { data } = $props();
</script>

<div class="h-full w-full grid place-items-center">
  <div
    class="card preset-filled-surface-100-900 border-[1px] border-surface-200-800 w-full max-w-lg p-8"
  >
    <header><h3 class="h3">Your devices</h3></header>
    <article class="mt-4">
      {#each data.paired as pairedDevice}
        <div class="p-4">A paired device</div>
      {:else}
        <div>
          <p class="p text-surface-400-600">
            No paired device, ... for now <a class="a" href="/pair"
              >Pair a device</a
            >
          </p>
        </div>
      {/each}
    </article>
    <footer class="mt-4">
      <a href="/pair" class="btn preset-filled-secondary-600-400 w-full"
        >Pair a new device</a
      >
    </footer>
  </div>
</div>
