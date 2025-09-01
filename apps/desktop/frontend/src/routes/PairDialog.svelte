<script lang="ts" module>
  import { device as dev } from "$lib/wails/wailsjs/go/models";
  import { writable } from "svelte/store";

  type PairDialogOptions = {
    info: dev.Info;
    accept: Function;
    decline: Function;
  };

  const pairDialogOptions = writable<PairDialogOptions | null>(null);
  export { pairDialogOptions };
  let options: PairDialogOptions | null = $state(null);

  pairDialogOptions.subscribe((value) => {
    options = value;
  });
</script>

<script lang="ts">
  import { Dialog } from "bits-ui";
  import { Check, X } from "lucide-svelte";
  import FoundDeviceTile from "$lib/components/tiles/FoundDeviceTile.svelte";

  function handleDecline() {
    if (!options) return;
    options.decline();
    pairDialogOptions.set(null);
  }

  function handleAccept() {
    if (!options) return;
    options.accept();
    pairDialogOptions.set(null);
  }
</script>

{#if options}
  <Dialog.Root
    open={!!options}
    onOpenChange={(open) => {
      if (!open) {
        pairDialogOptions.set(null);
      }
    }}
  >
    <Dialog.Portal>
      <Dialog.Overlay class="fixed inset-0 z-50 bg-black/30 backdrop-blur-sm" />
      <Dialog.Content
        class="card preset-filled-surface-200-800 border-surface-300-700 fixed left-1/2 top-1/2 z-50 w-full max-w-lg -translate-x-1/2 -translate-y-1/2  p-6 shadow-lg border-[1px]"
      >
        <div>
          <div class="flex flex-col space-y-4">
            <h3 class="h3">Pair with {options.info.Name}?</h3>
            <p class="text-surface-600-400">
              If you pair with <strong>{options.info.Name}</strong>, the device
              will be able to access services you have enabled. You can manage
              permissions in <a href="/settings" class="anchor">Settings</a>.
            </p>
            <FoundDeviceTile device={options.info} />
          </div>
          <div class="mt-6 flex justify-end space-x-2">
            <button
              type="button"
              class="btn preset-filled-error-300-700"
              onclick={handleDecline}
            >
              <X class="mr-2 h-4 w-4" />
              <span>Decline</span>
            </button>
            <button
              type="button"
              class="btn preset-filled-success-300-700"
              onclick={handleAccept}
            >
              <Check class="mr-2 h-4 w-4" />
              <span>Accept</span>
            </button>
          </div>
          <Dialog.Close
            class="absolute right-4 top-4 rounded-full p-1 transition-colors hover:bg-surface-300-700"
          >
            <X class="h-6 w-6" />
            <span class="sr-only">Close</span>
          </Dialog.Close>
        </div>
      </Dialog.Content>
    </Dialog.Portal>
  </Dialog.Root>
{/if}
