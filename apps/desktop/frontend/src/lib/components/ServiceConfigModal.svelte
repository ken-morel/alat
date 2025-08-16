<script lang="ts">
  import { selectedDeviceForPairing } from "$lib/state";
  import { GetServices, RequestPair } from "$lib/wailsjs/go/app/App";
  import { service } from "$lib/wailsjs/go/models";
  import { onMount } from "svelte";

  const SERVICES: {
    [key: string]: string;
  } = {
    rcfile: "Receive files",
  };

  let dialog: HTMLDialogElement;
  let services: service.Service[] = $state([]);
  let submitting = $state(false);
  let error = $state<string | null>(null);

  selectedDeviceForPairing.subscribe((device) => {
    if (device && dialog) {
      console.log(device);
      dialog.showModal();
      // Reset state when modal opens
      submitting = false;
      error = null;
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

  async function acceptPair(event: SubmitEvent) {
    event.preventDefault();
    if (!$selectedDeviceForPairing) return;

    submitting = true;
    error = null;

    try {
      await RequestPair($selectedDeviceForPairing, services);
      // On success, close the dialog
      closeDialog();
    } catch (err) {
      console.error(err);
      error = "Failed to send pair request. Please try again.";
    } finally {
      submitting = false;
    }
  }

  onMount(() => {
    GetServices().then((appServices: service.Service[]) => {
      // Initialize services, ensuring 'Enabled' is a boolean
      services = appServices.map((s) => {
        s.Enabled = s.Enabled || false;
        return s;
      });
    });
  });
</script>

<svelte:window on:keydown={handleKeydown} />

<dialog
  bind:this={dialog}
  onclose={closeDialog}
  onclick={(e) => e.target === dialog && closeDialog()}
>
  {#if $selectedDeviceForPairing}
    <div class="modal-content">
      <h3>Configure Services for {$selectedDeviceForPairing.Name}</h3>
      <p>
        Device Code: {$selectedDeviceForPairing.Code}<br />
        Address: <code>{$selectedDeviceForPairing.Address.Phrase}</code>
      </p>
      <ul>
        <h5>Supports</h5>
        {#each $selectedDeviceForPairing.Services as othersService}
          <li>{othersService.Name}</li>
        {/each}
      </ul>

      <form method="dialog" onsubmit={acceptPair}>
        <div class="form-group">
          {#each services as serviceItem}
            <label>
              <input type="checkbox" bind:checked={serviceItem.Enabled} />
              <span>{SERVICES[serviceItem.Name] || serviceItem.Name}</span>
            </label>
          {/each}
        </div>

        {#if error}
          <p class="error-message">{error}</p>
        {/if}

        <div class="modal-actions">
          <button
            type="button"
            class="btn"
            onclick={closeDialog}
            disabled={submitting}
          >
            Cancel
          </button>
          <button type="submit" class="btn btn-primary" disabled={submitting}>
            {#if submitting}
              Connecting...
            {:else}
              Connect
            {/if}
          </button>
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
ul
  margin-bottom: 2rem

.form-group
  margin-bottom: 1.5rem
  display: flex
  flex-direction: column
  align-items: flex-start
  gap: 1rem

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
  
  &:disabled
    opacity: 0.6
    cursor: not-allowed

.btn-primary
  background-color: theme.$primary
  border-color: theme.$primary-d1

  &:hover
    background-color: theme.$primary-l1

.error-message
  color: #f44336 // Standard error color
  margin-top: 1rem
  text-align: center

</style>
