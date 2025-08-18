<script lang="ts">
  import { goto } from "$app/navigation";
  import { WasSetup } from "$lib/wailsjs/go/app/App";

  WasSetup().then((setup) => {
    if (setup) goto("/dashboard");
    else goto("/setup");
  });
</script>

<p class="w3-padding-64 w3-center w3-xxlarge">Redirecting you...</p>
