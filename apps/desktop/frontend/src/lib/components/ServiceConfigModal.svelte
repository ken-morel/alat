<script lang="ts">
  import { selectedDeviceForPairing } from "$lib/state";
  import { GetServices } from "$lib/wailsjs/go/app/App";
  import { service } from "$lib/wailsjs/go/models";
  import { onMount } from "svelte";

  const SERVICES: {
    [key: string]: string;
  } = {
    rcfile: "Receive files",
  };

  let dialog: HTMLDialogElement;
  let services: service.Service[] = $state([]);

  selectedDeviceForPairing.subscribe((device) => {
    if (device && dialog) {
      dialog.showModal();
    } else if (dialog) {
      dialog.close();
    }
  });

  function closeDialog() {
    selectedDeviceForPairing.set(null);
  }

  function handleKeydown(event: KeyboardEvent) {
    if (event.key === "Escape") {
      closeDialog();
    }
  }
  function acceptPair(e: SubmitEvent) {
    e.preventDefault();
  }
  onMount(() => {
    GetServices().then((appServices: service.Service[]) => {
      services = appServices;
    });
  });
</script>

<svelte:window on:keydown={handleKeydown} />

<dialog bind:this={dialog} onclose={closeDialog} onclick={() => dialog.close()}>
  {#if $selectedDeviceForPairing}
    <div class="modal-content">
      <h3>Configure Services for {$selectedDeviceForPairing.Name}</h3>
      <p>Device Code: {$selectedDeviceForPairing.Code}</p>

      <form method="dialog" onsubmit={acceptPair}>
        <div class="form-group">
          {#each services as service, idx}
            <label>
              <input type="checkbox" />
              <span>{SERVICES[service.Name]}</span>
            </label>
          {/each}
        </div>

        <div class="modal-actions">
          <button type="button" class="btn" onclick={closeDialog}>
            Cancel
          </button>
          <button type="submit" class="btn btn-primary"> Connect </button>
        </div>
      </form>
    </div>
  {/if}
</dialog>

<style lang="sass">
@use "$lib/styles/theme"

dialog
  width: 450px
  border: 1px solid theme.$primary-d2
  background-color: theme.$background
  color: theme.$text-primary
  padding: 0

  &::backdrop
    background: rgba(0, 0, 0, 0.7)

.modal-content
  padding: 1.5rem

h3
  margin-top: 0
  color: theme.$primary-l3

p
  color: theme.$text-secondary
  margin-bottom: 2rem

.form-group
  margin-bottom: 1.5rem
  display: flex
  align-items: center

  label
    display: flex
    align-items: center
    cursor: pointer
    font-size: 1rem

  input[type="checkbox"]
    margin-right: 0.75rem
    width: 18px
    height: 18px
    accent-color: theme.$primary

.modal-actions
  margin-top: 2rem
  display: flex
  justify-content: flex-end
  gap: 1rem

.btn
  background-color: theme.$primary-d2
  color: theme.$text-primary
  border: 1px solid theme.$primary-d1
  padding: 0.5rem 1.5rem
  cursor: pointer
  transition: background-color 0.2s

  &:hover
    background-color: theme.$primary-d1

.btn-primary
  background-color: theme.$primary
  border-color: theme.$primary-d1

  &:hover
    background-color: theme.$primary-l1

</style>
