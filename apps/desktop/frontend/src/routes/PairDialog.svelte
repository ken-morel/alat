<script lang="ts" module>
  import { device as dev } from "$lib/wails/wailsjs/go/models";
  import { writable } from "svelte/store";
  import { fade } from "svelte/transition";
  import { X } from "phosphor-svelte";

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

{#if options}
  <dialog
    open={true}
    class="fixed backdrop-blur-xl place-items-center grid w-full bg-primary-100-900/70"
  >
    <div
      class="top-10 right-10 p-1 absolute backdrop-blur-sm rounded-full bg-error-400-600 opacity-50"
    >
      <X onclick={() => pairDialogOptions.set(null)} size={40} />
    </div>
    <div
      transition:fade
      class="card preset-filled-surface-200-800 border-[1px] border-surface-200-800 p-4"
    >
      <h3 class="h3">Do you want to pair with {options.info.Name}?</h3>
      <p class=" text-surface-600-400">
        If you pair with <span class="">{options.info.Name}</span> The device
        will be able to access the services you enabled in the
        <a class="" href="/setup">Settings</a>
      </p>
      <div class="flex w-full">
        <button class="btn preset-filled-success-300-700">Accept</button>
        <button class="btn preset-filled-error-300-700">Decline</button>
      </div>
    </div>
  </dialog>
{/if}
