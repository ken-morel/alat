<script lang="ts">
  import { goto } from "$app/navigation";
  import { selectedPairedDevice } from "$lib/state";
  $: {
    if ($selectedPairedDevice == null) goto("/");
  }
</script>

<div>The device page</div>
