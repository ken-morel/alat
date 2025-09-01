<script lang="ts">
  import { page } from "$app/state";
</script>

<div class="w-full h-full place-items-center grid">
  <div class="card preset-filled-surface-100-900 p-8">
    <h3 class="h3">Error</h3>
    <p class="p text-surface-400-600">
      An error occured: {page.error?.message}
    </p>
  </div>
</div>
